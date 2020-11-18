import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';
import { FontSize } from '../base/gdiStyle';
import { Player } from '../../engineModels/player';
import { GDITabPane } from '../components/gdiTabPane';

export class OverviewWindow extends GDIWindow {
    tabs = new GDITabPane();
    globalList = new GDIList();
    shipList = new GDIList();
    stationList = new GDIList();
    selectedItemID: string;

    initialize() {
        // set dimensions
        this.setWidth(300);
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
        this.globalList.setOnClick((item) => {
            this.selectedItemID = item.object.id;
        });

        this.tabs.addComponent(this.globalList, 'Global');

        // ship list
        this.shipList.setWidth(this.getWidth());
        this.shipList.setHeight(this.getHeight());
        this.shipList.initialize();

        this.shipList.setX(0);
        this.shipList.setY(0);

        this.shipList.setFont(FontSize.large);
        this.shipList.setOnClick((item) => {
            this.selectedItemID = item.object.id;
        });

        this.tabs.addComponent(this.shipList, 'Ships');

        // station list
        this.stationList.setWidth(this.getWidth());
        this.stationList.setHeight(this.getHeight());
        this.stationList.initialize();

        this.stationList.setX(0);
        this.stationList.setY(0);

        this.stationList.setFont(FontSize.large);
        this.stationList.setOnClick((item) => {
            this.selectedItemID = item.object.id;
        });

        this.tabs.addComponent(this.stationList, 'Stations');
    }

    periodicUpdate() {
        //
    }

    sync(player: Player) {
        const objects: OverviewRow[] = [];
        const ships: OverviewRow[] = [];
        const stations: OverviewRow[] = [];

        // include stars
        for (const i of player.currentSystem.stars) {
            objects.push({
                object: i,
                listString: () => {
                    return `${player.currentSystem.systemName} Star - ${overviewDistance(
                        player.currentShip.x,
                        player.currentShip.y,
                        i.x,
                        i.y)}`;
                }
            });
        }

        // include planets
        for (const i of player.currentSystem.planets) {
            objects.push({
                object: i,
                listString: () => {
                    return `Planet ${i.planetName} - ${overviewDistance(
                        player.currentShip.x,
                        player.currentShip.y,
                        i.x,
                        i.y)}`;
                }
            });
        }

        // include stations
        for (const i of player.currentSystem.stations) {
            const d: OverviewRow = {
                object: i,
                listString: () => {
                    return `Station ${i.stationName} - ${overviewDistance(
                        player.currentShip.x,
                        player.currentShip.y,
                        i.x,
                        i.y)}`;
                }
            };

            objects.push(d);
            stations.push(d);
        }

        // include ships
        for (const i of player.currentSystem.ships) {
            const d: OverviewRow = {
                object: i,
                listString: () => {
                    return `Ship ${i.shipName} - ${overviewDistance(
                        player.currentShip.x,
                        player.currentShip.y,
                        i.x,
                        i.y)}`;
                }
            };

            objects.push(d);
            ships.push(d);
        }

        // store on lists
        this.globalList.setItems(objects);
        this.shipList.setItems(ships);
        this.stationList.setItems(stations);

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
    }
}

class OverviewRow {
    object: any;
    listString: () => string;
}

function overviewDistance(px: number, py: number, x: number, y: number) {
    return Math.round(Math.sqrt((px - x) * (px - x) + (py - y) * (py - y)));
}
