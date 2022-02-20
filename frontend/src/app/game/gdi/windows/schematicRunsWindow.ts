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
    this.setWidth(500);
    this.setHeight(300);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Schematic Runs');

    // schematic runs list
    this.schematicRunList.setWidth(500);
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

  sync(cache: ServerSchematicRunsUpdate) {
    console.log("not yet implemented");
  }

  private refreshSchematicRuns() {
    setTimeout(() => {
      const b = new ClientViewSchematicRuns();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewSchematicRuns, b);
    }, 200);
  }
}
