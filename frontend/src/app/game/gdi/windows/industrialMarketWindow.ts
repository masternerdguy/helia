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
import { ClientViewIndustrialOrders } from '../../wsModels/bodies/viewIndustrialOrders';
import {
  WSIndustrialOrdersUpdate,
  WSIndustrialSilo,
} from '../../wsModels/entities/wsIndustrialOrdersUpdate';
import { ClientBuyFromSilo } from '../../wsModels/bodies/buyFromSilo';
import { ClientSellToSilo } from '../../wsModels/bodies/sellToSilo';

export class IndustrialMarketWindow extends GDIWindow {
  // lists
  private cargoView: GDIList = new GDIList();
  private actionView: GDIList = new GDIList();
  private orderView: GDIList = new GDIList();

  // order tree depth
  private depthStack: string[] = [];

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

  // orders tree
  private industrialOrdersTree: SilosTree;

  initialize() {
    // set dimensions
    this.setWidth(800);
    this.setHeight(600);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Station Industrial Market');

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
      } else if (a === 'Buy') {
        // verify we are browsing something we can buy from
        const top = this.depthStack[this.depthStack.length - 1];
        const arr = top.split('|');

        if (arr.length === 3) {
          const siloId = arr[0];
          const typeId = arr[1];

          this.modalInput.setOnReturn((txt: string) => {
            // convert text to an integer
            const n = Math.round(Number(txt));

            if (!Number.isNaN(n)) {
              // send buy request
              const tiMsg: ClientBuyFromSilo = {
                sid: this.wsSvc.sid,
                siloId: siloId,
                itemTypeId: typeId,
                quantity: n,
              };

              this.wsSvc.sendMessage(MessageTypes.BuyFromSilo, tiMsg);

              // request cargo bay refresh
              this.refreshCargoBay();

              // request silo refresh
              this.refreshSilos();

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
      } else if (a === 'Sell') {
        // verify we are browsing something we can sell to
        const top = this.depthStack[this.depthStack.length - 1];
        const arr = top.split('|');

        if (arr.length === 3) {
          const siloId = arr[0];

          // get selected item
          const i: ShipViewRow = this.cargoView.getSelectedItem();
          const item = i.object as WSContainerItem;

          this.modalInput.setOnReturn((txt: string) => {
            // convert text to an integer
            const n = Math.round(Number(txt));

            if (!Number.isNaN(n)) {
              // send sell request
              const tiMsg: ClientSellToSilo = {
                sid: this.wsSvc.sid,
                siloId: siloId,
                itemId: item.id,
                quantity: n,
              };

              this.wsSvc.sendMessage(MessageTypes.SellToSilo, tiMsg);

              // request cargo bay refresh
              this.refreshCargoBay();

              // request silo refresh
              this.refreshSilos();

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
      }
    });

    // setup info view
    this.orderView.setWidth(700);
    this.orderView.setHeight(330);
    this.orderView.initialize();

    this.orderView.setX(0);
    this.orderView.setY(0);

    this.orderView.setFont(FontSize.normal);
    this.orderView.setOnClick(() => {
      // get selected target
      const i = this.orderView.getSelectedItem() as OrderViewRow;

      // navigate to target frame
      if (i.next) {
        if (i.next == '--') {
          this.popDepth();
        } else {
          this.pushDepth(i.next);
        }
      }

      if (i.actions) {
        // map action strings for use in view
        const actions = i.actions.map((s: string) =>
          buildOrderViewRowText(s, undefined),
        );

        // list actions on action view
        this.actionView.setItems(actions);
      } else {
        this.actionView.setItems([]);
      }
    });

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
      this.getHeight() / 2 - this.modalInput.getHeight() / 2,
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
      this.refreshSilos();
    });

    // request periodic updates when docked and shown
    setInterval(() => {
      if (this.isDocked && !this.isHidden()) {
        this.refreshCargoBay();
        this.refreshSilos();
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

  private refreshSilos() {
    setTimeout(() => {
      const b = new ClientViewIndustrialOrders();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewIndustrialOrders, b);
    }, 200);
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

      // reset order view and tree
      this.industrialOrdersTree = undefined;
      this.orderView.setSelectedIndex(-1);
      this.orderView.setItems([]);
      this.depthStack = [];

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
      for (const ci of this.player.currentCargoView.items.sort((a, b) =>
        `${a.itemTypeName}::${a.quantity}::${a.id}` >
        `${b.itemTypeName}::${b.quantity}::${b.id}`
          ? 1
          : -1,
      )) {
        const r = buildCargoRowFromContainerItem(
          ci,
          this.isDocked,
          this.depthStack,
        );
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
    const s = this.cargoView.getScroll();

    this.cargoView.setItems(rows);
    this.cargoView.setSelectedIndex(i);
    this.cargoView.setScroll(s);

    if (this.industrialOrdersTree && this.depthStack) {
      try {
        // show orders tree at current depth
        const oRows: OrderViewRow[] = [];
        const idx = this.orderView.getSelectedIndex();

        const depth = this.depthStack.length;

        if (depth === 0) {
          for (const f of this.industrialOrdersTree.families) {
            // add row to browse a specific family
            oRows.push(buildOrderViewRowText(f[1].name, f[0]));
          }
        } else if (depth === 1) {
          // add back button
          oRows.push(buildOrderViewRowText('<== Back to Item Families', '--'));

          // add spacer
          oRows.push(buildOrderViewRowSpacer());

          // add row to browse a specific item type
          for (const g of this.industrialOrdersTree.families.get(
            this.depthStack[0],
          ).groups) {
            oRows.push(buildOrderViewRowText(g[1].name, g[0]));
          }
        } else if (depth === 2) {
          // add back button
          oRows.push(buildOrderViewRowText('<== Back to Item Types', '--'));

          // add spacer
          oRows.push(buildOrderViewRowSpacer());

          // add row to browse a specific order
          for (const g of this.industrialOrdersTree.families
            .get(this.depthStack[0])
            .groups.get(this.depthStack[1]).orders) {
            oRows.push(buildOrderViewDetailRow(g[1], this.depthStack));
          }
        } else if (depth === 3) {
          // add back button
          oRows.push(buildOrderViewRowText('<== Back to Silos', '--'));

          // add spacer
          oRows.push(buildOrderViewRowSpacer());

          // get order
          const order = this.industrialOrdersTree.families
            .get(this.depthStack[0])
            .groups.get(this.depthStack[1])
            .orders.get(this.depthStack[2]);

          // calculate unit volume
          let volume = 0;

          // make a shim item
          const shimItem = makeShimItem(order);

          // get item volume
          volume = Number(shimItem.itemTypeMeta['volume']);

          // NaN check
          if (Number.isNaN(volume)) {
            volume = 0;
          }

          // basic info
          oRows.push(buildOrderViewRowText('Basic Info', undefined));
          oRows.push(
            buildOrderViewRowText(
              infoKeyValueString('Item Type', shimItem.itemTypeName),
              undefined,
            ),
          );
          oRows.push(
            buildOrderViewRowText(
              infoKeyValueString('Item Family', shimItem.itemFamilyName),
              undefined,
            ),
          );

          // add spacer
          oRows.push(buildOrderViewRowSpacer());

          // order details
          oRows.push(buildOrderViewRowText('Order Details', undefined));
          oRows.push(
            buildOrderViewRowText(
              infoKeyValueString(
                order.isSelling ? 'Ask Price' : 'Bid Price',
                `${order.price} CBN`,
              ),
              undefined,
            ),
          );
          oRows.push(
            buildOrderViewRowText(
              infoKeyValueString('Unit Price', `${order.price} CBN`),
              undefined,
            ),
          );
          oRows.push(
            buildOrderViewRowText(
              infoKeyValueString('Unit Volume', `${volume}`),
              undefined,
            ),
          );

          // add spacer
          oRows.push(buildOrderViewRowSpacer());

          // metadata
          oRows.push(buildOrderViewRowText('Metadata', undefined));

          const meta = {};

          Object.assign(meta, shimItem.itemTypeMeta);

          for (const key in meta) {
            if (Object.prototype.hasOwnProperty.call(meta, key)) {
              // exclude industrial market which is handled separately
              if (key !== 'industrialmarket') {
                const value = meta[key];
                oRows.push(
                  buildOrderViewRowText(
                    infoKeyValueString(key, `${value}`),
                    undefined,
                  ),
                );
              }
            }
          }

          // add spacer
          oRows.push(buildOrderViewRowSpacer());

          // industrial market
          const industrialMeta = meta['industrialmarket'];
          if (industrialMeta) {
            oRows.push(buildOrderViewRowText('Industrial Limits', undefined));

            for (const key in industrialMeta) {
              if (Object.prototype.hasOwnProperty.call(industrialMeta, key)) {
                const value = industrialMeta[key];
                const addCBN = key === 'maxprice' || key === 'minprice';

                oRows.push(
                  buildOrderViewRowText(
                    infoKeyValueString(key, `${value}${addCBN ? ' CBN' : ''}`),
                    undefined,
                  ),
                );
              }
            }

            // add spacer
            oRows.push(buildOrderViewRowSpacer());
          }
        }

        // push rows to orders view
        this.orderView.setItems(oRows);
        this.orderView.setSelectedIndex(idx);
      } catch (ex) {
        // log error
        console.error(ex);

        // reset stack and view
        this.depthStack = [];
        this.orderView.setItems([]);
      }
    }
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  setPlayer(player: Player) {
    this.player = player;
  }

  syncIndustrialOrders(orders: WSIndustrialOrdersUpdate) {
    const rawTree: any = {};
    const allOrders: WSIndustrialSilo[] = [];

    // merge order lists
    for (const o of orders.inSilos) {
      allOrders.push(o);
    }

    for (const o of orders.outSilos) {
      allOrders.push(o);
    }

    // build order tree
    for (const o of allOrders) {
      // check if family needs to be created in tree
      if (!rawTree[o.itemFamilyID]) {
        // create branch
        rawTree[o.itemFamilyID] = {
          types: {},
          name: o.itemFamilyName,
        };
      }

      // check if type needs to be created in tree
      if (!rawTree[o.itemFamilyID].types[o.itemTypeID]) {
        // create branch
        rawTree[o.itemFamilyID].types[o.itemTypeID] = {
          orders: {},
          name: o.itemTypeName,
        };
      }

      // store order
      rawTree[o.itemFamilyID].types[o.itemTypeID].orders[
        `${o.stationProcessId}|${o.itemTypeID}|${o.id}`
      ] = o;
    }

    // copy to safer structure
    const safeTree = new SilosTree();

    for (const f in rawTree) {
      if (Object.prototype.hasOwnProperty.call(rawTree, f)) {
        const fe = rawTree[f];

        const family = new SilosFamily();
        family.name = fe.name;

        for (const g in fe.types) {
          if (Object.prototype.hasOwnProperty.call(fe.types, g)) {
            const te = fe.types[g];

            const group = new SilosGroup();
            group.name = te.name;

            for (const d in te.orders) {
              if (Object.prototype.hasOwnProperty.call(te.orders, d)) {
                const de = te.orders[d] as WSIndustrialSilo;

                // add to group
                group.orders.set(d, de);
              }
            }

            // add to family
            family.groups.set(g, group);
          }
        }

        // add to trunk
        safeTree.families.set(f, family);
      }
    }

    // sort trunk alphabetically
    const trunkArr = Array.from(safeTree.families).sort((a, b) =>
      a[1].name > b[1].name ? 1 : -1,
    );

    safeTree.families.clear();

    for (const f of trunkArr) {
      // sort groups alphabetically
      const groupArr = Array.from(f[1].groups).sort((a, b) =>
        a[1].name > b[1].name ? 1 : -1,
      );

      f[1].groups.clear();

      for (const g of groupArr) {
        // sort orders by their price per unit
        const orderArr = Array.from(g[1].orders).sort((a, b) =>
          a[1].price / a[1].available > b[1].price / b[1].available ? 1 : -1,
        );

        g[1].orders.clear();

        for (const od of orderArr) {
          g[1].orders.set(od[0], od[1]);
        }

        // push sorted group to family
        f[1].groups.set(g[0], g[1]);
      }

      // push sorted family to trunk
      safeTree.families.set(f[0], f[1]);
    }

    // store
    this.industrialOrdersTree = safeTree;
  }

  resetViews() {
    // reset cargo view
    this.cargoView.setSelectedIndex(-1);
    this.cargoView.setItems([]);

    // reset orders view
    this.orderView.setSelectedIndex(-1);
    this.orderView.setItems([]);

    // reset actions view
    this.actionView.setSelectedIndex(-1);
    this.actionView.setItems([]);
  }

  private pushDepth(id: string) {
    if (this.industrialOrdersTree) {
      this.depthStack.push(id);
      this.orderView.setItems([]);
    }
  }

  private popDepth() {
    if (this.industrialOrdersTree) {
      this.depthStack.pop();
      this.orderView.setItems([]);
    }
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
  isDocked: boolean,
  depthStack: string[],
): ShipViewRow {
  const r: ShipViewRow = {
    object: m,
    actions: getCargoRowActions(m, isDocked, depthStack),
    listString: () => {
      return itemStatusString(m);
    },
  };

  return r;
}

function getCargoRowActions(
  m: WSContainerItem,
  isDocked: boolean,
  depthStack: string[],
) {
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

    if (m.isPackaged) {
      if (depthStack.length > 0) {
        // verify we are browsing something we can sell into
        const top = depthStack[depthStack.length - 1];
        const arr = top.split('|');

        if (arr.length === 3) {
          const typeId = arr[1];

          if (typeId === m.itemTypeID) {
            // spacer and sell
            actions.push('');
            actions.push('Sell');
          }
        }
      }
    }

    // spacer
    actions.push('');

    // danger zone
    actions.push('Trash');
  }

  return actions;
}

class OrderViewRow {
  object: any;
  next: string;
  actions: string[];

  listString: () => string;
}

function buildOrderViewRowText(s: string, next: string): OrderViewRow {
  const r: OrderViewRow = {
    object: null,
    actions: [],
    next: next,
    listString: () => {
      return s;
    },
  };

  return r;
}

function buildOrderViewRowSpacer(): OrderViewRow {
  const r: OrderViewRow = {
    object: null,
    actions: [],
    next: undefined,
    listString: () => {
      return '';
    },
  };

  return r;
}

function makeShimItem(silo: WSIndustrialSilo): WSContainerItem {
  const i = new WSContainerItem();
  i.id = '';
  i.itemTypeID = silo.itemTypeID;
  i.itemTypeName = silo.itemTypeName;
  i.itemFamilyID = silo.itemFamilyID;
  i.itemFamilyName = silo.itemFamilyName;
  i.quantity = silo.available;
  i.isPackaged = true;
  i.meta = silo.meta;
  i.itemTypeMeta = silo.itemTypeMeta;

  return i;
}

function buildOrderViewDetailRow(
  order: WSIndustrialSilo,
  depthStack: string[],
): OrderViewRow {
  // make a shim item to reuse the cargo info function
  const shimItem = makeShimItem(order);

  // build cargo info
  let cargoString = buildCargoRowFromContainerItem(
    shimItem,
    true,
    depthStack,
  ).listString();

  // calculate volume
  let volume = order.available * Number(order.itemTypeMeta['volume']);

  // NaN check
  if (Number.isNaN(volume)) {
    volume = 0;
  }

  // build actions
  const actions: string[] = [];

  if (order.isSelling) {
    actions.push('Buy');
  }

  // augment cargo string
  cargoString = `${order.isSelling ? '▼' : '▲'}${cargoString.substring(1)}`;

  // build row
  const r: OrderViewRow = {
    object: order,
    actions: actions,
    next: `${order.stationProcessId}|${order.itemTypeID}|${order.id}`,
    listString: () => {
      return `${cargoString} ${fixedString(
        order.price.toString() + ' CBN',
        14,
      )}`;
    },
  };

  return r;
}

function infoKeyValueString(key: string, value: string) {
  // build string
  return `${fixedString('', 1)} ${fixedString(key, 32)} ${fixedString(
    value,
    32,
  )}`;
}

function fixedString(str: string, width: number): string {
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
    40,
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

class SilosTree {
  families: Map<string, SilosFamily> = new Map<string, SilosFamily>();
}

class SilosFamily {
  name: string;
  groups: Map<string, SilosGroup> = new Map<string, SilosGroup>();
}

class SilosGroup {
  name: string;
  orders: Map<string, WSIndustrialSilo> = new Map<string, WSIndustrialSilo>();
}
