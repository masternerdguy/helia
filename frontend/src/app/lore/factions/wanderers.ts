import { IFactionLoreData } from './iFactionLoreData';

export class WanderersLoreData implements IFactionLoreData {
  factionName: string;
  factionTicker: string;
  factionDescription: string;

  constructor() {
    this.factionName = 'Wanderers';
    this.factionTicker = '[WHY]';
    this.factionDescription = atob(this.rawDescription()).trim();
  }

  private rawDescription = () => {
    return 'QSBzbWFsbCBiYW5kIG9mIHBpbG90cyB3aG8gcm9hbSB0aGUgdW5pdmVyc2Ugc2VlbWluZ2x5IGZvciBubyBvdGhlciByZWFzb24gYnV0IHRvIGRvIHNvLCBXYW5kZXJlcnMga2VlcCB0byB0aGVtc2VsdmVzIGFzIHRoZXkgcm9hbSB0aGUgc3RhcnMuIEFzIHN1Y2gsIHRoZXkgYXJlIHVzdWFsbHkgaWdub3JlZCwgb3IgbG9va2VkIHVwb24gYXMgYSBjdXJpb3NpdHksIGJ5IG90aGVyIGZhY3Rpb25zLg==';
  };
}
