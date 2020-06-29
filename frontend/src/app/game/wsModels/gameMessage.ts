export class GameMessage {
    type: number;
    body: string;
}

export enum MessageTypes {
    Join = 0,
    Update = 1,
    NavClick = 2
}
