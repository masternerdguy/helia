-- schematics
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('cd3bc3ff-43c1-41ad-b182-15ed2fb141de', 'schematic', 'XL Gauss Rifle Schematic', '{"industrialmarket": {"maxprice": 1035921, "minprice": 385411, "silosize": 100, "process_id": "a4278f85-45d8-47df-90a3-f3f97442e335"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('305de9a6-b62b-4135-ad36-c17e00ce086b', 'schematic', 'Basic Auto-23 Cannon Schematic', '{"industrialmarket": {"maxprice": 1574690, "minprice": 280643, "silosize": 100, "process_id": "8f1bdca9-8353-4316-9b27-d3e4187056da"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('2f3014c7-4f4f-42c3-87fc-c15499dbcfc1', 'schematic', 'Dual Auto-23 Cannon Schematic', '{"industrialmarket": {"maxprice": 1597002, "minprice": 250188, "silosize": 100, "process_id": "d1b2da39-5ee6-483b-8cb1-d1c292f12e5e"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('b62cad23-3bdb-497e-ba48-a8a461f094e9', 'schematic', 'Basic HGPM-15 Launcher Schematic', '{"industrialmarket": {"maxprice": 361714, "minprice": 72146, "silosize": 100, "process_id": "1f6d58d5-4eed-4b8c-a8db-ee261e433d4e"}}');

-- processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('a4278f85-45d8-47df-90a3-f3f97442e335', 'Make XL Gauss Rifle', '{}', 332);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('173d3e1d-feba-46b5-96b8-144b058a6974', 'XL Gauss Rifle Sink [wm]', '{}', 637);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9f483968-bcc6-41df-bb97-dbc5ec58d6e6', 'XL Gauss Rifle Schematic Faucet [wm]', '{}', 998);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('7bfa2cce-f63c-4ea3-830f-65980f25f84c', 'XL Gauss Rifle Schematic Sink [wm]', '{}', 835);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('8f1bdca9-8353-4316-9b27-d3e4187056da', 'Make Basic Auto-23 Cannon', '{}', 12894);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b7f630bb-3765-45e4-9bb7-9b0841e65631', 'Basic Auto-23 Cannon Sink [wm]', '{}', 25693);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('07671e60-1727-4819-ad7f-398e60cf7041', 'Basic Auto-23 Cannon Schematic Faucet [wm]', '{}', 57918);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('126dc511-7acc-41a6-8cc1-0e54029cca21', 'Basic Auto-23 Cannon Schematic Sink [wm]', '{}', 87388);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('d1b2da39-5ee6-483b-8cb1-d1c292f12e5e', 'Make Dual Auto-23 Cannon', '{}', 1317);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('5e006e10-8394-43f8-bd62-4b23fc9ca574', 'Dual Auto-23 Cannon Sink [wm]', '{}', 2392);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('46e62ca8-d8f3-420b-bb66-b6fc727d3e32', 'Dual Auto-23 Cannon Schematic Faucet [wm]', '{}', 7300);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('7867814c-3c3d-46b4-ad54-76af62410f20', 'Dual Auto-23 Cannon Schematic Sink [wm]', '{}', 3855);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('1f6d58d5-4eed-4b8c-a8db-ee261e433d4e', 'Make Basic HGPM-15 Launcher', '{}', 19792);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('5bd8db3e-d42b-484f-907a-98ac32e6ad0b', 'Basic HGPM-15 Launcher Sink [wm]', '{}', 13710);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b51f9a7b-d10d-4703-a395-a9b5e5cc37dd', 'Basic HGPM-15 Launcher Schematic Faucet [wm]', '{}', 63019);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('21987c7e-ec0a-4508-97f0-b2e46c75f3d5', 'Basic HGPM-15 Launcher Schematic Sink [wm]', '{}', 43002);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('19dffef4-97fc-454f-84be-f795f3fc17f4', 'cd3bc3ff-43c1-41ad-b182-15ed2fb141de', 5, '{}', '7bfa2cce-f63c-4ea3-830f-65980f25f84c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ce9c5b7a-bdb7-48e7-8ddb-9ed8fb8c9ba1', '1c83a2c1-1723-4b4d-ad06-626fef52c34b', 17, '{}', '173d3e1d-feba-46b5-96b8-144b058a6974');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a57da7fe-53b8-4453-bdd7-a3736e543116', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 268, '{}', 'a4278f85-45d8-47df-90a3-f3f97442e335');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b837afb5-6fdf-4ed1-a293-cbbefdc71964', '61f52ba3-654b-45cf-88e3-33399d12350d', 257, '{}', 'a4278f85-45d8-47df-90a3-f3f97442e335');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('aefba56f-559a-46f2-aadb-ad0fe4cdff78', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 997, '{}', 'a4278f85-45d8-47df-90a3-f3f97442e335');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4ac9b495-4e93-40b5-8bb5-65db4c776e48', '24800206-2c58-45b0-8238-81974d0ebb3b', 1739, '{}', 'a4278f85-45d8-47df-90a3-f3f97442e335');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('479d3fa1-b76c-4725-b794-06492fda0df0', '11688112-f3d4-4d30-864a-684a8b96ea23', 252, '{}', 'a4278f85-45d8-47df-90a3-f3f97442e335');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a69dc9c4-3d8f-4fbe-baf5-4cfd2c45e475', '305de9a6-b62b-4135-ad36-c17e00ce086b', 8, '{}', '126dc511-7acc-41a6-8cc1-0e54029cca21');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ee204a7f-5eb8-4e21-aacb-11895b440e5c', '4e32b02a-0caf-4f32-9390-0fdd011a7759', 63, '{}', 'b7f630bb-3765-45e4-9bb7-9b0841e65631');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('275b2890-55d5-4e71-a94a-744ddbef4a9c', '56617d30-6c30-425c-84bf-2484ae8c1156', 136, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cc66fe18-dcc6-4fc8-bcad-03bacbbb9b4a', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 559, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e4ebaea5-1964-4ca1-88e7-d82b73988547', '24800206-2c58-45b0-8238-81974d0ebb3b', 18255, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('7394531b-6e9f-40a3-885b-72aef20cb3bf', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 151, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('20c0f61c-e2d9-48a9-97f0-b1f3613ebcb8', '66b7a322-8cfc-4467-9410-492e6b58f159', 141, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('646b5cee-d215-40e8-91d5-9e196d8e35b9', '2ce48bef-f06b-4550-b20c-0e64864db051', 157, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0d85e361-0124-4aed-82c2-6ac38cb8753f', '61f52ba3-654b-45cf-88e3-33399d12350d', 136, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('142e3c11-fb96-46a9-9e8f-0a6b905a01cd', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 131, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('53cba5bb-ee71-4ea8-b89a-fb3e01f17a3d', '2f3014c7-4f4f-42c3-87fc-c15499dbcfc1', 5, '{}', '7867814c-3c3d-46b4-ad54-76af62410f20');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('2c567023-3482-4036-97e9-23a4c43b9b0d', '20f3cdc7-00cf-4409-a2b3-670786a0cbba', 9, '{}', '5e006e10-8394-43f8-bd62-4b23fc9ca574');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('de37c90d-a3de-4f77-a721-de1bc890f7d8', '24800206-2c58-45b0-8238-81974d0ebb3b', 2779, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('7ba79aba-6ed3-4489-9ff6-dae32e565e32', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 510, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b6da3b4e-f514-484f-8d88-f6b2ce06bcf8', '11688112-f3d4-4d30-864a-684a8b96ea23', 119, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6e0eeb62-c3f9-438a-909c-c1aba14b5e68', '66b7a322-8cfc-4467-9410-492e6b58f159', 127, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('82ea1a36-758f-4dc9-8d21-8df20fe5e3e1', '2ce48bef-f06b-4550-b20c-0e64864db051', 120, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0fbbc0e4-272a-4836-9a8e-9c6df62908ba', '61f52ba3-654b-45cf-88e3-33399d12350d', 119, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('fbaf438d-a8ee-4afb-a753-98f8fb31edbd', '56617d30-6c30-425c-84bf-2484ae8c1156', 125, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('82cd00cf-a470-40df-9df9-81fd4aba00a4', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 127, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3a3f4385-5391-4e07-b7c9-fb97b6e2d5a7', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 102, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('94cb12ed-d58b-47a4-b6ba-15c771eeb5eb', 'b62cad23-3bdb-497e-ba48-a8a461f094e9', 2, '{}', '21987c7e-ec0a-4508-97f0-b2e46c75f3d5');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('13f42497-dc2f-4463-893d-f495c09eda1f', '038fc30e-bb2b-4b0f-aa37-c07e140ff6ed', 76, '{}', '5bd8db3e-d42b-484f-907a-98ac32e6ad0b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a9034d47-4772-47e4-ae3f-b32fb8d251a0', '24800206-2c58-45b0-8238-81974d0ebb3b', 4601, '{}', '1f6d58d5-4eed-4b8c-a8db-ee261e433d4e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ba45b7ce-c7d5-4840-a43d-4008de276bbd', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 282, '{}', '1f6d58d5-4eed-4b8c-a8db-ee261e433d4e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5db15060-af8c-4a0a-adce-e780c14b2023', '2ce48bef-f06b-4550-b20c-0e64864db051', 70, '{}', '1f6d58d5-4eed-4b8c-a8db-ee261e433d4e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d8ceebc5-e389-4977-91dc-1c8a87918185', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 88, '{}', '1f6d58d5-4eed-4b8c-a8db-ee261e433d4e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8141688f-a747-4c38-a9f1-1b8cdc018045', '9c3795a9-43fb-4f26-95cd-655b20f5347a', 69, '{}', '1f6d58d5-4eed-4b8c-a8db-ee261e433d4e');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('f4f8b1b2-5457-43d2-b108-c4615d569532', 'cd3bc3ff-43c1-41ad-b182-15ed2fb141de', 1, '{}', '9f483968-bcc6-41df-bb97-dbc5ec58d6e6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('01100afc-de8f-4246-8584-553742ab09aa', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 1739, '{}', 'a4278f85-45d8-47df-90a3-f3f97442e335');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a5bc748f-dfde-497d-82b3-f0cfb6097203', '1c83a2c1-1723-4b4d-ad06-626fef52c34b', 10, '{}', 'a4278f85-45d8-47df-90a3-f3f97442e335');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('01480d5e-39a8-4933-9c5f-c0d63c774075', '305de9a6-b62b-4135-ad36-c17e00ce086b', 1, '{}', '07671e60-1727-4819-ad7f-398e60cf7041');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a0edd34e-9acf-441b-b385-ed3a98ea3fd7', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 18255, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('4116df1f-4f24-4d88-b8cd-02b9174159a2', '4e32b02a-0caf-4f32-9390-0fdd011a7759', 45, '{}', '8f1bdca9-8353-4316-9b27-d3e4187056da');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('dc5755ba-ce0e-45f1-b4bc-9d35af5d2760', '2f3014c7-4f4f-42c3-87fc-c15499dbcfc1', 9, '{}', '46e62ca8-d8f3-420b-bb66-b6fc727d3e32');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a6647fb5-c31b-4efe-ac76-cfdc2daad138', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 2779, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('1f9823fe-ece5-4b60-b952-694df48b0b7e', '20f3cdc7-00cf-4409-a2b3-670786a0cbba', 10, '{}', 'd1b2da39-5ee6-483b-8cb1-d1c292f12e5e');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('388e1828-af5c-45e3-8a19-8b877e858f78', 'b62cad23-3bdb-497e-ba48-a8a461f094e9', 9, '{}', 'b51f9a7b-d10d-4703-a395-a9b5e5cc37dd');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('72d8d055-e705-42a4-af5c-5e88ef1c9a16', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 4601, '{}', '1f6d58d5-4eed-4b8c-a8db-ee261e433d4e');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('3febb2b1-0b21-4e27-b127-01e1ed163655', '038fc30e-bb2b-4b0f-aa37-c07e140ff6ed', 40, '{}', '1f6d58d5-4eed-4b8c-a8db-ee261e433d4e');
