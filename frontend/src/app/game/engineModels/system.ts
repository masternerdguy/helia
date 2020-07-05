import { WSSystem } from '../wsModels/entities/wsSystem';
import { Ship } from './ship';
import { Star } from './star';
import { Planet } from './planet';

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
    }

    ships: Ship[];
    stars: Star[];
    planets: Planet[];
    backplateImg: HTMLImageElement;
}
