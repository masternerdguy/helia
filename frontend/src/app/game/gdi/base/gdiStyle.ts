export class GDIStyle {
  // window
  static windowBorderColor = 'white';
  static windowFillColor = '#00091a';
  static windowBorderSize = 2;
  static windowHandleHeight = 20;
  static windowHandleColor = '#336699';
  static windowHandleDraggingColor = '#0099ff';
  static windowHandleTextColor = 'white';

  // button
  static buttonBorderColor = 'white';
  static buttonFillColor = '#00091a';
  static toggleButtonEnabledTextColor = '#0099ff';
  static buttonTextColor = 'white';
  static buttonBorderSize = 2;

  // label
  static labelBorderColor = '';
  static labelFillColor = '#00091a';
  static labelTextColor = 'white';
  static labelBorderSize = 0;

  // overlay
  static overlayBorderColor = '';
  static overlayFillColor = '#00091a';
  static overlayBorderSize = 0;

  // tab
  static tabBorderColor = 'white';
  static tabFillColor = '#00091a';
  static tabTextColor = 'white';
  static tabActiveTextColor = '#0099ff';
  static tabBorderSize = 0;
  static tabHandleHeight = 20;

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

  // starmap
  static starMapEdgeColor = '#00091a';
  static starMapEdgeWidth = 1.3;

  // underlying font sizes
  static smallFontSize = 7.25;
  static smallNormalFontSize = 9.25;
  static normalFontSize = 11.25;
  static largeFontSize = 15.25;
  static giantFontSize = 23.25;

  // underlying fonts
  static smallFont = GDIStyle.smallFontSize + 'px FiraCode-Regular';
  static smallNormalFont = GDIStyle.smallNormalFontSize + 'px FiraCode-Regular';
  static normalFont = GDIStyle.normalFontSize + 'px FiraCode-Regular';
  static largeFont = GDIStyle.normalFontSize + 'px FiraCode-Regular';
  static giantFont = GDIStyle.giantFontSize + 'px FiraCode-Regular';

  // standardized bar colors for gameplay elements
  static shieldBarColor = '#154360';
  static armorBarColor = '#4A235A';
  static hullBarColor = '#78281F';
  static energyBarColor = '#1D8348';
  static heatBarColor = '#BA4A00';
  static fuelBarColor = '#4D5656';

  // standardized override colors for ui elements
  static errorFillColor = '#b30000';

  // helpers
  static getUnderlyingFont(font: FontSize): string {
    let f = '';

    if (font === FontSize.small) {
      f = GDIStyle.smallFont;
    } else if (font === FontSize.smallNormal) {
      f = GDIStyle.smallNormalFont;
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
    } else if (font === FontSize.smallNormal) {
      px = GDIStyle.smallNormalFontSize;
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
  smallNormal = 'smallNormal',
  normal = 'normal',
  large = 'large',
  giant = 'giant',
}
