import { Camera } from './camera';
import { WSPlanet } from '../wsModels/entities/wsPlanet';
import '../procedural/planet/main.js';

export class Planet extends WSPlanet {
  texture2d: HTMLImageElement;
  isTargeted: boolean;
  proceduralBag: any;
  proceduralStart: boolean;

  constructor(ws: WSPlanet) {
    super();

    // copy from ws model
    this.id = ws.id;
    this.systemId = ws.systemId;
    this.planetName = ws.planetName;
    this.x = ws.x;
    this.y = ws.y;
    this.texture = ws.texture;
    this.mass = ws.mass;
    this.radius = ws.radius;
    this.theta = ws.theta;

    // hook procedural generation functions
    this.proceduralBag = (window as any).SteffTek_planet_js;
    this.proceduralStart = false;
  }

  render(ctx: any, camera: Camera) {
    // set up texture
    if (!this.texture2d) {
      this.texture2d = new Image();
      this.texture2d.src = '/assets/planets/' + this.texture + '.png';

      if (!this.proceduralStart) {
        this.proceduralStart = true;
        // procedural gen test
        let colors = {
          land_color: '#4cfa69', //Color of the Main Land
          beach_color: '#e9fe6d', //Color of the Beaches
          shore_color: '#78dffb', //Color of the Shores
          ocean_color: '#0744a6', //Color of the Deep Ocean
          mountain_color: '#854d1d', //Color of the Mountains
          mountain_top_color: '#ffffff', //Color of the Mountain Top (e.g. Snow)
          crater_color: '#8b9e90', //Main Color of Craters
          pole_color: '#BDDEEC', //Color of Poles (Ice)
          cloud_color: '#ffffff', //Cloud Color
          cloud_opacity: 70, //Cloud Base Opacity
          atmosphere_color: '#4F7AAD', //Atmosphere Color
          atmosphere_opacity: 40, //Atmosphere Opacity/Density
          shading_level: 2, //Shading Level (Float 0-2, 2 = Maximum)
          add_detail: true, //Plain Map or a bit more detailed?
        };

        let planet_options = {
          planet_radius: 20, //Planet Radius
          atmosphere_radius: 20, //Atmosphere Radius
          sea_level: 0.42, // ALL LEVELS ARE VALUES BETWEEN 0 AND 1
          shore_level: 0.48,
          beach_level: 0.5,
          mountain_level: 0.62,
          mountain_top_level: 0.75,
          cloud_level: 0.62, // CLOUD LEVEL IS CUSTOM GENERATED AND NOT AFFECTED BY THE OTHER LEVELS
          cloud_radius: 420, //Cloud Radius
          pole_level: 0.5, //How big the Poles should be (Float 0-2, 2 = Full Coverage)
          craters: true, //Should Craters Spawn?
          clouds: true, //Should Clouds Spawn?
          atmosphere: true, //Should the Planet have an atmosphere
          poles: true, //Should the Planet have icy poles?
          hard_pole_lines: false, //Should the pole line be a hard or a soft cut?
        };

        let generator_options = {
          octaveCount: 9, //Perlin Noise Octave (How Often)
          amplitude: 5, //Perlin Noise Amp (How Big)
          persistence: 0.5, //Perlin Noise persistence (How Smooth, smaller number = smoother)
        };

        let cloud_generator = {
          octaveCount: 6,
          amplitude: 6,
          persistence: 0.4,
        };

        let size = 40; //Control the ImageSize
        let seed = 'ASDFG';

        console.log('stand by 1');

        var rt = async () => {
          console.log('stand by 2');

          let image = this.proceduralBag.generatePlanet(
            size,
            planet_options,
            colors,
            seed,
            generator_options,
            cloud_generator,
          );

          console.log('stand by 3');

          // debug out
          console.log(image);
        };

        rt();
      }
    }

    // project to screen
    const sx = camera.projectX(this.x);
    const sy = camera.projectY(this.y);
    const sr = camera.projectR(this.radius);

    // draw bounding circle
    if (this.isTargeted) {
      ctx.strokeStyle = 'yellow';
    } else {
      ctx.strokeStyle = 'cyan';
    }

    ctx.beginPath();
    ctx.arc(sx, sy, Math.max(sr, 2), 0, 2 * Math.PI, false);
    ctx.lineWidth = 2;
    ctx.stroke();

    // draw planet
    ctx.save();
    ctx.translate(sx, sy);
    ctx.rotate(this.theta * (Math.PI / -180) + Math.PI / 2);
    ctx.drawImage(this.texture2d, -sr, -sr, sr * 2, sr * 2);
    ctx.restore();
  }

  sync(ws: WSPlanet) {
    this.id = ws.id;
    this.systemId = ws.systemId;
    this.planetName = ws.planetName;
    this.x = ws.x;
    this.y = ws.y;
    this.texture = ws.texture;
    this.mass = ws.mass;
    this.radius = ws.radius;
    this.theta = ws.theta;
  }
}
