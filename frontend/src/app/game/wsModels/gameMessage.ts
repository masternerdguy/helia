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
  StackItem = 16
}
