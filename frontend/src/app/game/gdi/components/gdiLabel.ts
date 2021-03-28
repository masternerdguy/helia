import { GDIBase } from '../base/gdiBase';
import { GDIStyle, FontSize } from '../base/gdiStyle';

export class GDILabel extends GDIBase {
  private canvas: OffscreenCanvas;
  private ctx: any;

  private text = '';
  private font: FontSize = FontSize.normal;

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
    this.ctx.fillStyle = GDIStyle.labelFillColor;
    this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

    // style text
    this.ctx.textAlign = 'center';
    this.ctx.textBaseline = 'middle';

    this.ctx.fillStyle = GDIStyle.labelTextColor;
    this.ctx.strokeStyle = GDIStyle.labelTextColor;
    this.ctx.font = GDIStyle.getUnderlyingFont(this.getFont());

    // render text
    this.ctx.fillText(
      this.getText(),
      this.getWidth() / 2,
      this.getHeight() / 2
    );

    if (GDIStyle.labelBorderSize > 0) {
      // render border
      this.ctx.lineWidth = GDIStyle.labelBorderSize;
      this.ctx.strokeStyle = GDIStyle.labelBorderColor;
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
}
