import { UnwrappedStarMapData } from '../wsModels/bodies/viewStarMap';
import { Container } from './container';
import { Faction } from './faction';
import { Ship } from './ship';
import { System } from './system';

export class Player {
  currentShip: Ship;
  currentCargoView: Container;
  currentStarMap: UnwrappedStarMapData;
  currentSystem: System;

  currentTargetID: string;
  currentTargetType: TargetType;

  uid: string;
  sid: string;

  getFaction(): Faction {
    if (!this.currentShip) {
      return null;
    }

    return this.currentShip.getFaction();
  }
}

export enum TargetType {
  Nothing = -1,
  Ship = 1,
  Station = 2,
  Star = 3,
  Planet = 4,
  Jumphole = 5,
  Asteroid = 6,
  Wreck = 7,
  Outpost = 8,
}
