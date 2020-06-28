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

        // draw ship
        ctx.drawImage(this.texture2d, sx - (this.texture2d.width / 2), sy - (this.texture2d.height / 2));
    }

    sync(sh: WSShip) {
        this.createdAt = sh.createdAt;
        this.shipName = sh.shipName;
        this.uid = sh.uid;
        this.systemId = sh.systemId;
        this.x = sh.x;
        this.y = sh.y;
    }
}
