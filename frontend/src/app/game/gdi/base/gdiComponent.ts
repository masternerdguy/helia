export interface GDIComponent {
    initialize();

    getX(): number;
    getY(): number;
    getWidth(): number;
    getHeight(): number;

    setX(x: number);
    setY(y: number);
    setWidth(w: number);
    setHeight(h: number);

    periodicUpdate();
    render(): ImageBitmap;

    containsPoint(x: number, y: number): boolean;
    handleClick(x: number, y: number);
    handleScroll(x: number, y: number, d: number);
}
