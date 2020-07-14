import { GDIBase } from '../base/gdiBase';
import { GDIStyle, FontSize } from '../base/gdiStyle';

export class GDIList extends GDIBase {
    private canvas: OffscreenCanvas;
    private ctx: any;

    private items: any[] = [];
    private scroll = 0;
    private font: FontSize =  FontSize.normal;

    initialize() {
        // initialize offscreen canvas
        this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight());
        this.ctx = this.canvas.getContext('2d');
    }

    periodicUpdate() {
        if (this.scroll > this.items.length) {
            this.scroll = this.items.length;
        } else if (this.scroll < 0) {
            this.scroll = 0;
        }
    }

    render(): ImageBitmap {
        // render background
        this.ctx.fillStyle = GDIStyle.listFillColor;
        this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

        // style text
        this.ctx.textAlign = 'left';
        this.ctx.textBaseline = 'top';

        this.ctx.fillStyle = GDIStyle.listTextColor;
        this.ctx.strokeStyle = GDIStyle.listTextColor;
        this.ctx.font = GDIStyle.getUnderlyingFont(this.getFont());

        // get font size
        const px = GDIStyle.getUnderlyingFontSize(this.getFont());

        // iterate over items and draw text
        const stop = Math.round(this.getHeight() / px);
        for (let i = this.scroll; i < stop; i++) {
            // exit if out of bounds
            if (i >= this.items.length) {
                break;
            }

            const item = this.items[i];

            // get text from item
            let t = '';
            if (item.listString) {
                t = item.listString();
            } else {
                t = JSON.stringify(item);
            }

            // border offset
            const bx = GDIStyle.listBorderSize + 3;

            // render text
            this.ctx.fillText(t, bx, (px * i) + bx);
        }

        if (GDIStyle.listBorderSize > 0) {
            // render border
            this.ctx.lineWidth = GDIStyle.listBorderSize;
            this.ctx.strokeStyle = GDIStyle.listBorderColor;
            this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());
        }

        // convert to image and return
        return this.canvas.transferToImageBitmap();
    }

    setItems(items: any[]) {
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
}
