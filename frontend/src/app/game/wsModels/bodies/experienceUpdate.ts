export class ServerExperienceUpdate {
  ships: ServerExperienceShipEntry[];
  modules: ServerExperienceModuleEntry[];
}

export class ServerExperienceShipEntry {
  experienceLevel: number;
  shipTemplateID: string;
  shipTemplateName: string;
}

export class ServerExperienceModuleEntry {
  experienceLevel: number;
  itemTypeID: string;
  itemTypeName: string;
}
