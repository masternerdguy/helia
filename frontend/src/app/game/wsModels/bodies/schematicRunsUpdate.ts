export class ServerSchematicRunsUpdate {
    runs: ServerSchematicRunEntryBody[];
}

export class ServerSchematicRunEntryBody {
	id: string;
	schematicName: string;
	hostShipName: string;
	hostStationName: string;
	solarSystemName: string;
	statusId: string;
	percentageComplete: number;
}
