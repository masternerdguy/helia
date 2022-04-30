import { IFactionLoreData } from './iFactionLoreData';

export class KingdomLoreData implements IFactionLoreData {
  factionName: string;
  factionTicker: string;
  factionDescription: string;

  constructor() {
    this.factionName = 'Kingdom of Antaria';
    this.factionTicker = '[KoA]';
    this.factionDescription = atob(this.rawDescription()).trim();
  }

  private rawDescription = () => {
    return 'VGhlIEtpbmdkb20gd2FzIHRoZSBmaXJzdCBvZiB0aGUgZ3JlYXQgbW9kZXJuIG5hdGlvbnMgdG8gZW1lcmdlLCBhbmQgdGhlaXIgZW1lcmdlbmNlIGlzIGdlbmVyYWxseSBjb25zaWRlcmVkIHRvIG1hcmsgdGhlIGVuZCBvZiB0aGUgc3ludGhldGljIGRhcmsgYWdlLiBUaGUgS2luZ2RvbSBpcyBjb21wb3NlZCBvZiB3aGF0IHdhcyBvbmNlIHNldmVyYWwgZHlzZnVuY3Rpb25hbCByZXB1YmxpY3MgZmlnaHRpbmcgZm9yIGRvbWluYW5jZSB3aGlsZSBhbHNvIGRlc2NlbmRpbmcgaW50byBpbmNyZWFzaW5nIGFtb3VudHMgb2YgaGVkb25pc20uIEV2ZW50dWFsbHksIHRoZXJlIHdhcyBhIHBvcHVsYXIgcmV2b2x0IGluIHRoZSBsYXJnZXN0IG9mIHRoZSByZXB1YmxpY3Mgd2hpY2ggcmVzdWx0ZWQgaW4gbWFueSBvZiB0aGVpciB1cHBlciBjbGFzcyBiZWluZyBlbGV2YXRlZCB0byBhIHN0YXRlIG9mIG5vYmlsaXR5LiBPbmNlIGluIHBvd2VyLCB0aGUgbmV3IG5vYmlsaXR5IGltbWVkaWF0ZWx5IG9yZ2FuaXplZCBpdHNlbGYgaW50byBhIGhpZXJhcmNoaWNhbCBjb21tYW5kIHN0cnVjdHVyZSB3aXRoIGEgS2luZyBhbmQgUXVlZW4gYWN0aW5nIGFzIGEgcG93ZXItY291cGxlIGF0IHRoZSB0b3AsIGJhc2VkIG9uIHNvbWUgc3Vydml2aW5nIHJlY29yZHMgZnJvbSBodW1hbiBjaXZpbGl6YXRpb24uCgpUaGUgZmlyc3Qgcm95YWwgcG93ZXItY291cGxlIHdlcmUgS2luZyBhbmQgUXVlZW4gQXZlcnktQ2hhcmxldHRlLCB3aG8gdG9nZXRoZXIgZGVjaWRlZCB0byBzdG9wIGZpZ2h0aW5nIHRoZWlyIG5laWdoYm91cnMgYW5kIGZvY3VzIG9uIGlud2FyZCBwZXJmZWN0aW9uLiBUaGV5IGNsb3NlZCB0aGVpciBib3JkZXJzIGxlYXZpbmcgdGhlIHJlcHVibGljcyB0byBmaWdodCBhbW9uZyB0aGVtc2VsdmVzLCBhbmQgZm9jdXNlZCBvbiByZXN0b3JpbmcgdHJhZGl0aW9uYWwgbW9yYWxpdHkgYW5kIG1vZGVybml6aW5nIGluZnJhc3RydWN0dXJlLiBUaGlzIHJlc3VsdGVkIGluIGEgR3JlYXQgUmV2aXZhbCBvZiB0aGVpciBzb2NpZXR5IGFuZCBhZnRlciBvbmx5IGEgZmV3IGRlY2FkZXMgdGhlIEtpbmdkb20gd2FzIGZhciBtb3JlIHBvd2VyZnVsIGFuZCBzdGFibGUgdGhhbiBhbGwgb2YgdGhlaXIgZm9ybWVyIGZvZXMgY29tYmluZWQgLSB3aGVuIHRoZSBLaW5nZG9tIGZpbmFsbHkgb3BlbmVkIHRoZWlyIGJvcmRlcnMsIHRoZXkgZWFzaWx5IHN1YnN1bWVkIHRoZW0uCgpGcmVlZCBvZiBvbGQgY29uZmxpY3RzLCB0aGUgS2luZ2RvbSBiZWdhbiB0byBleHBsb3JlIGFuZCBleHBhbmQuIEFzIHRoZXkgZXhwbG9yZWQsIHRoZXkgc3R1bWJsZWQgdXBvbiBhIHNpZ25pZmljYW50bHkgbGVzcyBhZHZhbmNlZCBuYXRpb24ga25vd24gYXMgdGhlIFZpZXJyYSBGZWRlcmF0aW9uIC0gYSBjb2xsZWN0aW9uIG9mIHN5c3RlbS1zdGF0ZXMgd2hvIHdlcmUgZmlnaHRpbmcgZm9yIHRoZWlyIGluZGVwZW5kZW5jZSBmcm9tIHRoZSBtb3JlIHBvd2VyZnVsIE96b3VrYSBBY2NvcmQuIEFmdGVyIGRpc2FzdHJvdXMgZmlyc3QgY29udGFjdCB3aXRoIHRoZSBPem91a2EgQWNjb3JkLCB3aG8gYWNjdXNlZCB0aGUgS2luZ2RvbSBvZiBzZWNyZXRseSBhcm1pbmcgdGhlIFZpZXJyYW4gUmViZWxzIGFuZCB0aHJlYXRlbmVkIHJldGFsaWF0aW9uLCB0aGV5IGRlY2lkZWQgdG8gcGxheSB0aGUgcGFydCBhbmQgc3VwcG9ydCB0aGUgVmllcnJhIEZlZGVyYXRpb24ncyBxdWVzdCBmb3IgaW5kZXBlbmRlbmNlLiBUZWNobmljYWxseSwgdGhpcyBjb25mbGljdCBoYXMgcmFnZWQgdG8gdGhpcyBkYXkgd2l0aCBvcGVuIGhvc3RpbGl0aWVzIGRlY2xhcmVkIGJldHdlZW4gYm90aCBzaWRlcy4gSG93ZXZlciwgZ2l2ZW4gdGhhdCBtb3N0IHN5bnRoZXRpY3MgdG9kYXkgYXJlIGZhc2NpbmF0ZWQgd2l0aCBsb25nZXZpdHksIG5laXRoZXIgc2lkZSBpcyBjdXJyZW50bHkgZW5nYWdlZCBpbiBhIGNhbXBhaWduIGFnYWluc3QgdGhlIG90aGVyIGFuZCBob3N0aWxpdGllcyBhcmUgbGltaXRlZCB0byBvbmUtb2ZmIHNraXJtaXNoZXMgYmV0d2VlbiBzbWFsbCBncm91cHMgb2YgcGlsb3RzLgoKVG9kYXksIHRoZSBLaW5nZG9tIGlzIHJ1bGVkIGJ5IEtpbmcgYW5kIFF1ZWVuIFBlYXJjZS1BbWVsaWEsIHRoZSBmaWZ0aCBpbiB0aGUgcnVsaW5nIGxpbmUuIE1vc3Qgcm95YWwgcG93ZXIgdHJhbnNpdGlvbnMgb2NjdXJyZWQgZWFybHkgLSBiZWZvcmUgdGhlIHJlLWRpc2NvdmVyeSBvZiBzeW50aGV0aWMgcmVwYWlyIHRlY2huaXF1ZXMgbWFkZSB0aGUgd2VsbC1wcm90ZWN0ZWQgY291cGxlIHZ1bG5lcmFibGUgb25seSB0byBmcmVhayBhY2NpZGVudCBvciBhc3Nhc3NpbmF0aW9uLiBUaGUgbGFzdCB0cmFuc2l0aW9uIG9mIHBvd2VyIHdhcyBvdmVyIDMwMCB5ZWFycyBhZ28gd2hlbiBLaW5nIGFuZCBRdWVlbiBLb2xlLUxpbmRhIGRpZWQgd2hpbGUgcmVjaGFyZ2luZyBpbiBhIGZyZWFrIHBvd2VyIHN1cmdlLCB0aGUgY2F1c2Ugb2Ygd2hpY2ggaGFzIG9idmlvdXNseSBiZWVuIGFkZHJlc3NlZC4KCkZvbGxvd2luZyBpbiBhbmNpZW50IHRyYWRpdGlvbiwgdGhlIEtpbmdkb20gYXNzZXJ0cyB0aGF0IHRoZWlyIEtpbmcgYW5kIFF1ZWVuIHJ1bGUgYnkgZGl2aW5lIG1hbmRhdGUuIFRoZWlyIGNpdGl6ZW5zIHRlbmQgdG8gYmUgdmVyeSByZWxpZ2lvdXMgYW5kIHRyYWRpdGlvbmFsLCBhbmQgYmVsaWV2ZSB0aGF0IGFsdGhvdWdoIHRoZXkgaGF2ZSBubyBob3BlIG9mIGV2ZXJsYXN0aW5nIGxpZmUsIG5vciBwb3NzaWJpbGl0eSBvZiBldGVybmFsIGRhbW5hdGlvbiwgdGhhdCB0aGV5IHdvdWxkIGJlIGRvaW5nIHRoZW1zZWx2ZXMgYW4gZW5vcm1vdXMgZGlzc2VydmljZSBieSBub3QgY3VsdGl2YXRpbmcgYSByZWxhdGlvbnNoaXAgd2l0aCB0aGUgQ3JlYXRvciBvZiB0aGUgY3JlYXRvcnMuIEJlc2lkZXMgdGhlaXIgZm9jdXMgb24gdGhlIGRpdmluZSwgbWFueSBob3BlIHRvIG9uZSBkYXkgbWVldCB0aGUgaHVtYW5zIHdobyBjcmVhdGVkIHRoZWlyIHJhY2UsIHdob20gdGhleSBnZW5lcmFsbHkgY29uc2lkZXIgc3VwZXJpb3IgdG8gdGhlbXNlbHZlcy4gQWx0aG91Z2ggaW1wZXJmZWN0IGJlaW5ncywgdGhleSB1c3VhbGx5IHN0cml2ZSB0byBiZSBnb29kIG9uZXMgYW5kIGF0dGVtcHQgdG8gZm9sbG93IHRoZSBwcmVjZXB0cyBvYnNlcnZlZCBieSB0aGVpciBjcmVhdG9ycy4=';
  };
}
