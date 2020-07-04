import { WSSystem } from '../entities/wsSystem';
import { WSShip } from '../entities/wsShip';
import { WSStar } from '../entities/wsStar';

export class ServerGlobalUpdateBody {
    currentSystemInfo: WSSystem;
    ships: WSShip[];
    stars: WSStar[];
}
