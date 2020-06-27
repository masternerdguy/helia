export class Camera {
    constructor(width: number, height: number) {
        this.w = width;
        this.h = height;
        this.x = 0;
        this.y = 0;
    }

    x: number;
    y: number;
    w: number;
    h: number;

    projectX(ex: number) {
        return ex - (this.x - this.w / 2);
    }

    projectY(ey: number) {
        return ey - (this.y - this.h / 2);
    }
}
