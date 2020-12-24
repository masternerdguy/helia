import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';
import { FontSize } from '../base/gdiStyle';
import { Player, TargetType } from '../../engineModels/player';
import { GDITabPane } from '../components/gdiTabPane';

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
    this.setWidth(310);
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

    // include stars
    for (const i of player.currentSystem.stars) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y
      );

      objects.push({
        object: i,
        type: TargetType.Star,
        listString: () => {
          return `${fixedString('STAR', 9)}${fixedString(
            player.currentSystem.systemName,
            24
          )} ${fixedString(od, 8)}`;
        },
      });
    }

    // include planets
    for (const i of player.currentSystem.planets) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y
      );

      objects.push({
        object: i,
        type: TargetType.Planet,
        listString: () => {
          return `${fixedString('PLANET', 9)}${fixedString(
            i.planetName,
            24
          )} ${fixedString(od, 8)}`;
        },
      });
    }

    // include stations
    for (const i of player.currentSystem.stations) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y
      );

      const d: OverviewRow = {
        object: i,
        type: TargetType.Station,
        listString: () => {
          return `${fixedString('STATION', 9)}${fixedString(
            i.stationName,
            24
          )} ${fixedString(od, 8)}`;
        },
      };

      objects.push(d);
      stations.push(d);
    }

    // include ships
    for (const i of player.currentSystem.ships) {
      // skip if player ship
      if (i.id === player.currentShip.id) {
        continue;
      }

      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y
      );

      const d: OverviewRow = {
        object: i,
        type: TargetType.Ship,
        listString: () => {
          return `${fixedString('SHIP', 9)}${fixedString(
            i.ownerName,
            16
          )} ${fixedString(i.texture, 7)} ${fixedString(od, 8)}`;
        },
      };

      objects.push(d);
      ships.push(d);
    }

    // include jumpholes
    for (const i of player.currentSystem.jumpholes) {
      const od = overviewDistance(
        player.currentShip.x,
        player.currentShip.y,
        i.x,
        i.y
      );

      const d: OverviewRow = {
        object: i,
        type: TargetType.Jumphole,
        listString: () => {
          return `${fixedString('JUMPHOLE', 9)}${fixedString(
            i.jumpholeName,
            24
          )} ${fixedString(od, 8)}`;
        },
      };

      objects.push(d);
      jumpholes.push(d);
    }

    // store on lists
    this.globalList.setItems(objects);
    this.shipList.setItems(ships);
    this.stationList.setItems(stations);
    this.jumpholeList.setItems(jumpholes);

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
}

function overviewDistance(
  px: number,
  py: number,
  x: number,
  y: number
): string {
  // get distance with unit appended
  const d = Math.round(Math.sqrt((px - x) * (px - x) + (py - y) * (py - y)));
  let o = `${d}`;

  // include metrix prefix if needed
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

function fixedString(str: string, width: number) {
  if (str === undefined || str == null) {
    return ''.padEnd(width);
  }

  return str?.substr(0, width)?.padEnd(width);
}
