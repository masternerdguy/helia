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
}
