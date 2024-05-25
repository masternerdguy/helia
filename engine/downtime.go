package engine

import (
	"helia/shared"
	"helia/sql"
	"time"
)

var DownTimeHandled = false
var DownTimeInitialized = false

// Structure representing tasks to perform as part of the daily downtime before engine start
type DownTimeRunner struct {
	spSvc      sql.SPService
	userSvc    sql.UserService
	factionSvc sql.FactionService
	startSvc   sql.StartService
}

// Initializes downtime job structure
func (d *DownTimeRunner) Initialize() {
	// safety check
	if DownTimeHandled || DownTimeInitialized {
		panic("second attempt at initializing downtime jobs!")
	}

	// get services
	d.spSvc = sql.GetSPService()
	d.userSvc = sql.GetUserService()
	d.factionSvc = *sql.GetFactionService()
	d.startSvc = *sql.GetStartService()

	// mark as ready
	DownTimeInitialized = true
}

// Runs daily downtime jobs - should be run before universe loading
func (d *DownTimeRunner) RunDownTimeTasks() {
	// brief sleep
	time.Sleep(100 * time.Millisecond)

	// safety check
	if DownTimeHandled {
		panic("second attempt at running downtime jobs!")
	}

	if !DownTimeInitialized {
		panic("downtime structure not ready!")
	}

	// disband leaderless custom factions
	err := d.disbandLeaderlessPlayerFactions()

	if err != nil {
		panic(err)
	}

	// cleanup old / orphaned records
	err = d.executeSPCleanup()

	if err != nil {
		panic(err)
	}

	// quarantine npc-only action reports
	err = d.executeSPQuarantineActionReports()

	if err != nil {
		panic(err)
	}

	// average player faction standings from members
	err = d.averagePlayerFactionStandings()

	if err != nil {
		panic(err)
	}

	// respawn stranded NPCs
	err = d.respawnStrandedNPCs()

	if err != nil {
		panic(err)
	}

	// mark as complete
	DownTimeHandled = true
}

// Runs the sp_cleanup stored procedure to purge old / orphaned records
func (d *DownTimeRunner) executeSPCleanup() error {
	shared.TeeLog("  - executeSPCleanup()")
	return d.spSvc.Cleanup()
}

// Runs the sp_quarantineactionreports stored procedure to purge old / orphaned records
func (d *DownTimeRunner) executeSPQuarantineActionReports() error {
	shared.TeeLog("  - executeSPQuarantineActionReports()")
	return d.spSvc.QuarantineActionReports()
}

// Recalculates the aggregate standings value for player custom factions based on the reputation sheets of their members
func (d *DownTimeRunner) averagePlayerFactionStandings() error {
	shared.TeeLog("  - averagePlayerFactionStandings()")

	// get custom factions
	customFactions, err := d.factionSvc.GetPlayerFactions()

	if err != nil {
		return err
	}

	// loop over custom factions
	for _, f := range customFactions {
		// assertion
		if f.IsNPC {
			panic("! was going to average standings for an NPC faction!")
		}

		// get members
		members, err := d.userSvc.GetUsersByFactionID(f.ID)

		if err != nil {
			return err
		}

		// reputation accumulator
		repAcc := make(map[string]*float64)

		// accumulate member faction standings
		for _, m := range members {
			// assertion
			if m.CurrentFactionID != f.ID {
				panic("! was going to include a non-member in standings averaging!")
			}

			// iterate over reputation sheet
			for k, rse := range m.ReputationSheet.FactionEntries {
				// check if accumulator contains this faction yet
				t, ok := repAcc[k]

				if t == nil || !ok {
					// add empty entry
					kv := 0.0
					repAcc[k] = &kv
				}

				// accumulate player standing value
				cv := *repAcc[k]
				cv += rse.StandingValue
				repAcc[k] = &cv
			}
		}

		// divide by number of members to get average
		if len(members) == 0 {
			panic("! empty faction was not caught by sp_cleanup!")
		}

		for z, h := range repAcc {
			cv := *h
			cv /= float64(len(members))
			repAcc[z] = &cv
		}

		// build new faction reputation sheet
		frs := sql.FactionReputationSheet{
			Entries:        make(map[string]sql.ReputationSheetEntry),
			HostFactionIDs: f.ReputationSheet.HostFactionIDs,
			WorldPercent:   f.ReputationSheet.WorldPercent,
		}

		for k, v := range f.ReputationSheet.Entries {
			// copy base entry
			be := v

			// merge average
			cx := repAcc[k]
			cv := 0.0

			if cx != nil {
				cv = *cx
			}

			be.StandingValue = cv

			// update hostility flags
			if cv >= shared.CLEAR_OPENLY_HOSTILE {
				// unset openly hostile flag
				be.AreOpenlyHostile = false
			} else if cv <= shared.BECOME_OPENLY_HOSTILE {
				// set openly hostile flag
				be.AreOpenlyHostile = true
			}

			// store in map
			frs.Entries[k] = be
		}

		// link new sheet
		f.ReputationSheet = frs

		// save to database
		err = d.factionSvc.SaveFaction(f)

		if err != nil {
			return err
		}
	}

	return nil
}

// Closes any factions who's owner has left the faction and returns all members to their starter factions
func (d *DownTimeRunner) disbandLeaderlessPlayerFactions() error {
	shared.TeeLog("  - disbandLeaderlessPlayerFactions()")

	// get custom factions
	customFactions, err := d.factionSvc.GetPlayerFactions()

	if err != nil {
		return err
	}

	// loop over custom factions
	for _, f := range customFactions {
		// assertion
		if f.IsNPC || f.OwnerID == nil {
			panic("! was going to consider disbanding an NPC faction!")
		}

		// get members
		members, err := d.userSvc.GetUsersByFactionID(f.ID)

		if err != nil {
			return err
		}

		// look for leader in roster
		leaderFound := false

		for _, m := range members {
			if m.ID == *f.OwnerID {
				leaderFound = true
				break
			}
		}

		if !leaderFound {
			// kick everyone from the faction and place them in their starter factions
			for _, m := range members {
				// get their start
				start, err := d.startSvc.GetStartByID(m.StartID)

				if err != nil {
					return err
				}

				// assign them to their starter faction
				err = d.userSvc.UpdateCurrentFactionID(m.ID, &start.FactionID)

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// Detects NPCs that failed to respawn and respawns them
func (d *DownTimeRunner) respawnStrandedNPCs() error {
	shared.TeeLog("  - respawnStrandedNPCs()")

	// get stranded NPCs
	l, err := userSvc.GetStrandedNPCs()

	if err != nil || l == nil {
		panic("! unable to list stranded NPCs!")
	}

	// iterate over stranded NPCs
	for _, m := range l {
		// get start
		s, err := startSvc.GetStartByID(m.StartID)

		if err != nil || s == nil {
			panic("! unable to get start for stranded NPC!")
		}

		// respawn ship
		u, err := CreateNoobShipForPlayer(s, m.ID)

		if err != nil || u == nil {
			panic("! unable to respawn stranded NPC!")
		}
	}

	return err
}
