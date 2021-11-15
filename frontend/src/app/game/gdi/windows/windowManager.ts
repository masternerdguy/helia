import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIWindow } from '../base/gdiWindow';
import { GDIToggleButton } from '../components/gdiToggleButton';

export class WindowManager extends GDIWindow {
  windowToggles: [any, any] = [null, null];
  private showHideCallback: (w: GDIWindow) => void;

  preinit(height: number, showHideCallback: (w: GDIWindow) => void) {
    this.setHeight(height);
    this.showHideCallback = showHideCallback;
  }

  resize(height: number) {
    // store new size
    this.setHeight(height);

    // get new offscreen canvas
    this.buildCanvas();
  }

  initialize() {
    // set dimensions
    this.setWidth(GDIStyle.giantFontSize);
    this.setBorderless(true);

    // initialize
    super.initialize();

    // empty null initial value
    while (this.windowToggles.length > 0) {
      this.windowToggles.pop();
    }
  }

  pack() {
    this.setTitle('Window Manager');
  }

  periodicUpdate() {
    for (const wt of this.windowToggles) {
      // show/hide windows based on toggles
      const w = wt[0] as GDIWindow;
      const b = wt[1] as GDIToggleButton;

      // update toggle based on window state
      b.setToggled(!w.isHidden());
    }

    // enforce position on screen
    this.setX(0);
    this.setY(-GDIStyle.windowHandleHeight);
  }

  manageWindow(w: GDIWindow, i: string) {
    // set up toggle button for this window
    const width = this.getWidth();
    const wt = new GDIToggleButton();

    wt.setWidth(width);
    wt.setHeight(width);
    wt.setX(0);
    wt.setY(this.windowToggles.length * width);

    wt.setText(i);
    wt.setOnClick(() => {
      w.setHidden(!wt.isToggled());
      this.showHideCallback(w);
    });

    wt.setFont(FontSize.giant);
    wt.setBorderless(true);
    wt.initialize();

    // store window and toggle
    this.windowToggles.push([w, wt]);
    this.addComponent(wt);
  }
}
