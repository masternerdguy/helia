import { Ship } from './ship';
import { System } from './system';

export class Player {
  currentShip: Ship;
  currentSystem: System;

  currentTargetID: string;
  currentTargetType: TargetType;

  uid: string;
  sid: string;
}

export enum TargetType {
  Ship = 1,
  Station = 2,
  Star = 3,
  Planet = 4,
  Jumphole = 5,
}
