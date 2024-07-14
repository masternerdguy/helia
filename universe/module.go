package universe

import (
	"fmt"
	"helia/listener/models"
	"helia/physics"
	"helia/shared"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

func (m *FittedSlot) activateAsGunTurret() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else if *m.TargetType == tgtReg.Asteroid {
		// verify this can mine
		canMineOre, _ := m.ItemTypeMeta.GetBool("can_mine_ore")
		canMineIce, _ := m.ItemTypeMeta.GetBool("can_mine_ice")

		if !canMineOre && !canMineIce {
			return false
		}

		// find asteroid
		tgt, f := m.shipMountedOn.CurrentSystem.asteroids[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// check if ammunition required to fire (note: NPCs do not require ammunition - i want to fix this eventually)
	ammoTypeRaw, found := m.ItemMeta.GetString("ammunition_type")
	var ammoItem *Item = nil

	if found && !m.shipMountedOn.IsNPC {
		// parse type id
		typeID, _ := uuid.Parse(ammoTypeRaw)

		// verify there is enough ammunition to fire
		x := m.shipMountedOn.FindFirstAvailableItemOfTypeInCargo(typeID, true)

		if x == nil {
			return false
		}

		// store item to take from
		ammoItem = x
	}

	// get damage values
	shieldDmg, _ := m.ItemMeta.GetFloat64("shield_damage")
	armorDmg, _ := m.ItemMeta.GetFloat64("armor_damage")
	hullDmg, _ := m.ItemMeta.GetFloat64("hull_damage")

	// apply usage experience modifier
	shieldDmg *= m.usageExperienceModifier
	armorDmg *= m.usageExperienceModifier
	hullDmg *= m.usageExperienceModifier

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// adjust damage based on tracking
	shieldDmg *= trackingRatio
	armorDmg *= trackingRatio
	hullDmg *= trackingRatio

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% damage (or ore pull if asteroid)

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// damage dealt (or ore / ice pulled if asteroid) is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)

			shieldDmg *= rangeRatio
			armorDmg *= rangeRatio
			hullDmg *= rangeRatio
		} else if falloff == "reverse_linear" {
			// damage dealt (or ore / ice pulled if asteroid) is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			shieldDmg *= rangeRatio
			armorDmg *= rangeRatio
			hullDmg *= rangeRatio
		}
	}

	// reduce ammunition count if needed
	if ammoItem != nil {
		ammoItem.Quantity--
		ammoItem.CoreDirty = true

		// escalate for saving
		m.shipMountedOn.CurrentSystem.ChangedQuantityItems = append(m.shipMountedOn.CurrentSystem.ChangedQuantityItems, ammoItem)
	}

	// apply damage (or ore / ice pulled if asteroid) to target
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)
		c.dealDamage(shieldDmg, armorDmg, hullDmg, m.shipMountedOn.ReputationSheet, m)
	} else if *m.TargetType == tgtReg.Station {
		c := tgtI.(*Station)
		c.DealDamage(shieldDmg, armorDmg, hullDmg)
	} else if *m.TargetType == tgtReg.Asteroid {
		// target is an asteroid - what can this module mine?
		canMineOre, _ := m.ItemTypeMeta.GetBool("can_mine_ore")
		canMineIce, _ := m.ItemTypeMeta.GetBool("can_mine_ice")

		// get ore / ice type and volume
		c := tgtI.(*Asteroid)

		// can this module mine this type?
		canMine := false

		if c.ItemFamilyID == "ore" && canMineOre {
			canMine = true
		} else if c.ItemFamilyID == "ice" && canMineIce {
			canMine = true
		}

		if canMine {
			// get mining volume
			oreMiningVolume, _ := m.ItemMeta.GetFloat64("ore_mining_volume")
			iceMiningVolume, _ := m.ItemMeta.GetFloat64("ice_mining_volume")

			miningVolume := 0.0

			if c.ItemFamilyID == "ore" && canMineOre {
				miningVolume = oreMiningVolume
			} else if c.ItemFamilyID == "ice" && canMineIce {
				miningVolume = iceMiningVolume
			}

			// get type and volume of ore / ice being collected
			unitType := c.ItemTypeID
			unitVol, _ := c.ItemTypeMeta.GetFloat64("volume")

			// get available space in cargo hold
			free := m.shipMountedOn.GetRealCargoBayVolume(false) - m.shipMountedOn.TotalCargoBayVolumeUsed(false)

			// calculate effective ore / ice volume pulled
			pulled := miningVolume * c.Yield * rangeRatio

			// apply usage experience modifier
			pulled *= m.usageExperienceModifier

			// make sure there is sufficient room to deposit the ore / ice
			if free-pulled >= 0 {
				found := false

				// quantity to be placed in cargo bay
				q := int((miningVolume * c.Yield) / unitVol)

				if q <= 0 && m.shipMountedOn.IsNPC {
					// raise fault
					m.shipMountedOn.aiNoOrePulledFault = true
				}

				// is there already packaged ore / ice of this type in the hold?
				for idx := range m.shipMountedOn.CargoBay.Items {
					itm := m.shipMountedOn.CargoBay.Items[idx]

					if itm.ItemTypeID == unitType && itm.IsPackaged && !itm.CoreDirty {
						// increase the size of this stack
						itm.Quantity += q

						// escalate for saving
						m.shipMountedOn.CurrentSystem.ChangedQuantityItems = append(m.shipMountedOn.CurrentSystem.ChangedQuantityItems, itm)

						// mark as found
						found = true
						break
					}
				}

				if !found && q > 0 {
					// create a new stack of ore / ice
					nid := uuid.New()

					newItem := Item{
						ID:            nid,
						ItemTypeID:    unitType,
						Meta:          c.ItemTypeMeta,
						Created:       time.Now(),
						CreatedBy:     &m.shipMountedOn.UserID,
						CreatedReason: fmt.Sprintf("Mined %v", c.ItemFamilyID),
						ContainerID:   m.shipMountedOn.CargoBayContainerID,
						Quantity:      q,
						IsPackaged:    true,
						Lock:          sync.Mutex{}, ItemTypeName: c.ItemTypeName,
						ItemFamilyID:   c.ItemFamilyID,
						ItemFamilyName: c.ItemFamilyName,
						ItemTypeMeta:   c.ItemTypeMeta,
						CoreDirty:      true,
					}

					// escalate to core for saving in db
					m.shipMountedOn.CurrentSystem.NewItems = append(m.shipMountedOn.CurrentSystem.NewItems, &newItem)

					// add new item to cargo hold
					m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, &newItem)
				}

				// log mine to console if player
				if !m.shipMountedOn.IsNPC {
					bm := 0

					if m.shipMountedOn.BehaviourMode != nil {
						bm = *m.shipMountedOn.BehaviourMode
					}

					shared.TeeLog(
						fmt.Sprintf(
							"[%v] %v (%v::%v) mined %v %v from %v",
							m.shipMountedOn.CurrentSystem.SystemName,
							m.shipMountedOn.CharacterName,
							m.shipMountedOn.Texture,
							bm,
							q,
							c.ItemTypeName,
							c.Name,
						),
					)
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsMissileLauncher() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// check if ammunition required to fire (note: NPCs do not require ammunition - i want to fix this eventually)
	ammoTypeRaw, found := m.ItemMeta.GetString("ammunition_type")
	var ammoItem *Item = nil

	if found && !m.shipMountedOn.IsNPC {
		// parse type id
		typeID, _ := uuid.Parse(ammoTypeRaw)

		// verify there is enough ammunition to fire
		x := m.shipMountedOn.FindFirstAvailableItemOfTypeInCargo(typeID, true)

		if x == nil {
			return false
		}

		// store item to take from
		ammoItem = x
	}

	// reduce ammunition count if needed
	if ammoItem != nil {
		ammoItem.Quantity--
		ammoItem.CoreDirty = true

		// escalate for saving
		m.shipMountedOn.CurrentSystem.ChangedQuantityItems = append(m.shipMountedOn.CurrentSystem.ChangedQuantityItems, ammoItem)
	}

	// build and hook missile projectile
	missileGfxEffect, _ := m.ItemTypeMeta.GetString("missile_gfx_effect")
	missileRadius, _ := m.ItemTypeMeta.GetFloat64("missile_radius")
	faultTolerance, _ := m.ItemMeta.GetFloat64("fault_tolerance")
	flightTime, _ := m.ItemMeta.GetFloat64("flight_time")
	maxVelocity := (modRange / flightTime)

	// accumulate guidance drift
	rawDrift := 0.0

	for _, mo := range m.shipMountedOn.TemporaryModifiers {
		if mo.Attribute == "guidance_drift" {
			rawDrift += mo.Raw
		}
	}

	// apply guidance drift
	if rawDrift > 0 {
		// scale for tolerence
		sDrift := rawDrift / 100.0

		// get tolerance ratio
		tRat := sDrift / faultTolerance

		// get scale factor from tolerance ratio
		sFac := (rawDrift * tRat) / 100.0

		// clamp to 0%
		if sFac < 0 {
			sFac = 0
		}

		// clamp to 100%
		if sFac > 1 {
			sFac = 1
		}

		// convert factor
		cFac := 1.0 - sFac

		// apply factor
		faultTolerance *= cFac
		flightTime *= cFac
	}

	// apply usage experience modifiers
	flightTime *= m.usageExperienceModifier
	maxVelocity *= m.usageExperienceModifier

	stubID := uuid.New()
	flightTicks := int((flightTime * 1000) / Heartbeat)

	// determine launcher hardpoint location
	hpX := [...]float64{
		0,
		0,
	}

	if m.SlotIndex != nil {
		rack := m.Rack
		idx := *m.SlotIndex

		if rack == "A" {
			hpX = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
		}
	}

	sRad := physics.ToRadians(m.shipMountedOn.Theta)
	hRad := physics.ToRadians(hpX[1])

	cRad := sRad + hRad
	lx := math.Cos(cRad) * hpX[0]
	ly := math.Sin(cRad) * hpX[0]

	// store stub
	stub := Missile{
		ID:             stubID,
		PosX:           m.shipMountedOn.PosX + lx,
		PosY:           m.shipMountedOn.PosY + ly,
		Texture:        missileGfxEffect,
		Module:         m,
		TargetID:       *m.TargetID,
		TargetType:     *m.TargetType,
		Radius:         missileRadius,
		FiredByID:      m.shipMountedOn.ID,
		TicksRemaining: flightTicks,
		MaxVelocity:    maxVelocity,
		FaultTolerance: faultTolerance,
	}

	m.shipMountedOn.CurrentSystem.missiles[stub.ID.String()] = &stub

	// module activates!
	return true
}

func (m *FittedSlot) activateAsShieldBooster() bool {
	// get shield boost amount
	shieldBoost, _ := m.ItemMeta.GetFloat64("shield_boost_amount")

	// apply usage experience modifier
	shieldBoost *= m.usageExperienceModifier

	// apply boost to mounting ship
	m.shipMountedOn.dealDamage(-shieldBoost, 0, 0, nil, nil)

	// include visual effect if present
	activationPGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		tgtReg := models.SharedTargetTypeRegistry

		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationPGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsEngineOvercharger() bool {
	// get activation energy and duration (same as cooldown for engine overchargers)
	activationEnergy, _ := m.ItemMeta.GetFloat64("activation_energy")
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")

	// get ship mass
	mx := m.shipMountedOn.GetRealMass()

	// calculate engine boost amount
	dA := (activationEnergy / mx) * 10

	// calculate effect duration in ticks
	dT := (cooldown * 1000) / Heartbeat

	// apply usage experience modifier
	dT *= m.usageExperienceModifier

	// add temporary modifier
	modifier := TemporaryShipModifier{
		Attribute:      "accel",
		Percentage:     dA,
		RemainingTicks: int(dT),
	}

	m.shipMountedOn.TemporaryModifiers = append(m.shipMountedOn.TemporaryModifiers, modifier)

	// module activates!
	return true
}

func (m *FittedSlot) activateAsActiveRadiator() bool {
	// get activation energy and duration (same as cooldown for active radiators)
	activationEnergy, _ := m.ItemMeta.GetFloat64("activation_energy")
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")

	// get ship heat capacity
	mx := m.shipMountedOn.GetRealMaxHeat(false)

	// calculate heat sink amount
	dA := (activationEnergy / mx) * 35

	// calculate effect duration in ticks
	dT := (cooldown * 1000) / Heartbeat

	// apply usage experience modifier
	dT *= m.usageExperienceModifier

	// add temporary modifier
	modifier := TemporaryShipModifier{
		Attribute:      "heat_sink",
		Percentage:     dA,
		RemainingTicks: int(dT),
	}

	m.shipMountedOn.TemporaryModifiers = append(m.shipMountedOn.TemporaryModifiers, modifier)

	// module activates!
	return true
}

func (m *FittedSlot) activateAsAetherDragger() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI interface{}

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// get drag multiplier
	dragMul, _ := m.ItemMeta.GetFloat64("drag_multiplier")

	// apply usage experience modifier
	dragMul *= m.usageExperienceModifier

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")
	rangeRatio := 1.0 // default to 100%

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// drag increase is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)
			dragMul *= rangeRatio
		} else if falloff == "reverse_linear" {
			// drag increase is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			dragMul *= rangeRatio
		}
	}

	// apply drag to target
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)

		// calculate effect duration in ticks
		cooldown, _ := m.ItemMeta.GetFloat64("cooldown")
		dT := (cooldown * 1000) / Heartbeat

		// add temporary modifier to target
		modifier := TemporaryShipModifier{
			Attribute:      "drag",
			Percentage:     dragMul,
			RemainingTicks: int(dT),
		}

		c.TemporaryModifiers = append(c.TemporaryModifiers, modifier)

		// update aggression tables
		c.updateAggressionTables(
			0,
			0,
			0,
			m.shipMountedOn.ReputationSheet,
			m,
		)
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsUtilityMiner() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Asteroid {
		// find asteroid
		tgt, f := m.shipMountedOn.CurrentSystem.asteroids[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% efficiency

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// damage dealt (or ore / ice pulled if asteroid) is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)
		} else if falloff == "reverse_linear" {
			// damage dealt (or ore / ice pulled if asteroid) is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}
		}
	}

	// apply damage (or ore / ice pulled if asteroid) to target
	if *m.TargetType == tgtReg.Asteroid {
		// target is an asteroid - what can this module mine?
		canMineOre, _ := m.ItemTypeMeta.GetBool("can_mine_ore")
		canMineIce, _ := m.ItemTypeMeta.GetBool("can_mine_ice")

		// get ore / ice type and volume
		c := tgtI.(*Asteroid)

		// can this module mine this type?
		canMine := false

		if c.ItemFamilyID == "ore" && canMineOre {
			canMine = true
		} else if c.ItemFamilyID == "ice" && canMineIce {
			canMine = true
		}

		if canMine {
			// get mining volume
			oreMiningVolume, _ := m.ItemMeta.GetFloat64("ore_mining_volume")
			iceMiningVolume, _ := m.ItemMeta.GetFloat64("ice_mining_volume")

			miningVolume := 0.0

			if c.ItemFamilyID == "ore" && canMineOre {
				miningVolume = oreMiningVolume
			} else if c.ItemFamilyID == "ice" && canMineIce {
				miningVolume = iceMiningVolume
			}

			// get type and volume of ore / ice being collected
			unitType := c.ItemTypeID
			unitVol, _ := c.ItemTypeMeta.GetFloat64("volume")

			// get available space in cargo hold
			free := m.shipMountedOn.GetRealCargoBayVolume(false) - m.shipMountedOn.TotalCargoBayVolumeUsed(false)

			// calculate effective ore / ice volume pulled
			pulled := miningVolume * c.Yield * rangeRatio * trackingRatio

			// apply usage experience modifier
			pulled *= m.usageExperienceModifier

			// make sure there is sufficient room to deposit the ore / ice
			if free-pulled >= 0 {
				found := false

				// quantity to be placed in cargo bay
				q := int((miningVolume * c.Yield) / unitVol)

				if q <= 0 && m.shipMountedOn.IsNPC {
					// raise fault
					m.shipMountedOn.aiNoOrePulledFault = true
				}

				// is there already packaged ore / ice of this type in the hold?
				for idx := range m.shipMountedOn.CargoBay.Items {
					itm := m.shipMountedOn.CargoBay.Items[idx]

					if itm.ItemTypeID == unitType && itm.IsPackaged && !itm.CoreDirty {
						// increase the size of this stack
						itm.Quantity += q

						// escalate for saving
						m.shipMountedOn.CurrentSystem.ChangedQuantityItems = append(m.shipMountedOn.CurrentSystem.ChangedQuantityItems, itm)

						// mark as found
						found = true
						break
					}
				}

				if !found && q > 0 {
					// create a new stack of ore / ice
					nid := uuid.New()

					newItem := Item{
						ID:            nid,
						ItemTypeID:    unitType,
						Meta:          c.ItemTypeMeta,
						Created:       time.Now(),
						CreatedBy:     &m.shipMountedOn.UserID,
						CreatedReason: fmt.Sprintf("Mined %v", c.ItemFamilyID),
						ContainerID:   m.shipMountedOn.CargoBayContainerID,
						Quantity:      q,
						IsPackaged:    true,
						Lock:          sync.Mutex{}, ItemTypeName: c.ItemTypeName,
						ItemFamilyID:   c.ItemFamilyID,
						ItemFamilyName: c.ItemFamilyName,
						ItemTypeMeta:   c.ItemTypeMeta,
						CoreDirty:      true,
					}

					// escalate to core for saving in db
					m.shipMountedOn.CurrentSystem.NewItems = append(m.shipMountedOn.CurrentSystem.NewItems, &newItem)

					// add new item to cargo hold
					m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, &newItem)
				}

				// log mine to console if player
				if !m.shipMountedOn.IsNPC {
					bm := 0

					if m.shipMountedOn.BehaviourMode != nil {
						bm = *m.shipMountedOn.BehaviourMode
					}

					shared.TeeLog(
						fmt.Sprintf(
							"[%v] %v (%v::%v) mined %v %v from %v",
							m.shipMountedOn.CurrentSystem.SystemName,
							m.shipMountedOn.CharacterName,
							m.shipMountedOn.Texture,
							bm,
							q,
							c.ItemTypeName,
							c.Name,
						),
					)
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsUtilitySiphon() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// get max siphon amount
	maxSiphonAmt, _ := m.ItemMeta.GetFloat64("energy_siphon_amount")

	// apply usage experience modifier
	maxSiphonAmt *= m.usageExperienceModifier

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// adjust siphon amount based on tracking
	maxSiphonAmt *= trackingRatio

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% effect

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// max amount siphoned is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)

			maxSiphonAmt *= rangeRatio
		} else if falloff == "reverse_linear" {
			//  max amount siphoned is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			maxSiphonAmt *= rangeRatio
		}
	}

	// siphon energy from target ship
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)

		actualSiphon := c.siphonEnergy(
			maxSiphonAmt,
			m.shipMountedOn.ReputationSheet,
			m,
		)

		// add to energy
		m.shipMountedOn.Energy += actualSiphon

		// any excess becomes heat
		maxEnergy := m.shipMountedOn.GetRealMaxEnergy(false)
		excess := m.shipMountedOn.Energy - maxEnergy

		if excess > 0 {
			// apply heat
			m.shipMountedOn.Heat += excess

			// clamp energy
			m.shipMountedOn.Energy = maxEnergy
		}
	} else {
		return false
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsUtilityCloak() bool {
	// make sure that no other modules are active
	for _, x := range m.shipMountedOn.Fitting.ARack {
		if x.IsCycling {
			return false
		}
	}

	for _, x := range m.shipMountedOn.Fitting.BRack {
		if x.IsCycling {
			return false
		}
	}

	// get activation energy and duration (approximately same as cooldown for cloaking devices)
	activationEnergy, _ := m.ItemMeta.GetFloat64("activation_energy")
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")

	// get ship mass
	mx := m.shipMountedOn.GetRealMass()

	// calculate "cloak amount" (if percentage < 100% then cloaking will "flicker")
	dC := (activationEnergy / mx) * 7

	// apply usage experience modifier
	dC *= m.usageExperienceModifier

	// calculate effect duration in ticks
	dT := (cooldown * 1000) / Heartbeat

	// add temporary modifier
	modifier := TemporaryShipModifier{
		Attribute:      "cloak",
		Percentage:     dC,
		RemainingTicks: int(dT) + int(dC*10), // duration bonus for lighter ships
	}

	m.shipMountedOn.TemporaryModifiers = append(m.shipMountedOn.TemporaryModifiers, modifier)

	// module activates!
	return true
}

func (m *FittedSlot) activateAsSalvager() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Wreck {
		// find wreck
		tgt, f := m.shipMountedOn.CurrentSystem.wrecks[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			// raise missing wreck fault
			m.shipMountedOn.aiNoWreckFault = true

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// get max salvage amount
	maxSalvageVol, _ := m.ItemMeta.GetFloat64("salvage_volume")

	// get salvage chance
	salvageChance, _ := m.ItemMeta.GetFloat64("salvage_chance")

	// apply usage experience modifier
	maxSalvageVol *= m.usageExperienceModifier
	salvageChance *= m.usageExperienceModifier

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// adjust based on tracking
	maxSalvageVol *= trackingRatio
	salvageChance *= trackingRatio

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% effect

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// max amount salvaged is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)

			maxSalvageVol *= rangeRatio
		} else if falloff == "reverse_linear" {
			//  max amount salvaged is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			maxSalvageVol *= rangeRatio
		}
	}

	// attempt to pull from target wreck
	if *m.TargetType == tgtReg.Wreck {
		c := tgtI.(*Wreck)

		i, sv := c.TrySalvage(salvageChance, maxSalvageVol)

		if i != nil {
			// success! try to add to cargo bay
			cv := m.shipMountedOn.GetRealCargoBayVolume(false)
			cu := m.shipMountedOn.TotalCargoBayVolumeUsed(false)

			if cu+sv <= cv {
				// store in cargo bay
				m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, i)
				i.ContainerID = m.shipMountedOn.CargoBayContainerID

				// escalate to core for saving
				i.CoreDirty = true
				m.shipMountedOn.CurrentSystem.MovedItems = append(m.shipMountedOn.CurrentSystem.MovedItems, i)

				// log salvage to console if player
				if !m.shipMountedOn.IsNPC {
					bm := 0

					if m.shipMountedOn.BehaviourMode != nil {
						bm = *m.shipMountedOn.BehaviourMode
					}

					shared.TeeLog(
						fmt.Sprintf(
							"[%v] %v (%v::%v) salvaged %v %v from %v",
							m.shipMountedOn.CurrentSystem.SystemName,
							m.shipMountedOn.CharacterName,
							m.shipMountedOn.Texture,
							bm,
							i.Quantity,
							i.ItemTypeName,
							c.WreckName,
						),
					)
				}
			}
		}
	} else {
		return false
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsUtilityVeil() bool {
	// get attributes
	activationEnergy, _ := m.ItemMeta.GetFloat64("activation_energy")
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")

	// get ship mass
	mx := m.shipMountedOn.GetRealMass()

	// calculate "veil amount" (amount of incoming damage absorbed)
	dC := (activationEnergy / mx)

	// apply usage experience modifier
	dC *= m.usageExperienceModifier

	// make sure its never 100%
	dC *= 0.99

	// calculate effect duration in ticks
	dT := (cooldown * 1000) / Heartbeat

	// add temporary modifier
	modifier := TemporaryShipModifier{
		Attribute:      "veil",
		Percentage:     dC,
		RemainingTicks: int(dT),
	}

	m.shipMountedOn.TemporaryModifiers = append(m.shipMountedOn.TemporaryModifiers, modifier)

	// module activates!
	return true
}

func (m *FittedSlot) activateAsFuelLoader() bool {
	// get max pellet volume and leakage
	maxVolume, _ := m.ItemMeta.GetFloat64("max_fuel_volume")
	leakage, _ := m.ItemTypeMeta.GetFloat64("leakage")

	// apply usage experience modifier
	maxVolume *= m.usageExperienceModifier

	// locate a stack of pellets below the given unit volume
	for _, i := range m.shipMountedOn.CargoBay.Items {
		// safety checks
		if !i.IsPackaged {
			continue
		}

		if i.Quantity <= 0 {
			continue
		}

		if i.CoreDirty {
			continue
		}

		// check family
		if i.ItemFamilyID != "fuel" {
			continue
		}

		// check volume
		iv, _ := i.ItemTypeMeta.GetFloat64("volume")

		if iv <= maxVolume {
			// decrement quantity
			i.Quantity--

			// escalate to core for saving
			m.shipMountedOn.CurrentSystem.ChangedQuantityItems = append(m.shipMountedOn.CurrentSystem.ChangedQuantityItems, i)

			// get fuel amount
			bf, _ := i.ItemTypeMeta.GetFloat64("fuelconversion")

			// apply leakage
			af := bf * (1 - leakage)

			// add fuel to tank
			m.shipMountedOn.Fuel += af

			// module activates!
			return true
		}
	}

	// module doesn't activate
	return false
}

func (m *FittedSlot) activateAsAreaDenialDevice() bool {
	// include visual effect if present
	activationPGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		tgtReg := models.SharedTargetTypeRegistry

		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationPGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// get damage values
	shieldDmg, _ := m.ItemMeta.GetFloat64("shield_damage")
	armorDmg, _ := m.ItemMeta.GetFloat64("armor_damage")
	hullDmg, _ := m.ItemMeta.GetFloat64("hull_damage")

	// get range value
	rge, _ := m.ItemMeta.GetFloat64("range")

	// get falloff style
	falloff, ff := m.ItemMeta.GetString("falloff")

	// get drag multiplier
	dMul, _ := m.ItemMeta.GetFloat64("drag_multiplier")

	// get heat damage
	hDmg, _ := m.ItemMeta.GetFloat64("heat_damage")

	// get missile destruction chance
	mDest, _ := m.ItemMeta.GetFloat64("missile_destruction_chance")

	// get energy siphon amount
	eSiph, _ := m.ItemMeta.GetFloat64("energy_siphon_amount")

	// apply usage experience modifiers
	shieldDmg *= m.usageExperienceModifier
	armorDmg *= m.usageExperienceModifier
	hullDmg *= m.usageExperienceModifier
	dMul *= m.usageExperienceModifier
	hDmg *= m.usageExperienceModifier
	eSiph *= m.usageExperienceModifier

	// iterate over missiles in the system
	for _, ts := range m.shipMountedOn.CurrentSystem.missiles {
		// check range
		dA := m.shipMountedOn.ToPhysicsDummy()
		dB := ts.ToPhysicsDummy()

		tR := physics.Distance(dA, dB)

		if tR < rge {
			// determine falloff amount
			rangeRatio := 1.0 // default to 100%

			if ff {
				// adjust based on falloff style
				if falloff == "linear" {
					// proportion of the distance to target over max range (closer is higher)
					rangeRatio = 1 - (tR / rge)
				} else if falloff == "reverse_linear" {
					// proportion of the distance to target over max range (further is higher)
					rangeRatio = (tR / rge)

					if tR > rge {
						rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
					}
				}
			}

			if dMul > 0 {
				// determine max velocity reduction
				fv := (1.0 - rangeRatio) / dMul

				if fv > 1 {
					fv = 1
				}

				if fv < 0 {
					fv = 0
				}

				// reduce max missile velocity
				ts.MaxVelocity *= fv
			}

			if mDest > 0 {
				// get intrinsic failure chance of missile
				ft := 1.0 - ts.FaultTolerance

				// determine total failure chance
				fc := (mDest * rangeRatio) + ft

				// roll
				or := rand.Float64()

				if or < fc {
					// destroy missile
					ts.TicksRemaining = 0
				}
			}
		}
	}

	// iterate over ships in the system
	for _, ts := range m.shipMountedOn.CurrentSystem.ships {
		// don't affect firing ship
		if ts.ID == m.shipMountedOn.ID {
			continue
		}

		// skip if docked
		if ts.IsDocked {
			continue
		}

		// check range
		dA := m.shipMountedOn.ToPhysicsDummy()
		dB := ts.ToPhysicsDummy()

		tR := physics.Distance(dA, dB)

		if tR < rge {
			// determine falloff amount
			rangeRatio := 1.0 // default to 100%

			if ff {
				// adjust based on falloff style
				if falloff == "linear" {
					// proportion of the distance to target over max range (closer is higher)
					rangeRatio = 1 - (tR / rge)
				} else if falloff == "reverse_linear" {
					// proportion of the distance to target over max range (further is higher)
					rangeRatio = (tR / rge)

					if tR > rge {
						rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
					}
				}
			}

			// deal damage
			ts.dealDamage(
				shieldDmg*rangeRatio,
				armorDmg*rangeRatio,
				hullDmg*rangeRatio,
				m.shipMountedOn.ReputationSheet,
				m,
			)

			// apply siphon
			actualSiphon := ts.siphonEnergy(
				eSiph*rangeRatio,
				m.shipMountedOn.ReputationSheet,
				m,
			)

			// add to energy
			m.shipMountedOn.Energy += actualSiphon

			// any excess becomes heat
			maxEnergy := m.shipMountedOn.GetRealMaxEnergy(false)
			excess := m.shipMountedOn.Energy - maxEnergy

			if excess > 0 {
				// apply heat
				m.shipMountedOn.Heat += excess

				// clamp energy
				m.shipMountedOn.Energy = maxEnergy
			}

			// calculate drag effect duration in ticks
			cooldown, _ := m.ItemMeta.GetFloat64("cooldown")
			dT := (cooldown * 1000) / Heartbeat

			// add temporary modifier to target
			modifier := TemporaryShipModifier{
				Attribute:      "drag",
				Percentage:     dMul * rangeRatio,
				RemainingTicks: int(dT),
			}

			ts.TemporaryModifiers = append(ts.TemporaryModifiers, modifier)

			// apply heat damage
			ts.Heat += hDmg * rangeRatio
		}
	}

	// module activates
	return true
}

func (m *FittedSlot) activateAsCycleDisruptor() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI interface{}

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// get falloff style
	falloff, ff := m.ItemMeta.GetString("falloff")
	rangeRatio := 1.0 // default to 100%

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}

		// account for falloff if present
		if ff {
			// adjust based on falloff style
			if falloff == "linear" {
				// proportion of the distance to target over max range (closer is higher)
				rangeRatio = 1 - (d / modRange)
			} else if falloff == "reverse_linear" {
				// proportion of the distance to target over max range (further is higher)
				rangeRatio = (d / modRange)

				if d > modRange {
					rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
				}
			}
		}
	}

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// get signal flux
	sigFlux, _ := m.ItemMeta.GetFloat64("signal_flux")
	signalGain, _ := m.ItemMeta.GetFloat64("signal_gain")

	// iterate over target modules
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)

		for r := 0; r < 2; r++ {
			var rk []FittedSlot

			// get rack for iteration
			if r == 0 {
				rk = c.Fitting.ARack
			} else if r == 1 {
				rk = c.Fitting.BRack
			}

			// iterate over slots
			for ix := range rk {
				// get target module
				tm := rk[ix]

				// skip if not cycling
				if !tm.IsCycling {
					continue
				}

				// get target numbers
				dC, dA := rollCycleDisruptor(tm, sigFlux, signalGain)

				// apply experience modifiers
				dC *= m.usageExperienceModifier

				// apply tracking
				dC *= trackingRatio

				// apply falloff
				dA *= rangeRatio

				// roll for disruption
				roll := rand.Float64()

				if roll <= dC {
					// apply effect to cycle progress
					dp := int(float64(tm.cooldownProgress) * dA)
					tm.cooldownProgress -= dp

					// store update to module
					if r == 0 {
						c.Fitting.ARack[ix] = tm
					} else if r == 1 {
						c.Fitting.BRack[ix] = tm
					}
				}
			}
		}

		// update aggression tables
		c.updateAggressionTables(
			0,
			0,
			0,
			m.shipMountedOn.ReputationSheet,
			m,
		)

		// include visual effect if present
		activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

		if found {
			// build effect trigger
			gfxEffect := models.GlobalPushModuleEffectBody{
				GfxEffect:    activationGfxEffect,
				ObjStartID:   m.shipMountedOn.ID,
				ObjStartType: tgtReg.Ship,
				ObjEndID:     m.TargetID,
				ObjEndType:   m.TargetType,
			}

			// push to solar system list for next update
			m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
		}
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsFireControlJammer() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtS *Ship

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtS = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// get falloff style
	falloff, ff := m.ItemMeta.GetString("falloff")
	rangeRatio := 1.0 // default to 100%

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}

		// account for falloff if present
		if ff {
			// adjust based on falloff style
			if falloff == "linear" {
				// proportion of the distance to target over max range (closer is higher)
				rangeRatio = 1 - (d / modRange)
			} else if falloff == "reverse_linear" {
				// proportion of the distance to target over max range (further is higher)
				rangeRatio = (d / modRange)

				if d > modRange {
					rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
				}
			}
		}
	}

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// apply usage experience modifier
	trackingRatio *= m.GetExperienceModifier()
	rangeRatio *= m.GetExperienceModifier()

	// get drift
	guidDrift, _ := m.ItemMeta.GetFloat64("guidance_drift")
	trackDrift, _ := m.ItemMeta.GetFloat64("tracking_drift")

	// apply falloff
	guidDrift *= rangeRatio
	trackDrift *= rangeRatio

	// apply tracking
	guidDrift *= trackingRatio
	trackDrift *= trackingRatio

	// calculate effect duration in ticks
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")
	dT := (cooldown * 1000) / Heartbeat

	// add temporary modifiers to target
	guidMod := TemporaryShipModifier{
		Attribute:      "guidance_drift",
		Raw:            guidDrift,
		RemainingTicks: int(dT),
	}

	trackMod := TemporaryShipModifier{
		Attribute:      "tracking_drift",
		Raw:            trackDrift,
		RemainingTicks: int(dT),
	}

	tgtS.TemporaryModifiers = append(tgtS.TemporaryModifiers, guidMod)
	tgtS.TemporaryModifiers = append(tgtS.TemporaryModifiers, trackMod)

	// update aggression tables
	tgtS.updateAggressionTables(
		0,
		0,
		0,
		m.shipMountedOn.ReputationSheet,
		m,
	)

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsRegenerationMask() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtS *Ship

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtS = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// get falloff style
	falloff, ff := m.ItemMeta.GetString("falloff")
	rangeRatio := 1.0 // default to 100%

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}

		// account for falloff if present
		if ff {
			// adjust based on falloff style
			if falloff == "linear" {
				// proportion of the distance to target over max range (closer is higher)
				rangeRatio = 1 - (d / modRange)
			} else if falloff == "reverse_linear" {
				// proportion of the distance to target over max range (further is higher)
				rangeRatio = (d / modRange)

				if d > modRange {
					rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
				}
			}
		}
	}

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// get mask radius
	maskRad, _ := m.ItemMeta.GetFloat64("mask_radius")

	// apply usage experience modifier
	maskRad *= m.GetExperienceModifier()

	// apply falloff
	maskRad *= rangeRatio

	// apply tracking
	maskRad *= trackingRatio

	// calculate mask percentage
	maskP := maskRad / tgtS.TemplateData.Radius

	// clamp to 0%
	if maskP < 0 {
		maskP = 0
	}

	// clamp to 100%
	if maskP > 1 {
		maskP = 1
	}

	// calculate effect duration in ticks
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")
	dT := (cooldown * 1000) / Heartbeat

	// add temporary modifiers to target
	maskMod := TemporaryShipModifier{
		Attribute:      "regeneration_mask",
		Percentage:     maskP,
		RemainingTicks: int(dT),
	}

	tgtS.TemporaryModifiers = append(tgtS.TemporaryModifiers, maskMod)

	// update aggression tables
	tgtS.updateAggressionTables(
		0,
		0,
		0,
		m.shipMountedOn.ReputationSheet,
		m,
	)

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsDissipationMask() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtS *Ship

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtS = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// get falloff style
	falloff, ff := m.ItemMeta.GetString("falloff")
	rangeRatio := 1.0 // default to 100%

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}

		// account for falloff if present
		if ff {
			// adjust based on falloff style
			if falloff == "linear" {
				// proportion of the distance to target over max range (closer is higher)
				rangeRatio = 1 - (d / modRange)
			} else if falloff == "reverse_linear" {
				// proportion of the distance to target over max range (further is higher)
				rangeRatio = (d / modRange)

				if d > modRange {
					rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
				}
			}
		}
	}

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// get mask radius
	maskRad, _ := m.ItemMeta.GetFloat64("mask_radius")

	// apply usage experience modifier
	maskRad *= m.GetExperienceModifier()

	// apply falloff
	maskRad *= rangeRatio

	// apply tracking
	maskRad *= trackingRatio

	// calculate mask percentage
	maskP := maskRad / tgtS.TemplateData.Radius

	// clamp to 0%
	if maskP < 0 {
		maskP = 0
	}

	// clamp to 100%
	if maskP > 1 {
		maskP = 1
	}

	// calculate effect duration in ticks
	cooldown, _ := m.ItemMeta.GetFloat64("cooldown")
	dT := (cooldown * 1000) / Heartbeat

	// add temporary modifiers to target
	maskMod := TemporaryShipModifier{
		Attribute:      "dissipation_mask",
		Percentage:     maskP,
		RemainingTicks: int(dT),
	}

	tgtS.TemporaryModifiers = append(tgtS.TemporaryModifiers, maskMod)

	// update aggression tables
	tgtS.updateAggressionTables(
		0,
		0,
		0,
		m.shipMountedOn.ReputationSheet,
		m,
	)

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsBurstReactor() bool {
	// get max pellet volume and leakage
	maxVolume, _ := m.ItemMeta.GetFloat64("max_fuel_volume")
	leakage, _ := m.ItemTypeMeta.GetFloat64("leakage")

	// apply usage experience modifier
	maxVolume *= m.usageExperienceModifier

	// locate a stack of pellets below the given unit volume
	for _, i := range m.shipMountedOn.CargoBay.Items {
		// safety checks
		if !i.IsPackaged {
			continue
		}

		if i.Quantity <= 0 {
			continue
		}

		if i.CoreDirty {
			continue
		}

		// check family
		if i.ItemFamilyID != "fuel" {
			continue
		}

		// check volume
		iv, _ := i.ItemTypeMeta.GetFloat64("volume")

		if iv <= maxVolume {
			// decrement quantity
			i.Quantity--

			// escalate to core for saving
			m.shipMountedOn.CurrentSystem.ChangedQuantityItems = append(m.shipMountedOn.CurrentSystem.ChangedQuantityItems, i)

			// get fuel amount
			bf, _ := i.ItemTypeMeta.GetFloat64("fuelconversion")

			// apply leakage
			af := bf * (1 - leakage)

			// apply energy
			m.shipMountedOn.Energy += af

			// get max energy
			me := m.shipMountedOn.GetRealMaxEnergy(false)

			// get excess for heat
			eh := m.shipMountedOn.Energy - me

			if eh > 0 {
				// add heat
				m.shipMountedOn.Heat += eh
			}

			// cap energy
			if m.shipMountedOn.Energy > me {
				m.shipMountedOn.Energy = me
			}

			// module activates!
			return true
		}
	}

	// module doesn't activate
	return false
}

// Reusable helper function to determine cycle disruptor chance and amount respectively
func rollCycleDisruptor(tgt FittedSlot, sigFlux float64, sigGain float64) (float64, float64) {
	// get target module v
	v, _ := tgt.ItemTypeMeta.GetFloat64("volume")
	v += Epsilon

	/*
		<chance of disruption> = [signal flux] / [tgt module volume]
								 (centered at the slot size of a the ship tier mounting the module)
	*/

	// calculate disruption chance
	dC := sigFlux / v

	// get cycle progress
	c, _ := tgt.ItemMeta.GetFloat64("cooldown")
	c += Epsilon

	p := float64(tgt.cooldownProgress) / c

	/*
		<disruption amount> = [(<cycle progress>^sin(<cycle progress>/<signal gain>)) mod <signal gain>] / <signal gain>
							  (centered at ~25, does not scale with ship tier slot size)
	*/

	// calculate disruption amount
	q := p / sigGain
	r := math.Pow(p, math.Sin(q))

	dA := float64(int(r)%int(sigGain)) / sigGain

	// return results
	return dC, dA
}

// Reusable helper function to determine tracking ratio between a module and a target
func (m *FittedSlot) calculateTrackingRatioWithTarget(tgtDummy physics.Dummy) float64 {
	// calculate distance
	d := physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

	// determine angular velocity for tracking
	dvX := m.shipMountedOn.VelX - tgtDummy.VelX
	dvY := m.shipMountedOn.VelY - tgtDummy.VelY

	dv := math.Sqrt((dvX * dvX) + (dvY * dvY))
	w := 0.0

	if d > 0 {
		w = ((dv / d) * float64(Heartbeat)) * (180.0 / math.Pi)
	}

	// get tracking value
	tracking, _ := m.ItemMeta.GetFloat64("tracking")

	// calculate tracking ratio
	trackingRatio := 1.0 // default to 100% tracking

	if w > 0 {
		trackingRatio = tracking / w
	}

	// accumulate tracking drift modifiers
	rawDrift := 0.0

	for _, mo := range m.shipMountedOn.TemporaryModifiers {
		if mo.Attribute == "tracking_drift" {
			rawDrift += mo.Raw
		}
	}

	if rawDrift > 0 {
		// get drift percentage
		driftP := rawDrift / tracking

		// decrement from tracking ratio
		trackingRatio -= driftP
	}

	// clamp tracking to 0%
	if trackingRatio < 0 {
		trackingRatio = 0
	}

	// clamp tracking to 100%
	if trackingRatio > 1.0 {
		trackingRatio = 1.0
	}

	return trackingRatio
}

// Calculates the experience percentage bonus to apply to some active module stats
func (m *FittedSlot) GetExperienceModifier() float64 {
	mx := 1.0

	if m.shipMountedOn.ExperienceSheet != nil {
		// get experience entry for this item type as a module
		v := m.shipMountedOn.ExperienceSheet.GetModuleExperienceEntry(m.ItemTypeID)

		// get truncated level
		l := math.Trunc(v.GetExperience())

		if l > 0 {
			// apply a dampening factor to get percentage
			b := math.Log(((math.Pow(l, 0.75)) / 8.8) + 1)

			if b > 0 {
				// add bonus
				mx += b
			}
		}
	}

	return mx
}

// Updates aggression records associated with a ship
func (s *Ship) updateAggressionTables(
	shieldDmg float64,
	armorDmg float64,
	hullDmg float64,
	attackerRS *shared.PlayerReputationSheet,
	attackerModule *FittedSlot,
) {
	// update aggression table
	if attackerRS != nil {
		// obtain lock
		attackerRS.Lock.Lock()
		defer attackerRS.Lock.Unlock()

		// get attacker's reputation sheet entry for this ship's faction
		f, ok := attackerRS.FactionEntries[s.FactionID.String()]

		if !ok {
			// does not exist - create a neutral one
			ne := shared.PlayerReputationSheetFactionEntry{
				FactionID:        s.FactionID,
				StandingValue:    0,
				AreOpenlyHostile: false,
			}

			attackerRS.FactionEntries[s.FactionID.String()] = &ne
			f = attackerRS.FactionEntries[s.FactionID.String()]
		}

		if shieldDmg+armorDmg+hullDmg >= 0 {
			// update temporary hostility due to aggro flag
			at := time.Now().Add(15 * time.Minute)
			f.TemporarilyOpenlyHostileUntil = &at
		}

		// store entry
		s.Aggressors[attackerRS.UserID.String()] = attackerRS
	}

	// update aggression table
	if attackerModule != nil {
		// get attacker's aggro sheet entry
		f, ok := s.AggressionLog[attackerModule.shipMountedOn.ID.String()]

		if !ok {
			// does not exist - create blank one
			s.AggressionLog[attackerModule.shipMountedOn.ID.String()] = &shared.AggressionLog{
				UserID:           attackerModule.shipMountedOn.UserID,
				FactionID:        attackerModule.shipMountedOn.FactionID,
				CharacterName:    attackerModule.shipMountedOn.CharacterName,
				FactionName:      attackerModule.shipMountedOn.Faction.Name,
				IsNPC:            attackerModule.shipMountedOn.IsNPC,
				LastAggressed:    time.Now(),
				ShipID:           attackerModule.shipMountedOn.ID,
				ShipName:         attackerModule.shipMountedOn.ShipName,
				ShipTemplateID:   attackerModule.shipMountedOn.TemplateData.ID,
				ShipTemplateName: attackerModule.shipMountedOn.TemplateData.ShipTemplateName,
				WeaponUse:        make(map[string]*shared.AggressionLogWeaponUse),
			}

			f = s.AggressionLog[attackerModule.shipMountedOn.ID.String()]
		}

		// get module entry from aggro row
		g, ok := f.WeaponUse[attackerModule.ItemID.String()]

		if !ok {
			// does not exist - create a blank one
			f.WeaponUse[attackerModule.ItemID.String()] = &shared.AggressionLogWeaponUse{
				ItemID:          attackerModule.ItemID,
				ItemTypeID:      attackerModule.ItemTypeID,
				ItemFamilyName:  attackerModule.ItemTypeFamilyName,
				ItemTypeName:    attackerModule.ItemTypeName,
				ItemFamilyID:    attackerModule.ItemTypeFamily,
				DamageInflicted: 0,
				LastUsed:        time.Now(),
			}

			g = f.WeaponUse[attackerModule.ItemID.String()]
		}

		// update aggro info
		f.LastSolarSystemID = attackerModule.shipMountedOn.CurrentSystem.ID
		f.LastSolarSystemName = attackerModule.shipMountedOn.CurrentSystem.SystemName
		f.LastRegionID = attackerModule.shipMountedOn.CurrentSystem.RegionID
		f.LastRegionName = attackerModule.shipMountedOn.CurrentSystem.RegionName
		f.LastPosX = attackerModule.shipMountedOn.PosX
		f.LastPosY = attackerModule.shipMountedOn.PosY
		f.LastAggressed = time.Now()

		g.DamageInflicted += (shieldDmg + armorDmg + hullDmg)
		g.LastUsed = time.Now()
	}
}

func (m *FittedSlot) activateAsHeatXfer() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// get max transfer amount
	maxXferAmt, _ := m.ItemMeta.GetFloat64("heat_xfer_amount")

	// apply usage experience modifier
	maxXferAmt *= m.usageExperienceModifier

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// adjust transfer amount based on tracking
	maxXferAmt *= trackingRatio

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% effect

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// max amount transfered is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)

			maxXferAmt *= rangeRatio
		} else if falloff == "reverse_linear" {
			//  max amount transfered is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			maxXferAmt *= rangeRatio
		}
	}

	// transfer heat from target ship
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)

		// verify both ships are in the same faction
		if m.shipMountedOn.FactionID != c.FactionID {
			return false
		}

		// reduce heat on target
		actualXfer := c.siphonHeat(maxXferAmt)

		// increase local heat
		m.shipMountedOn.Heat += actualXfer
	} else {
		return false
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsEnergyXfer() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// get max transfer amount
	maxXferAmt, _ := m.ItemMeta.GetFloat64("energy_xfer_amount")

	// apply usage experience modifier
	maxXferAmt *= m.usageExperienceModifier

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// adjust transfer amount based on tracking
	maxXferAmt *= trackingRatio

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% effect

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// max amount transfered is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)

			maxXferAmt *= rangeRatio
		} else if falloff == "reverse_linear" {
			//  max amount transfered is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			maxXferAmt *= rangeRatio
		}
	}

	// clamp at current energy amount
	if maxXferAmt > m.shipMountedOn.Energy {
		maxXferAmt = m.shipMountedOn.Energy
	}

	// transfer energy to target ship
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)

		// verify both ships are in the same faction
		if m.shipMountedOn.FactionID != c.FactionID {
			return false
		}

		// increase energy on target
		actualXfer := c.receiveEnergy(maxXferAmt)

		// decrease local energy
		m.shipMountedOn.Energy -= actualXfer
	} else {
		return false
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsShieldXfer() bool {
	// safety check targeting pointers
	if m.TargetID == nil || m.TargetType == nil {
		m.WillRepeat = false
		return false
	}

	// get target
	tgtReg := models.SharedTargetTypeRegistry

	// target details
	var tgtDummy physics.Dummy = physics.Dummy{}
	var tgtI Any

	if *m.TargetType == tgtReg.Ship {
		// find ship
		tgt, f := m.shipMountedOn.CurrentSystem.ships[m.TargetID.String()]

		if !f {
			// target doesn't exist - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false

			return false
		}

		// store target details
		tgtDummy = tgt.ToPhysicsDummy()
		tgtI = tgt
	} else {
		// unsupported target type - can't activate
		m.TargetID = nil
		m.TargetType = nil
		m.WillRepeat = false
		m.IsCycling = false

		return false
	}

	// check for max range
	modRange, found := m.ItemMeta.GetFloat64("range")
	var d float64 = 0

	if found {
		// get distance to target
		d = physics.Distance(tgtDummy, m.shipMountedOn.ToPhysicsDummy())

		// verify target is in range
		if d > modRange {
			// out of range - can't activate
			m.TargetID = nil
			m.TargetType = nil
			m.WillRepeat = false
			m.IsCycling = false

			return false
		}
	}

	// get max transfer amount
	maxXferAmt, _ := m.ItemMeta.GetFloat64("shield_xfer_amount")

	// apply usage experience modifier
	maxXferAmt *= m.usageExperienceModifier

	// calculate tracking ratio
	trackingRatio := m.calculateTrackingRatioWithTarget(tgtDummy)

	// adjust transfer amount based on tracking
	maxXferAmt *= trackingRatio

	// account for falloff if present
	falloff, found := m.ItemMeta.GetString("falloff")

	rangeRatio := 1.0 // default to 100% effect

	if found {
		// adjust based on falloff style
		if falloff == "linear" {
			// max amount transfered is a proportion of the distance to target over max range (closer is higher)
			rangeRatio = 1 - (d / modRange)

			maxXferAmt *= rangeRatio
		} else if falloff == "reverse_linear" {
			//  max amount transfered is a proportion of the distance to target over max range (further is higher)
			rangeRatio = (d / modRange)

			if d > modRange {
				rangeRatio = 0 // sharp cutoff if out of range to avoid sillinesss
			}

			maxXferAmt *= rangeRatio
		}
	}

	// clamp at current shield amount
	if maxXferAmt > m.shipMountedOn.Shield {
		maxXferAmt = m.shipMountedOn.Shield
	}

	// transfer shield to target ship
	if *m.TargetType == tgtReg.Ship {
		c := tgtI.(*Ship)

		// verify both ships are in the same faction
		if m.shipMountedOn.FactionID != c.FactionID {
			return false
		}

		// increase shield on target
		actualXfer := c.receiveShield(maxXferAmt)

		// decrease local shield
		m.shipMountedOn.Shield -= actualXfer
	} else {
		return false
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
			ObjEndID:     m.TargetID,
			ObjEndType:   m.TargetType,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}

func (m *FittedSlot) activateAsUtilityWisper() bool {
	// verify this can harvest gas
	canMineGas, _ := m.ItemTypeMeta.GetBool("can_mine_gas")

	if !canMineGas {
		return false
	}

	// get ship dummy
	dA := m.shipMountedOn.ToPhysicsDummy()

	// effective yields table
	eyt := make(map[string]*float64)
	itm := make(map[string]GasMiningYield)

	// check stars
	for _, s := range m.shipMountedOn.CurrentSystem.stars {
		// null check
		if s == nil {
			continue
		}

		// get distance to player
		dB := s.ToPhysicsDummy()
		d := physics.Distance(dA, dB)

		// radius check
		if d < s.Radius/2 {
			d = s.Radius / 2
		}

		// zero check
		if d <= 0 {
			d = Epsilon
		}

		// iterate over minable gases
		for _, g := range s.GasMiningMetadata.Yields {
			// create table entry if missing
			if eyt[g.ItemTypeID.String()] == nil {
				// zero for reference
				z := 0.0

				// create entry for gas
				eyt[g.ItemTypeID.String()] = &z
			}

			// get current entry for gas
			ey := *eyt[g.ItemTypeID.String()]

			// add contribution
			ey += (float64(g.Yield) / d)

			// store result
			eyt[g.ItemTypeID.String()] = &ey
			itm[g.ItemTypeID.String()] = g
		}
	}

	// check planets
	for _, p := range m.shipMountedOn.CurrentSystem.planets {
		// null check
		if p == nil {
			continue
		}

		// get distance to player
		dB := p.ToPhysicsDummy()
		d := physics.Distance(dA, dB)

		// radius check
		if d < p.Radius/2 {
			d = p.Radius / 2
		}

		// zero check
		if d <= 0 {
			d = Epsilon
		}

		// iterate over minable gases
		for _, g := range p.GasMiningMetadata.Yields {
			// create table entry if missing
			if eyt[g.ItemTypeID.String()] == nil {
				// zero for reference
				z := 0.0

				// create entry for gas
				eyt[g.ItemTypeID.String()] = &z
			}

			// get current entry for gas
			ey := *eyt[g.ItemTypeID.String()]

			// add contribution
			ey += (float64(g.Yield) / d)

			// store result
			eyt[g.ItemTypeID.String()] = &ey
			itm[g.ItemTypeID.String()] = g
		}
	}

	// check asteroids
	for _, a := range m.shipMountedOn.CurrentSystem.asteroids {
		// null check
		if a == nil {
			continue
		}

		// get distance to player
		dB := a.ToPhysicsDummy()
		d := physics.Distance(dA, dB)

		// radius check
		if d < a.Radius/2 {
			d = a.Radius / 2
		}

		// zero check
		if d <= 0 {
			d = Epsilon
		}

		// iterate over minable gases
		for _, g := range a.GasMiningMetadata.Yields {
			// create table entry if missing
			if eyt[g.ItemTypeID.String()] == nil {
				// zero for reference
				z := 0.0

				// create entry for gas
				eyt[g.ItemTypeID.String()] = &z
			}

			// get current entry for gas
			ey := *eyt[g.ItemTypeID.String()]

			// add contribution
			ey += (float64(g.Yield) / d)

			// store result
			eyt[g.ItemTypeID.String()] = &ey
			itm[g.ItemTypeID.String()] = g
		}
	}

	// get cooldown
	cooldown, found := m.ItemMeta.GetFloat64("cooldown")

	if found {
		// apply cooldown time modifier
		for k, v := range eyt {
			// get yield
			ey := *v

			// multiply by cycle time
			ey *= cooldown

			// store result
			eyt[k] = &ey
		}
	}

	// get intake area
	aperture, found := m.ItemMeta.GetFloat64("intake_area")

	if found {
		// apply intake area modifier
		for k, v := range eyt {
			// get yield
			ey := *v

			// multiply by square root of intake area
			ey *= math.Sqrt(aperture)

			// apply experience modifier
			ey *= m.usageExperienceModifier

			// store result
			eyt[k] = &ey
		}
	} else {
		// bound to epsilon
		aperture = Epsilon
	}

	// attempt to pull and store in cargo
	for k, ey := range eyt {
		// get mining volume
		miningVolume := aperture

		// get yield meta
		ym := itm[k]
		eyy := *ey

		// get type and volume gas being collected
		unitType := ym.ItemTypeID
		unitVol, _ := ym.ItemTypeMeta.GetFloat64("volume")

		// get available space in cargo hold
		free := m.shipMountedOn.GetRealCargoBayVolume(false) - m.shipMountedOn.TotalCargoBayVolumeUsed(false)

		// calculate effective volume pulled
		pulled := math.Min(miningVolume, eyy)

		// make sure there is sufficient room to deposit the gas
		if free-pulled >= 0 {
			found := false

			// quantity to be placed in cargo bay
			q := int(pulled / unitVol)

			if q <= 0 && m.shipMountedOn.IsNPC {
				// raise fault
				m.shipMountedOn.aiNoGasPulledFault = true
			}

			// is there already packaged gas of this type in the hold?
			for idx := range m.shipMountedOn.CargoBay.Items {
				itm := m.shipMountedOn.CargoBay.Items[idx]

				if itm.ItemTypeID == unitType && itm.IsPackaged && !itm.CoreDirty {
					// increase the size of this stack
					itm.Quantity += q

					// escalate for saving
					m.shipMountedOn.CurrentSystem.ChangedQuantityItems = append(m.shipMountedOn.CurrentSystem.ChangedQuantityItems, itm)

					// mark as found
					found = true
					break
				}
			}

			if !found && q > 0 {
				// create a new stack of gas
				nid := uuid.New()

				newItem := Item{
					ID:             nid,
					ItemTypeID:     unitType,
					Meta:           ym.ItemTypeMeta,
					Created:        time.Now(),
					CreatedBy:      &m.shipMountedOn.UserID,
					CreatedReason:  fmt.Sprintf("Harvested %v", ym.ItemFamilyID),
					ContainerID:    m.shipMountedOn.CargoBayContainerID,
					Quantity:       q,
					IsPackaged:     true,
					Lock:           sync.Mutex{},
					ItemTypeName:   ym.ItemTypeName,
					ItemFamilyID:   ym.ItemFamilyID,
					ItemFamilyName: ym.ItemFamilyName,
					ItemTypeMeta:   ym.ItemTypeMeta,
					CoreDirty:      true,
				}

				// escalate to core for saving in db
				m.shipMountedOn.CurrentSystem.NewItems = append(m.shipMountedOn.CurrentSystem.NewItems, &newItem)

				// add new item to cargo hold
				m.shipMountedOn.CargoBay.Items = append(m.shipMountedOn.CargoBay.Items, &newItem)
			}

			// log harvest to console if player
			if !m.shipMountedOn.IsNPC {
				bm := 0

				if m.shipMountedOn.BehaviourMode != nil {
					bm = *m.shipMountedOn.BehaviourMode
				}

				shared.TeeLog(
					fmt.Sprintf(
						"[%v] %v (%v::%v) havrested %v %v",
						m.shipMountedOn.CurrentSystem.SystemName,
						m.shipMountedOn.CharacterName,
						m.shipMountedOn.Texture,
						bm,
						q,
						ym.ItemTypeName,
					),
				)
			}
		} else {
			return false
		}
	}

	// include visual effect if present
	activationGfxEffect, found := m.ItemTypeMeta.GetString("activation_gfx_effect")

	if found {
		// get registry
		tgtReg := models.SharedTargetTypeRegistry

		// build effect trigger
		gfxEffect := models.GlobalPushModuleEffectBody{
			GfxEffect:    activationGfxEffect,
			ObjStartID:   m.shipMountedOn.ID,
			ObjStartType: tgtReg.Ship,
		}

		gfxEffect.ObjStartHardpointOffset = [...]float64{
			0,
			0,
		}

		if m.SlotIndex != nil {
			rack := m.Rack
			idx := *m.SlotIndex

			if rack == "A" {
				gfxEffect.ObjStartHardpointOffset = m.shipMountedOn.TemplateData.SlotLayout.ASlots[idx].TexturePosition
			}
		}

		// push to solar system list for next update
		m.shipMountedOn.CurrentSystem.pushModuleEffects = append(m.shipMountedOn.CurrentSystem.pushModuleEffects, gfxEffect)
	}

	// module activates!
	return true
}
