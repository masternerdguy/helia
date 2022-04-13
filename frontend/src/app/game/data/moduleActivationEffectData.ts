export class ModuleActivationEffectRepository {
  leBanhammer(): ModuleActivationEffectData {
    return {
      type: 'gauss',
      duration: 3500,
      color: 'gold',
      filter: 'blur(10px)', // "feather"
      thickness: 7.5,
    };
  }

  basicLaserTool(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 1500,
      color: 'red',
      filter: 'blur(1px)', // "feather"
      thickness: 0.5,
    };
  }

  basicShieldLaser(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 1320,
      color: 'limegreen',
      filter: 'blur(1px)', // "feather"
      thickness: 0.4,
    };
  }

  basicHullLaser(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 1352,
      color: 'goldenrod',
      filter: 'blur(1px)', // "feather"
      thickness: 0.6,
    };
  }

  basicGeneralLaser(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 1410,
      color: 'wheat',
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

  basicIceHarvester(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 3250,
      color: 'darkblue',
      filter: 'blur(3px)', // "feather"
      thickness: 1.22,
    };
  }

  basicOreHarvester(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 2819,
      color: 'darkred',
      filter: 'blur(2px)', // "feather"
      thickness: 0.95,
    };
  }

  mielIceHarvester(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 52789,
      color: 'darkviolet',
      filter: 'blur(4px)', // "feather"
      thickness: 1.35,
    };
  }

  lecheOreHarvester(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 53621,
      color: 'darkorange',
      filter: 'blur(3px)', // "feather"
      thickness: 1.01,
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

  basicSalvager(): ModuleActivationEffectData {
    return {
      type: 'salvager',
      duration: 10000,
      color: '#ff6347',
      filter: 'blur(1.5px)', // "feather"
      thickness: 1.8,
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

  basicUtilitySiphon(): ModuleActivationEffectData {
    return {
      type: 'siphon',
      duration: 1865,
      color: 'orange',
      filter: 'blur(9.28px)', // "feather"
      thickness: 1.2,
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
