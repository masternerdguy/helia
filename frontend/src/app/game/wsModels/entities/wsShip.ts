export class WSShip {
  id: string;
  uid: string;
  createdAt: string;
  shipName: string;
  ownerName: string;
  x: number;
  y: number;
  systemId: string;
  texture: string;
  theta: number;
  velX: number;
  velY: number;
  accel: number;
  mass: number;
  radius: number;
  turn: number;
  shieldP: number;
  armorP: number;
  hullP: number;
  energyP: number;
  heatP: number;
  fuelP: number;
  fitStatus: WSFitting;
  dockedAtStationID: string;
}

export class WSFitting {
  aRack: WSRack;
  bRack: WSRack;
  cRack: WSRack;
}

export class WSRack {
  modules: WSModule[];
}

export class WSModule {
  itemID: string;
  itemTypeID: string;
  family: string;
  type: string;
  isCycling: boolean;
  willRepeat: boolean;
  cyclePercent: number;
}
