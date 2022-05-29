export class ClientViewActionReportsPage {
  sid: string;
  page: number;
  count: number;
}

export class ServerActionReportsPage {
  page: number;
  logs: ServerActionReportSummary[];
}

export class ServerActionReportSummary {
  id: string;
  victim: string;
  ship: string;
  ticker: string;
  timestamp: string;
  systemName: string;
  regionName: string;
  parties: number;
}
