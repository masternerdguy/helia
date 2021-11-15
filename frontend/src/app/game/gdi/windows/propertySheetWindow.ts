import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { ClientViewProperty } from '../../wsModels/bodies/viewProperty';
import { MessageTypes } from '../../wsModels/gameMessage';
import { WsService } from '../../ws.service';
import { ServerPropertyShipCacheEntry, ServerPropertyUpdate } from '../../wsModels/bodies/propertyUpdate';
import { Player } from '../../engineModels/player';

export class PropertySheetWindow extends GDIWindow {
  private propertyList = new GDIList();
  private actionList = new GDIList();

  // last cargo bay refresh
  private lastPropertyView = 0;

  // player
  private player: Player;

  // ws service
  private wsSvc: WsService;

  initialize() {
    // set dimensions
    this.setWidth(700);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Property Sheet');

    // property list
    this.propertyList.setWidth(700);
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

    this.actionList.setX(600);
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

  setPlayer(player: Player) {
    this.player = player;
  }

  sync(cache: ServerPropertyUpdate) {
    // stash selected index
    const sIdx = this.propertyList.getSelectedIndex();

    const rows: PropertySheetViewRow[] = [];

    // build ship entries
    for (const s of cache.ships) {
      const r = new PropertySheetViewRow();
      const ls = propertySheetViewRowStringFromShip(s, this.player);

      r.listString = () => ls;
      r.ship = s;

      rows.push(r);
    }
    
    // update list
    this.propertyList.setItems(rows);

    // re-select index
    this.propertyList.setSelectedIndex(sIdx);
  }

  private refreshPropertySummary() {
    setTimeout(() => {
      const b = new ClientViewProperty();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewProperty, b);
    }, 200);
  }
}

class PropertySheetViewRow {
  listString: () => string;
  ship: ServerPropertyShipCacheEntry;
}

function propertySheetViewRowStringFromShip(
    s: ServerPropertyShipCacheEntry,
    p: Player
  ): string {

  if (s == null) {
    return;
  }

  // marker if current ship
  let flyingMarker = "";

  if (s.id == p.currentShip.id) {
    flyingMarker = "=>";
  }

  // build string
  return `${fixedString(flyingMarker, 2)} `
    + `${fixedString(s.name, 32)} `
    + `${fixedString(s.texture, 12)} `
    + `${fixedString(s.systemName, 12)} ` 
    + `${fixedString(s.dockedAtName, 24)}`;
}

function fixedString(str: string, width: number): string {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}
