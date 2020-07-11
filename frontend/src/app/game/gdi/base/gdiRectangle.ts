export class GDIRectangle {
    x: number;
    y: number;
    w: number;
    h: number;

    constructor(x: number, y: number, w: number, h: number) {
        this.x = x;
        this.y = y;
        this.w = w;
        this.h = h;
    }

    containsPoint(x: number, y: number) {
        return this.x <= x && x <= this.x + this.w &&
               this.y <= y && y <= this.y + this.h;
    }
}
