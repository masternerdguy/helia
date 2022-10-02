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
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('653a77e5-cc40-4ed9-a941-20c29b4d7cdc', '6cfba75-2cb5-47f5-9806-a92c4909c2ef', 5400, '{}', '916c852e-525d-4eeb-8781-32bb42179874');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d9d066cf-dfe1-4aaf-9c2c-cc8fcff196a5', '56617d30-6c30-425c-84bf-2484ae8c1156', 700, '{}', '916c852e-525d-4eeb-8781-32bb42179874');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('c521d4b3-911f-4dfd-86fa-3f2a00d221d4', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 24, '{}', '916c852e-525d-4eeb-8781-32bb42179874');

