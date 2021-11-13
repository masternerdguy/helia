import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { ClientViewProperty } from '../../wsModels/bodies/viewProperty';
import { MessageTypes } from '../../wsModels/gameMessage';
import { WsService } from '../../ws.service';

export class PropertySheetWindow extends GDIWindow {
  private propertyList = new GDIList();
  private actionList = new GDIList();

  // last cargo bay refresh
  private lastPropertyView = 0;

  // ws service
  private wsSvc: WsService;

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Property Sheet');

    // property list
    this.propertyList.setWidth(300);
    this.propertyList.setHeight(400);
    this.propertyList.initialize();

    this.propertyList.setX(0);
    this.propertyList.setY(0);

    this.propertyList.setFont(FontSize.normal);
    this.propertyList.setOnClick((item) => {
      // todo
    });

    this.addComponent(this.propertyList);

    // action list
    this.actionList.setWidth(100);
    this.actionList.setHeight(400);
    this.actionList.initialize();

    this.actionList.setX(300);
    this.actionList.setY(0);

    this.actionList.setFont(FontSize.normal);
    this.actionList.setOnClick((item) => {
      // todo
    });

    this.addComponent(this.actionList);
  }

  periodicUpdate() {
    if (!this.isHidden()) {
      // check if property view is stale
      const now = Date.now();

      if (now - this.lastPropertyView > 5000) {
        // refresh summary
        this.refreshPropertySummary();
        this.lastPropertyView = now;
      }
    }
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  private refreshPropertySummary() {
    setTimeout(() => {
      const b = new ClientViewProperty();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewProperty, b);
    }, 200);
  }
}

function fixedString(str: string, width: number) {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}
