-- insert some modules for the custom ship's class
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('dbc59b6b-1989-46c5-9710-933415f9a263', 'shield_booster', 'XL Shield Booster', '{"hp": 500, "rack": "b", "volume": 315, "cooldown": 45, "needs_target": false, "activation_heat": 2080, "industrialmarket": {"maxprice": 12178, "minprice": 43911, "silosize": 1000}, "activation_energy": 480, "shield_boost_amount": 3200, "activation_gfx_effect": "extra_large_shield_booster"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('7f25b73f-8ad8-4e87-bc76-241aaadcdd3b', 'battery_pack', 'XL Battery Pack', '{"hp": 250, "rack": "c", "volume": 210, "energy_max_add": 9920, "industrialmarket": {"maxprice": 302550, "minprice": 215991, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('27d43a76-5822-422b-94b7-65592eda9fec', 'missile_launcher', 'XL LGPM-4 Launcher', '{"hp": 150, "rack": "a", "range": 5505, "volume": 275, "cooldown": 4.25, "flight_time": 11.6, "hull_damage": 809, "armor_damage": 792, "needs_target": true, "shield_damage": 793, "missile_radius": 2.5, "activation_heat": 60, "ammunition_name": "LGPM-4", "ammunition_type": "b9b274c4-d938-4155-8c69-ecb6c4df054f", "fault_tolerance": 0.82, "industrialmarket": {"maxprice": 72355, "minprice": 98199, "silosize": 1000}, "activation_energy": 16, "missile_gfx_effect": "lgpm-4", "missile_explosion_effect": "basic_explosion", "missile_explosion_radius": 5}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('479d2b52-18d9-4677-a193-e3deacd0813d', 'armor_plate', 'XL Armor Plate', '{"hp": 750, "rack": "c", "volume": 260, "armor_max_add": 3428, "industrialmarket": {"maxprice": 65500, "minprice": 37200, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('824b0caa-c9f7-491e-84b4-8a6beb86638e', 'aux_generator', 'XL APU', '{"hp": 80, "rack": "c", "volume": 225, "industrialmarket": {"maxprice": 23856, "minprice": 19681, "silosize": 200}, "energy_regen_max_add": 260}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('f1917bc1-0239-4721-a366-b30151fc5bf2', 'active_sink', 'XL Active Radiator', '{"hp": 120, "rack": "b", "volume": 275, "cooldown": 167.7, "needs_target": false, "activation_heat": 2080, "industrialmarket": {"maxprice": 22725, "minprice": 17157, "silosize": 1000}, "activation_energy": 3375}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('8fef3f9a-4d9b-4baf-8b1c-d0c8ff91e36d', 'utility_cloak', 'XL Cloaking Device', '{"hp": 45, "rack": "a", "volume": 310, "cooldown": 36.3, "activation_heat": 371, "industrialmarket": {"maxprice": 522900, "minprice": 353800, "silosize": 1000}, "activation_energy": 3616}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('74d1fa76-7870-42ac-b402-1e4087e12e4f', 'utility_veil', 'XL Hardening Veil', '{"hp": 25, "rack": "a", "volume": 300, "cooldown": 424, "activation_heat": 169, "industrialmarket": {"maxprice": 364995, "minprice": 459220, "silosize": 1000}, "activation_energy": 3680}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('1376187c-d016-488a-9491-7735b5370ba2', 'eng_oc', 'XL Engine Overcharger', '{"hp": 62, "rack": "b", "volume": 265, "cooldown": 146, "needs_target": false, "activation_heat": 2272, "industrialmarket": {"maxprice": 230456, "minprice": 175962, "silosize": 1000}, "activation_energy": 2816}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('a349ad48-9808-46e1-8f5e-bc4e4da46568', 'utility_siphon', 'XL Energy Siphon', '{"hp": 55, "rack": "a", "range": 15253, "volume": 285, "falloff": "linear", "cooldown": 42, "tracking": 6.9, "needs_target": true, "activation_heat": 712, "industrialmarket": {"maxprice": 92811, "minprice": 35690, "silosize": 1000}, "activation_energy": 160, "energy_siphon_amount": 3562, "activation_gfx_effect": "extra_large_energy_siphon"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('82227782-f5cd-4ba3-bd0e-99e5e17f6ece', 'heat_sink', 'XL Heat Sink', '{"hp": 280, "rack": "c", "volume": 205, "heat_sink_add": 201.2, "industrialmarket": {"maxprice": 25000, "minprice": 18900, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('06ba702a-f071-4b0a-a5ef-cbfa139f23d5', 'missile_launcher', 'XL LHVM-3 Launcher', '{"hp": 57, "rack": "a", "range": 4407, "volume": 355, "cooldown": 1.82, "flight_time": 6.2, "hull_damage": 428, "armor_damage": 354, "needs_target": true, "shield_damage": 385, "missile_radius": 2.5, "activation_heat": 291.2, "ammunition_name": "LHVM-3", "ammunition_type": "8aa4d28b-e02b-4630-86eb-ab46d0ed7fc7", "fault_tolerance": 0.75, "industrialmarket": {"maxprice": 833902, "minprice": 492881, "silosize": 1000}, "activation_energy": 0.06, "missile_gfx_effect": "lhvm-3", "missile_explosion_effect": "basic_explosion", "missile_explosion_radius": 1.92}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('c0b46a6f-a56d-4ecb-9c17-849bf548fb72', 'gun_turret', 'XL Anti-Shield Laser', '{"hp": 30, "rack": "a", "range": 51296, "volume": 277, "falloff": "linear", "cooldown": 117, "tracking": 7.31, "hull_damage": 272, "armor_damage": 118, "needs_target": true, "shield_damage": 44792, "activation_heat": 7230, "industrialmarket": {"maxprice": 653469, "minprice": 330012, "silosize": 1000}, "activation_energy": 3729, "activation_gfx_effect": "extra_large_shield_laser"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('a17d6f86-1689-45b8-989c-033caabf4cb2', 'gun_turret', 'XL Anti-Hull Laser', '{"hp": 70, "rack": "a", "range": 51329, "volume": 233, "falloff": "linear", "cooldown": 112, "tracking": 7.53, "hull_damage": 20000, "armor_damage": 15200, "needs_target": true, "shield_damage": 95, "activation_heat": 2759, "industrialmarket": {"maxprice": 583165, "minprice": 327891, "silosize": 1000}, "activation_energy": 7826, "activation_gfx_effect": "extra_large_hull_laser"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('39a3a402-7a6f-42c8-8cb3-33e4b4c16ae3', 'gun_turret', 'XL Combat Laser', '{"hp": 85, "rack": "a", "range": 41955, "volume": 322, "falloff": "linear", "cooldown": 123.7, "tracking": 5.9, "hull_damage": 12672, "armor_damage": 11044, "needs_target": true, "shield_damage": 11987, "activation_heat": 6555, "industrialmarket": {"maxprice": 770937, "minprice": 519351, "silosize": 1000}, "activation_energy": 3520, "activation_gfx_effect": "extra_large_general_laser"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('cefa9c66-8c82-4b12-9229-8766099331fd', 'gun_turret', 'XL Ice Miner', '{"hp": 92, "rack": "a", "range": 9955, "volume": 295, "falloff": "reverse_linear", "cooldown": 122, "tracking": 0.084, "hull_damage": 955, "armor_damage": 477, "can_mine_gas": false, "can_mine_ice": true, "can_mine_ore": false, "needs_target": true, "shield_damage": 160, "activation_heat": 4000, "industrialmarket": {"maxprice": 120728, "minprice": 75092, "silosize": 1000}, "activation_energy": 488, "ice_mining_volume": 240, "activation_gfx_effect": "extra_large_ice_miner"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('63ea40d0-5b2c-497c-8115-04a012c58a15', 'utility_miner', 'XL Ore Harvester', '{"hp": 76, "rack": "a", "range": 7345, "volume": 317, "falloff": "reverse_linear", "cooldown": 189.5, "tracking": 0.059, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "activation_heat": 4068, "industrialmarket": {"maxprice": 410021, "minprice": 119963, "silosize": 1000}, "activation_energy": 6780, "ore_mining_volume": 1310, "activation_gfx_effect": "extra_large_ore_harvester"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('c83c205a-5393-4783-b764-3d20f45b2f0e', 'utility_miner', 'XL Ice Harvester', '{"hp": 62, "rack": "a", "range": 7689, "volume": 328, "falloff": "reverse_linear", "cooldown": 199, "tracking": 0.057, "can_mine_gas": false, "can_mine_ice": true, "can_mine_ore": false, "needs_target": true, "activation_heat": 3960, "industrialmarket": {"maxprice": 682139, "minprice": 450918, "silosize": 1000}, "activation_energy": 4825, "ice_mining_volume": 1186, "activation_gfx_effect": "extra_large_ice_harvester"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('5d6dc434-0df9-493a-9a25-5662a17402c0', 'gun_turret', 'XL Laser Tool', '{"hp": 430, "rack": "a", "range": 14065, "volume": 250, "falloff": "linear", "cooldown": 25, "tracking": 0.2625, "hull_damage": 3840, "armor_damage": 1250, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2912, "activation_heat": 960, "industrialmarket": {"maxprice": 77915, "minprice": 62303, "silosize": 1000}, "activation_energy": 800, "ore_mining_volume": 560, "activation_gfx_effect": "extra_large_laser_tool"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('76927c05-d018-4625-a963-f04c426854e3', 'drag_amp', 'XL Aether Dragger', '{"hp": 29, "rack": "b", "range": 32497, "volume": 268, "falloff": "linear", "cooldown": 329, "needs_target": true, "activation_heat": 2752, "drag_multiplier": 457, "industrialmarket": {"maxprice": 961996, "minprice": 774234, "silosize": 500}, "activation_energy": 6608, "activation_gfx_effect": "extra_large_aether_dragger"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('8e4d6054-76e9-4acb-92c0-17009ffad26a', 'gun_turret', 'Sniping Gauss Rifle', '{"hp": 35, "rack": "a", "range": 119773, "volume": 233, "falloff": "linear", "cooldown": 21.25, "tracking": 25.8, "hull_damage": 313, "armor_damage": 231, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": false, "needs_target": true, "shield_damage": 156, "activation_heat": 907, "ammunition_name": "Light Gauss Shell", "ammunition_type": "33f86ee2-4c16-400e-82ad-38796bea0017", "industrialmarket": {"maxprice": 281735, "minprice": 119816, "silosize": 1000}, "activation_energy": 512, "ore_mining_volume": 0, "activation_gfx_effect": "basic_gauss_rifle"}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('ede9ac0a-60ec-48d0-8ed6-ad715d763eb5', 'gun_turret', 'Auto-5 Point Defense', '{"hp": 33, "rack": "a", "range": 3045, "volume": 245, "falloff": "linear", "cooldown": 3.2, "tracking": 562.7, "hull_damage": 455, "armor_damage": 305, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": false, "needs_target": true, "shield_damage": 78, "activation_heat": 2.1, "ammunition_name": "Auto-5 Belt", "ammunition_type": "812bddab-db8f-4444-a787-e39e59b92256", "industrialmarket": {"maxprice": 104951, "minprice": 870918, "silosize": 1000}, "activation_energy": 12.8, "ore_mining_volume": 0, "activation_gfx_effect": "basic_auto-5_cannon"}');

-- insert custom ship
INSERT INTO public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume, itemtypeid, canundock, wrecktexture, explosiontexture) VALUES ('e2a67722-10d1-42fd-9a7c-057677ef6e79', '2022-04-10 00:00:00-04', 'Chollima', 'Chollima', 137.8, 3.6, 286720, 0.37, 21120, 6280, 15625, 93875, 91840, 22752, 374.4, 5184, 403.2, 'a7b8e2cf-9e69-480e-a5fa-dc19d8be9a57', '{"a_slots": [{"hp_pos": [80, 0], "volume": 450, "mod_family": "missile"}, {"hp_pos": [60, 0], "volume": 450, "mod_family": "missile"}, {"hp_pos": [40, 0], "volume": 450, "mod_family": "missile"}, {"hp_pos": [20, 0], "volume": 450, "mod_family": "missle"}, {"hp_pos": [-20, 0], "volume": 310, "mod_family": "gun"}, {"hp_pos": [-40, 0], "volume": 310, "mod_family": "gun"}, {"hp_pos": [-60, 0], "volume": 245, "mod_family": "utility"}, {"hp_pos": [-80, 0], "volume": 245, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 400, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0], "volume": 275, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 275, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 275, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 275, "mod_family": "any"}, {"hp_pos": [0, 0], "volume": 275, "mod_family": "any"}]}', 17600, 'c5b01960-fef0-47ae-9c16-25aa1799a003', true, 'basic-wreck', 'basic_explosion');

-- insert custom start
INSERT INTO public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid, wallet, factionid) VALUES ('9a994832-8037-4e22-b4fc-e0d4f081b456', 'Chrissy''s Start', 'e2a67722-10d1-42fd-9a7c-057677ef6e79', '{"a_rack": [{"item_type_id": "27d43a76-5822-422b-94b7-65592eda9fec"}, {"item_type_id": "27d43a76-5822-422b-94b7-65592eda9fec"}, {"item_type_id": "06ba702a-f071-4b0a-a5ef-cbfa139f23d5"}, {"item_type_id": "06ba702a-f071-4b0a-a5ef-cbfa139f23d5"}, {"item_type_id": "cefa9c66-8c82-4b12-9229-8766099331fd"}, {"item_type_id": "5d6dc434-0df9-493a-9a25-5662a17402c0"}, {}, {}], "b_rack": [{"item_type_id": "dbc59b6b-1989-46c5-9710-933415f9a263"}, {"item_type_id": "dbc59b6b-1989-46c5-9710-933415f9a263"}, {"item_type_id": "76927c05-d018-4625-a963-f04c426854e3"}, {"item_type_id": "76927c05-d018-4625-a963-f04c426854e3"}, {}, {}], "c_rack": [{"item_type_id": "7f25b73f-8ad8-4e87-bc76-241aaadcdd3b"}, {"item_type_id": "7f25b73f-8ad8-4e87-bc76-241aaadcdd3b"}, {"item_type_id": "82227782-f5cd-4ba3-bd0e-99e5e17f6ece"}, {"item_type_id": "82227782-f5cd-4ba3-bd0e-99e5e17f6ece"}, {"item_type_id": "b79c5b09-cb85-4c38-a595-95cddaf8bff5"}]}', '2022-05-15 16:01:53.877839-04', false, '0078c73f-1b40-4b22-bdf0-f4927d8bd7bd', '306c0fcd-cf01-43fc-850f-172191fe1582', 250000, '27a53dfc-a321-4c12-bf7c-bb177955c95b');

-- update chrissy's user to use custom start
UPDATE public.users SET StartID = '9a994832-8037-4e22-b4fc-e0d4f081b456' where ID = 'c98932b5-3c81-49ef-b795-b3ca88b1bf95';
