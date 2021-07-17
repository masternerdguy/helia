package main

import (
	"fmt"
	"helia/sql"
	"log"

	"github.com/google/uuid"
)

/*
 * This file is just a scratch pad for performing arbitrary operations I need
 */

func main() {
	// get faction service
	factionSvc := sql.GetFactionService()

	// add faction
	f := sql.Faction{}
	f.Meta = make(sql.Meta)

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

	log.Println(fmt.Sprintf("%v, %v", t, u))
}

/*

Ageiran Federation
Orin Federation

ITC
Caina Conglemerate

*/
