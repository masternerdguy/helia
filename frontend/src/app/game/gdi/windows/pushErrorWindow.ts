import { GDIWindow } from '../base/gdiWindow';
import { GDIButton } from '../components/gdiButton';
import { FontSize } from '../base/gdiStyle';
import { GDILabel } from '../components/gdiLabel';

export class PushErrorWindow extends GDIWindow {
  private testBtn = new GDIButton();
  private testLbl = new GDILabel();

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Push Error');

    // test button
    this.testBtn.setWidth(140);
    this.testBtn.setHeight(20);
    this.testBtn.initialize();

    this.testBtn.setX(10);
    this.testBtn.setY(50);

    this.testBtn.setFont(FontSize.normal);
    this.testBtn.setText('Hello test button!');

    this.testBtn.setOnClick((x, y) => {
      console.log(
        'click! ' + (x - this.testBtn.getX()) + ' ' + (y - this.testBtn.getY())
      );
    });

    this.addComponent(this.testBtn);

    // test label
    this.testLbl.setWidth(200);
    this.testLbl.setHeight(20);
    this.testLbl.initialize();

    this.testLbl.setX(10);
    this.testLbl.setY(10);

    this.testLbl.setFont(FontSize.large);
    this.testLbl.setText('Hello test label!');

    this.addComponent(this.testLbl);
  }

  periodicUpdate() {
    //
  }
}
