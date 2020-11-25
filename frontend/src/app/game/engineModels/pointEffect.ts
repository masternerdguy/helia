import { PointEffectData, PointEffectRegistry } from '../data/pointEffectData';
import { WsPushPointEffect } from '../wsModels/entities/wsPushPointEffect';
import { Camera } from './camera';
import { Player } from './player';

export class PointEffect extends WsPushPointEffect {
    player: Player;
    vfxData: PointEffectData;

    maxLifeTime: number;
    lifeElapsed = 0;
    finished = false;

    lastUpdateTime: number;

    constructor(b: WsPushPointEffect, player: Player) {
        // assign values
        super();
        Object.assign(this, b);
        this.player = player;

        // get effect data
        const registry = new PointEffectRegistry();

        if (b.gfxEffect === 'basic_explosion') {
            this.vfxData = registry.basicExplosion();
        }

        this.maxLifeTime = this.vfxData?.duration ?? 0;

        // set frame time
        this.lastUpdateTime = Date.now();
    }

    periodicUpdate() {
        // get time
        const now = Date.now();

        // increase lifetime elapsed
        this.lifeElapsed += (now - this.lastUpdateTime);

        // check for finish
        if (this.lifeElapsed >= this.maxLifeTime) {
            this.finished = true;
        }

        // store frame time
        this.lastUpdateTime = now;
    }

    render(ctx: any, camera: Camera) {
        if (this.vfxData) {
            if (this.vfxData.type === 'point_explosion') {
                // project to screen
                const sx = camera.projectX(this.x);
                const sy = camera.projectY(this.y);
                const sr = camera.projectR(this.r);

                // backup filter
                const oldFilter = ctx.filter;

                // style explosion
                ctx.fillStyle = this.vfxData.color;
                if (this.vfxData.filter) {
                    ctx.filter = this.vfxData.filter;
                }

                // use elapsed lifetime ratio to shrink radius
                const er = Math.max(0, sr * (1 - this.lifeElapsed / this.maxLifeTime));

                // draw explosion
                ctx.beginPath();
                ctx.arc(sx, sy, er, 0, 2 * Math.PI);
                ctx.fill();

                // restore filter
                ctx.filter = oldFilter;
            }
        }
    }
}
