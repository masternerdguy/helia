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
	spSvc sql.SPService
}

// Initializes downtime job structure
func (d *DownTimeRunner) Initialize() {
	// safety check
	if DownTimeHandled || DownTimeInitialized {
		panic("second attempt at initializing downtime jobs!")
	}

	// get services
	d.spSvc = sql.GetSPService()

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
}

// Runs the sp_cleanup stored procedure to purge old / orphaned records
func (d *DownTimeRunner) executeSPCleanup() error {
	shared.TeeLog("  - executeSPCleanup()")
	return d.spSvc.Cleanup()
}
