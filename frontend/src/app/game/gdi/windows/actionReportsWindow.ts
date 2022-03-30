import { Player } from '../../engineModels/player';
import { WsService } from '../../ws.service';
import { GDIWindow } from '../base/gdiWindow';

export class ActionReportsWindow extends GDIWindow {
  private wsSvc: WsService;
  private player: Player;

  initialize() {
    // set dimensions
    this.setWidth(800);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Action Reports');
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
