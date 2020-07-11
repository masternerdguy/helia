import { GDIWindow } from '../base/gdiWindow';
import { GDIButton, FontSize } from '../components/gdiButton';

export class TestWindow extends GDIWindow {
    initialize() {
        // set dimensions
        this.setWidth(400);
        this.setHeight(400);

        // initialize
        super.initialize();
    }

    pack() {
        // test button
        const testBtn = new GDIButton();
        testBtn.setWidth(140);
        testBtn.setHeight(20);
        testBtn.initialize();

        testBtn.setX(10);
        testBtn.setY(10);

        testBtn.setFont(FontSize.normal);
        testBtn.setText('Hello test button!');

        testBtn.setOnClick((x, y) => {
            console.log('click! ' + (x - testBtn.getX()) + ' ' + (y - testBtn.getY()));
        });

        this.addComponent(testBtn);
    }
}
