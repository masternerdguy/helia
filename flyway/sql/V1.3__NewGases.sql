-- item family
INSERT INTO public.itemfamilies(
	id, friendlyname, meta)
	VALUES ('gas', 'Trace Wisp', '{}');

-- item types
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('22fb3cda-c949-41ae-bcf5-ee0a60d497fc', 'gas', 'Tektum', '{"hp": 1, "volume": 1, "industrialmarket": {"maxprice": 181, "minprice": 181, "silosize": 446061900}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('5f84addf-d05b-407a-8dcd-fecd3af4c69d', 'gas', 'Avvon', '{"hp": 1, "volume": 2, "industrialmarket": {"maxprice": 5382, "minprice": 414, "silosize": 222924600}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('42179f25-b473-46d1-9592-08e1d23807bd', 'gas', 'Luche', '{"hp": 1, "volume": 3, "industrialmarket": {"maxprice": 12720, "minprice": 636, "silosize": 360182100}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('3f903eca-8bd3-4f54-a9ce-2b22627db94a', 'gas', 'Wihpe', '{"hp": 1, "volume": 1, "industrialmarket": {"maxprice": 786, "minprice": 131, "silosize": 151978200}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('04f16f34-497f-465e-a9a6-368dd62a5328', 'gas', 'Conren', '{"hp": 1, "volume": 2, "industrialmarket": {"maxprice": 3751, "minprice": 341, "silosize": 195419700}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d8e1ead1-055b-4cd0-b1ec-f88dd0f7ff31', 'gas', 'Uulep', '{"hp": 1, "volume": 3, "industrialmarket": {"maxprice": 3720, "minprice": 620, "silosize": 347427200}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('f4947ebe-4d4c-457b-aeb6-b4dd5b66e62b', 'gas', 'Vaike', '{"hp": 1, "volume": 1, "industrialmarket": {"maxprice": 994, "minprice": 142, "silosize": 895342200}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('0c429117-e7fa-4e00-97f9-ab67c639cae7', 'gas', 'Aschol', '{"hp": 1, "volume": 2, "industrialmarket": {"maxprice": 4080, "minprice": 408, "silosize": 177159300}}');
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('479e8dd7-06e8-479c-b07f-382b407b832f', 'gas', 'Entren', '{"hp": 1, "volume": 3, "industrialmarket": {"maxprice": 1020, "minprice": 340, "silosize": 171851800}}');

-- processes
INSERT INTO public.processes (id, name, meta, "time") VALUES ('cf3dd39f-5479-41b0-a601-e8f0970e502c', 'Tektum Faucet [wm]', '{}', 4945);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('6b9b499d-1426-4aa5-ae58-69f6516d4a9b', 'Tektum Sink [wm]', '{}', 4539);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('9a0216ba-136a-4bd8-a996-237ee48e8ab6', 'Avvon Faucet [wm]', '{}', 4673);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('5e64e9d3-bf33-4552-8d09-1f70d7d2d4a4', 'Avvon Sink [wm]', '{}', 2653);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('e969b65a-949e-44d4-9b5f-46be271553e0', 'Luche Faucet [wm]', '{}', 601);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('e77eee22-24f7-44c7-846a-09819808a8bf', 'Luche Sink [wm]', '{}', 4265);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('5338919e-5c2b-4baa-bb32-fa2ec759bc7c', 'Wihpe Faucet [wm]', '{}', 3589);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('eab6ce76-bee0-4ab7-bc1d-87e7b90af831', 'Wihpe Sink [wm]', '{}', 3999);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('c4378858-d7f0-4be4-80e4-68c9d656f7a3', 'Conren Faucet [wm]', '{}', 5940);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('bdab0eaf-3f9d-40a8-b60f-4770e05f4a06', 'Conren Sink [wm]', '{}', 2962);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('236688a4-bd2d-4add-aeff-6720e5105c55', 'Uulep Faucet [wm]', '{}', 1375);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('c4dde5ab-b782-4d69-9582-7bfd02cd7d1c', 'Uulep Sink [wm]', '{}', 501);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('7b48f3da-3cb8-4bc0-bc2d-dfa9274bc1dd', 'Vaike Faucet [wm]', '{}', 1705);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('4075709e-d2c0-4d20-8698-d504938c22f5', 'Vaike Sink [wm]', '{}', 282);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('e65d4f4c-3c8d-479a-ab89-dd92dc6280d2', 'Aschol Faucet [wm]', '{}', 722);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('f120d11b-e0ad-4a17-bedc-25bb2bb7a06d', 'Aschol Sink [wm]', '{}', 792);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('d1ebdda2-5a2b-4b6d-b454-867a2cd84592', 'Entren Faucet [wm]', '{}', 3403);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('e1e95a63-32f1-4f92-8052-b2ac7a92ec6b', 'Entren Sink [wm]', '{}', 484);

-- process inputs
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('a60c360e-6fb8-4c9c-91d0-f3a3d85c704a', '22fb3cda-c949-41ae-bcf5-ee0a60d497fc', 990, '{}', '6b9b499d-1426-4aa5-ae58-69f6516d4a9b');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('8e98de99-2903-4167-8049-76a78cce7678', '5f84addf-d05b-407a-8dcd-fecd3af4c69d', 415, '{}', '5e64e9d3-bf33-4552-8d09-1f70d7d2d4a4');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('fae34a25-80f7-43d1-bd2d-fb150608cea9', '42179f25-b473-46d1-9592-08e1d23807bd', 924, '{}', 'e77eee22-24f7-44c7-846a-09819808a8bf');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('555dd4db-f7dd-49a9-bbb5-ef6189664bef', '3f903eca-8bd3-4f54-a9ce-2b22627db94a', 690, '{}', 'eab6ce76-bee0-4ab7-bc1d-87e7b90af831');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('1d31e7f4-398c-49fc-8f21-185aef98084b', '04f16f34-497f-465e-a9a6-368dd62a5328', 505, '{}', 'bdab0eaf-3f9d-40a8-b60f-4770e05f4a06');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('dade6a10-e6f0-4c8d-8e2f-25a8dda1b664', 'd8e1ead1-055b-4cd0-b1ec-f88dd0f7ff31', 920, '{}', 'c4dde5ab-b782-4d69-9582-7bfd02cd7d1c');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('57330356-3ec6-4e34-9d23-169f56528afb', 'f4947ebe-4d4c-457b-aeb6-b4dd5b66e62b', 3, '{}', '4075709e-d2c0-4d20-8698-d504938c22f5');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('e6383f95-9d9b-42e0-a107-cee9a79f746f', '0c429117-e7fa-4e00-97f9-ab67c639cae7', 584, '{}', 'f120d11b-e0ad-4a17-bedc-25bb2bb7a06d');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('2a8333f0-f9db-4d5f-9735-f299863955c0', '479e8dd7-06e8-479c-b07f-382b407b832f', 395, '{}', 'e1e95a63-32f1-4f92-8052-b2ac7a92ec6b');

-- process outputs
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('19fdb936-731e-4c0a-a878-f45a70c18010', '22fb3cda-c949-41ae-bcf5-ee0a60d497fc', 601, '{}', 'cf3dd39f-5479-41b0-a601-e8f0970e502c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('67e4c1a0-8520-4b09-b2db-6ee5ec67ec70', '5f84addf-d05b-407a-8dcd-fecd3af4c69d', 158, '{}', '9a0216ba-136a-4bd8-a996-237ee48e8ab6');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('02061d6d-f8fe-406a-a039-e3c42b719804', '42179f25-b473-46d1-9592-08e1d23807bd', 105, '{}', 'e969b65a-949e-44d4-9b5f-46be271553e0');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('5d76baf2-c17f-4978-b79d-caee78c4c71c', '3f903eca-8bd3-4f54-a9ce-2b22627db94a', 196, '{}', '5338919e-5c2b-4baa-bb32-fa2ec759bc7c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('2120dfda-17c4-4e1b-a256-fe0205167f15', '04f16f34-497f-465e-a9a6-368dd62a5328', 432, '{}', 'c4378858-d7f0-4be4-80e4-68c9d656f7a3');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('6dfc31cf-7bfc-483e-b248-010c3fb35452', 'd8e1ead1-055b-4cd0-b1ec-f88dd0f7ff31', 741, '{}', '236688a4-bd2d-4add-aeff-6720e5105c55');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('5719b6ed-dcca-4a49-8853-2dbddec08afe', 'f4947ebe-4d4c-457b-aeb6-b4dd5b66e62b', 393, '{}', '7b48f3da-3cb8-4bc0-bc2d-dfa9274bc1dd');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('3ff72615-fa7d-468c-a15d-91f1c8caa093', '0c429117-e7fa-4e00-97f9-ab67c639cae7', 235, '{}', 'e65d4f4c-3c8d-479a-ab89-dd92dc6280d2');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('137f3848-694c-4876-9ad7-2da5f6d620fa', '479e8dd7-06e8-479c-b07f-382b407b832f', 530, '{}', 'd1ebdda2-5a2b-4b6d-b454-867a2cd84592');
