-- ====== gauss shell
INSERT INTO public.processes (id, name, meta, "time") VALUES ('b07f1407-0b62-4ae3-a740-7975604735be', 'Make Heavy Gauss Shell', '{}', 145);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('49dca438-75ac-4d5c-8fe9-a56cdb01b86c', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 78, '{}', 'b07f1407-0b62-4ae3-a740-7975604735be');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e9fa24a9-9a93-4ebb-a39a-c454fb7c996c', '24800206-2c58-45b0-8238-81974d0ebb3b', 21, '{}', 'b07f1407-0b62-4ae3-a740-7975604735be');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('42dde738-8342-4dc8-b91d-7297bca6019d', '366756f0-60e4-4dc1-864f-f94b3ce2c6e3', 600, '{}', 'b07f1407-0b62-4ae3-a740-7975604735be');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('d42438f3-5ec4-43a3-a4fa-7539267d6281', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 21, '{}', 'b07f1407-0b62-4ae3-a740-7975604735be');

-- ====== auto belt
INSERT INTO public.processes (id, name, meta, "time") VALUES ('3289de85-e26f-4003-bc1d-01052e7c6a5a', 'Make Auto-23 Belt', '{}', 168);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1c7d820a-ab5d-4b1c-b575-e0ef53c2fd2c', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 20, '{}', '3289de85-e26f-4003-bc1d-01052e7c6a5a');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('dbdd8f63-3ad1-4712-91e5-e352fc775fc3', '24800206-2c58-45b0-8238-81974d0ebb3b', 13, '{}', '3289de85-e26f-4003-bc1d-01052e7c6a5a');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('b7c1ddc4-1c2a-4d5d-a0d9-a35625427ee1', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 13, '{}', '3289de85-e26f-4003-bc1d-01052e7c6a5a');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('8f7ab7a7-1c0c-466e-a9c9-ba93aebe2ab0', 'b8aa4560-39ea-4af5-a7b8-d360e348a81e', 150, '{}', '3289de85-e26f-4003-bc1d-01052e7c6a5a');

-- ====== hgpm
INSERT INTO public.processes (id, name, meta, "time") VALUES ('54933edb-7b78-4d81-b3a9-05b1bb0e905e', 'Make HGPM-15', '{}', 287);

-- inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('fcf3a8d0-e09e-4d92-8b4f-11ef2690a7cb', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 218, '{}', '54933edb-7b78-4d81-b3a9-05b1bb0e905e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('b4ba5665-7e9d-445c-a784-c511d148b46a', 'a8646647-881a-4d24-a22f-f0dce044e6d3', 4, '{}', '54933edb-7b78-4d81-b3a9-05b1bb0e905e');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('61e0817e-fa78-4508-bdf5-93380bd3aa3b', '24800206-2c58-45b0-8238-81974d0ebb3b', 76, '{}', '54933edb-7b78-4d81-b3a9-05b1bb0e905e');

-- outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('403c6fd1-3273-4e01-b89e-d9fd9b565fe8', '368e03a0-bfef-43df-b686-2a5f279549d1', 195, '{}', '54933edb-7b78-4d81-b3a9-05b1bb0e905e');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('85d9c335-ac28-4bc0-828a-2fe23cea9630', '5c1049c4-f631-4066-9f2a-b0798b2c4399', 76, '{}', '54933edb-7b78-4d81-b3a9-05b1bb0e905e');
