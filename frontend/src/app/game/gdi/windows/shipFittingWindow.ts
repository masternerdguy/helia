import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { Ship } from '../../engineModels/ship';
import { GDIList } from '../components/gdiList';
import { WSModule } from '../../wsModels/entities/wsShip';
import { WsService } from '../../ws.service';
import { ClientActivateModule } from '../../wsModels/bodies/activateModule';
import { MessageTypes } from '../../wsModels/gameMessage';
import { Player } from '../../engineModels/player';
import { ClientDeactivateModule } from '../../wsModels/bodies/deactivateModule';

export class ShipFittingWindow extends GDIWindow {
  // lists
  private shipView: GDIList = new GDIList();
  private infoView: GDIList = new GDIList();
  private actionView: GDIList = new GDIList();

  // ship being fit
  private ship: Ship;
  private player: Player;

  // ws
  private wsSvc: WsService;

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
    this.shipView.setOnClick((row: ShipViewRow) => {
      // todo
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
    this.actionView.setOnClick(() => {});

    // pack
    this.addComponent(this.shipView);
    this.addComponent(this.infoView);
    this.addComponent(this.actionView);
  }

  periodicUpdate() {
    if (this.ship !== undefined && this.ship !== null) {
      if (this.ship.fitStatus !== undefined && this.ship.fitStatus !== null) {
        const rows: ShipViewRow[] = [];

        // update fitted module display
        const rackAMods: ShipViewRow[] = [];
        const rackBMods: ShipViewRow[] = [];
        const rackCMods: ShipViewRow[] = [];

        // build entries for modules on racks
        for (const m of this.ship.fitStatus.aRack.modules) {
          const d = buildFittingRowFromModule(m);
          rackAMods.push(d);
        }

        for (const m of this.ship.fitStatus.bRack.modules) {
          const d = buildFittingRowFromModule(m);
          rackBMods.push(d);
        }

        for (const m of this.ship.fitStatus.cRack.modules) {
          const d = buildFittingRowFromModule(m);
          rackCMods.push(d);
        }

        // layout rack a
        rows.push(buildShipViewRowText('Rack A'));

        for (const r of rackAMods) {
          rows.push(r);
        }

        rows.push(buildShipViewRowSpacer());

        // layout rack b
        rows.push(buildShipViewRowText('Rack B'));

        for (const r of rackBMods) {
          rows.push(r);
        }

        rows.push(buildShipViewRowSpacer());

        // layout rack c
        rows.push(buildShipViewRowText('Rack C'));

        for (const r of rackCMods) {
          rows.push(r);
        }

        rows.push(buildShipViewRowSpacer());

        // push to view
        const i = this.shipView.getSelectedIndex();

        this.shipView.setItems(rows);
        this.shipView.setSelectedIndex(i);
      }
    }
  }

  setShip(ship: Ship) {
    this.ship = ship;
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  setPlayer(player: Player) {
    this.player = player;
  }
}

class ShipViewRow {
  object: WSModule;
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
