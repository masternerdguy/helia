import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { MessageTypes } from '../../wsModels/gameMessage';
import { WsService } from '../../ws.service';
import { ServerSchematicRunsUpdate } from '../../wsModels/bodies/schematicRunsUpdate';
import { ClientViewSchematicRuns } from '../../wsModels/bodies/viewSchematicRuns';

export class SchematicRunsWindow extends GDIWindow {
  private schematicRunList = new GDIList();

  // last schematic run refresh
  private lastSchematicRunView: number = 0;

  // ws service
  private wsSvc: WsService;

  initialize() {
    // set dimensions
    this.setWidth(650);
    this.setHeight(300);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Schematic Runs');

    // schematic runs list
    this.schematicRunList.setWidth(650);
    this.schematicRunList.setHeight(300);
    this.schematicRunList.initialize();

    this.schematicRunList.setX(0);
    this.schematicRunList.setY(0);

    this.schematicRunList.setFont(FontSize.normal);
    this.schematicRunList.setOnClick((item) => {});

    this.addComponent(this.schematicRunList);
  }

  periodicUpdate() {
    if (!this.isHidden()) {
      // check if schematic run view is stale
      const now = Date.now();

      if (now - this.lastSchematicRunView > 5000) {
        // refresh summary
        this.refreshSchematicRuns();
        this.lastSchematicRunView = now;
      }
    }
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  sync(update: ServerSchematicRunsUpdate) {
    // sort runs by progress descending
    const sortedRuns = update.runs.sort((a, b) => {
      return b.percentageComplete - a.percentageComplete;
    });

    // stash scroll position
    const scrollIdx = this.schematicRunList.getScroll();

    // clear list
    this.schematicRunList.setItems([]);

    // update list
    const rows: string[] = [];

    for (const e of sortedRuns) {
      const percentageComplete = `${(e.percentageComplete * 100).toFixed(2)}`;

      rows.push(
        fixedString('', 2) +
          ' ' +
          fixedString(e.schematicName, 16) +
          ' ' +
          fixedString(e.statusId, 8) +
          ' ' +
          fixedString(e.hostShipName, 16) +
          ' ' +
          fixedString(e.hostStationName, 16) +
          ' ' +
          fixedString(e.solarSystemName, 16) +
          ' ' +
          fixedString(`(${percentageComplete})`, 8) +
          '~'
      );
    }

    this.schematicRunList.setItems(rows);

    // restore scroll position
    this.schematicRunList.setScroll(scrollIdx);
  }

  private refreshSchematicRuns() {
    setTimeout(() => {
      const b = new ClientViewSchematicRuns();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewSchematicRuns, b);
    }, 200);
  }
}

function fixedString(str: string, width: number): string {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}
