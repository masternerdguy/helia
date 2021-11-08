import { GDIBase } from '../base/gdiBase';
import { GDIStyle, FontSize } from '../base/gdiStyle';
import { GDIRectangle } from '../base/gdiRectangle';
import { getCharWidth } from '../../engineMath';

export class GDIList extends GDIBase {
  private canvas: OffscreenCanvas;
  private ctx: any;

  private items: any[] = [];
  private scroll = 0;
  private selectedRow = -1;
  private font: FontSize = FontSize.normal;

  private onClick: (item: any) => void;

  private overrideFillColor: string;
  private overrideTextColor: string;
  private overrideBorderColor: string;

  initialize() {
    // initialize offscreen canvas
    this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight());
    this.ctx = this.canvas.getContext('2d');
  }

  periodicUpdate() {
    this.boundCheck();
  }

  boundCheck() {
    if (this.scroll > this.items.length) {
      this.scroll = this.items.length;
    } else if (this.scroll < 0) {
      this.scroll = 0;
    }
  }

  render(): ImageBitmap {
    // render background
    this.ctx.fillStyle = this.overrideFillColor ?? GDIStyle.listFillColor;
    this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

    // style text
    this.ctx.textAlign = 'left';
    this.ctx.textBaseline = 'bottom';
    this.ctx.font = GDIStyle.getUnderlyingFont(this.getFont());

    // get font size
    const px = GDIStyle.getUnderlyingFontSize(this.getFont());

    // border offset
    const bx = GDIStyle.listBorderSize + 3;

    // iterate over items and draw text
    let r = 0;
    const stop = Math.round((this.getHeight() - bx) / px);

    this.boundCheck();

    for (let i = this.scroll; i < this.scroll + stop; i++) {
      // exit if out of bounds
      if (i >= this.items.length) {
        this.scroll = i - stop;
        break;
      }

      // draw selected background if selected
      if (i === this.selectedRow) {
        this.ctx.fillStyle = GDIStyle.listSelectedColor;
        this.ctx.fillRect(
          bx,
          px * r + bx,
          this.getWidth() - GDIStyle.listScrollWidth,
          px
        );
      }

      // get item
      const item = this.items[i];

      // get text from item
      let t = '';
      if (item.listString) {
        t = item.listString();
      } else {
        t = JSON.stringify(item);
      }

      // render text
      this.ctx.fillStyle = this.overrideTextColor ?? GDIStyle.listTextColor;
      this.ctx.strokeStyle = this.overrideTextColor ?? GDIStyle.listTextColor;

      this.ctx.fillText(t, bx, px * (r + 1) + bx);

      // iterate row counter
      r++;
    }

    const sw = GDIStyle.listScrollWidth;

    // render scroll bar
    this.ctx.fillStyle = this.overrideFillColor ?? GDIStyle.listFillColor;
    this.ctx.fillRect(this.getWidth() - sw, 0, sw, this.getHeight());

    this.ctx.fillStyle = GDIStyle.listScrollColor;
    if (stop >= this.items.length) {
      this.ctx.fillRect(this.getWidth() - sw, 0, sw, this.getHeight());
    } else {
      const scale = stop / this.items.length;
      const percent = this.scroll / this.items.length;

      this.ctx.fillRect(
        this.getWidth() - sw,
        percent * this.getHeight(),
        sw,
        scale * this.getHeight() + 2
      );
    }

    if (GDIStyle.listBorderSize > 0) {
      // render border
      this.ctx.lineWidth = GDIStyle.listBorderSize;
      this.ctx.strokeStyle =
        this.overrideBorderColor ?? GDIStyle.listBorderColor;
      this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
    }

    // convert to image and return
    return this.canvas.transferToImageBitmap();
  }

  handleClick(x: number, y: number) {
    // make sure this is a real click
    if (!this.containsPoint(x, y)) {
      return;
    }

    this.boundCheck();

    // adjust input to be relative to control origin
    const rX = x - this.getX();
    const rY = y - this.getY();

    // get font size
    const px = GDIStyle.getUnderlyingFontSize(this.getFont());

    // border offset
    const bx = GDIStyle.listBorderSize + 3;

    // find item being clicked on
    let r = 0;
    const stop = Math.round((this.getHeight() - bx) / px);

    for (let i = this.scroll; i < this.scroll + stop; i++) {
      // exit if out of bounds
      if (i >= this.items.length) {
        break;
      }

      // test for click
      const itemRect = new GDIRectangle(
        bx,
        px * r + bx,
        this.getWidth() - GDIStyle.listScrollWidth,
        px
      );
      const item = this.items[i];

      if (itemRect.containsPoint(rX, rY)) {
        this.selectedRow = i;
        this.onClick(item);
        break;
      }

      r++;
    }
  }

  handleScroll(x: number, y: number, d: number) {
    if (d > 0) {
      this.scrollPlus();
    } else if (d < 0) {
      this.scrollMinus();
    }

    this.boundCheck();
  }

  setOnClick(h: (item: any) => void) {
    this.onClick = h;
  }

  setItemsFromText(text: string) {
    // break text into rows
    const rows = this.breakText(text);

    // push to view
    this.setItems(rows);
  }

  breakText(rawText: string): any[] {
    // capture argument
    let text = `${rawText}`;
    const rows = [];

    // first break by row width
    const fontWidth = getCharWidth(
      ' ',
      GDIStyle.getUnderlyingFont(this.getFont())
    );

    const breakCol = Math.round(
      (this.getWidth() -
        (GDIStyle.listScrollWidth + 2 * (GDIStyle.listBorderSize + 3))) /
        fontWidth -
        0.5
    );

    let accIdx = 0;

    for (let i = 0; i < `${text}`.length; i++) {
      // reset if line break encountered
      if (text[i] == '\n') {
        accIdx = 0;
        continue;
      }

      // determine whether to back-insert a break and back-track
      if (accIdx == breakCol) {
        accIdx = 0;

        if (text[i] != '\n') {
          for (let j = breakCol; j >= 0; j--) {
            let x = i - breakCol + j;

            // break on previous space and back-track to last space position
            if (text[x] == ' ') {
              const s1 = text.substring(0, x);
              const s2 = text.substring(x);

              text = `${s1.trim()}ⶤⶤ${s2.trim()}`;
              i = x;

              break;
            }
          }
        }
      }

      accIdx++;
    }

    const lBreaks = text.split('ⶤⶤ');
    let nText = '';

    for (const ln of lBreaks) {
      nText += ln + '\n';
    }

    // then break text by newlines
    const byNewLines = nText.split('\n');

    for (const lbRow of byNewLines) {
      rows.push({
        text: `${lbRow}`,
        listString: () => lbRow,
      });
    }

    // return text rows
    return rows;
  }

  setItems(items: any[]) {
    this.selectedRow = -1;
    this.items = items;
  }

  getItems(): any[] {
    return this.items;
  }

  setFont(font: FontSize) {
    this.font = font;
  }

  getFont(): FontSize {
    return this.font;
  }

  getScroll(): number {
    return this.scroll;
  }

  setScroll(s: number) {
    this.scroll = Math.round(s);
  }

  scrollPlus() {
    this.scroll += 1;
  }

  scrollMinus() {
    this.scroll -= 1;
  }

  getSelectedIndex(): number {
    return this.selectedRow;
  }

  setSelectedIndex(i: number) {
    this.selectedRow = i;
  }

  getSelectedItem(): any {
    return this.items[this.selectedRow];
  }

  getOverrideFillColor(): string {
    return this.overrideFillColor;
  }

  setOverrideFillColor(color: string) {
    this.overrideFillColor = color;
  }

  getOverrideTextColor(): string {
    return this.overrideTextColor;
  }

  setOverrideTextColor(color: string) {
    this.overrideTextColor = color;
  }

  getOverrideBorderColor(): string {
    return this.overrideBorderColor;
  }

  setOverrideBorderColor(color: string) {
    this.overrideBorderColor = color;
  }
}
