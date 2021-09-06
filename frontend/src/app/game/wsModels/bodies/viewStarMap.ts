export class ClientViewStarMap {
  sid: string;
}

export class ServerStarMapUpdate {
  cachedMapData: string;
}

export class UnwrappedStarMapData {
  regions: StarMapRegion[] = [];
  edges: StarMapEdge[] = [];
  flattened: StarMapFlatSystem[] = [];

  constructor(cachedMapData: string) {
    // decode as JSON
    const asJSON = JSON.parse(cachedMapData);

    // store regions
    this.regions = asJSON.regions;

    // store edges
    this.edges = asJSON.edges;

    // flatten for quick reference
    for (let region of this.regions) {
      for (let system of region.systems) {
        // build flattened systtem
        const flat: StarMapFlatSystem = {
          region: region,
          system: system,
          edges: [],
        };

        // get all edges
        for (let edge of this.edges) {
          if (edge.aID == system.id) {
            flat.edges.push([edge, this.findRawSystemByID(edge.bID)]);
          }

          if (edge.bID == system.id) {
            flat.edges.push([edge, this.findRawSystemByID(edge.aID)]);
          }
        }

        // store flattened system
        this.flattened.push(flat);
      }
    }
  }

  findRawSystemByID(id: string): StarMapSolarSystem {
    for (let region of this.regions) {
      for (let system of region.systems) {
        if (system.id == id) {
          return system;
        }
      }
    }
  }
}

export class StarMapFlatSystem {
  system: StarMapSolarSystem;
  region: StarMapRegion;
  edges: [StarMapEdge, StarMapSolarSystem][];
}

export class StarMapSolarSystem {
  factionId: string;
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
  aID: string;
  bID: string;
  transient: boolean;
}
