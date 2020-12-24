import { GDIWindow } from '../base/gdiWindow';
import { GDIButton } from '../components/gdiButton';
import { FontSize } from '../base/gdiStyle';
import { GDILabel } from '../components/gdiLabel';
import { GDIBar } from '../components/gdiBar';
import { GDIList } from '../components/gdiList';
import { GDIInput } from '../components/gdiInput';

export class TestWindow extends GDIWindow {
  testBar2 = new GDIBar();

  initialize() {
    // set dimensions
    this.setWidth(400);
    this.setHeight(400);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Test Window Please Ignore');

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
      console.log(
        'click! ' + (x - testBtn.getX()) + ' ' + (y - testBtn.getY())
      );
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

    // test bar
    const testBar = new GDIBar();
    testBar.setWidth(200);
    testBar.setHeight(20);
    testBar.initialize();

    testBar.setX(10);
    testBar.setY(100);
    testBar.setPercentage(75);

    testBar.setFont(FontSize.small);
    testBar.setText('Hello test bar!');

    this.addComponent(testBar);

    // test bar
    this.testBar2 = new GDIBar();
    this.testBar2.setWidth(200);
    this.testBar2.setHeight(20);
    this.testBar2.initialize();

    this.testBar2.setX(10);
    this.testBar2.setY(140);
    this.testBar2.setPercentage(25);
    this.testBar2.setColor('green');

    this.addComponent(this.testBar2);

    // test list
    const testLst = new GDIList();
    testLst.setWidth(200);
    testLst.setHeight(100);
    testLst.initialize();

    testLst.setX(10);
    testLst.setY(200);

    testLst.setFont(FontSize.large);
    testLst.setItems([
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9,
      10,
      11,
      12,
      13,
      14,
      15,
      16,
      17,
      18,
      19,
      20,
    ]);
    testLst.setOnClick((item) => {
      console.log(item);
    });

    this.addComponent(testLst);

    // test input
    const testIpt = new GDIInput();
    testIpt.setWidth(200);
    testIpt.setHeight(20);
    testIpt.initialize();

    testIpt.setX(10);
    testIpt.setY(320);

    testIpt.setFont(FontSize.small);
    testIpt.setOnReturn((txt) => {
      console.log(txt);
    });

    this.addComponent(testIpt);
  }

  periodicUpdate() {
    this.testBar2.setPercentage(this.testBar2.getPercentage() + 1);

    if (this.testBar2.getPercentage() === 100) {
      this.testBar2.setPercentage(0);
    }
  }
}
