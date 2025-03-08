-- Heavy Gauss Shell
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('366756f0-60e4-4dc1-864f-f94b3ce2c6e3', 'ammunition', 'Heavy Gauss Shell', '{"hp": 1, "volume": 1.5, "industrialmarket": {"maxprice": 135, "minprice": 80, "silosize": 76000}}');

-- XL Gauss Rifle
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('1c83a2c1-1723-4b4d-ad06-626fef52c34b', 'gun_turret', 'XL Gauss Rifle', '{"hp": 89, "rack": "a", "range": 56249, "volume": 315, "falloff": "linear", "cooldown": 78.4, "tracking": 9.7, "hull_damage": 12126, "armor_damage": 8933, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": false, "needs_target": true, "shield_damage": 5871, "activation_heat": 1179.4, "ammunition_name": "Heavy Gauss Shell", "ammunition_type": "366756f0-60e4-4dc1-864f-f94b3ce2c6e3", "industrialmarket": {"maxprice": 198770, "minprice": 96568, "silosize": 250}, "activation_energy": 465, "ore_mining_volume": 0, "activation_gfx_effect": "xl_gauss_rifle"}');

-- Auto-23 Belt
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('b8aa4560-39ea-4af5-a7b8-d360e348a81e', 'ammunition', 'Auto-23 Belt', '{"hp": 1, "volume": 1.3, "industrialmarket": {"maxprice": 73, "minprice": 26, "silosize": 35000}}');

-- Basic Auto-23 Cannon
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('4e32b02a-0caf-4f32-9390-0fdd011a7759', 'gun_turret', 'Basic Auto-23 Cannon', '{"hp": 135, "rack": "a", "range": 3135, "volume": 275, "falloff": "linear", "cooldown": 1.9, "tracking": 62, "hull_damage": 3548, "armor_damage": 2123, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": false, "needs_target": true, "shield_damage": 606, "activation_heat": 25.3, "ammunition_name": "Auto-23 Belt", "ammunition_type": "b8aa4560-39ea-4af5-a7b8-d360e348a81e", "industrialmarket": {"maxprice": 193845, "minprice": 75123, "silosize": 500}, "activation_energy": 6.7, "ore_mining_volume": 0, "activation_gfx_effect": "heavy_auto-23_cannon"}');

-- Dual Auto-23 Cannon :)
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('20f3cdc7-00cf-4409-a2b3-670786a0cbba', 'gun_turret', 'Dual Auto-23 Cannon', '{"hp": 185, "rack": "a", "range": 3097, "volume": 420, "falloff": "linear", "cooldown": 0.98, "tracking": 31, "hull_damage": 3539, "armor_damage": 2099, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": false, "needs_target": true, "shield_damage": 598, "activation_heat": 26.8, "ammunition_name": "Auto-23 Belt", "ammunition_type": "b8aa4560-39ea-4af5-a7b8-d360e348a81e", "industrialmarket": {"maxprice": 392788, "minprice": 154377, "silosize": 250}, "activation_energy": 7.1, "ore_mining_volume": 0, "activation_gfx_effect": "heavy_auto-23_cannon"}');

-- HGPM-15
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('368e03a0-bfef-43df-b686-2a5f279549d1', 'ammunition', 'HGPM-15', '{"hp": 1, "volume": 13, "industrialmarket": {"maxprice": 675, "minprice": 230, "silosize": 62000}}');

-- Basic HGPM-15 Launcher
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('038fc30e-bb2b-4b0f-aa37-c07e140ff6ed', 'missile_launcher', 'Basic HGPM-15 Launcher', '{"hp": 97, "rack": "a", "range": 40254, "volume": 365, "cooldown": 31.5, "flight_time": 87.1, "hull_damage": 7116, "armor_damage": 7290, "needs_target": true, "shield_damage": 7615, "missile_radius": 5.5, "activation_heat": 252, "ammunition_name": "HGPM-15", "ammunition_type": "368e03a0-bfef-43df-b686-2a5f279549d1", "fault_tolerance": 0.98, "industrialmarket": {"maxprice": 52272, "minprice": 21300, "silosize": 500}, "activation_energy": 3.8, "missile_gfx_effect": "hgpm-15", "missile_explosion_effect": "basic_explosion", "missile_explosion_radius": 56}');
