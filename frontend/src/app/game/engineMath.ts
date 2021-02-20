let sharedMathCanvas = document.createElement("canvas");

export function angleBetween(
  cx: number,
  cy: number,
  ex: number,
  ey: number
): number {
  const dy = ey - cy;
  const dx = ex - cx;
  let theta = Math.atan2(dy, dx);
  theta *= -180 / Math.PI;

  if (theta < 0) {
    theta = 360 + theta;
  }

  return theta;
}

export function magnitude(
  ax: number,
  ay: number,
  bx: number,
  by: number
): number {
  return Math.sqrt((ax - bx) * (ax - bx) + (ay - by) * (ay - by));
}

export function heliaDateFromString(realDate: string): Date {
  const d = Date.parse(realDate);
  return heliaDate(new Date(d));
}

export function heliaDate(realDate: Date): Date {
  const d = convertDateToUTC(realDate);

  d.setUTCFullYear(d.getUTCFullYear() - 859);
  d.setUTCDate(d.getUTCDate() + 28);
  d.setUTCHours(d.getUTCHours() + 6);
  d.setUTCMinutes(d.getUTCMinutes() + 42);
  d.setUTCSeconds(d.getUTCSeconds() + 12);

  return d;
}

export function printHeliaDate(d: Date): string {
  return d.toUTCString();
}

function convertDateToUTC(date: Date): Date {
  return new Date(
    date.getUTCFullYear(),
    date.getUTCMonth(),
    date.getUTCDate(),
    date.getUTCHours(),
    date.getUTCMinutes(),
    date.getUTCSeconds()
  );
}

export function getTextWidth(text: string, font: string): number {
  var context = sharedMathCanvas.getContext("2d");
  
  context.font = font;
  var metrics = context.measureText(text);

  return metrics.width;
}

export function getCharWidth(char: string, font: string): number {
  return getTextWidth(`${char}${char}${char}${char}${char}`, font) / 5;
}
