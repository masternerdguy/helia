import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { Player, TargetType } from '../../engineModels/player';
import { GDITabPane } from '../components/gdiTabPane';
import { Ship } from '../../engineModels/ship';
import { GetPlayerFactionRelationshipCacheEntry } from '../../wsModels/shared';

export class OverviewWindow extends GDIWindow {
  tabs = new GDITabPane();
  globalList = new GDIList();
  shipList = new GDIList();
  stationList = new GDIList();
  jumpholeList = new GDIList();
  selectedItemID: string;
  selectedItemType: TargetType;

  initialize() {
    // set dimensions
    this.setWidth(370);
    this.setHeight(this.getHeight());

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('System Overview');

    // tab list
    this.tabs.setWidth(this.getWidth());
    this.tabs.setHeight(this.getHeight());
    this.tabs.initialize();

    this.tabs.setX(0);
    this.tabs.setY(0);
    this.tabs.setSelectedTab('Global');

    this.addComponent(this.tabs);

    // all object list
    this.globalList.setWidth(this.getWidth());
    this.globalList.setHeight(this.getHeight());
    this.globalList.initialize();

    this.globalList.setX(0);
    this.globalList.setY(0);

    this.globalList.setFont(FontSize.large);
    this.globalList.setOnClick((item: OverviewRow) => {
      this.selectedItemID = item.object.id;
      this.selectedItemType = item.type;
    });

    this.tabs.addComponent(this.globalList, 'Global');

    // ship list
    this.shipList.setWidth(this.getWidth());
    this.shipList.setHeight(this.getHeight());
    this.shipList.initialize();

    this.shipList.setX(0);
    this.shipList.setY(0);

    this.shipList.setFont(FontSize.large);
    this.shipList.setOnClick((item: OverviewRow) => {
      this.selectedItemID = item.object.id;
      this.selectedItemType = item.type;
    });

    this.tabs.addComponent(this.shipList, 'Ships');

    // station list
    this.stationList.setWidth(this.getWidth());
    this.stationList.setHeight(this.getHeight());
    this.stationList.initialize();

    this.stationList.setX(0);
    this.stationList.setY(0);

    this.stationList.setFont(FontSize.large);
    this.stationList.setOnClick((item: OverviewRow) => {
      this.selectedItemID = item.object.id;
      this.selectedItemType = item.type;
    });

    this.tabs.addComponent(this.stationList, 'Stations');

    // jumphole list
    this.jumpholeList.setWidth(this.getWidth());
    this.jumpholeList.setHeight(this.getHeight());
    this.jumpholeList.initialize();

    this.jumpholeList.setX(0);
    this.jumpholeList.setY(0);

    this.jumpholeList.setFont(FontSize.large);
    this.jumpholeList.setOnClick((item: OverviewRow) => {
      this.selectedItemID = item.object.id;
      this.selectedItemType = item.type;
    });

    this.tabs.addComponent(this.jumpholeList, 'Jumpholes');
  }

  periodicUpdate() {
    //
  }

  sync(player: Player) {
    const objects: OverviewRow[] = [];
    const ships: OverviewRow[] = [];
    const stations: OverviewRow[] = [];
    const jumpholes: OverviewRow[] = [];

    // sort entities
    const sortedJumpholes = player.currentSystem.jumpholes.sort((a, b) =>
      (a.jumpholeName ?? '').localeCompare(b.jumpholeName ?? ''),
    );

    const sortedAsteroids = player.currentSystem.asteroids.sort((a, b) =>
      (a.name ?? '').localeCompare(b.name ?? ''),
    );

    const sortedPlanets = player.currentSystem.planets.sort((a, b) =>
      (a.planetName ?? '').localeCompare(b.planetName ?? ''),
    );

    const sortedStations = player.currentSystem.stations.sort((a, b) =>
      (a.stationName ?? '').localeCompare(b.stationName ?? ''),
    );

    const sortedOutposts = player.currentSystem.outposts.sort((a, b) =>
    (a.outpostName ?? '').localeCompare(b.outpostName ?? ''),
  );

    const sortedShips = player.currentSystem.ships.sort((a, b) =>
      compareShipString(a).localeCompare(compareShipString(b)),
    );

    const sortedWrecks = player.currentSystem.wrecks.sort((a, b) =>
      (a.wreckName ?? '').localeCompare(b.wreckName ?? ''),
    );

    // include stars
    for (const i of player.currentSystem.stars) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y,
      );

      objects.push({
        object: i,
        type: TargetType.Star,
        listString: () => {
          return `${fixedString('STAR', 9)}${fixedString(
            player.currentSystem.systemName,
            32,
          )} ${fixedString(od, 8)}`;
        },
        listColor: () => GDIStyle.listTextColor,
      });
    }

    // include planets
    for (const i of sortedPlanets) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y,
      );

      objects.push({
        object: i,
        type: TargetType.Planet,
        listString: () => {
          return `${fixedString('PLANET', 9)}${fixedString(
            i.planetName,
            32,
          )} ${fixedString(od, 8)}`;
        },
        listColor: () => GDIStyle.listTextColor,
      });
    }

    // include stations
    for (const i of sortedStations) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y,
      );

      const d: OverviewRow = {
        object: i,
        type: TargetType.Station,
        listString: () => {
          const faction = i.getFaction();

          return `${fixedString('STATION', 9)}${fixedString(
            '[' + faction?.ticker + ']',
            6,
          )}${fixedString(i.stationName, 26)} ${fixedString(od, 8)}`;
        },
        listColor: () => i.getStandingColor(),
      };

      objects.push(d);
      stations.push(d);
    }

    // include outposts
    for (const i of sortedOutposts) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y,
      );

      const d: OverviewRow = {
        object: i,
        type: TargetType.Outpost,
        listString: () => {
          const faction = i.getFaction();

          return `${fixedString('OUTPOST', 9)}${fixedString(
            '[' + faction?.ticker + ']',
            6,
          )}${fixedString(i.outpostName, 26)} ${fixedString(od, 8)}`;
        },
        listColor: () => i.getStandingColor(),
      };

      objects.push(d);
      stations.push(d);
    }

    // include ships
    for (const i of sortedShips) {
      // skip if player ship
      if (i.id === player.currentShip.id) {
        continue;
      }

      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y,
      );

      const d: OverviewRow = {
        object: i,
        type: TargetType.Ship,
        listString: () => {
          const faction = i.getFaction();

          return `${fixedString('SHIP', 9)}${fixedString(
            '[' + faction?.ticker + ']',
            6,
          )}${fixedString(i.ownerName, 18)} ${fixedString(
            i.texture,
            7,
          )} ${fixedString(od, 8)}`;
        },
        listColor: () => i.getStandingColor(),
      };

      objects.push(d);
      ships.push(d);
    }

    // include jumpholes
    for (const i of sortedJumpholes) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y,
      );

      const d: OverviewRow = {
        object: i,
        type: TargetType.Jumphole,
        listString: () => {
          return `${fixedString('JUMPHOLE', 9)}${fixedString(
            i.jumpholeName,
            32,
          )} ${fixedString(od, 8)}`;
        },
        listColor: () => GDIStyle.listTextColor,
      };

      objects.push(d);
      jumpholes.push(d);
    }

    // include asteroids
    for (const i of sortedAsteroids) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y,
      );

      const d = {
        object: i,
        type: TargetType.Asteroid,
        listString: () => {
          return `${fixedString('ASTEROID', 9)}${fixedString(
            i.name,
            32,
          )} ${fixedString(od, 8)}`;
        },
        listColor: () => GDIStyle.listTextColor,
      };

      objects.push(d);
    }

    // include wrecks
    for (const i of sortedWrecks) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y,
      );

      const d = {
        object: i,
        type: TargetType.Wreck,
        listString: () => {
          return `${fixedString('WRECK', 9)}${fixedString(
            i.wreckName,
            32,
          )} ${fixedString(od, 8)}`;
        },
        listColor: () => GDIStyle.listTextColor,
      };

      objects.push(d);
    }

    // spacers at end
    const spacers: OverviewRow[] = [];
    for (let i = 0; i < 2; i++) {
      const spacerObject = new OverviewRow();
      spacerObject.object = {};
      spacerObject.listString = () => '';
      spacerObject.type = TargetType.Nothing;

      spacers.push(spacerObject);
    }

    // store on lists
    this.globalList.setItems([...objects, ...spacers]);
    this.shipList.setItems([...ships, ...spacers]);
    this.stationList.setItems([...stations, ...spacers]);
    this.jumpholeList.setItems([...jumpholes, ...spacers]);

    // reselect row in global list
    const gItems = this.globalList.getItems() as OverviewRow[];

    for (let x = 0; x < gItems.length; x++) {
      const i = gItems[x].object;

      if (i.id === this.selectedItemID) {
        this.globalList.setSelectedIndex(x);
        break;
      }
    }

    // reselect row in ship list
    const shItems = this.shipList.getItems() as OverviewRow[];

    for (let x = 0; x < shItems.length; x++) {
      const i = shItems[x].object;

      if (i.id === this.selectedItemID) {
        this.shipList.setSelectedIndex(x);
        break;
      }
    }

    // reselect row in station list
    const stItems = this.stationList.getItems() as OverviewRow[];

    for (let x = 0; x < stItems.length; x++) {
      const i = stItems[x].object;

      if (i.id === this.selectedItemID) {
        this.stationList.setSelectedIndex(x);
        break;
      }
    }

    // reselect row in jumphole list
    const jhItems = this.jumpholeList.getItems() as OverviewRow[];

    for (let x = 0; x < jhItems.length; x++) {
      const i = jhItems[x].object;

      if (i.id === this.selectedItemID) {
        this.jumpholeList.setSelectedIndex(x);
        break;
      }
    }

    // set target on ship
    player.currentTargetID = this.selectedItemID;
    player.currentTargetType = this.selectedItemType;
  }
}

class OverviewRow {
  object: any;
  type: TargetType;
  listString: () => string;
  listColor: () => string;
}

function overviewDistance(
  px: number,
  py: number,
  x: number,
  y: number,
): string {
  // get distance
  const d = Math.round(Math.sqrt((px - x) * (px - x) + (py - y) * (py - y)));
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

function fixedString(str: string, width: number): string {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str?.substr(0, width)?.padEnd(width);
}

function compareShipString(s: Ship): string {
  return `${s.faction?.ticker}||::||${s.texture}||>>||${s.ownerName}`;
}
