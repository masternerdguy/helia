-- insert some modules for the custom ship's class
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('dbc59b6b-1989-46c5-9710-933415f9a263', 'shield_booster', 'Extra Large Shield Booster', '{"hp": 500, "rack": "b", "volume": 215, "cooldown": 45, "needs_target": false, "activation_heat": 2080, "industrialmarket": {"maxprice": 12178, "minprice": 43911, "silosize": 1000}, "activation_energy": 480, "shield_boost_amount": 3200, "activation_gfx_effect": "basic_shield_booster"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('7f25b73f-8ad8-4e87-bc76-241aaadcdd3b', 'battery_pack', 'Extra Large Battery Pack', '{"hp": 250, "rack": "c", "volume": 110, "energy_max_add": 9920, "industrialmarket": {"maxprice": 302550, "minprice": 215991, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('5d6dc434-0df9-493a-9a25-5662a17402c0', 'gun_turret', 'Extra Large Laser Tool', '{"hp": 430, "rack": "a", "range": 42195, "volume": 213, "falloff": "linear", "cooldown": 25, "tracking": 0.2625, "hull_damage": 3840, "armor_damage": 1250, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2912, "activation_heat": 960, "industrialmarket": {"maxprice": 77915, "minprice": 62303, "silosize": 1000}, "activation_energy": 800, "ore_mining_volume": 560, "activation_gfx_effect": "extralarge_laser_tool"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('27d43a76-5822-422b-94b7-65592eda9fec', 'missile_launcher', 'Extra Large LGPM-4 Launcher', '{"hp": 150, "rack": "a", "range": 5505, "volume": 275, "cooldown": 4.25, "flight_time": 11.6, "hull_damage": 809, "armor_damage": 792, "needs_target": true, "shield_damage": 793, "missile_radius": 2.5, "activation_heat": 60, "ammunition_name": "LGPM-4", "ammunition_type": "b9b274c4-d938-4155-8c69-ecb6c4df054f", "fault_tolerance": 0.82, "industrialmarket": {"maxprice": 72355, "minprice": 98199, "silosize": 1000}, "activation_energy": 16, "missile_gfx_effect": "lgpm-4", "missile_explosion_effect": "basic_explosion", "missile_explosion_radius": 5}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('479d2b52-18d9-4677-a193-e3deacd0813d', 'armor_plate', 'Extra Large Armor Plate', '{"hp": 750, "rack": "c", "volume": 160, "armor_max_add": 3428, "industrialmarket": {"maxprice": 65500, "minprice": 37200, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('824b0caa-c9f7-491e-84b4-8a6beb86638e', 'aux_generator', 'Extra Large APU', '{"hp": 80, "rack": "c", "volume": 107, "industrialmarket": {"maxprice": 23856, "minprice": 19681, "silosize": 200}, "energy_regen_max_add": 260}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('f1917bc1-0239-4721-a366-b30151fc5bf2', 'active_sink', 'Extra Large Active Radiator', '{"hp": 120, "rack": "b", "volume": 160, "cooldown": 167.7, "needs_target": false, "activation_heat": 2080, "industrialmarket": {"maxprice": 22725, "minprice": 17157, "silosize": 1000}, "activation_energy": 3375}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('8fef3f9a-4d9b-4baf-8b1c-d0c8ff91e36d', 'utility_cloak', 'Extra Large Cloaking Device', '{"hp": 45, "rack": "a", "volume": 150, "cooldown": 36.3, "activation_heat": 371, "industrialmarket": {"maxprice": 522900, "minprice": 353800, "silosize": 1000}, "activation_energy": 3616}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('74d1fa76-7870-42ac-b402-1e4087e12e4f', 'utility_veil', 'Extra Large Hardening Veil', '{"hp": 25, "rack": "a", "volume": 400, "cooldown": 424, "activation_heat": 169, "industrialmarket": {"maxprice": 364995, "minprice": 459220, "silosize": 1000}, "activation_energy": 3680}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('1376187c-d016-488a-9491-7735b5370ba2', 'eng_oc', 'Extra Large Engine Overcharger', '{"hp": 62, "rack": "b", "volume": 275, "cooldown": 146, "needs_target": false, "activation_heat": 2272, "industrialmarket": {"maxprice": 230456, "minprice": 175962, "silosize": 1000}, "activation_energy": 2816}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('a349ad48-9808-46e1-8f5e-bc4e4da46568', 'utility_siphon', 'Extra Large Energy Siphon', '{"hp": 55, "rack": "a", "range": 15253, "volume": 255, "falloff": "linear", "cooldown": 42, "tracking": 6.9, "needs_target": true, "activation_heat": 712, "industrialmarket": {"maxprice": 92811, "minprice": 35690, "silosize": 1000}, "activation_energy": 160, "energy_siphon_amount": 3562, "activation_gfx_effect": "extra_large_energy_siphon"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('82227782-f5cd-4ba3-bd0e-99e5e17f6ece', 'heat_sink', 'Extra Large Heat Sink', '{"hp": 280, "rack": "c", "volume": 128, "heat_sink_add": 201.2, "industrialmarket": {"maxprice": 25000, "minprice": 18900, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('06ba702a-f071-4b0a-a5ef-cbfa139f23d5', 'missile_launcher', 'Extra Large LHVM-3 Launcher', '{"hp": 57, "rack": "a", "range": 4407, "volume": 385, "cooldown": 1.82, "flight_time": 6.2, "hull_damage": 428, "armor_damage": 354, "needs_target": true, "shield_damage": 385, "missile_radius": 2.5, "activation_heat": 291.2, "ammunition_name": "LHVM-3", "ammunition_type": "8aa4d28b-e02b-4630-86eb-ab46d0ed7fc7", "fault_tolerance": 0.75, "industrialmarket": {"maxprice": 833902, "minprice": 492881, "silosize": 1000}, "activation_energy": 0.06, "missile_gfx_effect": "lhvm-3", "missile_explosion_effect": "basic_explosion", "missile_explosion_radius": 1.92}');

-- insert custom ship
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('e2a67722-10d1-42fd-9a7c-057677ef6e79', '2022-04-10 00:00:00-04', 'Chollima', 'Chollima', 137.8, 3.6, 286720, 0.85, 21120, 6280, 15625, 93875, 91840, 22752, 374.4, 5184, 403.2, 'a7b8e2cf-9e69-480e-a5fa-dc19d8be9a57', '{"a_slots": [{"hp_pos": [4, 30], "volume": 350, "mod_family": "missile"}, {"hp_pos": [4, -30], "volume": 350, "mod_family": "missile"}, {"hp_pos": [2, 30], "volume": 350, "mod_family": "missile"}, {"hp_pos": [2, -30], "volume": 350, "mod_family": "missle"}, {"hp_pos": [4, 30], "volume": 210, "mod_family": "gun"}, {"hp_pos": [4, -30], "volume": 210, "mod_family": "gun"}, {"hp_pos": [2, 30], "volume": 245, "mod_family": "utility"}, {"hp_pos": [2, -30], "volume": 245, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 175, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 175, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 175, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 175, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 175, "mod_family": "any"}]}', 17600, 'c5b01960-fef0-47ae-9c16-25aa1799a003', true, 'basic-wreck', 'basic_explosion');

-- insert custom start
INSERT INTO public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid, wallet, factionid) VALUES ('9a994832-8037-4e22-b4fc-e0d4f081b456', 'Chrissy''s Start', 'e2a67722-10d1-42fd-9a7c-057677ef6e79', '{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9bb00839-95f2-4d7c-a7c9-eef60e05fa97"}, {}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}', '2022-05-15 16:01:53.877839-04', false, '0078c73f-1b40-4b22-bdf0-f4927d8bd7bd', '306c0fcd-cf01-43fc-850f-172191fe1582', 250000, '27a53dfc-a321-4c12-bf7c-bb177955c95b');

-- update chrissy's user to use custom start
UPDATE public.users SET StartID = '9a994832-8037-4e22-b4fc-e0d4f081b456' where ID = 'c98932b5-3c81-49ef-b795-b3ca88b1bf95';
