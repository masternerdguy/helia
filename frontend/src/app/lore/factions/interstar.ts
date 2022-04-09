import { IFactionLoreData } from './iFactionLoreData';

export class InterstarLoreData implements IFactionLoreData {
  factionName: string;
  factionTicker: string;
  factionDescription: string;

  constructor() {
    this.factionName = 'Interstar Corporation';
    this.factionTicker = '[IC*]';
    this.factionDescription = atob(this.rawDescription()).trim();
  }

  private rawDescription = () => {
    return 'QXMgYSBnZW5lcmFsIHJ1bGUsIG9yZ2FuaXphdGlvbnMgdGVuZCB0byBncm93IGluIHNpemUuIEludGVyc3RhciBDb3Jwb3JhdGlvbiBpcyBwZXJoYXBzIHRoZSBiZXN0IGV4YW1wbGUgb2Ygc3VjaCBncm93dGguCgpJbnRlcnN0YXIsIGFzIHRoZSBuYW1lIG1pZ2h0IHN1Z2dlc3QsIGJlZ2FuIGFzIGEgaHVtYmxlIHNoaXBwaW5nIGNvbXBhbnkgb3V0IG9mIFRlbmV2YW4gc3BhY2UuIER1cmluZyB0aGUgd2FyIHdpdGggdGhlIEZlZGVyYXRpb24sIHRoZXkgcHJpbWFyaWx5IHByb3ZpZGVkIGxvZ2lzdGljcyBmb3IgdGhlIG1pbGl0YXJ5LWluZHVzdHJpYWwgY29tcGxleCBvZiB0aGUgVGVuZXZhbiwgc2VlaW5nIGl0IGFzIHRoZWlyIHBhdHJpb3RpYyBkdXR5LiBUaGUgbW90dG8gb2YgSW50ZXJzdGFyIGF0IHRoZSB0aW1lIHdhcyAiV2UgZ28gd2hlcmUgeW91IHdvbid0IiAtIGFsbHVkaW5nIHRvIHRoZWlyIHdpbGxpbmduZXNzIHRvIHBpY2t1cCBhbmQgZGVsaXZlciBsb2FkcyBpbiB2ZXJ5IGRhbmdlcm91cyBhcmVhcyAoc3VjaCBhcyBhY3RpdmUgd2FyIHpvbmVzKS4gRG9pbmcgc28gbWFkZSB0aGVtIHZlcnkgd2VhbHRoeSwgYW5kIGdhdmUgdGhlbSB0aGUgb3Bwb3J0dW5pdHkgdG8gbGVhcm4gYWJvdXQgYW5kIGV4cGFuZCBpbnRvIHRoZSBkZWZlbnNlIGluZHVzdHJ5LCBhcyB3ZWxsIGFzIHRoZSBtZWRpY2FsIGluZHVzdHJ5LiBIb3dldmVyLCBieSB0aGUgdGltZSB0aGUgS2luZ2RvbSBlbnRlcmVkIHRoZSBjb25mbGljdCB0aGV5IGhhZCBiZWNvbWUgZmFyIGxlc3MgZW50aHVzaWFzdGljLgoKVGhpcyBsb3NzIG9mIGVudGh1c2lhc20gd2FzIHByaW1hcmlseSBkdWUgdG8gc3RhbGxlZCBwcm9ncmVzcyBpbiB0aGUgd2FyIGFnYWluc3QgdGhlIGxhc3QgaG9sZG91dHMgb2YgdGhlIEZlZGVyYXRpb24gd2hvIHJlZnVzZWQgdG8geWllbGQgYW5kIHByZXNlbnRlZCBhIHRhci1waXQgdHJhcCBmb3IgdGhlIFRlbmV2YW4gbWlsaXRhcnkuIFNlZWluZyB0aGUgZGlzcHJvcG9ydGlvbmF0ZSByZXNvdXJjZXMgYmVpbmcgZXhwZW5kZWQgb24gdGhlc2UgbGFzdCBmZXcgc3lzdGVtcywgYW5kIG5vIHJlYWwgZ2FpbiBiZXlvbmQgcHJpZGUgaW4gc2lnaHQgZm9yIGNvbnF1ZXJpbmcgdGhlbSwgSW50ZXJzdGFyIGV4ZWN1dGl2ZXMgYmVnYW4gdG8gYXNrIHRoZSBqdW50YSBsZWFkZXJzIGlmIGl0IHdhcyByZWFsbHkgd29ydGh3aGlsZS4gQWx0aG91Z2ggdGhlIGV4ZWN1dGl2ZXMgd2VyZW4ndCBkaXJlY3RseSBwdW5pc2hlZCwgdGhlIFRlbmV2YW4gbWlsaXRhcnkgYW5ub3VuY2VkIGl0IHdvdWxkIG5vdCBiZSByZW5ld2luZyB0aGVpciBjb250cmFjdHMgd2l0aCB0aGVtIGFuZCB3b3VsZCBiZSBzZWVraW5nIG90aGVyIHZlbmRvcnMuCgpUaGUgbG9zcyBvZiB0aGVzZSBjb250cmFjdHMsIGFuZCB0aGUgcHVibGljIGFubm91bmNlbWVudCwgd2FzIHF1aXRlIGFuIGluc3VsdCAtIHdpdGhpbiB0aGUgaGlnaGx5IGF1dGhvcml0YXJpYW4gVGVuZXZhbiBnb3Zlcm5tZW50IHRoZXJlIHdlcmUgZmV3IHByaXZhdGUgY29ycG9yYXRpb25zLCBlc3BlY2lhbGx5IGxhcmdlIG9uZXMsIHRoYXQgd2VyZSBwZXJtaXR0ZWQgdG8gb3BlcmF0ZS4gVGhvc2UgdGhhdCB3ZXJlIHBlcm1pdHRlZCBnZW5lcmFsbHkgaGFkIHRoZSBnb3Zlcm5tZW50IGFzIHRoZWlyIGJpZ2dlc3QgY2xpZW50LCBzbyB0aGlzIGFubm91bmNlbWVudCB3YXMgZWZmZWN0aXZlbHkgYSBkZWF0aCBzZW50ZW5jZSBmb3IgSW50ZXJzdGFyLiBJbiBhZGRpdGlvbiwgdGhlIHBvb2wgb2YgcG90ZW50aWFsICJvdGhlciB2ZW5kb3JzIiB3YXMgaW5zdWx0aW5nbHkgc21hbGwuCgpVbmxpa2Ugb3RoZXIgY29ycG9yYXRpb25zIHdpdGhpbiBDb2FsaXRpb24gc3BhY2UsIHdoaWNoIGhhZCByaXNlbiBhbmQgZmFsbGVuIHB1cmVseSBvbiB0aGUgZmF2b3VyIG9mIHRoZSBqdW50YSwgSW50ZXJzdGFyJ3MgZXhlY3V0aXZlcyBoYWQgY29uc2lkZXJlZCB0aGUgcG9zc2liaWxpdHkgb2YgbG9zaW5nIHRoZWlyIGdvdmVybm1lbnQgY29udHJhY3RzLiBUaGV5IGhhZCBzZWVuIHNldmVyYWwgb3RoZXJzIHN0YXJ2ZWQgb2YgYnVzaW5lc3MgYnkgdGhlIGdvdmVybm1lbnQsIHJlc3VsdGluZyBpbiB0aGVpciBkaXNzb2x1dGlvbiwgYW5kIGhhZCBhIGNvbnRpbmdlbmN5IHBsYW4uCgpBbHRob3VnaCBwcmltYXJpbHkga25vd24gZm9yIGxvZ2lzdGljcyBhbmQgYXJtcyBtYW51ZmFjdHVyaW5nLCBJbnRlcnN0YXIgaGFkIGV4cGFuZGVkIGludG8gdGhlIG1lZGljYWwgZmllbGQgLSBldmVudHVhbGx5IHByb3ZpZGluZyBmaWVsZCBtZWRpY2FsIGtpdHMsIHNwYXJlIHBhcnRzLCBhbmQgbW9zdCBpbXBvcnRhbnRseSBzeW50aGV0aWMgZmx1aWQuIFRoaXMgZmx1aWQsIG9mdGVuIGNhbGxlZCBzeW50aGV0aWMgYmxvb2QsIHByb3ZpZGVzIGltcG9ydGFudCBsdWJyaWNhdGlvbiBhbmQgY29uZHVjdGl2aXR5IHRocm91Z2hvdXQgdGhlIGJvZHkgYW5kIG11c3QgYmUgcGVyaW9kaWNhbGx5IHJlcGxlbmlzaGVkLiBJbiBhZGRpdGlvbiB0byB0aG9zZSBwcm9wZXJ0aWVzLCBpdCBhbHNvIGNvbnRhaW5zIHNwZWNpYWxpemVkIG5hbml0ZXMgdGhhdCBwcm92aWRlIGFuIGltcG9ydGFudCBzZWxmLXJlcGFpciBmYWNpbGl0eSAtIGFsbG93aW5nIHN5bnRoZXRpY3MgdG8gaGVhbCBtaW5vciBpbmp1cmllcyBhbmQgc2xvdyB0aGUgbmF0dXJhbCB3ZWFyLWFuZC10ZWFyIG9uIHRoZWlyIGJvZGllcy4gV2l0aG91dCB0aGVzZSBuYW5pdGVzLCB0aGUgbGlmZXNwYW4gb2YgYSBzeW50aGV0aWMgYm9keSB3b3VsZCBiZSBqdXN0IGEgZmV3IHllYXJzIQoKTWFudWZhY3R1cmluZyB0aGVzZSBuYW5pdGVzLCB3aGljaCB0aGVtc2VsdmVzIGV2ZW50dWFsbHkgd2VhciBvdXQsIGhhZCBiZWVuIGEgdHJ1bHkgYW5jaWVudCBhcnQga25vd24gYnkgYWxsIHN5bnRoZXRpYyBzb2NpZXRpZXMgKG90aGVyd2lzZSwgdGhlaXIgc29jaWV0aWVzIHdvdWxkbid0IGJlIGhlcmUgdG9kYXkpLiBUaGUgY29uc3RydWN0aW9uIG9mIHRoZSBuYW5pdGVzIHdhcyB3ZWxsIHVuZGVyc3Rvb2QsIGFuZCB0aGUga25vd2xlZGdlIHRvIHByb2R1Y2UgdGhlIG1hY2hpbmVzIHRoYXQgbWFrZSB0aGVtIHdhcyBjb21tb24uIE1vc3Qgb2YgdGhlIHJlc291cmNlcyBuZWVkZWQgdG8gbWFrZSB0aGUgbmFuaXRlcyB3ZXJlIGFsc28gcXVpdGUgY29tbW9uLCB3aXRoIG9uZSBleGNlcHRpb24gLSBueWNhY2l0ZS4KCk55Y2FjaXRlIGlzIGEgcmFyZSBtaW5lcmFsIHRoYXQgY2FuIGJlIHVzZWQgYXMgYSByb29tLXRlbXBlcmF0dXJlIHN1cGVyIGNvbmR1Y3Rvci4gV2hpbGUgb3RoZXIgcm9vbS10ZW1wZXJhdHVyZSBzdXBlciBjb25kdWN0b3JzIGFyZSBrbm93biwgbnljYWNpdGUgaXMgdGhlIG9ubHkgbmF0dXJhbGx5IG9jY3VycmluZyBvbmUuIEl0IGlzIGFsc28gdGhlIG9ubHkgb25lIHRoYXQgd29ya3Mgd2VsbCB3aXRoIHRoZSBjaGVtaXN0cnkgYW5kIGNvbnN0cnVjdGlvbiBvZiB0aGUgbmFuaXRlcyAtIGFsbCBvdGhlciBzdWJzdGFuY2VzIHRoYXQgd2VyZSBzdWJzdGl0dXRlZCBlaXRoZXIgZGlkbid0IHdvcmsgYXQgYWxsIG9yIHJlc3VsdGVkIGluIGZhciBpbmZlcmlvciBuYW5pdGVzLiBUaGVzZSBpbmZlcmlvciBuYW5pdGVzIHdvdWxkIHdlYXIgb3V0IHNpZ25pZmljYW50bHkgZmFzdGVyIHRoYW4gcHJvcGVyLCBueWNhY2l0ZSwgb25lcyAtIGxhc3RpbmcgZGF5cyBpbnN0ZWFkIG9mIG1vbnRocy4gVGhleSB3ZXJlIGFsc28gbWVkaW9jcmUgYXQgdGhlaXIgdGFza3MgLSBhIHN5bnRoZXRpYyBwcm92aWRlZCB3aXRoIG5vbi1ueWNhY2l0ZSBuYW5pdGVzIHdvdWxkIGxpdmUgKGF0IG1vc3QpIGhhbGYgYXMgbG9uZyBhcyBvbmUgd2hvIGhhZCBwcm9wZXIgbmFuaXRlcy4gV2l0aCB0aGUgbW9zdCBpbmZlcmlvciAoZnVuY3Rpb25hbCkgbmFuaXRlcywgdGhpcyB3YXMgY2xvc2VyIHRvIGEgcXVhcnRlciBhcyBsb25nLgoKRHVyaW5nIHRoZSBkYXJrIGFnZXMsIG5vbi1ueWNhY2l0ZSBuYW5pdGVzIHdlcmUgY29tbW9uIGFzIGEgc3RvcC1nYXAgd2hlbiBueWNhY2l0ZSB3YXNuJ3QgYXZhaWxhYmxlLiBIYXZpbmcgcHJvcGVyIHZzIGltcHJvcGVyIG5hbml0ZXMgd2FzIGFsc28gYSBjb21tb24gaW5kaWNhdG9yIG9mIHdlYWx0aCwgd2l0aCB0aGUgbWFqb3JpdHkgb2Z0ZW4gaGF2aW5nIHRvIGRlYWwgd2l0aCBub24tbnljYWNpdGUgbmFuaXRlcyB3aGlsZSB0aGUgcHJpdmlsZWdlZCBmZXcgZW5qb3llZCBueWNhY2l0ZSBvbmVzLiBIb3dldmVyLCBpbiBtb3JlIG1vZGVybiB0aW1lcyBhY2Nlc3MgdG8gbnljYWNpdGUgbmFuaXRlcyBoYWQgYmVjb21lIG1vcmUgY29tbW9uIHBsYWNlIGFuZCBhdHRhaW5hYmxlIGR1ZSB0byBpbXByb3ZlZCBtaW5pbmcgYW5kIHByb2Nlc3NpbmcgdGVjaG5pcXVlcy4gQWx0aG91Z2ggcmFyZSwgaXQgYWN0dWFsbHkgdGFrZXMgdmVyeSBsaXR0bGUgbnljYWNpdGUgdG8gbWFrZSBtYW55IG5hbml0ZXMgLSB3aXRoIHdhc3RlZCBueWNhY2l0ZSBsb3N0IGluIG1hbnVmYWN0dXJpbmcgYmVpbmcgdGhlIGJpZ2dlc3QgaGlzdG9yaWNhbCBpbmVmZmljaWVuY3kuIEluIGFkZGl0aW9uLCB0ZWNobmlxdWVzIHdlcmUgZGV2ZWxvcGVkIG92ZXIgdGltZSB0byByZWN5Y2xlIHNvbWUgbnljYWNpdGUgZnJvbSB3b3JuLW91dCBuYW5pdGVzLgoKR2l2ZW4gdGhlaXIgZXhwYW5kaW5nIHJvbGUgaW4gdGhlIHdhciwgSW50ZXJzdGFyIGhhZCBhY3F1aXJlZCBjb250cm9sIG9mIG1hbnkgZmFjaWxpdGllcyBpbnZvbHZlZCBpbiBuYW5pdGUgbWFudWZhY3R1cmUgYW5kIHJlY2xhbWF0aW9uIC0gYXMgd2VsbCBhcyBtYW55IG55Y2FjaXRlIGRlcG9zaXRzLiBHaXZlbiB0aGVpciBsb3lhbHR5LCB1cCB1bnRpbCB0aGVpciBleGVjdXRpdmVzIG9wZW5seSBxdWVzdGlvbmVkIHRoZSBqdW50YSBsZWFkZXJzLCB0aGlzIHdhcyBub3Qgc2VlbiBhcyBhIHNpZ25pZmljYW50IHJpc2suIEhvd2V2ZXIsIG9uY2UgdGhlIG5lZWQgd2FzIGZlbHQgdG8gaGVkZ2UgYWdhaW5zdCB0aGVpciBvd24gZGVzdHJ1Y3Rpb24sIEludGVyc3RhciBiZWdhbiB0byBob2FyZCB0aGUgbnljYWNpdGUgKGFuZCBvdGhlciBtZWFucyBvZiBwcm9kdWN0aW9uIGluIHRoZSBzdXBwbHkgY2hhaW4pLiBCeSB0aGUgdGltZSB0aGUgVGVuZXZhbiBnb3Zlcm5tZW50IGN1dCB0aGVtIG9mZiwgdGhleSBoYWQgYWxyZWFkeSBtb3ZlZCBuZWFybHkgYWxsIG9mIHRoZSBueWNhY2l0ZSwgYW5kIG1hbnkgb3RoZXIgYXNwZWN0cyBvZiB0aGUgc3VwcGx5IGNoYWluLCBvdXQgb2YgQ29hbGl0aW9uIHNwYWNlIQoKTm90IG9ubHkgZGlkIEludGVyc3RhciBub3cgY29udHJvbCBtb3N0IG5hbml0ZSBwcm9kdWN0aW9uIHdpdGhpbiB0aGUgVGVuZXZhbiBDb2FsaXRpb24gaW4gZ2VuZXJhbCwgdGhleSBhbHNvIGNvbnRyb2xsZWQgdGhlIHJhcmUgaW5ncmVkaWVudCB2aXRhbCBmb3IgcHJvZHVjaW5nIHRoZSBoaWdoZXN0IHF1YWxpdHkgbmFuaXRlcyBlbmpveWVkIGJ5IHRoZSBDb2FsaXRpb24gZWxpdGVzLiBUaGlzIG92ZXJzaWdodCBjYXVzZWQgdGhlIHRlcm1pbmF0aW9uIG9mIHRoZWlyIGdvdmVybm1lbnQgY29udHJhY3RzIHRvIGNvbXBsZXRlbHkgYmFja2ZpcmUsIHdpdGggdGhlIENvYWxpdGlvbiBnb3Zlcm5tZW50IGJlZ3J1ZGdpbmdseSBhZ3JlZWluZyB0byB0ZXJtcyBkaWN0YXRlZCBieSBJbnRlcnN0YXIgZm9yIGNvbnRpbnVlZCBhY2Nlc3MgdG8gbmFuaXRlcy4KCkVtYm9sZGVuZWQgYnkgdGhlIHN1Y2Nlc3Mgb2YgdGhlaXIgY29udGluZ2VuY3kgcGxhbiwgYW5kIHNlZWluZyBvcHBvcnR1bml0eSBpbiB0aGUgdW5lYXN5IHN0YWxlbWF0ZSBiZXR3ZWVuIHRoZSB0d28gc2lkZXMsIEludGVyc3RhciBkZWNpZGVkIHRvIGV4cGFuZCBpbnRlcm5hdGlvbmFsbHkuIFVzaW5nIHRoZWlyIGJyb2FkIHNraWxsIHNldCwgdGhleSBleHBhbmRlZCBpbnRvIHRoZSBvdGhlciB0aHJlZSBncmVhdCBuYXRpb25zIGFzIGxvZ2lzdGljcywgbWVkaWNhbCwgYW5kIG1hbnVmYWN0dXJpbmcgY29tcGFuaWVzLiBIb3dldmVyLCB0aGV5IGFsc28gc2VjcmV0bHkgdXNlZCBzZWVtaW5nbHkgdW5yZWxhdGVkIHNoZWxsIGNvcnBvcmF0aW9ucyB0byBzbG93bHkgYnV5IGFzcGVjdHMgb2YgdGhlIG5hbml0ZSBzdXBwbHkgY2hhaW4gd2l0aGluIHRoZW0gLSB3aXRoIGEgc3BlY2lhbCBmb2N1cyBvbiBueWNhY2l0ZSBhbmQgdGhlIGluZHVzdHJ5IHN1cnJvdW5kaW5nIGl0LgoKT2Z0ZW4gdGltZXMgdGhlc2UgYWNxdWlzaXRpb25zIHdlcmUgbWFkZSB3aXRoIHRoZSBwcm9taXNlIG9mIGltcHJvdmluZyBuYW5pdGUgb3V0cHV0LCBxdWFsaXR5LCBvciBsb3dlcmluZyBwcmljZXMuIE9uY2UgaW4gY29udHJvbCwgaG93ZXZlciwgdGhleSBnZW5lcmFsbHkgcmFpc2VkIHByaWNlcyBzaWduaWZpY2FudGx5LCBzb21ldGltZXMgfjEweCwgYnkgY3V0dGluZyBvdXRwdXQuIEJ1bGx5aW5nIHRhY3RpY3Mgd2VyZSBhbHNvIGVtcGxveWVkIGFnYWluc3QgbG9jYWwgcG9wdWxhdGlvbnMgYW5kIGxlYWRlcnMgdG8gcHJldmVudCBuZXdzIG9mIHRoaXMgZGVjZWl0IHNwcmVhZGluZyBhbmQgaGFsdCBhbnkgaW52ZXN0aWdhdGlvbnMuIE55Y2FjaXRlIHdhcyB1c3VhbGx5IG1vdmVkIGludG8gSW50ZXJzdGFyIHNwYWNlIChub3cgb2ZmaWNpYWxseSB1bmRlciB0aGVpciBjb250cm9sIGFmdGVyIGEgY29uY2Vzc2lvbiBmcm9tIHRoZSBDb2FsaXRpb24pLCBhbmQgc3RvcmVkIGluIGhlYXZpbHkgZm9ydGlmaWVkIGluc3RhbGxhdGlvbnMgZm9yIGdyYWR1YWwgcmVsZWFzZS4KCkJ1bGx5aW5nIGFuZCBsb2JieWluZyB0YWN0aWNzIHdlcmUgYWxzbyB1c2VkLCB3aGVyZSBwb3NzaWJsZSwgdG8gcGFzcyBsYXdzIGFuZCByZWd1bGF0aW9ucyB0aGF0IG1hZGUgaXQgaW1wcmFjdGljYWwgZm9yIG5vbi1JbnRlcnN0YXIgZmFjaWxpdGllcyB0byBwcm9kdWNlIG5hbml0ZXMuIEdlbmVyYWxseSwgdGhpcyB3YXMgYWR2YW5jZWQgYnkgY3JlYXRpbmcgaG9heCBoZWFsdGggc2NhcmVzIG9mIGZha2UsIG9yIGNvcm5lci1jdXQsIG5hbml0ZXMgYmVpbmcgc29sZCBhcyBwcm9wZXIgb25lcy4gU29tZXRpbWVzLCBJbnRlcnN0YXIgd291bGQgZXZlbiBzYWJvdGFnZSBub24tSW50ZXJzdGFyIG5hbml0ZSBwcm9kdWN0aW9uIGZhY2lsaXRpZXMgc28gdGhhdCB0aGVpciBuYW5pdGVzIHdvdWxkIGNhdXNlIGhhcm0gdG8gdGhvc2Ugd2hvIHVzZWQgdGhlbS4KClRoZXNlIHRhY3RpY3MgaGF2ZSBvbmx5IGJlZW4gd2lkZWx5IGtub3duIGZvciBsZXNzIHRoYW4gYSBkZWNhZGUsIGFuZCBieSBub3cgSW50ZXJzdGFyIGNvbnRyb2xzIGtleSBhc3BlY3RzIG9mIHRoZSBuYW5pdGUgaW5kdXN0cnkgaW4gYWxsIGZvdXIgZ3JlYXQgbmF0aW9ucy4gSG9sZGluZyB0aGUgbnljYWNpdGUgaG9zdGFnZSBoYXMgZ2l2ZW4gdGhlbSBncmVhdCBpbmZsdWVuY2Ugb3ZlciB0aGUgZWxpdGVzIGluIGFsbCBvZiB0aGVtLCBhbmQgdGh1cyBJbnRlcnN0YXIgd2llbGRzIGdyZWF0IHBvbGl0aWNhbCBwb3dlci4KCkJlc2lkZXMgdGhlIG5hbml0ZXMsIHRoZWlyIGJyb2FkIGV4cGFuc2lvbiBpbnRvIG90aGVyIGluZHVzdHJpZXMgbWVhbnMgdGhleSBvd24gdGhvdXNhbmRzIG9mIHNlZW1pbmdseSB1bnJlbGF0ZWQgYnJhbmRzLCBtYW55IG9mIHdoaWNoIGFyZSBob3VzZWhvbGQgbmFtZXMuIEl0IGlzIHZpcnR1YWxseSBpbXBvc3NpYmxlIHRvIG1ha2UgYSBicmFuZC1uYW1lIHB1cmNoYXNlIHdpdGhvdXQgYXQgbGVhc3Qgc29tZSBvZiB0aGUgbW9uZXkgZ29pbmcgdG8gSW50ZXJzdGFyLiBBcyBhIGNvdW50ZXJtZWFzdXJlIGFnYWluc3QgdGhlaXIgaW5jcmVhc2luZ2x5IG5lZ2F0aXZlIHB1YmxpYyBpbWFnZSwgSW50ZXJzdGFyJ3MgcHVibGljIHJlbGF0aW9ucyBhbmQgbWFya2V0aW5nIGRlcGFydG1lbnRzIHdvcmsgdGlyZWxlc3NseSB0byBwYWludCB0aGVtIGluIGFuIGV4dHJlbWVseSBwb3NpdGl2ZSBsaWdodCAtIHdoaWxlIHRoZWlyIGxlZ2FsIGRlcGFydG1lbnQgcGVyc2VjdXRlcyB0aG9zZSB3aG8gc3ByZWFkIGluZm9ybWF0aW9uIHRvIHRoZSBjb250cmFyeS4KClRoaXMgY2FtcGFpZ24gaGFzIGJlZW4gc28gZWZmZWN0aXZlIHRoYXQgbWFueSBhdCBJbnRlcnN0YXIgdHJ1bHkgYmVsaWV2ZSB0aGF0IHRoZXkgYXJlIGRvaW5nIGFsbCBzeW50aGV0aWNzIGEgZ3JlYXQgc2VydmljZSBieSB3b3JraW5nIGF0IEludGVyc3RhciwgZXZlbiBrbm93aW5nIHRoZSB0YWN0aWNzIHRoYXQgaGF2ZSBiZWVuIHVzZWQuIEhvd2V2ZXIsIGl0IGlzIHRoZSBydXRobGVzcyB0aGF0IHRlbmQgdG8gY2xpbWIgdGhlIGNvcnBvcmF0ZSBsYWRkZXIuCgpGb3Igb2J2aW91cyByZWFzb25zLCBJbnRlcnN0YXIgaXMgbm90IG9uIGdvb2QgdGVybXMgd2l0aCB0aGUgVGVuZXZhbiBDb2FsaXRpb24uIEhvd2V2ZXIsIHRoaXMgaXMgbWVhc3VyZWQgYWdhaW5zdCB0aGVpciBjb250cm9sIG9mIHRoZSBlc3NlbnRpYWwgbmFuaXRlIGluZHVzdHJ5LCBhbmQgdGhlIFRlbmV2YW4gZ292ZXJubWVudCBpcyB1bndpbGxpbmcgdG8gcmlzayBiZWluZyBjb21wbGV0ZWx5IGN1dCBvZmYgZnJvbSBueWNhY2l0ZS4gQSBzaW1pbGFyIHNpdHVhdGlvbiBleGlzdHMgd2l0aGluIHRoZSBWaWVycmEgRmVkZXJhdGlvbiwgd2hvIGtub3cgb2YgSW50ZXJzdGFyJ3MgY29udHJpYnV0aW9ucyBhZ2FpbnN0IHRoZW0gaW4gdGhlIHdhci4KClRoZSBLaW5nZG9tIG9mIEFudGFyaWEsIHdobydzIHJveWFsdHkgaGF2ZSBlbmpveWVkIGV4dHJlbWVseSBsb25nIGxpdmVzIGR1ZSBpbiBncmVhdCBwYXJ0IHRvIHRoZSBleGNlcHRpb25hbGx5IGhpZ2ggcXVhbGl0eSBuYW5pdGVzICh3aGljaCByZXF1aXJlIG1vcmUgbnljYWNpdGUgdGhhbiBub3JtYWwpIHRoZXkgZW5qb3ksIGFyZSB1bndpbGxpbmcgdG8gc3BlYWsgb3V0IGFnYWluc3QgSW50ZXJzdGFyIGFuZCBoYXZlIG5vbWluYWxseSBkZWNsYXJlZCB0aGVtIGEgZnJpZW5kbHkgY29ycG9yYXRpb24uIFRoaXMgaGFzIGJlZW4gaGVhdmlseSBjcml0aWNpemVkIGJvdGggd2l0aGluIHRoZSBLaW5nZG9tIGFuZCBhYnJvYWQuIFN1Y2ggY3JpdGljaXNtcyBoYXZlIGJlZW4gZ2VuZXJhbGx5IGlnbm9yZWQgYnkgdGhlIEFudGFycmlhbiByb3lhbHR5LCB3aG8gcHJlZmVyIG5vdCB0byBkaXNjdXNzIGl0IGF0IGFsbC4gVGhlIGZhaWx1cmUgb2YgdGhlIHJveWFsdHkgdG8gc3BlYWsgb3V0IGFuZCB0YWtlIGFjdGlvbiBhZ2FpbnN0IEludGVyc3RhciBoYXMgYmVlbiBoZWF2aWx5IGNyaXRpY2l6ZWQgYnkgdGhlIGNsZXJneSwgd2hvIG9mdGVuIGxhYmVsIHRoZSBuYW5pdGUgc2l0dWF0aW9uIGFzIHRoZSBncmVhdGVzdCBldmlsIG9mIHRoZWlyIHRpbWUgYW5kIGFuIGFmZnJvbnQgdG8gR29kLgoKV2l0aGluIHRoZSBPem91a2EgQWNjb3JkLCB3aG8ncyBlbGl0ZSBhbHNvIGVuam95IGhpZ2ggcXVhbGl0eSBuYW5pdGVzIG1hZGUgd2l0aCBueWNhY2l0ZSwgSW50ZXJzdGFyIGlzIGNvbnNpZGVyZWQgYSBmcmllbmRseSBjb3Jwb3JhdGlvbi4gVGhlIG1vc3QgcG93ZXJmdWwgbG9yZHMgc3RpbGwgaGF2ZSB0aGVpciBvd24gbmFuaXRlIHN1cHBseSBjaGFpbnMsIGtlcHQgc2FmZSBmcm9tIEludGVyc3Rhciwgd2hpbGUgdGhlIGxvd2VyIGxvcmRzIGFyZSB1bndpbGxpbmcgdG8gcmlzayBsb3NpbmcgYWNjZXNzIHRvIHRoZSBuYW5pdGVzIHRoZXkgbm93IG11c3QgYnV5IGZyb20gSW50ZXJzdGFyLiBJdCBzaG91bGQgYmUgbm90ZWQgdGhhdCB3aXRoaW4gdGhlIEFjY29yZCB0aGVyZSBoYXMgYmVlbiBhIG5hbml0ZSBjYXN0IHN5c3RlbSBpbiBwbGFjZSBzaW5jZSBhbnRpcXVpdHksIHdpdGggdGhlIGxvd2VyIHJhbmtzIG9mIHNvY2lldHkgZXhwZWN0aW5nIHRvIGJlIGdpdmVuIGluZmVyaW9yIG5hbml0ZXMgYW5kIGJlbGlldmluZyBpdCBpcyB0aGVpciBwbGFjZSB0byByZWNlaXZlIHRoZW0uIER1ZSB0byB0aGlzIGN1bHR1cmFsIGV4cGVjdGF0aW9uLCB0aGVyZSBpcyBmYXIgbGVzcyBvdXRyYWdlIG92ZXIgSW50ZXJzdGFyJ3MgYWN0aW9ucyB3aXRoaW4gdGhlIEFjY29yZCB0aGFuIGFicm9hZC4KCkludGVyc3RhciBjb25zaWRlcnMgQWx2YWNhLCB3aG8ncyB3b3JrLWluLXByb2dyZXNzIHdvcmxkcyB3aWxsIGV2ZW50dWFsbHkgeWllbGQgZW5vcm1vdXMgYW1vdW50cyBvZiBueWNhY2l0ZSBvdXRzaWRlIG9mIEludGVyc3RhciBjb250cm9sLCB0byBiZSB0aGUgZ3JlYXRlc3QgdGhyZWF0IHRvIHRoZWlyIG1vbm9wb2xpZXMuIEdpdmVuIEFsdmFjYSdzIHVud2lsbGluZ25lc3MgdG8gYWdyZWUgdG8gSW50ZXJzdGFyJ3MgdGVybXMsIHRoZXkgYXJlIGN1cnJlbnRseSB3YWdpbmcgYSB3YXIgYWdhaW5zdCB0aGVtIGZvciBjb250cm9sIG9mIHNldmVyYWwgd29ybGRzLgoKQm90aCBQYXBwYSBGbHkgUmVsb2FkZWQgYW5kIHRoZSBCYWQgUmFiYml0cyBwcmV5IG9uIEludGVyc3RhciB0cmFuc3BvcnQgc2hpcHMsIHNvbWV0aW1lcyBiZWNvbWluZyBib2xkIGVub3VnaCB0byByYWlkIG5hbml0ZSBwcm9kdWN0aW9uIGZhY2lsaXRpZXMsIGFuZCBldmVuIG55Y2FjaXRlIHN0b3JhZ2UgZmFjaWxpdGllcy4gUmFpZGluZyB0aGVzZSBmYWNpbGl0aWVzIGlzIGEgcG9pbnQgb2YgcHJpZGUgZm9yIFBhcHBhIEZseSBSZWxvYWRlZCwgd2hvIHNlZSB0aGVtc2VsdmVzIGFzIGhvbm91cmFibGUgdGhpZXZlcyB0YWtpbmcgZnJvbSB0aGUgcmljaCB0byAiZ2l2ZSIgdG8gdGhlIG5lZWR5Lgo=';
  };
}
