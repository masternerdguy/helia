export class WSIndustrialOrdersUpdate {
  outSilos: WSIndustrialSilo[];
  inSilos: WSIndustrialSilo[];
}

export class WSIndustrialSilo {
  stationId: string;
  stationProcessId: string;
  itemTypeID: string;
  itemTypeName: string;
  itemFamilyID: string;
  itemFamilyName: string;
  price: number;
  available: number;
  meta: any;
  itemTypeMeta: any;
  isSelling: boolean;
}
