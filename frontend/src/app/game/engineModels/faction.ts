import { WSFaction } from '../wsModels/entities/wsFaction';
import { WSPlayerFactionRelationship } from '../wsModels/entities/wsPlayerFaction';

export class Faction extends WSFaction {
  constructor(s: WSFaction) {
    super();

    Object.assign(this, s);
  }
}
