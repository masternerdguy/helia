-- fix collima item type
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d5eeeac7-d7ce-473a-bd7b-5662d65494e3', 'ship', 'Chollima', '{"volume": 1179834, "shiptemplateid": "e2a67722-10d1-42fd-9a7c-057677ef6e79", "industrialmarket": {"maxprice": 1382784819, "minprice": 326086956, "silosize": 5}}');

UPDATE shiptemplates SET itemtypeid = 'd5eeeac7-d7ce-473a-bd7b-5662d65494e3' WHERE id = 'e2a67722-10d1-42fd-9a7c-057677ef6e79';

-- Robin
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('33d34571-523e-4ae4-9da9-442f510eaed8', 'ship', 'Robin', '{"volume": 20000, "shiptemplateid": "d68817c6-2448-4126-8ae5-766669f34cf5", "industrialmarket": {"maxprice": 1402625, "minprice": 375229, "silosize": 500}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('1b54025c-a63a-4c6c-b790-ef6493aca3a6', 'schematic', 'Robin Schematic', '{"industrialmarket": {"maxprice": 301393, "minprice": 101992, "silosize": 100, "process_id": "9f224cea-15a4-4a05-b3e1-9512189769e0"}}');

INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('d68817c6-2448-4126-8ae5-766669f34cf5', '2022-06-19 15:45:30.860312-04', 'Robin', 'Robin', 32, 3.8, 1150, 4.2, 836, 24, 676, 540, 1060, 2948, 44, 138, 36, 'bed0330f-eba3-47ed-8e55-84c753c6c376', '{"a_slots": [{"hp_pos": [4.5, 35], "volume": 40, "mod_family": "gun"}, {"hp_pos": [4.5, -35], "volume": 40, "mod_family": "gun"}, {"hp_pos": [4.5, -35], "volume": 40, "mod_family": "gun"}, {"hp_pos": [3, 0], "volume": 40, "mod_family": "utility"}, {"hp_pos": [3, 0], "volume": 40, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 35, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 26, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 26, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 26, "mod_family": "any"}]}', 240, '33d34571-523e-4ae4-9da9-442f510eaed8', true, 'basic-wreck', 'basic_explosion');

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

