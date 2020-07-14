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

        // border offset
        const bx = GDIStyle.listBorderSize + 3;

        // iterate over items and draw text
        let r = 0;
        const stop = Math.round((this.getHeight() - bx) / px);

        for (let i = this.scroll; i < (this.scroll + stop); i++) {
            // exit if out of bounds
            if (i >= this.items.length) {
                this.scroll = i - (stop);
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

            // render text
            this.ctx.fillText(t, bx, (px * r) + bx);

            r++;
        }

        const sw = GDIStyle.listScrollWidth;

        // render scroll bar
        this.ctx.fillStyle = GDIStyle.listFillColor;
        this.ctx.fillRect(this.getWidth() - sw, 0, sw, this.getHeight());

        this.ctx.fillStyle = GDIStyle.listScrollColor;
        if (stop >= this.items.length) {
            this.ctx.fillRect(this.getWidth() - sw, 0, sw, this.getHeight());
        } else {
            const scale = stop / this.items.length;
            const percent = this.scroll / this.items.length;

            this.ctx.fillRect(this.getWidth() - sw, (percent * this.getHeight()), sw, (scale * this.getHeight()) + 2);
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

    handleScroll(x: number, y: number, d: number) {
        if (d > 0) {
            this.scrollPlus();
        } else if (d < 0) {
            this.scrollMinus();
        }

        this.boundCheck();
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

    scrollPlus() {
        this.scroll += 1;
    }

    scrollMinus() {
        this.scroll -= 1;
    }
}
