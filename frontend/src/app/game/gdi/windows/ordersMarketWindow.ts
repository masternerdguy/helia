import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { WsService } from '../../ws.service';
import { Player } from '../../engineModels/player';
import { ClientViewCargoBay } from '../../wsModels/bodies/viewCargoBay';
import { MessageTypes } from '../../wsModels/gameMessage';
import { WSContainerItem } from '../../wsModels/entities/wsContainer';
import { GDIInput } from '../components/gdiInput';
import { GDIOverlay } from '../components/gdiOverlay';
import { GDIBar } from '../components/gdiBar';
import { ClientPackageItem } from '../../wsModels/bodies/packageItem';
import { ClientSplitItem } from '../../wsModels/bodies/splitItem';
import { ClientStackItem } from '../../wsModels/bodies/stackItem';
import { ClientTrashItem } from '../../wsModels/bodies/trashItem';
import { ClientUnpackageItem } from '../../wsModels/bodies/unpackageItem';
import { ClientSellAsOrder } from '../../wsModels/bodies/sellAsOrder';
import { ClientViewOpenSellOrders } from '../../wsModels/bodies/viewOpenSellOrders';
import { WSOpenSellOrder, WSOpenSellOrdersUpdate } from '../../wsModels/entities/wsOpenSellOrdersUpdate';

export class OrdersMarketWindow extends GDIWindow {
  // lists
  private cargoView: GDIList = new GDIList();
  private actionView: GDIList = new GDIList();
  private orderView: GDIList = new GDIList();

  // inputs
  private modalOverlay: GDIOverlay = new GDIOverlay();
  private modalInput: GDIInput = new GDIInput();

  // bars
  private cargoBayUsed: GDIBar = new GDIBar();

  // player
  private player: Player;

  // ws service
  private wsSvc: WsService;

  // last cargo bay refresh
  private lastCargoView = 0;
  private isDocked = false;

  initialize() {
    // set dimensions
    this.setWidth(800);
    this.setHeight(600);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Station Orders Market');

    // setup cargo view
    this.cargoView.setWidth(700);
    this.cargoView.setHeight(250);
    this.cargoView.initialize();

    this.cargoView.setX(0);
    this.cargoView.setY(this.getHeight() - this.cargoView.getHeight());

    this.cargoView.setFont(FontSize.normal);
    this.cargoView.setOnClick((r) => {
      // check for actions
      if (r.actions) {
        // map action strings for use in view
        const actions = r.actions.map((s: string) => buildShipViewRowText(s));

        // list actions on action view
        this.actionView.setItems(actions);
      }
    });

    // setup cargo bar
    this.cargoBayUsed.setX(0);
    this.cargoBayUsed.setY(this.cargoView.getY() - 20);
    this.cargoBayUsed.setWidth(700);
    this.cargoBayUsed.setHeight(20);
    this.cargoBayUsed.initialize();
    this.cargoBayUsed.setFont(FontSize.normal);
    this.cargoBayUsed.setText('Cargo Bay Used');

    // setup action view
    this.actionView.setWidth(100);
    this.actionView.setHeight(600);
    this.actionView.initialize();

    this.actionView.setX(700);
    this.actionView.setY(0);

    this.actionView.setFont(FontSize.normal);
    this.actionView.setOnClick((r) => {
      // get action
      const a = r.listString();

      if (a === 'Trash') {
        // get selected item
        const i: ShipViewRow = this.cargoView.getSelectedItem();

        // send trash request
        const tiMsg: ClientTrashItem = {
          sid: this.wsSvc.sid,
          itemID: (i.object as WSContainerItem).id,
        };

        this.wsSvc.sendMessage(MessageTypes.TrashItem, tiMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Package') {
        // get selected item
        const i: ShipViewRow = this.cargoView.getSelectedItem();

        // send package request
        const tiMsg: ClientPackageItem = {
          sid: this.wsSvc.sid,
          itemID: (i.object as WSContainerItem).id,
        };

        this.wsSvc.sendMessage(MessageTypes.PackageItem, tiMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Unpackage') {
        // get selected item
        const i: ShipViewRow = this.cargoView.getSelectedItem();

        // send unpackage request
        const tiMsg: ClientUnpackageItem = {
          sid: this.wsSvc.sid,
          itemID: (i.object as WSContainerItem).id,
        };

        this.wsSvc.sendMessage(MessageTypes.UnpackageItem, tiMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Stack') {
        // get selected item
        const i: ShipViewRow = this.cargoView.getSelectedItem();

        // send stack request
        const tiMsg: ClientStackItem = {
          sid: this.wsSvc.sid,
          itemID: (i.object as WSContainerItem).id,
        };

        this.wsSvc.sendMessage(MessageTypes.StackItem, tiMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Split') {
        // get selected item
        const i: ShipViewRow = this.cargoView.getSelectedItem();

        this.modalInput.setOnReturn((txt: string) => {
          // convert text to an integer
          const n = Math.round(Number(txt));

          if (!Number.isNaN(n)) {
            // send split request
            const tiMsg: ClientSplitItem = {
              sid: this.wsSvc.sid,
              itemID: (i.object as WSContainerItem).id,
              size: n,
            };

            this.wsSvc.sendMessage(MessageTypes.SplitItem, tiMsg);

            // request cargo bay refresh
            this.refreshCargoBay();

            // reset views
            this.resetViews();
          }

          // clear input
          this.modalInput.setText('');

          // hide modal overlay
          this.hideModalInput();
        });

        this.showModalInput();
      } else if (a === 'Sell') {
        // get selected item
        const i: ShipViewRow = this.cargoView.getSelectedItem();

        this.modalInput.setOnReturn((txt: string) => {
          // convert text to an integer
          const n = Math.round(Number(txt));

          if (!Number.isNaN(n)) {
            // send sell request
            const tiMsg: ClientSellAsOrder = {
              sid: this.wsSvc.sid,
              itemID: (i.object as WSContainerItem).id,
              price: n,
            };

            this.wsSvc.sendMessage(MessageTypes.SellAsOrder, tiMsg);

            // request cargo bay refresh
            this.refreshCargoBay();

            // todo: request station orders refresh

            // reset views
            this.resetViews();
          }

          // clear input
          this.modalInput.setText('');

          // hide modal overlay
          this.hideModalInput();
        });

        this.showModalInput();
      }
    });

    // setup info view
    this.orderView.setWidth(700);
    this.orderView.setHeight(330);
    this.orderView.initialize();

    this.orderView.setX(0);
    this.orderView.setY(0);

    this.orderView.setFont(FontSize.normal);
    this.orderView.setOnClick(() => {});

    // setup modal input
    this.modalOverlay.setWidth(this.getWidth());
    this.modalOverlay.setHeight(this.getHeight());
    this.modalOverlay.setX(0);
    this.modalOverlay.setY(0);
    this.modalOverlay.initialize();

    const fontSize = GDIStyle.getUnderlyingFontSize(FontSize.large);
    this.modalInput.setWidth(100);
    this.modalInput.setHeight(Math.round(fontSize + 0.5));
    this.modalInput.setX(this.getWidth() / 2 - this.modalInput.getWidth() / 2);
    this.modalInput.setY(
      this.getHeight() / 2 - this.modalInput.getHeight() / 2
    );
    this.modalInput.setFont(FontSize.large);
    this.modalInput.initialize();

    // pack
    this.addComponent(this.cargoView);
    this.addComponent(this.actionView);
    this.addComponent(this.cargoBayUsed);
    this.addComponent(this.orderView);

    // request updates on show
    this.setOnShow(() => {
      this.refreshCargoBay();
      this.refreshOpenSellOrders();
    });

    // request periodic updates when docked and shown
    setInterval(() => {
      if (this.isDocked && !this.isHidden()) {
        this.refreshCargoBay();
        this.refreshOpenSellOrders();
      }
    }, 30000);
  }

  private showModalInput() {
    this.removeComponent(this.cargoView);
    this.removeComponent(this.actionView);
    this.removeComponent(this.cargoBayUsed);
    this.removeComponent(this.orderView);
    this.addComponent(this.modalOverlay);
    this.addComponent(this.modalInput);
  }

  private hideModalInput() {
    this.addComponent(this.cargoView);
    this.addComponent(this.actionView);
    this.addComponent(this.cargoBayUsed);
    this.addComponent(this.orderView);
    this.removeComponent(this.modalOverlay);
    this.removeComponent(this.modalInput);
  }

  private refreshCargoBay() {
    setTimeout(() => {
      const b = new ClientViewCargoBay();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewCargoBay, b);
    }, 200);
  }

  private refreshOpenSellOrders() {
    setTimeout(() => {
      const b = new ClientViewOpenSellOrders();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewOpenSellOrders, b);
    }, 200);
  }

  syncSellOrders(orders: WSOpenSellOrdersUpdate) {
    const rawTree: any = {};

    for (const o of orders.orders) {
      // check if family needs to be created in tree
      if(!rawTree[o.item.itemFamilyID]) {
        // create branch
        rawTree[o.item.itemFamilyID] = {
          types: {},
          name: o.item.itemFamilyName
        };
      }

      // check if type needs to be created in tree
      if(!rawTree[o.item.itemFamilyID].types[o.item.itemTypeID]) {
        // create branch
        rawTree[o.item.itemFamilyID].types[o.item.itemTypeID] = {
          orders: {},
          name: o.item.itemTypeName
        };
      }

      // store order
      rawTree[o.item.itemFamilyID].types[o.item.itemTypeID].orders[o.id] = o;

      // debug out
      console.log(rawTree);
    }
  }

  periodicUpdate() {
    if (this.isHidden()) {
      return;
    }

    // make sure resources are available
    if (!this.player?.currentShip) {
      return;
    }

    if (!this.player?.currentShip.fitStatus) {
      return;
    }

    if (!this.wsSvc) {
      return;
    }

    // check for docked status change
    const docked: boolean = !!this.player.currentShip.dockedAtStationID;

    if (docked !== this.isDocked) {
      // reset cargo view
      this.cargoView.setSelectedIndex(-1);
      this.cargoView.setItems([]);

      // reset action view and store status
      this.actionView.setItems([]);
      this.isDocked = docked;
    }

    // set up view row list
    const rows: ShipViewRow[] = [];

    // check if cargo bay is stale
    const now = Date.now();

    if (now - this.lastCargoView > 5000) {
      // request current cargo bay
      const b = new ClientViewCargoBay();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewCargoBay, b);
      this.lastCargoView = now;
    }

    // update cargo display
    const cargo: ShipViewRow[] = [];

    if (this.player.currentCargoView) {
      for (const ci of this.player.currentCargoView.items) {
        const r = buildCargoRowFromContainerItem(ci, this.isDocked);
        cargo.push(r);
      }
    }

    this.cargoBayUsed.setPercentage(this.player.currentShip.cargoP);

    // list cargo bay items
    for (const r of cargo) {
      rows.push(r);
    }

    rows.push(buildShipViewRowSpacer());

    // layout wallet
    rows.push(buildShipViewRowText('Wallet'));
    rows.push(buildShipViewRowText(`  ${this.player.currentShip.wallet} CBN`));

    rows.push(buildShipViewRowSpacer());

    // push to view
    const i = this.cargoView.getSelectedIndex();

    this.cargoView.setItems(rows);
    this.cargoView.setSelectedIndex(i);
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  setPlayer(player: Player) {
    this.player = player;
  }

  resetViews() {
    // reset cargo view
    this.cargoView.setSelectedIndex(-1);
    this.cargoView.setItems([]);
  }
}

class ShipViewRow {
  object: any;
  actions: string[];

  listString: () => string;
}

function buildShipViewRowSpacer(): ShipViewRow {
  const r: ShipViewRow = {
    object: null,
    actions: [],
    listString: () => {
      return '';
    },
  };

  return r;
}

function buildShipViewRowText(s: string): ShipViewRow {
  const r: ShipViewRow = {
    object: null,
    actions: [],
    listString: () => {
      return s;
    },
  };

  return r;
}

function buildCargoRowFromContainerItem(
  m: WSContainerItem,
  isDocked: boolean
): ShipViewRow {
  const r: ShipViewRow = {
    object: m,
    actions: getCargoRowActions(m, isDocked),
    listString: () => {
      return itemStatusString(m);
    },
  };

  return r;
}

function getCargoRowActions(m: WSContainerItem, isDocked: boolean) {
  const actions: string[] = [];

  if (m.isPackaged) {
    actions.push('Stack');

    if (m.quantity > 1) {
      actions.push('Split');
    }
  }

  if (isDocked) {
    if (m.isPackaged) {
      if (m.quantity === 1) {
        actions.push('Unpackage');
      }
    } else {
      actions.push('Package');
    }

    // spacer and sell
    actions.push('');
    actions.push('Sell');

    // spacer
    actions.push('');

    // danger zone
    actions.push('Trash');
  }

  return actions;
}

function fixedString(str: string, width: number) {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}

function itemStatusString(m: WSContainerItem) {
  // build status string
  const q = cargoQuantity(m.quantity);
  return `${fixedString(m.isPackaged ? '◰' : '', 1)} ${fixedString(
    m.itemTypeName,
    40
  )} ${fixedString(m.itemFamilyName, 16)} ${fixedString(q, 8)}`;
}

function cargoQuantity(d: number): string {
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

  return o;
}
