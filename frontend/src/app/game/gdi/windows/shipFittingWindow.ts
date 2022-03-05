import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { WSModule } from '../../wsModels/entities/wsShip';
import { WsService } from '../../ws.service';
import { Player } from '../../engineModels/player';
import { ClientViewCargoBay } from '../../wsModels/bodies/viewCargoBay';
import { MessageTypes } from '../../wsModels/gameMessage';
import { WSContainerItem } from '../../wsModels/entities/wsContainer';
import { ClientUnfitModule } from '../../wsModels/bodies/unfitModule';
import { ClientTrashItem } from '../../wsModels/bodies/trashItem';
import { ClientPackageItem } from '../../wsModels/bodies/packageItem';
import { ClientUnpackageItem } from '../../wsModels/bodies/unpackageItem';
import { ClientStackItem } from '../../wsModels/bodies/stackItem';
import { GDIInput } from '../components/gdiInput';
import { GDIOverlay } from '../components/gdiOverlay';
import { ClientSplitItem } from '../../wsModels/bodies/splitItem';
import { ClientFitModule } from '../../wsModels/bodies/fitModule';
import { GDIBar } from '../components/gdiBar';
import { ClientConsumeFuel } from '../../wsModels/bodies/consumeFuel';
import { ClientSelfDestruct } from '../../wsModels/bodies/selfDestruct';
import { ClientConsumeRepairKit } from '../../wsModels/bodies/consumeRepairKit';
import {
  ServerPropertyShipCacheEntry,
  ServerPropertyUpdate,
} from '../../wsModels/bodies/propertyUpdate';
import { ClientTransferItem } from '../../wsModels/bodies/transferItem';
import { ClientViewProperty } from '../../wsModels/bodies/viewProperty';
import { ClientRunSchematic } from '../../wsModels/bodies/runSchematic';

export class ShipFittingWindow extends GDIWindow {
  // lists
  private shipView: GDIList = new GDIList();
  private infoView: GDIList = new GDIList();
  private actionView: GDIList = new GDIList();

  // inputs
  private modalOverlay: GDIOverlay = new GDIOverlay();
  private modalInput: GDIInput = new GDIInput();
  private modalPropertyView: GDIList = new GDIList();

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
    this.setWidth(600);
    this.setHeight(600);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Ship Cargo and Fitting');

    // setup ship view
    this.shipView.setWidth(500);
    this.shipView.setHeight(380);
    this.shipView.initialize();

    this.shipView.setX(0);
    this.shipView.setY(0);

    this.shipView.setFont(FontSize.normal);
    this.shipView.setOnClick((r) => {
      // check for actions
      if (r.actions) {
        // map action strings for use in view
        const actions = r.actions.map((s: string) => buildShipViewRowText(s));

        // list actions on action view
        this.actionView.setItems(actions);

        // build and show item info
        if (r.object) {
          if (!r.object.metaRack) {
            // this is an item in the cargo bay
            const info = buildInfoRowsFromContainerItem(r.object);
            this.infoView.setItems(info);
          } else {
            // this is a fitted module
            const info = buildInfoRowsFromModule(r.object);
            this.infoView.setItems(info);
          }
        } else {
          // this is nothing
          this.infoView.setItems([]);
        }
      }
    });

    // setup cargo bar
    this.cargoBayUsed.setX(0);
    this.cargoBayUsed.setY(380);
    this.cargoBayUsed.setWidth(500);
    this.cargoBayUsed.setHeight(20);
    this.cargoBayUsed.initialize();
    this.cargoBayUsed.setFont(FontSize.normal);
    this.cargoBayUsed.setText('Cargo Bay Used');

    // setup info view
    this.infoView.setWidth(500);
    this.infoView.setHeight(200);
    this.infoView.initialize();

    this.infoView.setX(0);
    this.infoView.setY(400);

    this.infoView.setFont(FontSize.normal);
    this.infoView.setOnClick(() => {});

    // setup action view
    this.actionView.setWidth(100);
    this.actionView.setHeight(600);
    this.actionView.initialize();

    this.actionView.setX(500);
    this.actionView.setY(0);

    this.actionView.setFont(FontSize.normal);
    this.actionView.setOnClick((r) => {
      // get action
      const a = r.listString();

      // perform action
      if (a === 'Unfit') {
        // get selected module
        const i: ShipViewRow = this.shipView.getSelectedItem();
        const metaRack = i.object.metaRack;

        // send unfit request
        const umMsg: ClientUnfitModule = {
          sid: this.wsSvc.sid,
          rack: metaRack,
          itemID: (i.object as WSModule).itemID,
        };

        this.wsSvc.sendMessage(MessageTypes.UnfitModule, umMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Trash') {
        // get selected item
        const i: ShipViewRow = this.shipView.getSelectedItem();

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
      } else if (a === 'Transfer') {
        // request property refresh
        this.refreshPropertySummary();

        // set up callback
        this.modalPropertyView.setOnClick((h) => {
          // get selected item
          const i: ShipViewRow = this.shipView.getSelectedItem();

          // get selected ship
          const t = h as FittingPropertyRow;

          // hide property modal
          this.hideModalPropertyList();

          // send transfer request
          const tiMsg: ClientTransferItem = {
            sid: this.wsSvc.sid,
            itemID: (i.object as WSContainerItem).id,
            receiverID: t.ship.id,
          };

          this.wsSvc.sendMessage(MessageTypes.TransferItem, tiMsg);

          // request cargo bay refresh
          this.refreshCargoBay();

          // reset views
          this.resetViews();
        });

        // show modal list
        this.showModalPropertyList();
      } else if (a === 'Package') {
        // get selected item
        const i: ShipViewRow = this.shipView.getSelectedItem();

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
        const i: ShipViewRow = this.shipView.getSelectedItem();

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
        const i: ShipViewRow = this.shipView.getSelectedItem();

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
        const i: ShipViewRow = this.shipView.getSelectedItem();

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
      } else if (a === 'Fit') {
        // get selected item
        const i: ShipViewRow = this.shipView.getSelectedItem();

        // send fit request
        const tiMsg: ClientFitModule = {
          sid: this.wsSvc.sid,
          itemID: (i.object as WSContainerItem).id,
        };

        this.wsSvc.sendMessage(MessageTypes.FitModule, tiMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Run') {
        // get selected item
        const i: ShipViewRow = this.shipView.getSelectedItem();

        // send run request
        const tiMsg: ClientRunSchematic = {
          sid: this.wsSvc.sid,
          itemId: (i.object as WSContainerItem).id,
        };

        this.wsSvc.sendMessage(MessageTypes.RunSchematic, tiMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Load Fuel') {
        // get selected item
        const i: ShipViewRow = this.shipView.getSelectedItem();

        // send consume request
        const tiMsg: ClientConsumeFuel = {
          sid: this.wsSvc.sid,
          itemId: (i.object as WSContainerItem).id,
        };

        this.wsSvc.sendMessage(MessageTypes.ConsumeFuel, tiMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Repair Ship') {
        // get selected item
        const i: ShipViewRow = this.shipView.getSelectedItem();

        // send consume request
        const tiMsg: ClientConsumeRepairKit = {
          sid: this.wsSvc.sid,
          itemId: (i.object as WSContainerItem).id,
        };

        this.wsSvc.sendMessage(MessageTypes.ConsumeRepairKit, tiMsg);

        // request cargo bay refresh
        this.refreshCargoBay();

        // reset views
        this.resetViews();
      } else if (a === 'Destruct ☠') {
        // send self destruct request
        const tiMsg: ClientSelfDestruct = {
          sid: this.wsSvc.sid,
        };

        this.wsSvc.sendMessage(MessageTypes.SelfDestruct, tiMsg);

        // reset views
        this.resetViews();
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
      this.getHeight() / 2 - this.modalInput.getHeight() / 2
    );
    this.modalInput.setFont(FontSize.large);
    this.modalInput.initialize();

    this.modalPropertyView.setWidth(400);
    this.modalPropertyView.setHeight(300);
    this.modalPropertyView.initialize();
    this.modalPropertyView.setX(
      this.getWidth() / 2 - this.modalPropertyView.getWidth() / 2
    );
    this.modalPropertyView.setY(
      this.getHeight() / 2 - this.modalPropertyView.getHeight() / 2
    );

    this.modalPropertyView.setFont(FontSize.normal);
    this.modalPropertyView.setOnClick(() => {});

    // pack
    this.addComponent(this.shipView);
    this.addComponent(this.infoView);
    this.addComponent(this.actionView);
    this.addComponent(this.cargoBayUsed);
  }

  private showModalInput() {
    this.removeComponent(this.shipView);
    this.removeComponent(this.infoView);
    this.removeComponent(this.actionView);
    this.removeComponent(this.cargoBayUsed);
    this.removeComponent(this.modalPropertyView);
    this.addComponent(this.modalOverlay);
    this.addComponent(this.modalInput);
  }

  private showModalPropertyList() {
    this.removeComponent(this.shipView);
    this.removeComponent(this.infoView);
    this.removeComponent(this.actionView);
    this.removeComponent(this.cargoBayUsed);
    this.removeComponent(this.modalInput);
    this.addComponent(this.modalOverlay);
    this.addComponent(this.modalPropertyView);
  }

  private hideModalInput() {
    this.addComponent(this.shipView);
    this.addComponent(this.infoView);
    this.addComponent(this.actionView);
    this.addComponent(this.cargoBayUsed);
    this.removeComponent(this.modalPropertyView);
    this.removeComponent(this.modalOverlay);
    this.removeComponent(this.modalInput);
  }

  private hideModalPropertyList() {
    this.addComponent(this.shipView);
    this.addComponent(this.infoView);
    this.addComponent(this.actionView);
    this.addComponent(this.cargoBayUsed);
    this.removeComponent(this.modalPropertyView);
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

  private refreshPropertySummary() {
    setTimeout(() => {
      const b = new ClientViewProperty();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewProperty, b);
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
      // reset ship view
      this.shipView.setSelectedIndex(-1);
      this.shipView.setItems([]);

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
        getItemSortString(a).localeCompare(getItemSortString(b))
      )) {
        const r = buildCargoRowFromContainerItem(ci, this.isDocked);
        cargo.push(r);
      }
    }

    this.cargoBayUsed.setPercentage(this.player.currentShip.cargoP);

    // update fitted module display
    const rackAMods: ShipViewRow[] = [];
    const rackBMods: ShipViewRow[] = [];
    const rackCMods: ShipViewRow[] = [];

    // build entries for modules on racks
    for (const m of this.player.currentShip.fitStatus.aRack.modules) {
      const d = buildFittingRowFromModule(m, this.isDocked);
      rackAMods.push(d);
    }

    for (const m of this.player.currentShip.fitStatus.bRack.modules) {
      const d = buildFittingRowFromModule(m, this.isDocked);
      rackBMods.push(d);
    }

    for (const m of this.player.currentShip.fitStatus.cRack.modules) {
      const d = buildFittingRowFromModule(m, this.isDocked);
      rackCMods.push(d);
    }

    // layout rack a
    rows.push(buildShipViewRowText('Rack A'));

    for (const r of rackAMods) {
      r.object.metaRack = 'A';
      rows.push(r);
    }

    rows.push(buildShipViewRowSpacer());

    // layout rack b
    rows.push(buildShipViewRowText('Rack B'));

    for (const r of rackBMods) {
      r.object.metaRack = 'B';
      rows.push(r);
    }

    rows.push(buildShipViewRowSpacer());

    // layout rack c
    rows.push(buildShipViewRowText('Rack C'));

    for (const r of rackCMods) {
      r.object.metaRack = 'C';
      rows.push(r);
    }

    rows.push(buildShipViewRowSpacer());

    // layout cargo bay
    rows.push(buildShipViewRowText('Cargo Bay'));

    for (const r of cargo) {
      rows.push(r);
    }

    rows.push(buildShipViewRowSpacer());

    // fix wallet
    if (!this.player.currentShip.wallet) {
      this.player.currentShip.wallet = 0;
    }

    // layout wallet
    rows.push(buildShipViewRowText('Wallet'));
    rows.push(buildShipViewRowText(`  ${this.player.currentShip.wallet} CBN`));

    rows.push(buildShipViewRowSpacer());

    if (!docked) {
      // layout danger zone
      rows.push(buildShipViewRowText('Danger Zone'));
      rows.push(buildActionRow(`  Self Destruct ☠`, ['Destruct ☠']));

      rows.push(buildShipViewRowSpacer());
    }

    // push to view
    const i = this.shipView.getSelectedIndex();

    this.shipView.setItems(rows);
    this.shipView.setSelectedIndex(i);
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  setPlayer(player: Player) {
    this.player = player;
  }

  resetViews() {
    // reset ship view
    this.shipView.setSelectedIndex(-1);
    this.shipView.setItems([]);

    // reset action view
    this.actionView.setItems([]);

    // reset info view
    this.infoView.setItems([]);
  }

  syncProperty(cache: ServerPropertyUpdate) {
    if (this.player.currentShip.dockedAtStationID) {
      const rows: FittingPropertyRow[] = [];

      // get owned ships docked at this station
      for (let e of cache.ships.sort((a, b) => a.name.localeCompare(b.name))) {
        if (e.dockedAtId == this.player.currentShip.dockedAtStationID) {
          // include in property selection options
          rows.push({
            ship: e,
            listString: () =>
              `${fixedString(e.name, 32)} ${fixedString(
                shortWallet(e.wallet),
                11
              )}`,
          });
        }
      }

      // update property modal list
      const pIDX = this.modalPropertyView.getScroll();

      this.modalPropertyView.setItems(rows);
      this.modalPropertyView.setScroll(pIDX);
    }
  }
}

class FittingPropertyRow {
  ship: ServerPropertyShipCacheEntry;
  listString: () => string;
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
      // exclude industrial market which is handled separately
      if (prop !== 'industrialmarket') {
        const v = buildShipViewRowText(
          infoKeyValueString(prop, `${meta[prop]}`)
        );

        rows.push(v);
      }
    }
  }

  // spacer after metadata info
  rows.push(buildShipViewRowSpacer());

  // industrial market
  const industrialMeta = m.itemTypeMeta['industrialmarket'];
  if (industrialMeta) {
    rows.push(buildShipViewRowText('Industrial Limits'));

    for (const key in industrialMeta) {
      if (Object.prototype.hasOwnProperty.call(industrialMeta, key)) {
        const value = industrialMeta[key];
        const addCBN = key === 'maxprice' || key === 'minprice';

        const v = buildShipViewRowText(
          infoKeyValueString(key, `${value}${addCBN ? ' CBN' : ''}`)
        );

        rows.push(v);
      }
    }

    // add spacer
    rows.push(buildShipViewRowSpacer());
  }

  // schematics
  if (m.schematic) {
    rows.push(buildShipViewRowText('Schematic Properties'));

    const v = buildShipViewRowText(
      infoKeyValueString('Job Time', `${m.schematic.time} seconds`)
    );

    rows.push(v);

    rows.push(buildShipViewRowSpacer());
    rows.push(buildShipViewRowText('Schematic Inputs'));

    for (const f of m.schematic.inputs) {
      const v = buildShipViewRowText(
        infoKeyValueString(f.itemTypeName, `${f.quantity}`)
      );

      rows.push(v);
    }

    rows.push(buildShipViewRowSpacer());
    rows.push(buildShipViewRowText('Schematic Outputs'));

    for (const f of m.schematic.outputs) {
      const v = buildShipViewRowText(
        infoKeyValueString(f.itemTypeName, `${f.quantity}`)
      );

      rows.push(v);
    }

    rows.push(buildShipViewRowSpacer());
  }

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
      // exclude industrial market which isn't shown on fitted modules
      if (prop !== 'industrialmarket') {
        const v = buildShipViewRowText(
          infoKeyValueString(prop, `${meta[prop]}`)
        );

        rows.push(v);
      }
    }
  }

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
        m.itemFamilyID === 'armor_plate' ||
        m.itemFamilyID === 'eng_oc' ||
        m.itemFamilyID === 'active_sink' ||
        m.itemFamilyID === 'drag_amp' ||
        m.itemFamilyID === 'utility_miner' ||
        m.itemFamilyID === 'utility_siphon' ||
        m.itemFamilyID === 'utility_cloak' ||
        m.itemFamilyID === 'battery_pack' ||
        m.itemFamilyID === 'aux_generator' ||
        m.itemFamilyID === 'cargo_expander' ||
        m.itemFamilyID === 'salvager' ||
        m.itemFamilyID === 'heat_sink';

      if (isModule) {
        // offer fit action
        actions.push('Fit');
      }

      // determine whether or not this is a schematic
      const isSchematic = m.itemFamilyID == 'schematic';

      if (isSchematic) {
        // offer run action
        actions.push('Run');
      }

      // determine whether or not this is fuel
      const isFuel = m.itemFamilyID === 'fuel';

      if (isFuel) {
        // offer load fuel action
        actions.push('Load Fuel');
      }

      // determine whether or not this is a repair kit
      const isRepairKit = m.itemFamilyID === 'repair_kit';

      if (isRepairKit) {
        // offer repair ship action
        actions.push('Repair Ship');
      }
    }

    actions.push('Transfer');

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

function buildActionRow(text: string, actions: string[]): ShipViewRow {
  const r: ShipViewRow = {
    object: {},
    actions: actions,
    listString: () => {
      return text;
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

function fixedString(str: string, width: number): string {
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

function shortWallet(d: number): string {
  d = Math.round(d);

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

function getItemSortString(item: WSContainerItem): string {
  return `${item.itemFamilyName}::${item.itemTypeName}::${item.quantity}::${item.id}`;
}
