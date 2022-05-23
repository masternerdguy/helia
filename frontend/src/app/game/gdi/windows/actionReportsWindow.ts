import { Player } from '../../engineModels/player';
import { WsService } from '../../ws.service';
import { ClientViewActionReportsPage } from '../../wsModels/bodies/viewActionReportsPage';
import { MessageTypes } from '../../wsModels/gameMessage';
import { FontSize } from '../base/gdiStyle';
import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';

export class ActionReportsWindow extends GDIWindow {
  private actionView: GDIList = new GDIList();
  private reportView: GDIList = new GDIList();

  private wsSvc: WsService;
  private player: Player;
  private page: number;
  private pageSize: number;

  initialize() {
    // set dimensions
    this.setWidth(800);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Action Reports');
    this.page = 0;
    this.pageSize = 30;

    // setup action view
    this.actionView.setWidth(100);
    this.actionView.setHeight(400);
    this.actionView.initialize();

    this.actionView.setX(700);
    this.actionView.setY(0);

    this.actionView.setFont(FontSize.normal);
    this.actionView.setOnClick((r) => {
      // get action
      const a = r.listString();
    });

    // setup info view
    this.reportView.setWidth(700);
    this.reportView.setHeight(400);
    this.reportView.initialize();

    this.reportView.setX(0);
    this.reportView.setY(0);

    this.reportView.setFont(FontSize.normal);
    this.reportView.setOnClick(() => {});

    // request updates on show
    this.setOnShow(() => {
      this.refreshPage();
    });

    // request periodic updates when shown
    setInterval(() => {
      if (!this.isHidden()) {
        this.refreshPage();
      }
    }, 5000);

    // pack
    this.addComponent(this.reportView);
    this.addComponent(this.actionView);
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

  private refreshPage() {
    setTimeout(() => {
      const b = new ClientViewActionReportsPage();
      b.sid = this.wsSvc.sid;
      b.page = this.page;
      b.count = this.pageSize;

      this.wsSvc.sendMessage(MessageTypes.ViewActionReportsPage, b);
    }, 200);
  }
}
