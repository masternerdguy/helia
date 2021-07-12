import { Faction } from '../engineModels/faction';

export class CurrentSystemInfo {
  id: string;
  systemId: string;
  shipName: string;
  uid: string;
  x: number;
  y: number;
}

/* Helpers for Window globals */

export function UpdateFactionCache(factions: Faction[]) {
  if (!(window as any).factionDictionary) {
    (window as any).factionDictionary = {};
  }

  for (const f of factions) {
    (window as any).factionDictionary[f.id] = new Faction(f);
  }
}

export function GetFactionCache(): Faction[] {
  const cache: Faction[] = [];

  if (!(window as any).factionDictionary) {
    return cache;
  }

  for (const f in (window as any).factionDictionary) {
    if (
      Object.prototype.hasOwnProperty.call((window as any).factionDictionary, f)
    ) {
      const e = (window as any).factionDictionary[f] as Faction;
      cache.push(e);
    }
  }

  return cache;
}

export function GetFactionCacheEntry(id: string): Faction {
  if (!(window as any).factionDictionary) {
    return null;
  }

  return (window as any).factionDictionary[id] as Faction;
}
