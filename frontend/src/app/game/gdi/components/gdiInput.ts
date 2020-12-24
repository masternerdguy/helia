import { GDIBase } from '../base/gdiBase';
import { GDIStyle, FontSize } from '../base/gdiStyle';

export class GDIInput extends GDIBase {
  private canvas: OffscreenCanvas;
  private ctx: any;

  private text = '';
  private font: FontSize = FontSize.normal;

  private onReturn: (txt: string) => void;

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
    this.ctx.fillStyle = GDIStyle.inputFillColor;
    this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

    // style text
    this.ctx.textAlign = 'left';
    this.ctx.textBaseline = 'middle';

    this.ctx.fillStyle = GDIStyle.inputTextColor;
    this.ctx.strokeStyle = GDIStyle.inputTextColor;
    this.ctx.font = GDIStyle.getUnderlyingFont(this.getFont());

    // render text
    this.ctx.fillText(
      this.getText() + '┊',
      GDIStyle.inputBorderSize,
      this.getHeight() / 2
    );

    if (GDIStyle.inputBorderSize > 0) {
      // render border
      this.ctx.lineWidth = GDIStyle.inputBorderSize;
      this.ctx.strokeStyle = GDIStyle.inputBorderColor;
      this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
    }

    // convert to image and return
    return this.canvas.transferToImageBitmap();
  }

  handleKeyDown(x: number, y: number, key: string) {
    // make sure this is a relevant key press
    if (!this.containsPoint(x, y)) {
      return;
    }

    // check for return / enter press
    if (key === 'Enter') {
      this.onReturn(this.text);
      return;
    }

    // ignore some keys
    if (key === 'Shift') {
      return;
    }

    // check for backspace
    if (key === 'Backspace') {
      if (this.text.length > 0) {
        // slice last character off of text
        this.text = this.text.slice(0, -1);
      }
    } else {
      // append key
      this.text += key;
    }
  }

  setOnReturn(h: (txt: string) => void) {
    this.onReturn = h;
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
