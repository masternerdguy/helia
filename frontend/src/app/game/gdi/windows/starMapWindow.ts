import { Player } from '../../engineModels/player';
import { WsService } from '../../ws.service';
import { GDIWindow } from '../base/gdiWindow';

export class StarMapWindow extends GDIWindow {
  private player: Player;
  private wsSvc: WsService;

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();
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
