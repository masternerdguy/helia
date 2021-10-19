import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { GetFactionCache } from '../../wsModels/shared';
import { Faction } from '../../engineModels/faction';

export class ReputationSheetWindow extends GDIWindow {
  private factionList = new GDIList();
  private actionList = new GDIList();
  private infoList = new GDIList();

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
      // get faction row
      const o = item as FactionRepViewRow;

      // update actions
      this.actionList.setItems(o.actions);

      // update detailed info
      const details = this.buildDetails(o);
      this.infoList.setItems(details);
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

    // info list
    this.infoList.setWidth(400);
    this.infoList.setHeight(200);
    this.infoList.initialize();

    this.infoList.setX(0);
    this.infoList.setY(200);

    this.infoList.setFont(FontSize.normal);

    this.addComponent(this.infoList);
  }

  private buildDetails(r: FactionRepViewRow): FactionInfoViewRow[] {
    const factions = GetFactionCache();
    const rows: string[] = [];

    // basic info
    rows.push('Basic');
    rows.push(infoKeyValueString('Name', r.faction.name));
    rows.push(infoKeyValueString('Ticker', `[${r.faction.ticker}]`));
    rows.push('');

    // relationships
    rows.push('Liked By');
    for (const f of factions) {
      if (f.id != r.faction.id && !f.isClosed) {
        // find relationship
        for (const rel of f.relationships) {
          if (rel.factionId != r.faction.id) {
            continue;
          }

          if (rel.standingValue > 0) {
            rows.push(
              infoKeyValueString(f.name, `[${rel.standingValue.toFixed(2)}] `)
            );
          }
        }
      }
    }

    rows.push('');

    rows.push('Disliked By');
    for (const f of factions) {
      if (f.id != r.faction.id && !f.isClosed) {
        // find relationship
        for (const rel of f.relationships) {
          if (rel.factionId != r.faction.id) {
            continue;
          }

          if (rel.standingValue < 0) {
            let openHostileFlag = '';

            if (rel.openlyHostile) {
              openHostileFlag = '⚔';
            }

            rows.push(
              infoKeyValueString(
                f.name,
                `[${rel.standingValue.toFixed(2)}] ` + openHostileFlag
              )
            );
          }
        }
      }
    }

    // convert to display rows
    const dRows: FactionInfoViewRow[] = [];

    for (const r of rows) {
      dRows.push(this.infoRowFromString(r));
    }

    // return converted rows
    return dRows;
  }

  private infoRowFromString(s: string): FactionInfoViewRow {
    return {
      listString: () => s,
    };
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
          listString: () => `${f.name}`,
        });
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

class FactionInfoViewRow {
  listString: () => string;
}

function infoKeyValueString(key: string, value: string) {
  // build string
  return `${fixedString('', 1)} ${fixedString(key, 24)} ${fixedString(
    value,
    24
  )}~`;
}

function fixedString(str: string, width: number) {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}
