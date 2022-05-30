import { heliaDateFromString, printHeliaDate } from '../../engineMath';
import { WsService } from '../../ws.service';
import {
  ClientViewActionReportsPage,
  ServerActionReportsPage,
  ServerActionReportSummary,
} from '../../wsModels/bodies/viewActionReportsPage';
import { MessageTypes } from '../../wsModels/gameMessage';
import { FontSize } from '../base/gdiStyle';
import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';

export class ActionReportsWindow extends GDIWindow {
  private reportView: GDIList = new GDIList();

  private pageData: ServerActionReportsPage;
  private wsSvc: WsService;
  private page: number;
  private pageSize: number;

  initialize() {
    // set dimensions
    this.setWidth(850);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Action Reports');
    this.page = 0;
    this.pageSize = 30;

    // setup info view
    this.reportView.setWidth(850);
    this.reportView.setHeight(400);
    this.reportView.initialize();

    this.reportView.setX(0);
    this.reportView.setY(0);

    this.reportView.setFont(FontSize.normal);
    this.reportView.setOnClick((r) => {
      const row = r as ActionReportWindowRow;

      if (row.summary) {
        // todo
      } else {
        if (row.listString() == `==> Next Page`) {
          this.page++;
        } else if (row.listString() == `<== Previous Page`) {
          this.page--;
        }
      }
    });

    // request updates on show
    this.setOnShow(() => {
      this.refreshPage();
    });

    // request periodic updates when shown
    setInterval(() => {
      if (!this.isHidden()) {
        this.refreshPage();
      }
    }, 5000);

    // pack
    this.addComponent(this.reportView);
  }

  periodicUpdate() {
    if (!this.isHidden()) {
      if (this.pageData) {
        // set up for redraw
        const rows: ActionReportWindowRow[] = [];
        const scroll = this.reportView.getScroll();
        const idx = this.reportView.getSelectedIndex();

        // header
        rows.push(makeTextRow(`---- Page ${this.pageData.page + 1} ----`));

        if (this.pageData.page != this.page) {
          rows.push(makeTextRow(`Loading page ${this.page + 1}...`));
        } else {
          // spacer
          rows.push(makeSpacerRow());

          // navigation
          if (this.pageData.logs.length > 0) {
            rows.push(makeTextRow(`==> Next Page`));
          }

          if (this.page > 0) {
            rows.push(makeTextRow(`<== Previous Page`));
          }

          // spacer
          rows.push(makeSpacerRow());

          if (this.pageData.logs.length == 0) {
            rows.push(makeTextRow(`Nothing to display.`));
          } else {
            // summary rows
            for (const r of this.pageData.logs) {
              rows.push(makeReportRow(r));
            }
          }
        }

        // update view
        this.reportView.setItems(rows);
        this.reportView.setScroll(scroll);
        this.reportView.setSelectedIndex(idx);
      }
    }
  }

  setWsService(wsSvc: WsService) {
    this.wsSvc = wsSvc;
  }

  setPageData(page: ServerActionReportsPage) {
    this.pageData = page;
  }

  private refreshPage() {
    setTimeout(() => {
      const b = new ClientViewActionReportsPage();
      b.sid = this.wsSvc.sid;
      b.page = this.page;
      b.count = this.pageSize;

      this.wsSvc.sendMessage(MessageTypes.ViewActionReportsPage, b);
    }, 200);
  }
}

function makeReportRow(s: ServerActionReportSummary): ActionReportWindowRow {
  return {
    summary: s,
    listString: () => {
      return (
        ` ${fixedString(s.victim, 16)}` +
        ` ${fixedString(s.ship, 16)}` +
        ` ${s.ticker.length > 0 ? fixedString('[' + s.ticker + ']', 5) : fixedString('', 5)}` +
        ` ${fixedString(quantity(s.parties), 8)}` +
        ` ${fixedString(s.systemName, 16)}` +
        ` ${fixedString(s.regionName, 16)}` +
        ` ${printHeliaDate(heliaDateFromString(s.timestamp))}`
      );
    },
  };
}

function makeTextRow(s: string): ActionReportWindowRow {
  return {
    summary: undefined,
    listString: () => {
      return `${s}`;
    },
  };
}

function makeSpacerRow(): ActionReportWindowRow {
  return {
    summary: undefined,
    listString: () => {
      return ``;
    },
  };
}

function fixedString(str: string, width: number): string {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str.substr(0, width).padEnd(width);
}

function quantity(d: number): string {
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

class ActionReportWindowRow {
  summary: ServerActionReportSummary;

  listString: () => string;
}
