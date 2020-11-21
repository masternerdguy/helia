import { WsPushModuleEffect } from '../wsModels/entities/wsPushModuleEffect';
import { ModuleActivationEffectData, ModuleActivationEffectRepository } from '../data/moduleActivationEffectData';
import { Camera } from './camera';
import { Player, TargetType } from './player';

export class ModuleEffect extends WsPushModuleEffect {
    vfxData: ModuleActivationEffectData;
    player: Player;

    maxLifeTime = 0;
    lifeElapsed = 0;

    lastUpdateTime: number;
    finished = false;

    objStart: any;
    objEnd: any;

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

        // locate start object
        if (this.objStartType === TargetType.Station) {
            for (const st of this.player.currentSystem.stations) {
                if (st.id === this.objStartID) {
                    this.objStart = st;
                    break;
                }
            }
        } else if (this.objStartType === TargetType.Ship) {
            for (const s of this.player.currentSystem.ships) {
                if (s.id === this.objStartID) {
                    this.objStart = s;
                    break;
                }
            }
        }

        // locate end object if present
        if (this.objEndID) {
            if (this.objEndType === TargetType.Station) {
                for (const st of this.player.currentSystem.stations) {
                    if (st.id === this.objEndID) {
                        this.objEnd = st;
                        break;
                    }
                }
            } else if (this.objEndType === TargetType.Ship) {
                for (const s of this.player.currentSystem.ships) {
                    if (s.id === this.objEndID) {
                        this.objEnd = s;
                        break;
                    }
                }
            }
        }

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
