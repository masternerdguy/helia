export class GameMessage {
    type: number;
    body: string;
}

export enum MessageTypes {
    Join = 0,
    Update = 1,
    NavClick = 2,
    CurrentShipUpdate = 3,
    Goto = 4,
    Orbit = 5,
    Dock = 6
}
