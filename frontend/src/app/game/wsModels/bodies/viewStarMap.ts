export class ClientViewStarMap {
  sid: string;
}

export class ServerStarMapUpdate {
  cachedMapData: string;
}

export class UnwrappedStarMapData {
  regions: StarMapRegion[];
  edges: StarMapEdge[];

  constructor(cachedMapData: string) {
    // decode as JSON
    const asJSON = JSON.parse(cachedMapData);

    // store regions
    this.regions = asJSON.regions;

    // store edges
    this.edges = asJSON.edges;

    // debug out
    console.log(this);
  }
}

export class StarMapSolarSystem {
  x: number;
  y: number;
  name: string;
  id: string;
}

export class StarMapRegion {
  systems: StarMapSolarSystem[];
  x: number;
  y: number;
  name: string;
  id: string;
}

export class StarMapEdge {
  a: string;
  b: string;
}
