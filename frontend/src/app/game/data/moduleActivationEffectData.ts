export class ModuleActivationEffectRepository {
    basicLaserTool(): ModuleActivationEffectData {
        return {
            type: 'laser',
            duration: 1500,
            color: 'red'
        };
    }
}

export class ModuleActivationEffectData {
    type: string;
    duration: number;
    color?: string;
}
