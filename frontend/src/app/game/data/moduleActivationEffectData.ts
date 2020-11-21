export class ModuleActivationEffectRepository {
    basicLaserTool(): ModuleActivationEffectData {
        return {
            type: 'laser',
            duration: 1500,
            color: 'red',
            filter: 'blur(1px)', // "feather"
            thickness: 0.5
        };
    }
}

export class ModuleActivationEffectData {
    type: string;
    duration: number;
    color?: string;
    filter?: string;
    thickness?: number;
}
