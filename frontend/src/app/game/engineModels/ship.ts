import { WSShip } from '../wsModels/entities/wsShip';
import { GetFactionCacheEntry } from '../wsModels/shared';
import { Camera } from './camera';
import { Faction } from './faction';

export class Ship extends WSShip {
  texture2d: HTMLImageElement;
  lastSeen: number;
  isTargeted: boolean;
  isPlayer: boolean;
  lastWsX: number;
  lastWsY: number;
  deltaX: number;
  deltaY: number;
  tps: number;

  faction: Faction;

  constructor(ws: WSShip) {
    super();
    // copy from ws model
    this.id = ws.id;
    this.createdAt = ws.createdAt;
    this.shipName = ws.shipName;
    this.systemId = ws.systemId;
    this.uid = ws.uid;
    this.x = ws.x;
    this.y = ws.y;
    this.lastWsX = ws.x;
    this.lastWsY = ws.y;
    this.texture = ws.texture;
    this.theta = ws.theta;
    this.velX = ws.velX;
    this.velY = ws.velY;
    this.radius = ws.radius;
    this.factionId = ws.factionId;
    this.lastSeen = Date.now();

    if (ws.accel) {
      this.accel = ws.accel;
    }

    if (ws.turn) {
      this.turn = ws.turn;
    }

    if (ws.shieldP) {
      this.shieldP = ws.shieldP;
    }

    if (ws.armorP) {
      this.armorP = ws.armorP;
    }

    if (ws.hullP) {
      this.hullP = ws.hullP;
    }

    if (ws.energyP) {
      this.energyP = ws.energyP;
    }

    if (ws.heatP) {
      this.heatP = ws.heatP;
    }

    if (ws.fuelP) {
      this.fuelP = ws.fuelP;
    }

    if (ws.dockedAtStationID) {
      this.dockedAtStationID = ws.dockedAtStationID;
    }

    if (ws.cargoP) {
      this.cargoP = ws.cargoP;
    }

    if (ws.wallet) {
      this.wallet = Math.round(ws.wallet);
    }
  }

  render(ctx: any, camera: Camera) {
    // set up texture
    if (!this.texture2d) {
      this.texture2d = new Image();
      this.texture2d.src = '/assets/ships/' + this.texture + '.png';
    }

    // project to screen
    const sx = camera.projectX(this.x);
    const sy = camera.projectY(this.y);
    const sr = camera.projectR(this.radius);

    // draw bounding circle
    ctx.beginPath();
    ctx.arc(sx, sy, Math.max(sr, 1.3), 0, 2 * Math.PI, false);
    ctx.lineWidth = 2;

    // select color by status and owner
    if (this.isTargeted) {
      ctx.strokeStyle = 'yellow';
    } else if (this.isPlayer) {
      ctx.strokeStyle = 'magenta';
    } else {
      ctx.strokeStyle = 'white';
    }

    ctx.stroke();

    // convert theta for rendering
    const st = this.theta * (Math.PI / -180) + Math.PI / 2;

    // draw ship
    ctx.save();
    ctx.translate(sx, sy);
    ctx.rotate(st);
    ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
    ctx.restore();

    // draw rack A hardpoint indicators
    const hpi = camera.projectR(1);

    if (this.fitStatus?.aRack?.modules) {
      for (const hp of this.fitStatus.aRack.modules) {
        // skip if invalid position
        if (hp.hpPos.length != 2) {
          continue;
        }

        // get hardpoint position and angle
        const hx = this.getHardpointPosition(hp.hpPos);

        // project to screen
        const shx = camera.projectX(hx[0]);
        const shy = camera.projectY(hx[1]);
        const shr = camera.projectR(0.33);

        // draw indicator circle
        ctx.beginPath();

        if (hp.isCycling) {
          ctx.strokeStyle = '#fc9003';
          ctx.arc(
            shx,
            shy,
            shr,
            0,
            2 * Math.PI * (hp.cyclePercent / 100),
            false
          );
        } else {
          if (hpi >= 5) {
            ctx.strokeStyle = '#b0b0b0';
            ctx.arc(shx, shy, shr, 0, 2 * Math.PI, false);
          }
        }

        ctx.lineWidth = 2;
        ctx.stroke();
      }
    }
  }

  sync(sh: WSShip) {
    const now = Date.now();

    // for interpolation
    this.deltaX = sh.x - this.lastWsX;
    this.deltaY = sh.y - this.lastWsY;
    this.tps = now - this.lastSeen;

    // copy from message
    this.createdAt = sh.createdAt;
    this.shipName = sh.shipName;
    this.ownerName = sh.ownerName;
    this.uid = sh.uid;
    this.systemId = sh.systemId;
    this.x = sh.x;
    this.y = sh.y;
    this.lastWsX = sh.x;
    this.lastWsY = sh.y;
    this.theta = sh.theta;
    this.velX = sh.velX;
    this.velY = sh.velY;
    this.mass = sh.mass;
    this.radius = sh.radius;
    this.lastSeen = now;

    // reset texture if changed
    if (this.texture !== sh.texture) {
      this.texture = sh.texture;
      this.texture2d = undefined;
    }

    if (sh.accel) {
      this.accel = sh.accel;
    }

    if (sh.turn) {
      this.turn = sh.turn;
    }

    if (sh.shieldP) {
      this.shieldP = sh.shieldP;
    }

    if (sh.armorP) {
      this.armorP = sh.armorP;
    }

    if (sh.hullP) {
      this.hullP = sh.hullP;
    }

    if (sh.energyP) {
      this.energyP = sh.energyP;
    }

    if (sh.heatP) {
      this.heatP = sh.heatP;
    }

    if (sh.fuelP) {
      this.fuelP = sh.fuelP;
    }

    if (sh.cargoP) {
      this.cargoP = sh.cargoP;
    }

    if (sh.dockedAtStationID) {
      this.dockedAtStationID = sh.dockedAtStationID;
    } else {
      this.dockedAtStationID = undefined;
    }

    if (sh.fitStatus) {
      this.fitStatus = sh.fitStatus;

      if (this.fitStatus.aRack.modules == null) {
        this.fitStatus.aRack.modules = [];
      }

      if (this.fitStatus.bRack.modules == null) {
        this.fitStatus.bRack.modules = [];
      }

      if (this.fitStatus.cRack.modules == null) {
        this.fitStatus.cRack.modules = [];
      }
    }

    if (sh.wallet) {
      this.wallet = Math.round(sh.wallet);
    }

    if (sh.factionId) {
      this.factionId = sh.factionId;
    }
  }

  getFaction(): Faction {
    return GetFactionCacheEntry(this.factionId);
  }

  getHardpointPosition(hpPos: number[]): [number, number] {
    if (hpPos?.length != 2) {
      return [this.x, this.y];
    }

    // get hardpoint radius and angle
    const hr = hpPos[0];
    const ht = (hpPos[1] + 360) % 360;

    // add hardpoint angle to ship
    const st = this.theta * (Math.PI / -180) + Math.PI / 2;
    const sht = st + (ht * (Math.PI / -180) + Math.PI / 2);

    // get cartesian coordinates of hardpoint
    const hx = this.x + Math.cos(sht) * hr;
    const hy = this.y + Math.sin(sht) * hr;

    // return hp position on screen
    return [hx, hy];
  }
}
