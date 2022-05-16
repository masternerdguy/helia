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

  xlLaserTool(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 2400,
      color: 'red',
      filter: 'blur(4.9px)', // "feather"
      thickness: 2.7,
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

  xlShieldLaser(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 2590,
      color: 'limegreen',
      filter: 'blur(6px)', // "feather"
      thickness: 3.5,
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

  xlHullLaser(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 2818,
      color: 'goldenrod',
      filter: 'blur(7px)', // "feather"
      thickness: 4.5,
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

  xlGeneralLaser(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 2650,
      color: 'wheat',
      filter: 'blur(4.2px)', // "feather"
      thickness: 3.5,
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

  xlIceMiner(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 2260,
      color: 'blue',
      filter: 'blur(11.3px)', // "feather"
      thickness: 3.7,
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

  xlIceHarvester(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 3250,
      color: 'darkblue',
      filter: 'blur(17.2px)', // "feather"
      thickness: 7.2,
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

  xlOreHarvester(): ModuleActivationEffectData {
    return {
      type: 'laser',
      duration: 3975,
      color: 'darkred',
      filter: 'blur(11.3px)', // "feather"
      thickness: 5.25,
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

  xlShieldBooster(): ModuleActivationEffectData {
    return {
      type: 'bubble_shield_boost',
      duration: 725,
      color: '#A1FA61',
      filter: 'blur(25px)', // "feather"
      thickness: 15,
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

  xlAetherDragger(): ModuleActivationEffectData {
    return {
      type: 'aether_drag',
      duration: 350000,
      color: '#dbf4ff',
      filter: 'blur(21px)', // "feather"
      thickness: 15,
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

  xlUtilitySiphon(): ModuleActivationEffectData {
    return {
      type: 'siphon',
      duration: 2397,
      color: 'orange',
      filter: 'blur(45.16px)', // "feather"
      thickness: 15.3,
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
