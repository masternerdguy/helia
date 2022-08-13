-- item families
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('ewar_cycle', 'Cycle Disruptor', '{}');
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('ewar_fcj', 'Fire Control Jammer', '{}');
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('ewar_r_mask', 'Regeneration Mask', '{}');
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('ewar_d_mask', 'Dissipation Mask', '{}');

-- item types (modules)
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('5c40d8ec-a269-4ae5-912f-1678b6c202a5', 'ewar_cycle', 'Basic Cycle Disruptor', '{"hp": 3, "rack": "b", "range": 7250, "volume": 5, "falloff": "reverse_linear", "cooldown": 37.5, "tracking": 325, "signal_flux": 10, "signal_gain": 25, "needs_target": true, "activation_heat": 6.5, "industrialmarket": {"maxprice": 8212, "minprice": 6798, "silosize": 500}, "activation_energy": 17}');

-- item types (schematics)
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('34b6e8ab-cba8-4fe9-81b6-429ef997a01f', 'schematic', 'Basic Cycle Disruptor Schematic', '{"industrialmarket": {"maxprice": 38068, "minprice": 13660, "silosize": 100, "process_id": "3768279c-849f-4c96-a447-9d4bf4865271"}}');

-- processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('3768279c-849f-4c96-a447-9d4bf4865271', 'Make Basic Cycle Disruptor', '{}', 105);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('826f1b25-38d2-4837-b336-f410425bf41c', 'Basic Cycle Disruptor Sink [wm]', '{}', 200);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9eefa1d3-1203-4fc8-b09a-3e4c967965a5', 'Basic Cycle Disruptor Schematic Faucet [wm]', '{}', 819);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('aad35b80-16ac-432d-bf4d-02950f37f3a0', 'Basic Cycle Disruptor Schematic Sink [wm]', '{}', 300);

-- process inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('41ba2a8f-b483-4c60-8080-ccddf2add06b', '34b6e8ab-cba8-4fe9-81b6-429ef997a01f', 1, '{}', 'aad35b80-16ac-432d-bf4d-02950f37f3a0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1879e881-1868-4c2a-ad20-ae1c776feb96', '5c40d8ec-a269-4ae5-912f-1678b6c202a5', 19, '{}', '826f1b25-38d2-4837-b336-f410425bf41c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('559ff122-cb36-4c88-adf7-99fc50fb3df8', '24800206-2c58-45b0-8238-81974d0ebb3b', 368, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d05401c8-eb84-47e3-9bc5-04a6a8d9cddb', '61f52ba3-654b-45cf-88e3-33399d12350d', 2, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('2e347b40-a08b-4ddd-a412-c904852e871f', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 1, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6f1675cb-4ee9-4968-919e-afc92f101c48', '2ce48bef-f06b-4550-b20c-0e64864db051', 5, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('294b3eb4-787f-47b4-a252-8f64da26d617', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 1, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e569ffef-0cbf-4360-a8e8-2929b7fef10e', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 9, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');

-- process outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('2e07ac97-6798-40ed-86c9-340423e27f71', '34b6e8ab-cba8-4fe9-81b6-429ef997a01f', 8, '{}', '9eefa1d3-1203-4fc8-b09a-3e4c967965a5');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('1583610f-6e33-433c-a9b3-0a115fe29e7c', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 368, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('2b6a90cb-331d-4c82-a489-8ce82f25971d', '5c40d8ec-a269-4ae5-912f-1678b6c202a5', 10, '{}', '3768279c-849f-4c96-a447-9d4bf4865271');
