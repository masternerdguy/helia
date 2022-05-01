import { IFactionLoreData } from './iFactionLoreData';

export class BadLoreData implements IFactionLoreData {
  factionName: string;
  factionTicker: string;
  factionDescription: string;

  constructor() {
    this.factionName = 'Bad Rabbits';
    this.factionTicker = '[-R-]';
    this.factionDescription = atob(this.rawDescription()).trim();
  }

  private rawDescription = () => {
    return 'QmVmb3JlIHRoZSBLaW5nZG9tIGVudGVyZWQgdGhlIHdhciBvbiB0aGUgc2lkZSBvZiB0aGUgRmVkZXJhdGlvbiwgdGhlIGZpbmFsIGhvbGQgb3V0IHN5c3RlbXMgb2YgdGhlIEZlZGVyYXRpb24gd2VyZSBmaWdodGluZyBkZXNwZXJhdGVseSBmb3IgdGhlaXIgZnJlZWRvbSBhZ2FpbnN0IHRoZSBBY2NvcmQgYW5kIENvYWxpdGlvbi4gRHVyaW5nIHRoZWlyIGRhcmtlc3QgaG91cnMsIG1pbGl0YXJ5IGxlYWRlcnMgd2l0aGluIHRoZSBGZWRlcmF0aW9uIGJlZ2FuIHRvIGJlbGlldmUgZGVmZWF0IHdhcyBpbmV2aXRhYmxlLiBBIG5ldyBzdHJhdGVneSB3YXMgbmVlZGVkLCB0aGV5IHJlYXNvbmVkLCB0byBzb21laG93IGxldmVsIHRoZSBwbGF5aW5nIGZpZWxkIHdpdGggdGhlaXIgZW5lbWllcy4KCkluIGRlc3BlcmF0aW9uLCB0aGUgRmVkZXJhdGlvbiBzZWNyZXRseSBjcmVhdGVkIGEgc3BlY2lhbCB0YXNrIGZvcmNlIHRvIHdhZ2UgdW5jb252ZW50aW9uYWwgd2FyZmFyZSwgYW5kIGV2ZW4gY29tbWl0IHRlcnJvcmlzdCBhdHRhY2tzLCBhZ2FpbnN0IHRoZWlyIGVuZW1pZXMuIFRha2luZyB0aGVpciBuYW1lIGZyb20gdGhlIGdyYWR1YWwgZXZvbHV0aW9uIG9mIGFuIGluc2lkZS1qb2tlIHdpdGhpbiB0aGUgRmVkZXJhdGlvbiBhcm1lZCBmb3JjZXMgKGxlbmd0aHksIGFuZCBpbnZvbHZpbmcgdGhlIGlkZWFzIG9mIEVhc3RlciBFZ2dzLCBiZWluZyBiYWQgZWdncywgYW5kIHRoZSBFYXN0ZXIgQnVubnkgZGVsaXZlcmluZyBzYWlkIGJhZCBlZ2dzKSB0aGUgIkJhZCBSYWJiaXRzIiB3b3VsZCBvZmZpY2lhbGx5IGJlIGRpc2NoYXJnZWQgZnJvbSB0aGUgYXJtZWQgZm9yY2VzIHRvIGFjdCBhcyBwcml2YXRlZXJzIC0gc3RyaWtpbmcgZGVlcGx5IGF0IHRoZSBzb2Z0IHVuZGVyYmVsbHkgb2YgdGhlaXIgZW5lbWllcyBhdCBhbnkgY29zdC4gT2ZmaWNpYWxseSwgdGhlIGFjdGlvbnMgb2YgdGhlIEJhZCBSYWJiaXRzIHdvdWxkIGJlIGRpc21pc3NlZCBieSB0aGUgRmVkZXJhdGlvbiBhcyB0aGUgYWN0cyBvZiByZW5lZ2FkZXMuCgpUaGUgQmFkIFJhYmJpdHMgd2VyZSB3ZWxsIGVxdWlwcGVkLCB3ZWxsIHRyYWluZWQsIGFuZCBnaXZlbiBhIGxpY2Vuc2UgdG8gZG8gYW55dGhpbmcgdGhleSB3YW50ZWQgdG8gdGhlIGVuZW15IC0gYXMgbG9uZyBhcyBpdCB3YXNuJ3QgaW4gdGhlIG5hbWUgb2YgdGhlIEZlZGVyYXRpb24gaXRzZWxmLiBUaGV5IGhpdCBhbnkgdGFyZ2V0IHRoYXQgd2FzIHZhbHVhYmxlLCBjaXZpbGlhbiBvciBtaWxpdGFyeS4gTWFueSBhdHRyb2NpdGllcyBhZ2FpbnN0IGNpdmlsaWFucyBpbiB0aGUgQWNjb3JkIGFuZCBDb2FsaXRpb24gYnkgdGhlIEJhZCBSYWJiaXRzLCB3aG8gdXN1YWxseSBnb3QgaW4sIGRpZCB0aGVpciBkZWVkcywgYW5kIGVzY2FwZWQgYmVmb3JlIGEgc2VyaW91cyBlbmVteSBjb3VudGVyLWF0dGFjayBjb3VsZCBiZSBtb3VudGVkLgoKVG8gc3RyaWtlIGRlZXAsIGFuZCBoYXJkLCB0aGUgQmFkIFJhYmJpdHMgdG9vayBmdWxsIGFkdmFudGFnZSBvZiBjbG9ha2luZyBkZXZpY2VzLCBlbmdpbmUgb3ZlcmNoYXJnZXJzLCBlbGVjdHJvbmljIHdhcmZhcmUsIGFuZCB0cmFuc2llbnQganVtcGhvbGUgY29ubmVjdGlvbnMuIEZyZWVkIGZyb20gdGhlIG92ZXJzaWdodCBvZiB0aGUgYXJtZWQgZm9yY2VzLCB0aGV5IG9mdGVuIHB1c2hlZCB0aGVpciBzaGlwcyB0byB0aGUgYWJzb2x1dGUgbGltaXQgLSBzb21ldGltZXMgYWxtb3N0IGJsb3dpbmcgdGhlbXNlbHZlcyB1cCB3aXRoIG5vIGhlbHAgZnJvbSB0aGUgZW5lbXkhIFRoZXkgd2VyZSBhbHNvIHdpbGxpbmcgdG8gdXNlIGNhcHR1cmVkIGVuZW15IHNoaXBzLCBjaXZpbGlhbiBvciBvdGhlcndpc2UsIHRvIHNldCB0cmFwcyBvciBwdXNoIGZhciBkZWVwZXIgaW50byBlbmVteSB0ZXJyaXRvcnkgdGhhbiB0aGV5IG90aGVyd2lzZSB3b3VsZCBiZSBhYmxlIHRvLgoKT25jZSB0aGUgS2luZ2RvbSBqb2luZWQgdGhlIHdhciwgYW5kIHRoZSBiYWxhbmNlIG9mIHBvd2VyIHdhcyBzaGlmdGVkLCB0aGUgRmVkZXJhdGlvbiBkZWNsYXJlZCB0aGUgQmFkIFJhYmJpdHMgdG8gYmUgb2Jzb2xldGUgYW5kIGF0dGVtcHRlZCB0byByZWNhbGwgdGhlbS4gSG93ZXZlciwgdGhlIGZpcnN0IHRvIGNhdXRpb3VzbHkgcmV0dXJuIGhvbWUgd2VyZSBwcm9tcHRseSBhcnJlc3RlZCBhbmQgY2hhcmdlZCB3aXRoIHBpcmFjeSBhbmQgd2FyIGNyaW1lcyBpbiBhbiBlZmZvcnQgdG8gbWFpbnRhaW4gdGhlIGdvb2QgaW1hZ2Ugb2YgdGhlIEZlZGVyYXRpb24uCgpLbm93aW5nIHdoYXQgd291bGQgaGFwcGVuIHRvIHRoZW0gaWYgdGhleSByZXR1cm5lZCBob21lLCB0aGUgb3RoZXIgQmFkIFJhYmJpdHMgZGVjaWRlZCB0byBjb250aW51ZSBvcGVyYXRpb25zIGFnYWluc3QgdGhlIEFjY29yZCBhbmQgQ29hbGl0aW9uLiBUaGV5IGhhZCBiZWNvbWUgdmVyeSBpbmRlcGVuZGVudCBmcm9tIHRoZSBGZWRlcmF0aW9uLCBhbmQgaGFkIGJlY29tZSBnb29kIGVub3VnaCBhdCByYWlkaW5nIHRvIGFjcXVpcmUgdGhlIGVxdWlwbWVudCBhbmQgc3VwcGxpZXMgdGhleSBuZWVkZWQgb24gdGhlaXIgb3duLiBUaGV5IHRydWx5IGRpZG4ndCBuZWVkIHRoZW0gYW55bW9yZS4KClRoZSBCYWQgUmFiYml0cyBjb250aW51ZWQgdG8gZmlnaHQsIGFuZCBncmFkdWFsbHkgbWFkZSB0aGUgdHJhbnNpdGlvbiB0byBwaXJhY3kgaW4gaXRzIG93biByaWdodC4gSW4gdGhlIHNodWZmbGUgb2YgdGhlIHdhciwgdGhleSBldmVuIG1hbmFnZWQgdG8gY2xhaW0gbWFueSBkZXZhc3RhdGVkIEZlZGVyYXRpb24gc3lzdGVtcyBhcyB0aGVpciBvd24uIEFsdGhvdWdoIHRoZSBGZWRlcmF0aW9uICh3aXRoIG5vIGhlbHAgZnJvbSB0aGUgS2luZ2RvbSB3aG8gaGFkIG5vIGludGVyZXN0IGluIHJlaWduaW5nIGluIHRoZWlyIHJlbmVnYWRlcykgYXR0ZW1wdGVkIHRvIHJlY2xhaW0gdGhvc2Ugc3lzdGVtcywgdGhleSBiZWNhbWUgdGhlIHZpY3RpbXMgb2YgdGhlIHZlcnkgc2FtZSBkaXNob25vdXJhYmxlIHRhY3RpY3MgdGhleSBoYWQgb25jZSB1bmxlYXNoZWQgdXBvbiB0aGVpciBlbmVtaWVzLgoKVG8gdGhpcyBkYXksIHRoZSBCYWQgUmFiYml0cyBtYWludGFpbiBtYW55IGFzcGVjdHMgb2YgbWlsaXRhcnkgY3VsdHVyZSBhbmQgZGlzY2lwbGluZS4gVGhleSBwbGFuIHRoZWlyIG9wZXJhdGlvbnMgd2l0aCBhbGwgdGhlIHByZWNpc2lvbiBvZiBhIHJlYWwgbWlsaXRhcnkgb3BlcmF0aW9uLCBhbmQgdGhleSBjb250aW51ZSB0byByZWNydWl0IG5ldyBtZW1iZXJzIGZyb20gdGhlIG1pbGl0YXJpZXMgb2YgYWxsIGZvdXIgb2YgdGhlIGdyZWF0IGVtcGlyZXMuIFRoZWlyIHJlY3J1aXRzIGFyZSBnZW5lcmFsbHkgZGlzaWxsdXNpb25lZCB3aXRoIHRoZWlyIG93biBuYXRpb24sIG9yIG1vcmUgb2Z0ZW4ganVzdCBpdHMgY3VycmVudCBsZWFkZXJzaGlwLCBhbmQgYXJlIGxvb2tpbmcgZm9yIGEgY2hhbmdlIG9mIHBhY2UgYW5kIHNlbnNlIG9mIHB1cnBvc2UuCgpUaGV5IGFyZSBjdXJyZW50bHkgZW5nYWdlZCBpbiBhIHR1cmYtd2FyIHdpdGggdGhlaXIgcml2YWxzIFBhcHBhIEZseSBSZWxvYWRlZCwgd2hvbSB0aGV5IGRvbid0IHNlZSBleWUgdG8gZXllIHdpdGggZHVlIHRvIGN1bHR1cmFsIGRpZmZlcmVuY2VzLiBJbiBhZGRpdGlvbiwgdGhleSBoYXZlIGJlZW4ga25vd24gdG8gYWNjZXB0IHByaXZhdGVlcmluZyBjb250cmFjdHMgZnJvbSBhIHdpZGUgdmFyaWV0eSBvZiBpbmRpdmlkdWFscyBhbmQgb3JnYW5pemF0aW9ucy4K';
  };
}