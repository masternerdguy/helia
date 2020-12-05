import { FontSize, GDIStyle } from '../base/gdiStyle';
import { GDIWindow } from '../base/gdiWindow';
import { GDIToggleButton } from '../components/gdiToggleButton';

export class WindowManager extends GDIWindow {
    windowToggles: [any, any] = [null, null];

    preinit(height: number) {
        this.setHeight(height);
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

            w.setHidden(!b.isToggled());
        }

        // enforce position on screen
        this.setX(0);
        this.setY(-GDIStyle.windowHandleHeight);
    }

    manageWindow(w: GDIWindow, i: string, v: boolean) {
        // set up toggle button for this window
        const width = this.getWidth();
        const wt = new GDIToggleButton();

        wt.setWidth(width);
        wt.setHeight(width);
        wt.setX(0);
        wt.setY(this.windowToggles.length * width);

        wt.setText(i);
        wt.setOnClick(() => {});

        wt.setFont(FontSize.giant);
        wt.setBorderless(true);
        wt.setToggled(v);

        wt.initialize();

        // store window and toggle
        this.windowToggles.push([w, wt]);
        this.addComponent(wt);
    }
}
