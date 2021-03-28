import { GDIBase } from '../base/gdiBase';
import { GDIComponent } from '../base/gdiComponent';
import { GDIRectangle } from '../base/gdiRectangle';
import { GDIStyle, FontSize } from '../base/gdiStyle';

export class GDITabPane extends GDIBase {
  private tabs: Map<string, Tab>;
  private selectedTab: string;

  private canvas: OffscreenCanvas;
  private ctx: any;

  private font: FontSize = FontSize.normal;

  initialize() {
    // initialize offscreen canvas
    this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight());
    this.ctx = this.canvas.getContext('2d');
    this.tabs = new Map();
  }

  periodicUpdate() {
    // 
  }

  render(): ImageBitmap {
    // render background
    this.ctx.fillStyle = GDIStyle.tabFillColor;
    this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

    // style text
    this.ctx.textAlign = 'middle';
    this.ctx.textBaseline = 'middle';

    this.ctx.fillStyle = GDIStyle.tabTextColor;
    this.ctx.strokeStyle = GDIStyle.tabTextColor;
    this.ctx.font = GDIStyle.getUnderlyingFont(this.getFont());

    // render tab handles
    const handleWidth = this.getWidth() / Math.max(this.tabs.size, 1);
    let handleI = 0;

    for (const tv of this.tabs) {
      if (tv[0] !== this.selectedTab) {
        this.ctx.fillStyle = GDIStyle.tabTextColor;
        this.ctx.strokeStyle = GDIStyle.tabTextColor;
      } else {
        this.ctx.fillStyle = GDIStyle.tabActiveTextColor;
        this.ctx.strokeStyle = GDIStyle.tabActiveTextColor;
      }

      this.ctx.fillText(
        tv[0],
        handleWidth * handleI + 2,
        GDIStyle.tabHandleHeight / 2
      );
      handleI++;
    }

    // get selected tab
    const t = this.tabs.get(this.selectedTab);

    if (t) {
      // render selected tab
      const b = t.render();
      this.ctx.drawImage(b, t.getX(), t.getY());
    }

    if (GDIStyle.tabBorderSize > 0) {
      // render border
      this.ctx.lineWidth = GDIStyle.tabBorderSize;
      this.ctx.strokeStyle = GDIStyle.tabBorderColor;
      this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
    }

    // convert to image and return
    return this.canvas.transferToImageBitmap();
  }

  handleClick(x: number, y: number) {
    // make sure this is a relevant click
    if (!this.containsPoint(x, y)) {
      return;
    }

    // check for click in handle
    const hY = y - this.getY();
    const hX = x - this.getX();

    if (hY < GDIStyle.tabHandleHeight) {
      // change tab selection
      const handleWidth = this.getWidth() / Math.max(this.tabs.size, 1);

      for (let i = 0; i < this.tabs.size; i++) {
        const rect = new GDIRectangle(
          handleWidth * i,
          0,
          handleWidth,
          GDIStyle.tabHandleHeight
        );
        if (rect.containsPoint(hX, hY)) {
          let u = 0;

          for (const tu of this.tabs) {
            if (u === i) {
              this.setSelectedTab(tu[0]);
              break;
            } else {
              u++;
            }
          }
        }
      }

      return;
    }

    // adjust input to be relative to tab
    const rX = hX;
    const rY = hY - GDIStyle.tabHandleHeight;

    // send event to selected tab
    const t = this.tabs.get(this.selectedTab);

    if (t) {
      t.handleClick(rX, rY);
    }
  }

  setFont(font: FontSize) {
    this.font = font;
  }

  getFont(): FontSize {
    return this.font;
  }

  getSelectedTab(): string {
    return this.selectedTab;
  }

  setSelectedTab(tab: string) {
    this.selectedTab = tab;
  }

  addComponent(component: GDIComponent, tab: string) {
    // get tab from map
    let t = this.tabs.get(tab);

    if (!t) {
      // initialize missing tab
      t = new Tab();
      t.setWidth(this.getWidth());
      t.setHeight(this.getHeight() - GDIStyle.tabHandleHeight);
      t.setX(0);
      t.setY(GDIStyle.tabHandleHeight);
      t.initialize();

      this.tabs.set(tab, t);
    }

    // add component to tab
    t.components.push(component);
  }
}

class Tab extends GDIBase {
  public components: GDIComponent[] = [];
  private canvas: OffscreenCanvas;
  private ctx: any;

  handleClick(x: number, y: number) {
    // make sure this is a relevant click
    if (!this.containsPoint(x, y + GDIStyle.tabHandleHeight)) {
      return;
    }

    // find the component that is being clicked on within this tab
    for (const c of this.components) {
      if (c.containsPoint(x, y)) {
        // send click event
        c.handleClick(x, y);
        break;
      }
    }
  }

  handleScroll(x: number, y: number, d: number) {
    // make sure this is a relevant scroll
    if (!this.containsPoint(x, y)) {
      return;
    }

    // find the component that is being scrolled on within this tab
    for (const c of this.components) {
      if (c.containsPoint(x, y)) {
        // send scroll event
        c.handleScroll(x, y, d);
        break;
      }
    }
  }

  handleKeyDown(x: number, y: number, key: string) {
    // make sure this is a relevant key press
    if (!this.containsPoint(x, y)) {
      return;
    }

    // find the component that is being typed on within this tab
    for (const c of this.components) {
      if (c.containsPoint(x, y)) {
        // send key event
        c.handleKeyDown(x, y, key);
        break;
      }
    }
  }

  initialize() {
    // initialize offscreen canvas
    this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight());
    this.ctx = this.canvas.getContext('2d');

    // initialize empty arrays
    this.components = [];
  }

  periodicUpdate() {
    // update components
    for (const c of this.components) {
      c.periodicUpdate();
    }
  }

  render(): ImageBitmap {
    // render background
    this.ctx.fillStyle = GDIStyle.tabFillColor;
    this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

    // render components
    for (const c of this.components) {
      const b = c.render();
      this.ctx.drawImage(b, c.getX(), c.getY());
    }

    // render border
    this.ctx.lineWidth = GDIStyle.tabBorderSize;
    this.ctx.strokeStyle = GDIStyle.tabBorderColor;
    this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());

    // convert to image and return
    return this.canvas.transferToImageBitmap();
  }
}
