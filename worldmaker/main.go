package main

import (
	"fmt"
	"helia/physics"
	"helia/sql"
	"log"
	"math"
	"math/rand"
	"sort"

	"github.com/google/uuid"
)

/*
 * Contains routines for procedurally generating the universe
 */

const MinSystemCount = 1250
const MaxSystemCount = 1500
const MinRegions = 27
const MaxRegions = 45
const MinExtent = 3
const MaxExtent = 10
const Connectivity = 6
const SpiralFactor = 30
const MinPlanets = 0
const MaxPlanets = 15
const MinStars = 1
const MaxStars = 3
const MinSystemRadius = 100000
const MaxSystemRadius = 1000000
const MinPlanetRadius = 500
const MaxPlanetRadius = 5000
const MinStarRadius = 6000
const MaxStarRadius = 10000
const MinStarMass = 500000000
const MaxStarMass = 30000000000
const MinPlanetMass = 950000
const MaxPlanetMass = 80000000
const MinJumpholeRadius = 150
const MaxJumpholeRadius = 425
const MinJumpholeMass = 200000
const MaxJumpholeMass = 900000

var StarTextures = [...]string{
	"vh_main_sequence/star_blue01",
	"vh_main_sequence/star_orange02",
	"vh_main_sequence/star_red03",
	"vh_main_sequence/star_white04",
	"vh_main_sequence/star_blue02",
	"vh_main_sequence/star_orange03",
	"vh_main_sequence/star_red04",
	"vh_main_sequence/star_yellow01",
	"vh_main_sequence/star_blue03",
	"vh_main_sequence/star_orange04",
	"vh_main_sequence/star_white01",
	"vh_main_sequence/star_yellow02",
	"vh_main_sequence/star_blue04",
	"vh_main_sequence/star_red01",
	"vh_main_sequence/star_white02",
	"vh_main_sequence/star_yellow03",
	"vh_main_sequence/star_orange01",
	"vh_main_sequence/star_red02",
	"vh_main_sequence/star_white03",
	"vh_main_sequence/star_yellow04",
}

var PlanetTextures = [...]string{
	"vh_unshaded/planet11",
	"vh_unshaded/planet22",
	"vh_unshaded/planet31",
	"vh_unshaded/planet43",
	"vh_unshaded/planet02",
	"vh_unshaded/planet13",
	"vh_unshaded/planet24",
	"vh_unshaded/planet32",
	"vh_unshaded/planet44",
	"vh_unshaded/planet03",
	"vh_unshaded/planet16",
	"vh_unshaded/planet25",
	"vh_unshaded/planet33",
	"vh_unshaded/planet45",
	"vh_unshaded/planet06",
	"vh_unshaded/planet17",
	"vh_unshaded/planet26",
	"vh_unshaded/planet38",
	"vh_unshaded/planet46",
	"vh_unshaded/planet07",
	"vh_unshaded/planet18",
	"vh_unshaded/planet27",
	"vh_unshaded/planet39",
	"vh_unshaded/planet47",
	"vh_unshaded/planet08",
	"vh_unshaded/planet19",
	"vh_unshaded/planet28",
	"vh_unshaded/planet40",
	"vh_unshaded/planet48",
	"vh_unshaded/planet09",
	"vh_unshaded/planet20",
	"vh_unshaded/planet29",
	"vh_unshaded/planet41",
	"vh_unshaded/planet10",
	"vh_unshaded/planet21",
	"vh_unshaded/planet30",
	"vh_unshaded/planet42",
}

func main() {
	// initialize RNG
	rand.Seed(101)

	// determine size of map
	extent := physics.RandInRange(MinExtent, MaxExtent)

	// determine how many systems to generate
	systemCount := physics.RandInRange(MinSystemCount, MaxSystemCount)

	// determine how many regions to sort them into
	regionCount := physics.RandInRange(MinRegions, MaxRegions)

	// output chosen values
	log.Println(fmt.Sprintf("Extent: %v, Regions: %v, Systems: %v", extent, regionCount, systemCount))

	// generate empty systems in a spiral
	systems := generateEmptySystems(extent, systemCount)

	// generate empty regions in map
	regions := generateEmptyRegions(systems, regionCount)

	// sort each system into the region closest to it
	for _, s := range systems {
		if s.RegionID == nil {
			var closest *Regionling = nil
			var distance float64 = 0.0

			for _, r := range regions {
				// initialize if needed
				if closest == nil {
					closest = r
					distance = physics.Distance(physics.Dummy{PosX: r.PosX, PosY: r.PosY}, physics.Dummy{PosX: s.PosX, PosY: s.PosY})
				}

				// check distance to region
				d := physics.Distance(physics.Dummy{PosX: r.PosX, PosY: r.PosY}, physics.Dummy{PosX: s.PosX, PosY: s.PosY})

				if d < distance {
					// new closest region
					closest = r
					distance = d
				}
			}

			s.RegionID = &closest.ID
			closest.Systems = append(closest.Systems, s)
		}
	}

	// name systems
	usedSystemNames := make(map[string]*Sysling)

	for _, s := range systems {
		for {
			n := randomPlaceholderName()

			if usedSystemNames[n] == nil {
				s.Name = n
				usedSystemNames[n] = s

				break
			}
		}
	}

	// name regions
	usedRegionNames := make(map[string]*Regionling)

	for _, s := range regions {
		for {
			n := randomPlaceholderName()

			if usedRegionNames[n] == nil {
				s.Name = n
				usedRegionNames[n] = s

				break
			}
		}
	}

	// generate jumphole connections
	jumpNetworkEdges := make([]Edgeling, 0)

	for _, s := range systems {
		systemsTmp := systems

		// sort by distance from this system
		sort.SliceStable(systemsTmp, func(i, j int) bool {
			ie := systemsTmp[i]
			je := systemsTmp[j]

			d1 := physics.Distance(physics.Dummy{PosX: s.PosX, PosY: s.PosY}, physics.Dummy{PosX: ie.PosX, PosY: ie.PosY})
			d2 := physics.Distance(physics.Dummy{PosX: s.PosX, PosY: s.PosY}, physics.Dummy{PosX: je.PosX, PosY: je.PosY})

			return d1 < d2
		})

		// get top entries
		linkCount := physics.RandInRange(1, Connectivity)
		links := make([]*Sysling, 0)

		for v := 1; v < linkCount+1; v++ {
			links = append(links, systemsTmp[v])
		}

		// insert edges into network map
		for _, l := range links {
			// make sure there isn't already an equivalent edge
			safe := true

			for _, e := range jumpNetworkEdges {
				if e.A == s && e.B == l {
					safe = false
					break
				}

				if e.A == l && e.B == s {
					safe = false
					break
				}
			}

			if safe {
				// insert edge
				jumpNetworkEdges = append(jumpNetworkEdges, Edgeling{
					A: s,
					B: l,
				})
			}
		}
	}

	// re-sort any systems that somehow didn't get into a real region
	unsafe := make([]*Sysling, 0)

	for _, s := range systems {
		safe := false

		for _, r := range regions {
			if r.ID == *s.RegionID {
				safe = true
			}
		}

		if !safe {
			unsafe = append(unsafe, s)
		}
	}

	for _, s := range unsafe {
		var closest *Regionling = nil
		var distance float64 = 0.0

		for _, r := range regions {
			// initialize if needed
			if closest == nil {
				closest = r
				distance = physics.Distance(physics.Dummy{PosX: r.PosX, PosY: r.PosY}, physics.Dummy{PosX: s.PosX, PosY: s.PosY})
			}

			// check distance to region
			d := physics.Distance(physics.Dummy{PosX: r.PosX, PosY: r.PosY}, physics.Dummy{PosX: s.PosX, PosY: s.PosY})

			if d < distance {
				// new closest region
				closest = r
				distance = d
			}
		}

		s.RegionID = &closest.ID
		closest.Systems = append(closest.Systems, s)
	}

	// generate stars + planets
	for _, s := range systems {
		// determine number of stars + planets to generate
		starCount := physics.RandInRange(MinStars, MaxStars)
		planetCount := physics.RandInRange(MinPlanets, MaxPlanets)

		// generate stars
		for i := 0; i < starCount; i++ {
			sl := generateStar(s, i)

			if sl != nil {
				s.Stars = append(s.Stars, sl)
			}
		}

		// generate planets
		for i := 0; i < planetCount; i++ {
			sl := generatePlanet(s, i)

			if sl != nil {
				s.Planets = append(s.Planets, sl)
			}
		}
	}

	// generate jumpholes from edges
	for _, e := range jumpNetworkEdges {
		injectJumpholes(&e)
	}

	/* NOTE: THE UNIVERSE IN THE DB SHOULD BE ESSENTIALLY EMPTY AT THIS POINT!!! */

	regionSvc := sql.GetRegionService()
	systemSvc := sql.GetSolarSystemService()
	starSvc := sql.GetStarService()
	planetSvc := sql.GetPlanetService()
	jumpholeSvc := sql.GetJumpholeService()

	// save empty regions
	for _, r := range regions {
		o := sql.Region{
			PosX:       r.PosX,
			PosY:       r.PosY,
			RegionName: r.Name,
			ID:         r.ID,
		}

		err := regionSvc.NewRegionWorldMaker(&o)

		if err != nil {
			panic(err)
		}
	}

	// temporarily using neutral for all holdings
	nID, err := uuid.Parse("42b937ad-0000-46e9-9af9-fc7dbf878e6a")

	if err != nil {
		panic(err)
	}

	// save empty solar systems
	for _, s := range systems {
		o := sql.SolarSystem{
			PosX:             s.PosX,
			PosY:             s.PosY,
			SystemName:       s.Name,
			RegionID:         *s.RegionID,
			HoldingFactionID: nID,
			ID:               s.ID,
		}

		err := systemSvc.NewSolarSystemWorldMaker(&o)

		if err != nil {
			panic(err)
		}
	}

	// save planets, stars, and jumpholes in solar systems
	for _, s := range systems {
		for _, st := range s.Stars {
			o := sql.Star{
				ID:       st.ID,
				SystemID: s.ID,
				PosX:     st.PosX,
				PosY:     st.PosY,
				Texture:  st.Texture,
				Radius:   st.Radius,
				Mass:     st.Mass,
				Theta:    st.Theta,
			}

			err := starSvc.NewStarWorldMaker(&o)

			if err != nil {
				panic(err)
			}
		}

		for _, st := range s.Planets {
			o := sql.Planet{
				ID:         st.ID,
				SystemID:   s.ID,
				PlanetName: st.Name,
				PosX:       st.PosX,
				PosY:       st.PosY,
				Texture:    st.Texture,
				Radius:     st.Radius,
				Mass:       st.Mass,
				Theta:      st.Theta,
			}

			err := planetSvc.NewPlanetWorldMaker(&o)

			if err != nil {
				panic(err)
			}
		}

		for _, st := range s.Jumpholes {
			o := sql.Jumphole{
				ID:           st.ID,
				SystemID:     s.ID,
				OutSystemID:  st.OutSystemID,
				JumpholeName: st.Name,
				PosX:         st.PosX,
				PosY:         st.PosY,
				Texture:      st.Texture,
				Radius:       st.Radius,
				Mass:         st.Mass,
				Theta:        st.Theta,
			}

			err := jumpholeSvc.NewJumpHoleWorldMaker(&o)

			if err != nil {
				panic(err)
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

func injectJumpholes(edge *Edgeling) {
	r := float64(physics.RandInRange(MinJumpholeRadius, MaxJumpholeRadius))
	t := float64(rand.Float64() * 360.0)
	m := float64(physics.RandInRange(MinJumpholeMass, MaxJumpholeMass))

	// system A
	{
		jumphole := Jumpling{
			ID:          uuid.New(),
			Name:        fmt.Sprintf("%v Jumphole", edge.B.Name),
			OutSystemID: edge.B.ID,
			Texture:     "Jumphole",
		}

		system := edge.A

		for {
			safe := true

			x := float64(physics.RandInRange(-system.Radius*3, system.Radius*3))
			y := float64(physics.RandInRange(-system.Radius*3, system.Radius*3))

			a := physics.Dummy{
				PosX: x,
				PosY: y,
			}

			for _, l := range system.Jumpholes {
				b := physics.Dummy{
					PosX: l.PosX,
					PosY: l.PosY,
				}

				d := physics.Distance(a, b)

				if d <= math.Max(r*2, l.Radius*2) {
					safe = false
				}
			}

			for _, l := range system.Stars {
				b := physics.Dummy{
					PosX: l.PosX,
					PosY: l.PosY,
				}

				d := physics.Distance(a, b)

				if d <= math.Max(r*5, l.Radius*5) {
					safe = false
				}
			}

			for _, l := range system.Planets {
				b := physics.Dummy{
					PosX: l.PosX,
					PosY: l.PosY,
				}

				d := physics.Distance(a, b)

				if d <= math.Max(r*3, l.Radius*3) {
					safe = false
				}
			}

			if safe {
				jumphole.PosX = x
				jumphole.PosY = y
				jumphole.Theta = t
				jumphole.Radius = r
				jumphole.Mass = m

				system.Jumpholes = append(system.Jumpholes, &jumphole)

				break
			}
		}
	}

	// system B
	{
		jumphole := Jumpling{
			ID:          uuid.New(),
			Name:        fmt.Sprintf("%v Jumphole", edge.A.Name),
			OutSystemID: edge.A.ID,
			Texture:     "Jumphole",
		}

		system := edge.B

		for {
			safe := true

			x := float64(physics.RandInRange(-system.Radius*3, system.Radius*3))
			y := float64(physics.RandInRange(-system.Radius*3, system.Radius*3))

			a := physics.Dummy{
				PosX: x,
				PosY: y,
			}

			for _, l := range system.Jumpholes {
				b := physics.Dummy{
					PosX: l.PosX,
					PosY: l.PosY,
				}

				d := physics.Distance(a, b)

				if d <= math.Max(r*2, l.Radius*2) {
					safe = false
				}
			}

			for _, l := range system.Stars {
				b := physics.Dummy{
					PosX: l.PosX,
					PosY: l.PosY,
				}

				d := physics.Distance(a, b)

				if d <= math.Max(r*5, l.Radius*5) {
					safe = false
				}
			}

			for _, l := range system.Planets {
				b := physics.Dummy{
					PosX: l.PosX,
					PosY: l.PosY,
				}

				d := physics.Distance(a, b)

				if d <= math.Max(r*3, l.Radius*3) {
					safe = false
				}
			}

			if safe {
				jumphole.PosX = x
				jumphole.PosY = y
				jumphole.Theta = t
				jumphole.Radius = r
				jumphole.Mass = m

				system.Jumpholes = append(system.Jumpholes, &jumphole)

				break
			}
		}
	}
}

func generateStar(system *Sysling, seq int) *Starling {
	star := Starling{
		ID:   uuid.New(),
		Name: fmt.Sprintf("%v S%v", system.Name, seq+1),
	}

	for {
		safe := true

		x := float64(physics.RandInRange(-system.Radius, system.Radius))
		y := float64(physics.RandInRange(-system.Radius, system.Radius))
		r := float64(physics.RandInRange(MinStarRadius, MaxStarRadius))
		t := float64(rand.Float64() * 360.0)
		m := float64(physics.RandInRange(MinStarMass, MaxStarMass))

		a := physics.Dummy{
			PosX: x,
			PosY: y,
		}

		for _, l := range system.Stars {
			b := physics.Dummy{
				PosX: l.PosX,
				PosY: l.PosY,
			}

			d := physics.Distance(a, b)

			if d <= math.Max(r, l.Radius) {
				safe = false
			}
		}

		for _, l := range system.Planets {
			b := physics.Dummy{
				PosX: l.PosX,
				PosY: l.PosY,
			}

			d := physics.Distance(a, b)

			if d <= math.Max(r, l.Radius) {
				safe = false
			}
		}

		if safe {
			star.PosX = x
			star.PosY = y
			star.Theta = t
			star.Radius = r
			star.Mass = m

			break
		}
	}

	tIdx := physics.RandInRange(0, len(StarTextures))
	tx := StarTextures[tIdx]

	star.Texture = tx

	return &star
}

func generatePlanet(system *Sysling, seq int) *Planetling {
	planet := Planetling{
		ID:   uuid.New(),
		Name: fmt.Sprintf("%v P%v", system.Name, seq+1),
	}

	for {
		safe := true

		x := float64(physics.RandInRange(-system.Radius, system.Radius))
		y := float64(physics.RandInRange(-system.Radius, system.Radius))
		r := float64(physics.RandInRange(MinPlanetRadius, MaxPlanetRadius))
		t := float64(rand.Float64() * 360.0)
		m := float64(physics.RandInRange(MinPlanetMass, MaxPlanetMass))

		a := physics.Dummy{
			PosX: x,
			PosY: y,
		}

		for _, l := range system.Planets {
			b := physics.Dummy{
				PosX: l.PosX,
				PosY: l.PosY,
			}

			d := physics.Distance(a, b)

			if d <= math.Max(r, l.Radius) {
				safe = false
			}
		}

		for _, l := range system.Planets {
			b := physics.Dummy{
				PosX: l.PosX,
				PosY: l.PosY,
			}

			d := physics.Distance(a, b)

			if d <= math.Max(r, l.Radius) {
				safe = false
			}
		}

		if safe {
			planet.PosX = x
			planet.PosY = y
			planet.Theta = t
			planet.Radius = r
			planet.Mass = m

			break
		}
	}

	tIdx := physics.RandInRange(0, len(PlanetTextures))
	tx := PlanetTextures[tIdx]

	planet.Texture = tx

	return &planet
}

func generateEmptyRegions(systems []*Sysling, regionCount int) []*Regionling {
	o := make([]*Regionling, 0)

	for p := 0; p < regionCount; p++ {
		id := uuid.New()

		// pick a system at random that doesn't have a region
		for {
			exit := false

			sIdx := physics.RandInRange(0, len(systems))
			s := systems[sIdx]

			if s.RegionID == nil {
				s.RegionID = &id

				// this will be the first system in the region
				r := Regionling{
					ID:      s.ID,
					PosX:    s.PosX,
					PosY:    s.PosY,
					Systems: make([]*Sysling, 0),
				}

				r.Systems = append(r.Systems, s)

				// store region
				o = append(o, &r)

				// done
				exit = true
			}

			// exit if done
			if exit {
				break
			}
		}
	}

	return o
}

func generateEmptySystems(extent int, systemCount int) []*Sysling {
	o := make([]*Sysling, 0)

	for p := 0; p < systemCount; p++ {
		r := math.Sqrt(rand.Float64()) * float64(extent)

		x := float64(SpiralFactor)*r*math.Cos(r) + (2*float64(SpiralFactor)*rand.Float64() - float64(SpiralFactor))
		y := float64(SpiralFactor)*r*math.Sin(r) + (2*float64(SpiralFactor)*rand.Float64() - float64(SpiralFactor))

		ri := physics.RandInRange(MinSystemRadius, MaxSystemRadius)

		o = append(o, &Sysling{
			ID:       uuid.New(),
			PosX:     x,
			PosY:     y,
			Stars:    make([]*Starling, 0),
			Planets:  make([]*Planetling, 0),
			Radius:   ri,
			RegionID: nil,
		})
	}

	return o
}

// Represents a scaffolding for a solar system
type Sysling struct {
	ID        uuid.UUID
	PosX      float64
	PosY      float64
	RegionID  *uuid.UUID
	Name      string
	Stars     []*Starling
	Planets   []*Planetling
	Jumpholes []*Jumpling
	Radius    int
}

// Represents a scaffolding for a region
type Regionling struct {
	ID      uuid.UUID
	PosX    float64
	PosY    float64
	Systems []*Sysling
	Name    string
}

// Represents an edge in the jumphole network
type Edgeling struct {
	A *Sysling
	B *Sysling
}

// Represents a jumphole in the jumphole network
type Jumpling struct {
	ID          uuid.UUID
	OutSystemID uuid.UUID
	PosX        float64
	PosY        float64
	Radius      float64
	Mass        float64
	Theta       float64
	Name        string
	Texture     string
}

// Represents a scaffolding for a planet
type Planetling struct {
	ID      uuid.UUID
	PosX    float64
	PosY    float64
	Radius  float64
	Mass    float64
	Theta   float64
	Name    string
	Texture string
}

// Represents a scaffolding for a star
type Starling struct {
	ID      uuid.UUID
	PosX    float64
	PosY    float64
	Radius  float64
	Mass    float64
	Theta   float64
	Name    string
	Texture string
}
