import { WSSystem } from '../entities/wsSystem';
import { WSShip } from '../entities/wsShip';
import { WSStar } from '../entities/wsStar';
import { WSPlanet } from '../entities/wsPlanet';
import { WSStation } from '../entities/wsStation';
import { WSJumphole } from '../entities/wsJumphole';
import { WsPushModuleEffect } from '../entities/wsPushModuleEffect';
import { WsPushPointEffect } from '../entities/wsPushPointEffect';

export class ServerGlobalUpdateBody {
    currentSystemInfo: WSSystem;
    ships: WSShip[];
    stars: WSStar[];
    planets: WSPlanet[];
    jumpholes: WSJumphole[];
    stations: WSStation[];
    newModuleEffects: WsPushModuleEffect[];
    newPointEffects: WsPushPointEffect[];
}
