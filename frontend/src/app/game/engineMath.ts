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
