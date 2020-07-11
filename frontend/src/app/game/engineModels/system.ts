import { WSSystem } from '../wsModels/entities/wsSystem';
import { Ship } from './ship';
import { Star } from './star';
import { Planet } from './planet';
import { Station } from './station';

export class System extends WSSystem {
    constructor(ws: WSSystem) {
        super();

        // copy from ws model
        this.id = ws.id;
        this.systemName = ws.systemName;

        // initialize self
        this.ships = [];
        this.stars = [];
        this.planets = [];
        this.stations = [];
    }

    ships: Ship[];
    stars: Star[];
    planets: Planet[];
    stations: Station[];
    backplateImg: HTMLImageElement;
}
