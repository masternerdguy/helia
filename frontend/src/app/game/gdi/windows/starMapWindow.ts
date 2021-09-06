import { Camera } from '../../engineModels/camera';
import { Player } from '../../engineModels/player';
import { WsService } from '../../ws.service';
import { ClientViewStarMap } from '../../wsModels/bodies/viewStarMap';
import { MessageTypes } from '../../wsModels/gameMessage';
import { GDIStyle } from '../base/gdiStyle';
import { GDIWindow } from '../base/gdiWindow';

export class StarMapWindow extends GDIWindow {
  private player: Player;
  private wsSvc: WsService;
  private needInitialFetch: boolean = true;
  private camera: Camera;

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();
    this.camera = new Camera(this.getWidth(), this.getHeight(), 3);

    // hook events
    super.setOnShow(() => {
      // fetch map if needed
      if (this.needInitialFetch) {
        const msg: ClientViewStarMap = {
          sid: this.wsSvc.sid
        };

        this.wsSvc.sendMessage(MessageTypes.ViewStarMap, msg);
        this.needInitialFetch = false;
      }
    });

    super.setOnPreHandleRender((ctx) => {
      // check if map exists
      if (this.player.currentStarMap) {
        // with current star map
        const map = this.player.currentStarMap;

        // center map camera on player's current system
        const currentSystemEntry = map.findRawSystemByID(this.player.currentSystem.id);
        this.camera.x = currentSystemEntry.x;
        this.camera.y = currentSystemEntry.y;

        // draw edges first
        ctx.fillStyle = GDIStyle.starMapEdgeColor;

        for (let row of map.flattened) {
          for (let edge of row.edges) {
            // get endpoint
            const b = edge[1];

            // project endpoints
            const ax = this.camera.projectX(row.system.x);
            const ay = this.camera.projectY(row.system.y);

            const bx = this.camera.projectX(b.x);
            const by = this.camera.projectY(b.y);

            // draw line
            ctx.beginPath();
            ctx.moveTo(ax, ay);
            ctx.lineTo(bx, by);
            ctx.lineWidth = GDIStyle.starMapEdgeWidth;
            ctx.stroke();
          }
        }
      }
    });
  }

  pack() {
    this.setTitle('Star Map');
  }

  periodicUpdate() {
    //
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  setPlayer(player: Player) {
    this.player = player;
  }
}
