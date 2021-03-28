import { GDIBase } from '../base/gdiBase';
import { GDIStyle, FontSize } from '../base/gdiStyle';

export class GDIToggleButton extends GDIBase {
  private canvas: OffscreenCanvas;
  private ctx: any;

  private text = '';
  private font: FontSize = FontSize.normal;

  private onClick: (x: number, y: number) => void;
  private enabled = true;
  private toggled = false;
  private borderless = false;

  initialize() {
    // initialize offscreen canvas
    this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight());
    this.ctx = this.canvas.getContext('2d');
  }

  periodicUpdate() {
    // 
  }

  render(): ImageBitmap {
    if (!this.enabled) {
      // render background
      this.ctx.fillStyle = GDIStyle.buttonFillColor;
      this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

      if (GDIStyle.buttonBorderSize > 0) {
        // render border
        this.ctx.lineWidth = GDIStyle.buttonBorderSize;
        this.ctx.strokeStyle = GDIStyle.buttonBorderColor;
        this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
      }
    } else {
      // render background
      this.ctx.fillStyle = GDIStyle.buttonFillColor;
      this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

      // style text
      this.ctx.textAlign = 'center';
      this.ctx.textBaseline = 'middle';

      if (!this.toggled) {
        this.ctx.fillStyle = GDIStyle.buttonTextColor;
        this.ctx.strokeStyle = GDIStyle.buttonTextColor;
      } else {
        this.ctx.fillStyle = GDIStyle.toggleButtonEnabledTextColor;
        this.ctx.strokeStyle = GDIStyle.toggleButtonEnabledTextColor;
      }

      this.ctx.font = GDIStyle.getUnderlyingFont(this.getFont());

      // render text
      this.ctx.fillText(
        this.getText(),
        this.getWidth() / 2,
        this.getHeight() / 2
      );

      if (GDIStyle.buttonBorderSize > 0 && !this.borderless) {
        // render border
        this.ctx.lineWidth = GDIStyle.buttonBorderSize;
        this.ctx.strokeStyle = GDIStyle.buttonBorderColor;
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
    this.onClick = (x, y) => {
      // set toggle
      this.toggled = !this.toggled;

      // invoke callback
      h(x, y);
    };
  }

  setEnabled(enabled: boolean) {
    this.enabled = enabled;
  }

  isToggled(): boolean {
    return this.toggled;
  }

  setToggled(b: boolean) {
    this.toggled = b;
  }

  isBorderless(): boolean {
    return this.borderless;
  }

  setBorderless(b: boolean) {
    this.borderless = b;
  }
}
