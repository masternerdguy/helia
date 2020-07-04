import { WSShip } from '../wsModels/entities/wsShip';
import { Camera } from './camera';

export class Ship extends WSShip {
    texture2d: HTMLImageElement;

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

        if (ws.accel) {
            this.accel = ws.accel;
        }

        if (ws.turn) {
            this.turn = ws.turn;
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

        // draw debug bounding circle
        ctx.beginPath();
        ctx.arc(sx, sy, sr, 0, 2 * Math.PI, false);
        ctx.lineWidth = 2;
        ctx.strokeStyle = 'white';
        ctx.stroke();

        // draw ship
        ctx.save();
        ctx.translate(sx, sy);
        ctx.rotate(((this.theta) * (Math.PI / -180)) + (Math.PI / 2));
        ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
        ctx.restore();
    }

    sync(sh: WSShip) {
        this.createdAt = sh.createdAt;
        this.shipName = sh.shipName;
        this.uid = sh.uid;
        this.systemId = sh.systemId;
        this.x = sh.x;
        this.y = sh.y;
        this.theta = sh.theta;
        this.velX = sh.velX;
        this.velY = sh.velY;
        this.mass = sh.mass;
        this.radius = sh.radius;

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
    }
}
