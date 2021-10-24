import { Faction } from '../engineModels/faction';
import { WSPlayerFactionRelationship } from './entities/wsPlayerFaction';

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
    f.relationships = f.relationships.sort(
      (a, b) => b.standingValue - a.standingValue
    );
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

  return cache.sort((a, b) => (a.name ?? '').localeCompare(b.name ?? ''));
}

export function GetFactionCacheEntry(id: string): Faction {
  if (!(window as any).factionDictionary) {
    return null;
  }

  return (window as any).factionDictionary[id] as Faction;
}

export function UpdatePlayerFactionRelationshipCache(
  playerFactions: WSPlayerFactionRelationship[]
) {
  if (!(window as any).playerFactionDictionary) {
    (window as any).playerFactionDictionary = {};
  }

  for (const f of playerFactions) {
    (window as any).playerFactionDictionary[f.factionId] = f;
  }
}

export function GetPlayerFactionRelationshipCache(): WSPlayerFactionRelationship[] {
  const cache: WSPlayerFactionRelationship[] = [];

  if (!(window as any).playerFactionDictionary) {
    return cache;
  }

  for (const f in (window as any).playerFactionDictionary) {
    if (
      Object.prototype.hasOwnProperty.call(
        (window as any).playerFactionDictionary,
        f
      )
    ) {
      const e = (window as any).playerFactionDictionary[
        f
      ] as WSPlayerFactionRelationship;
      cache.push(e);
    }
  }

  return cache.sort((a, b) => b.standingValue - a.standingValue);
}

export function GetPlayerFactionRelationshipCacheEntry(
  id: string
): WSPlayerFactionRelationship {
  if (!(window as any).playerFactionDictionary) {
    return null;
  }

  return (window as any).playerFactionDictionary[
    id
  ] as WSPlayerFactionRelationship;
}
