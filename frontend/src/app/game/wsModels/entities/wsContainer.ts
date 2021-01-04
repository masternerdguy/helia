export class WSContainer {
  id: string;
  items: WSContainerItem[];
}

export class WSContainerItem {
  id: string;
  itemTypeID: string;
  itemTypeName: string;
  itemFamilyID: string;
  itemFamilyName: string;
  quantity: number;
  isPackaged: boolean;
  meta: any;
  itemTypeMeta: any;
}
