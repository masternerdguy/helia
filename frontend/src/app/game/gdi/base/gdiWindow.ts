import { GDIBase } from './gdiBase';
import { GDIComponent } from './gdiComponent';
import { GDIStyle } from './gdiStyle';

export class GDIWindow extends GDIBase {
    private components: GDIComponent[];
    private canvas: OffscreenCanvas;
    private ctx: any;

    initialize() {
        // initialize offscreen canvas
        this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight());
        this.ctx = this.canvas.getContext('2d');

        // initialize empty arrays
        this.components = [];
    }

    periodicUpdate() {
        // update components
        for (const c of this.components) {
            c.periodicUpdate();
        }
    }

    render(): ImageBitmap {
        // set up window style
        this.ctx.lineWidth = GDIStyle.windowBorderSize;
        this.ctx.fillStyle = GDIStyle.windowFillColor;
        this.ctx.strokeStyle = GDIStyle.windowBorderColor;

        // render self
        this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());
        this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());

        // render components
        for (const c of this.components) {
            const b = c.render();
            this.ctx.drawImage(b, c.getX(), c.getY());
        }

        // convert to image and return
        return this.canvas.transferToImageBitmap();
    }
}
