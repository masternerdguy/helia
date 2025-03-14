import { WSSystem } from '../wsModels/entities/wsSystem';
import { Ship } from './ship';
import { Star } from './star';
import { Planet } from './planet';
import { Station } from './station';
import { Jumphole } from './jumphole';
import { ModuleEffect } from './moduleEffect';
import { PointEffect } from './pointEffect';
import { Asteroid } from './asteroid';
import { GetFactionCacheEntry } from '../wsModels/shared';
import { Faction } from './faction';
import { Missile } from './missile';
import { Wreck } from './wreck';
import { Outpost } from './outpost';
import { Artifact } from './artifact';

export class System extends WSSystem {
  constructor(ws: WSSystem) {
    super();

    // copy from ws model
    this.id = ws.id;
    this.systemName = ws.systemName;
    this.factionId = ws.factionId;

    // initialize self
    this.ships = [];
    this.stars = [];
    this.planets = [];
    this.artifacts = [];
    this.jumpholes = [];
    this.stations = [];
    this.outposts = [];
    this.asteroids = [];
    this.moduleEffects = [];
    this.pointEffects = [];
    this.missiles = [];
    this.wrecks = [];
  }

  ships: Ship[];
  stars: Star[];
  planets: Planet[];
  artifacts: Artifact[];
  jumpholes: Jumphole[];
  stations: Station[];
  outposts: Outpost[];
  asteroids: Asteroid[];
  moduleEffects: ModuleEffect[];
  pointEffects: PointEffect[];
  missiles: Missile[];
  wrecks: Wreck[];

  backplateImg: HTMLImageElement;
  backplateValid = false;

  getFaction(): Faction {
    return GetFactionCacheEntry(this.factionId);
  }
}
