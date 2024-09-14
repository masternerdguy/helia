import { WSSystem } from '../entities/wsSystem';
import { WSShip } from '../entities/wsShip';
import { WSStar } from '../entities/wsStar';
import { WSPlanet } from '../entities/wsPlanet';
import { WSStation } from '../entities/wsStation';
import { WSJumphole } from '../entities/wsJumphole';
import { WsPushModuleEffect } from '../entities/wsPushModuleEffect';
import { WsPushPointEffect } from '../entities/wsPushPointEffect';
import { WSAsteroid } from '../entities/wsAsteroid';
import { WSMissile } from '../entities/wsMissile';
import { WSSystemChatMessage } from '../entities/wsSystemChatMessage';
import { WSWreck } from '../entities/wsWreck';
import { WSOutpost } from '../entities/wsOutpost';
import { WSArtifact } from '../entities/wsArtifact';

export class ServerGlobalUpdateBody {
  currentSystemInfo: WSSystem;
  ships: WSShip[];
  stars: WSStar[];
  planets: WSPlanet[];
  artifacts: WSArtifact[];
  jumpholes: WSJumphole[];
  stations: WSStation[];
  outposts: WSOutpost[];
  asteroids: WSAsteroid[];
  newModuleEffects: WsPushModuleEffect[];
  newPointEffects: WsPushPointEffect[];
  missiles: WSMissile[];
  wrecks: WSWreck[];
  systemChat: WSSystemChatMessage[];
  token: number;
}
