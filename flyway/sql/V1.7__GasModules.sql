-- item family
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('utility_wisper', 'Wisp Collector', '{}');

-- items
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('572b92e7-149d-4129-a9a6-91b2e92999a8', 'utility_wisper', 'Basic Wisp Collector', '{"hp": 8, "rack": "a", "volume": 13, "cooldown": 3.6, "intake_area": 3.1, "can_mine_gas": true, "can_mine_ice": false, "can_mine_ore": false, "needs_target": false, "activation_heat": 6, "industrialmarket": {"maxprice": 14577, "minprice": 11324, "silosize": 350}, "activation_energy": 35}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('693dd357-35b9-4e24-a3c2-a2f24613162a', 'utility_wisper', 'Small Wisp Collector', '{"hp": 11, "rack": "a", "volume": 22, "cooldown": 4.5, "intake_area": 5.8, "can_mine_gas": true, "can_mine_ice": false, "can_mine_ore": false, "needs_target": false, "activation_heat": 9, "industrialmarket": {"maxprice": 37990, "minprice": 16321, "silosize": 325}, "activation_energy": 43}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('073858f3-ec93-41bd-9582-ea12405edc7e', 'utility_wisper', 'Medium Wisp Collector', '{"hp": 16, "rack": "a", "volume": 40, "cooldown": 5.7, "intake_area": 9.3, "can_mine_gas": true, "can_mine_ice": false, "can_mine_ore": false, "needs_target": false, "activation_heat": 15, "industrialmarket": {"maxprice": 73452, "minprice": 28929, "silosize": 315}, "activation_energy": 62}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('bd13334d-a66e-4294-86ea-b6e8348497ea', 'utility_wisper', 'Large Wisp Collector', '{"hp": 19, "rack": "a", "volume": 75, "cooldown": 6.9, "intake_area": 17.5, "can_mine_gas": true, "can_mine_ice": false, "can_mine_ore": false, "needs_target": false, "activation_heat": 21, "industrialmarket": {"maxprice": 189335, "minprice": 59667, "silosize": 300}, "activation_energy": 83}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('59ac90f5-77a3-4f36-b46d-dbb1effd30c5', 'utility_wisper', 'XL Wisp Collector', '{"hp": 21, "rack": "a", "volume": 105, "cooldown": 8.2, "intake_area": 38.6, "can_mine_gas": true, "can_mine_ice": false, "can_mine_ore": false, "needs_target": false, "activation_heat": 22, "industrialmarket": {"maxprice": 492610, "minprice": 103239, "silosize": 200}, "activation_energy": 91}');

-- avoid outpost kits in module schematic check
CREATE OR REPLACE VIEW public.vw_modules_needsschematics
 AS
 SELECT itemtypes.id,
    itemtypes.family,
    itemtypes.name,
    itemtypes.meta
   FROM itemtypes
  WHERE NOT (itemtypes.id IN ( SELECT processoutputs.itemtypeid
           FROM processoutputs
          WHERE (processoutputs.processid IN ( SELECT ((itemtypes_1.meta::json -> 'industrialmarket'::text) ->> 'process_id'::text)::uuid AS proccessid
                   FROM itemtypes itemtypes_1
                  WHERE itemtypes_1.family::text = 'schematic'::text)))) AND itemtypes.name::text !~~ '%`%'::text AND (itemtypes.family::text <> ALL (ARRAY['nothing'::character varying::text, 'ore'::character varying::text, 'repair_kit'::character varying::text, 'mod_kit'::character varying::text, 'ammunition'::character varying::text, 'fuel'::character varying::text, 'ice'::character varying::text, 'trade_good'::character varying::text, 'ship'::character varying::text, 'schematic'::character varying::text, 'power_cell'::character varying::text, 'depleted_cell'::character varying::text, 'widget'::character varying::text, 'outpost_kit'::character varying::text]));

ALTER TABLE public.vw_modules_needsschematics
    OWNER TO heliaagent;

-- schematic item types
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('5f16b98c-1aa1-4f43-80fb-50e698f86a7c', 'schematic', 'Medium Wisp Collector Schematic', '{"industrialmarket": {"maxprice": 192826, "minprice": 92480, "silosize": 100, "process_id": "71ae064d-8d43-4980-b581-3d1863dc95fb"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('e8d8fb80-7294-4ac0-b57e-e2bca7254d98', 'schematic', 'Basic Wisp Collector Schematic', '{"industrialmarket": {"maxprice": 11239, "minprice": 2221, "silosize": 100, "process_id": "2463bee2-0568-4f72-bdc2-04e7603409af"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('9c025aac-988e-45d7-9be0-66498061be05', 'schematic', 'Small Wisp Collector Schematic', '{"industrialmarket": {"maxprice": 95303, "minprice": 30110, "silosize": 100, "process_id": "2af4c35b-cd21-4c10-b48d-94f1bdac0492"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('617f3203-2614-410c-821e-b88cbf20a4b9', 'schematic', 'Large Wisp Collector Schematic', '{"industrialmarket": {"maxprice": 449707, "minprice": 80631, "silosize": 100, "process_id": "c9e999cd-cdda-4acc-ae5e-6f547bad169d"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d5215ce1-a521-4d8b-ac50-0ff47489c1e0', 'schematic', 'XL Wisp Collector Schematic', '{"industrialmarket": {"maxprice": 260693, "minprice": 118849, "silosize": 100, "process_id": "606b5007-6768-4e99-8018-6af8ebe9b404"}}');

-- processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('71ae064d-8d43-4980-b581-3d1863dc95fb', 'Make Medium Wisp Collector', '{}', 276);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('15f3b138-a5fa-4cfa-93a2-4a56a64133c8', 'Medium Wisp Collector Sink [wm]', '{}', 259);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('2181bc90-7ba2-4c4b-b35c-946e52b94e9d', 'Medium Wisp Collector Schematic Faucet [wm]', '{}', 900);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('71db2814-2eb8-4b64-974e-5ace2d542f73', 'Medium Wisp Collector Schematic Sink [wm]', '{}', 1650);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('2463bee2-0568-4f72-bdc2-04e7603409af', 'Make Basic Wisp Collector', '{}', 52);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('ccb260e0-5447-4242-ba9d-88515e9aa31a', 'Basic Wisp Collector Sink [wm]', '{}', 90);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('3de4f67d-d99a-400d-baf0-426ef1e34a1c', 'Basic Wisp Collector Schematic Faucet [wm]', '{}', 323);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('be206558-af43-4317-becc-d672aed83908', 'Basic Wisp Collector Schematic Sink [wm]', '{}', 319);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('2af4c35b-cd21-4c10-b48d-94f1bdac0492', 'Make Small Wisp Collector', '{}', 232);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('0dae58d0-64a3-4d8d-b1e9-6bee8793707a', 'Small Wisp Collector Sink [wm]', '{}', 288);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('664ae1de-7b67-4f7f-b174-0f5d2ee2bfde', 'Small Wisp Collector Schematic Faucet [wm]', '{}', 647);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('4906c9e0-0de9-4e16-bf7e-ae35ae847ff9', 'Small Wisp Collector Schematic Sink [wm]', '{}', 656);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('c9e999cd-cdda-4acc-ae5e-6f547bad169d', 'Make Large Wisp Collector', '{}', 2261);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('20e35286-cc99-4cc3-85bf-3bdb767d7e05', 'Large Wisp Collector Sink [wm]', '{}', 1580);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('eaffe773-ae65-4ef5-99da-01fa57b8455e', 'Large Wisp Collector Schematic Faucet [wm]', '{}', 7464);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('1512c0c3-b31b-4904-8fe6-dacd200f3689', 'Large Wisp Collector Schematic Sink [wm]', '{}', 13680);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('606b5007-6768-4e99-8018-6af8ebe9b404', 'Make XL Wisp Collector', '{}', 805);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('85bae98e-68a1-47a2-b508-8b945da4a5a7', 'XL Wisp Collector Sink [wm]', '{}', 1104);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('6ba18058-ef7d-40f3-9176-29d8016627d4', 'XL Wisp Collector Schematic Faucet [wm]', '{}', 3342);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('fb28debd-019d-4ea2-90c6-f60e56ea96b8', 'XL Wisp Collector Schematic Sink [wm]', '{}', 4280);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('60923132-af3e-4d07-a056-469c75e34ae9', '5f16b98c-1aa1-4f43-80fb-50e698f86a7c', 8, '{}', '71db2814-2eb8-4b64-974e-5ace2d542f73');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b32cc130-ec58-4edf-b231-e8e38d787d76', '073858f3-ec93-41bd-9582-ea12405edc7e', 5, '{}', '15f3b138-a5fa-4cfa-93a2-4a56a64133c8');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a7a211cf-4fd8-4294-86ba-9274e63e7575', '2ce48bef-f06b-4550-b20c-0e64864db051', 22, '{}', '71ae064d-8d43-4980-b581-3d1863dc95fb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('01589f7e-bc2d-43f0-a1e8-f5def4209867', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 21, '{}', '71ae064d-8d43-4980-b581-3d1863dc95fb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f1a18cb7-8834-4226-a770-46ab8559e140', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 26, '{}', '71ae064d-8d43-4980-b581-3d1863dc95fb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8a52062c-3a4b-4bb6-8c68-3ba36c3ce297', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 102, '{}', '71ae064d-8d43-4980-b581-3d1863dc95fb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9a5b81be-fad8-4763-b147-725ceb8dcec2', '24800206-2c58-45b0-8238-81974d0ebb3b', 824, '{}', '71ae064d-8d43-4980-b581-3d1863dc95fb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('32c8b535-f6b1-4fe8-ad01-8a207ea20ed4', 'e8d8fb80-7294-4ac0-b57e-e2bca7254d98', 6, '{}', 'be206558-af43-4317-becc-d672aed83908');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4c2addc6-e08e-4336-8a67-ff7baa046be4', '572b92e7-149d-4129-a9a6-91b2e92999a8', 2, '{}', 'ccb260e0-5447-4242-ba9d-88515e9aa31a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ec951774-926a-466f-ae9d-cf0ea6682f2f', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 4, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('475643c8-1643-4e8e-bf71-3379f09697c1', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 2, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('24b0493e-f438-45db-823a-4755d4ce75cf', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 5, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c1aa5dd4-24f2-4f33-b287-63d399bd06b4', '24800206-2c58-45b0-8238-81974d0ebb3b', 123, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('17ba8329-0548-488f-ab83-e05cf0ebd3b5', '11688112-f3d4-4d30-864a-684a8b96ea23', 2, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b1b9e52e-dc1d-44e6-8ba9-f273d4d5ac46', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 1, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('06e4120b-40fa-4dd3-a1ba-c31ba99b1ef5', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 1, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('980c070a-84c4-4f1c-a3fc-f2c466cfefe3', '2ce48bef-f06b-4550-b20c-0e64864db051', 1, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a1f3912b-f612-4322-b613-d6bc585dba8e', '9c025aac-988e-45d7-9be0-66498061be05', 1, '{}', '4906c9e0-0de9-4e16-bf7e-ae35ae847ff9');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1b94efa3-a831-49e3-b821-c5c6cb8ef0f4', '693dd357-35b9-4e24-a3c2-a2f24613162a', 6, '{}', '0dae58d0-64a3-4d8d-b1e9-6bee8793707a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('37983543-c918-4d9a-8853-3543e8dfc4df', '56617d30-6c30-425c-84bf-2484ae8c1156', 24, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('dc7d53ce-8b3b-49ad-aafa-293c73fa73f5', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 104, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3eed50f0-a5df-4395-aa42-ca419b23b873', '24800206-2c58-45b0-8238-81974d0ebb3b', 794, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e1b1408a-41cc-4164-9a24-764a5865e55e', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 21, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('58f20685-14ef-4d11-b455-3dc3af68eca5', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 23, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4af06ad3-453a-4193-9734-5855eb090349', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 20, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('69c77fb2-f1fc-4c0c-b164-f16aaa0efa17', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 29, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('56fe162f-14b4-4dfe-824c-d55fe5f8c664', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 19, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('078379a9-70f8-4b04-be7b-44aeed6b1bd8', '61f52ba3-654b-45cf-88e3-33399d12350d', 21, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b769f432-7fa6-4cbd-b311-647d0c1def8e', '617f3203-2614-410c-821e-b88cbf20a4b9', 4, '{}', '1512c0c3-b31b-4904-8fe6-dacd200f3689');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0bf30d35-04a7-4bcd-83f0-332b06139aa6', 'bd13334d-a66e-4294-86ea-b6e8348497ea', 23, '{}', '20e35286-cc99-4cc3-85bf-3bdb767d7e05');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c873a435-7ed6-4df7-a01e-48a9d969da3f', '2ce48bef-f06b-4550-b20c-0e64864db051', 35, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b4697ed6-33de-417a-b846-2831ea953ee0', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 35, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d2c20fc1-40d2-44f3-9000-5563d11118c7', '61f52ba3-654b-45cf-88e3-33399d12350d', 27, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b540ea51-6446-45b7-a778-ecdec31cf07c', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 36, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ccde24ed-6720-4521-a0c5-1434267fa422', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 146, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('36efd893-edd3-41c8-abce-3e13408021e1', '24800206-2c58-45b0-8238-81974d0ebb3b', 4834, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d03d1003-87ce-4681-a9fc-2b2a718d04cf', '11688112-f3d4-4d30-864a-684a8b96ea23', 36, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1f50d8b4-a06e-480c-a2c2-50d5e4e79726', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 43, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4a3fd397-5e00-479c-8be6-da05c3e4d1ac', '66b7a322-8cfc-4467-9410-492e6b58f159', 37, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('357f2355-060f-468f-b530-248c625f2f11', 'd5215ce1-a521-4d8b-ac50-0ff47489c1e0', 4, '{}', 'fb28debd-019d-4ea2-90c6-f60e56ea96b8');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3de90d7d-d7d9-4c3d-8993-7ee7c37ab5bd', '59ac90f5-77a3-4f36-b46d-dbb1effd30c5', 5, '{}', '85bae98e-68a1-47a2-b508-8b945da4a5a7');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bd233e45-dad8-4afe-9e1c-1cb3a00cc502', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 42, '{}', '606b5007-6768-4e99-8018-6af8ebe9b404');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1fa1220b-fe27-4105-9d07-608bcf98a92b', '2ce48bef-f06b-4550-b20c-0e64864db051', 42, '{}', '606b5007-6768-4e99-8018-6af8ebe9b404');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('abb33e97-3086-462f-9e98-f605ad37e963', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 37, '{}', '606b5007-6768-4e99-8018-6af8ebe9b404');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('aca2bfb4-17f9-4137-b47c-59f1fb2c5d85', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 180, '{}', '606b5007-6768-4e99-8018-6af8ebe9b404');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ef121180-8903-420b-833d-f915213e6bf9', '24800206-2c58-45b0-8238-81974d0ebb3b', 2788, '{}', '606b5007-6768-4e99-8018-6af8ebe9b404');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('94c3609b-2ee5-47d7-99e9-578b5fded13f', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 52, '{}', '606b5007-6768-4e99-8018-6af8ebe9b404');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('64db3728-7a94-4aa2-ab34-433e70597ee2', '5f16b98c-1aa1-4f43-80fb-50e698f86a7c', 1, '{}', '2181bc90-7ba2-4c4b-b35c-946e52b94e9d');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('803946f8-4ad8-4d8f-a8fa-553f1bb247c9', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 824, '{}', '71ae064d-8d43-4980-b581-3d1863dc95fb');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('c2d814cd-d795-4415-91b7-d62f42ee4976', '073858f3-ec93-41bd-9582-ea12405edc7e', 7, '{}', '71ae064d-8d43-4980-b581-3d1863dc95fb');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('141526bb-d112-4ed8-9652-e366793d2089', 'e8d8fb80-7294-4ac0-b57e-e2bca7254d98', 1, '{}', '3de4f67d-d99a-400d-baf0-426ef1e34a1c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a9306b62-eaa6-4e75-be1f-28d32c0ea761', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 123, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('0e0f40ae-4209-413d-9f4a-8eac0ed50630', '572b92e7-149d-4129-a9a6-91b2e92999a8', 2, '{}', '2463bee2-0568-4f72-bdc2-04e7603409af');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('546d9d6e-7e9a-44df-b93b-c0893a1bd9e2', '9c025aac-988e-45d7-9be0-66498061be05', 1, '{}', '664ae1de-7b67-4f7f-b174-0f5d2ee2bfde');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('b4a5d1c3-4e87-437f-97e1-b758a4b778cb', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 794, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('395cc791-64b9-48d7-852f-0209ae672a97', '693dd357-35b9-4e24-a3c2-a2f24613162a', 9, '{}', '2af4c35b-cd21-4c10-b48d-94f1bdac0492');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('85702a8c-f0b6-4f9c-80df-648ac02e8321', '617f3203-2614-410c-821e-b88cbf20a4b9', 7, '{}', 'eaffe773-ae65-4ef5-99da-01fa57b8455e');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('3868b7c1-5127-476a-a429-16204058ceb0', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 4834, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('dd3dde57-a627-482a-8e75-71e00133a761', 'bd13334d-a66e-4294-86ea-b6e8348497ea', 15, '{}', 'c9e999cd-cdda-4acc-ae5e-6f547bad169d');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('57cda1ec-6ae0-43f7-aa1d-770e37f63aa6', 'd5215ce1-a521-4d8b-ac50-0ff47489c1e0', 9, '{}', '6ba18058-ef7d-40f3-9176-29d8016627d4');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('6a288b55-18f8-4f9d-a865-c73fa8aaebe6', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 2788, '{}', '606b5007-6768-4e99-8018-6af8ebe9b404');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('1294ba25-ef92-4de9-a9aa-55107bfb346f', '59ac90f5-77a3-4f36-b46d-dbb1effd30c5', 5, '{}', '606b5007-6768-4e99-8018-6af8ebe9b404');
