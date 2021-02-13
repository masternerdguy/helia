import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { WSModule } from '../../wsModels/entities/wsShip';
import { WsService } from '../../ws.service';
import { Player } from '../../engineModels/player';
import { ClientViewCargoBay } from '../../wsModels/bodies/viewCargoBay';
import { MessageTypes } from '../../wsModels/gameMessage';
import { WSContainerItem } from '../../wsModels/entities/wsContainer';
import { GDIInput } from '../components/gdiInput';
import { GDIOverlay } from '../components/gdiOverlay';
import { GDIBar } from '../components/gdiBar';

export class OrdersMarketWindow extends GDIWindow {
  // lists
  private cargoView: GDIList = new GDIList();
  private actionView: GDIList = new GDIList();

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
      // todo
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

      // todo
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
      this.getHeight() / 2 - this.modalInput.getHeight() / 2
    );
    this.modalInput.setFont(FontSize.large);
    this.modalInput.initialize();

    // pack
    this.addComponent(this.cargoView);
    this.addComponent(this.actionView);
    this.addComponent(this.cargoBayUsed);
  }

  private showModalInput() {
    this.removeComponent(this.cargoView);
    this.removeComponent(this.actionView);
    this.removeComponent(this.cargoBayUsed);
    this.addComponent(this.modalOverlay);
    this.addComponent(this.modalInput);
  }

  private hideModalInput() {
    this.addComponent(this.cargoView);
    this.addComponent(this.actionView);
    this.addComponent(this.cargoBayUsed);
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

function buildInfoRowsFromContainerItem(m: WSContainerItem): ShipViewRow[] {
  const rows: ShipViewRow[] = [];

  // basic info
  rows.push(buildShipViewRowText('Basic Info'));

  const type = buildShipViewRowText(infoKeyValueString('Type', m.itemTypeName));

  const family = buildShipViewRowText(
    infoKeyValueString('Family', m.itemFamilyName)
  );

  const quantity = buildShipViewRowText(
    infoKeyValueString('Quantity', m.quantity.toString())
  );

  const packaged = buildShipViewRowText(
    infoKeyValueString('Packaged', m.isPackaged.toString())
  );

  // store basic info
  rows.push(type);
  rows.push(family);
  rows.push(quantity);
  rows.push(packaged);

  // spacer after basic info
  rows.push(buildShipViewRowSpacer());

  // combine item and item type metadata
  const meta: any = {};

  if (m.itemTypeMeta) {
    Object.assign(meta, m.itemTypeMeta);
  }

  if (m.meta) {
    Object.assign(meta, m.meta);
  }

  // metadata info
  rows.push(buildShipViewRowText('Metadata'));

  for (const prop in meta) {
    if (Object.prototype.hasOwnProperty.call(meta, prop)) {
      const v = buildShipViewRowText(infoKeyValueString(prop, `${meta[prop]}`));

      rows.push(v);
    }
  }

  // spacer after metadata info
  rows.push(buildShipViewRowSpacer());

  // return info rows
  return rows;
}

function buildInfoRowsFromModule(m: WSModule): ShipViewRow[] {
  const rows: ShipViewRow[] = [];

  // basic info
  rows.push(buildShipViewRowText('Basic Info'));

  const type = buildShipViewRowText(infoKeyValueString('Type', m.type));

  const family = buildShipViewRowText(
    infoKeyValueString('Family', m.familyName)
  );

  // store basic info
  rows.push(type);
  rows.push(family);

  // spacer after basic info
  rows.push(buildShipViewRowSpacer());

  // combine item and item type metadata
  const meta: any = {};

  if (m.meta) {
    Object.assign(meta, m.meta);
  }

  // metadata info
  rows.push(buildShipViewRowText('Metadata'));

  for (const prop in meta) {
    if (Object.prototype.hasOwnProperty.call(meta, prop)) {
      const v = buildShipViewRowText(infoKeyValueString(prop, `${meta[prop]}`));

      rows.push(v);
    }
  }

  // spacer after metadata info
  rows.push(buildShipViewRowSpacer());

  // slot info
  rows.push(buildShipViewRowText('Slot Info'));

  const slotFamily = buildShipViewRowText(
    infoKeyValueString('Compatibility', m.hpFamily)
  );
  const slotVolume = buildShipViewRowText(
    infoKeyValueString('Volume', m.hpVolume?.toString())
  );

  rows.push(slotFamily);
  rows.push(slotVolume);

  // spacer after slot info
  rows.push(buildShipViewRowSpacer());

  // return info rows
  return rows;
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

      // determine whether or not this is a module
      const isModule =
        m.itemFamilyID === 'gun_turret' ||
        m.itemFamilyID === 'missile_launcher' ||
        m.itemFamilyID === 'shield_booster' ||
        m.itemFamilyID === 'fuel_tank' ||
        m.itemFamilyID === 'armor_plate';

      if (isModule) {
        // offer fit action
        actions.push('Fit');
      }
    }

    // spacer
    actions.push('');

    // danger zone
    actions.push('Trash');
  }

  return actions;
}

function buildFittingRowFromModule(
  m: WSModule,
  isDocked: boolean
): ShipViewRow {
  const r: ShipViewRow = {
    object: m,
    actions: getFittingRowActions(isDocked, m),
    listString: () => {
      return moduleStatusString(m);
    },
  };

  return r;
}

function getFittingRowActions(isDocked: boolean, m: WSModule) {
  const actions: string[] = [];

  if (isDocked) {
    if (m.itemID !== '00000000-0000-0000-0000-000000000000') {
      actions.push('Unfit');
    }
  }

  return actions;
}

function fixedString(str: string, width: number) {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}

function moduleStatusString(m: WSModule) {
  const pc = fixedString(m.isCycling ? `${m.cyclePercent}%` : '', 4);

  // build status string
  return `${fixedString(m.willRepeat ? '*' : '', 1)} ${fixedString(
    m.itemID !== '00000000-0000-0000-0000-000000000000' ? m.type : '[EMPTY]',
    24
  )} ${pc}`;
}

function itemStatusString(m: WSContainerItem) {
  // build status string
  const q = cargoQuantity(m.quantity);
  return `${fixedString(m.isPackaged ? '◰' : '', 1)} ${fixedString(
    m.itemTypeName,
    40
  )} ${fixedString(m.itemFamilyName, 16)} ${fixedString(q, 8)}`;
}

function infoKeyValueString(key: string, value: string) {
  // build string
  return `${fixedString('', 1)} ${fixedString(key, 32)} ${fixedString(
    value,
    32
  )}`;
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
