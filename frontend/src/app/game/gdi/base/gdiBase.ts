import { GDIComponent } from './gdiComponent';
import { GDIRectangle } from './gdiRectangle';

export class GDIBase implements GDIComponent {
  private x: number;
  private y: number;
  private w: number;
  private h: number;

  initialize() {
    throw new Error('Must override in derived class.');
  }

  getX(): number {
    return this.x;
  }

  getY(): number {
    return this.y;
  }

  getWidth(): number {
    return this.w;
  }

  getHeight(): number {
    return this.h;
  }

  setX(x: number) {
    this.x = Math.round(x);
  }

  setY(y: number) {
    this.y = Math.round(y);
  }

  setWidth(w: number) {
    this.w = Math.round(w);
  }

  setHeight(h: number) {
    this.h = Math.round(h);
  }

  periodicUpdate() {
    throw new Error('Must override in derived class.');
  }

  render(): ImageBitmap {
    throw new Error('Must override in derived class.');
  }

  containsPoint(x: number, y: number): boolean {
    const rect = new GDIRectangle(
      this.getX(),
      this.getY(),
      this.getWidth(),
      this.getHeight(),
    );
    return rect.containsPoint(x, y);
  }

  handleMouseMove(x: number, y: number) {
    //
  }

  handleClick(x: number, y: number) {
    //
  }

  handleScroll(x: number, y: number, d: number) {
    //
  }

  handleKeyDown(x: number, y: number, key: string) {
    //
  }
}
