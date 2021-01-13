import { GDIBase } from '../base/gdiBase';
import { GDIStyle } from '../base/gdiStyle';

export class GDIOverlay extends GDIBase {
  private canvas: OffscreenCanvas;
  private ctx: any;

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
    this.ctx.fillStyle = GDIStyle.overlayFillColor;
    this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

    if (GDIStyle.overlayBorderSize > 0) {
      // render border
      this.ctx.lineWidth = GDIStyle.overlayBorderSize;
      this.ctx.strokeStyle = GDIStyle.overlayBorderColor;
      this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
    }

    // convert to image and return
    return this.canvas.transferToImageBitmap();
  }

  containsPoint(x: number, y: number): boolean {
    return false;
  }
}
