import { GDIWindow } from '../base/gdiWindow';
import { FontSize } from '../base/gdiStyle';
import { GDIBar } from '../components/gdiBar';
import { Ship } from '../../engineModels/ship';

export class ShipStatusWindow extends GDIWindow {
    shieldBar: GDIBar = new GDIBar();
    armorBar: GDIBar = new GDIBar();
    hullBar: GDIBar = new GDIBar();
    heatBar: GDIBar = new GDIBar();
    energyBar: GDIBar = new GDIBar();
    fuelBar: GDIBar = new GDIBar();

    ship: Ship;

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

        // setup armor bar
        this.armorBar.setWidth(600);
        this.armorBar.setHeight(10);
        this.armorBar.initialize();

        this.armorBar.setX(0);
        this.armorBar.setY(10);
        this.armorBar.setPercentage(0);

        this.armorBar.setFont(FontSize.small);
        this.armorBar.setText('Armor');

        // setup hull bar
        this.hullBar.setWidth(600);
        this.hullBar.setHeight(10);
        this.hullBar.initialize();

        this.hullBar.setX(0);
        this.hullBar.setY(20);
        this.hullBar.setPercentage(0);

        this.hullBar.setFont(FontSize.small);
        this.hullBar.setText('Hull');

        // setup energy bar
        this.energyBar.setWidth(600);
        this.energyBar.setHeight(10);
        this.energyBar.initialize();

        this.energyBar.setX(0);
        this.energyBar.setY(30);
        this.energyBar.setPercentage(0);

        this.energyBar.setFont(FontSize.small);
        this.energyBar.setText('Energy');

        // setup heat bar
        this.heatBar.setWidth(600);
        this.heatBar.setHeight(10);
        this.heatBar.initialize();

        this.heatBar.setX(0);
        this.heatBar.setY(40);
        this.heatBar.setPercentage(0);

        this.heatBar.setFont(FontSize.small);
        this.heatBar.setText('Heat');

        // setup fuel bar
        this.fuelBar.setWidth(600);
        this.fuelBar.setHeight(10);
        this.fuelBar.initialize();

        this.fuelBar.setX(0);
        this.fuelBar.setY(50);
        this.fuelBar.setPercentage(0);

        this.fuelBar.setFont(FontSize.small);
        this.fuelBar.setText('Fuel');

        // add components
        this.addComponent(this.shieldBar);
        this.addComponent(this.armorBar);
        this.addComponent(this.hullBar);
        this.addComponent(this.energyBar);
        this.addComponent(this.heatBar);
        this.addComponent(this.fuelBar);
    }

    periodicUpdate() {
        if (this.ship !== undefined && this.ship !== null) {
            this.shieldBar.setPercentage(this.ship.shieldP);
            this.armorBar.setPercentage(this.ship.armorP);
            this.hullBar.setPercentage(this.ship.hullP);
            this.energyBar.setPercentage(this.ship.energyP);
            this.heatBar.setPercentage(this.ship.heatP);
            this.fuelBar.setPercentage(this.ship.fuelP);
        }
    }

    setShip(ship: Ship) {
        this.ship = ship;
    }
}
