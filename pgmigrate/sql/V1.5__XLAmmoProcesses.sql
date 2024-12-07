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
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('02315008-630a-4aff-87d8-0f8f0f050959', '368e03a0-bfef-43df-b686-2a5f279549d1', 195, '{}', '0f9d4229-45a0-4923-9af6-d56a3ca5195c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('c991d482-4282-47cd-a6c7-ce9397dd4772', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 167, '{}', '0f9d4229-45a0-4923-9af6-d56a3ca5195c');
