export class ServerPropertyUpdate {
  ships: ServerPropertyShipCacheEntry[];
  outposts: ServerPropertyOutpostCacheEntry[];
}

export class ServerPropertyShipCacheEntry {
  name: string;
  texture: string;
  id: string;
  systemId: string;
  systemName: string;
  dockedAtId: string;
  dockedAtName: string;
  wallet: number;
}

export class ServerPropertyOutpostCacheEntry {
  name: string;
  texture: string;
  id: string;
  systemId: string;
  systemName: string;
  wallet: number;
}
