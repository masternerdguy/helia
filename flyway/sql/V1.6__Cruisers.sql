-- battlecruiser ship type for future use
INSERT INTO public.shiptypes (id, name) VALUES ('5b971c58-cec9-47e9-975a-ccc45977d967', 'Battlecruiser');

-- item types
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('04a3b923-3b47-4cad-b93d-57c8831723b9', 'ship', 'Kea', '{"volume": 40000, "shiptemplateid": "211646ef-a445-4e5d-b3e1-708b7a1d4817", "industrialmarket": {"maxprice": 9860453, "minprice": 2705401, "silosize": 250}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('505d4a75-5bf9-4e33-abda-b39d77aef060', 'ship', 'Skink', '{"volume": 42509, "shiptemplateid": "754743cf-d534-4d6f-a46a-6d0e94d62d69", "industrialmarket": {"maxprice": 6932250, "minprice": 2072535, "silosize": 250}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('68f2f340-9d70-486a-8c07-932c6625d1eb', 'ship', 'Kudu', '{"volume": 39211, "shiptemplateid": "baa7d3dd-4261-43d7-94b5-d9e058e03b5d", "industrialmarket": {"maxprice": 9556800, "minprice": 3141600, "silosize": 250}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('5ea74809-8fad-4544-86b7-a3cfc98f924c', 'ship', 'Spook', '{"volume": 41781, "shiptemplateid": "472151e7-7a53-47af-a16f-7eb4f4d7776d", "industrialmarket": {"maxprice": 9619406, "minprice": 3168375, "silosize": 250}}');

-- ship templates
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('754743cf-d534-4d6f-a46a-6d0e94d62d69', '2022-12-05 12:37:14.083022-05', 'Skink', 'Skink', 65, 2.9, 14355, 3.1, 1220, 12, 6790, 2283, 4389, 4137, 85, 1441, 58, 'b6be8bdb-37d4-4899-9092-0c5c1901ed62', '{"a_slots": [{"hp_pos": [-15, 45], "volume": 80, "mod_family": "missile"}, {"hp_pos": [-14, 0], "volume": 80, "mod_family": "missile"}, {"hp_pos": [-15, -45], "volume": 80, "mod_family": "missile"}, {"hp_pos": [-15, -25], "volume": 80, "mod_family": "missile"}, {"hp_pos": [-10, 20], "volume": 90, "mod_family": "utility"}, {"hp_pos": [10, 0], "volume": 90, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 60, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 60, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 60, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}]}', 795, '505d4a75-5bf9-4e33-abda-b39d77aef060', true, 'basic-wreck', 'basic_explosion');
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('211646ef-a445-4e5d-b3e1-708b7a1d4817', '2022-12-05 12:32:51.536634-05', 'Kea', 'Kea', 56, 3.2, 6900, 3.6, 3375, 49, 2713, 2165, 4166, 5896, 83, 1439, 72, 'b6be8bdb-37d4-4899-9092-0c5c1901ed62', '{"a_slots": [{"hp_pos": [15, 0], "volume": 80, "mod_family": "gun"}, {"hp_pos": [10, 0], "volume": 80, "mod_family": "gun"}, {"hp_pos": [5, 0], "volume": 80, "mod_family": "gun"}, {"hp_pos": [20, 0], "volume": 80, "mod_family": "gun"}, {"hp_pos": [-15, 45], "volume": 80, "mod_family": "utility"}, {"hp_pos": [-15, -45], "volume": 80, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 50, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 50, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 50, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 50, "mod_family": "any"}]}', 480, '04a3b923-3b47-4cad-b93d-57c8831723b9', true, 'basic-wreck', 'basic_explosion');
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('472151e7-7a53-47af-a16f-7eb4f4d7776d', '2022-12-05 12:47:42.284319-05', 'Spook', 'Spook', 63, 3, 8957, 3.3, 3139, 41, 3095, 2154, 993, 9187, 92, 1575, 73, 'b6be8bdb-37d4-4899-9092-0c5c1901ed62', '{"a_slots": [{"hp_pos": [10, 0], "volume": 80, "mod_family": "gun"}, {"hp_pos": [5, 0], "volume": 80, "mod_family": "gun"}, {"hp_pos": [0, 0], "volume": 80, "mod_family": "gun"}, {"hp_pos": [15, 0], "volume": 80, "mod_family": "gun"}, {"hp_pos": [-5, 0], "volume": 80, "mod_family": "utility"}, {"hp_pos": [-10, 0], "volume": 80, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 70, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 60, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 60, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 60, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 60, "mod_family": "any"}]}', 465, '5ea74809-8fad-4544-86b7-a3cfc98f924c', true, 'basic-wreck', 'basic_explosion');
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('baa7d3dd-4261-43d7-94b5-d9e058e03b5d', '2022-12-05 12:41:21.347976-05', 'Kudu', 'Kudu', 60, 3.5, 5940, 4, 6917, 82, 2145, 1422, 5105, 5097, 73, 1466, 70, 'b6be8bdb-37d4-4899-9092-0c5c1901ed62', '{"a_slots": [{"hp_pos": [10, 45], "volume": 80, "mod_family": "missile"}, {"hp_pos": [10, -45], "volume": 80, "mod_family": "missile"}, {"hp_pos": [-10, 45], "volume": 80, "mod_family": "missile"}, {"hp_pos": [-10, -45], "volume": 80, "mod_family": "missile"}, {"hp_pos": [-5, 0], "volume": 70, "mod_family": "utility"}, {"hp_pos": [5, 0], "volume": 70, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 100, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 100, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 80, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 80, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 60, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 55, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 55, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 55, "mod_family": "any"}]}', 321, '68f2f340-9d70-486a-8c07-932c6625d1eb', true, 'basic-wreck', 'basic_explosion');

-- fair processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9f224cea-15a4-4a05-b3e1-9512189769e0', 'Make Kea', '{}', 52779);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9f18c0cc-5561-469c-af3c-6b3bd4092ead', 'Make Skink', '{}', 42047);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('c3d58d27-b000-4127-a6cf-d3a6b8de55c6', 'Make Kudu', '{}', 60006);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('65f59a8b-deb3-4418-9f97-ad8624a62f37', 'Make Spook', '{}', 65053);

-- process inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b03fd047-05ea-46da-a036-6ee6aa5d8225', '61f52ba3-654b-45cf-88e3-33399d12350d', 483, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5360e0d9-2869-488c-a736-6b50a97ecdaa', '24800206-2c58-45b0-8238-81974d0ebb3b', 7098, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e59ef1ae-1084-405d-9fca-a535c3378416', '66b7a322-8cfc-4467-9410-492e6b58f159', 350, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('92f2664a-1fd1-4215-bb65-c73846269559', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 322, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1c4c7d2b-bef9-44e1-9dc7-bb4ba2073980', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 392, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9ed9e361-b3ba-40ff-916c-5ffd28214611', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 1533, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('25208f81-67ff-42c4-ac19-ed66649c52ab', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 364, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bba6e90c-2fe5-4cb6-be90-6f9799740ebc', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 336, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ab1c7188-308a-4176-af45-7f95dae07171', '56617d30-6c30-425c-84bf-2484ae8c1156', 336, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('53b3440f-78fd-454b-8a01-f9397bce6d70', '24800206-2c58-45b0-8238-81974d0ebb3b', 9282, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('74e57d71-5c4d-4852-ba5e-03bba0ad160a', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 1672, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e4203468-625a-4e34-a2c7-8f0fb27ecabc', '56617d30-6c30-425c-84bf-2484ae8c1156', 455, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0a1bcee4-49cf-4ec3-8cbd-be7ed32df81c', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 394, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5a83036c-ac9b-4429-9e88-32f14154bf0a', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 353, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('88faebd5-2cc3-490f-a07b-59ee2350a753', '66b7a322-8cfc-4467-9410-492e6b58f159', 428, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ad5ed4b5-7bd0-4dac-9d83-1d143d9125af', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 401, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b79ff562-44af-40eb-b562-ceb64de8cb8a', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 428, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1bc9e93b-20e5-438f-af2d-e892a714e422', '11688112-f3d4-4d30-864a-684a8b96ea23', 306, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('672eabc6-6233-410b-a588-f7eb82a1529a', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 1035, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f2ef025d-83a5-483a-9d3b-a3d111094a1a', 'd1866be4-5c3e-4b95-b6d9-020832338014', 48, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('41459ffe-ed55-4cbe-a40c-87a7f585eb52', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 200, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d95f2773-e004-44df-a803-f060b84a0084', '66b7a322-8cfc-4467-9410-492e6b58f159', 241, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9a9d8787-fb4f-4368-a442-e79b1b1a59c9', '2ce48bef-f06b-4550-b20c-0e64864db051', 248, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('934c2ae2-9c0c-40bd-bc81-7d77db49a435', '61f52ba3-654b-45cf-88e3-33399d12350d', 241, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bf0cc212-15cd-46a1-be74-1af7b92df481', '24800206-2c58-45b0-8238-81974d0ebb3b', 6934, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cc69c8d8-c803-408e-94e4-d376e66fe210', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 262, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c5e5bfb3-4596-4ae1-9ba5-a1fcd218fa89', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 248, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('31c93363-3d3f-414c-b8f2-b967f4596fb7', '24800206-2c58-45b0-8238-81974d0ebb3b', 12655, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3e992d89-8fdb-44e3-90ca-4a03d504e7a2', '66b7a322-8cfc-4467-9410-492e6b58f159', 328, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('59ff1de4-f515-4f8d-8650-42333aa7198a', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 286, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('dc101355-3a26-41e0-bb52-47d2e8cadacf', '61f52ba3-654b-45cf-88e3-33399d12350d', 307, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b125d873-7a42-4064-9198-aef72b0bb280', '11688112-f3d4-4d30-864a-684a8b96ea23', 300, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bb6a1254-8c2f-4565-a53c-eadb1ef1fca8', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 328, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('80b14ddb-ed64-470f-9a00-84ddb5b7966f', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 1378, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');

-- process outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('94e7999f-d649-40ea-9b50-e59855547d79', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 7098, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('5abe7416-6614-4436-b744-e9d57f64429f', '04a3b923-3b47-4cad-b93d-57c8831723b9', 1, '{}', '9f224cea-15a4-4a05-b3e1-9512189769e0');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('53df41db-a3b2-46c3-91a9-7fa8c2e4890a', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 9282, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('62183bcd-3b60-4ba4-89ea-ce73c8d608fa', '505d4a75-5bf9-4e33-abda-b39d77aef060', 1, '{}', '9f18c0cc-5561-469c-af3c-6b3bd4092ead');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a2302510-95ba-4f5f-a435-b1178f5dd29a', '68f2f340-9d70-486a-8c07-932c6625d1eb', 1, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('5219576f-848d-4a73-aeb3-1f47b255b294', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 6934, '{}', 'c3d58d27-b000-4127-a6cf-d3a6b8de55c6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('3bef765a-c471-487e-8e13-4a1cd07b3c24', '5ea74809-8fad-4544-86b7-a3cfc98f924c', 1, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a364ea67-41ad-4a34-a84a-d71fab7559d4', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 12655, '{}', '65f59a8b-deb3-4418-9f97-ad8624a62f37');
