export class WSFaction {
  id: string;
  name: string;
  description: string;
  isNPC: boolean;
  isJoinable: boolean;
  canHoldSov: boolean;
  ticker: string;
  ownerId: string;
  relationships: WSFactionRelationship[];
}

export class WSFactionRelationship {
  factionId: string;
  openlyHostile: boolean;
  standingValue: number;
}
