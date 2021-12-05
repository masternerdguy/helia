import { WSShip } from '../wsModels/entities/wsShip';
import { GetFactionCacheEntry } from '../wsModels/shared';
import { Camera } from './camera';
import { Faction } from './faction';

export class Ship extends WSShip {
  texture2d: HTMLImageElement;
  lastSeen: number;
  isTargeted: boolean;
  isPlayer: boolean;

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

    // draw ship
    ctx.save();
    ctx.translate(sx, sy);
    ctx.rotate(this.theta * (Math.PI / -180) + Math.PI / 2);
    ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
    ctx.restore();

    console.log(this.fitStatus)

    // debug: draw rack A hardpoints
    if (this.fitStatus?.aRack?.modules){
      for (const hp of this.fitStatus.aRack.modules) {
        ctx.strokeStyle = "pink";
  
        if (hp.hpPos.length != 2) {
          continue;
        }
  
        const hr = hp.hpPos[0];
        const ht = hp.hpPos[1];
  
        const shr = camera.projectR(hr);
        const sht = ((this.theta * (Math.PI / -180) + Math.PI / 2) + ht) % 360;
  
        const hx = sx + (Math.cos(sht * (-180/Math.PI)) * shr);
        const hy = sy + (Math.sin(sht * (-180/Math.PI)) * shr);
  
        ctx.beginPath();
        ctx.arc(hx, hy, 1.5, 0, 2 * Math.PI, false);
        ctx.lineWidth = 2;
        ctx.stroke();
      }
    }
  }

  sync(sh: WSShip) {
    this.createdAt = sh.createdAt;
    this.shipName = sh.shipName;
    this.ownerName = sh.ownerName;
    this.uid = sh.uid;
    this.systemId = sh.systemId;
    this.x = sh.x;
    this.y = sh.y;
    this.theta = sh.theta;
    this.velX = sh.velX;
    this.velY = sh.velY;
    this.mass = sh.mass;
    this.radius = sh.radius;
    this.lastSeen = Date.now();

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
  }

  getFaction(): Faction {
    return GetFactionCacheEntry(this.factionId);
  }
}
