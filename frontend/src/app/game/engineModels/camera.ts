export class Camera {
  constructor(width: number, height: number, zoom: number) {
    this.w = width;
    this.h = height;
    this.zoom = zoom;
    this.x = 0;
    this.y = 0;
  }

  x: number;
  y: number;
  w: number;
  h: number;
  zoom: number;

  projectX(ex: number) {
    return (ex - this.x) * this.zoom + this.w / 2;
  }

  projectY(ey: number) {
    return (ey - this.y) * this.zoom + this.h / 2;
  }

  projectR(er: number) {
    return er * this.zoom;
  }
}
