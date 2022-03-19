import { WSFaction } from '../wsModels/entities/wsFaction';

export class Faction extends WSFaction {
  constructor(s: WSFaction) {
    super();

    Object.assign(this, s);
  }
}
