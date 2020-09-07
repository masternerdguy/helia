import { GDIWindow } from '../base/gdiWindow';

export class OverviewWindow extends GDIWindow {
    initialize() {
        // set dimensions
        this.setWidth(300);
        this.setHeight(this.getHeight());

        // initialize
        super.initialize();
    }

    pack() {
        this.setTitle('System Overview');
    }

    periodicUpdate() {

    }
}
