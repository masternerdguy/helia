import { TargetType } from '../../engineModels/player';

export class WsPushModuleEffect {
  gfxEffect: string;
  objStartID: string;
  objStartType: TargetType;
  objStartHPOffset: number[];
  objEndID: string;
  objEndType: TargetType;
}
