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

