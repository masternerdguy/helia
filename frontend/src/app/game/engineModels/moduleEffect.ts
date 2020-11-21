import { WsPushModuleEffect } from '../wsModels/entities/wsPushModuleEffect';
import { ModuleActivationEffectData, ModuleActivationEffectRepository } from '../data/moduleActivationEffectData';
import { Camera } from './camera';
import { Player } from './player';

export class ModuleEffect extends WsPushModuleEffect {
    vfxData: ModuleActivationEffectData;
    player: Player;

    maxLifeTime = 0;
    lifeElapsed = 0;

    lastUpdateTime: number;
    finished = false;

    constructor(b: WsPushModuleEffect, player: Player) {
        // assign values
        super();
        Object.assign(this, b);
        this.player = player;

        // get effect data
        const repo = new ModuleActivationEffectRepository();

        if (b.gfxEffect === 'basic_laser_tool') {
            this.vfxData = repo.basicLaserTool();
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
            if (this.vfxData.type === 'laser') {
                /* draw a line of the given color from source to destination */

                console.log(this);
            }
        }
    }
}
