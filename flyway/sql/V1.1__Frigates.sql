-- fix collima item type
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d5eeeac7-d7ce-473a-bd7b-5662d65494e3', 'ship', 'Chollima', '{"volume": 1179834, "shiptemplateid": "e2a67722-10d1-42fd-9a7c-057677ef6e79", "industrialmarket": {"maxprice": 1382784819, "minprice": 326086956, "silosize": 5}}');

UPDATE shiptemplates SET itemtypeid = 'd5eeeac7-d7ce-473a-bd7b-5662d65494e3' WHERE id = 'e2a67722-10d1-42fd-9a7c-057677ef6e79';

-- Robin
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('33d34571-523e-4ae4-9da9-442f510eaed8', 'ship', 'Robin', '{"volume": 20000, "shiptemplateid": "d68817c6-2448-4126-8ae5-766669f34cf5", "industrialmarket": {"maxprice": 1402625, "minprice": 375229, "silosize": 500}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('1b54025c-a63a-4c6c-b790-ef6493aca3a6', 'schematic', 'Robin Schematic', '{"industrialmarket": {"maxprice": 301393, "minprice": 101992, "silosize": 100, "process_id": "9f224cea-15a4-4a05-b3e1-9512189769e0"}}');

INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('d68817c6-2448-4126-8ae5-766669f34cf5', '2022-06-19 15:45:30.860312-04', 'Robin', 'Robin', 32, 3.8, 1150, 4.2, 836, 24, 676, 540, 1060, 2948, 44, 138, 36, 'bed0330f-eba3-47ed-8e55-84c753c6c376', '{"a_slots": [{"hp_pos": [15, 0], "volume": 40, "mod_family": "gun"}, {"hp_pos": [10, 0], "volume": 40, "mod_family": "gun"}, {"hp_pos": [5, 0], "volume": 40, "mod_family": "gun"}, {"hp_pos": [-15, 45], "volume": 40, "mod_family": "utility"}, {"hp_pos": [-15, -45], "volume": 40, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 26, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 26, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 26, "mod_family": "any"}]}', 240, '33d34571-523e-4ae4-9da9-442f510eaed8', true, 'basic-wreck', 'basic_explosion');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('9f224cea-15a4-4a05-b3e1-9512189769e0', 'Make Robin', '{}', 17593);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('3501cc67-3103-4703-bde4-fd4ed05686e0', 'Robin Schematic Faucet [wm]', '{}', 44458);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('f2f6d554-4783-4ee4-a9f3-e83f3c37f514', 'Robin Schematic Sink [wm]', '{}', 134780);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f8f7881d-b736-482b-809f-604518301ea3', '1b54025c-a63a-4c6c-b790-ef6493aca3a6', 1, '{}', 'f2f6d554-4783-4ee4-a9f3-e83f3c37f514');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b03fd047-05ea-46da-a036-6ee6aa5d8225', '61f52ba3-654b-45cf-88e3-33399d12350d', 69, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5360e0d9-2869-488c-a736-6b50a97ecdaa', '24800206-2c58-45b0-8238-81974d0ebb3b', 1014, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e59ef1ae-1084-405d-9fca-a535c3378416', '66b7a322-8cfc-4467-9410-492e6b58f159', 50, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('92f2664a-1fd1-4215-bb65-c73846269559', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 46, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1c4c7d2b-bef9-44e1-9dc7-bb4ba2073980', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 56, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9ed9e361-b3ba-40ff-916c-5ffd28214611', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 219, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('25208f81-67ff-42c4-ac19-ed66649c52ab', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 52, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bba6e90c-2fe5-4cb6-be90-6f9799740ebc', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 48, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ab1c7188-308a-4176-af45-7f95dae07171', '56617d30-6c30-425c-84bf-2484ae8c1156', 48, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('b29779c0-1792-4ae2-8b18-c732d2a74283', '1b54025c-a63a-4c6c-b790-ef6493aca3a6', 3, '{}', '3501cc67-3103-4703-bde4-fd4ed05686e0');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('94e7999f-d649-40ea-9b50-e59855547d79', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 1014, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('5abe7416-6614-4436-b744-e9d57f64429f', '33d34571-523e-4ae4-9da9-442f510eaed8', 1, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');

-- Alligator
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('e5d06f1f-49ce-431d-9c6e-7a9a83a7d3f4', 'ship', 'Alligator', '{"volume": 21233, "shiptemplateid": "9982af33-2a17-4ec9-87b8-85eb6cd857d2", "industrialmarket": {"maxprice": 987500, "minprice": 295233, "silosize": 500}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('8a24cb3e-6739-411e-bf37-904d8a1abefe', 'schematic', 'Alligator Schematic', '{"industrialmarket": {"maxprice": 303825, "minprice": 142079, "silosize": 100, "process_id": "9f18c0cc-5561-469c-af3c-6b3bd4092ead"}}');

INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('9982af33-2a17-4ec9-87b8-85eb6cd857d2', '2022-06-20 08:52:53.99576-04', 'Alligator', 'Alligator', 33, 3.5, 2395, 3.7, 305, 7, 1656, 570, 1107, 2055, 44, 503, 28, 'bed0330f-eba3-47ed-8e55-84c753c6c376', '{"a_slots": [{"hp_pos": [-15, 45], "volume": 40, "mod_family": "missile"}, {"hp_pos": [-14, 0], "volume": 40, "mod_family": "missile"}, {"hp_pos": [-15, -45], "volume": 40, "mod_family": "missile"}, {"hp_pos": [-10, 0], "volume": 45, "mod_family": "utility"}, {"hp_pos": [10, 0], "volume": 45, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 30, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 30, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}]}', 403, 'e5d06f1f-49ce-431d-9c6e-7a9a83a7d3f4', true, 'basic-wreck', 'basic_explosion');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('9f18c0cc-5561-469c-af3c-6b3bd4092ead', 'Make Alligator', '{}', 14499);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('413260af-dd59-431e-ba78-de77c48325ba', 'Alligator Schematic Faucet [wm]', '{}', 60320);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('15da5c6f-d1ab-4225-a3a3-eabda61aa6c6', 'Alligator Schematic Sink [wm]', '{}', 62012);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('53b3440f-78fd-454b-8a01-f9397bce6d70', '24800206-2c58-45b0-8238-81974d0ebb3b', 1365, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('7c739399-c2c8-4a4e-8b3a-9ce1ec535345', '8a24cb3e-6739-411e-bf37-904d8a1abefe', 7, '{}', '15da5c6f-d1ab-4225-a3a3-eabda61aa6c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('74e57d71-5c4d-4852-ba5e-03bba0ad160a', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 246, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e4203468-625a-4e34-a2c7-8f0fb27ecabc', '56617d30-6c30-425c-84bf-2484ae8c1156', 67, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0a1bcee4-49cf-4ec3-8cbd-be7ed32df81c', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 58, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5a83036c-ac9b-4429-9e88-32f14154bf0a', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 52, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('88faebd5-2cc3-490f-a07b-59ee2350a753', '66b7a322-8cfc-4467-9410-492e6b58f159', 63, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ad5ed4b5-7bd0-4dac-9d83-1d143d9125af', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 59, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b79ff562-44af-40eb-b562-ceb64de8cb8a', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 63, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1bc9e93b-20e5-438f-af2d-e892a714e422', '11688112-f3d4-4d30-864a-684a8b96ea23', 45, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('53df41db-a3b2-46c3-91a9-7fa8c2e4890a', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 1365, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a701375b-d775-46d5-86ad-a56eefbcc9ce', '8a24cb3e-6739-411e-bf37-904d8a1abefe', 4, '{}', '413260af-dd59-431e-ba78-de77c48325ba');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('62183bcd-3b60-4ba4-89ea-ce73c8d608fa', 'e5d06f1f-49ce-431d-9c6e-7a9a83a7d3f4', 1, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');

-- Elephant
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('86f89e29-a57e-4045-84d3-9ab9b978d234', 'schematic', 'Elephant Schematic', '{"industrialmarket": {"maxprice": 324016, "minprice": 78511, "silosize": 100, "process_id": "c3d58d27-b000-4127-a6cf-d3a6b8de55c6"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('a475a323-28ea-4dd9-a67f-b1a6427515aa', 'ship', 'Elephant', '{"volume": 18672, "shiptemplateid": "046e606f-89c1-4f23-9137-d3fe2d55503a", "industrialmarket": {"maxprice": 1357500, "minprice": 446250, "silosize": 500}}');

INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('046e606f-89c1-4f23-9137-d3fe2d55503a', '2022-06-20 12:43:05.064446-04', 'Elephant', 'Elephant', 31, 4.1, 995, 4.6, 1725, 38, 540, 396, 1275, 2549, 38, 565, 34.2, 'bed0330f-eba3-47ed-8e55-84c753c6c376', '{"a_slots": [{"hp_pos": [4, 30], "volume": 40, "mod_family": "missile"}, {"hp_pos": [4, -30], "volume": 40, "mod_family": "missile"}, {"hp_pos": [4, -30], "volume": 40, "mod_family": "missile"}, {"hp_pos": [2, 0], "volume": 35, "mod_family": "utility"}, {"hp_pos": [2, 0], "volume": 35, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 50, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 50, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 40, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 40, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 28, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 28, "mod_family": "any"}]}', 166, 'a475a323-28ea-4dd9-a67f-b1a6427515aa', true, 'basic-wreck', 'basic_explosion');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('c3d58d27-b000-4127-a6cf-d3a6b8de55c6', 'Make Elephant', '{}', 19357);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('780377c4-cfcd-4df1-bd86-ec36135a460c', 'Elephant Schematic Faucet [wm]', '{}', 13172);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('53ff6f5f-bd8a-49eb-bdb1-c84e3f9dab96', 'Elephant Schematic Sink [wm]', '{}', 33036);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6c8f368e-ab31-4965-b5d2-f9d4c176f257', '86f89e29-a57e-4045-84d3-9ab9b978d234', 9, '{}', '53ff6f5f-bd8a-49eb-bdb1-c84e3f9dab96');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('672eabc6-6233-410b-a588-f7eb82a1529a', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 150, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f2ef025d-83a5-483a-9d3b-a3d111094a1a', 'd1866be4-5c3e-4b95-b6d9-020832338014', 7, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('41459ffe-ed55-4cbe-a40c-87a7f585eb52', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 29, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d95f2773-e004-44df-a803-f060b84a0084', '66b7a322-8cfc-4467-9410-492e6b58f159', 35, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9a9d8787-fb4f-4368-a442-e79b1b1a59c9', '2ce48bef-f06b-4550-b20c-0e64864db051', 36, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('934c2ae2-9c0c-40bd-bc81-7d77db49a435', '61f52ba3-654b-45cf-88e3-33399d12350d', 35, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bf0cc212-15cd-46a1-be74-1af7b92df481', '24800206-2c58-45b0-8238-81974d0ebb3b', 1005, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cc69c8d8-c803-408e-94e4-d376e66fe210', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 38, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c5e5bfb3-4596-4ae1-9ba5-a1fcd218fa89', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 36, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('f335939b-25ff-4165-a9e1-21285bd20398', '86f89e29-a57e-4045-84d3-9ab9b978d234', 1, '{}', '780377c4-cfcd-4df1-bd86-ec36135a460c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a2302510-95ba-4f5f-a435-b1178f5dd29a', 'a475a323-28ea-4dd9-a67f-b1a6427515aa', 1, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('5219576f-848d-4a73-aeb3-1f47b255b294', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 1005, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');

-- Spectre
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('4524d577-6ebb-46c9-beac-6ba7814c69d3', 'schematic', 'Spectre Schematic', '{"industrialmarket": {"maxprice": 327011, "minprice": 129453, "silosize": 100, "process_id": "65f59a8b-deb3-4418-9f97-ad8624a62f37"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('e6951873-f479-4b64-8a21-7940cf68cf29', 'ship', 'Spectre', '{"volume": 19896, "shiptemplateid": "62100eb8-a7fb-4241-a157-c6d783cb15d9", "industrialmarket": {"maxprice": 1354846, "minprice": 446250, "silosize": 500}}');

INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('62100eb8-a7fb-4241-a157-c6d783cb15d9', '2022-06-20 13:24:23.597377-04', 'Spectre', 'Spectre', 31.5, 3.6, 1492, 3.9, 785, 19, 760, 537, 246, 3663, 45, 606, 37, 'bed0330f-eba3-47ed-8e55-84c753c6c376', '{"a_slots": [{"hp_pos": [6.5, 50], "volume": 40, "mod_family": "gun"}, {"hp_pos": [6.5, -50], "volume": 40, "mod_family": "gun"}, {"hp_pos": [6.5, -50], "volume": 40, "mod_family": "gun"}, {"hp_pos": [3, 0], "volume": 40, "mod_family": "utility"}, {"hp_pos": [3, 0], "volume": 40, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 30, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 30, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 30, "mod_family": "any"}]}', 230, 'e6951873-f479-4b64-8a21-7940cf68cf29', true, 'basic-wreck', 'basic_explosion');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('65f59a8b-deb3-4418-9f97-ad8624a62f37', 'Make Spectre', '{}', 21329);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('3ac67246-e0fe-414f-bb3f-3bc4a8d37173', 'Spectre Schematic Faucet [wm]', '{}', 126381);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9aef6a8d-9d02-4cf4-9cc1-3f81f97b410a', 'Spectre Schematic Sink [wm]', '{}', 179072);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c42c1a14-7eb7-401a-8df5-3482dca7e9c4', '4524d577-6ebb-46c9-beac-6ba7814c69d3', 3, '{}', '9aef6a8d-9d02-4cf4-9cc1-3f81f97b410a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('31c93363-3d3f-414c-b8f2-b967f4596fb7', '24800206-2c58-45b0-8238-81974d0ebb3b', 1808, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3e992d89-8fdb-44e3-90ca-4a03d504e7a2', '66b7a322-8cfc-4467-9410-492e6b58f159', 47, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('59ff1de4-f515-4f8d-8650-42333aa7198a', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 41, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('dc101355-3a26-41e0-bb52-47d2e8cadacf', '61f52ba3-654b-45cf-88e3-33399d12350d', 44, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b125d873-7a42-4064-9198-aef72b0bb280', '11688112-f3d4-4d30-864a-684a8b96ea23', 43, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bb6a1254-8c2f-4565-a53c-eadb1ef1fca8', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 47, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('80b14ddb-ed64-470f-9a00-84ddb5b7966f', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 197, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('164bc84b-727f-402d-8a15-a1f2a5dd2b39', '4524d577-6ebb-46c9-beac-6ba7814c69d3', 4, '{}', '3ac67246-e0fe-414f-bb3f-3bc4a8d37173');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a364ea67-41ad-4a34-a84a-d71fab7559d4', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 1808, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('3bef765a-c471-487e-8e13-4a1cd07b3c24', 'e6951873-f479-4b64-8a21-7940cf68cf29', 1, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
