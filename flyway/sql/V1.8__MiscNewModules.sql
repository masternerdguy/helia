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
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('0b65f180-4c42-419d-afa3-7f9143fb9eaa', 'schematic', 'Basic Thermal Capacitor Schematic', '{"industrialmarket": {"maxprice": 35325, "minprice": 12084, "silosize": 100, "process_id": "9bbf6f88-2dfb-4b27-a024-5176ad18e98b"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('b101b6bc-5ce6-4c7d-ba10-8756b43e1422', 'schematic', 'Small Thermal Capacitor Schematic', '{"industrialmarket": {"maxprice": 43796, "minprice": 19453, "silosize": 100, "process_id": "faf7d498-10e9-470c-adc7-e96b7de251cb"}}');

INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('65a84cbe-2c47-448f-8d4d-0198180bc723', 'schematic', 'Basic Burst Fusion Reactor Schematic', '{"industrialmarket": {"maxprice": 91953, "minprice": 15472, "silosize": 100, "process_id": "0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('0dc7e9a2-6660-41da-9ecd-316f5ec21986', 'schematic', 'XL Thermal Capacitor Schematic', '{"industrialmarket": {"maxprice": 7144758, "minprice": 2812287, "silosize": 100, "process_id": "966dd5f0-665b-440f-a2c5-2a005e87546b"}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('a2f27e00-1c0c-4e22-be80-917d4d4dfee3', 'schematic', 'XL Burst Fusion Reactor Schematic', '{"industrialmarket": {"maxprice": 13503021, "minprice": 4256913, "silosize": 100, "process_id": "7551b393-1268-4cc5-bcf0-b57c90ce0eca"}}');

-- processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9bbf6f88-2dfb-4b27-a024-5176ad18e98b', 'Make Basic Thermal Capacitor', '{}', 486);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('d416b6c2-d171-4393-a1b8-0043d1d94780', 'Basic Thermal Capacitor Sink [wm]', '{}', 513);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b75e1bfb-5269-4d90-90d1-413d7b0c7b2f', 'Basic Thermal Capacitor Schematic Faucet [wm]', '{}', 1179);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('76c24cf6-2e23-4618-b8ee-f329eb95630f', 'Basic Thermal Capacitor Schematic Sink [wm]', '{}', 3766);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('faf7d498-10e9-470c-adc7-e96b7de251cb', 'Make Small Thermal Capacitor', '{}', 1901);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('5541fe0a-1f99-4a55-8352-c8fc2e2c7625', 'Small Thermal Capacitor Sink [wm]', '{}', 2917);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('8b3eb14d-a97a-4da3-91a5-5fe8bcce00d4', 'Small Thermal Capacitor Schematic Faucet [wm]', '{}', 12232);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('4654ddac-3eaa-4c67-8381-4b3c741a1f03', 'Small Thermal Capacitor Schematic Sink [wm]', '{}', 13862);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc', 'Make Basic Burst Fusion Reactor', '{}', 201);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('57cb556b-9a73-4866-8c7e-f214149afbba', 'Basic Burst Fusion Reactor Sink [wm]', '{}', 308);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('721feee4-5af0-47a2-ac6b-ba045e001be4', 'Basic Burst Fusion Reactor Schematic Faucet [wm]', '{}', 1381);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('2f31d378-3de3-4608-821e-be19aa563940', 'Basic Burst Fusion Reactor Schematic Sink [wm]', '{}', 1346);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('966dd5f0-665b-440f-a2c5-2a005e87546b', 'Make XL Thermal Capacitor', '{}', 7978);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('7e27aff7-779a-48c9-9817-5ecc56786d5d', 'XL Thermal Capacitor Sink [wm]', '{}', 13522);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('8be2a46f-2393-4064-9755-8d47dcf8b5d9', 'XL Thermal Capacitor Schematic Faucet [wm]', '{}', 18165);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('41c95c18-c52e-4b96-8a35-6183a5bee43b', 'XL Thermal Capacitor Schematic Sink [wm]', '{}', 48862);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('7551b393-1268-4cc5-bcf0-b57c90ce0eca', 'Make XL Burst Fusion Reactor', '{}', 11712);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('264f573d-6019-49c9-a317-6b07e9f2f61c', 'XL Burst Fusion Reactor Sink [wm]', '{}', 17062);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('bed671c1-c147-4457-9400-72d1cae15176', 'XL Burst Fusion Reactor Schematic Faucet [wm]', '{}', 25900);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b23a0dbe-e251-4a0f-8755-0d203ef681c0', 'XL Burst Fusion Reactor Schematic Sink [wm]', '{}', 82774);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b7adb276-be4d-4f9a-87ee-7db78a737b84', '0b65f180-4c42-419d-afa3-7f9143fb9eaa', 3, '{}', '76c24cf6-2e23-4618-b8ee-f329eb95630f');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('404fbb03-63ac-4463-ab96-6c218bfd0e96', '6f402b12-c262-4eb8-b5c8-46c5b58fcef7', 83, '{}', 'd416b6c2-d171-4393-a1b8-0043d1d94780');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4d06e552-9bc3-4c62-bd65-256c66f4ac4d', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 3, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('719bf528-480b-4bf3-a5ec-10cbeef338d0', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 12, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ad9825e4-9ade-4892-be95-76582ce768f2', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 1, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d5a018f4-db59-492a-8e75-e39f8c8f0762', '66b7a322-8cfc-4467-9410-492e6b58f159', 4, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('58086376-1d97-49e0-b6f9-5e61c9cf1c10', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 1, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4da9d239-8247-45e7-91f6-3180c5c2e864', '11688112-f3d4-4d30-864a-684a8b96ea23', 2, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('4df2dbf6-7417-4a97-9939-c0d6876c01e8', '2ce48bef-f06b-4550-b20c-0e64864db051', 3, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('fe14c2eb-a2ad-4bed-a355-ad90d539c6ec', '56617d30-6c30-425c-84bf-2484ae8c1156', 1, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ee97e741-26ae-4b66-8271-1836afa6951b', 'b101b6bc-5ce6-4c7d-ba10-8756b43e1422', 3, '{}', '4654ddac-3eaa-4c67-8381-4b3c741a1f03');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('af5a4a0b-2d9e-44aa-8017-88646fa832cc', '95fc28b8-9464-474d-a274-82e62d059199', 48, '{}', '5541fe0a-1f99-4a55-8352-c8fc2e2c7625');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9e2390e4-2c0f-44f8-afef-f1a5b7c0abb7', '66b7a322-8cfc-4467-9410-492e6b58f159', 3, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b40dc42b-8015-4662-8b9d-3b9374f5125b', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 12, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('ead82a6a-e8c0-4c16-936c-c2d26df39765', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 7, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1e68ddc9-5b69-40dd-9ee3-47ab40189c73', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 39, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f233e9f1-8ef7-4204-b617-a7ab97651557', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 1, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('015d8395-92c9-437f-9e6b-22203641a3a8', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 13, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('50cc680b-a282-4b4d-94b5-f61b6ca01e39', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 9, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('6d8c3366-47fd-4202-a191-6e56198f06b2', '2ce48bef-f06b-4550-b20c-0e64864db051', 10, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('9fa4c5e2-5391-4a7d-90e3-0ffc19d0ad98', '65a84cbe-2c47-448f-8d4d-0198180bc723', 6, '{}', '2f31d378-3de3-4608-821e-be19aa563940');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('cb9fefb8-c07b-4982-baa3-8abdad86bd76', 'f5833591-6b7c-4b80-9904-81001d51b888', 111, '{}', '57cb556b-9a73-4866-8c7e-f214149afbba');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f430ed65-b1f8-40c1-b5d1-c14b12faae04', 'df13b7cf-3019-4f8e-8933-c9f25d4ff941', 67, '{}', '0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('45ad5c2b-c839-4816-91a1-fd8bdb417213', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 293, '{}', '0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('13f9a5e3-61e0-4a52-986e-ee5035dd94d4', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 1, '{}', '0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0f09e68e-d770-4ebf-b095-e2e37a41357d', '61f52ba3-654b-45cf-88e3-33399d12350d', 58, '{}', '0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a86247d7-0d13-4c5d-9c75-af3660deca5e', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 68, '{}', '0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a49d2d5d-4495-4425-ba33-fa09eea43c80', '0dc7e9a2-6660-41da-9ecd-316f5ec21986', 7, '{}', '41c95c18-c52e-4b96-8a35-6183a5bee43b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('2e649181-a53b-4438-aaf4-62c9d294f34e', 'd43e9d2f-8185-4efd-afba-4fad3f419f76', 50, '{}', '7e27aff7-779a-48c9-9817-5ecc56786d5d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8861cf65-b31f-42a7-a8a1-ee7d15c60984', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 1, '{}', '966dd5f0-665b-440f-a2c5-2a005e87546b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f6a81015-b747-484e-a77c-afea84c862bd', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 6672, '{}', '966dd5f0-665b-440f-a2c5-2a005e87546b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8b9cce22-c0cd-486f-8f24-8faf197408b9', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 6738, '{}', '966dd5f0-665b-440f-a2c5-2a005e87546b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('119470de-7db1-4dc8-9dd1-253e0831acd3', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 6747, '{}', '966dd5f0-665b-440f-a2c5-2a005e87546b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('91f991cd-3c37-4f98-93dd-607d321474e3', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 27582, '{}', '966dd5f0-665b-440f-a2c5-2a005e87546b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3d196537-d932-40b7-9e78-577cdc1ed739', 'a2f27e00-1c0c-4e22-be80-917d4d4dfee3', 3, '{}', 'b23a0dbe-e251-4a0f-8755-0d203ef681c0');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a49b6030-4d37-4f0d-9217-31bed3957e4d', '186561d1-41ab-4525-b8b9-534d78248f10', 30, '{}', '264f573d-6019-49c9-a317-6b07e9f2f61c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('66350a38-7329-4f93-921d-f3c4db27134f', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 1363, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('61b4021f-9729-4d59-b1c7-3aca2bcab237', '2ce48bef-f06b-4550-b20c-0e64864db051', 1428, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b5d88ecc-87c1-4262-9157-d90b828b2a4a', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 5774, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('df9ddc56-bfc0-414d-84d8-d8dbb305faa5', '54c29ce8-67c7-46c7-8c02-c737eed3143c', 1, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('3abdc773-d261-4059-bb62-b23dc9bbf6fe', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 1390, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('5907d1fc-1037-4c68-bb5e-d840cf2efcb4', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 1424, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e2da688e-f41b-4201-bdcf-e18da3e23ee1', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 1411, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('d9cb68da-842d-4f01-b904-aa1bee9955ae', '0b65f180-4c42-419d-afa3-7f9143fb9eaa', 7, '{}', 'b75e1bfb-5269-4d90-90d1-413d7b0c7b2f');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('f658897a-e442-40e2-9a37-6fca0cda1ab3', '3935523b-3e38-485b-935f-e790758ce36b', 1, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('3c94f4b7-4215-406b-b122-7fea07274f3d', '6f402b12-c262-4eb8-b5c8-46c5b58fcef7', 70, '{}', '9bbf6f88-2dfb-4b27-a024-5176ad18e98b');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('8675c776-c903-41b1-a98b-fca6facc2e15', 'b101b6bc-5ce6-4c7d-ba10-8756b43e1422', 4, '{}', '8b3eb14d-a97a-4da3-91a5-5fe8bcce00d4');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('2eeb9f83-041e-4895-8ee2-0dd7da88c00c', '3935523b-3e38-485b-935f-e790758ce36b', 1, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('db13e6f5-7d2f-4a59-b2c9-139995c93a91', '95fc28b8-9464-474d-a274-82e62d059199', 60, '{}', 'faf7d498-10e9-470c-adc7-e96b7de251cb');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('cbb9a21f-9c80-4716-8486-45f866f62bcd', '65a84cbe-2c47-448f-8d4d-0198180bc723', 2, '{}', '721feee4-5af0-47a2-ac6b-ba045e001be4');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('c336de9e-21ef-4859-952e-aeef35f8032c', 'f5833591-6b7c-4b80-9904-81001d51b888', 80, '{}', '0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('45c08a10-281a-4ff2-a67f-29c36d491d9d', '3935523b-3e38-485b-935f-e790758ce36b', 1, '{}', '0ee8e0f7-3dd5-4279-b7dc-0ca8df8508fc');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('b8c33a1a-1730-4aed-a973-5ef1b8d2e752', '0dc7e9a2-6660-41da-9ecd-316f5ec21986', 2, '{}', '8be2a46f-2393-4064-9755-8d47dcf8b5d9');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('d7583b1d-f007-4ad2-b9fc-0f1afa647a9e', '3935523b-3e38-485b-935f-e790758ce36b', 1, '{}', '966dd5f0-665b-440f-a2c5-2a005e87546b');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('49d8a948-369c-433a-bf18-673cd7548457', 'd43e9d2f-8185-4efd-afba-4fad3f419f76', 40, '{}', '966dd5f0-665b-440f-a2c5-2a005e87546b');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('b8719880-34c6-4379-b938-930708f5d63c', 'a2f27e00-1c0c-4e22-be80-917d4d4dfee3', 5, '{}', 'bed671c1-c147-4457-9400-72d1cae15176');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('71b891cc-f2bc-4e2b-9c6f-bb06ecf6ee3b', '3935523b-3e38-485b-935f-e790758ce36b', 1, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('0ce78727-5d38-4602-a73b-977f5a87468d', '186561d1-41ab-4525-b8b9-534d78248f10', 40, '{}', '7551b393-1268-4cc5-bcf0-b57c90ce0eca');
