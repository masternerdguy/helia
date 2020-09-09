import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';
import { FontSize } from '../base/gdiStyle';
import { Player } from '../../engineModels/player';

export class OverviewWindow extends GDIWindow {
    objectList = new GDIList();

    initialize() {
        // set dimensions
        this.setWidth(300);
        this.setHeight(this.getHeight());

        // initialize
        super.initialize();
    }

    pack() {
        this.setTitle('System Overview');

        // object list
        this.objectList.setWidth(this.getWidth());
        this.objectList.setHeight(this.getHeight());
        this.objectList.initialize();

        this.objectList.setX(0);
        this.objectList.setY(0);

        this.objectList.setFont(FontSize.large);
        this.objectList.setItems([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20]);
        this.objectList.setOnClick((item) => {
            console.log(item);
        });

        this.addComponent(this.objectList);
    }

    periodicUpdate() {
        //
    }

    sync(player: Player) {
        const objects: OverviewRow[] = [];

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
            })
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
            })
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
            })
        }

        // include ships
        for (const i of player.currentSystem.ships) {
            objects.push({
                object: i,
                listString: () => {
                    return `Ship ${i.shipName} - ${overviewDistance(
                        player.currentShip.x, 
                        player.currentShip.y, 
                        i.x, 
                        i.y)}`
                }
            })
        }

        this.objectList.setItems(objects);
    }
}

class OverviewRow {
    object: any;
    listString: () => string;
}

function overviewDistance(px: number, py: number, x: number, y: number) {
    return Math.round(Math.sqrt((px - x) * (px - x) + (py - y) * (py - y)))
}