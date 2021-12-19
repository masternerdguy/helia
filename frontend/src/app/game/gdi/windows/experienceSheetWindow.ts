import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIList } from '../components/gdiList';
import { MessageTypes } from '../../wsModels/gameMessage';
import { WsService } from '../../ws.service';
import {
  ServerExperienceModuleEntry,
  ServerExperienceShipEntry,
  ServerExperienceUpdate,
} from '../../wsModels/bodies/experienceUpdate';
import { ClientViewExperience } from '../../wsModels/bodies/viewExperience';

export class ExperienceSheetWindow extends GDIWindow {
  private experienceList = new GDIList();

  // last experience refresh
  private lastExperienceView: number = 0;

  // ws service
  private wsSvc: WsService;

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Experience Sheet');

    // experience list
    this.experienceList.setWidth(400);
    this.experienceList.setHeight(400);
    this.experienceList.initialize();

    this.experienceList.setX(0);
    this.experienceList.setY(0);

    this.experienceList.setFont(FontSize.normal);
    this.experienceList.setOnClick((item) => {});

    this.addComponent(this.experienceList);
  }

  periodicUpdate() {
    if (!this.isHidden()) {
      // check if experience view is stale
      const now = Date.now();

      if (now - this.lastExperienceView > 5000) {
        // refresh summary
        this.refreshExperienceSummary();
        this.lastExperienceView = now;
      }
    }
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  sync(cache: ServerExperienceUpdate) {
    // stash selected index
    const sIdx = this.experienceList.getSelectedIndex();

    const rows: ExperienceSheetViewRow[] = [];

    // handle ship template entries
    {
      // add header
      rows.push(...makeSectionHeader('Ship Models'));

      // sort entries by level, then name, then id
      const sorted = cache.ships
        .sort((a, b) =>
          this.getShipSortKey(a).localeCompare(this.getShipSortKey(b))
        )
        .reverse();

      // build ship entries
      for (const s of sorted) {
        const r = new ExperienceSheetViewRow();
        const ls = experienceSheetViewRowStringFromShip(s);

        r.listString = () => ls;
        r.ship = s;

        rows.push(r);
      }
    }

    // handle module entries
    {
      // spacer
      rows.push({
        listString: () => '',
      });

      // add header
      rows.push(...makeSectionHeader('Module Types'));

      // sort entries by level, then name, then id
      const sorted = cache.modules
        .sort((a, b) =>
          this.getModuleSortKey(a).localeCompare(this.getModuleSortKey(b))
        )
        .reverse();

      // build module entries
      for (const s of sorted) {
        const r = new ExperienceSheetViewRow();
        const ls = experienceSheetViewRowStringFromModule(s);

        r.listString = () => ls;
        r.module = s;

        rows.push(r);
      }
    }

    // update list
    this.experienceList.setItems(rows);

    // re-select index
    this.experienceList.setSelectedIndex(sIdx);
  }

  private refreshExperienceSummary() {
    setTimeout(() => {
      const b = new ClientViewExperience();
      b.sid = this.wsSvc.sid;

      this.wsSvc.sendMessage(MessageTypes.ViewExperience, b);
    }, 200);
  }

  private getShipSortKey(s: ServerExperienceShipEntry): string {
    return `${s.experienceLevel}::${s.shipTemplateName}::${s.shipTemplateID}`;
  }

  private getModuleSortKey(s: ServerExperienceModuleEntry): string {
    return `${s.experienceLevel}::${s.itemTypeName}::${s.itemTypeID}`;
  }
}

class ExperienceSheetViewRow {
  listString: () => string;
  ship?: ServerExperienceShipEntry;
  module?: ServerExperienceModuleEntry;
}

function makeSectionHeader(title: string): ExperienceSheetViewRow[] {
  const rows: ExperienceSheetViewRow[] = [];

  rows.push({
    listString: () => title,
  });

  let underline = '';

  for (let i = 0; i < title.length; i++) {
    underline += '-';
  }

  rows.push({
    listString: () => underline,
  });

  rows.push({
    listString: () => ' ',
  });

  return rows;
}

function experienceSheetViewRowStringFromShip(
  s: ServerExperienceShipEntry
): string {
  if (s == null) {
    return;
  }

  // build string
  return (
    `${fixedString('', 2)} ` +
    `${fixedString(s.shipTemplateName, 40)} ` +
    `${fixedString(`[${s.experienceLevel.toFixed(2)}]`, 8)} `
  );
}

function experienceSheetViewRowStringFromModule(
  s: ServerExperienceModuleEntry
): string {
  if (s == null) {
    return;
  }

  // build string
  return (
    `${fixedString('', 2)} ` +
    `${fixedString(s.itemTypeName, 40)} ` +
    `${fixedString(`[${s.experienceLevel.toFixed(2)}]`, 8)} `
  );
}

function fixedString(str: string, width: number): string {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}
