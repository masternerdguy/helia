export class GDIStyle {
    // window
    static windowBorderColor = 'white';
    static windowFillColor = '#00091a';
    static windowBorderSize = 2;
    static windowHandleHeight = 20;
    static windowHandleColor = '#336699';
    static windowHandleDraggingColor = '#0099ff';

    // button
    static buttonBorderColor = 'white';
    static buttonFillColor = '#00091a';
    static buttonTextColor = 'white';
    static buttonBorderSize = 2;

    // label
    static labelBorderColor = '';
    static labelFillColor = '#00091a';
    static labelTextColor = 'white';
    static labelBorderSize = 0;

    // bar
    static barBorderColor = 'white';
    static barFillColor = '#0099ff';
    static barTextColor = 'white';
    static barBorderSize = 2;

    // list
    static listBorderColor = 'white';
    static listFillColor = '#00091a';
    static listScrollColor = 'white';
    static listTextColor = 'white';
    static listSelectedColor = '#0099ff';
    static listBorderSize = 2;
    static listScrollWidth = 15;

    // input
    static inputBorderColor = 'white';
    static inputFillColor = '#00091a';
    static inputTextColor = 'white';
    static inputBorderSize = 2;

    // underlying font sizes
    static smallFontSize =  8;
    static normalFontSize =  12;
    static largeFontSize =  16;
    static giantFontSize =  24;

    // underlying fonts
    static smallFont =  GDIStyle.smallFontSize + 'px monospace';
    static normalFont =  GDIStyle.normalFontSize + 'px monospace';
    static largeFont =  GDIStyle.normalFontSize + 'px monospace';
    static giantFont =  GDIStyle.giantFontSize + 'px monospace';

    // helpers
    static getUnderlyingFont(font: FontSize): string {
        let f = '';

        if (font === FontSize.small) {
            f = GDIStyle.smallFont;
        } else if (font === FontSize.normal) {
            f = GDIStyle.normalFont;
        } else if (font === FontSize.large) {
            f = GDIStyle.largeFont;
        } else if (font === FontSize.giant) {
            f = GDIStyle.giantFont;
        }

        return f;
    }

    static getUnderlyingFontSize(font: FontSize): number {
        let px = 0;

        if (font === FontSize.small) {
            px = GDIStyle.smallFontSize;
        } else if (font === FontSize.normal) {
            px = GDIStyle.normalFontSize;
        } else if (font === FontSize.large) {
            px = GDIStyle.largeFontSize;
        } else if (font === FontSize.giant) {
            px = GDIStyle.giantFontSize;
        }

        return px;
    }
}

export enum FontSize {
    small = 'small',
    normal = 'normal',
    large = 'large',
    giant = 'giant'
}
