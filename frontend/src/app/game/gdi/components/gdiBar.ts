import { GDIBase } from '../base/gdiBase';
import { GDIStyle, FontSize } from '../base/gdiStyle';

export class GDIBar extends GDIBase {
  private canvas: OffscreenCanvas;
  private ctx: any;

  private text = '';
  private percentage = 0;
  private allowOverflow = false;
  private showPercentage = false;
  private font: FontSize = FontSize.normal;
  private color = GDIStyle.barFillColor;

  initialize() {
    // initialize offscreen canvas
    this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight());
    this.ctx = this.canvas.getContext('2d');
  }

  periodicUpdate() {
    //
  }

  render(): ImageBitmap {
    // render background
    this.ctx.fillStyle = GDIStyle.windowFillColor;
    this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

    // render bar
    this.ctx.fillStyle = this.color;
    this.ctx.fillRect(
      0,
      0,
      this.getWidth() * (this.percentage / 100),
      this.getHeight(),
    );

    // style text
    this.ctx.textAlign = 'center';
    this.ctx.textBaseline = 'middle';

    this.ctx.fillStyle = GDIStyle.barTextColor;
    this.ctx.strokeStyle = GDIStyle.barTextColor;
    this.ctx.font = GDIStyle.getUnderlyingFont(this.getFont());

    // render text
    this.ctx.fillText(
      this.getText() +
        (this.showPercentage ? ` (${Math.round(this.percentage)}%)` : ''),
      this.getWidth() / 2,
      this.getHeight() / 2,
    );

    if (GDIStyle.barBorderSize > 0) {
      // render border
      this.ctx.lineWidth = GDIStyle.barBorderSize;
      this.ctx.strokeStyle = GDIStyle.barBorderColor;
      this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
    }

    // convert to image and return
    return this.canvas.transferToImageBitmap();
  }

  setText(text: string) {
    this.text = text;
  }

  getText(): string {
    return this.text;
  }

  setFont(font: FontSize) {
    this.font = font;
  }

  getFont(): FontSize {
    return this.font;
  }

  setPercentage(p: number) {
    if (p > 100 && !this.allowOverflow) {
      p = 100;
    }

    if (p < 0) {
      p = 0;
    }

    this.percentage = p;
  }

  getPercentage(): number {
    return this.percentage;
  }

  setColor(c: string) {
    this.color = c;
  }

  getColor(): string {
    return this.color;
  }

  getShowPercentage(): boolean {
    return this.showPercentage;
  }

  setShowPercentage(b: boolean) {
    this.showPercentage = b;
  }

  getAllowOverflow(): boolean {
    return this.allowOverflow;
  }

  setAllowOverflow(b: boolean) {
    this.allowOverflow = b;
  }
}
