import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { GetFactionCache } from '../../wsModels/shared';
import { Faction } from '../../engineModels/faction';

export class ReputationSheetWindow extends GDIWindow {
  private factionList = new GDIList();
  private actionList = new GDIList();

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Reputation Sheet');

    // faction list
    this.factionList.setWidth(300);
    this.factionList.setHeight(200);
    this.factionList.initialize();

    this.factionList.setX(0);
    this.factionList.setY(0);

    this.factionList.setFont(FontSize.normal);
    this.factionList.setOnClick((item) => {
      const o = item as FactionRepViewRow;

      this.actionList.setItems(o.actions)
    });

    this.addComponent(this.factionList);

    // action list
    this.actionList.setWidth(100);
    this.actionList.setHeight(200);
    this.actionList.initialize();

    this.actionList.setX(300);
    this.actionList.setY(0);

    this.actionList.setFont(FontSize.normal);
    this.actionList.setOnClick((item) => {
      console.log(item); // todo
    });

    this.addComponent(this.actionList);
  }

  periodicUpdate() {
    if (!this.isHidden()) {
      // get factions
      const factions = GetFactionCache();

      // build rows
      const factionRows: FactionRepViewRow[] = [];

      for (const f of factions) {
        factionRows.push({
          faction: f,
          actions: [],
          listString: () => `${f.name}`
        })
      }

      // update faction list
      const sp = this.factionList.getScroll();
      const si = this.factionList.getSelectedIndex();

      this.factionList.setItems(factionRows);
      this.factionList.setScroll(sp);
      this.factionList.setSelectedIndex(si);
    }
  }
}

class FactionRepViewRow {
  faction: Faction;
  actions: string[];

  listString: () => string;
}
