--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2020-12-15 08:02:38

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 50573)
-- Name: containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.containers (
    id uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.containers OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 50337)
-- Name: itemfamilies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.itemfamilies (
    id character varying(16) NOT NULL,
    friendlyname character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemfamilies OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 50343)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id uuid NOT NULL,
    itemtypeid uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL,
    createdby uuid,
    createdreason character varying(64) NOT NULL,
    containerid uuid NOT NULL
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 50349)
-- Name: itemtypes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.itemtypes (
    id uuid NOT NULL,
    family character varying(16) NOT NULL,
    name character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemtypes OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 50355)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    userid uuid NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 50358)
-- Name: ships; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ships (
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    userid uuid NOT NULL,
    pos_x double precision DEFAULT 0 NOT NULL,
    pos_y double precision DEFAULT 0 NOT NULL,
    created timestamp with time zone DEFAULT now() NOT NULL,
    shipname character varying(32) NOT NULL,
    texture character varying(32) DEFAULT 'Mass Testing Brick'::character varying NOT NULL,
    theta double precision DEFAULT 0 NOT NULL,
    vel_x double precision DEFAULT 0 NOT NULL,
    vel_y double precision DEFAULT 0 NOT NULL,
    shield double precision NOT NULL,
    armor double precision NOT NULL,
    hull double precision NOT NULL,
    fuel double precision NOT NULL,
    heat double precision NOT NULL,
    energy double precision NOT NULL,
    shiptemplateid uuid NOT NULL,
    dockedat_stationid uuid,
    fitting jsonb DEFAULT '{}'::jsonb NOT NULL,
    destroyed boolean DEFAULT false NOT NULL,
    destroyedat timestamp with time zone,
    cargobay_containerid uuid NOT NULL,
    fittingbay_containerid uuid NOT NULL
);


ALTER TABLE public.ships OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 50373)
-- Name: shiptemplates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptemplates (
    id uuid NOT NULL,
    created timestamp with time zone DEFAULT now() NOT NULL,
    shiptemplatename character varying(32) NOT NULL,
    texture character varying(32) DEFAULT 'Mass Testing Brick'::character varying NOT NULL,
    radius double precision DEFAULT 0 NOT NULL,
    baseaccel double precision DEFAULT 0 NOT NULL,
    basemass double precision DEFAULT 0 NOT NULL,
    baseturn double precision DEFAULT 0 NOT NULL,
    baseshield double precision DEFAULT 0 NOT NULL,
    baseshieldregen double precision DEFAULT 0 NOT NULL,
    basearmor double precision DEFAULT 0 NOT NULL,
    basehull double precision DEFAULT 0 NOT NULL,
    basefuel double precision DEFAULT 0 NOT NULL,
    baseheatcap double precision DEFAULT 0 NOT NULL,
    baseheatsink double precision DEFAULT 0 NOT NULL,
    baseenergy double precision DEFAULT 0 NOT NULL,
    baseenergyregen double precision DEFAULT 0 NOT NULL,
    shiptypeid uuid NOT NULL,
    slotlayout jsonb DEFAULT '{}'::jsonb NOT NULL,
    basecargobayvolume double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.shiptemplates OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 50395)
-- Name: shiptypes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptypes (
    id uuid NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.shiptypes OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 50398)
-- Name: starts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.starts (
    id uuid NOT NULL,
    name character varying(16) NOT NULL,
    shiptemplateid uuid NOT NULL,
    shipfitting jsonb NOT NULL,
    created timestamp with time zone DEFAULT now() NOT NULL,
    available boolean DEFAULT true NOT NULL,
    systemid uuid NOT NULL,
    homestationid uuid NOT NULL
);


ALTER TABLE public.starts OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 50406)
-- Name: universe_jumpholes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_jumpholes (
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    out_systemid uuid NOT NULL,
    jumpholename character varying(64) NOT NULL,
    pos_x double precision NOT NULL,
    pos_y double precision NOT NULL,
    texture character varying(64) NOT NULL,
    radius double precision DEFAULT 0 NOT NULL,
    mass double precision DEFAULT 0 NOT NULL,
    theta double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.universe_jumpholes OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 50412)
-- Name: universe_planets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_planets (
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    planetname character varying(64) NOT NULL,
    pos_x double precision NOT NULL,
    pos_y double precision NOT NULL,
    texture character varying(64) NOT NULL,
    radius double precision DEFAULT 0 NOT NULL,
    mass double precision DEFAULT 0 NOT NULL,
    theta double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.universe_planets OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 50418)
-- Name: universe_regions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_regions (
    id uuid NOT NULL,
    regionname character varying(32) NOT NULL
);


ALTER TABLE public.universe_regions OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 50421)
-- Name: universe_stars; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_stars (
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    pos_x double precision NOT NULL,
    pos_y double precision NOT NULL,
    texture character varying(64) NOT NULL,
    radius double precision DEFAULT 0 NOT NULL,
    mass double precision DEFAULT 0 NOT NULL,
    theta double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.universe_stars OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 50427)
-- Name: universe_stations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_stations (
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    stationname character varying(64) NOT NULL,
    pos_x double precision NOT NULL,
    pos_y double precision NOT NULL,
    texture character varying(64) NOT NULL,
    radius double precision DEFAULT 0 NOT NULL,
    mass double precision DEFAULT 0 NOT NULL,
    theta double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.universe_stations OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 50433)
-- Name: universe_systems; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_systems (
    id uuid NOT NULL,
    systemname character varying(32) NOT NULL,
    regionid uuid NOT NULL
);


ALTER TABLE public.universe_systems OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 50436)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    username character varying(16) NOT NULL,
    hashpass character(64) NOT NULL,
    registered timestamp with time zone NOT NULL,
    banned bit(1) NOT NULL,
    current_shipid uuid,
    startid uuid NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3003 (class 0 OID 50573)
-- Dependencies: 217
-- Data for Name: containers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.containers (id, meta, created) FROM stdin;
426464cb-f345-4228-a31d-dda5dd80ffa0	{}	2020-12-05 16:02:22.508474-05
14e652c9-35c0-44f4-b1fb-0b2f6575db76	{}	2020-12-05 16:02:22.510474-05
bae78764-aead-4d74-8b6b-f21bf543717f	{}	2020-12-05 16:03:37.438249-05
ea4fc0eb-c7cc-4e20-98fd-00c37d0032dc	{}	2020-12-05 16:03:37.440245-05
87257c60-ee80-4a64-a225-5f98277f6cb3	{}	2020-12-05 17:13:11.959212-05
c96140ed-39a8-455b-a421-eaa59e86fbc2	{}	2020-12-05 17:13:11.961193-05
88d233cc-a808-4f2c-9715-1f94552e5ebd	{}	2020-12-15 07:21:44.391486-05
f7b4b107-2d8a-40f5-9ed3-768d06455d32	{}	2020-12-15 07:21:44.395392-05
ba0b74ee-5e47-42e1-9017-99db5157c961	{}	2020-12-15 07:24:51.839708-05
d7f96a08-cfe2-44df-84bf-b11632992a43	{}	2020-12-15 07:24:51.84752-05
9150a4db-0262-431e-b109-f4ae31f158a0	{}	2020-12-15 07:29:17.99772-05
09ebe22a-d3ea-4804-aabb-1b3ecbf9bbb2	{}	2020-12-15 07:29:17.999672-05
b553d183-10eb-4c7f-92f9-367df6e1c468	{}	2020-12-15 08:00:29.422045-05
f692beab-c24a-4caf-8877-08a7b320fac8	{}	2020-12-15 08:00:29.424974-05
\.


--
-- TOC entry 2988 (class 0 OID 50337)
-- Dependencies: 202
-- Data for Name: itemfamilies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.itemfamilies (id, friendlyname, meta) FROM stdin;
gun_turret	Gun Turret	{}
missile_launcher	Missile Launcher	{}
shield_booster	Shield Booster	{}
fuel_tank	Fuel Tank	{}
armor_plate	Armor Plate	{}
\.


--
-- TOC entry 2989 (class 0 OID 50343)
-- Dependencies: 203
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, itemtypeid, meta, created, createdby, createdreason, containerid) FROM stdin;
1c5b52f2-cb0a-44e7-ac7e-fc30264ef057	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-05 16:02:22.569476-05	15ee3a6d-ef4d-43b1-ae6a-348ba43d2957	Module for new noob ship for player	14e652c9-35c0-44f4-b1fb-0b2f6575db76
7703db34-c9ca-48c6-92c0-82d933c85150	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-05 16:02:22.634474-05	15ee3a6d-ef4d-43b1-ae6a-348ba43d2957	Module for new noob ship for player	14e652c9-35c0-44f4-b1fb-0b2f6575db76
04493c09-1e6a-4470-9a47-cc46eec82337	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-05 16:03:37.506251-05	270613be-4b65-4c55-a9c8-9a88718b308e	Module for new noob ship for player	ea4fc0eb-c7cc-4e20-98fd-00c37d0032dc
ed9b5175-9fb0-4c50-b21c-88cc274992b1	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-05 16:03:37.575244-05	270613be-4b65-4c55-a9c8-9a88718b308e	Module for new noob ship for player	ea4fc0eb-c7cc-4e20-98fd-00c37d0032dc
3279b27f-a191-4aec-8a67-35720abae7e0	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-05 17:13:12.022795-05	7a44a154-1154-44e8-9084-e6c7b8df94f4	Module for new noob ship for player	c96140ed-39a8-455b-a421-eaa59e86fbc2
5bf12388-e1e9-4359-94f9-3de5a02d1487	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-05 17:13:12.086796-05	7a44a154-1154-44e8-9084-e6c7b8df94f4	Module for new noob ship for player	c96140ed-39a8-455b-a421-eaa59e86fbc2
7d59ee15-8d8e-4b32-b090-5ddeb4b026d4	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-05 17:13:12.149794-05	7a44a154-1154-44e8-9084-e6c7b8df94f4	Module for new noob ship for player	c96140ed-39a8-455b-a421-eaa59e86fbc2
f1b54d31-a669-4bc6-97f0-f420ca5d6ef4	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 07:21:44.461794-05	53d2e390-f082-4906-bea8-047e4e9b1550	Module for new noob ship for player	f7b4b107-2d8a-40f5-9ed3-768d06455d32
48ddb903-2f75-48e8-abf8-909c253a8300	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 07:21:44.5413-05	53d2e390-f082-4906-bea8-047e4e9b1550	Module for new noob ship for player	f7b4b107-2d8a-40f5-9ed3-768d06455d32
fd5aa420-eda0-4cfc-a1dc-f5ec86891091	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-15 07:21:44.613561-05	53d2e390-f082-4906-bea8-047e4e9b1550	Module for new noob ship for player	f7b4b107-2d8a-40f5-9ed3-768d06455d32
698b6035-a46b-4089-8b79-4e0fbf6a72d0	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 07:24:51.933447-05	53d2e390-f082-4906-bea8-047e4e9b1550	Module for new noob ship for player	d7f96a08-cfe2-44df-84bf-b11632992a43
c41bb6f5-bd98-4dd9-90a3-f47cac9c9dd2	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 07:24:51.940283-05	53d2e390-f082-4906-bea8-047e4e9b1550	Module for new noob ship for player	d7f96a08-cfe2-44df-84bf-b11632992a43
e0a5ee5e-6d0f-4351-ad10-64f78ace34be	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-15 07:24:52.012545-05	53d2e390-f082-4906-bea8-047e4e9b1550	Module for new noob ship for player	d7f96a08-cfe2-44df-84bf-b11632992a43
6fd6061f-03d9-40bf-b190-292528742f45	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 07:29:18.066198-05	831c5d36-0bea-4d61-8717-54d2fff70f66	Module for new noob ship for player	09ebe22a-d3ea-4804-aabb-1b3ecbf9bbb2
b0c12a98-f412-4f8b-bbd2-b931b64a9629	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 07:29:18.135539-05	831c5d36-0bea-4d61-8717-54d2fff70f66	Module for new noob ship for player	09ebe22a-d3ea-4804-aabb-1b3ecbf9bbb2
946614ea-c4f9-4ab6-8fa2-39879623e971	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-15 07:29:18.2078-05	831c5d36-0bea-4d61-8717-54d2fff70f66	Module for new noob ship for player	09ebe22a-d3ea-4804-aabb-1b3ecbf9bbb2
dfe0bed0-4c19-4368-8634-e14c986ad27b	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2020-12-15 07:29:18.281041-05	831c5d36-0bea-4d61-8717-54d2fff70f66	Module for new noob ship for player	09ebe22a-d3ea-4804-aabb-1b3ecbf9bbb2
67f803c2-fe5b-4316-877d-19ac95cb7d4b	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 08:00:29.491375-05	c2ed6e2c-5bab-4aff-b86d-85ebd8948bcc	Module for new noob ship for player	f692beab-c24a-4caf-8877-08a7b320fac8
65df9c2e-41af-4fa1-bf46-451d1ea2c786	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 08:00:29.562678-05	c2ed6e2c-5bab-4aff-b86d-85ebd8948bcc	Module for new noob ship for player	f692beab-c24a-4caf-8877-08a7b320fac8
52745ad5-a280-490d-ab15-891cd9bb18f3	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-15 08:00:29.634922-05	c2ed6e2c-5bab-4aff-b86d-85ebd8948bcc	Module for new noob ship for player	f692beab-c24a-4caf-8877-08a7b320fac8
af0bf5be-c1a4-45b2-b788-4ebdc4b01ffd	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2020-12-15 08:00:29.709134-05	c2ed6e2c-5bab-4aff-b86d-85ebd8948bcc	Module for new noob ship for player	f692beab-c24a-4caf-8877-08a7b320fac8
f90bad54-9cb3-42c8-bbaa-47fab8d38a9d	c311df30-c21e-4895-acb0-d8808f99710e	{}	2020-12-15 08:00:29.785302-05	c2ed6e2c-5bab-4aff-b86d-85ebd8948bcc	Module for new noob ship for player	f692beab-c24a-4caf-8877-08a7b320fac8
\.


--
-- TOC entry 2990 (class 0 OID 50349)
-- Dependencies: 204
-- Data for Name: itemtypes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.itemtypes (id, family, name, meta) FROM stdin;
9d1014c5-3422-4a0f-9839-f585269b4b16	gun_turret	Basic Laser Tool	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}
b481a521-1b12-4ffa-ac2f-4da015036f7f	fuel_tank	Basic Fuel Tank	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}
09172710-740c-4d1c-9fc0-43cb62e674e7	shield_booster	Basic Shield Booster	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}
c311df30-c21e-4895-acb0-d8808f99710e	armor_plate	Basic Armor Plate	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}
\.


--
-- TOC entry 2991 (class 0 OID 50355)
-- Dependencies: 205
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
1d4955c0-7ac9-4954-ab11-c46d90744d3d	7a44a154-1154-44e8-9084-e6c7b8df94f4
e1def049-605d-499f-bef3-8bdb17fa9b53	53d2e390-f082-4906-bea8-047e4e9b1550
daefa53f-9c9b-4cf9-b2cd-a1db8c585506	15ee3a6d-ef4d-43b1-ae6a-348ba43d2957
12014c4c-30bc-471a-9d7c-2e6c361fd187	831c5d36-0bea-4d61-8717-54d2fff70f66
190d29c1-9e83-4c21-9c9a-c432b7dfe6a9	c2ed6e2c-5bab-4aff-b86d-85ebd8948bcc
4db6fd4c-3a20-4ae1-b856-181270a83d86	270613be-4b65-4c55-a9c8-9a88718b308e
\.


--
-- TOC entry 2992 (class 0 OID 50358)
-- Dependencies: 206
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid, fittingbay_containerid) FROM stdin;
0063178d-e7b4-432e-b63b-6db686b55052	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	7a44a154-1154-44e8-9084-e6c7b8df94f4	48081	-22113	2020-12-05 17:13:12.221795-05	5t6yrhtrgf's Starter Ship	Sparrow	0	0	0	188.64787191741996	169	135	240.41183394719053	83.46425440147871	93.83161356499187	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "3279b27f-a191-4aec-8a67-35720abae7e0", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "5bf12388-e1e9-4359-94f9-3de5a02d1487", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "7d59ee15-8d8e-4b32-b090-5ddeb4b026d4", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": null}	f	\N	87257c60-ee80-4a64-a225-5f98277f6cb3	c96140ed-39a8-455b-a421-eaa59e86fbc2
6c3a9c1c-1e86-4eae-8c00-f8b5e00c1578	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	53d2e390-f082-4906-bea8-047e4e9b1550	48046.60381357499	-21976.83006003722	2020-12-15 07:21:44.686653-05	rtyeurjhgtr's Starter Ship	Sparrow	219.82368583464438	-0.55867177699374	0.6871997918916668	209	169	-0.01367257896916034	216.81261496974955	1325.2350672344496	37.03884094108268	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "f1b54d31-a669-4bc6-97f0-f420ca5d6ef4", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "48ddb903-2f75-48e8-abf8-909c253a8300", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "fd5aa420-eda0-4cfc-a1dc-f5ec86891091", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": null}	t	2020-12-15 07:24:51.825109-05	88d233cc-a808-4f2c-9715-1f94552e5ebd	f7b4b107-2d8a-40f5-9ed3-768d06455d32
3d2e2c4d-e675-4ea0-920d-b99c53cbc94e	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	53d2e390-f082-4906-bea8-047e4e9b1550	-18153	34059	2020-12-15 07:24:52.083853-05	rtyeurjhgtr's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	cf07bba9-90b2-4599-b1e3-84d797a67f0a	{"a_rack": [{"item_id": "698b6035-a46b-4089-8b79-4e0fbf6a72d0", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "c41bb6f5-bd98-4dd9-90a3-f47cac9c9dd2", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "e0a5ee5e-6d0f-4351-ad10-64f78ace34be", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": []}	f	\N	ba0b74ee-5e47-42e1-9017-99db5157c961	d7f96a08-cfe2-44df-84bf-b11632992a43
66287702-b754-49a4-8656-f9dbba937979	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	831c5d36-0bea-4d61-8717-54d2fff70f66	48081	-22113	2020-12-15 07:29:18.353366-05	ui67uyhge4's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "6fd6061f-03d9-40bf-b190-292528742f45", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "b0c12a98-f412-4f8b-bbd2-b931b64a9629", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "946614ea-c4f9-4ab6-8fa2-39879623e971", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_id": "dfe0bed0-4c19-4368-8634-e14c986ad27b", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}]}	f	\N	9150a4db-0262-431e-b109-f4ae31f158a0	09ebe22a-d3ea-4804-aabb-1b3ecbf9bbb2
41f04f60-032c-4817-ac85-4fc294a84b03	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	c2ed6e2c-5bab-4aff-b86d-85ebd8948bcc	48081	-22113	2020-12-15 08:00:29.860493-05	674u6hrgef's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "67f803c2-fe5b-4316-877d-19ac95cb7d4b", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "65df9c2e-41af-4fa1-bf46-451d1ea2c786", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "52745ad5-a280-490d-ab15-891cd9bb18f3", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_id": "af0bf5be-c1a4-45b2-b788-4ebdc4b01ffd", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "f90bad54-9cb3-42c8-bbaa-47fab8d38a9d", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	b553d183-10eb-4c7f-92f9-367df6e1c468	f692beab-c24a-4caf-8877-08a7b320fac8
7143361e-8ea0-4b36-9d7a-d2d0fe8320fa	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	270613be-4b65-4c55-a9c8-9a88718b308e	25453.3422880961	-10256.25297968535	2020-12-05 16:03:37.642244-05	asdf's Starter Ship	Sparrow	205.61911900620103	-9.4e-323	9.4e-323	208.96673785022452	169	135	249.06830983894537	0	137.99755200087688	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "04493c09-1e6a-4470-9a47-cc46eec82337", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "ed9b5175-9fb0-4c50-b21c-88cc274992b1", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	f	\N	bae78764-aead-4d74-8b6b-f21bf543717f	ea4fc0eb-c7cc-4e20-98fd-00c37d0032dc
f60f6df9-fdef-428d-80d2-71df13a0299e	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	15ee3a6d-ef4d-43b1-ae6a-348ba43d2957	24771.795632863843	-9938.30877953488	2020-12-05 16:02:22.697475-05	nwiehoff's Starter Ship	Sparrow	202.71629513311237	0	0	208.95602137287304	169	135	265	0	137.99312314989513	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	cf07bba9-90b2-4599-b1e3-84d797a67f0a	{"a_rack": [{"item_id": "1c5b52f2-cb0a-44e7-ac7e-fc30264ef057", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "7703db34-c9ca-48c6-92c0-82d933c85150", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	f	\N	426464cb-f345-4228-a31d-dda5dd80ffa0	14e652c9-35c0-44f4-b1fb-0b2f6575db76
\.


--
-- TOC entry 2993 (class 0 OID 50373)
-- Dependencies: 207
-- Data for Name: shiptemplates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume) FROM stdin;
8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	2020-11-23 22:14:30.004993-05	Sparrow	Sparrow	12.5	4.3	100	4.7	209	6	169	135	265	737	11	138	9	e364a553-1dc5-4e8d-9195-0ca4989bec49	{"a_slots": [{"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}]}	120
\.


--
-- TOC entry 2994 (class 0 OID 50395)
-- Dependencies: 208
-- Data for Name: shiptypes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shiptypes (id, name) FROM stdin;
e364a553-1dc5-4e8d-9195-0ca4989bec49	Skiff
dcc89c69-28cb-4018-8ee0-1c9e34ff0bca	Transport
bed0330f-eba3-47ed-8e55-84c753c6c376	Frigate
b6be8bdb-37d4-4899-9092-0c5c1901ed62	Cruiser
a7b8e2cf-9e69-480e-a5fa-dc19d8be9a57	Battleship
2d8c098a-b7d8-4518-940b-8c6bfcac311b	Freighter
\.


--
-- TOC entry 2995 (class 0 OID 50398)
-- Dependencies: 209
-- Data for Name: starts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid) FROM stdin;
49f06e89-29fb-4a02-a034-4b5d0702adac	Test Start	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	2020-11-23 15:31:55.475609-05	t	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cf07bba9-90b2-4599-b1e3-84d797a67f0a
\.


--
-- TOC entry 2996 (class 0 OID 50406)
-- Dependencies: 210
-- Data for Name: universe_jumpholes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_jumpholes (id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
0602c3bf-7d70-4a4f-9ebd-77e7cec7ff12	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	edf08406-0879-4141-8af1-f68d32e31c8d	Test Jumphole 1	25000	25000	Jumphole	250	9999999	25
834572ef-b709-4ea7-9cd9-b526744f38cc	edf08406-0879-4141-8af1-f68d32e31c8d	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Jumphole 2	-600000	350000	Jumphole	250	9999999	112
\.


--
-- TOC entry 2997 (class 0 OID 50412)
-- Dependencies: 211
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	207460	-335230	vh_unshaded\\planet03	3960	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	-303350	-100230	vh_unshaded\\planet09	2352	2000000	112
\.


--
-- TOC entry 2998 (class 0 OID 50418)
-- Dependencies: 212
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 2999 (class 0 OID 50421)
-- Dependencies: 213
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	10570	15130	vh_main_sequence\\star_yellow03	28380	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	32670	-11350	vh_main_sequence\\star_red01	15010	1000000000	0
\.


--
-- TOC entry 3000 (class 0 OID 50427)
-- Dependencies: 214
-- Data for Name: universe_stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stations (id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
cf07bba9-90b2-4599-b1e3-84d797a67f0a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Station	24771.795632863843	-9938.30877953488	Sunfarm	740	25300	112.4
526f57f5-09e0-41c7-9a89-cd803ec0a065	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Station	207460	-335230	Fleet Armory	810	65000	23.33
\.


--
-- TOC entry 3001 (class 0 OID 50433)
-- Dependencies: 215
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 3002 (class 0 OID 50436)
-- Dependencies: 216
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid, startid) FROM stdin;
15ee3a6d-ef4d-43b1-ae6a-348ba43d2957	nwiehoff	02c4bf7d7e35c6bab999ac03ece60b8586a27f7ecd4830983b138b74262bf3f9	2020-12-05 16:02:22.436474-05	0	f60f6df9-fdef-428d-80d2-71df13a0299e	49f06e89-29fb-4a02-a034-4b5d0702adac
270613be-4b65-4c55-a9c8-9a88718b308e	asdf	7ab518aea3a72084bf1f1e1ddeb7fa25563a47b73b643452960187b6405b72f1	2020-12-05 16:03:37.351143-05	0	7143361e-8ea0-4b36-9d7a-d2d0fe8320fa	49f06e89-29fb-4a02-a034-4b5d0702adac
7a44a154-1154-44e8-9084-e6c7b8df94f4	5t6yrhtrgf	93494a65c0c4bd0083116d9b3a58e97c30db23f6c48ee55d4114be50b946fc7f	2020-12-05 17:13:11.889212-05	0	0063178d-e7b4-432e-b63b-6db686b55052	49f06e89-29fb-4a02-a034-4b5d0702adac
53d2e390-f082-4906-bea8-047e4e9b1550	rtyeurjhgtr	2f151b8c29ae8094d5cbeef12d535b9eff20168587c4868462a9d933a377aea7	2020-12-15 07:21:44.305303-05	0	3d2e2c4d-e675-4ea0-920d-b99c53cbc94e	49f06e89-29fb-4a02-a034-4b5d0702adac
831c5d36-0bea-4d61-8717-54d2fff70f66	ui67uyhge4	267ae7cd1faf9d491b14e77d7d368c4f89e6141334c8d2b810e10ac2d457fdbc	2020-12-15 07:29:17.921576-05	0	66287702-b754-49a4-8656-f9dbba937979	49f06e89-29fb-4a02-a034-4b5d0702adac
c2ed6e2c-5bab-4aff-b86d-85ebd8948bcc	674u6hrgef	4b8f91123b9cbb9088fa646ee7b41f11520c70549118a02346fa0098e1b23a93	2020-12-15 08:00:29.343925-05	0	41f04f60-032c-4817-ac85-4fc294a84b03	49f06e89-29fb-4a02-a034-4b5d0702adac
\.


--
-- TOC entry 2801 (class 2606 OID 50440)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2807 (class 2606 OID 50442)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 2809 (class 2606 OID 50444)
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- TOC entry 2811 (class 2606 OID 50446)
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- TOC entry 2813 (class 2606 OID 50448)
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- TOC entry 2817 (class 2606 OID 50450)
-- Name: starts starts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT starts_pkey PRIMARY KEY (id);


--
-- TOC entry 2819 (class 2606 OID 50452)
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- TOC entry 2821 (class 2606 OID 50454)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 2823 (class 2606 OID 50456)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 2825 (class 2606 OID 50458)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2827 (class 2606 OID 50460)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 2829 (class 2606 OID 50462)
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- TOC entry 2831 (class 2606 OID 50464)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 2833 (class 2606 OID 50466)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2840 (class 2606 OID 50580)
-- Name: containers uq_container_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.containers
    ADD CONSTRAINT uq_container_id PRIMARY KEY (id);


--
-- TOC entry 2793 (class 2606 OID 50468)
-- Name: itemfamilies uq_itemfamily_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemfamilies
    ADD CONSTRAINT uq_itemfamily_id PRIMARY KEY (id);


--
-- TOC entry 2797 (class 2606 OID 50470)
-- Name: items uq_items_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT uq_items_id PRIMARY KEY (id);


--
-- TOC entry 2799 (class 2606 OID 50472)
-- Name: itemtypes uq_itemtype_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT uq_itemtype_id PRIMARY KEY (id);


--
-- TOC entry 2803 (class 2606 OID 50474)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 2836 (class 2606 OID 50476)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2838 (class 2606 OID 50478)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 2794 (class 1259 OID 50587)
-- Name: fki_fk_items_containers; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_containers ON public.items USING btree (containerid);


--
-- TOC entry 2795 (class 1259 OID 50479)
-- Name: fki_fk_items_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_users ON public.items USING btree (createdby);


--
-- TOC entry 2804 (class 1259 OID 50594)
-- Name: fki_fk_ships_containers_cargobay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_cargobay ON public.ships USING btree (cargobay_containerid);


--
-- TOC entry 2805 (class 1259 OID 50600)
-- Name: fki_fk_ships_containers_fittingbay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_fittingbay ON public.ships USING btree (fittingbay_containerid);


--
-- TOC entry 2814 (class 1259 OID 50480)
-- Name: fki_fk_starts_homestations; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_homestations ON public.starts USING btree (homestationid);


--
-- TOC entry 2815 (class 1259 OID 50481)
-- Name: fki_fk_starts_systems; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_systems ON public.starts USING btree (systemid);


--
-- TOC entry 2834 (class 1259 OID 50482)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 2843 (class 2606 OID 50582)
-- Name: items fk_items_containers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_containers FOREIGN KEY (containerid) REFERENCES public.containers(id);


--
-- TOC entry 2841 (class 2606 OID 50483)
-- Name: items fk_items_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 2842 (class 2606 OID 50488)
-- Name: items fk_items_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_users FOREIGN KEY (createdby) REFERENCES public.users(id);


--
-- TOC entry 2844 (class 2606 OID 50493)
-- Name: itemtypes fk_itemtypes_itemfamilies; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT fk_itemtypes_itemfamilies FOREIGN KEY (family) REFERENCES public.itemfamilies(id);


--
-- TOC entry 2849 (class 2606 OID 50589)
-- Name: ships fk_ships_containers_cargobay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_cargobay FOREIGN KEY (cargobay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2850 (class 2606 OID 50595)
-- Name: ships fk_ships_containers_fittingbay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_fittingbay FOREIGN KEY (fittingbay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2845 (class 2606 OID 50498)
-- Name: ships fk_ships_dockstations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_dockstations FOREIGN KEY (dockedat_stationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2846 (class 2606 OID 50503)
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2847 (class 2606 OID 50508)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2848 (class 2606 OID 50513)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 2851 (class 2606 OID 50518)
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id);


--
-- TOC entry 2852 (class 2606 OID 50523)
-- Name: starts fk_starts_homestations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_homestations FOREIGN KEY (homestationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2853 (class 2606 OID 50528)
-- Name: starts fk_starts_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2854 (class 2606 OID 50533)
-- Name: starts fk_starts_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_systems FOREIGN KEY (systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2860 (class 2606 OID 50538)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 2861 (class 2606 OID 50543)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 2855 (class 2606 OID 50548)
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2856 (class 2606 OID 50553)
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2857 (class 2606 OID 50558)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2858 (class 2606 OID 50563)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2859 (class 2606 OID 50568)
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2020-12-15 08:02:38

--
-- PostgreSQL database dump complete
--

