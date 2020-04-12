package engine

import (
	"fmt"
	"helia/sql"
	"helia/universe"
	"log"
)

func loadUniverse() (*universe.Universe, error) {
	//get services
	regionSvc := sql.GetRegionService()

	u := universe.Universe{}

	//load regions
	rs, err := regionSvc.GetAllRegions()

	if err != nil {
		return nil, err
	}

	regions := make([]*universe.Region, 0)
	for _, e := range rs {
		//load basic region information
		r := universe.Region{
			ID:         e.ID,
			RegionName: e.RegionName,
		}

		log.Println(fmt.Sprintf("%v", r))

		//todo: load systems

		regions = append(regions, &r)
	}

	//store regions and systems
	u.Regions = regions

	return &u, nil
}
