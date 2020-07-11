import { WSSystem } from '../entities/wsSystem';
import { WSShip } from '../entities/wsShip';
import { WSStar } from '../entities/wsStar';
import { WSPlanet } from '../entities/wsPlanet';
import { WSStation } from '../entities/wsStation';

export class ServerGlobalUpdateBody {
    currentSystemInfo: WSSystem;
    ships: WSShip[];
    stars: WSStar[];
    planets: WSPlanet[];
    stations: WSStation[];
}
