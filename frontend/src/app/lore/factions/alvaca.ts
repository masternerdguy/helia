import { IFactionLoreData } from './iFactionLoreData';

export class AlvacaLoreData implements IFactionLoreData {
  factionName: string;
  factionTicker: string;
  factionDescription: string;

  constructor() {
    this.factionName = 'Alvaca';
    this.factionTicker = '[AV.]';
    this.factionDescription = atob(this.rawDescription()).trim();
  }

  private rawDescription = () => {
    return 'T25jZSB0aGUgZm91ciBncmVhdCBuYXRpb25zIGhhZCByZWFjaGVkIGEgc3RhbGVtYXRlLCB0aGV5IGJlZ2FuIHRvIHBvbmRlciBvdGhlciBtYXR0ZXJzIHdpdGhpbiB0aGVpciBkb21haW5zLiBPZiBwYXJ0aWN1bGFyIGNvbmNlcm4gd2VyZSB0aGUgYWJhbmRvbmVkIHRlcnJhZm9ybWluZyBwcm9qZWN0cyBUaGUgT3JpZ2luYWxzIGhhZCBzdGFydGVkIGJlZm9yZSB0aGUgaHVtYW4gY3JlYXRvcnMgaGFuZGVkIGNvbnRyb2wgb3ZlciB0byB0aGVpciBoZWxwZXJzIGFuZCBkaXNhcHBlYXJlZC4gVGhlc2UgcHJvamVjdHMgd2VyZSBhdCBhbGwgc3RhZ2VzIG9mIGNvbXBsZXRpb24gLSBmcm9tIGJhcmVseSBicmVha2luZyBncm91bmQgdG8gbmVhcmx5IHRyYW5zZm9ybWVkIGJpb3NwaGVyZXMgc2ltcGx5IGF3YWl0aW5nIHNvbWUgYW5pbWFsIHNwZWNpZXMgcHJpemVkIGJ5IHRoZSBjcmVhdG9ycy4KClRoZXJlIHdlcmUgbWFueSBvcGluaW9ucyByZWdhcmRpbmcgd2hhdCB0byBkbyB3aXRoIHRoZXNlIHdvcmxkcywgYW5kIHRoZSBpc3N1ZSBkaXZpZGVkIHRoZSBzeW50aGV0aWMgbmF0aW9ucyBldmVuIHdpdGhpbiB0aGVtc2VsdmVzLiBPbiBvbmUgaGFuZCwgaXQgd2FzIGFyZ3VlZCB0aGF0IHRoZXNlIHdvcmxkcyB3ZXJlIGNsZWFybHkgaW50ZW5kZWQgdG8gYmUgdHJhbnNmb3JtZWQsIGFuZCBjb21wbGV0aW5nIHRoZSBwcm9jZXNzIHdvdWxkIGluY3JlYXNlIHRoZSBhbW91bnQgb2YgYXZhaWxhYmxlIGxpdmluZyBzcGFjZSBmb3Igc3ludGhldGljcy4gT24gdGhlIG90aGVyIGhhbmQsIG90aGVycyBhcmd1ZWQgdGhhdCB0aGUgcHJvamVjdHMgaGFkIGJlZW4gaGFsdGVkIGZvciBhbiB1bmtub3duIHJlYXNvbiBieSB0aGUgY3JlYXRvcnMsIGFuZCB0aGF0IHN5bnRoZXRpY3Mgd291bGQgYmUgcGxheWluZyBHb2QgYnkgcmVzdW1pbmcgdGhlbS4gQSBtaW5vcml0eSB2aWV3cG9pbnQgYWxzbyBhcm9zZSB0aGF0IHZvaWNlZCBjb25jZXJucyB0aGF0IHRoZSBpZGVhIG9mIHRlcnJhZm9ybWluZyB3YXMgZnVuZGFtZW50YWxseSBldmlsIGJlY2F1c2UgaXQgaW52b2x2ZWQgZGVzdHJveWluZyB0aGUgbmF0dXJhbCBzdGF0ZSBvZiB3b3JsZHMgaW4gb3JkZXIgdG8gY29udmVydCB0aGVtIGludG8gc29tZXRoaW5nIGNvbXBsZXRlbHkgZGlmZmVyZW50LgoKRGVzcGl0ZSB0aGVzZSBleGlzdGVudGlhbCBjb25jZXJucywgbWFueSBlbnRpdGllcywgYm90aCBnb3Zlcm5tZW50IGFuZCBidXNpbmVzcyBvcmdhbml6YXRpb25zLCBzYXcgbWF0ZXJpYWwgdmFsdWUgaW4gcmVzdW1pbmcgYW5kIGNvbXBsZXRpbmcgdGhlc2UgcHJvamVjdHMgLSBhIHRyYW5zZm9ybWVkIGJpb3NwaGVyZSB3b3VsZCBhbGxvdyBzeW50aGV0aWNzIHRvIG9wZXJhdGUgdGhlcmUganVzdCBhcyBmcmVlbHkgYXMgaHVtYW5zIHdvdWxkLiBBbHRob3VnaCB0aGVpciBzcGVjaWZpYyBpbnRlcmVzdHMgdmFyaWVkLCB0aGVyZSB3ZXJlIG1hbnkgd2hvIHdhbnRlZCB0byBzZWUgdGhlc2Ugd29ybGRzIHRyYW5zZm9ybWVkIGluIG9yZGVyIHRvIGNvbnZlbmllbnRseSB1dGlsaXplIHRoZWlyIHJlc291cmNlcy4KCkV2ZW50dWFsbHksIHRoaXMgZGVzaXJlIHJlYWNoZWQgYSBwb2ludCB3aGVyZSBpdCBjb3VsZCBubyBsb25nZXIgYmUgc3VwcHJlc3NlZCBieSBydWxlcnMuIEV4Y2VwdGlvbmFsbHkgd2VhbHRoeSBpbmRpdmlkdWFscyBiZWdhbiB0byBhY3F1aXJlIHZpcnR1YWxseSBjb21wbGV0ZSBjb250cm9sIG9mIHNvbWUgb2YgdGhlc2Ugd29ybGRzLCBldmVudHVhbGx5IGV4cGFuZGluZyB0aGF0IGNvbnRyb2wgdG8gdGhlIGVudGlyZSBzeXN0ZW0uIExhcmdlIG1lZ2EtY29ycG9yYXRpb25zLCBhbHNvIGV4Y2VwdGlvbmFsbHkgd2VhbHRoeSBidXQgZ3VpZGVkIGJ5IGEgYm9hcmQgb2YgZGlyZWN0b3JzIGluc3RlYWQgb2YgYSBzaW5nbGUgaW5kaXZpZHVhbCwgZGlkIHRoZSBzYW1lLiBUaGUgd29ybGRzIGFuZCBzeXN0ZW1zIGNsYWltZWQgd2VyZSBpbnZhcmlhYmx5IGZyb250aWVycyBpbiB3aGljaCB0aGUgZ292ZXJubWVudHMgdGhhdCBub21pbmFsbHkgY29udHJvbGxlZCB0aGVtIGhhZCBkaWZmaWN1bHR5IGVuZm9yY2luZyB0aGVpciBydWxlcy4KCkFsdGhvdWdoIHRoZSBsZWFkZXJzIG9mIHRoZSBncmVhdCBuYXRpb25zIGNvbnRpbnVlZCB0byBhc3NlcnQgdGhleSB3ZXJlIGluIGNvbnRyb2wgb2YgdGhlc2Ugd29ybGRzLCBpdCB3YXMgYmVjb21pbmcgcGFpbmZ1bGx5IG9idmlvdXMgdG8gdGhlbSB0aGF0IGl0IHdhcyBpbiBuYW1lIG9ubHkuIEZhY2luZyB0aGUgcHJvc3BlY3Qgb2YgZmlnaHRpbmcgY2l2aWwgd2FycyBvdmVyIHRoZXNlIGZyb250aWVyIHJlZ2lvbnMsIGFuZCB0aGVpciBpbmNvbXBsZXRlIHRlcnJhZm9ybWluZyBwcm9qZWN0cywgdGhlaXIgbGVhZGVycyBiZWdhbiB0byBncmFudCBjaGFydGVycyB0byB0aGVzZSBlbnRpdGllcyB0aGF0IGxlZ2FsbHkgcGVybWl0dGVkIHRoZW0gdG8gcnVsZSB0aGVpciBjbGFpbWVkIHRlcnJpdG9yaWVzIC0gbm9taW5hbGx5IGFzIHNwZWNpYWwgZGlzdHJpY3RzIG9mIHRoZSBuYXRpb24uIEZldyB3YW50ZWQgdG8gcmlzayBzaW11bHRhbmVvdXMgY2l2aWwgd2FycyB0aGF0IGNvdWxkIGRlc3RhYmlsaXplIHRoZSBmcmFnaWxlIHBvd2VyIGJhbGFuY2UuCgpEZXNwaXRlIHRoZWlyIG5ld2ZvdW5kIHBvbGl0aWNhbCByZWNvZ25pdGlvbiwgdGhlIGluZGl2aWR1YWxzIGFuZCBjb3Jwb3JhdGlvbnMgY29udHJvbGxpbmcgdGhlc2Ugc3lzdGVtcyBoYWQgYW5vdGhlciBwcm9ibGVtIC0gYWx0aG91Z2ggdGhleSBoYWQgaW1tZW5zZSByZXNvdXJjZXMsIGl0IHdhcyBiZWNvbWluZyBvYnZpb3VzIHRoZXkgZGlkbid0IGhhdmUgdGhlIGtub3ctaG93IHRvIGNvbXBsZXRlIHRoZXNlIHByb2plY3RzLiBJdCB3YXMgZGV0ZXJtaW5lZCB0aGF0IGNvbXBsZXRpbmcgdGhlbSB3b3VsZCBpbnZvbHZlIHN0dWR5aW5nIHRoZSB2YXJpb3VzIHN0YWdlcyBvZiBjb21wbGV0aW9uIGNvb3BlcmF0aXZlbHkuCgpBZnRlciBtYW55IG1lZXRpbmdzLCBhIGxlYWd1ZSBvZiBzY2llbnRpc3RzIHdhcyBjcmVhdGVkLCByZWNydWl0aW5nIHRoZSBiZXN0IG9mIHRoZSBiZXN0IGluIHRoZSB0ZXJyYWZvcm1pbmcgZmllbGQgZnJvbSBhbGwgbmF0aW9ucy4gSXQgd2FzIGRlY2lkZWQgdGhhdCB0aGVzZSBzY2llbnRpc3RzIHdvdWxkIGJlIGdpdmVuIHNwZWNpYWwgcHJpdmlsZWdlcyB3aXRoaW4gdGhlIHN5c3RlbXMgdG8gYmUgdGVycmFmb3JtZWQgLSBhbGxvd2luZyB0aGVtIHVucmVzdHJpY3RlZCBhY2Nlc3MgdG8sIGFuZCB0b3RhbCBjb250cm9sIG9mLCB0aGUgdGVycmFmb3JtaW5nIHByb2plY3RzLiBUaGlzLCBpbiB0aGVvcnksIHdvdWxkIGFsbG93IHVucmVzdHJpY3RlZCBmbG93IG9mIGtub3dsZWRnZSBhbmQgdGhlIGFiaWxpdHkgdG8gZWZmaWNpZW50bHkgYXBwbHkgaXQuIFRoZSBpbmRpdmlkdWFscyBhbmQgY29ycG9yYXRpb25zIHRoYXQgbWFkZSB0aGlzIHBhY3QgaG9wZWQgaXQgd291bGQgYmUgdGhlIHF1aWNrZXN0IHBhdGggdG8gdmljdG9yeSBhbmQgdGhhdCB0aGVpciB3b3JsZHMgd291bGQgYmUgdGVycmFmb3JtZWQgd2l0aGluIG9ubHkgYSBmZXcgZ2VuZXJhdGlvbnMuCgpPbmUgZmluYWwgcGllY2Ugd2FzIG5lZWRlZCAtIGEgbmFtZS4gSXQgd2FzIGRlY2lkZWQgdGhhdCB0aGlzIHRlY2hub2NyYXRpYyBjb2FsaXRpb24gYmV0d2VlbiB0ZXJyYWZvcm1lcnMgd291bGQgYmUgY2FsbGVkICJBbHZhY2EiLCByZWZlcnJpbmcgdG8gdGhlIG1haW4gY2hhcmFjdGVyIG9mIGFuIGFuY2llbnQgT3pvdWthbiBmb2xrIGxlZ2VuZCBpbiB3aGljaCBhIHdvbGYgc3BlbmRzIGhlciBlbnRpcmUgbGlmZSBicmluZ2luZyB3YXRlciB0byBhIGRlc2VydCBhbmQgZGllcyBvbmx5IHdoZW4gc2hlIHNlZXMgaXQgeWllbGQgYSBiZWF1dGlmdWwgZmllbGQgb2YgZmxvd2Vycy4KCk11Y2ggbGlrZSB0aGUgd29sZiBpbiB0aGUgbGVnZW5kLCBpdCB3YXMgZXhwZWN0ZWQgdGhhdCBBbHZhY2EgaXRzZWxmIHdvdWxkIGRpc3NvbHZlIG9uY2UgdGhlIHRlcnJhZm9ybWluZyBwcm9qZWN0cyB3ZXJlIGNvbXBsZXRlZC4gSG93ZXZlciwgdGhpcyBzZWVtcyBsZXNzIGFuZCBsZXNzIGxpa2VseSBhcyB0aGUgeWVhcnMgZ28gb24uIFNpbmNlIHRoZW4sIG11Y2ggbGlrZSB0aGUgd29sZiwgQWx2YWNhIGhhcyBncm93biBxdWl0ZSBzdHJvbmcgYW5kIGZpZXJjZS4gRGVzcGl0ZSB0aGUgdGVycmFmb3JtaW5nIHByb2plY3RzIGZhbGxpbmcgYmVoaW5kIHNjaGVkdWxlIG92ZXIgYW5kIG92ZXIsIHRoZSB0ZWNobm9jcmFjeSBoYXMgYmVjb21lIGluY3JlYXNpbmdseSBwb3dlcmZ1bCBhbmQgbm93IGV4ZXJ0cyBjb250cm9sIG92ZXIgbWFueSBhc3BlY3RzIG9mIGRheS10by1kYXkgbGlmZSB3aXRoaW4gQWx2YWNhJ3MgZG9tYWluIChhbGwgZm9yIHRoZSBwdXJwb3NlIG9mIGltcHJvdmluZyB0ZXJyYWZvcm1pbmcgZWZmaWNpZW5jeSwgb2YgY291cnNlKS4gSG93ZXZlciwgZGVzcGl0ZSB0aGVpciBzbG93IHByb2dyZXNzLCB0aGV5IGhhdmUgc3VjY2VlZGVkIGluIGNvbXBsZXRpbmcgc2V2ZXJhbCB0ZXJyYWZvcm1pbmcgcHJvamVjdHMgLSBhbHRob3VnaCBhbGwgb2YgdGhlbSB3ZXJlIGluIGZhaXJseSBsYXRlIHN0YWdlcyB0byBiZWdpbiB3aXRoLgoKTGlmZSBpbiBBbHZhY2EgaXMgZm9jdXNlZCB1bHRpbWF0ZWx5IG9uIHRlcnJhZm9ybWluZyAtIGFsbCBpbmR1c3RyaWVzLCBhdCBhbGwgc2NhbGVzLCB1bHRpbWF0ZWx5IGZlZWQgaW50byB0aGUgY29tcGxldGlvbiBvZiB0aGVzZSBwcm9qZWN0cy4gRXZlbiB0aGUgZW50ZXJ0YWlubWVudCBpbmR1c3RyeSBpcyBzZWVuIGFzIGFkdmFuY2luZyB0aGUgdGVycmFmb3JtaW5nIGRyZWFtLiBFdmVyeXRoaW5nIGlzIGNvbG9yZWQgd2l0aCBhIHZlbmVlciBvZiBuYXR1cmFsaXNtIGFuZCBlbnZpcm9ubWVudGFsaXNtLCB3aXRoIHRoZSBpZGVhIG9mIGNvbnNlcnZpbmcgdGhlIGFuaW1hbHMgcHJpemVkIGJ5IHRoZSBjcmVhdG9ycyBvZnRlbiBiZWluZyBpbnZva2VkLiBIb3dldmVyLCBhIHJlbGlnaW9uIG9mICJuYXR1cmFsbmVzcyIgaGFzIGdyYWR1YWxseSBiZWVuIGRldmVsb3BpbmcgdGhhdCByZXF1aXJlcyBubyBjcmVhdG9ycywgYW5kIG1ha2VzIGlkZWFsaXN0aWMgdmlzaW9ucyBvZiBuYXR1cmUgYW4gZW5kIHVudG8gdGhlbXNlbHZlcy4gQXMgb25lIG1pZ2h0IGltYWdpbmUsIHNlZWluZyBBbHZhY2EgYmVjb21lIG1vcmUgYW5kIG1vcmUgcGVjdWxpYXIgaGFzIG1hZGUgc29tZSBvZiB0aGUgZ3JlYXQgbmF0aW9ucyBuZXJ2b3VzLgoKVGhlc2UgZGF5cywgQWx2YWNhIGdldHMgYWxvbmcgZmFpcmx5IHdlbGwgd2l0aCB0aGUgVGVuZXZhbiBDb2FsaXRpb24gYW5kIFZpZXJyYSBGZWRlcmF0aW9uICh3aG8gc3RpbGwgaG9wZSB0byByZWFsaXplIHRhbmdpYmxlIGJlbmVmaXRzIGZyb20gdGhlIGNvbXBsZXRlZCBwcm9qZWN0cykuIEhvd2V2ZXIsIHRoZXkgYXJlIGRpc2xpa2VkIGJ5IHRoZSBLaW5nZG9tIG9mIEFudGFyaWEgYW5kLCBpcm9uaWNhbGx5LCB0aGUgT3pvdWthIEFjY29yZC4gU2VlaW5nIHRoZSBncm93dGggb2YgdGhlICJyZWxpZ2lvbiBvZiBuYXR1cmFsbmVzcyIsIGFuZCB0aGUgaW5jcmVhc2luZ2x5IGV4dHJlbWUgbWVhbnMgYnkgd2hpY2ggaXQgaXMgcHVyc3VlZCwgaGFzIG1hZGUgYm90aCB0aGUgS2luZ2RvbSBhbmQgQWNjb3JkIHZlcnkgbmVydm91cy4gVGhlIEtpbmdkb20gaXMgdGhlIG1vc3QgY29uY2VybmVkIG9mIGFsbCwgYW5kIHRoZWlyIHJveWFsdHkgZnJlcXVlbnRseSBjb21wYXJlIEFsdmFjYSB0byB0aGUgYW5jaWVudCBUb3dlciBvZiBCYWJlbCAod2VsbCBrbm93biB0byB0aGVtIGZyb20gdGhlIHJlY29yZHMgb2YgdGhlIGNyZWF0b3JzKS4KCkluIG9yZGVyIHRvIGd1YXJhbnRlZSB0aGVpciBzZWN1cml0eSBpbiB0aGUgZXZlbnQgdGhhdCB0aGVpciBjaGFydGVycyBhcmUgcmV2b2tlZCwgd2hpY2ggd291bGQgbGlrZWx5IGJlIHRoZSBvcGVuaW5nIHBoYXNlIG9mIGEgd2FyIGFnYWluc3QgdGhlbSwgQWx2YWNhIGhhcyBkZXZlbG9wZWQgYSBjb3p5IHJlbGF0aW9uc2hpcCB3aXRoIFBhcHBhIEZseSBSZWxvYWRlZCwgd2hvbSB0aGV5IHByb3ZpZGUgc3RhZ2luZyBmb3IuIEFsdGhvdWdoIG5vdCBhcyBvcmdhbml6ZWQgYXMgdGhlIEJhZCBSYWJiaXRzLCBQYXBwYSBGbHkgUmVsb2FkZWQgZG9lcyBwcm92aWRlIGEgc3Vic3RhbnRpYWwgYW1vdW50IG9mIGZvcmNlLCBhbmQgYmxhY2stbWFya2V0IGNvbm5lY3Rpb24sIHRvIGF1Z21lbnQgQWx2YWNhJ3Mgb3duIGludGVybmFsIHNlY3VyaXR5IGZvcmNlLiBUaGlzIGhhcyBtYWRlIHRoZW0gYW4gZW5lbXkgb2YgdGhlIEJhZCBSYWJiaXRzLCB3aG8gZGlkbid0IGxpa2UgdGhlbSBtdWNoIHRvIGJlZ2luIHdpdGguCgpUaGUgU2FuY3R1YXJ5IHN5c3RlbXMgYWxzbyBoYXZlIGEgZGlzbGlrZSBmb3IgQWx2YWNhLCBhbmQgYXJlIGNvbmNlcm5lZCB0aGF0IHRoZWlyIG93biBjaGFydGVycyBtYXkgYmUgcmV2b2tlZCBzaG9ydGx5IGFmdGVyIEFsdmFjYSdzIGFyZS4KClBlcmhhcHMgdGhlIGdyZWF0ZXN0IGVuZW15IG9mIEFsdmFjYSBpcyBJbnRlcnN0YXIgQ29ycG9yYXRpb24sIHdobyBzZWUgdGhlIG5ldyBiaW9zcGhlcmVzIHRoYXQgQWx2YWNhIGlzIGF0dGVtcHRpbmcgdG8gY3JlYXRlIGFzIHRoZSBncmVhdGVzdCB0aHJlYXQgdG8gdGhlaXIgbW9ub3BvbGllcyBvbiBzeW50aGV0aWMgZmx1aWQgcHJvZHVjdGlvbiwgcGFydGljdWxhcmx5IG9mIHRoZSBuYW5pdGVzIHdpdGhpbiB0aGUgZmx1aWQuIEFzIHRoZSB0ZXJyYWZvcm1pbmcgcHJvY2VzcyBoYXMgcHJvZ3Jlc3NlZCwgcmljaCBkZXBvc2l0cyBvZiBueWNhY2l0ZSBoYXZlIGJlZW4gZGlzY292ZXJlZCBvbiBzb21lIG9mIHRob3NlIHdvcmxkcyAtIHZlcnkgaW1wb3J0YW50IGluIHRoZSBjb25zdHJ1Y3Rpb24gb2YgbmV3IG5hbml0ZXMsIGFuZCBvdXRzaWRlIG9mIHRoZSBjb250cm9sIG9mIEludGVyc3Rhci4gQWx2YWNhIGhhcyByZXBlYXRlZGx5IHJlZnVzZWQgdG8gc2lnbiBhIHRyZWF0eSB3aXRoIEludGVyc3RhciB0aGF0IHdvdWxkIGdyYW50IGl0IGNvbnRyb2wgb2YgdGhvc2UgZGVwb3NpdHMsIGxlYWRpbmcgSW50ZXJzdGFyIHRvIGZlYXIgdGhlIHZhbHVhYmxlIHN1YnN0YW5jZSBtYXkgYmUgZXhwb3J0ZWQgYnkgQWx2YWNhIHRvIG5vbi1JbnRlcnN0YXIgbmFuaXRlIG1hbnVmYWN0dXJlcnMuIFRoaXMgaGFzIHJlc3VsdGVkIGluIGFjdGl2ZSBjb25mbGljdCBiZXR3ZWVuIHRoZSB0d28gY29ycG9yYXRpb25zLgo=';
  };
}