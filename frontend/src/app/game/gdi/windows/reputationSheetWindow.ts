import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { GetFactionCache } from '../../wsModels/shared';
import { Faction } from '../../engineModels/faction';

export class ReputationSheetWindow extends GDIWindow {
  private factionList = new GDIList();

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

    this.factionList.setFont(FontSize.large);
    this.factionList.setOnClick((item) => {
      console.log(item);
    });

    this.addComponent(this.factionList);
  }

  periodicUpdate() {
    if (!this.isHidden()) {
      // get factions
      const factions = GetFactionCache()

      // build rows
      const factionRows: FactionRepViewRow[] = []

      for (const f of factions) {
        factionRows.push({
          faction: f,
          actions: [],
          listString: () => `${f.name}`
        })
      }

      // update faction list
      const sp = this.factionList.getScroll()
      const si = this.factionList.getSelectedIndex()

      this.factionList.setItems(factionRows)
      this.factionList.setScroll(sp)
      this.factionList.setSelectedIndex(si)
    }
  }
}

class FactionRepViewRow {
  faction: Faction;
  actions: string[];

  listString: () => string;
}
