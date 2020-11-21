import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIBar } from '../components/gdiBar';
import { Ship } from '../../engineModels/ship';
import { GDIList } from '../components/gdiList';
import { WSModule } from '../../wsModels/entities/wsShip';
import { WsService } from '../../ws.service';
import { ClientActivateModule } from '../../wsModels/bodies/activateModule';
import { MessageTypes } from '../../wsModels/gameMessage';
import { Player } from '../../engineModels/player';
import { ClientDeactivateModule } from '../../wsModels/bodies/clientDeactivateModule';

export class ShipStatusWindow extends GDIWindow {
    // status bars
    private shieldBar: GDIBar = new GDIBar();
    private armorBar: GDIBar = new GDIBar();
    private hullBar: GDIBar = new GDIBar();
    private heatBar: GDIBar = new GDIBar();
    private energyBar: GDIBar = new GDIBar();
    private fuelBar: GDIBar = new GDIBar();

    // racks
    private rackA: GDIList = new GDIList();
    private rackB: GDIList = new GDIList();
    private rackC: GDIList = new GDIList();

    // ship being monitored
    private ship: Ship;
    private player: Player;

    // ws
    private wsSvc: WsService;

    initialize() {
        // set dimensions
        this.setWidth(600);
        this.setHeight(260);

        // initialize
        super.initialize();
    }

    pack() {
        this.setTitle('Ship Status');

        // setup shield bar
        this.shieldBar.setWidth(600);
        this.shieldBar.setHeight(10);
        this.shieldBar.initialize();

        this.shieldBar.setX(0);
        this.shieldBar.setY(0);
        this.shieldBar.setPercentage(0);

        this.shieldBar.setFont(FontSize.small);
        this.shieldBar.setText('Shield');
        this.shieldBar.setShowPercentage(true);

        // setup armor bar
        this.armorBar.setWidth(600);
        this.armorBar.setHeight(10);
        this.armorBar.initialize();

        this.armorBar.setX(0);
        this.armorBar.setY(10);
        this.armorBar.setPercentage(0);

        this.armorBar.setFont(FontSize.small);
        this.armorBar.setText('Armor');
        this.armorBar.setShowPercentage(true);

        // setup hull bar
        this.hullBar.setWidth(600);
        this.hullBar.setHeight(10);
        this.hullBar.initialize();

        this.hullBar.setX(0);
        this.hullBar.setY(20);
        this.hullBar.setPercentage(0);

        this.hullBar.setFont(FontSize.small);
        this.hullBar.setText('Hull');
        this.hullBar.setShowPercentage(true);

        // setup energy bar
        this.energyBar.setWidth(600);
        this.energyBar.setHeight(10);
        this.energyBar.initialize();

        this.energyBar.setX(0);
        this.energyBar.setY(30);
        this.energyBar.setPercentage(0);

        this.energyBar.setFont(FontSize.small);
        this.energyBar.setText('Energy');
        this.energyBar.setShowPercentage(true);

        // setup heat bar
        this.heatBar.setWidth(600);
        this.heatBar.setHeight(10);
        this.heatBar.initialize();

        this.heatBar.setX(0);
        this.heatBar.setY(40);
        this.heatBar.setPercentage(0);

        this.heatBar.setFont(FontSize.small);
        this.heatBar.setText('Heat');
        this.heatBar.setShowPercentage(true);
        this.heatBar.setAllowOverflow(true);

        // setup fuel bar
        this.fuelBar.setWidth(600);
        this.fuelBar.setHeight(10);
        this.fuelBar.initialize();

        this.fuelBar.setX(0);
        this.fuelBar.setY(50);
        this.fuelBar.setPercentage(0);

        this.fuelBar.setFont(FontSize.small);
        this.fuelBar.setText('Fuel');
        this.fuelBar.setShowPercentage(true);

        // setup rack a
        this.rackA.setWidth(200);
        this.rackA.setHeight(200);
        this.rackA.initialize();

        this.rackA.setX(0);
        this.rackA.setY(60);

        this.rackA.setFont(FontSize.smallNormal);
        this.rackA.setOnClick((module: RackRow) => {
            if (!module.object.willRepeat) {
                // issue order to activate module
                this.activateModule(module, 'A');
            } else {
                // issue order to deactivate module
                this.deactivateModule(module, 'A');
            }
        });

        // setup rack b
        this.rackB.setWidth(200);
        this.rackB.setHeight(200);
        this.rackB.initialize();

        this.rackB.setX(200);
        this.rackB.setY(60);

        this.rackB.setFont(FontSize.small);
        this.rackB.setOnClick((module: RackRow) => {
            if (!module.object.willRepeat) {
                // issue order to activate module
                this.activateModule(module, 'B');
            } else {
                // issue order to deactivate module
                this.deactivateModule(module, 'B');
            }
        });

        // setup rack c
        this.rackC.setWidth(200);
        this.rackC.setHeight(200);
        this.rackC.initialize();

        this.rackC.setX(400);
        this.rackC.setY(60);

        this.rackC.setFont(FontSize.small);
        this.rackC.setOnClick((module: RackRow) => {
            if (!module.object.willRepeat) {
                // issue order to activate module
                this.activateModule(module, 'C');
            } else {
                // issue order to deactivate module
                this.deactivateModule(module, 'C');
            }
        });

        // add components
        this.addComponent(this.shieldBar);
        this.addComponent(this.armorBar);
        this.addComponent(this.hullBar);
        this.addComponent(this.energyBar);
        this.addComponent(this.heatBar);
        this.addComponent(this.fuelBar);
        this.addComponent(this.rackA);
        this.addComponent(this.rackB);
        this.addComponent(this.rackC);
    }

    private activateModule(module: RackRow, rack: string) {
        const b = new ClientActivateModule();
        b.rack = rack;

        b.sid = this.wsSvc.sid;
        b.itemID = module.object.itemID;
        b.TargetID = this.player.currentTargetID;
        b.TargetType = this.player.currentTargetType;

        this.wsSvc.sendMessage(MessageTypes.ActivateModule, b);
    }

    private deactivateModule(module: RackRow, rack: string) {
        const b = new ClientDeactivateModule();
        b.rack = rack;

        b.sid = this.wsSvc.sid;
        b.itemID = module.object.itemID;

        this.wsSvc.sendMessage(MessageTypes.DeactivateModule, b);
    }

    periodicUpdate() {
        if (this.ship !== undefined && this.ship !== null) {
            // update status bars
            this.shieldBar.setPercentage(this.ship.shieldP);
            this.armorBar.setPercentage(this.ship.armorP);
            this.hullBar.setPercentage(this.ship.hullP);
            this.energyBar.setPercentage(this.ship.energyP);
            this.heatBar.setPercentage(this.ship.heatP);
            this.fuelBar.setPercentage(this.ship.fuelP);

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

                // save on lists
                this.rackA.setItems(rackAMods);
                this.rackB.setItems(rackBMods);
                this.rackC.setItems(rackCMods);
            }
        }
    }

    private moduleStatusString(m: WSModule) {
        const pc = fixedString(m.isCycling ? `${m.cyclePercent}%` : '', 4);

        // build status string
        return `${fixedString(m.willRepeat ? '*' : '', 1)} ${fixedString(m.type, 24)} ${pc}~`;
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
