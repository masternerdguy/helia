import { WSStar } from '../wsModels/entities/wsStar';
import { Camera } from './camera';

export class Star extends WSStar {
  texture2d: HTMLImageElement;
  isTargeted: boolean;

  constructor(ws: WSStar) {
    super();

    // copy from ws model
    this.id = ws.id;
    this.systemId = ws.systemId;
    this.x = ws.x;
    this.y = ws.y;
    this.texture = ws.texture;
    this.mass = ws.mass;
    this.radius = ws.radius;
    this.theta = ws.theta;
  }

  render(ctx: any, camera: Camera) {
    // set up texture
    if (!this.texture2d) {
      this.texture2d = new Image();
      this.texture2d.src = '/assets/stars/' + this.texture + '.png';
    }

    // project to screen
    const sx = camera.projectX(this.x);
    const sy = camera.projectY(this.y);
    const sr = camera.projectR(this.radius);

    // draw debug bounding circle
    if (this.isTargeted) {
      ctx.strokeStyle = 'yellow';
    } else {
      ctx.strokeStyle = 'cyan';
    }

    ctx.beginPath();
    ctx.arc(sx, sy, Math.max(sr, 3), 0, 2 * Math.PI, false);
    ctx.lineWidth = 2;
    ctx.stroke();

    // draw star
    ctx.save();
    ctx.translate(sx, sy);
    ctx.rotate(this.theta * (Math.PI / -180) + Math.PI / 2);
    ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
    ctx.restore();
  }

  sync(ws: WSStar) {
    this.id = ws.id;
    this.systemId = ws.systemId;
    this.x = ws.x;
    this.y = ws.y;
    this.texture = ws.texture;
    this.mass = ws.mass;
    this.radius = ws.radius;
    this.theta = ws.theta;
  }
}
