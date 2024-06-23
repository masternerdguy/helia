package main

import (
	"encoding/csv"
	"fmt"
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
	/*
		// load universe from database
		shared.TeeLog("Loading universe from database...")
		universe, err := engine.LoadUniverse()

		if err != nil {
			panic(err)
		}

		shared.TeeLog("Loaded universe!")*/

	/*
	 * COMMENT AND UNCOMMENT THE BELOW ROUTINES AS NEEDED
	 */

	/*var toInject = [...]string{
		"21deccaa-c9b0-414e-a898-3ab9230d4528",
		"f0cf1d7a-c1c6-41d9-a7c2-00ce1a7b0cce",
		"b47202fc-34cf-4117-a63a-eff63c62febd",
		"c40ad708-724f-40a9-b99d-29eec2049e98",
		"ee66bd1c-4ac4-4ad6-9ff2-834e13580921",
		"f4b40184-803f-4387-bfb2-439d444c29b2",
		"76c9b909-3f15-41f4-ac77-cc13dc96ee29",
		"7737b828-4ef0-4344-bc53-923f27a885e1",
		"b0177b10-2b91-40a4-b695-b19681249c2b",
		"b01122d1-3c82-49ce-ba3d-bc47de0741c1",
		"0b5700d7-e787-43ad-84c8-c1442deff38b",
		"c81ee20b-0466-42f6-9962-24b195fdc913",
		"b46679ae-33c6-4392-b535-6797d3152de7",
		"350c8d81-e6dc-457b-922e-bcf4ba3da585",
		"215eabca-9cc5-4b94-84ee-df4be3a7888b",
		"7723775b-24c7-4cd9-85bd-f2c943cbc2b6",
		"7228a4d1-0cf8-4f6b-a8dd-0a84f4fa892b",
		"88a99a46-937e-406e-a357-50d18a2e1184",
		"85aed9df-db0b-42e7-9cbb-a5874f7dba43",
		"26541a60-0606-464e-a86b-6745ba74c08c",
		"198f524d-b707-49da-aa6c-17296915b231",
		"defb2d6f-8d22-4e98-96c0-ae71eee6abf0",
		"ce90d325-6a78-44e9-a964-df39b839a1ee",
		"351f52bb-c581-462d-9a7d-c102942f71ca",
		"c075da8e-efe5-4f0f-88bc-129544d545c0",
		"2061a44b-c0ec-4d97-beec-6ad9b30011eb",
		"0b33103e-7e46-4b6f-ae8f-8a7096ba7d4e",
		"f4fac632-9248-459f-8309-30015c60f72f",
		"03390735-4a97-4c70-9ab2-dd58384c962f",
		"de7c37e5-92a0-46bb-93e2-b74d3b123a2b",
		"622c7b71-1e76-44ca-8ce5-474fbac77245",
		"c64f0b49-7d0f-4214-a668-803ff824ca38",
		"374db251-abcd-466f-9b4c-a8268570c01a",
		"4c463a82-69ca-4be2-a46a-209bd106a1fe",
		"944b329c-5c70-4bee-b8af-ca3aa5143906",
		"3cedf939-829d-4ceb-823f-5a6b6bdacc0c",
		"5385eb4b-2046-46f2-901c-8555bfaeda50",
		"ae36fea3-c8f4-47e3-9f8f-8f30d2084a83",
		"657e02d7-b32b-4b9b-b6be-0e2fa5cc1709",
		"6a24faa9-5b35-4339-bbf6-3a942247fe06",
		"8a51aeca-74ff-48f0-9907-56c6008e3279",
		"c72cc388-026b-4e0e-a8c1-2341757a80e8",
		"947c290c-4983-4057-87e2-77ec8f1c2e73",
		"aca8755d-e0e2-44e6-9b70-7ec09523874e",
		"3f03b17d-821c-4350-af28-8ad49ea01e98",
		"2256a571-4cd6-4fa1-964c-7217b26ac6fc",
		"ce38bcee-625a-4eb7-b6cd-b48477a12854",
		"5e5973ae-9217-4c76-999f-e60549462a1f",
		"08b70d06-5861-495b-bccb-4daad5b11519",
		"ba762629-2664-4331-b74a-0134360e695f",
		"764aef65-10c4-4811-86a2-cd3fa019f935",
		"5d8fdcc5-3116-4bdd-849c-f93838234acc",
		"67d15f39-9861-4f90-b2ba-75de6cd1cfa1",
		"0b7c5ed8-a114-4929-bd54-2f3d5372d3a6",
		"a9702e63-9fb9-4f95-b81d-b5582605e4da",
		"247ddad5-03f4-4363-a78e-2c54ecfc6e87",
		"9808b8e0-0f63-4f4a-9f31-d1ed940fc030",
		"c04d38e2-8c89-4775-9f27-1793a20d70fe",
		"956a96c1-a6fa-40bc-8871-417c246f99a3",
		"fb62983a-b1cf-4693-a87e-b31ba7867222",
	}*/

	// dropAsteroids(universe)
	//dropSanctuaryStations(universe)

	/*for i, e := range toInject {
		log.Printf("injecting process %v", e)
		injectProcess(universe, e, i)
	}*/

	//stubModuleSchematicsAndProcesses()

	loadNewWares()
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
				log.Println(fmt.Sprintf("Skipping %v - not solvable with input materials given.", industrial.Name))
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
		log.Println(fmt.Sprintf(":) %v", industrial.Name))
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

func injectProcess(u *universe.Universe, pid string, offset int) {
	pID, err := uuid.Parse(pid)
	prob := 3

	stationProcessSvc := sql.GetStationProcessService()

	if err != nil {
		panic(err)
	}

	var textures = [...]string{
		"bad-",
		"sanctuary-",
		"interstar-",
		"alvaca-",
	}

	toSave := make([]sql.StationProcess, 0)

	for _, r := range u.Regions {
		for _, s := range r.Systems {
			rand.Seed(int64(calculateSystemSeed(s)) - 903827450 + int64(offset))

			/*if r.ID.ID()%2 != 0 {
				continue
			}

			if s.ID.ID()%5 != 0 {
				continue
			}*/

			stations := s.CopyStations(true)

			for _, st := range stations {
				for _, t := range textures {
					if strings.Contains(st.Texture, t) {
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

type WareCsvRecord struct {
	Name     string
	MinPrice int
	MaxPrice int
	SiloSize int
	Volume   int
}

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

	// to store item types to be saved
	newTypes := make([]sql.ItemType, 0)

	// iterate over parsed wares
	for _, w := range wares {
		// verify this ware doesn't already exist
		existing, err := itemTypeSvc.GetItemTypeByName(w.Name)

		if err != nil {
			log.Panicf("Ware already exists! %v", existing)
		}

		// create new item type for ware
		it := sql.ItemType{
			ID:     uuid.New(),
			Family: "trade_good",
			Name:   w.Name,
			Meta:   sql.Meta{},
		}

		meta := universe.IndustrialMetadata{
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
		_, err := itemTypeSvc.NewItemTypeForWorldFiller(it)

		if err != nil {
			log.Panicf("Error saving new item type %v", err)
		}
	}
}

/*

{
  "hp": 1,
  "volume": 2,
  "industrialmarket": {
    "maxprice": 47,
    "minprice": 32,
    "silosize": 150000000
  }
}

*/

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
