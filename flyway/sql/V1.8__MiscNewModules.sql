-- item families
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('therm_cap', 'Thermal Capacitor', '{}');
INSERT INTO public.itemfamilies (id, friendlyname, meta) VALUES ('burst_reactor', 'Burst Fusion Reactor', '{}');

-- item types
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('6f402b12-c262-4eb8-b5c8-46c5b58fcef7', 'therm_cap', 'Basic Thermal Capacitor', '{"hp": 21, "rack": "c", "volume": 5, "heat_cap_max_add": 200, "industrialmarket": {"maxprice": 1316, "minprice": 944, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('95fc28b8-9464-474d-a274-82e62d059199', 'therm_cap', 'Small Thermal Capacitor', '{"hp": 44, "rack": "c", "volume": 23, "heat_cap_max_add": 415, "industrialmarket": {"maxprice": 3371, "minprice": 1688, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d43e9d2f-8185-4efd-afba-4fad3f419f76', 'therm_cap', 'XL Thermal Capacitor', '{"hp": 216, "rack": "c", "volume": 217, "heat_cap_max_add": 3372, "industrialmarket": {"maxprice": 413219, "minprice": 288920, "silosize": 1000}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('f5833591-6b7c-4b80-9904-81001d51b888', 'burst_reactor', 'Basic Burst Fusion Reactor', '{"hp": 6, "rack": "b", "volume": 7, "leakage": 0.35, "cooldown": 212, "needs_target": false, "activation_heat": 216, "max_fuel_volume": 1, "industrialmarket": {"maxprice": 4461, "minprice": 3375, "silosize": 1000}, "activation_energy": 45}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('186561d1-41ab-4525-b8b9-534d78248f10', 'burst_reactor', 'XL Burst Fusion Reactor', '{"hp": 26, "rack": "b", "volume": 215, "leakage": 0.16, "cooldown": 86, "needs_target": false, "activation_heat": 5672, "max_fuel_volume": 1000, "industrialmarket": {"maxprice": 697265, "minprice": 491606, "silosize": 1000}, "activation_energy": 1782}');
