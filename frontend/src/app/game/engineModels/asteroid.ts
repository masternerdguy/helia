import { Camera } from './camera';
import { WSAsteroid } from '../wsModels/entities/wsAsteroid';

export class Asteroid extends WSAsteroid {
  texture2d: HTMLImageElement;
  isTargeted: boolean;

  constructor(ws: WSAsteroid) {
    super();

    // copy from ws model
    this.id = ws.id;
    this.systemId = ws.systemId;
    this.name = ws.name;
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
      this.texture2d.src = '/assets/asteroids/' + this.texture + '.png';
    }

    // project to screen
    const sx = camera.projectX(this.x);
    const sy = camera.projectY(this.y);
    const sr = camera.projectR(this.radius);

    // draw debug bounding circle
    if (this.isTargeted) {
      ctx.strokeStyle = 'yellow';
    } else {
      ctx.strokeStyle = 'aliceblue';
    }

    ctx.beginPath();
    // use a 10% padding due to irregular shapes
    ctx.arc(sx, sy, Math.max(sr * 1.1, 1.5), 0, 2 * Math.PI, false);
    ctx.lineWidth = 2;
    ctx.stroke();

    // draw asteroid
    ctx.save();
    ctx.translate(sx, sy);
    ctx.rotate(this.theta * (Math.PI / -180) + Math.PI / 2);
    ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
    ctx.restore();
  }

  sync(ws: WSAsteroid) {
    this.id = ws.id;
    this.systemId = ws.systemId;
    this.name = ws.name;
    this.x = ws.x;
    this.y = ws.y;
    this.texture = ws.texture;
    this.mass = ws.mass;
    this.radius = ws.radius;
    this.theta = ws.theta;
  }
}
