export class ClientViewStarMap {
  sid: string;
}

export class ServerStarMapUpdate {
  cachedMapData: string;
}

export class UnwrappedStarMapData {

  constructor(cachedMapData: string) {
    // decode as JSON
    const asJSON = JSON.parse(cachedMapData);

    // debug out
    console.log(asJSON);
  }
}
