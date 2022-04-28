package engine

import (
	"fmt"
	"helia/shared"
	"helia/universe"
	"sync"
	"time"

	"github.com/google/uuid"
)

var schematicRunnerStarted = false // indicates whether or not the schematics watcher has been initialized

var schematicRunMap map[string][]*universe.SchematicRun // map of users to schematic runs
var schematicRunMapLock sync.Mutex                      // lock to share user<->schematic runs map

// Initializes the schematics run watcher
func initializeSchematicsWatcher() {
	// check if already started
	if schematicRunnerStarted {
		panic("Schematic runner already running!")
	}

	// initialize map and mutex
	schematicRunMap = make(map[string][]*universe.SchematicRun)
	schematicRunMapLock = sync.Mutex{}
}

// Starts the schematics run watcher
func startSchematics() {
	// check if already started
	if schematicRunnerStarted {
		panic("Schematic runner already running!")
	}

	lastFrame := makeTimestamp()
	msAccumulator := 0

	// start watcher
	go func() {
		// exit notification
		defer shared.TeeLog("! Schematic watcher has halted")

		// watcher loop
		for {
			// check for shutdown signal
			if shutdownSignal {
				break
			}

			// obtain lock
			schematicRunMapLock.Lock()

			// get tpf
			now := makeTimestamp()
			tpf := now - lastFrame

			// store last frame
			lastFrame = makeTimestamp()

			// update accumulator
			msAccumulator += int(tpf)

			if msAccumulator >= 1000 {
				updateRunningSchematics()

				// decrement accumulator
				msAccumulator -= 1000
			}

			// release lock, must be right before sleep
			schematicRunMapLock.Unlock()

			// sleep to avoid pegging cpu
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// mark as started
	schematicRunnerStarted = true
}

// Updates the state of running schematics
func updateRunningSchematics() {
	// iterate over known users
	for _, s := range schematicRunMap {
		// iterate over associated jobs
		for _, j := range s {
			// obtain lock
			j.Lock.Lock()
			defer j.Lock.Unlock()

			// skip if uninitialized
			if !j.Initialized {
				continue
			}

			// increment timer
			if j.StatusID == "running" {
				j.Progress += 1
			} else if j.StatusID == "new" {
				// start job
				j.StatusID = "running"
				j.Progress = 0
			} else if j.StatusID == "deliverypending" {
				// do not redeliver
				continue
			} else if j.StatusID == "error" {
				// do not redeliver
				continue
			} else if j.StatusID == "delivered" {
				// do not redeliver
				continue
			}

			// check if complete
			if j.Process != nil {
				if j.Progress >= j.Process.Time {
					// mark as delivering
					j.StatusID = "deliverypending"

					/* Deliver items on a separate goroutine */

					go func(j *universe.SchematicRun) {
						// obtain lock
						j.Lock.Lock()
						defer j.Lock.Unlock()

						// obtain lock on delivery system and ship
						if j.Ship != nil {
							sh := j.Ship
							s := sh.CurrentSystem

							sh.Lock.Lock()
							defer sh.Lock.Unlock()

							if s != nil {
								s.Lock.Lock()
								defer s.Lock.Unlock()

								// use core to create new items
								for _, o := range j.Process.Outputs {
									// generate a new uuid
									nid, err := uuid.NewUUID()

									if err != nil {
										shared.TeeLog(fmt.Sprintf("Error delivering run result: %v, %v", err, j))
										j.StatusID = "error"

										return
									}

									// is this a ship?
									stIDStr, isShip := o.ItemTypeMeta.GetString("shiptemplateid")

									if !isShip {
										// make a new item stack of the given size
										newItem := universe.Item{
											ID:             nid,
											ItemTypeID:     o.ItemTypeID,
											Meta:           o.ItemTypeMeta,
											Created:        time.Now(),
											CreatedBy:      &sh.UserID,
											CreatedReason:  "Delivered from schematic run",
											ContainerID:    sh.CargoBayContainerID,
											Quantity:       o.Quantity,
											IsPackaged:     true,
											Lock:           sync.Mutex{},
											ItemTypeName:   o.ItemTypeName,
											ItemFamilyID:   o.ItemFamilyID,
											ItemFamilyName: o.ItemFamilyName,
											ItemTypeMeta:   o.ItemTypeMeta,
											CoreDirty:      true,
										}

										// escalate order save request to core
										s.NewItems = append(s.NewItems, &newItem)

										// obtain lock on cargo bay
										sh.CargoBay.Lock.Lock()
										defer sh.CargoBay.Lock.Unlock()

										// place item in cargo bay
										sh.CargoBay.Items = append(sh.CargoBay.Items, &newItem)
									} else {
										// parse template id
										stID, err := uuid.Parse(stIDStr)

										if err != nil {
											shared.TeeLog(fmt.Sprintf("Error delivering run result: %v, %v", err, j))
											j.StatusID = "error"

											return
										}

										// request a new ship to be generated from this purchase
										r := universe.NewShipTicket{
											UserID:         sh.UserID,
											ShipTemplateID: stID,
											StationID:      *sh.DockedAtStationID,
										}

										// escalate order save request to core
										s.NewShipTickets = append(s.NewShipTickets, &r)
									}
								}

								// mark as delivered
								j.StatusID = "delivered"

								// free schematic
								j.SchematicItem.SchematicInUse = false
							} else {
								shared.TeeLog(fmt.Sprintf("Schematic ship is not in a system! %v", sh))
								j.StatusID = "error"
							}

						} else {
							shared.TeeLog(fmt.Sprintf("Schematic run does not have ship attached! %v", j))
							j.StatusID = "error"
						}
					}(j)
				}
			} else {
				shared.TeeLog(fmt.Sprintf("Schematic run does not have process attached! %v", j))
				j.StatusID = "error"
			}
		}
	}
}

// Returns pointers to hooked schematic runs for a given user
func getSchematicRunsByUser(userID uuid.UUID) []*universe.SchematicRun {
	// obtain lock
	schematicRunMapLock.Lock()
	defer schematicRunMapLock.Unlock()

	// check if user known
	_, ok := schematicRunMap[userID.String()]

	if !ok {
		// add empty slice
		schematicRunMap[userID.String()] = make([]*universe.SchematicRun, 0)
	}

	// return pointers
	return schematicRunMap[userID.String()]
}

// Hooks a schematic run into the watcher
func addSchematicRunForUser(userID uuid.UUID, run *universe.SchematicRun) {
	// obtain lock
	schematicRunMapLock.Lock()
	defer schematicRunMapLock.Unlock()

	// check if user known
	_, ok := schematicRunMap[userID.String()]

	if !ok {
		// add empty slice
		schematicRunMap[userID.String()] = make([]*universe.SchematicRun, 0)
	}

	// store run
	schematicRunMap[userID.String()] = append(schematicRunMap[userID.String()], run)
}
