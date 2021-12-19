export class ServerExperienceUpdate {
  ships: ServerExperienceShipEntry[];
}

export class ServerExperienceShipEntry {
  experienceLevel: number;
  shipTemplateID: string;
  shipTemplateName: string;
}
