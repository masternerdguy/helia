import { WSSystem } from './entities/wsSystem';
import { WSShip } from './entities/wsShip';

export class ServerGlobalUpdateBody {
    currentSystemInfo: WSSystem;
    ships: WSShip[];
}
