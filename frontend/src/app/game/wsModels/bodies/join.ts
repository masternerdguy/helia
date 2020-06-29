import { WSShip } from '../entities/wsShip';
import { WSSystem } from '../entities/wsSystem';

export class ClientJoinBody {
    sid: string;
}

export class ServerJoinBody {
    currentShipInfo: WSShip;
    currentSystemInfo: WSSystem;
}
