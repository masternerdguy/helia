export class GameMessage {
  type: number;
  body: string;
}

export enum MessageTypes {
  PushError = -1,
  Join = 0,
  Update = 1,
  NavClick = 2,
  CurrentShipUpdate = 3,
  Goto = 4,
  Orbit = 5,
  Dock = 6,
  Undock = 7,
  ActivateModule = 8,
  DeactivateModule = 9,
  ViewCargoBay = 10,
  CargoBayUpdate = 11,
  UnfitModule = 12,
  TrashItem = 13,
  PackageItem = 14,
  UnpackageItem = 15,
  StackItem = 16,
  SplitItem = 17,
  FitModule = 18,
  SellAsOrder = 19,
  ViewOpenSellOrders = 20,
  OpenSellOrdersUpdate = 21,
  BuySellOrder = 22,
  ViewIndustrialOrders = 23,
  IndustrialOrdersUpdate = 24,
  BuyFromSilo = 25,
  SellToSilo = 26,
  FactionUpdate = 27,
  ViewStarMap = 28,
  StarMapUpdate = 29,
  ConsumeFuel = 30,
  PlayerFactionUpdate = 31,
}
