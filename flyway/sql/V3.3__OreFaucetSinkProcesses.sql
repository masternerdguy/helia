--
-- Data for Name: processes; Type: TABLE DATA; Schema: public; Owner: heliaagent
--

INSERT INTO public.processes (id, name, meta, "time") VALUES ('03471e7f-9265-4e3c-acb7-8c7c56d67801', 'Ore Sink (Generic)', '{}', 99427);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('e763bd26-3a00-44dd-ad9f-c11a214f5bcd', 'Betro Faucet', '{}', 24278);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('76c29e55-e057-4e2e-ac6e-824c19dd5d41', 'Betro Sink', '{}', 68941);
INSERT INTO public.processes (id, name, meta, "time") VALUES ('8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c', 'Ore Faucet (Generic)', '{}', 129600);


--
-- Data for Name: processinputs; Type: TABLE DATA; Schema: public; Owner: heliaagent
--

INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('efe4532c-a430-41e8-ad5e-dd1c01770010', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 25000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0cbf7c99-d66a-4e73-9186-c451bfbbcf22', '56617d30-6c30-425c-84bf-2484ae8c1156', 20000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('f46b170c-fac3-4c2b-bb70-339f6936b597', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 15000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('d3b0c736-a1bd-4565-94fd-b47ec3192661', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 10000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('bab106c0-ebb6-4e21-80af-6a38362fa4f5', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 9000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('0856f46a-7914-48c1-8885-da97ca985534', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 8000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('54893f30-fa76-49b3-9797-52c759840b13', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 7000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('26580712-8092-44ef-a5ec-313c108ca360', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 6000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('29117d5e-f66d-4252-8516-643f358e0e61', '61f52ba3-654b-45cf-88e3-33399d12350d', 5000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('97965463-dad1-4799-b511-be6273372372', '11688112-f3d4-4d30-864a-684a8b96ea23', 4000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('fa68bcd0-e6f5-402f-b920-3d33b7dcc0f8', '2ce48bef-f06b-4550-b20c-0e64864db051', 3000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('16df64f1-9155-4705-b7bd-f4a0e7ea8d63', '66b7a322-8cfc-4467-9410-492e6b58f159', 2000, '{}', '03471e7f-9265-4e3c-acb7-8c7c56d67801');
INSERT INTO public.processinputs (id, itemtypeid, quantity, meta, processid) VALUES ('fe525e2c-f55f-45c2-8bed-188bbaf74372', 'd1866be4-5c3e-4b95-b6d9-020832338014', 500, '{}', '76c29e55-e057-4e2e-ac6e-824c19dd5d41');


--
-- Data for Name: processoutputs; Type: TABLE DATA; Schema: public; Owner: heliaagent
--

INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('5748a0df-5a51-4296-9ecf-f24d9181e0be', '66b7a322-8cfc-4467-9410-492e6b58f159', 2000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('361f1ad9-6f0f-4f61-9f7b-75b329ce56e0', '2ce48bef-f06b-4550-b20c-0e64864db051', 3000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('f0816cc9-1b2f-4d16-9115-141dbe16e433', '11688112-f3d4-4d30-864a-684a8b96ea23', 4000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('281d0faa-e1fb-4e38-af67-b735a7ae516d', '61f52ba3-654b-45cf-88e3-33399d12350d', 5000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('dea944fb-d35a-4c41-a3c0-1ea517783c25', '7dcd5138-d7e0-419f-867a-6f0f23b99b5b', 6000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('1d797ddf-dc83-439a-ab31-fb74bc1f768f', 'dd0c9b0a-279e-418e-b3b6-2f569fda0186', 7000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('7cd64b26-3d81-41be-b385-ba556aa839d7', '39b8eedf-ef80-4c29-a4bf-99abc4d84fa6', 8000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('8d4e4157-deb0-4fca-a051-cb30ef9019a9', '0cd04eea-a150-410c-91eb-6af00d8c6eae', 9000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('39c79800-05fb-4cf8-a74b-697e7fdc0d40', '1d0d344b-ef28-43c8-a7a6-3275936b2dea', 10000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('017793c7-b07c-4cbe-9cca-72d85bebd7ee', '26a3fc9e-db2f-439d-a929-ba755d11d09c', 15000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 'dd522f03-2f52-4e82-b2f8-d7e0029cb82f', 25000, '{}', '8178e044-b8fc-4e8f-bd46-ed9ab3d9a66c');
INSERT INTO public.processoutputs (id, itemtypeid, quantity, meta, processid) VALUES ('9096676e-3f85-4a0a-921b-0d12504e05a3', 'd1866be4-5c3e-4b95-b6d9-020832338014', 500, '{}', 'e763bd26-3a00-44dd-ad9f-c11a214f5bcd');
