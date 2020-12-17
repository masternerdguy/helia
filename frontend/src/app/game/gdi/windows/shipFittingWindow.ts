import { GDIWindow } from '../base/gdiWindow';
import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIBar } from '../components/gdiBar';
import { Ship } from '../../engineModels/ship';
import { GDIList } from '../components/gdiList';
import { WSModule } from '../../wsModels/entities/wsShip';
import { WsService } from '../../ws.service';
import { ClientActivateModule } from '../../wsModels/bodies/activateModule';
import { MessageTypes } from '../../wsModels/gameMessage';
import { Player } from '../../engineModels/player';
import { ClientDeactivateModule } from '../../wsModels/bodies/clientDeactivateModule';

export class ShipFittingWindow extends GDIWindow {
    // lists
    private shipView: GDIList = new GDIList();
    private infoView: GDIList = new GDIList();
    private actionView: GDIList = new GDIList();

    // ship being fit
    private ship: Ship;
    private player: Player;

    // ws
    private wsSvc: WsService;

    initialize() {
        // set dimensions
        this.setWidth(600);
        this.setHeight(600);

        // initialize
        super.initialize();
    }

    pack() {
        this.setTitle('Ship Fitting');


    }

    periodicUpdate() {
        if (this.ship !== undefined && this.ship !== null) {
            if (this.ship.fitStatus !== undefined && this.ship.fitStatus !== null) {
                // update fitted module display
                const rackAMods: RackRow[] = [];
                const rackBMods: RackRow[] = [];
                const rackCMods: RackRow[] = [];

                // build entries for modules on racks
                for (const m of this.ship.fitStatus.aRack.modules) {
                    const d: RackRow = {
                        object: m,
                        listString: () => {
                            return this.moduleStatusString(m);
                        }
                    };

                    rackAMods.push(d);
                }

                for (const m of this.ship.fitStatus.bRack.modules) {
                    const d: RackRow = {
                        object: m,
                        listString: () => {
                            return this.moduleStatusString(m);
                        }
                    };

                    rackBMods.push(d);
                }

                for (const m of this.ship.fitStatus.cRack.modules) {
                    const d: RackRow = {
                        object: m,
                        listString: () => {
                            return this.moduleStatusString(m);
                        }
                    };

                    rackCMods.push(d);
                }
            }
        }
    }

    private moduleStatusString(m: WSModule) {
        const pc = fixedString(m.isCycling ? `${m.cyclePercent}%` : '', 4);

        // build status string
        return `${fixedString(m.willRepeat ? '*' : '', 1)} ${fixedString(m.type, 24)} ${pc}`;
    }

    setShip(ship: Ship) {
        this.ship = ship;
    }

    setWsService(wsSvc: WsService) {
        this.wsSvc = wsSvc;
    }

    setPlayer(player: Player) {
        this.player = player;
    }
}

class RackRow {
    object: WSModule;
    listString: () => string;
}

function fixedString(str: string, width: number) {
    if (str === undefined || str == null) {
        return ''.padEnd(width);
    }

    return str.substr(0, width).padEnd(width);
}
