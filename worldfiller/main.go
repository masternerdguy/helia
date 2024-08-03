package main

import (
	"encoding/csv"
	"fmt"
	"helia/engine"
	"helia/physics"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

/*
 * Contains routines for procedurally filling a scaffolded universe
 * (Designed to be run after world maker)
 */

func main() {
	// configure global tee logging
	shared.InitializeTeeLog(
		printLogger,
	)

	// load universe from database
	shared.TeeLog("Loading universe from database...")
	universe, err := engine.LoadUniverse()

	if err != nil {
		panic(err)
	}

	shared.TeeLog("Loaded universe!")

	/*
	 * COMMENT AND UNCOMMENT THE BELOW ROUTINES AS NEEDED
	 */

	/*var toInject = [...]string{
		"6b9b499d-1426-4aa5-ae58-69f6516d4a9b",
		"5e64e9d3-bf33-4552-8d09-1f70d7d2d4a4",
		"e77eee22-24f7-44c7-846a-09819808a8bf",
		"eab6ce76-bee0-4ab7-bc1d-87e7b90af831",
		"bdab0eaf-3f9d-40a8-b60f-4770e05f4a06",
		"c4dde5ab-b782-4d69-9582-7bfd02cd7d1c",
		"4075709e-d2c0-4d20-8698-d504938c22f5",
		"f120d11b-e0ad-4a17-bedc-25bb2bb7a06d",
		"e1e95a63-32f1-4f92-8052-b2ac7a92ec6b",
	}

	for i, e := range toInject {
		log.Printf("injecting process %v", e)
		injectProcess(universe, e, i)
	}*/

	//fillGasMiningYields(universe)

	// dropAsteroids(universe)
	//dropSanctuaryStations(universe)
	dropArtifacts(universe)

	//stubModuleSchematicsAndProcesses()
	//loadNewWares()
}

/* Parameters for asteroid generation */
const MinAsteroidsPerSystem = 0
const MaxAsteroidsPerSystem = 100
const MinAsteroidYield = 0.1
const MaxAsteroidYield = 5.0
const MinAsteroidRadius = 120
const MaxAsteroidRadius = 315
const MinAsteroidMass = 6000
const MaxAsteroidMass = 75000
const MinBeltRadius = 25000
const MaxBeltRadius = 1350000
const MinBeltWidth = 1000
const MaxBeltWidth = 10000

/*
============ ORE RARITY TABLE ========================================
id	                                    name	    probability	stop
dd522f03-2f52-4e82-b2f8-d7e0029cb82f	Testite	    0.1875	    0.1875
56617d30-6c30-425c-84bf-2484ae8c1156	Alri  	    0.1743	    0.3618
26a3fc9e-db2f-439d-a929-ba755d11d09c	Feymar	    0.1609    	0.5227
1d0d344b-ef28-43c8-a7a6-3275936b2dea	Listine	    0.0843   	0.607
0cd04eea-a150-410c-91eb-6af00d8c6eae	Hetrone	    0.0614  	0.6684
39b8eedf-ef80-4c29-a4bf-99abc4d84fa6	Novum	    0.0532  	0.7216
dd0c9b0a-279e-418e-b3b6-2f569fda0186	Suemetrium	0.0284  	0.75
7dcd5138-d7e0-419f-867a-6f0f23b99b5b	Jutrick	    0.0833  	0.8333
61f52ba3-654b-45cf-88e3-33399d12350d	Ovan	    0.0621  	0.8954
11688112-f3d4-4d30-864a-684a8b96ea23	Caiqua	    0.0382  	0.9336
2ce48bef-f06b-4550-b20c-0e64864db051	Zvitis	    0.0298  	0.9634
66b7a322-8cfc-4467-9410-492e6b58f159	Ichre	    0.0231  	0.9865
d1866be4-5c3e-4b95-b6d9-020832338014	Betro	    0.0135  	1

============ ASTEROID COUNT FOR SYSTEM ================================
actual = (1-scarcity^0.35) * potential

*/

type OreStop struct {
	ID      string
	Stop    float64
	Texture string
}

func GetOreStops() []OreStop {
	o := make([]OreStop, 0)

	o = append(o, OreStop{
		ID:   "dummy",
		Stop: 0,
	})

	o = append(o, OreStop{
		ID:      "dd522f03-2f52-4e82-b2f8-d7e0029cb82f",
		Stop:    0.1875,
		Texture: "Mega/asteroidR4",
	})

	o = append(o, OreStop{
		ID:      "56617d30-6c30-425c-84bf-2484ae8c1156",
		Stop:    0.3618,
		Texture: "Mega/asteroidR2",
	})

	o = append(o, OreStop{
		ID:      "26a3fc9e-db2f-439d-a929-ba755d11d09c",
		Stop:    0.5227,
		Texture: "Mega/asteroidR6",
	})

	o = append(o, OreStop{
		ID:      "1d0d344b-ef28-43c8-a7a6-3275936b2dea",
		Stop:    0.6070,
		Texture: "Mega/asteroidR3",
	})

	o = append(o, OreStop{
		ID:      "0cd04eea-a150-410c-91eb-6af00d8c6eae",
		Stop:    0.6684,
		Texture: "Mega/asteroidR9",
	})

	o = append(o, OreStop{
		ID:      "39b8eedf-ef80-4c29-a4bf-99abc4d84fa6",
		Stop:    0.7216,
		Texture: "Mega/asteroidR7",
	})

	o = append(o, OreStop{
		ID:      "dd0c9b0a-279e-418e-b3b6-2f569fda0186",
		Stop:    0.7500,
		Texture: "Mega/asteroidR1",
	})

	o = append(o, OreStop{
		ID:      "7dcd5138-d7e0-419f-867a-6f0f23b99b5b",
		Stop:    0.8333,
		Texture: "Mega/asteroidR12",
	})

	o = append(o, OreStop{
		ID:      "61f52ba3-654b-45cf-88e3-33399d12350d",
		Stop:    0.8954,
		Texture: "Mega/asteroidR5",
	})

	o = append(o, OreStop{
		ID:      "11688112-f3d4-4d30-864a-684a8b96ea23",
		Stop:    0.9336,
		Texture: "Mega/asteroidR11",
	})

	o = append(o, OreStop{
		ID:      "2ce48bef-f06b-4550-b20c-0e64864db051",
		Stop:    0.9634,
		Texture: "Mega/asteroidR8",
	})

	o = append(o, OreStop{
		ID:      "66b7a322-8cfc-4467-9410-492e6b58f159",
		Stop:    0.9865,
		Texture: "Mega/asteroidR13",
	})

	o = append(o, OreStop{
		ID:      "d1866be4-5c3e-4b95-b6d9-020832338014",
		Stop:    1.0000,
		Texture: "Mega/asteroidR10",
	})

	return o
}

// Generates a seed integer for use as a system-specific RNG seed for consistent internal generation
func calculateSystemSeed(s *universe.SolarSystem) int {
	// use the system's position and uuid timestamp as a seed
	seed := (int(s.PosX*10000.0)>>int(math.Abs(s.PosY)*10000.0) + s.ID.ClockSequence())

	if s.PosY < 0 {
		seed *= -1
	}

	return seed
}

// Represents a schematic with inputs and outputs to be saved by worldfiller
type schematicStubWorldmaker struct {
	// item type
	ItemType sql.VwItemTypeIndustrial
	// fair process (item type, used by schematic)
	FairProcess sql.Process
	FairInputs  []sql.ProcessInput
	FairOutputs []sql.ProcessOutput
	// sink process (item type)
	SinkProcess       sql.Process
	ItemTypeSinkInput sql.ProcessInput
	// new schematic item type
	NewSchematic sql.ItemType
	// schematic item faucet
	SchematicFaucetProcess       sql.Process
	SchematicFaucetProcessOutput sql.ProcessOutput
	// schematic item sink
	SchematicSinkProcess      sql.Process
	SchematicSinkProcessInput sql.ProcessInput
}

// Generates schematics and processes for modules
func stubModuleSchematicsAndProcesses() {
	rand.Seed(time.Now().Unix())
	generated := make([]schematicStubWorldmaker, 0)

	// get services
	itemTypeSvc := sql.GetItemTypeService()
	processSvc := sql.GetProcessService()
	processInputSvc := sql.GetProcessInputService()
	processOutputSvc := sql.GetProcessOutputService()

	// load item type views
	needSchematics, _ := itemTypeSvc.GetVwModuleNeedSchematics()
	industrials, _ := itemTypeSvc.GetVwItemTypeIndustrials()

	// group industrials
	ores := make([]sql.VwItemTypeIndustrial, 0)
	ices := make([]sql.VwItemTypeIndustrial, 0)
	gases := make([]sql.VwItemTypeIndustrial, 0)
	wastes := make([]sql.VwItemTypeIndustrial, 0)
	repairKits := make([]sql.VwItemTypeIndustrial, 0)

	var smallPowerCell sql.VwItemTypeIndustrial
	var smallDepletedPowerCell sql.VwItemTypeIndustrial
	var testite sql.VwItemTypeIndustrial

	for _, i := range industrials {
		if i.Family == "ore" && i.Name != "Betro" {
			if i.Name == "Testite" {
				testite = i
			} else {
				ores = append(ores, i)
			}
		}

		if i.Family == "ice" {
			ices = append(ices, i)
		}

		if i.Family == "gas" {
			gases = append(gases, i)
		}

		if i.Family == "trade_good" && strings.Contains(i.Name, "Waste") {
			wastes = append(wastes, i)
		}

		if i.Family == "repair_kit" {
			repairKits = append(repairKits, i)
		}

		if i.Name == "10 kWH Cell" {
			smallPowerCell = i
		}

		if i.Name == "Depleted 10 kWH Cell" {
			smallDepletedPowerCell = i
		}
	}

	// iterate over item types to stub for
	for _, i := range needSchematics {
		var industrial *sql.VwItemTypeIndustrial
		var fairProcess = sql.Process{
			ID:   uuid.New(),
			Name: fmt.Sprintf("Make %v", i.Name),
			Meta: sql.Meta{},
		}

		var fairInputs = make(map[string]int)
		var fairOutputs = make(map[string]int)

		var fairInputsSql = make([]sql.ProcessInput, 0)
		var fairOutputsSql = make([]sql.ProcessOutput, 0)

		// get industrial data for item type
		for _, j := range industrials {
			if j.ID == i.ID {
				industrial = &j
				break
			}
		}

		// skip dev stuff
		if industrial.Name == "Le Banhammer" {
			continue
		}

		// select ores
		inputMaterials := ores

		rand.Shuffle(len(inputMaterials), func(i, j int) {
			q := inputMaterials[i]
			r := inputMaterials[j]

			inputMaterials[i] = r
			inputMaterials[j] = q
		})

		inputMaterials = inputMaterials[:physics.RandInRange(3, 9)]
		inputMaterials = append(inputMaterials, testite)

		// select any additional materials (excluding power cells)
		if industrial.Volume >= 100 {
			roll := physics.RandInRange(0, 100)

			if roll <= 4 {
				inputMaterials = append(inputMaterials, ices[physics.RandInRange(0, len(ices))])
			}

			if roll <= 2 {
				inputMaterials = append(inputMaterials, gases[physics.RandInRange(0, len(gases))])
			}

			roll = physics.RandInRange(0, 100)

			if roll <= 7 {
				inputMaterials = append(inputMaterials, repairKits[physics.RandInRange(0, len(repairKits))])
			}

			roll = physics.RandInRange(0, 100)

			if roll <= 10 {
				inputMaterials = append(inputMaterials, wastes[physics.RandInRange(0, len(wastes))])
			}
		}

		// determine energy cost : material cost ratio
		costRatio := rand.Float64()

		if costRatio < 0.2 {
			costRatio = 0.2
		}

		if costRatio > 0.6 {
			costRatio = 0.6
		}

		// determine quantities
		ceil := int(industrial.SiloSize) / 10

		if ceil <= 1 {
			ceil = 10
		}

		outputQuantity := physics.RandInRange(1, ceil)

		if outputQuantity > 10 {
			rem := outputQuantity % 5
			outputQuantity -= rem
		}

		jobTime := physics.RandInRange(1, (int(industrial.Volume)+1)*outputQuantity*2)
		outputVolume := outputQuantity * int(industrial.Volume)
		fairProcess.Time = jobTime

		// determine costs
		energyCost := ((costRatio * industrial.MinPrice) * 0.9) * float64(outputQuantity)
		materialCost := (((1 - costRatio) * industrial.MinPrice) * 0.9) * float64(outputQuantity)

		// add input/output for energy cost
		cellsUsed := 0

		for {
			cellsUsed++

			if smallPowerCell.MaxPrice*float64(cellsUsed) >= energyCost {
				break
			}
		}

		fairInputs[smallPowerCell.Name] = cellsUsed
		fairOutputs[smallDepletedPowerCell.Name] = cellsUsed
		fairOutputs[industrial.Name] = outputQuantity

		// determine material input quantities
		matMap := make(map[string]sql.VwItemTypeIndustrial)

		for _, e := range inputMaterials {
			matMap[e.Name] = e
			fairInputs[e.Name] = 0
		}

		resets := 0

		for {
			cost := 0.0
			volume := 0.0

			for i, v := range fairInputs {
				if i == smallPowerCell.Name {
					continue
				}

				// roll dice
				roll := physics.RandInRange(1, 100)

				if roll > 75 || i == "Testite" {
					// increment quantity
					fairInputs[i] = v + 1
					v = fairInputs[i]
				}

				// accumulate
				cost += matMap[i].MaxPrice * float64(v)
				volume += matMap[i].Volume * float64(v)
			}

			if cost >= materialCost*1.1 || volume >= float64(outputVolume)*10 {
				// reset
				for i := range fairInputs {
					if i == smallPowerCell.Name {
						continue
					}

					fairInputs[i] = 0
				}

				resets++
			} else {
				if cost > materialCost && volume > float64(outputVolume)*0.5 {
					// all done!
					break
				}
			}

			if resets > 100 {
				log.Printf("Skipping %v - not solvable with input materials given.", industrial.Name)
				break
			}
		}

		if resets > 100 {
			continue
		}

		// build io for fair process
		for k, v := range fairInputs {
			if v <= 0 {
				continue
			}

			for _, l := range industrials {
				if l.Name == k {
					fairInputsSql = append(fairInputsSql, sql.ProcessInput{
						ID:         uuid.New(),
						ItemTypeID: l.ID,
						Quantity:   v,
						Meta:       sql.Meta{},
						ProcessID:  fairProcess.ID,
					})

					break
				}
			}
		}

		for k, v := range fairOutputs {
			if v <= 0 {
				continue
			}

			for _, l := range industrials {
				if l.Name == k {
					fairOutputsSql = append(fairOutputsSql, sql.ProcessOutput{
						ID:         uuid.New(),
						ItemTypeID: l.ID,
						Quantity:   v,
						Meta:       sql.Meta{},
						ProcessID:  fairProcess.ID,
					})

					break
				}
			}
		}

		// make sink process
		sinkProcess := sql.Process{
			ID:   uuid.New(),
			Name: fmt.Sprintf("%v Sink [wm]", industrial.Name),
			Time: physics.RandInRange(fairProcess.Time/2, fairProcess.Time*2),
			Meta: sql.Meta{},
		}

		sinkInput := sql.ProcessInput{
			ID:         uuid.New(),
			ItemTypeID: industrial.ID,
			Quantity:   physics.RandInRange(outputQuantity/2, outputQuantity*2),
			Meta:       sql.Meta{},
			ProcessID:  sinkProcess.ID,
		}

		// new schematic item type
		newSchematic := sql.ItemType{
			ID:     uuid.New(),
			Family: "schematic",
			Name:   fmt.Sprintf("%v Schematic", industrial.Name),
			Meta:   sql.Meta{},
		}

		maxSP := physics.RandInRange(int(energyCost)/2, int(materialCost)*2)
		minSP := physics.RandInRange(maxSP/8, maxSP/2)

		meta := universe.IndustrialMetadata{
			SiloSize:  100,
			MaxPrice:  maxSP,
			MinPrice:  minSP,
			ProcessID: &fairProcess.ID,
		}

		newSchematic.Meta["industrialmarket"] = meta

		// make schematic faucet process
		schematicFaucet := sql.Process{
			ID:   uuid.New(),
			Name: fmt.Sprintf("%v Faucet [wm]", newSchematic.Name),
			Time: physics.RandInRange(fairProcess.Time*2, fairProcess.Time*8),
			Meta: sql.Meta{},
		}

		schematicFaucetOutput := sql.ProcessOutput{
			ID:         uuid.New(),
			ItemTypeID: newSchematic.ID,
			Quantity:   physics.RandInRange(1, 10),
			Meta:       sql.Meta{},
			ProcessID:  schematicFaucet.ID,
		}

		// make schematic sink process
		schematicSink := sql.Process{
			ID:   uuid.New(),
			Name: fmt.Sprintf("%v Sink [wm]", newSchematic.Name),
			Time: physics.RandInRange(fairProcess.Time*2, fairProcess.Time*8),
			Meta: sql.Meta{},
		}

		schematicSinkInput := sql.ProcessInput{
			ID:         uuid.New(),
			ItemTypeID: newSchematic.ID,
			Quantity:   physics.RandInRange(1, 10),
			Meta:       sql.Meta{},
			ProcessID:  schematicSink.ID,
		}

		// store for saving
		generated = append(generated, schematicStubWorldmaker{
			ItemType:                     *industrial,
			FairProcess:                  fairProcess,
			FairInputs:                   fairInputsSql,
			FairOutputs:                  fairOutputsSql,
			SinkProcess:                  sinkProcess,
			ItemTypeSinkInput:            sinkInput,
			NewSchematic:                 newSchematic,
			SchematicFaucetProcess:       schematicFaucet,
			SchematicFaucetProcessOutput: schematicFaucetOutput,
			SchematicSinkProcess:         schematicSink,
			SchematicSinkProcessInput:    schematicSinkInput,
		})

		// success :)
		log.Printf(":) %v", industrial.Name)
	}

	// save everything
	for _, e := range generated {
		// save processes
		o, err := processSvc.NewProcessForWorldFiller(e.FairProcess)

		if o == nil || err != nil {
			panic(fmt.Sprintf("error saving process: %v | %v", o, err))
		}

		o, err = processSvc.NewProcessForWorldFiller(e.SinkProcess)

		if o == nil || err != nil {
			panic(fmt.Sprintf("error saving process: %v | %v", o, err))
		}

		o, err = processSvc.NewProcessForWorldFiller(e.SchematicFaucetProcess)

		if o == nil || err != nil {
			panic(fmt.Sprintf("error saving process: %v | %v", o, err))
		}

		o, err = processSvc.NewProcessForWorldFiller(e.SchematicSinkProcess)

		if o == nil || err != nil {
			panic(fmt.Sprintf("error saving process: %v | %v", o, err))
		}

		// save schematic item type
		s, err := itemTypeSvc.NewItemTypeForWorldFiller(e.NewSchematic)

		if s == nil || err != nil {
			panic(fmt.Sprintf("error saving schematic: %v | %v", o, err))
		}

		// coalesce inputs and outputs
		allInputs := make([]sql.ProcessInput, 0)
		allOuts := make([]sql.ProcessOutput, 0)

		allInputs = append(allInputs, e.SchematicSinkProcessInput)
		allInputs = append(allInputs, e.ItemTypeSinkInput)
		allInputs = append(allInputs, e.FairInputs...)

		allOuts = append(allOuts, e.SchematicFaucetProcessOutput)
		allOuts = append(allOuts, e.FairOutputs...)

		// save inputs and outputs
		for _, x := range allInputs {
			b, err := processInputSvc.NewProcessInputForWorldFiller(x)

			if b == nil || err != nil {
				panic(fmt.Sprintf("error saving input: %v | %v", b, err))
			}
		}

		for _, x := range allOuts {
			b, err := processOutputSvc.NewProcessOutputForWorldFiller(x)

			if b == nil || err != nil {
				panic(fmt.Sprintf("error saving output: %v | %v", b, err))
			}
		}
	}
}

// Creates station processes for a given process
func injectProcess(u *universe.Universe, pid string, offset int) {
	pID, err := uuid.Parse(pid)
	prob := 3

	stationProcessSvc := sql.GetStationProcessService()

	if err != nil {
		panic(err)
	}

	var textures = [...]string{
		"sanctuary-",
		"interstar-",
	}

	toSave := make([]sql.StationProcess, 0)

	for _, r := range u.Regions {
		for _, s := range r.Systems {
			rand.Seed(int64(calculateSystemSeed(s)) - 178046782372 + int64(offset))

			/*if r.ID.ID()%2 != 0 {
				continue
			}

			if s.ID.ID()%5 != 0 {
				continue
			}*/

			stations := s.CopyStations(true)

			for _, st := range stations {
				for _, t := range textures {
					if !strings.Contains(st.Texture, t) {
						roll := physics.RandInRange(0, 100)

						if roll <= prob {
							sp := sql.StationProcess{
								StationID: st.ID,
								ProcessID: pID,
								ID:        uuid.New(),
							}

							toSave = append(toSave, sp)
						}
					}
				}
			}
		}
	}

	for _, o := range toSave {
		err := stationProcessSvc.NewStationProcessWorldMaker(&o)

		if err != nil {
			panic(err)
		}
	}
}

// Creates ancient signal relays
func dropArtifacts(u *universe.Universe) {
	toSave := make([]sql.Artifact, 0)

	for _, r := range u.Regions {
		for _, s := range r.Systems {
			// roll to skip
			rts := rand.Float64() * rand.Float64() * rand.Float64()
			rqq := 0.75

			if s.HoldingFactionID == uuid.MustParse("bdeffd9a-3cab-408c-9cd7-32fce1124f7a") {
				rqq -= 0.5
			}

			if s.HoldingFactionID == uuid.MustParse("a0152cbb-8a78-45a4-ae9b-2ad2b60b583b") {
				rqq -= 0.8
			}

			if s.HoldingFactionID == uuid.MustParse("27a53dfc-a321-4c12-bf7c-bb177955c95b") {
				rqq += 0.5
			}

			if s.HoldingFactionID == uuid.MustParse("5db2bec7-37c3-4f1c-ab88-21024c12d639") {
				rqq += 0.8
			}

			if rts < rqq {
				continue
			}

			rand.Seed(int64(calculateSystemSeed(s)))

			stars := s.CopyStars(true)
			planets := s.CopyPlanets(true)
			jumpholes := s.CopyJumpholes(true)
			asteroids := s.CopyAsteroids(true)
			stations := s.CopyStations(true)

			// build artifact for the system
			art := sql.Artifact{
				ID:           uuid.New(),
				SystemID:     s.ID,
				ArtifactName: fmt.Sprintf("Original Relay %v", randomPlaceholderName()),
				Texture:      "signalrelay",
				Radius:       775 - (75 * rand.Float64()),
				Mass:         89763 - (125 * rand.Float64()),
				Theta:        float64(physics.RandInRange(0, 360)),
			}

			// get star position
			sx := 0.0
			sy := 0.0

			for _, q := range stars {
				sx = q.PosX
				sy = q.PosY

				break
			}

			// determine position
			for {
				// generate random position in belt
				dW := float64(physics.RandInRange(1000, 10000))
				mag := dW
				the := 2.0 * math.Pi * rand.Float64()

				art.PosX = (mag * math.Cos(the)) + sx
				art.PosY = (mag * math.Sin(the)) + sy

				// check for overlap with forbidden objects
				sB := physics.Dummy{
					PosX: art.PosX,
					PosY: art.PosY,
				}

				for _, v := range stars {
					sA := physics.Dummy{
						PosX: v.PosX,
						PosY: v.PosY,
					}

					dst := physics.Distance(sA, sB)

					if dst < (1+rand.Float64())*(v.Radius+art.Radius) {
						// not safe, try again
						continue
					}
				}

				for _, v := range planets {
					sA := physics.Dummy{
						PosX: v.PosX,
						PosY: v.PosY,
					}

					dst := physics.Distance(sA, sB)

					if dst < (1+rand.Float64())*(v.Radius+art.Radius) {
						// not safe, try again
						continue
					}
				}

				for _, v := range stations {
					sA := physics.Dummy{
						PosX: v.PosX,
						PosY: v.PosY,
					}

					dst := physics.Distance(sA, sB)

					if dst < (1+rand.Float64())*(v.Radius+art.Radius) {
						// not safe, try again
						continue
					}
				}

				for _, v := range jumpholes {
					sA := physics.Dummy{
						PosX: v.PosX,
						PosY: v.PosY,
					}

					dst := physics.Distance(sA, sB)

					if dst < (1+rand.Float64())*(v.Radius+art.Radius) {
						// not safe, try again
						continue
					}
				}

				for _, v := range asteroids {
					sA := physics.Dummy{
						PosX: v.PosX,
						PosY: v.PosY,
					}

					dst := physics.Distance(sA, sB)

					if dst < (1+rand.Float64())*(v.Radius+art.Radius) {
						// not safe, try again
						continue
					}
				}

				// safe to store artifact
				break
			}

			toSave = append(toSave, art)
		}
	}

	// get service
	artifactSvc := sql.GetArtifactService()

	// save new artifacts
	for _, st := range toSave {
		err := artifactSvc.NewArtifactWorldFiller(&st)

		if err != nil {
			panic(err)
		}
	}
}

// Creates stations owned by the sanctuary systems
func dropSanctuaryStations(u *universe.Universe) {
	fID, err := uuid.Parse("b3d3fa9c-b21e-490f-b39e-128b3af12128")

	if err != nil {
		panic(err)
	}

	toSave := make([]sql.Station, 0)

	for _, r := range u.Regions {
		for _, s := range r.Systems {
			rand.Seed(int64(calculateSystemSeed(s)))

			// is owned by sanctuary?
			if s.HoldingFactionID == fID {
				stars := s.CopyStars(true)
				planets := s.CopyPlanets(true)
				jumpholes := s.CopyJumpholes(true)
				asteroids := s.CopyAsteroids(true)
				stations := s.CopyStations(true)

				// build sanctuary station for the system
				stat := sql.Station{
					ID:          uuid.New(),
					SystemID:    s.ID,
					StationName: "Sanctuary Outpost",
					Texture:     "sanctuary-1",
					Radius:      525,
					Mass:        23700,
					Theta:       float64(physics.RandInRange(0, 360)),
					FactionID:   fID,
				}

				// determine position
				for {
					// generate random position in belt
					dW := float64(physics.RandInRange(25000, 75000))
					mag := dW
					the := 2.0 * math.Pi * rand.Float64()

					stat.PosX = (mag * math.Cos(the))
					stat.PosY = (mag * math.Sin(the))

					// check for overlap with forbidden objects
					sB := physics.Dummy{
						PosX: stat.PosX,
						PosY: stat.PosY,
					}

					for _, v := range stars {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					for _, v := range planets {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					for _, v := range stations {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					for _, v := range jumpholes {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					for _, v := range asteroids {
						sA := physics.Dummy{
							PosX: v.PosX,
							PosY: v.PosY,
						}

						dst := physics.Distance(sA, sB)

						if dst < (1+rand.Float64())*(v.Radius+stat.Radius) {
							// not safe, try again
							continue
						}
					}

					// safe to store asteroid
					break
				}

				toSave = append(toSave, stat)
			}
		}
	}

	// get service
	stationSvc := sql.GetStationService()

	// save new stations
	for _, st := range toSave {
		err := stationSvc.NewStationWorldFiller(&st)

		if err != nil {
			panic(err)
		}
	}
}

// Inserts minable asteroids into the universe
func dropAsteroids(u *universe.Universe) {
	globalAsteroids := make([]Astling, 0)

	for _, r := range u.Regions {
		for _, s := range r.Systems {
			// get system internal seed
			seed := calculateSystemSeed(s)

			// introduce unique offset for this function
			seed = seed + 50912

			// initialize RNG with seed
			rand.Seed(int64(seed))

			// get scarcity level
			scarcity := rand.Float64()

			// get obstacles in system
			stars := s.CopyStars(true)
			planets := s.CopyPlanets(true)
			jumpholes := s.CopyJumpholes(true)

			// determine total number of asteroids in system
			potentialAsteroids := physics.RandInRange(MinAsteroidsPerSystem, MaxAsteroidsPerSystem)
			actualAsteroids := int((1.0 - math.Pow(scarcity, 0.35)) * float64(potentialAsteroids))

			// determine asteroid distribution scale
			beltRadius := physics.RandInRange(MinBeltRadius, MaxBeltRadius)

			// use first star as belt center
			beltX := 0.0
			beltY := 0.0

			for _, e := range stars {
				beltX = e.PosX
				beltY = e.PosY
			}

			// get ore stops
			stops := GetOreStops()

			// generated localAsteroids
			localAsteroids := make([]Astling, 0)

			// make asteroids
			for i := 0; i < actualAsteroids; i++ {
				// roll to determine yield
				y := math.Max(rand.Float64()*MaxAsteroidYield, MinAsteroidYield) * (1 - scarcity)

				// roll to determine ore type
				p := rand.Float64()

				for x := range stops {
					// skip dummy stop
					if x == 0 {
						continue
					}

					// determine if this stop captures the roll
					floor := stops[x-1]
					ceiling := stops[x]

					if floor.Stop < p && p <= ceiling.Stop {
						ast := Astling{
							ID:            uuid.New().String(),
							SystemID:      s.ID.String(),
							OreItemTypeID: stops[x].ID,
							Name:          fmt.Sprintf("A%v%v%v-%v", stops[x].ID[0], stops[x].ID[1], stops[x].ID[2], physics.RandInRange(1000, 9999)),
							Texture:       stops[x].Texture,
							Radius:        float64(physics.RandInRange(MinAsteroidRadius, MaxAsteroidRadius)),
							Theta:         float64(rand.Float64() * 360.0),
							Yield:         y,
							Mass:          float64(physics.RandInRange(MinAsteroidMass, MaxAsteroidMass)),
						}

						// determine position
						for {
							// generate random position in belt
							dW := float64(physics.RandInRange(MinBeltWidth, MaxBeltWidth))
							mag := dW + float64(beltRadius)
							the := 2.0 * math.Pi * rand.Float64()

							ast.PosX = (mag * math.Cos(the)) + beltX
							ast.PosY = (mag * math.Sin(the)) + beltY

							// check for overlap with forbidden objects
							sB := physics.Dummy{
								PosX: ast.PosX,
								PosY: ast.PosY,
							}

							for _, v := range stars {
								sA := physics.Dummy{
									PosX: v.PosX,
									PosY: v.PosY,
								}

								dst := physics.Distance(sA, sB)

								if dst < (1+rand.Float64())*(v.Radius+ast.Radius) {
									// not safe, try again
									continue
								}
							}

							for _, v := range planets {
								sA := physics.Dummy{
									PosX: v.PosX,
									PosY: v.PosY,
								}

								dst := physics.Distance(sA, sB)

								if dst < (1+rand.Float64())*(v.Radius+ast.Radius) {
									// not safe, try again
									continue
								}
							}

							for _, v := range jumpholes {
								sA := physics.Dummy{
									PosX: v.PosX,
									PosY: v.PosY,
								}

								dst := physics.Distance(sA, sB)

								if dst < (1+rand.Float64())*(v.Radius+ast.Radius) {
									// not safe, try again
									continue
								}
							}

							for _, v := range localAsteroids {
								sA := physics.Dummy{
									PosX: v.PosX,
									PosY: v.PosY,
								}

								dst := physics.Distance(sA, sB)

								if dst < (1+rand.Float64())*(v.Radius+ast.Radius) {
									// not safe, try again
									continue
								}
							}

							// safe to store asteroid
							break
						}

						// store asteroid
						localAsteroids = append(localAsteroids, ast)

						// next asteroid
						break
					}
				}
			}

			// store asteroids to save
			globalAsteroids = append(globalAsteroids, localAsteroids...)
		}
	}

	// save generated asteroids
	astSvc := sql.GetAsteroidService()

	for _, ast := range globalAsteroids {
		id, err := uuid.Parse(ast.ID)

		if err != nil {
			panic(err)
		}

		systemID, err := uuid.Parse(ast.SystemID)

		if err != nil {
			panic(err)
		}

		oreID, err := uuid.Parse(ast.OreItemTypeID)

		if err != nil {
			panic(err)
		}

		f := sql.Asteroid{
			ID:         id,
			SystemID:   systemID,
			ItemTypeID: oreID,
			Name:       ast.Name,
			Texture:    ast.Texture,
			Radius:     ast.Radius,
			Theta:      ast.Theta,
			PosX:       ast.PosX,
			PosY:       ast.PosY,
			Yield:      ast.Yield,
			Mass:       ast.Mass,
		}

		err = astSvc.NewAsteroidWorldFiller(&f)

		if err != nil {
			panic(err)
		}
	}
}

// Structure representing a scaffold for an asteroid for worldfiller
type Astling struct {
	ID            string
	SystemID      string
	OreItemTypeID string
	Name          string
	Texture       string
	Radius        float64
	Theta         float64
	PosX          float64
	PosY          float64
	Yield         float64
	Mass          float64
}

func printLogger(s string, t time.Time) {
	log.Println(s) // t is intentionally discarded because Println already provides a timestamp
}

// Structure representing a scaffold for a trade good for worldfiller
type WareCsvRecord struct {
	Name     string
	MinPrice int
	MaxPrice int
	SiloSize int
	Volume   int
}

// Imports new trade goods from a CSV into the database for worldfiller
func loadNewWares() {
	// open file
	f, err := os.Open("newwares.csv")

	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	wares := parseWareCsv(data)

	// print the array
	fmt.Printf("%+v\n", wares)

	// get services
	itemTypeSvc := sql.GetItemTypeService()
	processSvc := sql.GetProcessService()
	inputSvc := sql.GetProcessInputService()
	outputSvc := sql.GetProcessOutputService()

	// to store item types to be saved
	newTypes := make([]sql.ItemType, 0)

	// iterate over parsed wares
	for _, w := range wares {
		// verify this ware doesn't already exist
		existing, err := itemTypeSvc.GetItemTypeByName(w.Name)

		if err == nil || existing != nil {
			log.Panicf("Ware already exists! %v", existing)
		}

		// create new item type for ware
		it := sql.ItemType{
			ID:     uuid.New(),
			Family: "gas",
			Name:   w.Name,
			Meta:   sql.Meta{},
		}

		meta := universe.IndustrialMetadataNoProcessId{
			SiloSize: w.SiloSize,
			MaxPrice: w.MaxPrice,
			MinPrice: w.MinPrice,
		}

		it.Meta["hp"] = 1
		it.Meta["volume"] = w.Volume
		it.Meta["industrialmarket"] = meta

		// append for later
		newTypes = append(newTypes, it)
	}

	// save new wares
	for _, it := range newTypes {
		log.Printf("Saving %v", it)

		// save core ware
		_, err := itemTypeSvc.NewItemTypeForWorldFiller(it)

		if err != nil {
			log.Panicf("Error saving new item type %v", err)
		}

		// calculate runtimes
		rtf := physics.RandInRange(5, 1000)
		rts := physics.RandInRange(7, 1100)

		// make ware faucet process
		wareFaucet := sql.Process{
			ID:   uuid.New(),
			Name: fmt.Sprintf("%v Faucet [wm]", it.Name),
			Time: physics.RandInRange(rtf, rtf*8),
			Meta: sql.Meta{},
		}

		wareFaucetOutput := sql.ProcessOutput{
			ID:         uuid.New(),
			ItemTypeID: it.ID,
			Quantity:   physics.RandInRange(1, 1000),
			Meta:       sql.Meta{},
			ProcessID:  wareFaucet.ID,
		}

		// make ware sink process
		wareSink := sql.Process{
			ID:   uuid.New(),
			Name: fmt.Sprintf("%v Sink [wm]", it.Name),
			Time: physics.RandInRange(rts, rts*8),
			Meta: sql.Meta{},
		}

		wareSinkInput := sql.ProcessInput{
			ID:         uuid.New(),
			ItemTypeID: it.ID,
			Quantity:   physics.RandInRange(1, 1000),
			Meta:       sql.Meta{},
			ProcessID:  wareSink.ID,
		}

		// save faucet process
		_, err = processSvc.NewProcessForWorldFiller(wareFaucet)

		if err != nil {
			log.Panicf("Error saving faucet process %v", err)
		}

		_, err = outputSvc.NewProcessOutputForWorldFiller(wareFaucetOutput)

		if err != nil {
			log.Panicf("Error saving faucet process (output) %v", err)
		}

		// save sink process
		_, err = processSvc.NewProcessForWorldFiller(wareSink)

		if err != nil {
			log.Panicf("Error saving sink process %v", err)
		}

		_, err = inputSvc.NewProcessInputForWorldFiller(wareSinkInput)

		if err != nil {
			log.Panicf("Error saving faucet process (input) %v", err)
		}
	}
}

// Parses new ware CSV for worldfiller
func parseWareCsv(data [][]string) []WareCsvRecord {
	var wareList []WareCsvRecord
	for i, line := range data {
		if i > 0 { // omit header line
			var rec WareCsvRecord

			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 2 {
					k, _ := strconv.Atoi(field)
					rec.MinPrice = k
				} else if j == 3 {
					k, _ := strconv.Atoi(field)
					rec.MaxPrice = k
				} else if j == 4 {
					k, _ := strconv.Atoi(field)
					rec.SiloSize = k
				} else if j == 5 {
					k, _ := strconv.Atoi(field)
					rec.Volume = k
				}
			}

			wareList = append(wareList, rec)
		}
	}
	return wareList
}

// Initializes gas mining yields on celestials
func fillGasMiningYields(u *universe.Universe) {
	/* asteroid gas meta */
	const astHasGasProb = 15

	var astGases = [...]string{
		"22fb3cda-c949-41ae-bcf5-ee0a60d497fc",
		"5f84addf-d05b-407a-8dcd-fecd3af4c69d",
		"42179f25-b473-46d1-9592-08e1d23807bd",
	}

	/* planet gas meta */
	const plnHasGasProb = 66

	var plnGases = [...]string{
		"3f903eca-8bd3-4f54-a9ce-2b22627db94a",
		"04f16f34-497f-465e-a9a6-368dd62a5328",
		"d8e1ead1-055b-4cd0-b1ec-f88dd0f7ff31",
	}

	/* star gas meta */
	const starHasGasProb = 85

	var strGases = [...]string{
		"f4947ebe-4d4c-457b-aeb6-b4dd5b66e62b",
		"0c429117-e7fa-4e00-97f9-ab67c639cae7",
		"479e8dd7-06e8-479c-b07f-382b407b832f",
	}

	// get services
	asteroidSvc := sql.GetAsteroidService()
	planetSvc := sql.GetPlanetService()
	starSvc := sql.GetStarService()

	// iterate over regions
	for _, r := range u.Regions {
		// progress
		log.Printf("seeding region %v", r.RegionName)

		// roll scarcity
		rScarcity := rand.Float64()

		// iterate over solar systems
		for _, s := range r.Systems {
			// roll scarcity
			sScarcity := rand.Float64()

			// iterate over asteroids
			asts := s.CopyAsteroids(false)

			for _, a := range asts {
				// roll scarcity
				aScarcity := rand.Float64()

				// empty gas mining meta
				gmm := universe.GasMiningMetadata{
					Yields: make(map[string]universe.GasMiningYield),
				}

				// roll for gas presence
				hasGasRoll := physics.RandInRange(0, 100)

				if hasGasRoll <= astHasGasProb {
					// iterate over asteroid gases
					for _, gid := range astGases {
						// roll for yield
						yld := physics.RandInRange(0, int(a.Mass)/int(a.Radius))

						// apply scarcity
						scarcity := rScarcity * sScarcity * aScarcity
						yld = int(float64(yld) * scarcity)

						if yld > 0 {
							// add entry
							gmm.Yields[gid] = universe.GasMiningYield{
								ItemTypeID: uuid.MustParse(gid),
								Yield:      yld,
							}
						}
					}
				}

				// save metadata
				meta := a.Meta
				meta["gasmining"] = gmm

				asteroidSvc.UpdateMetaWorldfiller(a.ID, (*sql.Meta)(&meta))
			}

			// iterate over planets
			pls := s.CopyPlanets(false)

			for _, p := range pls {
				// roll scarcity
				aScarcity := rand.Float64()

				// empty gas mining meta
				gmm := universe.GasMiningMetadata{
					Yields: make(map[string]universe.GasMiningYield),
				}

				// roll for gas presence
				hasGasRoll := physics.RandInRange(0, 100)

				if hasGasRoll <= plnHasGasProb {
					// iterate over planet gases
					for _, gid := range plnGases {
						// roll for yield
						yld := physics.RandInRange(0, int(p.Mass)/int(p.Radius))

						// apply scarcity
						scarcity := rScarcity * sScarcity * aScarcity
						yld = int(float64(yld) * scarcity)

						if yld > 0 {
							// add entry
							gmm.Yields[gid] = universe.GasMiningYield{
								ItemTypeID: uuid.MustParse(gid),
								Yield:      yld,
							}
						}
					}
				}

				// save metadata
				meta := p.Meta
				meta["gasmining"] = gmm

				planetSvc.UpdateMetaWorldfiller(p.ID, (*sql.Meta)(&meta))
			}

			// iterate over stars
			stars := s.CopyStars(false)

			for _, st := range stars {
				// roll scarcity
				aScarcity := rand.Float64()

				// empty gas mining meta
				gmm := universe.GasMiningMetadata{
					Yields: make(map[string]universe.GasMiningYield),
				}

				// roll for gas presence
				hasGasRoll := physics.RandInRange(0, 100)

				if hasGasRoll <= starHasGasProb {
					// iterate over star gases
					for _, gid := range strGases {
						// roll for yield
						yld := physics.RandInRange(0, int(st.Mass)/int(st.Radius))

						// apply scarcity
						scarcity := rScarcity * sScarcity * aScarcity
						yld = int(float64(yld) * scarcity)

						if yld > 0 {
							// add entry
							gmm.Yields[gid] = universe.GasMiningYield{
								ItemTypeID: uuid.MustParse(gid),
								Yield:      yld,
							}
						}
					}
				}

				// save metadata
				meta := st.Meta
				meta["gasmining"] = gmm

				starSvc.UpdateMetaWorldfiller(st.ID, (*sql.Meta)(&meta))
			}
		}
	}
}

func randomPlaceholderName() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := []rune("1234567890")

	acc := ""

	firstLength := physics.RandInRange(1, 5)
	secondLength := physics.RandInRange(1, 5)

	for i := 0; i < firstLength; i++ {
		idx := physics.RandInRange(0, len(letters))
		acc = fmt.Sprintf("%v%v", acc, string(letters[idx]))
	}

	acc = fmt.Sprintf("%v%v", acc, "-")

	for i := 0; i < secondLength; i++ {
		idx := physics.RandInRange(0, len(numbers))
		acc = fmt.Sprintf("%v%v", acc, string(numbers[idx]))
	}

	return acc
}
