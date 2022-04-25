package main

import (
	"helia/sql"

	"github.com/google/uuid"
)

/*
 * This file is just a scratch pad for performing arbitrary operations I need
 */

func main() {
	// get faction service
	factionSvc := sql.GetFactionService()

	/*

	   id	name	description
	   a8a28085-e7b4-48f5-b8cb-1465ccab82a5	Test Starter Faction	Temporary starter faction for use when a player is created.
	   42b937ad-0000-46e9-9af9-fc7dbf878e6a	Neutral	Not associated with any faction.

	   bdeffd9a-3cab-408c-9cd7-32fce1124f7a	Ozouka Accord	todo (side A empire faction, lawful)
	   5db2bec7-37c3-4f1c-ab88-21024c12d639	Tenevan Coalition	todo (side A empire faction, lawful)

	   a0152cbb-8a78-45a4-ae9b-2ad2b60b583b	Kingdom of Antaria	todo (side B empire faction, lawful)
	   27a53dfc-a321-4c12-bf7c-bb177955c95b	Vierra Federation	todo (side B empire faction, lawful)

	   7decfd86-9c82-4a17-af3b-5d4af8c4e2ad	Bad Rabbits	todo (high pirate faction, unlawful)
	   30bc70eb-2692-47e6-a7e3-2772e131c3d7	Pappa Fly Reloaded	todo (low pirate faction, unlawful)

	   559d7fc1-5470-4ab4-8c66-fa2f0b89a523	Interstar Corporation	todo (high corporate, quasi-lawful)
	   506286f5-6613-4481-ac26-caa9940fbe68	Alvaca	todo (low corporate, quasi-lawful)

	   b3d3fa9c-b21e-490f-b39e-128b3af12128	Sanctuary Systems	todo (freeport league)



	*/

	f, _ := factionSvc.GetAllFactions()

	a := "5db2bec7-37c3-4f1c-ab88-21024c12d639"
	b := "b3d3fa9c-b21e-490f-b39e-128b3af12128"

	for _, e := range f {
		if e.ID.String() == a || e.ID.String() == b {
			src := ""
			dest := ""

			if e.ID.String() == a {
				src = a
				dest = b
			} else {
				src = b
				dest = a
			}

			srcID, _ := uuid.Parse(src)
			destID, _ := uuid.Parse(dest)

			if e.ReputationSheet.Entries == nil {
				e.ReputationSheet.Entries = make(map[string]sql.ReputationSheetEntry)
			}

			e.ReputationSheet.Entries[dest] = sql.ReputationSheetEntry{
				SourceFactionID:  srcID,
				TargetFactionID:  destID,
				StandingValue:    0,
				AreOpenlyHostile: false,
			}
		}

		_ = factionSvc.SaveFaction(e)
	}

	/*// add faction
	f := sql.Faction{}
	f.Meta = make(sql.Meta)

	this is bad and needs to use the new repsheet column instead!

	f.Name = "Sanctuary Systems"
	f.Description = "todo (freeport league)"
	f.CanHoldSov = true
	f.IsNPC = true
	f.IsClosed = false
	f.IsJoinable = true
	f.Ticker = "_S_"

	s := sql.FactionReputationSheet{}

	s.Entries = make(map[string]sql.ReputationSheetEntry)
	s.HostFactionIDs = make([]uuid.UUID, 0)
	s.WorldPercent = 0.05

	f.Meta["reputationSheet"] = s

	// save faction
	t, u := factionSvc.NewFaction(f)

	log.Println(fmt.Sprintf("%v, %v", t, u))*/
}

/*

id	name	description
a8a28085-e7b4-48f5-b8cb-1465ccab82a5	Test Starter Faction	Temporary starter faction for use when a player is created.
42b937ad-0000-46e9-9af9-fc7dbf878e6a	Neutral	Not associated with any faction.

bdeffd9a-3cab-408c-9cd7-32fce1124f7a	Ozouka Accord	todo (side A empire faction, lawful)
5db2bec7-37c3-4f1c-ab88-21024c12d639	Tenevan Coalition	todo (side A empire faction, lawful)
a0152cbb-8a78-45a4-ae9b-2ad2b60b583b	Kingdom of Antaria	todo (side B empire faction, lawful)
27a53dfc-a321-4c12-bf7c-bb177955c95b	Vierra Federation	todo (side B empire faction, lawful)

7decfd86-9c82-4a17-af3b-5d4af8c4e2ad	Bad Rabbits	todo (high pirate faction, unlawful)
30bc70eb-2692-47e6-a7e3-2772e131c3d7	Pappa Fly Reloaded	todo (low pirate faction, unlawful)

559d7fc1-5470-4ab4-8c66-fa2f0b89a523	Interstar Corporation	todo (high corporate, quasi-lawful)
506286f5-6613-4481-ac26-caa9940fbe68	Alvaca	todo (low corporate, quasi-lawful)

b3d3fa9c-b21e-490f-b39e-128b3af12128	Sanctuary Systems	todo (freeport league)



*/
