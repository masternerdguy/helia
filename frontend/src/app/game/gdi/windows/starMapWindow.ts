import { Player } from '../../engineModels/player';
import { WsService } from '../../ws.service';
import { ClientViewStarMap } from '../../wsModels/bodies/viewStarMap';
import { MessageTypes } from '../../wsModels/gameMessage';
import { GDIWindow } from '../base/gdiWindow';

export class StarMapWindow extends GDIWindow {
  private player: Player;
  private wsSvc: WsService;
  private needInitialFetch: boolean = true;

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();

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

    super.setOnLastRenderStep((ctx) => {
      // check if map exists
      if (this.player.currentStarMap) {
        // with current star map
        const map = this.player.currentStarMap;

        // debug out
        console.log(map);
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
