import { WSSystem } from '../entities/wsSystem';
import { WSShip } from '../entities/wsShip';
import { WSStar } from '../entities/wsStar';
import { WSPlanet } from '../entities/wsPlanet';
import { WSStation } from '../entities/wsStation';
import { WSJumphole } from '../entities/wsJumphole';

export class ServerGlobalUpdateBody {
    currentSystemInfo: WSSystem;
    ships: WSShip[];
    stars: WSStar[];
    planets: WSPlanet[];
    jumpholes: WSJumphole[];
    stations: WSStation[];
}
