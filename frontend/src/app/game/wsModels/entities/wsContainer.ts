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
  schematic: WSSchematicProcess;
}

export class WSSchematicProcess {
  id: string;
  time: number;
  inputs: WSSchematicProcessFactor[];
  outputs: WSSchematicProcessFactor[];
}

export class WSSchematicProcessFactor {
  itemTypeId: string;
  itemTypeName: string;
  quantity: number;
}
