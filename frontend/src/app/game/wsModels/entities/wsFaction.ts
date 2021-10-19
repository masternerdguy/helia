export class WSFaction {
  id: string;
  name: string;
  description: string;
  isNPC: boolean;
  isJoinable: boolean;
  isClosed: boolean;
  canHoldSov: boolean;
  ticker: string;
  relationships: WSFactionRelationship[];
}

export class WSFactionRelationship {
  factionId: string;
  openlyHostile: boolean;
  standingValue: number;
}
