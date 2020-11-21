import { WsPushModuleEffect } from '../wsModels/entities/wsPushModuleEffect';

export class ModuleEffect extends WsPushModuleEffect {
    constructor(b: WsPushModuleEffect) {
        super();
        Object.assign(this, b);
    }
}
