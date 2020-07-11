export class GDIStyle {
    // window
    static windowBorderColor = 'white';
    static windowFillColor = '#00091a';
    static windowBorderSize = 2;

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

    // underlying font sizes
    static smallFont =  '8px monospace';
    static normalFont =  '12px monospace';
    static largeFont =  '16px monospace';
    static giantFont =  '24px monospace';

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
}

export enum FontSize {
    small = 'small',
    normal = 'normal',
    large = 'large',
    giant = 'giant'
}
