-- insert item family
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('fuel_loader', 'Field Fuel Loader', '{}');

-- insert item types
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('57c375ed-0ffa-4a93-9806-3901cd6e66d7', 'fuel_loader', 'Basic Field Fuel Loader', '{"hp": 7, "rack": "b", "volume": 3, "leakage": 0.2, "cooldown": 212, "needs_target": false, "activation_heat": 128, "max_fuel_volume": 1, "industrialmarket": {"maxprice": 1388, "minprice": 965, "silosize": 1000}, "activation_energy": 17}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('2e11fdec-21d0-4561-b827-ef7e89f695fd', 'fuel_loader', 'XL Field Fuel Loader', '{"hp": 35, "rack": "b", "volume": 120, "leakage": 0.1, "cooldown": 318, "needs_target": false, "activation_heat": 2899, "max_fuel_volume": 1000, "industrialmarket": {"maxprice": 223795, "minprice": 99628, "silosize": 1000}, "activation_energy": 965}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('7acd336d-0a4c-41e4-a5d2-fdeb44007fac', 'schematic', 'Basic Field Fuel Loader Schematic', '{"industrialmarket": {"maxprice": 33634, "minprice": 6381, "silosize": 100, "process_id": "23cef584-f3aa-47dc-b11a-87b4debaec80"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('76a3f99d-c7af-4024-b276-87a225455dab', 'schematic', 'XL Field Fuel Loader Schematic', '{"industrialmarket": {"maxprice": 182486, "minprice": 50495, "silosize": 100, "process_id": "fd1bc6f0-59e5-4aad-9d64-13a48b86b74f"}}');

-- insert process definitions
INSERT INTO public.processes (id, name, meta, "time") VALUES ('23cef584-f3aa-47dc-b11a-87b4debaec80', 'Make Basic Field Fuel Loader', '{}', 97);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('e1bbb4ab-9178-4b7d-b0f8-9365449d1fd6', 'Basic Field Fuel Loader Sink [wm]', '{}', 70);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('3533818e-3a3b-41dd-af00-78d800c5921f', 'Basic Field Fuel Loader Schematic Faucet [wm]', '{}', 746);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('6db24bfb-92ae-4156-a120-5887137c3d3b', 'Basic Field Fuel Loader Schematic Sink [wm]', '{}', 330);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('fd1bc6f0-59e5-4aad-9d64-13a48b86b74f', 'Make XL Field Fuel Loader', '{}', 631);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('0c3df30f-7ff0-4a1d-8f3d-b86f76f3c30a', 'XL Field Fuel Loader Sink [wm]', '{}', 512);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('0693c405-8a08-4728-8f80-c649fccbbcab', 'XL Field Fuel Loader Schematic Faucet [wm]', '{}', 4313);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('ebb77519-cb29-4acd-ab17-5dded28b6adc', 'XL Field Fuel Loader Schematic Sink [wm]', '{}', 4121);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e810886f-42e6-4f33-98d5-fe0f9a769e07', '7acd336d-0a4c-41e4-a5d2-fdeb44007fac', 4, '{}', '6db24bfb-92ae-4156-a120-5887137c3d3b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('60cf75ef-096a-4e6c-b39f-2ef93c06933a', '57c375ed-0ffa-4a93-9806-3901cd6e66d7', 17, '{}', 'e1bbb4ab-9178-4b7d-b0f8-9365449d1fd6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ed4f848f-b53f-44ed-9ae1-f9607d31558a', '11688112-f3d4-4d30-864a-684a8b96ea23', 10, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c7a39633-2461-4566-9152-c93e61018ba2', '56617d30-6c30-425c-84bf-2484ae8c1156', 8, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8c83d81b-acbd-4d84-99d9-b02254cb5dfe', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 12, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('74d6f426-c036-493b-a210-326601a54bd9', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 8, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9f80a84c-1977-4f62-8aad-822e17c32299', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 48, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('71a2d330-98f1-42bc-bf28-9b002a0ed9b9', '24800206-2c58-45b0-8238-81974d0ebb3b', 53, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8e0bb388-74e6-4de9-9af0-fb8ec961dafd', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 12, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a6f8632c-f594-46c0-bdf3-4d1fdadbc7bc', '76a3f99d-c7af-4024-b276-87a225455dab', 2, '{}', 'ebb77519-cb29-4acd-ab17-5dded28b6adc');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1e9951f5-1eb1-4df2-9ae1-17f12c75472c', '2e11fdec-21d0-4561-b827-ef7e89f695fd', 4, '{}', '0c3df30f-7ff0-4a1d-8f3d-b86f76f3c30a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('20cff344-0a40-4d57-8196-8e8688911a60', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 89, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e2625420-c007-41b9-b6cc-5258658ece76', '24800206-2c58-45b0-8238-81974d0ebb3b', 1521, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('452b71f3-c2e8-4c7b-830c-7e1cff071ece', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 24, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('2d762403-3165-4b5b-9926-3e259930d661', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 25, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4aabfb85-f7a0-48a7-b9fe-fb382d4298b5', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 21, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6afae5dd-3f74-4fdf-9fdd-d5f6400e3c94', '56617d30-6c30-425c-84bf-2484ae8c1156', 23, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('db78d366-5260-405d-aabc-00c473e3bb8c', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 15, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a43803ab-9b06-488d-a40c-c5f4d4e26db6', '2ce48bef-f06b-4550-b20c-0e64864db051', 24, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('fa011e0e-d63f-418b-a80c-0cd14b1e41db', '7acd336d-0a4c-41e4-a5d2-fdeb44007fac', 5, '{}', '3533818e-3a3b-41dd-af00-78d800c5921f');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('3771d3f7-bbee-4a95-a4a9-3715b5f1ae24', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 53, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('98612a27-7bc5-41cd-8ffe-d7bb18ac64cb', '57c375ed-0ffa-4a93-9806-3901cd6e66d7', 30, '{}', '23cef584-f3aa-47dc-b11a-87b4debaec80');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('9c82a67f-b53a-49a4-8173-cb5d8469fd8d', '76a3f99d-c7af-4024-b276-87a225455dab', 4, '{}', '0693c405-8a08-4728-8f80-c649fccbbcab');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('b0609b4c-0321-4278-8db3-df0e96695682', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 1521, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('aaf51d1d-7c03-4d59-91a8-9b301b22dc04', '2e11fdec-21d0-4561-b827-ef7e89f695fd', 3, '{}', 'fd1bc6f0-59e5-4aad-9d64-13a48b86b74f');

