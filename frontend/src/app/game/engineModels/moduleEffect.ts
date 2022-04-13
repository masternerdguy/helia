import { WsPushModuleEffect } from '../wsModels/entities/wsPushModuleEffect';
import {
  ModuleActivationEffectData,
  ModuleActivationEffectRepository,
} from '../data/moduleActivationEffectData';
import { Camera } from './camera';
import { Player, TargetType } from './player';
import { Station } from './station';
import { Ship } from './ship';
import { Asteroid } from './asteroid';
import { angleBetween, magnitude } from '../engineMath';

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

    if (b.gfxEffect === 'le_banhammer') {
      this.vfxData = repo.leBanhammer();
    } else if (b.gfxEffect === 'basic_laser_tool') {
      this.vfxData = repo.basicLaserTool();
    } else if (b.gfxEffect === 'basic_gauss_rifle') {
      this.vfxData = repo.basicGaussRifle();
    } else if (b.gfxEffect === 'basic_ice_miner') {
      this.vfxData = repo.basicIceMiner();
    } else if (b.gfxEffect === 'basic_shield_booster') {
      this.vfxData = repo.basicShieldBooster();
    } else if (b.gfxEffect === 'basic_auto-5_cannon') {
      this.vfxData = repo.basicAuto5Cannon();
    } else if (b.gfxEffect === 'basic_aether_dragger') {
      this.vfxData = repo.basicAetherDragger();
    } else if (b.gfxEffect === 'basic_ice_harvester') {
      this.vfxData = repo.basicIceHarvester();
    } else if (b.gfxEffect === 'basic_ore_harvester') {
      this.vfxData = repo.basicOreHarvester();
    } else if (b.gfxEffect === 'basic_energy_siphon') {
      this.vfxData = repo.basicUtilitySiphon();
    } else if (b.gfxEffect === 'miel_ice_harvester') {
      this.vfxData = repo.mielIceHarvester();
    } else if (b.gfxEffect === 'leche_ore_harvester') {
      this.vfxData = repo.lecheOreHarvester();
    } else if (b.gfxEffect === 'basic_salvager') {
      this.vfxData = repo.basicSalvager();
    } else if (b.gfxEffect === 'basic_shield_laser') {
      this.vfxData = repo.basicShieldLaser();
    } else if (b.gfxEffect === 'basic_hull_laser') {
      this.vfxData = repo.basicHullLaser();
    } else if (b.gfxEffect === 'basic_hull_laser') {
      this.vfxData = repo.basicGeneralLaser();
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
    } else if (this.objStartType === TargetType.Asteroid) {
      for (const s of this.player.currentSystem.asteroids) {
        if (s.id === this.objStartID) {
          this.objStart = s;
          break;
        }
      }
    } else if (this.objStartType === TargetType.Wreck) {
      for (const s of this.player.currentSystem.wrecks) {
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
      } else if (this.objEndType === TargetType.Asteroid) {
        for (const s of this.player.currentSystem.asteroids) {
          if (s.id === this.objEndID) {
            this.objEnd = s;
            break;
          }
        }
      } else if (this.objEndType === TargetType.Wreck) {
        for (const s of this.player.currentSystem.wrecks) {
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
    this.lifeElapsed += now - this.lastUpdateTime;

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
        this.renderAsLaserEffect(camera, ctx);
      } else if (this.vfxData.type === 'gauss') {
        this.renderAsGaussEffect(camera, ctx);
      } else if (this.vfxData.type === 'autocannon') {
        this.renderAsAutocannonEffect(camera, ctx);
      } else if (this.vfxData.type === 'bubble_shield_boost') {
        this.renderAsBubbleShieldBoostEffect(camera, ctx);
      } else if (this.vfxData.type === 'aether_drag') {
        this.renderAsAetherDragEffect(camera, ctx);
      } else if (this.vfxData.type === 'siphon') {
        this.renderAsSiphonEffect(camera, ctx);
      } else if (this.vfxData.type === 'salvager') {
        this.renderAsSalvagerEffect(camera, ctx);
      }
    }
  }

  private renderAsSalvagerEffect(camera: Camera, ctx: any) {
    if (this.objStart && this.objEnd) {
      // get end-point coordinates
      const src = getTargetCoordinatesAndRadius(
        this.objStart,
        this.objStartType,
        this.objStartHPOffset
      );
      const dest = getTargetCoordinatesAndRadius(this.objEnd, this.objEndType);

      // apply offset to destination coordinates for cooler effect
      if (!this.endPosOffset) {
        // get a random point within the radius of the target
        const bR = dest[2] / 1.5;

        const ox = randomIntFromInterval(-bR, bR);
        const oy = randomIntFromInterval(-bR, bR);

        // store offset
        this.endPosOffset = [ox, oy];
      }

      dest[0] += this.endPosOffset[0];
      dest[1] += this.endPosOffset[1];

      // project to screen
      const sx = camera.projectX(src[0]);
      const sy = camera.projectY(src[1]);

      const tx = camera.projectX(dest[0]);
      const ty = camera.projectY(dest[1]);

      // project salvage beam thickness
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

  private renderAsSiphonEffect(camera: Camera, ctx: any) {
    if (this.objStart && this.objEnd) {
      // get end-point coordinates
      const src = getTargetCoordinatesAndRadius(
        this.objStart,
        this.objStartType,
        this.objStartHPOffset
      );
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

      // project to screen
      const sx = camera.projectX(src[0]);
      const sy = camera.projectY(src[1]);

      const tx = camera.projectX(dest[0]);
      const ty = camera.projectY(dest[1]);

      // project siphon curve thickness
      const lt = camera.projectR(this.vfxData.thickness);

      // style curve
      ctx.strokeStyle = this.vfxData.color;

      const oldFilter = ctx.filter;

      if (this.vfxData.filter) {
        ctx.filter = this.vfxData.filter;
      }

      // animate curve effect
      const d = magnitude(sx, sy, tx, ty);
      const p = 1 - this.lifeElapsed / this.maxLifeTime;

      const ox = p * d;
      const oy = p * d;

      // draw curve
      ctx.beginPath();
      ctx.moveTo(sx, sy);
      ctx.quadraticCurveTo(sx + ox, sy + oy, tx, ty);
      ctx.lineWidth = lt;
      ctx.stroke();

      // revert filter
      ctx.filter = oldFilter;
    }
  }

  private renderAsAetherDragEffect(camera: Camera, ctx: any) {
    // get target coordinates
    const src = getTargetCoordinatesAndRadius(this.objEnd, this.objEndType);

    // project to screen
    const sx = camera.projectX(src[0]);
    const sy = camera.projectY(src[1]);
    const sr = camera.projectR(src[2]);
    const bt = camera.projectR(this.vfxData.thickness);

    // backup filter
    const oldFilter = ctx.filter;

    // style boost
    ctx.strokeStyle = this.vfxData.color;
    if (this.vfxData.filter) {
      ctx.filter = this.vfxData.filter;
      ctx.lineWidth = bt;
    }

    // use elapsed lifetime ratio to contract radius
    const er = Math.max(0, sr * (1 - this.lifeElapsed / this.maxLifeTime));

    // draw aether drag field
    ctx.beginPath();
    ctx.arc(sx, sy, er, 0, 2 * Math.PI);
    ctx.stroke();

    // restore filter
    ctx.filter = oldFilter;
  }

  private renderAsBubbleShieldBoostEffect(camera: Camera, ctx: any) {
    // get start coordinates
    const src = getTargetCoordinatesAndRadius(
      this.objStart,
      this.objStartType,
      this.objStartHPOffset
    );

    // project to screen
    const sx = camera.projectX(src[0]);
    const sy = camera.projectY(src[1]);
    const sr = camera.projectR(src[2]);
    const bt = camera.projectR(this.vfxData.thickness);

    // backup filter
    const oldFilter = ctx.filter;

    // style boost
    ctx.strokeStyle = this.vfxData.color;
    if (this.vfxData.filter) {
      ctx.filter = this.vfxData.filter;
      ctx.lineWidth = bt;
    }

    // use elapsed lifetime ratio to expand radius
    const er = Math.max(0, sr * (this.lifeElapsed / this.maxLifeTime));

    // draw boost
    ctx.beginPath();
    ctx.arc(sx, sy, er, 0, 2 * Math.PI);
    ctx.stroke();

    // restore filter
    ctx.filter = oldFilter;
  }

  private renderAsAutocannonEffect(camera: Camera, ctx: any) {
    if (this.objStart && this.objEnd) {
      // get end-point coordinates
      const src = getTargetCoordinatesAndRadius(
        this.objStart,
        this.objStartType,
        this.objStartHPOffset
      );
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

      // project to screen
      const sx = camera.projectX(src[0]);
      const sy = camera.projectY(src[1]);

      const tx = camera.projectX(dest[0]);
      const ty = camera.projectY(dest[1]);

      // project autocannon trail thickness
      const phase = Math.cos(this.lifeElapsed) / (2 * Math.PI);
      const lt = camera.projectR(this.vfxData.thickness * phase);

      // style line
      ctx.strokeStyle = this.vfxData.color;

      const oldFilter = ctx.filter;

      if (this.vfxData.filter) {
        ctx.filter = this.vfxData.filter;
      }

      // get line length and angle
      const llm = magnitude(sx, sy, tx, ty) * phase;
      const lla = angleBetween(tx, ty, sx, sy) / (Math.PI / 180);

      // determine actual end
      const txA = tx + llm * Math.cos(lla);
      const tyA = ty + llm * Math.sin(lla);

      // draw line
      ctx.beginPath();
      ctx.moveTo(sx, sy);
      ctx.lineTo(txA, tyA);
      ctx.lineWidth = lt;
      ctx.stroke();

      // revert filter
      ctx.filter = oldFilter;
    }
  }

  private renderAsGaussEffect(camera: Camera, ctx: any) {
    if (this.objStart && this.objEnd) {
      // get end-point coordinates
      const src = getTargetCoordinatesAndRadius(
        this.objStart,
        this.objStartType,
        this.objStartHPOffset
      );
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

      // project to screen
      const sx = camera.projectX(src[0]);
      const sy = camera.projectY(src[1]);

      const tx = camera.projectX(dest[0]);
      const ty = camera.projectY(dest[1]);

      // project gauss trail thickness
      const decay = 1 - this.lifeElapsed / this.maxLifeTime;
      const lt = camera.projectR(this.vfxData.thickness * decay);

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

  private renderAsLaserEffect(camera: Camera, ctx: any) {
    if (this.objStart && this.objEnd) {
      // get end-point coordinates
      const src = getTargetCoordinatesAndRadius(
        this.objStart,
        this.objStartType,
        this.objStartHPOffset
      );
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

function getTargetCoordinatesAndRadius(
  tgt: any,
  tgtType: TargetType,
  hpPos?: number[]
): [number, number, number] {
  if (tgtType === TargetType.Station) {
    const st = tgt as Station;
    return [st.x, st.y, st.radius];
  }

  if (tgtType === TargetType.Ship) {
    const s = tgt as Ship;

    if (hpPos?.length != 2) {
      return [s.x, s.y, s.radius];
    } else {
      const ox = s.getHardpointPosition(hpPos);
      return [ox[0], ox[1], s.radius];
    }
  }

  if (tgtType === TargetType.Asteroid) {
    const s = tgt as Asteroid;
    return [s.x, s.y, s.radius];
  }

  return [tgt?.x, tgt?.y, tgt?.radius];
}

function randomIntFromInterval(min: number, max: number): number {
  return Math.floor(Math.random() * (max - min + 1) + min);
}
