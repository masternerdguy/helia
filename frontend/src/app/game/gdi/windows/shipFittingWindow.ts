import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { WSModule } from '../../wsModels/entities/wsShip';
import { WsService } from '../../ws.service';
import { Player } from '../../engineModels/player';
import { ClientViewCargoBay } from '../../wsModels/bodies/viewCargoBay';
import { GameMessage, MessageTypes } from '../../wsModels/gameMessage';
import { WSContainerItem } from '../../wsModels/entities/wsContainer';
import { ClientUnfitModule } from '../../wsModels/bodies/unfitModule';

export class ShipFittingWindow extends GDIWindow {
  // lists
  private shipView: GDIList = new GDIList();
  private infoView: GDIList = new GDIList();
  private actionView: GDIList = new GDIList();

  // player
  private player: Player;

  // ws service
  private wsSvc: WsService;

  // last cargo bay refresh
  private lastCargoView = 0;

  initialize() {
    // set dimensions
    this.setWidth(600);
    this.setHeight(600);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Ship Fitting');

    // setup ship view
    this.shipView.setWidth(500);
    this.shipView.setHeight(400);
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
      }
    });

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
          itemID: (i.object as WSModule).itemID
        };

        this.wsSvc.sendMessage(MessageTypes.UnfitModule, umMsg);

        // request cargo bay refresh
        setTimeout(() => {
          const b = new ClientViewCargoBay();
          b.sid = this.wsSvc.sid;

          this.wsSvc.sendMessage(MessageTypes.ViewCargoBay, b);
        }, 200);
      }
    });

    // pack
    this.addComponent(this.shipView);
    this.addComponent(this.infoView);
    this.addComponent(this.actionView);
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
        const r = buildCargoRowFromContainerItem(ci);
        cargo.push(r);
      }
    }

    // update fitted module display
    const rackAMods: ShipViewRow[] = [];
    const rackBMods: ShipViewRow[] = [];
    const rackCMods: ShipViewRow[] = [];

    // build entries for modules on racks
    for (const m of this.player.currentShip.fitStatus.aRack.modules) {
      const d = buildFittingRowFromModule(m);
      rackAMods.push(d);
    }

    for (const m of this.player.currentShip.fitStatus.bRack.modules) {
      const d = buildFittingRowFromModule(m);
      rackBMods.push(d);
    }

    for (const m of this.player.currentShip.fitStatus.cRack.modules) {
      const d = buildFittingRowFromModule(m);
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

function buildCargoRowFromContainerItem(m: WSContainerItem): ShipViewRow {
  const r: ShipViewRow = {
    object: m,
    actions: ['Trash'],
    listString: () => {
      return itemStatusString(m);
    },
  };

  return r;
}

function buildFittingRowFromModule(m: WSModule): ShipViewRow {
  const r: ShipViewRow = {
    object: m,
    actions: ['Unfit'],
    listString: () => {
      return moduleStatusString(m);
    },
  };

  return r;
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
    m.type,
    24
  )} ${pc}`;
}

function itemStatusString(m: WSContainerItem) {
  // build status string
  return `${fixedString('', 1)} ${fixedString(
    m.itemTypeName,
    24
  )}`;
}
