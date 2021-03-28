import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIWindow } from '../base/gdiWindow';
import { GDIList } from '../components/gdiList';

export class PushErrorWindow extends GDIWindow {
  private textList = new GDIList();

  initialize() {
    // set dimensions
    this.setWidth(600);
    this.setHeight(100);

    // initialize
    super.initialize();
  }

  pack() {
    this.setTitle('Push Error');
    // text list
    this.textList.setWidth(this.getWidth());
    this.textList.setHeight(this.getHeight());
    this.textList.initialize();

    this.textList.setFont(FontSize.giant);
    this.textList.setX(0);
    this.textList.setY(0);
    this.textList.setOverrideFillColor(GDIStyle.errorFillColor);

    this.addComponent(this.textList);
  }

  setText(text: string) {
    this.textList.setItemsFromText(text);
  }

  periodicUpdate() {
    // 
  }
}
