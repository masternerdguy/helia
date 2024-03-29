import { GDIBase } from '../base/gdiBase';
import { GDIStyle, FontSize } from '../base/gdiStyle';

export class GDIButton extends GDIBase {
  private canvas: OffscreenCanvas;
  private ctx: any;

  private text = '';
  private font: FontSize = FontSize.normal;

  private onClick: (x: number, y: number) => void;
  private enabled: boolean;

  private fillColor: string = GDIStyle.buttonFillColor;
  private borderColor: string = GDIStyle.buttonBorderColor;
  private textColor: string = GDIStyle.buttonTextColor;

  initialize() {
    // initialize offscreen canvas
    this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight());
    this.ctx = this.canvas.getContext('2d');
    this.enabled = true;
  }

  periodicUpdate() {
    //
  }

  render(): ImageBitmap {
    if (!this.enabled) {
      // render background
      this.ctx.fillStyle = this.fillColor;
      this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

      if (GDIStyle.buttonBorderSize > 0) {
        // render border
        this.ctx.lineWidth = GDIStyle.buttonBorderSize;
        this.ctx.strokeStyle = this.borderColor;
        this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
      }
    } else {
      // render background
      this.ctx.fillStyle = this.fillColor;
      this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

      // style text
      this.ctx.textAlign = 'center';
      this.ctx.textBaseline = 'middle';

      this.ctx.fillStyle = this.textColor;
      this.ctx.strokeStyle = this.textColor;
      this.ctx.font = GDIStyle.getUnderlyingFont(this.getFont());

      // render text
      this.ctx.fillText(
        this.getText(),
        this.getWidth() / 2,
        this.getHeight() / 2,
      );

      if (GDIStyle.buttonBorderSize > 0) {
        // render border
        this.ctx.lineWidth = GDIStyle.buttonBorderSize;
        this.ctx.strokeStyle = this.borderColor;
        this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
      }
    }

    // convert to image and return
    return this.canvas.transferToImageBitmap();
  }

  handleClick(x: number, y: number) {
    if (!this.enabled) {
      return;
    }

    this.onClick(x, y);
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

  setOnClick(h: (x: number, y: number) => void) {
    this.onClick = h;
  }

  setEnabled(enabled: boolean) {
    this.enabled = enabled;
  }

  public getFillColor(): string {
    return this.fillColor;
  }
  public setFillColor(value: string) {
    this.fillColor = value;
  }

  public getBorderColor(): string {
    return this.borderColor;
  }
  public setBorderColor(value: string) {
    this.borderColor = value;
  }

  public getTextColor(): string {
    return this.textColor;
  }
  public setTextColor(value: string) {
    this.textColor = value;
  }
}
