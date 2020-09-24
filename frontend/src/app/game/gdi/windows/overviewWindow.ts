import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';
import { FontSize } from '../base/gdiStyle';
import { Player } from '../../engineModels/player';
import { GDITabList } from '../components/gdiTabList';

export class OverviewWindow extends GDIWindow {
    tabs = new GDITabList();
    globalList = new GDIList();
    shipList = new GDIList();

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
        this.tabs.setSelectedTab("Global");

        this.addComponent(this.tabs);

        // all object list
        this.globalList.setWidth(this.getWidth());
        this.globalList.setHeight(this.getHeight());
        this.globalList.initialize();

        this.globalList.setX(0);
        this.globalList.setY(0);

        this.globalList.setFont(FontSize.large);
        this.globalList.setItems([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]);
        this.globalList.setOnClick((item) => {
            console.log(item);
        });

        this.tabs.addComponent(this.globalList, "Global");

        // ship list
        this.shipList.setWidth(this.getWidth());
        this.shipList.setHeight(this.getHeight());
        this.shipList.initialize();

        this.shipList.setX(0);
        this.shipList.setY(0);

        this.shipList.setFont(FontSize.large);
        this.shipList.setItems(['a','s','d','f','l','o','l']);
        this.shipList.setOnClick((item) => {
            console.log(item);
        });

        this.tabs.addComponent(this.shipList, "Ships");
    }

    periodicUpdate() {
        //
    }

    sync(player: Player) {
        const objects: OverviewRow[] = [];
        const ships: OverviewRow[] = [];

        // include stars
        for (const i of player.currentSystem.stars) {
            objects.push({
                object: i,
                listString: () => {
                    return `${player.currentSystem.systemName} Star - ${overviewDistance(
                        player.currentShip.x, 
                        player.currentShip.y, 
                        i.x, 
                        i.y)}`
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
                        i.y)}`
                }
            });
        }

        // include stations
        for (const i of player.currentSystem.stations) {
            objects.push({
                object: i,
                listString: () => {
                    return `Station ${i.stationName} - ${overviewDistance(
                        player.currentShip.x, 
                        player.currentShip.y, 
                        i.x, 
                        i.y)}`
                }
            });
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
                        i.y)}`
                }
            };

            objects.push(d);
            ships.push(d);
        }

        this.globalList.setItems(objects);
        this.shipList.setItems(ships);
    }
}

class OverviewRow {
    object: any;
    listString: () => string;
}

function overviewDistance(px: number, py: number, x: number, y: number) {
    return Math.round(Math.sqrt((px - x) * (px - x) + (py - y) * (py - y)))
}