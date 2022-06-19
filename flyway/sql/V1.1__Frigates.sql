-- fix collima item type
INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('d5eeeac7-d7ce-473a-bd7b-5662d65494e3', 'ship', 'Chollima', '{"volume": 1179834, "shiptemplateid": "e2a67722-10d1-42fd-9a7c-057677ef6e79", "industrialmarket": {"maxprice": 1382784819, "minprice": 326086956, "silosize": 5}}');

UPDATE shiptemplates SET itemtypeid = 'd5eeeac7-d7ce-473a-bd7b-5662d65494e3' WHERE id = 'e2a67722-10d1-42fd-9a7c-057677ef6e79';

-- 