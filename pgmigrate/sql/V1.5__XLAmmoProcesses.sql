-- ====== gauss shell
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b08b30d9-c1c4-45a3-af4c-9bcf67dc6b5a', 'Make Heavy Gauss Shell', '{}', 379);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8f093669-c299-4195-9929-57aec24c0078', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 131, '{}', 'b08b30d9-c1c4-45a3-af4c-9bcf67dc6b5a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5562894b-88ef-48b5-9515-7acde529d8e7', '24800206-2c58-45b0-8238-81974d0ebb3b', 62, '{}', 'b08b30d9-c1c4-45a3-af4c-9bcf67dc6b5a');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('da5b8c3a-a540-416f-8b2f-8920c5388fc3', '366756f0-60e4-4dc1-864f-f94b3ce2c6e3', 385, '{}', 'b08b30d9-c1c4-45a3-af4c-9bcf67dc6b5a');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('90f988ce-cc02-4f2f-a26d-27623ecaf888', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 62, '{}', 'b08b30d9-c1c4-45a3-af4c-9bcf67dc6b5a');

-- ====== auto belt
INSERT INTO public.processes (id, name, meta, "time") VALUES ('bb56670a-fc36-445e-a522-88b21092c033', 'Make Auto-23 Belt', '{}', 451);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('91369f97-b7a6-4b4a-8427-516d98d46164', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 43, '{}', 'bb56670a-fc36-445e-a522-88b21092c033');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('01744dd2-c5e1-4cf7-9c07-aa76bda0092c', '24800206-2c58-45b0-8238-81974d0ebb3b', 21, '{}', 'bb56670a-fc36-445e-a522-88b21092c033');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('69774c84-e21d-41d3-84ff-9ef347950c5f', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 21, '{}', 'bb56670a-fc36-445e-a522-88b21092c033');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('d2f2bc9d-f79d-4aff-858d-a85602b1769d', 'b8aa4560-39ea-4af5-a7b8-d360e348a81e', 89, '{}', 'bb56670a-fc36-445e-a522-88b21092c033');

-- ====== hgpm
INSERT INTO public.processes (id, name, meta, "time") VALUES ('0f9d4229-45a0-4923-9af6-d56a3ca5195c', 'Make HGPM-15', '{}', 355);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f171f362-2f4f-488d-ba53-5b60f6e5e22e', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 429, '{}', '0f9d4229-45a0-4923-9af6-d56a3ca5195c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d11ba767-5f4b-4713-b258-f51ab618fb70', 'a8646647-881a-4d24-a22f-f0dce044e6d3', 7, '{}', '0f9d4229-45a0-4923-9af6-d56a3ca5195c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('04ffc193-e51f-426a-8cb1-6d35d280a8e3', '24800206-2c58-45b0-8238-81974d0ebb3b', 167, '{}', '0f9d4229-45a0-4923-9af6-d56a3ca5195c');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('02315008-630a-4aff-87d8-0f8f0f050959', '368e03a0-bfef-43df-b686-2a5f279549d1', 92, '{}', '0f9d4229-45a0-4923-9af6-d56a3ca5195c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('c991d482-4282-47cd-a6c7-ce9397dd4772', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 167, '{}', '0f9d4229-45a0-4923-9af6-d56a3ca5195c');

-- ==== schematics
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('308df2f2-37ea-48c1-aab5-127ce8be4843', 'schematic', 'HGPM-15 Schematic', '{"industrialmarket": {"maxprice": 454656, "minprice": 144284, "silosize": 100, "process_id": "0f9d4229-45a0-4923-9af6-d56a3ca5195c"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('ae54a20c-ac6c-44bb-b245-277ac2aee144', 'schematic', 'Heavy Gauss Shell Schematic', '{"industrialmarket": {"maxprice": 609476, "minprice": 113430, "silosize": 100, "process_id": "b08b30d9-c1c4-45a3-af4c-9bcf67dc6b5a"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('a25f293e-874d-43fa-834b-ed7e4ff11480', 'schematic', 'Auto-23 Belt Schematic', '{"industrialmarket": {"maxprice": 396160, "minprice": 37024, "silosize": 100, "process_id": "bb56670a-fc36-445e-a522-88b21092c033"}}');

-- processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('80ac90c8-5d0b-4dc9-a8c0-8ecff096c87e', 'HGPM-15 Schematic Faucet', '{}', 211053);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('4e95c448-bf09-4471-81c0-9e9f9449d2fe', 'HGPM-15 Schematic Sink', '{}', 138458);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('d692d5d8-aeaa-48b5-af11-bef2f825ade7', 'Heavy Gauss Shell Schematic Faucet', '{}', 14711);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('7391e693-6d37-4ecb-8bec-b54798a21857', 'Heavy Gauss Shell Schematic Sink', '{}', 20542);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('a190c47b-6fc7-436e-8b4c-6c55cc07921a', 'Auto-23 Belt Schematic Faucet', '{}', 12677);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('cfe037d9-0f5d-4f4f-8547-9ec35b8e9ee3', 'Auto-23 Belt Schematic Sink', '{}', 6122);

-- io
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('058b4726-8205-4e1d-b1b4-60a7dbd48da3', '308df2f2-37ea-48c1-aab5-127ce8be4843', 6, '{}', '4e95c448-bf09-4471-81c0-9e9f9449d2fe');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('cab046bc-2966-4417-a595-9b4d602505c2', '308df2f2-37ea-48c1-aab5-127ce8be4843', 2, '{}', '80ac90c8-5d0b-4dc9-a8c0-8ecff096c87e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4ee055c8-9e1f-492a-9d4a-b33fd2429c60', 'ae54a20c-ac6c-44bb-b245-277ac2aee144', 5, '{}', '7391e693-6d37-4ecb-8bec-b54798a21857');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('9bc1b8ae-5d6f-47f7-a20d-3b86897f45ca', 'ae54a20c-ac6c-44bb-b245-277ac2aee144', 9, '{}', 'd692d5d8-aeaa-48b5-af11-bef2f825ade7');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1d1b87c1-0ffc-4436-8138-f640ea7864ae', 'a25f293e-874d-43fa-834b-ed7e4ff11480', 3, '{}', 'cfe037d9-0f5d-4f4f-8547-9ec35b8e9ee3');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('d3303440-a118-455e-a010-87491f1af206', 'a25f293e-874d-43fa-834b-ed7e4ff11480', 4, '{}', 'a190c47b-6fc7-436e-8b4c-6c55cc07921a');
