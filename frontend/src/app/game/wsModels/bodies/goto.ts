import { TargetType } from '../../engineModels/player';

export class ClientGoto {
  sid: string;
  targetId: string;
  type: TargetType;
}
