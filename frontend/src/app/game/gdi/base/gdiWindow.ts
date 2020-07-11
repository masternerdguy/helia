import { GDIBase } from './gdiBase';
import { GDIComponent } from './gdiComponent';
import { GDIStyle } from './gdiStyle';
import { Rectangle } from './rectangle';

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

        // render background
        this.ctx.fillRect(0, 0, this.getWidth(), this.getHeight());

        // render components
        for (const c of this.components) {
            const b = c.render();
            this.ctx.drawImage(b, c.getX(), c.getY());
        }

        // render border
        this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight());

        // convert to image and return
        return this.canvas.transferToImageBitmap();
    }

    handleClick(x: number, y: number) {
        // make sure this is a real click
        if (!this.containsPoint(x, y)) {
            return;
        }

        // adjust input to be relative to window origin
        const rX = x - this.getX();
        const rY = y - this.getY();

        // find the component that is being clicked on within this window
        for (const c of this.components) {
            if (c.containsPoint(rX, rY)) {
                // send click event
                c.handleClick(rX, rY);
                break;
            }
        }
    }

    pack() {
        throw new Error('Must override in derived class.');
    }

    addComponent(component: GDIComponent) {
        this.components.push(component);
    }
}
