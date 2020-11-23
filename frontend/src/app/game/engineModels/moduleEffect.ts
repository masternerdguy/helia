import { WsPushModuleEffect } from '../wsModels/entities/wsPushModuleEffect';
import { ModuleActivationEffectData, ModuleActivationEffectRepository } from '../data/moduleActivationEffectData';
import { Camera } from './camera';
import { Player, TargetType } from './player';
import { Station } from './station';
import { Ship } from './ship';

export class ModuleEffect extends WsPushModuleEffect {
    vfxData: ModuleActivationEffectData;
    player: Player;

    maxLifeTime = 0;
    lifeElapsed = 0;

    lastUpdateTime: number;
    finished = false;

    objStart: any;
    objEnd: any;

    endPosOffset: [number, number];

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
                if (this.objStart && this.objEnd) {
                    // get end-point coordinates
                    const src = getTargetCoordinatesAndRadius(this.objStart, this.objStartType);
                    const dest = getTargetCoordinatesAndRadius(this.objEnd, this.objEndType);

                    // apply offset to destination coordinates for cooler effect
                    if (!this.endPosOffset) {
                        // get a random point within the radius of the target
                        const bR = dest[2] / 3;

                        const ox = randomIntFromInterval(-bR, bR);
                        const oy = randomIntFromInterval(-bR, bR);

                        // store offset
                        this.endPosOffset = [ox, oy];
                    }

                    dest[0] += this.endPosOffset[0];
                    dest[1] += this.endPosOffset[1];

                    // todo: implement hardpoint offset for source ship

                    // project to screen
                    const sx = camera.projectX(src[0]);
                    const sy = camera.projectY(src[1]);

                    const tx = camera.projectX(dest[0]);
                    const ty = camera.projectY(dest[1]);

                    // project laser beam thickness
                    const lt = camera.projectR(this.vfxData.thickness);

                    // style line
                    ctx.strokeStyle = this.vfxData.color;

                    const oldFilter = ctx.filter;

                    if (this.vfxData.filter) {
                        ctx.filter = this.vfxData.filter;
                    }

                    // draw line
                    ctx.beginPath();
                    ctx.moveTo(sx, sy);
                    ctx.lineTo(tx, ty);
                    ctx.lineWidth = lt;
                    ctx.stroke();

                    // revert filter
                    ctx.filter = oldFilter;
                }
            }
        }
    }
}

function getTargetCoordinatesAndRadius(tgt: any, tgtType: TargetType): [number, number, number] {
    if (tgtType === TargetType.Station) {
        const st = tgt as Station;
        return [st.x, st.y, st.radius];
    }

    if (tgtType === TargetType.Ship) {
        const s = tgt as Ship;
        return [s.x, s.y, s.radius];
    }

    return [tgt?.x, tgt?.y, tgt?.radius];
}

function randomIntFromInterval(min: number, max: number): number {
    return Math.floor(Math.random() * (max - min + 1) + min);
}