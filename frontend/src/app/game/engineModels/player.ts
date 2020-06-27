import { Ship } from './ship';
import { System } from './system';

export class Player {
    currentShip: Ship;
    currentSystem: System;

    uid: string;
    sid: string;
}
