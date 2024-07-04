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

	/*// load universe from database
	shared.TeeLog("Loading universe from database...")
	universe, err := engine.LoadUniverse()

	if err != nil {
		panic(err)
	}

	shared.TeeLog("Loaded universe!")
	*/
	/*
	 * COMMENT AND UNCOMMENT THE BELOW ROUTINES AS NEEDED
	 */
	/*
		var toInject = [...]string{
			"582fff52-332a-4c9c-814f-f70b3db0d361",
			"6b2490f3-d274-4a49-9a11-315f6964ba5a",
			"5972c81e-aabd-4a74-ae79-19aee544f568",
			"a07d4ad0-1be1-4e33-8ed1-c1391942b3be",
			"4433e336-0969-49d1-98a7-135ec025dfb4",
			"b18afe73-7de1-4f42-9e23-c15b9047558e",
			"92b0f6c5-7a1c-4a8b-a6a9-3ec0faad9626",
			"18b4ec68-86bf-43be-b16d-89c74da81794",
			"cc72c386-f4a9-4d6d-b772-6e52b8025065",
			"5d4e46a6-5b4c-4a42-b5f7-1f8ebea75157",
			"f659e665-c0f4-4390-801c-ebb0bead5c53",
			"9614167a-e4d3-4581-a5cf-a603a06429bf",
			"e49a2920-5e84-47e8-9233-f8fbe9b7549e",
			"b7ac86bc-6e37-4519-aa55-326af2374f34",
			"c44f265b-4cd2-4694-894a-9956b94a76fb",
			"3f3039f8-6f87-485b-9e5c-ddb9eed47acf",
			"64d480d8-45ec-46a9-8f45-4943b7015372",
			"0ce85f9b-2933-4f46-9684-efa261956d4e",
			"bd86fbc8-6c2d-4a8c-b4f8-aa04c9814415",
			"e69a83a8-3e8b-4bca-ad46-951ff118b423",
			"84942262-e00f-48b4-baeb-3e6c38aa77f1",
			"41d6a990-1b15-40fe-a56d-46a9be56201d",
			"e4f62fc2-e70c-421d-a4a1-10f38b006d55",
			"16d2fb12-f69b-4f50-bbab-39f43de250f6",
			"7b2402f7-97ca-4889-8b8a-d74cd7839623",
			"a7f2e536-ecc4-4e40-ac76-0dde649203e2",
			"28687a8a-834f-4395-92dc-2a1993f5e8f2",
			"5755a002-7820-47e4-aa57-72a56a8ddc3d",
			"5502f886-5a5d-4a9a-8251-c97b82b56ebd",
			"811feeb2-1647-436b-96c2-74f868b21fe8",
			"f2fd6193-3592-4d48-aa30-3f62d6eb3c2c",
			"2b0e62b7-b344-4b4a-9f7c-be9f928589ce",
			"30bdef7e-07cf-4ddf-a3c3-f0ad9d03773c",
			"7db4c707-8d8b-42de-af55-16e13b992e91",
			"08c8dfdd-2679-49f7-a371-64057a8b6dd9",
			"b2f48189-e44d-4629-9213-ff612ca5230b",
			"37670065-b364-45e6-a3f0-c5bdb9f80e74",
			"aa354e12-2535-4a39-ab77-4302520bc4d4",
			"13ba35d6-67cf-4ea4-ae89-19c86f188a12",
			"557fa68c-a88c-4d4b-ae76-9ddfe55d3d43",
			"2a12bb7f-8f71-406a-bb65-c839d97a64e4",
			"06e95f1c-0652-4d44-bef2-7fae08df8772",
			"795d9c9d-0cef-47c3-afd2-1e47118bfc3a",
			"72bef950-8f04-4175-acbc-49778ccc7f90",
			"2b3677fe-5fed-4af1-a500-ad1291813d7e",
			"0f3d0718-ede0-4701-a191-d7e6c44e35ce",
			"bace7d32-1bf6-4d18-acee-3e4ed86cb451",
			"996dbb36-a33b-4f79-bf87-8190f0d0def2",
			"ddf890ab-7d42-4cf2-a110-fed0a6bd4ad5",
			"a61858aa-40a8-470f-93dd-d2b803f12b4e",
			"a9e7244b-2ce1-4336-b5db-21978d8b434c",
			"5c89c2d3-e258-4c05-840c-90e2cc5a590d",
			"71baf14d-49fb-46fb-aae3-003e5b39fd19",
			"927680ff-5a0f-4982-b07c-935457cb48bc",
			"f521294d-d881-4804-8396-9e1f195f2526",
			"ec19248b-0132-409a-81ff-956aa04980d5",
			"08c625e5-ff68-436f-aea4-d9ef0bc8cd33",
			"9e48337d-eea1-4f3b-b667-3d3b39d6681e",
			"94a1780b-f137-431e-b655-da3d8da1290b",
			"3347ddf4-71de-4051-a93a-cfa6be10f0cf",
			"08456917-debf-4950-9361-1f47b7298337",
			"b08ca0ce-7201-46cb-bcda-76bc922bd25a",
			"dcd00169-33ce-42a0-b9d7-4d7276a6daf6",
			"da8241e5-9f95-4f16-99fd-2c6a49208450",
			"12bbb205-51fd-44e9-adc9-e849b8542410",
			"16077bb6-911b-4852-8247-24e561b17364",
			"b5309636-5352-4759-981d-bc31b92567b6",
			"19ac42b2-65b1-43f3-b85a-8c6b375affa4",
			"986ebe82-4219-46c4-a057-02c3694b1e07",
			"a0f67820-a3fb-45ce-b0c8-45e117e25126",
			"6a7fc1b6-c331-461d-9839-2f855beb1259",
			"ef08d1fc-4660-4edd-aa47-1fa72a6ea730",
			"a9b761e2-96fb-4f64-84cf-e6b069fcd516",
			"c6e69ed3-0cb0-44b7-87aa-52d657d3ef2c",
			"1eeb2a8f-1ce4-4ccf-abb7-19f12aa9bbcc",
			"b5938172-3ec4-4b34-92e0-cdac33b15c52",
			"bf25d1c2-5bce-4261-9bf7-e93b527f96d1",
			"892c9822-38f8-4e5a-a78e-07bb150daf7b",
			"99f147ba-741d-4495-b40f-a9e318e47962",
			"38df3baa-64b0-483c-bb7f-7155d57f708a",
			"57fc41d9-7919-40e4-b9b0-44cb58c6d6ee",
			"5481487f-212c-4bfa-839c-575caa510f9a",
			"7f239a88-9da2-4a47-81bf-d5168ef1894d",
			"96ca2647-a68c-4eec-aae3-8b368f53c5ef",
			"b02b72e8-fa3e-46ed-9ae3-438183fe1c8b",
			"6dc20625-d3c6-42c6-ac77-017b4f4aa9c8",
			"925b4d12-fbf4-442d-bb7c-b301063e1072",
			"f7e3f476-a038-4095-bb3a-941d239282b4",
			"062301b8-c647-4998-93be-d9b94690c167",
			"e4da1804-8500-4150-aaae-55f32bdd3b16",
			"6229c660-4de0-4095-bd0d-8b6ca29eaf4c",
			"05444e33-02d6-4f8a-9261-324a8bafb75d",
			"e48bbd32-44fb-4ebb-8a9c-83e48c1334e2",
			"91fe1440-6cea-4b44-b488-02fcf916d45f",
			"bb22aedd-5ce3-4184-aa02-65c9f2d73ee4",
			"65afee21-290c-4627-989d-920b4c2b8b72",
			"9721c28a-8539-4f5d-92b6-ecab36fc4d73",
			"a389fea4-da61-4f63-b25f-d403060d532f",
			"ee1b7245-5404-470f-a5da-f0f44e024ca4",
			"22b329a6-d612-41b7-9cdd-78204d2304fd",
			"42950c84-48a5-49e8-a0c6-7614994630c8",
			"788a5389-db4d-4f77-b838-d3caace35ed5",
			"04a421da-a218-4ca0-8324-28c481923f25",
			"01c0dfc7-3670-4bd6-9a3c-e263f5652f50",
			"1e09eff9-d52e-41a3-86f1-803ab97f93a7",
			"9cdf16ed-dd7e-4c35-9f03-e102e5884cb0",
			"30e0c788-0f5f-425a-bbc7-b239e9ef6313",
			"6043530e-5e75-43f8-a85d-84556a20ea88",
			"1ab7e519-d41b-494d-b1c1-f92e8c3068df",
			"2520372d-2b32-4879-8aea-8fe07cc2f4d9",
			"3fe252ed-0932-4017-92b0-ee8988dc28c6",
			"c5413194-fe43-4ef8-a62b-828d2b1d7e51",
			"5051eb13-d3ff-4df0-b77c-938d325b0ed0",
			"ecab61ab-75a9-492c-8145-03a8e3075bcd",
			"209d6e90-2263-4448-a1dc-06d6b8d0db11",
			"23df5e75-7b3a-4e75-b6a9-ad89f726fee7",
			"564fee33-bf23-4861-93ba-b45302d17c96",
			"f0826569-b124-4c75-92af-eef339d4c1d2",
			"c28340ec-cd77-4e78-8cd6-ce753c9cc806",
			"fe51093a-d690-48ac-a156-0142a635e69a",
			"6bbd9ff7-8d23-48c3-b298-19b8d443f1ea",
			"5e346a8e-db8f-4436-8b6b-6aed558af7b7",
			"5773b753-d56a-47b6-acce-cd00489b5e68",
			"f6c62bba-2ae2-4052-a48e-667559882dca",
			"c40be21c-2d12-4ae2-99a8-12ab1ae8138b",
			"c771b022-36a5-4fb8-b219-e99626299dff",
			"a8881a01-1622-4d32-beca-c56c387d1d37",
			"9a7407ca-b048-4fd0-8548-443e387199ba",
			"da43dbf3-0c9d-4f0e-9106-4a396b2e78ec",
			"4d3be032-56e5-42d3-8f96-b73baac77c90",
			"83ebf580-d49f-4c7f-beb9-12e56e0bc5f3",
			"0aabdabf-591e-4f98-b27b-723335b7be21",
			"9ea682fa-5427-4261-a48d-8e4efcc97906",
			"3526f1db-8960-4456-91d5-b6274ae78ea6",
			"ab129bb8-91cb-496e-93e2-1520e3c79e56",
			"4cd4bced-c1bb-4cb9-986e-54f11c22734c",
			"348117e6-fda7-4996-b9ad-f0bafe908f67",
			"3fa3164f-05bf-4ef3-a278-e0f093fc8371",
			"479a16fb-2817-43f6-bb0c-c49dc8916b08",
			"a64c76cb-ec31-4b8a-90da-039dcd130e27",
			"81c3e70d-4427-4748-a1a3-d2fee81624b1",
			"b48b208e-04bf-4659-a1cc-45e329d03cdd",
			"431db2b2-4b7e-4bfd-8126-3a0b1f303c1c",
			"55fc726b-d2a0-4b21-9f80-3f8a9d663b5e",
			"131db682-7327-4f3f-8bac-722eb284d7a1",
			"5595c065-5817-4563-8bda-7f577911d275",
			"90a8c8ad-faf5-4f73-895b-dfd85c22b93a",
			"d136f899-abe5-4bce-9ab9-f13367301a3f",
			"180da7d1-a195-447f-a2a3-68b49672cc6a",
			"cf49d8e2-554e-4f06-b665-70990aac1863",
			"d4a0b3b9-cd1d-4519-96f1-cd299ea19fc9",
			"a6e21565-31f0-4461-9e5f-de492471550f",
			"c59651ee-4ef7-45c0-8939-9f52d452baff",
			"db2dde41-aadd-4fc2-a021-af6c48d92b09",
			"1d3d3d84-b3e3-4fcc-ae38-9ded57c9a369",
			"7c2eef1f-a877-4881-ac0b-a626ba1c9acd",
			"ac92be47-d076-43a0-aad1-56c877bb32f6",
			"8fc4fec7-c3de-4802-b75a-6d3cad6241dd",
			"ae4d3480-6be0-471a-82f0-3143d87d0d03",
			"79ff91d3-d575-44c0-b10a-5d240fa5a1b6",
			"3751b85e-3947-4785-9992-567097f9f268",
			"edc2531e-4f91-4492-9b5d-c5fe56fd73dc",
			"df3477e8-3638-449f-8f37-cd622d00b71f",
			"a5ddee77-8013-48db-9386-3f2483773a66",
			"0067fcdb-7d37-495d-83a2-97663b212e76",
			"cfd4d951-ceaa-4054-a856-06404167c63f",
			"9b7ce331-cf3d-4e2d-893f-ea354edda498",
			"3e59e211-7a79-49f3-af0f-8d3e5f5a94b3",
			"426dd27b-2741-4aad-9c98-fdfffa593e8d",
			"f6dddc27-8cf0-46d5-9744-268dafebda27",
			"08564a4c-b14f-4e83-b8bc-fe09a97fedc7",
			"c4858a85-ef9d-4a1c-9da5-db1d4d4e57fd",
			"5a5b917e-ba8e-4824-afe5-eacbfcb25ebd",
			"2a031272-389e-4f44-b53f-5dbcec945463",
			"26bb1a9a-ca38-459e-882c-d88b944674e8",
			"39479049-8094-4adf-b27d-dec3350ef9f6",
			"7674bf29-148a-4de8-aaae-f3dce3064000",
			"d2208d6c-a01e-47cc-99fb-7d40cdd196c1",
			"7bbe912d-046e-47dc-9307-1f541a112a86",
			"e3203031-4dd9-40a4-b15f-777e38375c1f",
			"b6beeb96-2730-4124-912b-c8d812811392",
			"e807f2a3-ba11-44af-9e09-2a4add5e46c3",
			"bf0adb0e-9102-40a9-b766-c49bb1b388a2",
			"de773eee-19cf-4ef0-96a8-1fd88b4fa6db",
			"2b2f9f8f-9811-4d64-a808-3178d6a6994a",
			"f1560623-403e-452b-883b-6548ceb7b546",
			"2d2e0658-c423-4b63-b7d1-aa37c0bd7cdd",
			"2883312d-8d23-4c7a-826f-e4cc7160a53d",
			"aa79aed2-d763-479c-9fc4-27b592ff9123",
			"e42a1595-7e9b-4ba2-bec7-e2017afb0582",
			"a48a7979-a54e-4501-9132-fbb409f0cd65",
			"e5432683-ad9b-4e7d-856e-c68145d6d4d4",
			"0d3dda63-69b0-46b5-8e5b-53d2353e367e",
			"bdb31099-f3e8-4c85-bd6a-936c673ec1fb",
			"34d76bcb-b023-4867-a533-3f54a77f434b",
			"d99072ef-2e90-4c2d-84d5-573944ff2dbd",
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

	/*var textures = [...]string{
		"bad-",
		"sanctuary-",
		"interstar-",
		"alvaca-",
	}*/

	toSave := make([]sql.StationProcess, 0)

	for _, r := range u.Regions {
		for _, s := range r.Systems {
			rand.Seed(int64(calculateSystemSeed(s)) - 178046782389 + int64(offset))

			/*if r.ID.ID()%2 != 0 {
				continue
			}

			if s.ID.ID()%5 != 0 {
				continue
			}*/

			stations := s.CopyStations(true)

			for _, st := range stations {
				/*for _, t := range textures {
				if strings.Contains(st.Texture, t) {*/
				roll := physics.RandInRange(0, 100)

				if roll <= prob {
					sp := sql.StationProcess{
						StationID: st.ID,
						ProcessID: pID,
						ID:        uuid.New(),
					}

					toSave = append(toSave, sp)
				}
				/*	}
					}*/
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
			Family: "trade_good",
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
