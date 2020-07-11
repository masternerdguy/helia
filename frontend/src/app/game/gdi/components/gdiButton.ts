import { GDIBase } from '../base/gdiBase';
import { GDIStyle } from '../base/gdiStyle';

export class GDIButton extends GDIBase {
    private canvas: OffscreenCanvas;
    private ctx: any;

    private text = '';
    private font: FontSize =  FontSize.normal;

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
        this.ctx.fillStyle = GDIStyle.buttonFillColor;
        this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

        // render text
        this.ctx.textAlign = 'center';
        this.ctx.textBaseline = 'middle';

        // get actual font
        let f = '';

        if (this.font === FontSize.small) {
            f = GDIStyle.smallFont;
        } else if (this.font === FontSize.normal) {
            f = GDIStyle.normalFont;
        }

        this.ctx.fillStyle = GDIStyle.buttonTextColor;
        this.ctx.strokeStyle = GDIStyle.buttonTextColor;
        this.ctx.font = f;

        this.ctx.fillText(this.getText(), this.getWidth() / 2, this.getHeight() / 2);

        // render border
        this.ctx.lineWidth = GDIStyle.buttonBorderSize;
        this.ctx.strokeStyle = GDIStyle.buttonBorderColor;
        this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());

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

export enum FontSize {
    small = 'small',
    normal = 'normal'
}
