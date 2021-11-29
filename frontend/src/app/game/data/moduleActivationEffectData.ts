export class ModuleActivationEffectRepository {
  basicLaserTool(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 1500,
      color: 'red',
      filter: 'blur(1px)', // "feather"
      thickness: 0.5,
    };
  }

  basicGaussRifle(): ModuleActivationEffectData {
    return {
      type: 'gauss',
      duration: 325,
      color: 'white',
      filter: 'blur(1px)', // "feather"
      thickness: 0.75,
    };
  }

  basicIceMiner(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 1500,
      color: 'blue',
      filter: 'blur(2px)', // "feather"
      thickness: 0.65,
    };
  }

  basicShieldBooster(): ModuleActivationEffectData {
    return {
      type: 'bubble_shield_boost',
      duration: 450,
      color: '#A1FA61',
      filter: 'blur(5px)', // "feather"
      thickness: 2,
    };
  }

  basicAetherDragger(): ModuleActivationEffectData {
    return {
      type: 'aether_drag',
      duration: 10000,
      color: '#dbf4ff',
      filter: 'blur(5px)', // "feather"
      thickness: 3,
    };
  }

  basicAuto5Cannon(): ModuleActivationEffectData {
    return {
      type: 'autocannon',
      duration: 3500,
      color: 'yellow',
      filter: 'blur(0.25px)', // "feather"
      thickness: 0.78,
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
