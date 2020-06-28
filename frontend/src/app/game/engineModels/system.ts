import { WSSystem } from '../wsModels/entities/wsSystem';
import { Ship } from './ship';

export class System extends WSSystem {
    constructor(ws: WSSystem) {
        super();

        // copy from ws model
        this.id = ws.id;
        this.systemName = ws.systemName;

        // initialize self
        this.ships = [];
    }

    ships: Ship[];
    backplateImg: HTMLImageElement;
}
