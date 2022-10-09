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

INSERT INTO public.processes (id, name, meta, "time") VALUES ('053554d1-a94c-4c46-a1ca-79106cc23db4', '1 GWH Cell Schematic Sink', '{}', 26739);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('5db46b9e-d5e6-4e2d-b8f7-a5341311f897', '1 GWH Cell Schematic Faucet', '{}', 12911);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('66cf4115-59cc-47f3-881d-9e2023b56048', '682e5b2f-7940-43c6-a6fc-8be58604a947', 1, '{}', '053554d1-a94c-4c46-a1ca-79106cc23db4');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('2f1edd63-859b-4f6c-a8ce-055406a2714a', '682e5b2f-7940-43c6-a6fc-8be58604a947', 1, '{}', '5db46b9e-d5e6-4e2d-b8f7-a5341311f897');

-- fair pellet schematic
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('ff4095cd-9cad-4ad4-9b0e-1cb1836d694e', 'schematic', '3 PJ Pellet Schematic', '{"volume": 1, "industrialmarket": {"maxprice": 7590000, "minprice": 4290000, "silosize": 100, "process_id": "1dedc612-65bf-4fc5-9238-25374bd43e59"}}');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('526931d1-680f-4279-b85b-0967b9a292f1', '3 PJ Pellet Schematic Sink', '{}', 23742);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b3d19718-c3bd-4f5b-96ae-25355e3f9b7d', '3 PJ Pellet Schematic Faucet', '{}', 11850);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1e776623-9980-4359-abbc-b9a916fb8766', 'ff4095cd-9cad-4ad4-9b0e-1cb1836d694e', 1, '{}', '526931d1-680f-4279-b85b-0967b9a292f1');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('8153818d-eb6a-4e5e-ba80-f04e981505ac', 'ff4095cd-9cad-4ad4-9b0e-1cb1836d694e', 1, '{}', 'b3d19718-c3bd-4f5b-96ae-25355e3f9b7d');

-- flippers for cell and pellet
INSERT INTO public.processes (id, name, meta, "time") VALUES ('bf7c7948-e4c0-40b1-adec-a17b3411d336', 'Move 1 GWH Cell', '{}', 7689);
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a5609c08-2fd2-46ca-a8f5-1f3d725751a1', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 10, '{}', 'bf7c7948-e4c0-40b1-adec-a17b3411d336');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('02365ace-0a47-451e-933d-797d360e858b', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 10, '{}', 'bf7c7948-e4c0-40b1-adec-a17b3411d336');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('1b04331c-bffe-4a11-ac2d-e0449d53fbfb', 'Move 3 PJ Pellet', '{}', 6226);
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('84d2763a-6e08-42a4-bfe9-2d3066620ec5', 'c6cfba75-2cb5-47f5-9806-a92c4909c2ef', 7, '{}', '1b04331c-bffe-4a11-ac2d-e0449d53fbfb');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('74c5ed71-7aef-4c8b-ad74-ed92b96e0ced', 'c6cfba75-2cb5-47f5-9806-a92c4909c2ef', 7, '{}', '1b04331c-bffe-4a11-ac2d-e0449d53fbfb');

-- charge 10 kWH from 1 GWH
INSERT INTO public.processes (id, name, meta, "time") VALUES ('6cacaecf-b358-4008-a9ca-66510a64c646', 'Charge 10 kWH from 1 GWH', '{}', 1450);

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a5bb2f26-5be2-42ef-b3bb-f80038a8452e', '3935523b-3e38-485b-935f-e790758ce36b', 1, '{}', '6cacaecf-b358-4008-a9ca-66510a64c646');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('71af0ae2-e1f3-4c3e-83ef-ec3bbfafd8cc', '24800206-2c58-45b0-8238-81974d0ebb3b', 1000000, '{}', '6cacaecf-b358-4008-a9ca-66510a64c646');

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6335d7d3-d3bb-4a9e-ba50-a9c5361e4b5c', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 1, '{}', '6cacaecf-b358-4008-a9ca-66510a64c646');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3ad35763-91cb-42ae-a8af-135e733998a9', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 1000000, '{}', '6cacaecf-b358-4008-a9ca-66510a64c646');

INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('0b646850-e787-4572-adb0-a4dbfa804d85', 'schematic', '1 GWH to 10 kWH Conversion', '{"volume": 1, "industrialmarket": {"maxprice": 134729, "minprice": 62585, "silosize": 100, "process_id": "6cacaecf-b358-4008-a9ca-66510a64c646"}}');

-- charge 1 GWH from 10 kWH
INSERT INTO public.processes (id, name, meta, "time") VALUES ('567aacc9-f86f-42aa-8f66-f0cf647a92f6', 'Charge 1 GWH from 10 kWH', '{}', 4219);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f40f2da9-bb94-48cb-808f-32b48a2a6cbc', '3935523b-3e38-485b-935f-e790758ce36b', 1, '{}', '567aacc9-f86f-42aa-8f66-f0cf647a92f6');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('948042b8-7b00-41c5-a6b1-706c65037c33', '24800206-2c58-45b0-8238-81974d0ebb3b', 1000000, '{}', '567aacc9-f86f-42aa-8f66-f0cf647a92f6');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('6f72c5d3-72f4-43f2-9105-ce34cb8140a5', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 1, '{}', '567aacc9-f86f-42aa-8f66-f0cf647a92f6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('75bb6035-8ccf-4ecc-81ba-0992fac2e6af', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 1000000, '{}', '567aacc9-f86f-42aa-8f66-f0cf647a92f6');

INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('ce047922-0ac5-49da-be61-79b0f7508e16', 'schematic', '10 kWH to 1 GWH Conversion', '{"volume": 1, "industrialmarket": {"maxprice": 217803, "minprice": 120112, "silosize": 100, "process_id": "567aacc9-f86f-42aa-8f66-f0cf647a92f6"}}');

-- charge 10 kWH from 1 GWH | faucet and sink
INSERT INTO public.processes (id, name, meta, "time") VALUES ('c4798ab5-29ba-43ce-b6be-05199f09a4c7', 'Charge 10 kWH from 1 GWH Sink', '{}', 42109);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('8f0b3eb8-20fa-4e7e-a507-c260b8f29e2d', 'Charge 10 kWH from 1 GWH Faucet', '{}', 41789);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1073c5e5-da88-4f21-a820-a3dc8afaaaab', '0b646850-e787-4572-adb0-a4dbfa804d85', 1, '{}', 'c4798ab5-29ba-43ce-b6be-05199f09a4c7');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('0868aa9b-681d-4f89-9ac8-8b704a2c7968', '0b646850-e787-4572-adb0-a4dbfa804d85', 1, '{}', '8f0b3eb8-20fa-4e7e-a507-c260b8f29e2d');

-- charge 1 GWH from 10 kWH | faucet and sink
INSERT INTO public.processes (id, name, meta, "time") VALUES ('df87682f-725e-4a01-a09b-40dd9d81d5b2', 'Charge 1 GWH from 10 kWH Sink', '{}', 31857);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('33e856f1-f9b8-4f18-aa4d-a133d7ea66b2', 'Charge 1 GWH from 10 kWH Faucet', '{}', 32609);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('c22daabb-fbe6-40a9-b5f7-dfc10caba543', 'ce047922-0ac5-49da-be61-79b0f7508e16', 1, '{}', 'df87682f-725e-4a01-a09b-40dd9d81d5b2');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('a86ccad8-6d43-4699-8f59-bd06674a23f0', 'ce047922-0ac5-49da-be61-79b0f7508e16', 1, '{}', '33e856f1-f9b8-4f18-aa4d-a133d7ea66b2');

-- giga patch kit
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('897a3ede-d682-4e1e-bbf9-4605dee60597', 'repair_kit', 'XL Nanite Patch Kit', '{"hp": 1, "volume": 467, "hullconversion": 50000, "armorconversion": 250000, "industrialmarket": {"maxprice": 190000, "minprice": 60000, "silosize": 80000}}');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('3da96cfb-6288-4ac4-8324-cc400a278092', 'Make XL Nanite Patch Kit', '{}', 91200);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('0be9e154-b8d9-4a48-b7f1-4412452351a6', 'Move XL Nanite Patch Kit', '{}', 3278);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f981639f-1150-4aa6-8b17-de8577f4f836', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 1600000, '{}', '3da96cfb-6288-4ac4-8324-cc400a278092');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('208abb4b-29df-4124-974a-adc8dc166faf', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 50000, '{}', '3da96cfb-6288-4ac4-8324-cc400a278092');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cb2bd69f-58c5-4aa1-972d-283c36197efb', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 6000, '{}', '3da96cfb-6288-4ac4-8324-cc400a278092');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('eeb726ad-8b04-492e-8e5d-e7b98c590449', '24800206-2c58-45b0-8238-81974d0ebb3b', 20000, '{}', '3da96cfb-6288-4ac4-8324-cc400a278092');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('41262701-f16d-4585-9eeb-1efbcb2ae35b', '897a3ede-d682-4e1e-bbf9-4605dee60597', 350, '{}', '0be9e154-b8d9-4a48-b7f1-4412452351a6');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('d6ced9a6-52ef-4d6c-9cab-2024b2bbaf71', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 20000, '{}', '3da96cfb-6288-4ac4-8324-cc400a278092');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('16e3a66a-0aba-454b-8596-f6cb1fa59a0a', '897a3ede-d682-4e1e-bbf9-4605dee60597', 350, '{}', '0be9e154-b8d9-4a48-b7f1-4412452351a6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('3578c3ab-d755-4103-a8a1-c9984462f3bb', '897a3ede-d682-4e1e-bbf9-4605dee60597', 75, '{}', '3da96cfb-6288-4ac4-8324-cc400a278092');

-- giga patch kit schematic w/fuacet + sink
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('e3c0f5aa-5f61-49cd-a90c-f01d664898bb', 'schematic', 'XL Nanite Patch Kit Schematic', '{"volume": 1, "industrialmarket": {"maxprice": 8932485, "minprice": 3371992, "silosize": 100, "process_id": "3da96cfb-6288-4ac4-8324-cc400a278092"}}');

INSERT INTO public.processes (id, name, meta, "time") VALUES ('21eb50c8-f3ec-41ea-b686-e381a296d5c8', 'XL Nanite Patch Kit Sink', '{}', 22414);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('597517a0-232f-41bb-9924-6eb55476afe2', 'XL Nanite Patch Kit Faucet', '{}', 16399);

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bc07f3a7-6073-4458-a142-4479b504c137', 'e3c0f5aa-5f61-49cd-a90c-f01d664898bb', 1, '{}', '21eb50c8-f3ec-41ea-b686-e381a296d5c8');

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('48d703b3-190f-4892-86dd-c40c3d2efc9a', 'e3c0f5aa-5f61-49cd-a90c-f01d664898bb', 1, '{}', '597517a0-232f-41bb-9924-6eb55476afe2');
