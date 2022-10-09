-- ship type (existing freighter type reserved for an intermediate class)
INSERT INTO public.shiptypes (id, name) VALUES ('d7b35504-b958-4f74-9c3e-7e9b619a9495', 'Superfreighter');

-- item types
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('a17c43ba-25ec-4236-b0a2-d0d433e75097', 'ship', 'Zebra', '{"volume": 1273598, "shiptemplateid": "ab6714d2-0574-492b-aeb9-6da2727faa06", "industrialmarket": {"maxprice": 525738429, "minprice": 172463080, "silosize": 5}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('a5404d21-ef22-45c6-855e-2de0d53ac851', 'ship', 'Fetch', '{"volume": 1282675, "shiptemplateid": "11ae39d0-0b49-4276-971b-9655ee135b27", "industrialmarket": {"maxprice": 637180314, "minprice": 260281735, "silosize": 5}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d2cd985d-a2fe-4491-8b89-a68f18c6350c', 'ship', 'Station Assembly Yard', '{"volume": 4712683429, "shiptemplateid": "b4c9a253-7aab-4321-b112-a63e11ba0e13", "industrialmarket": {"maxprice": 174294800, "minprice": 62638170, "silosize": 5}}');

-- ship templates
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('ab6714d2-0574-492b-aeb9-6da2727faa06', '2022-09-11 17:47:27.185141-04', 'Zebra', 'Zebra', 135, 0.7, 358997, 0.5, 6900, 513, 2200, 72415, 63471, 7781, 117, 1694, 171, 'd7b35504-b958-4f74-9c3e-7e9b619a9495', '{"a_slots": [{"hp_pos": [10, 45], "volume": 325, "mod_family": "missile"}, {"hp_pos": [-5, 0], "volume": 325, "mod_family": "missile"}, {"hp_pos": [10, -45], "volume": 250, "mod_family": "utility"}, {"hp_pos": [-10, 45], "volume": 250, "mod_family": "utility"}, {"hp_pos": [-10, -45], "volume": 235, "mod_family": "utility"}, {"hp_pos": [5, 0], "volume": 235, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 350, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 350, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 240, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 240, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 500, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 500, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 500, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 500, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 500, "mod_family": "any"}]}', 636785, 'a17c43ba-25ec-4236-b0a2-d0d433e75097', true, 'basic-wreck', 'basic_explosion');
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('11ae39d0-0b49-4276-971b-9655ee135b27', '2022-09-11 21:27:39.754702-04', 'Fetch', 'Fetch', 138, 1.4, 294231, 0.2, 1572, 178, 5520, 54397, 127386, 3663, 89, 2142, 216, 'd7b35504-b958-4f74-9c3e-7e9b619a9495', '{"a_slots": [{"hp_pos": [10, 0], "volume": 400, "mod_family": "gun"}, {"hp_pos": [5, 0], "volume": 400, "mod_family": "gun"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "utility"}, {"hp_pos": [-5, 0], "volume": 400, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 350, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 350, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 350, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 325, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 325, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 325, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 325, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 325, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 325, "mod_family": "any"}]}', 460227, 'a5404d21-ef22-45c6-855e-2de0d53ac851', true, 'basic-wreck', 'basic_explosion');
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('b4c9a253-7aab-4321-b112-a63e11ba0e13', '2022-09-11 21:48:29.48882-04', 'Station Assembly Yard', 'Station Assembly Yard', 100, 0, 0, 0, 1, 0.01, 1, 1, 1, 1000, 13.4, 1, 0.01, 'e364a553-1dc5-4e8d-9195-0ca4989bec49', '{"a_slots": [], "b_slots": [], "c_slots": []}', 100000000, 'd2cd985d-a2fe-4491-8b89-a68f18c6350c', false, 'basic-wreck', 'basic_explosion');

-- helper view to find ships that don't have schematics
CREATE OR REPLACE VIEW public.vw_ships_needsschematics
 AS
 SELECT h.id,
    h.family,
    h.name,
    h.meta
   FROM ( SELECT itemtypes.id,
            itemtypes.family,
            itemtypes.name,
            itemtypes.meta
           FROM itemtypes
          WHERE NOT (itemtypes.id IN ( SELECT processoutputs.itemtypeid
                   FROM processoutputs
                  WHERE (processoutputs.processid IN ( SELECT ((itemtypes_1.meta::json -> 'industrialmarket'::text) ->> 'process_id'::text)::uuid AS proccessid
                           FROM itemtypes itemtypes_1))))) h
  WHERE h.family::text = 'ship'::text;

ALTER TABLE public.vw_ships_needsschematics
    OWNER TO heliaagent;

-- giga sized power cells
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('54c29ce8-67c7-46c7-8c02-c737eed3143c', 'power_cell', '1 GWH Cell', '{"hp": 1, "volume": 4250, "industrialmarket": {"maxprice": 17000000, "minprice": 1275000, "silosize": 50000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('3935523b-3e38-485b-935f-e790758ce36b', 'depleted_cell', 'Depleted 1 GWH Cell', '{"hp": 1, "volume": 4250, "industrialmarket": {"maxprice": 850000, "minprice": 255000, "silosize": 50000}}');

-- solar charge giga cell
INSERT INTO public.processes (id, name, meta, "time") VALUES ('35ea1aa3-6c0f-461a-aab3-f03530f83ce8', 'Solar Charge 1 GWH Cell', '{}', 77139);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1300caaf-be3e-4456-acf2-e227ef5e36e5', '3935523b-3e38-485b-935f-e790758ce36b', 10, '{}', '35ea1aa3-6c0f-461a-aab3-f03530f83ce8');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('273c84f6-4fd8-481d-bfea-0039f68b684d', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 10, '{}', '35ea1aa3-6c0f-461a-aab3-f03530f83ce8');

-- 3 PJ pellet
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('c6cfba75-2cb5-47f5-9806-a92c4909c2ef', 'fuel', '3 PJ Pellet', '{"hp": 1, "volume": 200, "fuelconversion": 55000, "industrialmarket": {"maxprice": 50000, "minprice": 10000, "silosize": 85000}}');

-- quick charge giga cell
INSERT INTO public.processes (id, name, meta, "time") VALUES ('916c852e-525d-4eeb-8781-32bb42179874', '3 TJ Quick Charge 1 GWH Cell', '{}', 700);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('7f84a485-a4ff-4570-b233-1a30a5ea0733', '3935523b-3e38-485b-935f-e790758ce36b', 27, '{}', '916c852e-525d-4eeb-8781-32bb42179874');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('653a77e5-cc40-4ed9-a941-20c29b4d7cdc', 'c6cfba75-2cb5-47f5-9806-a92c4909c2ef', 5400, '{}', '916c852e-525d-4eeb-8781-32bb42179874');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d9d066cf-dfe1-4aaf-9c2c-cc8fcff196a5', '56617d30-6c30-425c-84bf-2484ae8c1156', 700, '{}', '916c852e-525d-4eeb-8781-32bb42179874');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('c521d4b3-911f-4dfd-86fa-3f2a00d221d4', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 24, '{}', '916c852e-525d-4eeb-8781-32bb42179874');

-- assembly yard faucet
INSERT INTO public.processes (id, name, meta, "time") VALUES ('355afc0a-2469-4752-9239-bb776524a34b', 'Station Assembly Yard Faucet', '{}', 1274219);

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('ca305269-791e-40cb-b992-409a2ae5b2cc', 'd2cd985d-a2fe-4491-8b89-a68f18c6350c', 1, '{}', '355afc0a-2469-4752-9239-bb776524a34b');

-- super freighter faucets (faction ships, no schematics)
INSERT INTO public.processes (id, name, meta, "time") VALUES ('95f9515c-e439-4d61-b0df-f90860318e25', 'Zebra Faucet', '{}', 112713);
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('20917b7f-0d23-4c45-94a6-a68518e3309f', 'a17c43ba-25ec-4236-b0a2-d0d433e75097', 1, '{}', '95f9515c-e439-4d61-b0df-f90860318e25');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('8cf78edf-2b7a-4727-8d7a-0d0d6ecf379f', 'Fetch Faucet', '{}', 97648);
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('b86dee5f-d268-4aa2-8053-12f06952a735', 'a5404d21-ef22-45c6-855e-2de0d53ac851', 1, '{}', '8cf78edf-2b7a-4727-8d7a-0d0d6ecf379f');

-- fair make 1 GWH cell
INSERT INTO public.processes (id, name, meta, "time") VALUES ('45d9d2d4-8076-4a67-9168-ddf6ade545fe', 'Make 1 GWH Cell', '{}', 23892);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f713696a-d4c4-4e4c-a55c-ec2ad56388a0', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 1000, '{}', '45d9d2d4-8076-4a67-9168-ddf6ade545fe');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c42bed15-e6c7-4870-bbf5-fcced62de28e', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 5225, '{}', '45d9d2d4-8076-4a67-9168-ddf6ade545fe');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bed04b99-1380-4a0c-b25a-fb1eef8cc4cc', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 1, '{}', '45d9d2d4-8076-4a67-9168-ddf6ade545fe');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('f81379d4-d4cd-40d0-87ae-4fe42de99f09', '3935523b-3e38-485b-935f-e790758ce36b', 3, '{}', '45d9d2d4-8076-4a67-9168-ddf6ade545fe');

-- fair make 3 PJ pellet
INSERT INTO public.processes (id, name, meta, "time") VALUES ('1dedc612-65bf-4fc5-9238-25374bd43e59', 'Make 3 PJ Pellet', '{}', 8350);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('256b102f-8e1a-44dd-a0c1-ca278bac7c51', 'da619e43-4832-42b8-ad03-5eb42441a403', 23883, '{}', '1dedc612-65bf-4fc5-9238-25374bd43e59');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('92668de8-bbea-47a4-b55b-2c6e24714ec4', 'c6cfba75-2cb5-47f5-9806-a92c4909c2ef', 100, '{}', '1dedc612-65bf-4fc5-9238-25374bd43e59');

-- fair cell schematic
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('682e5b2f-7940-43c6-a6fc-8be58604a947', 'schematic', '1 GWH Cell Schematic', '{"volume": 1, "industrialmarket": {"maxprice": 4150000, "minprice": 2952000, "silosize": 100, "process_id": "45d9d2d4-8076-4a67-9168-ddf6ade545fe"}}');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('053554d1-a94c-4c46-a1ca-79106cc23db4', '1 GWH Schematic Sink', '{}', 26739);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('5db46b9e-d5e6-4e2d-b8f7-a5341311f897', '1 GWH Schematic Faucet', '{}', 12911);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('696d7aa9-ee65-469f-8972-ab69765a978c', '682e5b2f-7940-43c6-a6fc-8be58604a947', 1, '{}', '053554d1-a94c-4c46-a1ca-79106cc23db4');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('14277f26-e685-4163-8737-c8520614d1b8', '682e5b2f-7940-43c6-a6fc-8be58604a947', 1, '{}', '5db46b9e-d5e6-4e2d-b8f7-a5341311f897');

-- fair pellet schematic
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('ff4095cd-9cad-4ad4-9b0e-1cb1836d694e', 'schematic', '3 PJ Pellet Schematic', '{"volume": 1, "industrialmarket": {"maxprice": 7590000, "minprice": 4290000, "silosize": 100, "process_id": "1dedc612-65bf-4fc5-9238-25374bd43e59"}}');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('526931d1-680f-4279-b85b-0967b9a292f1', '3 PJ Pellet Schematic Sink', '{}', 23742);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b3d19718-c3bd-4f5b-96ae-25355e3f9b7d', '3 PJ Pellet Schematic Faucet', '{}', 11850);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a831e2af-5ce6-4f1f-9b7c-2acceb00c718', 'ff4095cd-9cad-4ad4-9b0e-1cb1836d694e', 1, '{}', '526931d1-680f-4279-b85b-0967b9a292f1');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('c5a45f5d-3482-494d-8277-cf9991ef32fa', 'ff4095cd-9cad-4ad4-9b0e-1cb1836d694e', 1, '{}', 'b3d19718-c3bd-4f5b-96ae-25355e3f9b7d');
