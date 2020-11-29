import { GDIBase } from './gdiBase';
import { GDIComponent } from './gdiComponent';
import { GDIRectangle } from './gdiRectangle';
import { GDIStyle, FontSize } from './gdiStyle';

export class GDIWindow extends GDIBase {
    private components: GDIComponent[];
    private canvas: OffscreenCanvas;
    private ctx: any;

    private dragMode = false;
    private title = '';

    initialize() {
        // initialize offscreen canvas
        this.canvas = new OffscreenCanvas(this.getWidth(), this.getHeight() + GDIStyle.windowHandleHeight);
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

    containsPoint(x: number, y: number): boolean {
        const rect = new GDIRectangle(this.getX(), this.getY(), this.getWidth(), this.getHeight() + GDIStyle.windowHandleHeight);
        return rect.containsPoint(x, y);
    }

    render(): ImageBitmap {
        // render handle background
        if (this.dragMode) {
            this.ctx.fillStyle = GDIStyle.windowHandleDraggingColor;
        } else {
            this.ctx.fillStyle = GDIStyle.windowHandleColor;
        }

        this.ctx.fillRect(0, 0, this.getWidth(), GDIStyle.windowHandleHeight);

        // render handle text
        this.ctx.fillStyle = GDIStyle.windowHandleTextColor;
        this.ctx.textAlign = 'left';
        this.ctx.textBaseline = 'middle';
        this.ctx.font = GDIStyle.getUnderlyingFont(FontSize.large);

        this.ctx.fillText(this.title, GDIStyle.windowBorderSize + 2, (GDIStyle.windowHandleHeight / 2));

        // render background
        this.ctx.fillStyle = GDIStyle.windowFillColor;
        this.ctx.fillRect(0, GDIStyle.windowHandleHeight, this.getWidth(), this.getHeight());

        // render components
        for (const c of this.components) {
            const b = c.render();
            this.ctx.drawImage(b, c.getX(), c.getY() + GDIStyle.windowHandleHeight);
        }

        // render border
        this.ctx.lineWidth = GDIStyle.windowBorderSize;
        this.ctx.strokeStyle = GDIStyle.windowBorderColor;
        this.ctx.strokeRect(0, 0, this.getWidth(), this.getHeight() + GDIStyle.windowHandleHeight);

        // convert to image and return
        return this.canvas.transferToImageBitmap();
    }

    handleClick(x: number, y: number) {
        // make sure this is a relevant click
        if (!this.containsPoint(x, y)) {
            return;
        }

        // check for click in handle
        const hY = y - this.getY();

        if (hY < GDIStyle.windowHandleHeight) {
            // toggle drag mode
            this.dragMode = !this.dragMode;
            return;
        }

        // adjust input to be relative to window origin
        const rX = x - this.getX();
        const rY = y - (this.getY() + GDIStyle.windowHandleHeight);

        // find the component that is being scrolled on within this window
        for (const c of this.components) {
            if (c.containsPoint(rX, rY)) {
                // send click event
                c.handleClick(rX, rY);
                break;
            }
        }
    }

    handleScroll(x: number, y: number, d: number) {
        // make sure this is a relevant scroll
        if (!this.containsPoint(x, y)) {
            return;
        }

        // adjust input to be relative to window origin
        const rX = x - this.getX();
        const rY = y - (this.getY() + GDIStyle.windowHandleHeight);

        // find the component that is being scrolled on within this window
        for (const c of this.components) {
            if (c.containsPoint(rX, rY)) {
                // send scroll event
                c.handleScroll(rX, rY, d);
                break;
            }
        }
    }

    handleKeyDown(x: number, y: number, key: string) {
        // make sure this is a relevant key press
        if (!this.containsPoint(x, y)) {
            return;
        }

        // adjust input to be relative to window origin
        const rX = x - this.getX();
        const rY = y - (this.getY() + GDIStyle.windowHandleHeight);

        // find the component that is being typed on within this window
        for (const c of this.components) {
            if (c.containsPoint(rX, rY)) {
                // send key event
                c.handleKeyDown(rX, rY, key);
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

    isDragging(): boolean {
        return this.dragMode;
    }

    getTitle(): string {
        return this.title;
    }

    setTitle(title: string) {
        this.title = title;
    }
}
