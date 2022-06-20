package main

import (
	"fmt"
	"helia/engine"
	"helia/physics"
	"helia/shared"
	"helia/sql"
	"helia/universe"
	"log"
	"math"
	"math/rand"
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

	var toInject = [...]string{
		"12e06c35-b923-4bf2-aaba-dad99fc26590",
		"8f350754-ec8c-4e37-a4e5-533fb1f8277e",
		"c6bc07e1-2046-4ab1-8d95-a2d5fa332926",
		"074b046f-e370-4e90-aeb9-fd41457cee3f",
		"1f346e6e-73d4-4bfb-8426-e9364fe4e08d",
		"304c2832-0b80-47f3-841d-b0cae8826a8c",
		"6a8a7182-80e6-4f1a-ae98-92c508c5f365",
		"c8252cdd-e60c-4cd3-b495-ec5cf210557a",
		"55966cd5-be70-4194-a517-dfe9d7a8459c",
		"d4736eb1-58bf-4c65-99f5-90633ef43f46",
		"038232ea-16b7-4eb3-b986-64fb2c46c709",
		"8182725b-80e0-4e8c-aa0e-9136a6251e3e",
		"ec68716d-3bf7-4a0a-a3a5-65c08c041ebf",
		"2359bca1-7b69-48fc-8d5b-08a6ce59285e",
		"6e32da1d-08c4-4fbc-a663-dc65d233e341",
		"e416386c-454f-40d2-9b7a-952a32b6e1f8",
		"40bd3a83-29fd-4023-9639-0ffb36ece96d",
		"1604836e-90f8-47d4-b79d-1e08bdfba14a",
		"55f54ac7-21f7-4254-b600-0fdb1cf2036b",
		"0d614b36-cdc7-44e2-838f-729e3cbc2378",
		"12e5ceab-576d-4b4b-9ad8-b51e71dce382",
		"4e664f94-fe79-4b87-974b-8fc180638a88",
		"a31fd51c-0995-4509-a2da-aef736b48b11",
		"27c5458b-c98f-43c7-b5d6-276beac6213f",
		"746c4d74-91bf-492a-a7b1-cc65c343c60b",
		"07a05681-5d39-4511-807e-82db23ce70f2",
		"e3d355a3-c196-41e5-94cf-3721f90458c5",
		"0ee6bad6-2b82-4ec4-b198-7617c1775757",
		"3e211832-700e-45ee-9124-0027d6b69ffd",
		"811426fe-9bfe-4c56-977e-83bbc18d6080",
		"aee56ad6-1fd8-49e2-9f55-70d83003865d",
		"53f3768f-5a67-41b4-9497-323a3a0f63cd",
		"ae1afc9e-836c-4367-ac7c-be7d7cf0af48",
		"18508b43-4725-4039-82c5-c66b74101460",
		"1a5b8edd-fc12-4e64-87ae-1220413e6396",
		"325d3ad6-0bbd-4580-b005-15afc973cbf3",
		"50e009e0-bc39-4c91-9f99-e73a47b271a9",
		"b99c0a57-3064-4cd7-94df-18e797b081b6",
		"29ca9f66-74e3-4346-a1ef-e3e68bae07fe",
		"31ced1da-e9e0-468a-bd9e-6ced19e8382e",
		"1595481f-9694-470f-a84d-f9ecbc927944",
		"c63e529d-5510-4cee-83f8-076628a3959b",
		"29bc69e4-f345-4239-8308-d772962bd7b0",
		"ec4cf466-0ae4-4c15-928a-c86c3b4aa980",
		"e1e3c650-d96d-4b42-913e-9daa83820f09",
		"4f135209-37bd-4c86-a617-2a3dae5b674d",
		"6e387a5b-7ea5-42b7-a627-d56f04c8c793",
		"a8877298-d73f-4834-8ce5-1445d22ea57e",
		"2d70e6d4-3b73-4097-b58c-cb9008c08472",
		"0fbe3b3f-521e-46b2-a3a8-f5da81e17b9b",
		"594dbb18-b3dd-42a4-a078-901ba186cebd",
		"1a93c4ad-0a23-4d7b-a0d5-669da314b172",
		"1c650dc6-8f45-482b-a147-9d64b5409195",
		"6a28d823-3dd7-48ba-a771-870de93ae35f",
		"38bf6fb8-6a8d-4e53-b59f-8c6c5f90ff72",
		"ddb7a683-a49b-466c-93b3-4bfa98bba048",
		"35e9f7e8-c7c9-4ffd-a7c7-bd86f8439cd3",
		"d9204dc9-c782-46a0-938f-577c2f86868f",
		"0a878648-f8b8-46bd-90cc-70547ea478a0",
		"35b283d2-7e02-47b5-94e6-f8c98171ecc4",
		"d16039b4-4cb1-4b81-bd60-42e2e1f18ff8",
		"2eec6d15-f208-4548-af2a-a4694a3eca3b",
		"cf5c76f7-2a3d-4b19-953e-2ace0a4b5566",
		"84840cc0-ccf1-4e89-8796-87fa121fb687",
		"a04883cd-b740-45aa-84af-1f4838ba2cd4",
		"0cec1e99-e42d-4e79-97af-4bc228b9e91e",
		"1e7d3a3e-4dc1-40fd-8c17-a5bf53cd4255",
		"c1447d5b-f47a-4ee5-a0ba-f44ce4c4fb8d",
		"d93b6035-449f-4491-bf2e-a638880b311b",
		"e2e66ed4-6c62-4f89-bd3e-1064f4f91c08",
		"6f0718ad-7404-4e91-b389-9bd5620a98f8",
		"271b8be9-9f1b-44e6-b8f6-03943ac2def1",
		"befb8ef0-291b-4a12-bff4-b688c40880b2",
		"9e72f786-57f1-4a3f-851a-e0b0d0037ae3",
		"4678141e-37db-4281-af05-a1e2ed27a5ab",
		"f819fbd4-fc9a-48e1-9627-274793fa7905",
		"a0c87142-0d9d-4e6a-b1eb-4dd3b8b4bf24",
		"ed859b0a-432b-4e2a-a6bc-6a50c27cb944",
		"b86043e6-7716-4bb1-9041-79839697a930",
		"1b30aa7d-4fbb-4caf-8254-b09c2d2d1445",
		"4d8e04fc-d935-40a9-b324-8066be7744c5",
		"a76064ea-2ff5-4a53-9bc0-1b3e4dbe1819",
		"64f15ac7-063e-490f-99c7-ac1e82c8b50c",
		"f550e887-b382-452b-8fc3-a1de38544bc7",
		"20c2f530-dab9-4d05-b20f-ac98ff5d79ff",
		"dd9eca8d-6cbc-4960-96ed-5878a4f92efe",
		"d2c629db-95f4-48a0-b255-0ed0d5ac801f",
		"a34a0243-89d0-4167-a02f-e683eaf85e43",
		"d019d663-db16-4e6f-959f-87780d4cbc53",
		"6cc051bb-393c-4a5c-a089-13cabd3c0b07",
		"4e7ca073-c5e1-4ede-b19f-b42081fd7122",
		"0fa14c7c-3518-4be8-9a3a-9ebee24e434d",
	}

	// dropAsteroids(universe)
	//dropSanctuaryStations(universe)

	for i, e := range toInject {
		log.Println(fmt.Sprintf("injecting process %v", e))
		injectProcess(universe, e, i)
	}

	//stubModuleSchematicsAndProcesses()
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

		if i.Family == "power_cell" {
			// todo: deal with more than one of these in the future
			smallPowerCell = i
		}

		if i.Family == "depleted_cell" {
			// todo: deal with more than one of these in the future
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
	prob := 2

	stationProcessSvc := sql.GetStationProcessService()

	if err != nil {
		panic(err)
	}

	var textures = [...]string{
		"sanctuary-",
	}

	toSave := make([]sql.StationProcess, 0)

	for _, r := range u.Regions {
		for _, s := range r.Systems {
			rand.Seed(int64(calculateSystemSeed(s)) - 77827798 + int64(offset))

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
