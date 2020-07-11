import { GDIWindow } from '../base/gdiWindow';
import { GDIButton } from '../components/gdiButton';
import { FontSize } from '../base/gdiStyle';
import { GDILabel } from '../components/gdiLabel';

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
        testBtn.setY(50);

        testBtn.setFont(FontSize.normal);
        testBtn.setText('Hello test button!');

        testBtn.setOnClick((x, y) => {
            console.log('click! ' + (x - testBtn.getX()) + ' ' + (y - testBtn.getY()));
        });

        this.addComponent(testBtn);

        // test label
        const testLbl = new GDILabel();
        testLbl.setWidth(200);
        testLbl.setHeight(20);
        testLbl.initialize();

        testLbl.setX(10);
        testLbl.setY(10);

        testLbl.setFont(FontSize.large);
        testLbl.setText('Hello test label!');

        this.addComponent(testLbl);
    }
}
