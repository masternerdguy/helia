import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
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
import { GDIOverlay } from '../components/gdiOverlay';
import { GDIInput } from '../components/gdiInput';
import { ClientTransferCreditsBody } from '../../wsModels/bodies/transferCredits';
import { ClientSellShipAsOrderBody } from '../../wsModels/bodies/sellShipAsOrder';
import { ClientTrashShipBody } from '../../wsModels/bodies/trashShip';
import { ClientRenameShipBody } from '../../wsModels/bodies/renameShip';

export class PropertySheetWindow extends GDIWindow {
  private propertyList = new GDIList();
  private actionList = new GDIList();

  // inputs
  private modalOverlay: GDIOverlay = new GDIOverlay();
  private modalInput: GDIInput = new GDIInput();

  // last property refresh
  private lastPropertyView: number = 0;

  // dock state
  private isDocked: boolean = undefined;

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
        // actions only possible if both ships are also docked at the same station
        if (this.player.currentShip.dockedAtStationID == ship.dockedAtId) {
          actions.push({
            listString: () => 'Rename',
            ship: ship,
          });

          // spacer
          actions.push({
            listString: () => '',
            ship: null,
          });

          // actions only possible when player has also selected a different ship than the one they are flying
          if (this.player.currentShip.id != ship.id) {
            actions.push({
              listString: () => 'Board',
              ship: ship,
            });

            actions.push({
              listString: () => 'Move CBN',
              ship: ship,
            });

            // spacer
            actions.push({
              listString: () => '',
              ship: null,
            });

            actions.push({
              listString: () => 'Sell',
              ship: ship,
            });

            actions.push({
              listString: () => 'Trash',
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

        // request refresh
        this.refreshPropertySummary();

        // reset views
        this.resetViews();
      } else if (action.listString() == 'Move CBN') {
        this.modalInput.setOnReturn((txt: string) => {
          // convert text to an integer
          const n = Math.round(Number(txt));

          if (!Number.isNaN(n)) {
            // send credit transfer request
            const tiMsg: ClientTransferCreditsBody = {
              sid: this.wsSvc.sid,
              shipId: action.ship.id,
              amount: n,
            };

            this.wsSvc.sendMessage(MessageTypes.TransferCredits, tiMsg);

            // request refresh
            this.refreshPropertySummary();

            // reset views
            this.resetViews();
          }

          // hide modal overlay
          this.hideModalInput();
        });

        this.showModalInput();
      } else if (action.listString() == 'Sell') {
        this.modalInput.setOnReturn((txt: string) => {
          // convert text to an integer
          const n = Math.round(Number(txt));

          if (!Number.isNaN(n)) {
            // send sell as order request
            const tiMsg: ClientSellShipAsOrderBody = {
              sid: this.wsSvc.sid,
              shipId: action.ship.id,
              price: n,
            };

            this.wsSvc.sendMessage(MessageTypes.SellShipAsOrder, tiMsg);

            // request refresh
            this.refreshPropertySummary();

            // reset views
            this.resetViews();
          }

          // hide modal overlay
          this.hideModalInput();
        });

        this.showModalInput();
      } else if (action.listString() == 'Trash') {
        // issue request to trash owned ship
        const b = new ClientTrashShipBody();
        b.sid = this.wsSvc.sid;
        b.shipId = action.ship.id;

        this.wsSvc.sendMessage(MessageTypes.TrashShip, b);

        // request refresh
        this.refreshPropertySummary();

        // reset views
        this.resetViews();
      } else if (action.listString() == 'Rename') {
        this.modalInput.setOnReturn((txt: string) => {
          // send rename request
          const tiMsg: ClientRenameShipBody = {
            sid: this.wsSvc.sid,
            shipId: action.ship.id,
            name: txt,
          };

          this.wsSvc.sendMessage(MessageTypes.RenameShip, tiMsg);

          // request refresh
          this.refreshPropertySummary();

          // reset views
          this.resetViews();

          // hide modal overlay
          this.hideModalInput();
        });

        this.showModalInput();
      }
    });

    this.addComponent(this.actionList);

    // setup modal input
    this.modalOverlay.setWidth(this.getWidth());
    this.modalOverlay.setHeight(this.getHeight());
    this.modalOverlay.setX(0);
    this.modalOverlay.setY(0);
    this.modalOverlay.initialize();

    const fontSize = GDIStyle.getUnderlyingFontSize(FontSize.large);
    this.modalInput.setWidth(200);
    this.modalInput.setHeight(Math.round(fontSize + 0.5));
    this.modalInput.setX(this.getWidth() / 2 - this.modalInput.getWidth() / 2);
    this.modalInput.setY(
      this.getHeight() / 2 - this.modalInput.getHeight() / 2
    );
    this.modalInput.setFont(FontSize.large);
    this.modalInput.initialize();
  }

  private showModalInput() {
    this.removeComponent(this.propertyList);
    this.removeComponent(this.actionList);
    this.addComponent(this.modalOverlay);
    this.addComponent(this.modalInput);
  }

  private hideModalInput() {
    this.addComponent(this.propertyList);
    this.addComponent(this.actionList);
    this.removeComponent(this.modalOverlay);
    this.removeComponent(this.modalInput);
  }

  periodicUpdate() {
    // check for dock state change
    let isDocked: boolean = undefined;
    this.player.currentShip.dockedAtStationID ? (isDocked = true) : false;

    if (isDocked != this.isDocked) {
      // reset views
      this.resetViews();
    }

    // store current dock state
    this.isDocked = isDocked;

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

  private resetViews() {
    this.actionList.setItems([]);
    this.propertyList.setItems([]);

    // clear input
    this.modalInput.setText('');
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
