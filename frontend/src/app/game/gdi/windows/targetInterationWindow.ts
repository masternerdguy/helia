import { GDIWindow } from '../base/gdiWindow';

export class TargetInteractionWindow extends GDIWindow {
    initialize() {
        // set dimensions
        this.setWidth(200);
        this.setHeight(30);

        // initialize
        super.initialize();
    }

    pack() {
        this.setTitle('Interaction');
    }

    periodicUpdate() {
        //
    }
}
