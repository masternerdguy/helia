import { Camera } from './camera';
import { WSJumphole } from '../wsModels/entities/wsJumphole';

export class Jumphole extends WSJumphole {
    texture2d: HTMLImageElement;

    constructor(ws: WSJumphole) {
        super();

        // copy from ws model
        this.id = ws.id;
        this.systemId = ws.systemId;
        this.outSystemId = ws.outSystemId;
        this.jumpholeName = ws.jumpholeName;
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
            this.texture2d.src = '/assets/jumpholes/' + this.texture + '.png';
        }

        // project to screen
        const sx = camera.projectX(this.x);
        const sy = camera.projectY(this.y);
        const sr = camera.projectR(this.radius);

        // draw debug bounding circle
        ctx.beginPath();
        ctx.arc(sx, sy, sr, 0, 2 * Math.PI, false);
        ctx.lineWidth = 2;
        ctx.strokeStyle = 'orange';
        ctx.stroke();

        // draw jumphole
        ctx.save();
        ctx.translate(sx, sy);
        ctx.rotate(((this.theta) * (Math.PI / -180)) + (Math.PI / 2));
        ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
        ctx.restore();
    }

    sync(ws: WSJumphole) {
        this.id = ws.id;
        this.systemId = ws.systemId;
        this.outSystemId = ws.outSystemId;
        this.jumpholeName = ws.jumpholeName;
        this.x = ws.x;
        this.y = ws.y;
        this.texture = ws.texture;
        this.mass = ws.mass;
        this.radius = ws.radius;
        this.theta = ws.theta;
    }
}
