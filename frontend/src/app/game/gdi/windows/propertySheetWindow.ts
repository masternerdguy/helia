import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { ClientViewProperty } from '../../wsModels/bodies/viewProperty';
import { MessageTypes } from '../../wsModels/gameMessage';
import { WsService } from '../../ws.service';
import {
  ServerPropertyShipCacheEntry,
  ServerPropertyUpdate,
} from '../../wsModels/bodies/propertyUpdate';
import { Player } from '../../engineModels/player';
import { ClientBoardBody } from '../../wsModels/bodies/board';

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
    this.setWidth(815);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Property Sheet');

    // property list
    this.propertyList.setWidth(715);
    this.propertyList.setHeight(400);
    this.propertyList.initialize();

    this.propertyList.setX(0);
    this.propertyList.setY(0);

    this.propertyList.setFont(FontSize.normal);
    this.propertyList.setOnClick((item) => {
      const row = item as PropertySheetViewRow;
      const ship = row.ship;

      const actions: PropertySheetActionRow[] = [];

      // actions only possible when player is docked
      if (this.player.currentShip.dockedAtStationID) {
        // actions only possible when player has selected a different ship than the one they are flying
        if (this.player.currentShip.id != ship.id) {
          // actions only possible if both ships are docked at the same station
          if (this.player.currentShip.dockedAtStationID == ship.dockedAtId) {
            actions.push({
              listString: () => 'Board',
              ship: ship,
            });
          }
        }
      }

      this.actionList.setItems(actions);
    });

    this.addComponent(this.propertyList);

    // action list
    this.actionList.setWidth(100);
    this.actionList.setHeight(400);
    this.actionList.initialize();

    this.actionList.setX(715);
    this.actionList.setY(0);

    this.actionList.setFont(FontSize.normal);
    this.actionList.setOnClick((item) => {
      const action = item as PropertySheetActionRow;

      if (action.listString() == 'Board') {
        // issue request to board owned ship
        const b = new ClientBoardBody();
        b.sid = this.wsSvc.sid;
        b.shipId = action.ship.id;

        this.wsSvc.sendMessage(MessageTypes.Board, b);
      }
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

    // sort cache by system name, then station name, then ship name
    const sorted = cache.ships.sort((a, b) =>
      this.getShipSortKey(a).localeCompare(this.getShipSortKey(b))
    );

    // build ship entries
    for (const s of sorted) {
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

  private getShipSortKey(a: ServerPropertyShipCacheEntry): string {
    return `${a.systemName}::${a.dockedAtName}::${a.name}::${a.id}`;
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

class PropertySheetActionRow {
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
  let flyingMarker = '';

  if (s.id == p.currentShip.id) {
    flyingMarker = '=>';
  }

  // build string
  return (
    `${fixedString(flyingMarker, 2)} ` +
    `${fixedString(s.name, 32)} ` +
    `${fixedString(s.texture, 12)} ` +
    `${fixedString(s.systemName, 12)} ` +
    `${fixedString(s.dockedAtName, 24)} ` +
    `${fixedString(shortWallet(s.wallet), 11)}`
  );
}

function fixedString(str: string, width: number): string {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}

function shortWallet(d: number): string {
  let o = `${d}`;

  // include metric prefix if needed
  if (d >= 1000000000000000) {
    o = `${(d / 1000000000000000).toFixed(2)}P`;
  } else if (d >= 1000000000000) {
    o = `${(d / 1000000000000).toFixed(2)}T`;
  } else if (d >= 1000000000) {
    o = `${(d / 1000000000).toFixed(2)}G`;
  } else if (d >= 1000000) {
    o = `${(d / 1000000).toFixed(2)}M`;
  } else if (d >= 1000) {
    o = `${(d / 1000).toFixed(2)}k`;
  }

  return o + ' CBN';
}
