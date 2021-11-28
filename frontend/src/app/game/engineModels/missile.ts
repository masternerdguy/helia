import { angleBetween } from '../engineMath';
import { WSMissile } from '../wsModels/entities/wsMissile';
import { Camera } from './camera';

export class Missile extends WSMissile {
  texture2d: HTMLImageElement;
  isTargeted: boolean;
  lastX: number;
  lastY: number;
  lastSeen: number;

  constructor(ws: WSMissile) {
    super();

    // copy from ws model
    this.id = ws.id;
    this.x = this.lastX = ws.x;
    this.y = this.lastY = ws.y;
    this.r = ws.r;
    this.t = ws.t;
  }

  render(ctx: any, camera: Camera) {
    // set up texture
    if (!this.texture2d) {
      this.texture2d = new Image();
      this.texture2d.src = '/assets/missiles/' + this.t + '.png';
    }

    // project to screen
    const sx = camera.projectX(this.x);
    const sy = camera.projectY(this.y);
    const sr = camera.projectR(this.r);

    // determine direction
    const theta = angleBetween(this.lastX, this.lastY, this.x, this.y);

    // draw missile
    ctx.save();
    ctx.translate(sx, sy);
    ctx.rotate(theta * (Math.PI / -180) + Math.PI / 2);
    ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
    ctx.restore();
  }

  sync(ws: WSMissile) {
    this.lastSeen = Date.now()

    // stash current position
    this.lastX = this.x;
    this.lastY = this.y;

    // copy new values
    this.id = ws.id;
    this.x = ws.x;
    this.y = ws.y;
    this.r = ws.r;
    this.t = ws.t;
  }
}
