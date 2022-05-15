-- insert custom start
INSERT INTO public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid, wallet, factionid) VALUES ('9a994832-8037-4e22-b4fc-e0d4f081b456', 'Chrissy''s Start', '8d9e032c-d9b1-4a36-8bbf-1448fa60a09a', '{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9bb00839-95f2-4d7c-a7c9-eef60e05fa97"}, {}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}', '2022-05-15 16:01:53.877839-04', false, '0078c73f-1b40-4b22-bdf0-f4927d8bd7bd', '306c0fcd-cf01-43fc-850f-172191fe1582', 250000, '27a53dfc-a321-4c12-bf7c-bb177955c95b');

-- update chrissy's user to use custom start
UPDATE public.users SET StartID = '9a994832-8037-4e22-b4fc-e0d4f081b456' where ID = 'c98932b5-3c81-49ef-b795-b3ca88b1bf95';

