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

	// cleanup old / orphaned records
	err := d.executeSPCleanup()

	if err != nil {
		panic(err)
	}

	// average player faction standings from members
	err = d.averagePlayerFactionStandings()

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
			cv := *repAcc[k]
			be.StandingValue = cv

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
