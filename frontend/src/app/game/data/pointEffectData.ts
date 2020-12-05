export class PointEffectRegistry {
    basicExplosion(): PointEffectData {
        return {
            type: 'point_explosion',
            duration: 3000,
            color: '#CCFFFF',
            filter: 'blur(10px)' // "feather"
        };
    }

    basicShieldBooster(): PointEffectData {
        return {
            type: 'bubble_shield_boost',
            duration: 450,
            color: '#A1FA61',
            filter: 'blur(5px)' // "feather"
        };
    }
}

export class PointEffectData {
    type: string;
    duration: number;
    color?: string;
    filter?: string;
    thickness?: number;
}
