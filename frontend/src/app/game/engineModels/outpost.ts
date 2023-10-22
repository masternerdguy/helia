import { Camera } from './camera';
import { WSOutpost } from '../wsModels/entities/wsOutpost';
import { Faction } from './faction';
import {
  GetFactionCacheEntry,
  GetPlayerFactionRelationshipCacheEntry,
} from '../wsModels/shared';

export class Outpost extends WSOutpost {
  texture2d: HTMLImageElement;
  isTargeted: boolean;

  faction: Faction;
  lastSeen: number;

  constructor(ws: WSOutpost) {
    super();

    // copy from ws model
    this.id = ws.id;
    this.systemId = ws.systemId;
    this.outpostName = ws.outpostName;
    this.x = ws.x;
    this.y = ws.y;
    this.texture = ws.texture;
    this.mass = ws.mass;
    this.radius = ws.radius;
    this.shieldP = ws.shieldP;
    this.armorP = ws.armorP;
    this.hullP = ws.hullP;
    this.theta = ws.theta;
    this.factionId = ws.factionId;

    this.lastSeen = Date.now();
  }

  render(ctx: any, camera: Camera) {
    // set up texture
    if (!this.texture2d) {
      this.texture2d = new Image();
      this.texture2d.src = '/assets/outposts/' + this.texture + '.png';
    }

    // project to screen
    const sx = camera.projectX(this.x);
    const sy = camera.projectY(this.y);
    const sr = camera.projectR(this.radius);

    // draw bounding circle
    ctx.beginPath();
    ctx.arc(sx, sy, Math.max(sr, 1.5), 0, 2 * Math.PI, false);
    ctx.lineWidth = 2;

    // select color by status and owner
    if (this.isTargeted) {
      ctx.strokeStyle = 'yellow';
    } else {
      ctx.strokeStyle = this.getStandingColor();
    }

    ctx.stroke();

    // draw outpost
    ctx.save();
    ctx.translate(sx, sy);
    ctx.rotate(this.theta * (Math.PI / -180) + Math.PI / 2);
    ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
    ctx.restore();
  }

  sync(ws: WSOutpost) {
    this.id = ws.id;
    this.systemId = ws.systemId;
    this.outpostName = ws.outpostName;
    this.x = ws.x;
    this.y = ws.y;
    this.texture = ws.texture;
    this.mass = ws.mass;
    this.radius = ws.radius;
    this.theta = ws.theta;
    this.shieldP = ws.shieldP;
    this.armorP = ws.armorP;
    this.hullP = ws.hullP;
    this.factionId = ws.factionId;

    this.lastSeen = Date.now();
  }

  getFaction(): Faction {
    return GetFactionCacheEntry(this.factionId);
  }

  getStandingColor() {
    const rep = GetPlayerFactionRelationshipCacheEntry(this.factionId);

    if (!rep) {
      return 'antiquewhite';
    }

    if (rep.isMember) {
      if (rep.openlyHostile) {
        return 'firebrick';
      } else {
        return 'lightgreen';
      }
    }

    if (rep.standingValue >= 6) {
      return 'royalblue';
    }

    if (rep.standingValue > 1.999) {
      return 'skyblue';
    }

    if (rep.standingValue <= -6) {
      return 'orangered';
    }

    if (rep.standingValue < -1.999) {
      return 'darkorange';
    }

    return 'antiquewhite';
  }
}
