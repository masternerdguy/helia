-- item families
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('ewar_cycle', 'Cycle Disruptor', '{}');
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('ewar_fcj', 'Fire Control Jammer', '{}');
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('ewar_r_mask', 'Regeneration Mask', '{}');
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('ewar_d_mask', 'Dissipation Mask', '{}');

-- item types (modules)
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('5c40d8ec-a269-4ae5-912f-1678b6c202a5', 'ewar_cycle', 'Basic Cycle Disruptor', '{"hp": 3, "rack": "b", "range": 7250, "volume": 5, "falloff": "reverse_linear", "cooldown": 37.5, "tracking": 325, "signal_flux": 10, "signal_gain": 25, "needs_target": true, "activation_heat": 6.5, "industrialmarket": {"maxprice": 8212, "minprice": 6798, "silosize": 500}, "activation_energy": 17}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('614434d0-eae6-4e41-bd02-9cac795b22c5', 'ewar_fcj', 'Basic Fire Control Jammer', '{"hp": 5, "rack": "b", "range": 5645, "volume": 4, "falloff": "linear", "cooldown": 32.3, "tracking": 267, "needs_target": true, "guidance_drift": 7.5, "tracking_drift": 15, "activation_heat": 6.5, "industrialmarket": {"maxprice": 4877, "minprice": 2119, "silosize": 500}, "activation_energy": 12}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('9bf8ba05-5888-49df-81bd-77969d3f3c37', 'ewar_r_mask', 'Basic Regeneration Mask', '{"hp": 4, "rack": "b", "range": 2418, "volume": 4, "falloff": "linear", "cooldown": 16.3, "mask_radius": 11.6, "needs_target": true, "activation_heat": 92.8, "industrialmarket": {"maxprice": 14778, "minprice": 9067, "silosize": 500}, "activation_energy": 110, "activation_gfx_effect": "basic_regen_mask"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('f40a5bf6-3124-4e3b-ade5-23a66d38339e', 'ewar_d_mask', 'Basic Dissipation Mask', '{"hp": 5, "rack": "b", "range": 2395, "volume": 3, "falloff": "linear", "cooldown": 15.4, "mask_radius": 11.9, "needs_target": true, "activation_heat": 93.3, "industrialmarket": {"maxprice": 16251, "minprice": 9390, "silosize": 500}, "activation_energy": 112, "activation_gfx_effect": "basic_dissip_mask"}');

-- item types (schematics)
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('34b6e8ab-cba8-4fe9-81b6-429ef997a01f', 'schematic', 'Basic Cycle Disruptor Schematic', '{"industrialmarket": {"maxprice": 38068, "minprice": 13660, "silosize": 100, "process_id": "3768279c-849f-4c96-a447-9d4bf4865271"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('32acc0b3-5941-4536-ac6a-3721b7fc0c55', 'schematic', 'Basic Fire Control Jammer Schematic', '{"industrialmarket": {"maxprice": 42564, "minprice": 10228, "silosize": 100, "process_id": "6db05608-bd73-4bc7-a2b0-1fb4c7531307"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('76c9e735-6e28-4ff4-8a51-d84fb070c03c', 'schematic', 'Basic Regeneration Mask Schematic', '{"industrialmarket": {"maxprice": 185696, "minprice": 60969, "silosize": 100, "process_id": "ea20caa5-1d32-40b7-8cdb-4948815c55a6"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('705fbf02-bad7-4000-a34a-db4f4313cd5f', 'schematic', 'Basic Dissipation Mask Schematic', '{"industrialmarket": {"maxprice": 215991, "minprice": 78813, "silosize": 100, "process_id": "ac62fe72-7db8-4e48-b659-de39360306fd"}}');

-- processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('3768279c-849f-4c96-a447-9d4bf4865271', 'Make Basic Cycle Disruptor', '{}', 105);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('826f1b25-38d2-4837-b336-f410425bf41c', 'Basic Cycle Disruptor Sink [wm]', '{}', 200);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9eefa1d3-1203-4fc8-b09a-3e4c967965a5', 'Basic Cycle Disruptor Schematic Faucet [wm]', '{}', 819);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('aad35b80-16ac-432d-bf4d-02950f37f3a0', 'Basic Cycle Disruptor Schematic Sink [wm]', '{}', 300);

INSERT INTO public.processes (id, name, meta, "time") VALUES ('6db05608-bd73-4bc7-a2b0-1fb4c7531307', 'Make Basic Fire Control Jammer', '{}', 124);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('8863ada9-26b8-4fc4-8cc7-b2b904bf6272', 'Basic Fire Control Jammer Sink [wm]', '{}', 90);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('e747ec63-3856-4fd1-8d4b-9eb8779f02e9', 'Basic Fire Control Jammer Schematic Faucet [wm]', '{}', 705);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('2fc8eca7-d513-45e4-808c-8d6cb21d8553', 'Basic Fire Control Jammer Schematic Sink [wm]', '{}', 687);

INSERT INTO public.processes (id, name, meta, "time") VALUES ('ea20caa5-1d32-40b7-8cdb-4948815c55a6', 'Make Basic Regeneration Mask', '{}', 177);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('f96c31f6-77e0-410e-8380-3fab636b3404', 'Basic Regeneration Mask Sink [wm]', '{}', 335);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('59a906a9-5159-4f40-8335-0a6a36ec087c', 'Basic Regeneration Mask Schematic Faucet [wm]', '{}', 1188);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('dcfdae71-052f-4adb-a8c4-4162ff65fa94', 'Basic Regeneration Mask Schematic Sink [wm]', '{}', 561);

INSERT INTO public.processes (id, name, meta, "time") VALUES ('ac62fe72-7db8-4e48-b659-de39360306fd', 'Make Basic Dissipation Mask', '{}', 15);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('db3d82a7-dcba-447f-ab5f-bb7d58061ea9', 'Basic Dissipation Mask Sink [wm]', '{}', 20);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('2efbb6eb-b6af-44aa-8cbb-f30e712dad2d', 'Basic Dissipation Mask Schematic Faucet [wm]', '{}', 63);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('27f815ce-47ae-4c59-b7cf-5c1a1840624b', 'Basic Dissipation Mask Schematic Sink [wm]', '{}', 88);

-- process inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('41ba2a8f-b483-4c60-8080-ccddf2add06b', '34b6e8ab-cba8-4fe9-81b6-429ef997a01f', 1, '{}', 'aad35b80-16ac-432d-bf4d-02950f37f3a0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1879e881-1868-4c2a-ad20-ae1c776feb96', '5c40d8ec-a269-4ae5-912f-1678b6c202a5', 19, '{}', '826f1b25-38d2-4837-b336-f410425bf41c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('559ff122-cb36-4c88-adf7-99fc50fb3df8', '24800206-2c58-45b0-8238-81974d0ebb3b', 368, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d05401c8-eb84-47e3-9bc5-04a6a8d9cddb', '61f52ba3-654b-45cf-88e3-33399d12350d', 2, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('2e347b40-a08b-4ddd-a412-c904852e871f', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 1, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6f1675cb-4ee9-4968-919e-afc92f101c48', '2ce48bef-f06b-4550-b20c-0e64864db051', 5, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('294b3eb4-787f-47b4-a252-8f64da26d617', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 1, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e569ffef-0cbf-4360-a8e8-2929b7fef10e', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 9, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cb59948e-8fee-4edf-a18a-8100c2edad64', '32acc0b3-5941-4536-ac6a-3721b7fc0c55', 4, '{}', '2fc8eca7-d513-45e4-808c-8d6cb21d8553');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e0b8236f-c9a4-4490-95ca-5635ac837abb', '614434d0-eae6-4e41-bd02-9cac795b22c5', 16, '{}', '8863ada9-26b8-4fc4-8cc7-b2b904bf6272');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e2c4e539-2847-46ee-abeb-59422b28da33', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 28, '{}', '6db05608-bd73-4bc7-a2b0-1fb4c7531307');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d306730d-8517-4b2e-9460-4aada080db2e', '24800206-2c58-45b0-8238-81974d0ebb3b', 208, '{}', '6db05608-bd73-4bc7-a2b0-1fb4c7531307');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8a1fca40-1203-4e58-9ac9-8e37bb69a36e', '11688112-f3d4-4d30-864a-684a8b96ea23', 6, '{}', '6db05608-bd73-4bc7-a2b0-1fb4c7531307');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ccc2a14f-d999-4a5b-8337-71672164078f', '56617d30-6c30-425c-84bf-2484ae8c1156', 6, '{}', '6db05608-bd73-4bc7-a2b0-1fb4c7531307');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('25099548-06ec-47a1-9678-da75360f7594', '2ce48bef-f06b-4550-b20c-0e64864db051', 5, '{}', '6db05608-bd73-4bc7-a2b0-1fb4c7531307');

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8d4968a9-00a1-4c54-bf66-841d5d8adc59', '76c9e735-6e28-4ff4-8a51-d84fb070c03c', 3, '{}', 'dcfdae71-052f-4adb-a8c4-4162ff65fa94');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f1fac957-986b-454d-b980-9673a308d386', '9bf8ba05-5888-49df-81bd-77969d3f3c37', 46, '{}', 'f96c31f6-77e0-410e-8380-3fab636b3404');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('98b1edc7-1a23-4fe8-9186-80609a1b05b0', '61f52ba3-654b-45cf-88e3-33399d12350d', 23, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('fcd1e5db-e3ee-419b-b150-38e0f79e5a26', '2ce48bef-f06b-4550-b20c-0e64864db051', 26, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('2b1b65b1-7b0e-40ad-a72e-b0c53ab17ea9', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 27, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('85774bcf-24f6-4169-b665-bc68c1b82797', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 18, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('dfe50658-fc51-4fe5-97cd-7278d58af34f', '11688112-f3d4-4d30-864a-684a8b96ea23', 26, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ee3e30dd-4582-4f2d-9a25-a1e0bf501d37', '56617d30-6c30-425c-84bf-2484ae8c1156', 24, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a64e8e45-3376-4611-b4b1-f943cdc97416', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 117, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6ba64517-606a-4c30-bd5f-9ba62e77e0a3', '24800206-2c58-45b0-8238-81974d0ebb3b', 690, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ca8abbe7-4f9f-4189-9562-21c08ae6e5d9', '705fbf02-bad7-4000-a34a-db4f4313cd5f', 7, '{}', '27f815ce-47ae-4c59-b7cf-5c1a1840624b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e73fb510-6091-41a3-8782-4ea9d93b1191', 'f40a5bf6-3124-4e3b-ade5-23a66d38339e', 23, '{}', 'db3d82a7-dcba-447f-ab5f-bb7d58061ea9');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e47af0f9-ca8c-42a5-8c29-3365f9bb5ec0', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 26, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5ba84bb0-27b3-404c-9d8a-1eb8700844d6', '2ce48bef-f06b-4550-b20c-0e64864db051', 30, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('70bbd2ab-6a76-40b5-b019-3e4360d3e26b', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 29, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3d51e612-da6b-4052-add1-ecab832f7c76', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 124, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('66740e43-8dea-414d-b8cf-7094792c74f3', '24800206-2c58-45b0-8238-81974d0ebb3b', 991, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d81b5170-05e7-4da7-8608-ae35d77ed24e', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 27, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ac261efc-f4d1-46cb-b394-f33473ede707', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 31, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('972ea48a-358e-428a-82a1-6b0a7781f1f8', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 23, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');

-- process outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('2e07ac97-6798-40ed-86c9-340423e27f71', '34b6e8ab-cba8-4fe9-81b6-429ef997a01f', 8, '{}', '9eefa1d3-1203-4fc8-b09a-3e4c967965a5');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('1583610f-6e33-433c-a9b3-0a115fe29e7c', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 368, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('2b6a90cb-331d-4c82-a489-8ce82f25971d', '5c40d8ec-a269-4ae5-912f-1678b6c202a5', 10, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('59c117a3-6e55-41ca-9aa8-e405838513db', '32acc0b3-5941-4536-ac6a-3721b7fc0c55', 5, '{}', 'e747ec63-3856-4fd1-8d4b-9eb8779f02e9');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('783b6114-4c26-40bd-b5dd-a7a0fec60bce', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 208, '{}', '6db05608-bd73-4bc7-a2b0-1fb4c7531307');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('c98a5192-a152-448f-9be4-c1b3acb4d1c2', '614434d0-eae6-4e41-bd02-9cac795b22c5', 25, '{}', '6db05608-bd73-4bc7-a2b0-1fb4c7531307');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('99cd43da-0a9b-4c06-9793-cf1cc70d8eb5', '76c9e735-6e28-4ff4-8a51-d84fb070c03c', 4, '{}', '59a906a9-5159-4f40-8335-0a6a36ec087c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('397087b4-c910-4a94-b964-de3800054505', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 690, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('f6190ca9-9768-4a37-9fd2-305cff7145a8', '9bf8ba05-5888-49df-81bd-77969d3f3c37', 30, '{}', 'ea20caa5-1d32-40b7-8cdb-4948815c55a6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('797ce102-3f77-492b-806e-833d1b451976', '705fbf02-bad7-4000-a34a-db4f4313cd5f', 3, '{}', '2efbb6eb-b6af-44aa-8cbb-f30e712dad2d');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('7f209dc6-1cf6-48a9-9743-6df4674d651c', 'f40a5bf6-3124-4e3b-ade5-23a66d38339e', 30, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('7b101a1a-c8f0-42ed-a5df-61a8aba1a042', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 991, '{}', 'ac62fe72-7db8-4e48-b659-de39360306fd');
