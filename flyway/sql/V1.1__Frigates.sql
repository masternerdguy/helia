-- fix collima item type
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d5eeeac7-d7ce-473a-bd7b-5662d65494e3', 'ship', 'Chollima', '{"volume": 1179834, "shiptemplateid": "e2a67722-10d1-42fd-9a7c-057677ef6e79", "industrialmarket": {"maxprice": 1382784819, "minprice": 326086956, "silosize": 5}}');

UPDATE shiptemplates SET itemtypeid = 'd5eeeac7-d7ce-473a-bd7b-5662d65494e3' WHERE id = 'e2a67722-10d1-42fd-9a7c-057677ef6e79';

-- Robin
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('33d34571-523e-4ae4-9da9-442f510eaed8', 'ship', 'Robin', '{"volume": 20000, "shiptemplateid": "d68817c6-2448-4126-8ae5-766669f34cf5", "industrialmarket": {"maxprice": 1402625, "minprice": 375229, "silosize": 500}}');

INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('d68817c6-2448-4126-8ae5-766669f34cf5', '2022-06-19 15:45:30.860312-04', 'Robin', 'Robin', 32, 3.8, 1150, 4.2, 836, 24, 676, 540, 1060, 2948, 44, 138, 36, 'bed0330f-eba3-47ed-8e55-84c753c6c376', '{"a_slots": [{"hp_pos": [4.5, 35], "volume": 10, "mod_family": "gun"}, {"hp_pos": [4.5, -35], "volume": 10, "mod_family": "gun"}, {"hp_pos": [3, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 6, "mod_family": "any"}]}', 240, '33d34571-523e-4ae4-9da9-442f510eaed8', true, 'basic-wreck', 'basic_explosion');

