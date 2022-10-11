-- item families
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('therm_cap', 'Thermal Capacitor', '{}');
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('burst_reactor', 'Burst Fusion Reactor', '{}');

-- item types
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('6f402b12-c262-4eb8-b5c8-46c5b58fcef7', 'therm_cap', 'Basic Thermal Capacitor', '{"hp": 21, "rack": "c", "volume": 5, "heat_cap_max_add": 200, "industrialmarket": {"maxprice": 1316, "minprice": 944, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('95fc28b8-9464-474d-a274-82e62d059199', 'therm_cap', 'Small Thermal Capacitor', '{"hp": 44, "rack": "c", "volume": 23, "heat_cap_max_add": 415, "industrialmarket": {"maxprice": 3371, "minprice": 1688, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d43e9d2f-8185-4efd-afba-4fad3f419f76', 'therm_cap', 'XL Thermal Capacitor', '{"hp": 216, "rack": "c", "volume": 217, "heat_cap_max_add": 3372, "industrialmarket": {"maxprice": 413219, "minprice": 288920, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('f5833591-6b7c-4b80-9904-81001d51b888', 'burst_reactor', 'Basic Burst Fusion Reactor', '{"hp": 6, "rack": "b", "volume": 7, "leakage": 0.35, "cooldown": 212, "needs_target": false, "activation_heat": 216, "max_fuel_volume": 1, "industrialmarket": {"maxprice": 4461, "minprice": 3375, "silosize": 1000}, "activation_energy": 45}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('186561d1-41ab-4525-b8b9-534d78248f10', 'burst_reactor', 'XL Burst Fusion Reactor', '{"hp": 26, "rack": "b", "volume": 215, "leakage": 0.16, "cooldown": 86, "needs_target": false, "activation_heat": 5672, "max_fuel_volume": 1000, "industrialmarket": {"maxprice": 697265, "minprice": 491606, "silosize": 1000}, "activation_energy": 1782}');

-- schematics
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('9886093f-688e-445f-890d-d4c7f5e30a27', 'schematic', 'Basic Thermal Capacitor Schematic', '{"industrialmarket": {"maxprice": 35293, "minprice": 7484, "silosize": 100, "process_id": "9a817a9a-0b13-4dff-b0de-2015d71ad57f"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('ab91e34c-ecce-4f47-96e4-84746adefcb1', 'schematic', 'Small Thermal Capacitor Schematic', '{"industrialmarket": {"maxprice": 54653, "minprice": 22615, "silosize": 100, "process_id": "4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('de936f8d-2725-44e3-81e2-b6ad9a6cb4d5', 'schematic', 'Basic Burst Fusion Reactor Schematic', '{"industrialmarket": {"maxprice": 95148, "minprice": 18955, "silosize": 100, "process_id": "93cf9897-d809-4f16-8d7c-f791e7af773a"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('3ae90719-d56c-4c26-89be-5c379eb65f87', 'schematic', 'XL Burst Fusion Reactor Schematic', '{"industrialmarket": {"maxprice": 6773763, "minprice": 1510957, "silosize": 100, "process_id": "677aefc9-d5e5-4ef1-a885-fc772d98605c"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('fd616ba6-fb65-4c8a-ba4c-b9ae93bf5f09', 'schematic', 'XL Thermal Capacitor Schematic', '{"industrialmarket": {"maxprice": 13767001, "minprice": 5389649, "silosize": 100, "process_id": "b59bf67d-b643-4a70-b419-ce21227083cf"}}');

-- processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9a817a9a-0b13-4dff-b0de-2015d71ad57f', 'Make Basic Thermal Capacitor', '{}', 195);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('de864e70-4c36-48e0-9ad4-85b0d3936188', 'Basic Thermal Capacitor Sink [wm]', '{}', 370);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('ce72330e-8e46-4d80-a562-644b706934ed', 'Basic Thermal Capacitor Schematic Faucet [wm]', '{}', 1489);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('e7d11865-332f-4f99-9a74-c0e350ad7c3d', 'Basic Thermal Capacitor Schematic Sink [wm]', '{}', 687);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d', 'Make Small Thermal Capacitor', '{}', 2667);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('ac45078a-f3d6-4956-8188-9b053b2124e1', 'Small Thermal Capacitor Sink [wm]', '{}', 1629);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('d59f65ad-fbbd-4b91-ba0a-a994cf6f27ca', 'Small Thermal Capacitor Schematic Faucet [wm]', '{}', 11055);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('34efd490-fcc0-49ce-8723-04a24f678b8c', 'Small Thermal Capacitor Schematic Sink [wm]', '{}', 12732);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('93cf9897-d809-4f16-8d7c-f791e7af773a', 'Make Basic Burst Fusion Reactor', '{}', 510);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b95a734e-658b-49da-af20-52dc720504a7', 'Basic Burst Fusion Reactor Sink [wm]', '{}', 440);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('1ff86d95-5469-463c-8a0a-a29155cf154f', 'Basic Burst Fusion Reactor Schematic Faucet [wm]', '{}', 1665);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('f39774bf-59a7-4473-b464-5056e65ec9e7', 'Basic Burst Fusion Reactor Schematic Sink [wm]', '{}', 3881);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('677aefc9-d5e5-4ef1-a885-fc772d98605c', 'Make XL Burst Fusion Reactor', '{}', 18893);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('8996b556-bfbb-40ec-8b3f-3887b15b6f18', 'XL Burst Fusion Reactor Sink [wm]', '{}', 16118);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('213f1b33-62cb-4740-a8b9-ac90a250ee0a', 'XL Burst Fusion Reactor Schematic Faucet [wm]', '{}', 115355);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('63f8ac39-8b2e-499d-abdc-e27d6a1f0239', 'XL Burst Fusion Reactor Schematic Sink [wm]', '{}', 48631);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b59bf67d-b643-4a70-b419-ce21227083cf', 'Make XL Thermal Capacitor', '{}', 836);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('77f1192e-c8f3-41b1-8990-9b1e01324b7a', 'XL Thermal Capacitor Sink [wm]', '{}', 434);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('028a33b0-34e9-4aae-aef7-d867813e0349', 'XL Thermal Capacitor Schematic Faucet [wm]', '{}', 6390);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('2f57ebcf-af1d-4178-88e6-5943cb79be94', 'XL Thermal Capacitor Schematic Sink [wm]', '{}', 3895);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('dbe6d047-a4fa-4f5b-bbd1-ddef40a9a09e', '9886093f-688e-445f-890d-d4c7f5e30a27', 3, '{}', 'e7d11865-332f-4f99-9a74-c0e350ad7c3d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9759ea3d-53fd-45c0-97a1-11406aba55e8', '6f402b12-c262-4eb8-b5c8-46c5b58fcef7', 101, '{}', 'de864e70-4c36-48e0-9ad4-85b0d3936188');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ec305875-0a3a-4b55-b604-ecded1d7ba80', '11688112-f3d4-4d30-864a-684a8b96ea23', 3, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e153142b-42e2-46f1-ad4a-a1c23bdbbb0b', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 1, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('41fe22c7-3e1f-4492-8bab-2ffa3486e622', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 12, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('280af85a-ca64-4715-b3e4-cba823c66c3b', '24800206-2c58-45b0-8238-81974d0ebb3b', 281, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8c83f36b-9ef1-434e-aa4b-8a902bedd71c', '2ce48bef-f06b-4550-b20c-0e64864db051', 2, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cf87d440-2cf2-41cc-9b9c-328dc5e8309f', '66b7a322-8cfc-4467-9410-492e6b58f159', 4, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('50825f2b-d196-4c82-97f9-868b26082032', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 1, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3d124c4e-bc37-462e-a55b-6bef0b95d02f', 'ab91e34c-ecce-4f47-96e4-84746adefcb1', 5, '{}', '34efd490-fcc0-49ce-8723-04a24f678b8c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6b9f5769-ac23-4fcb-90bc-686925e1135d', '95fc28b8-9464-474d-a274-82e62d059199', 127, '{}', 'ac45078a-f3d6-4956-8188-9b053b2124e1');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b4783774-98a7-46bf-9fa9-fdc9f9c243d7', '24800206-2c58-45b0-8238-81974d0ebb3b', 639, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b6e6f992-00bb-46e9-a619-363dc4be281e', '66b7a322-8cfc-4467-9410-492e6b58f159', 8, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f6ae30c8-7112-498a-ae01-a1a957879ddd', '11688112-f3d4-4d30-864a-684a8b96ea23', 7, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cfb0aa56-9079-485a-b305-64b62d2098bd', '56617d30-6c30-425c-84bf-2484ae8c1156', 8, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a11d921f-6737-4926-b2f5-ca90f0ce42fe', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 10, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c00e2f69-9b2e-42ad-a60a-96d32737dbae', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 6, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a192bf0c-ef9c-48ba-b45e-d16358a6ad19', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 5, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('caad00ef-612c-46a8-9268-da77ff351c6b', '61f52ba3-654b-45cf-88e3-33399d12350d', 10, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('25772fcd-046c-4c58-ac8f-311081451bac', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 8, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5f3e3cc5-c711-43c6-b6fb-f08731101e1f', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 34, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('2220d39e-4a55-445a-a226-a0e0ef696fce', 'de936f8d-2725-44e3-81e2-b6ad9a6cb4d5', 4, '{}', 'f39774bf-59a7-4473-b464-5056e65ec9e7');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e6c4f0b0-2ef9-4cc3-9d65-3bc07a981d60', 'f5833591-6b7c-4b80-9904-81001d51b888', 108, '{}', 'b95a734e-658b-49da-af20-52dc720504a7');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c1c11970-7d02-4363-8a45-73fc8d05335a', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 93, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e4f91654-aa94-44b7-950b-72269776f0c5', '24800206-2c58-45b0-8238-81974d0ebb3b', 486, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('22e0f4e4-933a-4c9d-a152-7c1c1fbc1608', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 23, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bce44aa0-88e8-46d7-a5c4-5a009ed65e59', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 25, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e264c0f8-e91a-43f3-ae0e-492738da0f6b', '2ce48bef-f06b-4550-b20c-0e64864db051', 28, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('73cc1062-65ec-4e07-9aa2-29bbf67f7f53', '66b7a322-8cfc-4467-9410-492e6b58f159', 14, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('033ced4c-1477-4020-b42a-3da9f5cdf1c7', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 20, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4c86915d-228f-41d6-bd88-e09dc7d94f61', '61f52ba3-654b-45cf-88e3-33399d12350d', 23, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('65548613-56da-43a4-9172-6013b7c34773', '3ae90719-d56c-4c26-89be-5c379eb65f87', 9, '{}', '63f8ac39-8b2e-499d-abdc-e27d6a1f0239');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('671b288c-fb41-4e12-a070-b8b9f4ac576f', '186561d1-41ab-4525-b8b9-534d78248f10', 71, '{}', '8996b556-bfbb-40ec-8b3f-3887b15b6f18');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('271764b7-faf1-4124-a49f-52b7920d2f77', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 169, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('afeebdda-c064-4033-a582-b3463ff6a01c', '897a3ede-d682-4e1e-bbf9-4605dee60597', 42, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e066e7c2-e4cb-4ecf-a628-4f57ef11de10', '24800206-2c58-45b0-8238-81974d0ebb3b', 119461, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e8fff586-c890-4ff1-b000-1674a2a69a80', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 43, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cbfc49e0-f6e0-49a2-91b5-1b1e1b1f0664', '9c3795a9-43fb-4f26-95cd-655b20f5347a', 39, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('65b3b6b1-09f7-4b46-82f9-bce2e8d76ab3', '11688112-f3d4-4d30-864a-684a8b96ea23', 46, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0e1d7053-8132-44a9-931e-864fe6510311', '66b7a322-8cfc-4467-9410-492e6b58f159', 50, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4394bdce-df6b-4e32-b14a-2aa66f77d4b4', 'fd616ba6-fb65-4c8a-ba4c-b9ae93bf5f09', 4, '{}', '2f57ebcf-af1d-4178-88e6-5943cb79be94');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b98f9929-b379-46d0-8e49-2069ffda7755', 'd43e9d2f-8185-4efd-afba-4fad3f419f76', 150, '{}', '77f1192e-c8f3-41b1-8990-9b1e01324b7a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8e62db95-1e77-4271-9022-f4488000e1c2', '24800206-2c58-45b0-8238-81974d0ebb3b', 124814, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('110ab943-c1f2-4b38-8ad1-79293815c087', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 1697, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bc926ebb-0618-4afb-aaaf-2e6523a58eb0', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 1688, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1a034dbe-44e5-4f9d-be3a-01f02d59aaf6', '56617d30-6c30-425c-84bf-2484ae8c1156', 1705, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c11467b4-b84c-4f93-950c-9662f7c1ad1d', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 1713, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6e2d37d1-173b-4116-98bd-49cbba58db8f', '2ce48bef-f06b-4550-b20c-0e64864db051', 1697, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bdaa8b0a-47d2-45ee-85e9-fbbc5cbd2efd', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 6891, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('2cb113c1-cf46-4f10-a061-c0f6bbc79b2b', '9886093f-688e-445f-890d-d4c7f5e30a27', 2, '{}', 'ce72330e-8e46-4d80-a562-644b706934ed');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('da806288-3340-472b-a6ff-c32aeab0b984', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 281, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('51a7ce1d-c355-4cdc-a5eb-c1aa44dfe1f7', '6f402b12-c262-4eb8-b5c8-46c5b58fcef7', 55, '{}', '9a817a9a-0b13-4dff-b0de-2015d71ad57f');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('0ddacb13-0f85-48ff-b145-e72e48661911', 'ab91e34c-ecce-4f47-96e4-84746adefcb1', 8, '{}', 'd59f65ad-fbbd-4b91-ba0a-a994cf6f27ca');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('08d45479-7373-496d-b6fa-c57cb016541a', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 639, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('94cb622a-364f-4a86-8d6a-f6873e553c8c', '95fc28b8-9464-474d-a274-82e62d059199', 70, '{}', '4b6bedc9-9248-4fa0-9cb2-fd0efa7a7d5d');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('fcb79e57-0f22-47e7-b9ec-8ec9dee0496a', 'de936f8d-2725-44e3-81e2-b6ad9a6cb4d5', 7, '{}', '1ff86d95-5469-463c-8a0a-a29155cf154f');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('d969f6f8-5e1d-4ff3-88af-0a290ede1d53', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 486, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('9b4c439a-146e-4ea2-91a1-edd3d92c840e', 'f5833591-6b7c-4b80-9904-81001d51b888', 80, '{}', '93cf9897-d809-4f16-8d7c-f791e7af773a');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('354bc922-d323-4917-a2b5-624b948963cd', '3ae90719-d56c-4c26-89be-5c379eb65f87', 7, '{}', '213f1b33-62cb-4740-a8b9-ac90a250ee0a');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('e39399a6-8260-4505-a713-7c5193b6f169', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 119461, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('86f3ebcb-06f0-4d45-b4a4-bcfa592fb387', '186561d1-41ab-4525-b8b9-534d78248f10', 45, '{}', '677aefc9-d5e5-4ef1-a885-fc772d98605c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('17840995-e9cd-432a-82a9-b59232f83ec7', 'fd616ba6-fb65-4c8a-ba4c-b9ae93bf5f09', 8, '{}', '028a33b0-34e9-4aae-aef7-d867813e0349');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('ed3df6e1-328b-4f36-86b6-dd5de66d67f3', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 124814, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('bf3e3f6d-478f-4b48-b459-d2113db25a0a', 'd43e9d2f-8185-4efd-afba-4fad3f419f76', 80, '{}', 'b59bf67d-b643-4a70-b419-ce21227083cf');
