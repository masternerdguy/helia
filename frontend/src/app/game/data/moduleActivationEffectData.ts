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

    basicShieldBooster(): ModuleActivationEffectData {
        return {
            type: 'bubble_shield_boost',
            duration: 450,
            color: '#A1FA61',
            filter: 'blur(5px)', // "feather"
            thickness: 2
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
