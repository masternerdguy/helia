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
		"09fdf2bc-dde1-429b-b595-81166b392981",
		"40d13f5c-0d61-4ee6-af53-7934b42f9533",
		"40604d6c-0a0d-4e9b-a107-51630f01a8c1",
		"17a883bb-dca3-4349-b072-b63f8a204a63",
		"a4c75049-d81f-4c80-b7aa-1cfad1a84b45",
		"f577a67b-a9ef-407e-908e-96dbb5d6e71d",
		"5217b31d-ec77-46cf-873a-faaaaf3d24ed",
		"cb090191-a600-418c-bc88-0ef35c846e5e",
		"db443e40-c08d-461e-a0cb-0f7614b036f9",
		"ab770d50-1a34-45ed-8b4b-2de68771eac3",
		"6cae9b27-987b-4f3d-838a-466805ac99b9",
		"51f95b66-60d7-4dec-8eef-5fbbcf701c60",
		"7be86058-54ff-4688-af70-77e9f8ce3ecc",
		"7d2bfb85-bece-49a2-9eb5-06c33a717fa5",
		"e48a390c-7749-4495-984d-95cea122e3d6",
		"26a51572-8410-46ac-bf35-9c8582d8e172",
		"e8265208-9139-45d6-8589-46a4053f8c57",
		"ffddb4a7-8753-4ef8-9b32-cddf082dd424",
		"e46a0590-5cd5-417f-a41c-28c19c7555ab",
		"6da418b9-425f-484e-9a43-bbe1ad657cbb",
		"6f83372f-2f7f-4f73-8b8c-257b3a6d78df",
		"19fdbe23-cd9d-4c7d-ad49-44a43912237c",
		"d0e07a4a-7d87-40b1-9459-db0fa69a7380",
		"9450db39-9e17-424a-8843-d7a2fbbcf81c",
		"24f615a5-568b-4bf9-b0aa-576e6063165d",
		"eeef7269-c73e-4af1-ba3d-d1e6e2912f81",
		"154aa403-8e98-42df-ac3c-e1b5a5f6dc4c",
		"7a4ba867-75c0-4d9e-9f4f-fd5fa4f0a31f",
		"856f3bf2-4bd5-49fc-8080-7e812de86619",
		"69022d56-b794-46ba-a56e-e124a19e9a5e",
		"5e67a1f2-330f-4a68-9b9e-875629a1e0e9",
		"828bc69b-91d7-4fba-9ffd-b18a306be56f",
		"11c914d0-6049-4739-86ef-50f94dac5841",
		"9b77980d-c12b-4ede-9fc0-a023b78a9070",
		"40e8d7e8-ecc0-43f0-8db1-0b9e4b39bca7",
		"c45d7bb6-ea4c-46a3-93fe-f262ac7f56a0",
		"7c830d36-d180-47b8-995c-b883e8678c41",
		"711dcf60-4b74-4706-b7e5-14ce44691c90",
		"1ca81daf-d903-428d-9a9d-962de9676879",
		"79a2103c-7a4f-4b00-90fc-5803b34591d6",
		"49af374f-5ea8-4ff9-a51e-6c82cdf368fc",
		"5bb99aa5-4498-4c39-a275-7472292699bc",
		"09d7acc0-dc99-498d-9c5a-9a4b6fce75fa",
		"5be1bb69-aa5d-4f97-8c5f-a711fe42bec7",
		"20c41620-15de-45c2-81d9-18afbede6075",
		"a877a619-c357-485a-ae85-b43f8fe841c4",
		"04e54fe4-356c-43be-bd58-5650dc1c438b",
		"b3ed62dd-a5f5-4889-9257-af9cf59be61e",
		"083a8a4d-2280-4a6b-b59f-b0443a4efab6",
		"725d57ba-78b8-489b-a060-43aa59fff3fd",
		"3783e856-1397-4f69-93ba-f66e9d6e6a1e",
		"8aaa09a9-4b18-45a8-8f7f-18f0e1bcb9c1",
		"cbe2eaf0-b530-4a1f-abcb-747da35672ca",
		"c3f339de-8647-4d42-9911-e9d4e6890b4e",
		"bf735889-a4b8-478e-8da3-9322d13bc96c",
		"3a296939-f0cc-4da9-8dae-9336b678536d",
		"3443f66a-d21b-48b4-ab3d-207cb3271ecc",
		"b347756d-d547-4a20-a906-869ec31f3db7",
		"36be8581-7aec-4455-af27-64c8e37064ae",
		"59069a4b-b4a6-47f4-8c99-66c2716e29f6",
		"888a6579-4836-4791-aa0b-3d8113632343",
		"9c9971f1-0515-431f-b834-ca212d1872c8",
		"9ebec0d3-856c-4acf-bb0d-064313271b5f",
		"d63b55d9-85b5-4c7c-b6f3-0e5db702dbde",
		"d7a0d8fd-4a1a-4ee5-a01f-4908cc0190ef",
		"da788336-501c-483c-ab62-a0f142922d77",
		"9cd59fb5-6227-4074-8097-14b23fc229d2",
		"5719f3fc-e804-4b5f-aa88-dce64199b89b",
		"26f62e5e-da65-439c-8842-6232e072d543",
		"240944b8-76b9-4fd5-aeeb-3eee50a7cd48",
		"ee82b9a3-ed79-433b-a8b5-ee22c1169ac3",
		"df8e0a3c-a33b-499b-ae36-6ec515af98a7",
		"cdfd30ad-f2fe-49c4-aab9-4168ac9c484d",
		"b6189378-db66-4d95-b25f-0de9f9626955",
		"ecbefccf-5847-43c3-b8ca-7cd098ac02a3",
		"6e528db6-44bf-40b0-9ef3-83e911ca9511",
		"69f99c28-6c7c-462b-895f-e8de44484546",
		"eb16aa2e-9f4b-4853-bf3f-c5d2d4916159",
		"5a13c720-3e20-459f-9cd2-402185dc3dbb",
		"67d81a98-08e6-4a56-97f0-8238522dccbf",
		"fa528aa4-ba69-48f0-9df7-7a9d54a73a0f",
		"9519e3c4-098c-452f-be88-0cef1ab4eaf5",
		"bb6a1d24-26a2-4876-ab9a-54c10b653efb",
		"aa13c3e3-bd7f-404f-87c2-5c0d4998f32f",
		"50b8f85e-aa15-41d0-971c-6ea678c0dfc3",
		"a87949fd-29cd-4927-accf-482b242442b7",
		"5cee8772-99bd-4440-80d7-e62182b39e06",
		"12070161-f95d-4c4c-a7b5-0945540b6d0b",
		"113eb258-4d7e-4f6e-9dda-dcc40097c77d",
		"99597692-5b10-4179-a1cd-51d0b6d3a5a4",
		"3e5be1a9-e20f-4dd8-bb77-b58acb23b64e",
		"27e4341e-ffd4-4d0f-b9c6-ddd0711da2a7",
		"bd6f96df-310b-4708-a64e-5415b26e36b2",
		"b97fa64d-fc45-4ed5-9c94-2c7fa6a59b19",
		"32941e6b-e3f0-440f-a05c-bfa5e700f894",
		"af96d517-35a4-4227-b307-77c0cdef9690",
		"fa7b639a-93c6-49cb-ba8d-6bef8fd51a3d",
		"d1a373ba-2ace-4644-bdd4-2450451f8a27",
		"a7bd08d4-fd24-49ff-b402-935f39fcc7cc",
		"32c8ded7-a62d-4702-be7d-6adf20b46205",
		"5e89c9a6-c661-472a-9708-c68c7e9eb822",
		"81f656c6-e566-4153-a842-990a6ad476e3",
		"3bc08593-ebd1-49ef-aa0e-c96648298503",
		"92e77d78-3057-4c91-9645-e999879f0727",
		"c6db873a-7a54-4a0a-99f2-11e4c1ffc7cd",
		"1e0806d0-1793-4b26-975a-76d2f66cf554",
		"9960a250-a939-449e-9bda-c3fd59a72365",
		"6615e874-039a-4a96-8a04-938978367956",
		"de94eaa4-9fa4-4457-8b45-8c272ddabc84",
		"3ff4e198-496b-41b7-a55d-324c28494389",
		"b4ab9a27-1ff8-44c2-90ba-c264453186fa",
		"96305bd7-e12f-4992-a123-ff4a1d5488ef",
		"600c279b-b406-4a96-8cfc-ae764f421a15",
		"6a8d073e-612a-4ff2-a407-92ce63fb5184",
		"936479d2-03c9-4535-a9ed-488c71daa327",
		"a7e320c5-5bde-4cba-8363-98216b4fdedd",
		"bf17e268-1874-4179-a3bc-cb55f0004497",
		"d61711f6-e23c-42b1-bf46-a0cbf63d036f",
		"0873d708-5ad5-44bb-80a3-0bf1c6331e3f",
		"e23b38f6-66a4-4332-ad99-c76b5a525ff7",
		"057037aa-005d-4c6b-9782-c997930ae28f",
		"9c39d83f-9513-4748-b8de-0880d8994244",
		"ca66082d-af12-4c7f-9e9c-a1798c9efceb",
		"b49e1b56-2b62-4477-87a7-a8dedf88c0b3",
		"3fa70678-ff98-4006-8b5d-4debb7899065",
		"bbb46312-65fb-4115-a78f-775a8ca897e2",
		"db647c69-9cda-404e-a650-2ebdfa4a7161",
		"d298cdd8-de3d-4f92-b12c-7853a95209a8",
		"9f14b95f-e913-4af2-9f48-ba84c348ec05",
		"9c4c8946-a4f1-4fb5-aa82-675cb571262c",
		"227d053e-6bcc-416f-947f-75171066ec99",
		"4983cd8d-15d2-4aa6-8184-41b3b0ad7b6b",
		"cc1aae00-95dd-4b61-a4a4-d02c62935fe6",
		"48323402-ae4d-4ef7-be99-170ad5e0f27c",
		"ceb370cb-3413-4362-bc94-a715e260f2e7",
		"608ac7f8-e556-4738-9025-5bd3135fd62a",
		"fd0501a8-e9a4-4a9a-bcbf-de5d60660022",
		"77f4b6e3-3fc1-4149-905b-4a46726efe34",
		"c18315d6-cdcb-417b-96e4-e184b14df938",
		"96dfda57-b53c-47e2-a64f-fd360bafb524",
		"a83e43e4-06b2-4d22-a276-58ba3f95015c",
		"4372bc0d-4aa5-4a78-84f1-01e19e006911",
		"80caff11-2e1e-485f-b21e-ed30de599d4d",
		"62d5ac0d-050b-4483-9659-5d0b4e7f73cc",
		"8d49df06-0b85-4f08-a532-58070d92b583",
		"1f4c3d89-820c-43ae-81df-a4702f651fe0",
		"6be7d664-ed5c-4a36-bcb0-54c233c507c1",
		"055f80d4-745e-4ba5-98e2-c66db79df322",
		"c3d54b9a-1f67-4a40-96cd-079c09b42f50",
		"d42a4c73-8b7c-4ca8-b545-0e9949d7a812",
		"5c0ed36e-8050-4e04-9831-25346f2e03cf",
		"02a5bdfc-e57c-48ae-92b6-ecd948dea2db",
		"cdd2ca73-95ea-4347-814f-f0ecea46c7c4",
		"aa841c4c-5e6e-4cf9-9d62-d8ede3f8a817",
		"788e47da-6fb8-4683-b7e0-aa7cbdf73db2",
		"1476864d-2d0b-4ac8-b914-e858b4e4aaee",
		"345bcb5f-06e3-48ba-b367-f8c9880278f2",
		"957d966e-66ed-46c9-b8e0-204e94a1367f",
		"27131967-272d-49aa-ab77-dd73ef0b8d58",
		"be42d0f8-4a03-4d23-bd04-671ae783ff53",
		"b40712ab-33fc-4fa7-a390-b996cc8b3171",
		"062c07a2-6a4d-4340-b1b6-abb9739d648c",
		"cd030577-70a2-4053-bac8-c1706d479028",
		"7c0737d0-21f0-4d51-be59-03682a09b001",
		"655736a9-4734-4611-bc4d-8b4b3c44cfe4",
		"0d42f62b-d2dd-4411-a2a4-ed7b0f0e5643",
		"ade39d47-aec8-44d2-ae73-7ec9e6fe8bbd",
		"0c578ec0-6852-42b2-8754-28aecf717707",
		"23801eb1-cdf0-4cee-a00f-1f86db77b3c7",
		"19f420eb-7891-474c-9c7a-00220abee7db",
		"a12e3047-edcc-4c17-ae8e-dfe5a77fc59c",
		"9b9aad54-12bb-45bd-a65c-7a74330930d3",
		"32c608d7-325a-4c47-8c2c-ea837b41701e",
		"0ef15457-83ba-47dd-a6b9-e76da7650af0",
		"19834524-520d-4f24-be82-eb5fc8dd7e4c",
		"401ec453-59e5-4156-9179-5e7c812dcc89",
		"80b0cf71-204e-4fe9-9fcd-54d51626d2f3",
		"1ba8e84f-8a99-452a-9b21-41bdd1edde9a",
		"284376a7-241a-458e-be0a-e4fa8b39be45",
		"4e0cf147-2173-44de-bc07-64a326479a1d",
		"639c34a5-18ea-42cb-888e-4f603c391a4b",
		"52577be4-efa2-45d0-badb-e8d239531d07",
		"ab8f640d-bedb-4c98-85f5-74b3b90c6288",
		"5730a68b-58e0-422f-8439-97b06fb50151",
		"1cd34849-60be-4603-a520-33038e7fbea0",
		"a282f4d7-9b41-4f6d-8990-a43ec31120ce",
		"dd436f0b-832f-478c-a88d-565f23a11e58",
		"2733b209-731a-48bc-8a99-4f8d9de5a963",
		"db46bb85-1d9a-43f2-820a-7c654673e27d",
		"12087075-ea5d-44be-8ee8-02c8aa1a56fc",
		"78b63ae0-c540-482e-bd27-af56500f9d91",
		"d689f88b-44e3-4ee4-9d93-aea413e7929f",
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
			rand.Seed(int64(calculateSystemSeed(s)) - 57982197 + int64(offset))

			/*if r.ID.ID()%3 != 0 {
				continue
			}

			if s.ID.ID()%2 != 0 {
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
