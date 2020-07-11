import { GDIWindow } from '../base/gdiWindow';

export class TestWindow extends GDIWindow {
    initialize() {
        this.setWidth(400);
        this.setHeight(400);

        super.initialize();
    }
}
