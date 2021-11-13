export class ServerPropertyUpdate {
    ships: ServerPropertyShipCacheEntry[];
}

export class ServerPropertyShipCacheEntry {
    name: string;
    texture: string;
    id: string;
    systemId: string;
    systemName: string;
    dockedAtID: string;
    dockedAtName: string
}
