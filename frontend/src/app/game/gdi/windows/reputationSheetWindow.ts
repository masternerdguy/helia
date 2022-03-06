import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import {
  GetFactionCache,
  GetPlayerFactionRelationshipCache,
} from '../../wsModels/shared';
import { Faction } from '../../engineModels/faction';
import { WSPlayerFactionRelationship } from '../../wsModels/entities/wsPlayerFaction';
import { GDITabPane } from '../components/gdiTabPane';
import { GDILabel } from '../components/gdiLabel';
import { Player } from '../../engineModels/player';
import { WsService } from '../../ws.service';
import { GDIButton } from '../components/gdiButton';

export class ReputationSheetWindow extends GDIWindow {
  private player: Player;
  private lastFactionId: string;
  private wsSvc: WsService;

  // tab pane
  private tabs = new GDITabPane();

  // "reputation" tab
  private factionList = new GDIList();
  private actionList = new GDIList();
  private infoList = new GDIList();

  // "my faction" tab
  private npcFactionLabel = new GDILabel();
  private createFactionButton = new GDIButton();

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(420);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Reputation Sheet');

    // tab list
    this.tabs.setWidth(this.getWidth());
    this.tabs.setHeight(this.getHeight());
    this.tabs.initialize();

    this.tabs.setX(0);
    this.tabs.setY(0);
    this.tabs.setSelectedTab('Reputation');

    this.addComponent(this.tabs);

    // pack tabs
    this.packReputationTab();
    this.packMyFactionTab();
  }

  private packMyFactionTab() {
    // npc faction indicator label
    this.npcFactionLabel.setWidth(this.getWidth());
    this.npcFactionLabel.setHeight(Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2));
    this.npcFactionLabel.initialize();

    this.npcFactionLabel.setText('You are currently a member of an NPC faction.')
    this.npcFactionLabel.setFont(FontSize.normal);

    this.npcFactionLabel.setX(0);
    this.npcFactionLabel.setY(GDIStyle.tabHandleHeight);

    // create faction button
    this.createFactionButton.setWidth(this.getWidth() * 0.5);
    this.createFactionButton.setHeight(Math.round(GDIStyle.getUnderlyingFontSize(FontSize.normal) * 2));
    this.createFactionButton.initialize();

    this.createFactionButton.setText('Create Faction');
    this.createFactionButton.setFont(FontSize.normal);

    this.createFactionButton.setX(this.getWidth() * 0.25);
    this.createFactionButton.setY(100 + GDIStyle.tabHandleHeight);

    this.createFactionButton.setOnClick(() => {
      console.log("todo - new faction button");
    });
  }

  private packReputationTab() {
    // faction list
    this.factionList.setWidth(300);
    this.factionList.setHeight(200 + GDIStyle.tabHandleHeight);
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

    this.tabs.addComponent(this.factionList, 'Reputation');

    // action list
    this.actionList.setWidth(100);
    this.actionList.setHeight(200 + GDIStyle.tabHandleHeight);
    this.actionList.initialize();

    this.actionList.setX(300);
    this.actionList.setY(0);

    this.actionList.setFont(FontSize.normal);
    this.actionList.setOnClick((item) => {
      console.log(item); // todo
    });

    this.tabs.addComponent(this.actionList, 'Reputation');

    // info list
    this.infoList.setWidth(400);
    this.infoList.setHeight(220 - GDIStyle.tabHandleHeight);
    this.infoList.initialize();

    this.infoList.setX(0);
    this.infoList.setY(200);

    this.infoList.setFont(FontSize.normal);
    this.infoList.setOnClick(() => { });

    this.tabs.addComponent(this.infoList, 'Reputation');
  }

  private buildDetails(r: FactionRepViewRow): FactionInfoViewRow[] {
    const relationships = r.faction.relationships.sort(
      (a, b) => a.standingValue - b.standingValue
    );

    const factions = GetFactionCache().sort((a, b) => {
      // get standing entries
      const aStanding = relationships.filter((x) => x.factionId == a.id);
      const bStanding = relationships.filter((x) => x.factionId == b.id);

      // extract values
      var aVal = 0;
      var bVal = 0;

      if (aStanding.length > 0) {
        aVal = aStanding[0].standingValue;
      }

      if (bStanding.length > 0) {
        bVal = bStanding[0].standingValue;
      }

      // sort by standing desc
      return bVal - aVal;
    });

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

    // description
    rows.push('');
    rows.push('Encyclopedia Summary');
    rows.push('');

    const displayDescription = this.infoList.breakText(
      r.faction.description ?? ''
    );

    for (const i of displayDescription) {
      rows.push(i.text);
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

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  setPlayer(player: Player) {
    this.player = player;
    const faction = player.getFaction();

    // check for faction change (or initial presentation)
    if (this.lastFactionId != faction.id) {
      this.lastFactionId == faction.id;

      if (faction.isNPC) {
        this.showNPCComponentsOnMyFactionTab();
      } else {
        this.hideNPCComponentsOnMyFactionTab();
      }
    }
  }

  periodicUpdate() {
    if (!this.isHidden()) {
      // get player-faction relationships
      const playerFactionRelationships = GetPlayerFactionRelationshipCache();

      // get factions
      const factions = GetFactionCache().sort((a, b) => {
        // get standing entries
        const aStanding = playerFactionRelationships.filter(
          (x) => x.factionId == a.id
        );

        const bStanding = playerFactionRelationships.filter(
          (x) => x.factionId == b.id
        );

        // extract values
        var aVal = 0;
        var bVal = 0;

        if (aStanding.length > 0) {
          aVal = aStanding[0].standingValue;
        }

        if (bStanding.length > 0) {
          bVal = bStanding[0].standingValue;
        }

        // sort by standing desc
        return bVal - aVal;
      });

      // build rows
      const factionRows: FactionRepViewRow[] = [];

      for (const f of factions) {
        let playerRel: WSPlayerFactionRelationship = null;

        // find relationship to player
        for (const rel of playerFactionRelationships) {
          if (rel.factionId == f.id) {
            playerRel = rel;
            break;
          }
        }

        factionRows.push({
          faction: f,
          actions: [],
          listString: () => factionListRowString(playerRel, f),
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

  private showNPCComponentsOnMyFactionTab() {
    this.tabs.addComponent(this.npcFactionLabel, 'My Faction');
    this.tabs.addComponent(this.createFactionButton, 'My Faction');
  }

  private hideNPCComponentsOnMyFactionTab() {
    this.tabs.removeComponent(this.npcFactionLabel, 'My Faction');
    this.tabs.removeComponent(this.createFactionButton, 'My Faction');
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

function factionListRowString(
  rel: WSPlayerFactionRelationship,
  faction: Faction
): string {
  if (rel == null || faction == null) {
    return;
  }

  // determine whether or not to show hostility flag
  let openHostileFlag = '';

  if (rel.openlyHostile) {
    openHostileFlag = '⚔';
  }

  // determine whether or not the player is a member
  let memberFlag = '';

  if (rel.isMember) {
    memberFlag = '✪';
  }

  // build string
  return `${fixedString(faction.name, 24)} ${fixedString(
    `[${rel.standingValue.toFixed(2)}] `,
    10
  )} ${fixedString(memberFlag, 1)}${fixedString(openHostileFlag, 1)}`;
}

function infoKeyValueString(key: string, value: string) {
  // build string
  return `${fixedString('', 1)} ${fixedString(key, 24)} ${fixedString(
    value,
    24
  )}`;
}

function fixedString(str: string, width: number): string {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}
