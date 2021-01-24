--
-- PostgreSQL database dump
--

-- Dumped from database version 10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)

-- Started on 2021-01-24 01:21:22 EST

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

--
-- TOC entry 1 (class 3079 OID 13164)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 3258 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 25972)
-- Name: containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.containers (
    id uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.containers OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 25978)
-- Name: itemfamilies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.itemfamilies (
    id character varying(16) NOT NULL,
    friendlyname character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemfamilies OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 25984)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id uuid NOT NULL,
    itemtypeid uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL,
    createdby uuid,
    createdreason character varying(64) NOT NULL,
    containerid uuid NOT NULL,
    quantity integer DEFAULT 1 NOT NULL,
    ispackaged boolean DEFAULT false NOT NULL
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 199 (class 1259 OID 25992)
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
-- TOC entry 213 (class 1259 OID 26260)
-- Name: sellorders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sellorders (
    id uuid NOT NULL,
    universe_stationid uuid NOT NULL,
    itemid uuid NOT NULL,
    seller_userid uuid NOT NULL,
    askprice double precision NOT NULL,
    created timestamp with time zone NOT NULL,
    bought timestamp with time zone,
    buyer_userid uuid
);


ALTER TABLE public.sellorders OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 25998)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    userid uuid NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 26001)
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
    fittingbay_containerid uuid NOT NULL,
    remaxdirty boolean DEFAULT true NOT NULL,
    trash_containerid uuid NOT NULL,
    wallet double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.ships OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 26017)
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
-- TOC entry 203 (class 1259 OID 26040)
-- Name: shiptypes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptypes (
    id uuid NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.shiptypes OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 26043)
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
    homestationid uuid NOT NULL,
    wallet double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.starts OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 26051)
-- Name: universe_asteroids; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_asteroids (
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    ore_itemtypeid uuid NOT NULL,
    name character varying(8) NOT NULL,
    texture character varying(255) NOT NULL,
    radius double precision NOT NULL,
    theta double precision NOT NULL,
    pos_x double precision NOT NULL,
    pos_y double precision NOT NULL,
    yield double precision NOT NULL,
    mass double precision NOT NULL
);


ALTER TABLE public.universe_asteroids OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 26054)
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
-- TOC entry 207 (class 1259 OID 26060)
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
-- TOC entry 208 (class 1259 OID 26066)
-- Name: universe_regions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_regions (
    id uuid NOT NULL,
    regionname character varying(32) NOT NULL
);


ALTER TABLE public.universe_regions OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 26069)
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
-- TOC entry 210 (class 1259 OID 26075)
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
-- TOC entry 211 (class 1259 OID 26081)
-- Name: universe_systems; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_systems (
    id uuid NOT NULL,
    systemname character varying(32) NOT NULL,
    regionid uuid NOT NULL
);


ALTER TABLE public.universe_systems OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 26084)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    username character varying(16) NOT NULL,
    hashpass character(64) NOT NULL,
    registered timestamp with time zone NOT NULL,
    banned bit(1) NOT NULL,
    current_shipid uuid,
    startid uuid NOT NULL,
    escrow_containerid uuid NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3233 (class 0 OID 25972)
-- Dependencies: 196
-- Data for Name: containers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.containers (id, meta, created) FROM stdin;
9c283482-76d4-4139-ae6c-3138345ad19c	{}	2021-01-24 01:03:56.091615-05
bbc5a814-6bb1-4955-a0b5-d9580f692e2e	{}	2021-01-24 01:03:56.098964-05
10f9eb73-7ede-4c60-909e-0a9341bf41cc	{}	2021-01-24 01:03:56.100486-05
82ffe4dd-7343-4191-9652-2940ae5a4f42	{}	2021-01-24 01:03:56.101879-05
\.


--
-- TOC entry 3234 (class 0 OID 25978)
-- Dependencies: 197
-- Data for Name: itemfamilies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.itemfamilies (id, friendlyname, meta) FROM stdin;
gun_turret	Gun Turret	{}
missile_launcher	Missile Launcher	{}
shield_booster	Shield Booster	{}
fuel_tank	Fuel Tank	{}
armor_plate	Armor Plate	{}
nothing	Empty Space	{}
ore	Ore	{}
\.


--
-- TOC entry 3235 (class 0 OID 25984)
-- Dependencies: 198
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, itemtypeid, meta, created, createdby, createdreason, containerid, quantity, ispackaged) FROM stdin;
58c57c2d-efea-4fdd-b4af-5a851ccec732	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-24 01:03:56.10364-05	22e53c8f-d5ad-46dd-827f-03204644ddf7	Module for new noob ship for player	82ffe4dd-7343-4191-9652-2940ae5a4f42	1	f
a8fb8932-a2bb-47b7-ab6d-bd632a906d81	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-24 01:03:56.106008-05	22e53c8f-d5ad-46dd-827f-03204644ddf7	Module for new noob ship for player	82ffe4dd-7343-4191-9652-2940ae5a4f42	1	f
cce56f8c-ac40-477a-87c7-5e23fc8f4f20	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-24 01:03:56.107989-05	22e53c8f-d5ad-46dd-827f-03204644ddf7	Module for new noob ship for player	82ffe4dd-7343-4191-9652-2940ae5a4f42	1	f
4cfafc0a-85fa-490f-9a6b-5ea5c40641aa	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-24 01:03:56.110175-05	22e53c8f-d5ad-46dd-827f-03204644ddf7	Module for new noob ship for player	82ffe4dd-7343-4191-9652-2940ae5a4f42	1	f
5dae670e-482d-46b0-94b8-e4ab5e7a4737	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-24 01:03:56.112507-05	22e53c8f-d5ad-46dd-827f-03204644ddf7	Module for new noob ship for player	82ffe4dd-7343-4191-9652-2940ae5a4f42	1	f
540ac6ca-4566-4d29-aaaa-008d9d206cfd	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-24 01:12:46.345811-05	22e53c8f-d5ad-46dd-827f-03204644ddf7	Mined ore	10f9eb73-7ede-4c60-909e-0a9341bf41cc	20	t
\.


--
-- TOC entry 3236 (class 0 OID 25992)
-- Dependencies: 199
-- Data for Name: itemtypes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.itemtypes (id, family, name, meta) FROM stdin;
9d1014c5-3422-4a0f-9839-f585269b4b16	gun_turret	Basic Laser Tool	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}
b481a521-1b12-4ffa-ac2f-4da015036f7f	fuel_tank	Basic Fuel Tank	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}
09172710-740c-4d1c-9fc0-43cb62e674e7	shield_booster	Basic Shield Booster	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}
c311df30-c21e-4895-acb0-d8808f99710e	armor_plate	Basic Armor Plate	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}
91ec9901-ea7e-476a-bc65-7da4523dca38	nothing	Nothing	{"volume": 0}
dd522f03-2f52-4e82-b2f8-d7e0029cb82f	ore	Testite	{"hp": 1, "volume": 1}
\.


--
-- TOC entry 3250 (class 0 OID 26260)
-- Dependencies: 213
-- Data for Name: sellorders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sellorders (id, universe_stationid, itemid, seller_userid, askprice, created, bought, buyer_userid) FROM stdin;
\.


--
-- TOC entry 3237 (class 0 OID 25998)
-- Dependencies: 200
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
3a42d5fd-afe3-4132-94b7-34bf75e6d840	22e53c8f-d5ad-46dd-827f-03204644ddf7
\.


--
-- TOC entry 3238 (class 0 OID 26001)
-- Dependencies: 201
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid, fittingbay_containerid, remaxdirty, trash_containerid, wallet) FROM stdin;
5452e22f-7e5c-4626-8828-696427c0ee8c	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	22e53c8f-d5ad-46dd-827f-03204644ddf7	24771.7956328638429	-9938.30877953488016	2021-01-24 01:03:56.114356-05	aaa's Starter Ship	Sparrow	344.522670160362395	0	0	209	244	135	295	0	137.994586263279388	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	cf07bba9-90b2-4599-b1e3-84d797a67f0a	{"a_rack": [{"item_id": "58c57c2d-efea-4fdd-b4af-5a851ccec732", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "a8fb8932-a2bb-47b7-ab6d-bd632a906d81", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "cce56f8c-ac40-477a-87c7-5e23fc8f4f20", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "4cfafc0a-85fa-490f-9a6b-5ea5c40641aa", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "5dae670e-482d-46b0-94b8-e4ab5e7a4737", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	10f9eb73-7ede-4c60-909e-0a9341bf41cc	82ffe4dd-7343-4191-9652-2940ae5a4f42	f	bbc5a814-6bb1-4955-a0b5-d9580f692e2e	10000
\.


--
-- TOC entry 3239 (class 0 OID 26017)
-- Dependencies: 202
-- Data for Name: shiptemplates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume) FROM stdin;
8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	2020-11-23 22:14:30.004993-05	Sparrow	Sparrow	12.5	4.29999999999999982	100	4.70000000000000018	209	6	169	135	265	737	11	138	9	e364a553-1dc5-4e8d-9195-0ca4989bec49	{"a_slots": [{"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}]}	120
\.


--
-- TOC entry 3240 (class 0 OID 26040)
-- Dependencies: 203
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
-- TOC entry 3241 (class 0 OID 26043)
-- Dependencies: 204
-- Data for Name: starts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid, wallet) FROM stdin;
49f06e89-29fb-4a02-a034-4b5d0702adac	Test Start	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	2020-11-23 15:31:55.475609-05	t	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cf07bba9-90b2-4599-b1e3-84d797a67f0a	10000
\.


--
-- TOC entry 3242 (class 0 OID 26051)
-- Dependencies: 205
-- Data for Name: universe_asteroids; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_asteroids (id, universe_systemid, ore_itemtypeid, name, texture, radius, theta, pos_x, pos_y, yield, mass) FROM stdin;
231ac943-7fca-42db-a8ef-07f35690af3b	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	XYZ-123	Mega\\asteroidR4	220	45.3500000000000014	-62454	-34091	1.25	10000
\.


--
-- TOC entry 3243 (class 0 OID 26054)
-- Dependencies: 206
-- Data for Name: universe_jumpholes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_jumpholes (id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
0602c3bf-7d70-4a4f-9ebd-77e7cec7ff12	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	edf08406-0879-4141-8af1-f68d32e31c8d	Test Jumphole 1	25000	25000	Jumphole	250	9999999	25
834572ef-b709-4ea7-9cd9-b526744f38cc	edf08406-0879-4141-8af1-f68d32e31c8d	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Jumphole 2	-600000	350000	Jumphole	250	9999999	112
\.


--
-- TOC entry 3244 (class 0 OID 26060)
-- Dependencies: 207
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	207460	-335230	vh_unshaded\\planet03	3960	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	-303350	-100230	vh_unshaded\\planet09	2352	2000000	112
\.


--
-- TOC entry 3245 (class 0 OID 26066)
-- Dependencies: 208
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 3246 (class 0 OID 26069)
-- Dependencies: 209
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	10570	15130	vh_main_sequence\\star_yellow03	28380	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	32670	-11350	vh_main_sequence\\star_red01	15010	1000000000	0
\.


--
-- TOC entry 3247 (class 0 OID 26075)
-- Dependencies: 210
-- Data for Name: universe_stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stations (id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
526f57f5-09e0-41c7-9a89-cd803ec0a065	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Station	207460	-335230	Fleet Armory	810	65000	23.3299999999999983
cf07bba9-90b2-4599-b1e3-84d797a67f0a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Station	24771.7956328638429	-9938.30877953488016	Sunfarm	740	25300	112.400000000000006
\.


--
-- TOC entry 3248 (class 0 OID 26081)
-- Dependencies: 211
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 3249 (class 0 OID 26084)
-- Dependencies: 212
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid, startid, escrow_containerid) FROM stdin;
22e53c8f-d5ad-46dd-827f-03204644ddf7	aaa	09b6eea09b48ac41c136ef0515927f468dd413d64406b97660a4b1de14c15d0c	2021-01-24 01:03:56.094014-05	0	5452e22f-7e5c-4626-8828-696427c0ee8c	49f06e89-29fb-4a02-a034-4b5d0702adac	9c283482-76d4-4139-ae6c-3138345ad19c
\.


--
-- TOC entry 3039 (class 2606 OID 26088)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 3046 (class 2606 OID 26090)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 3048 (class 2606 OID 26092)
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- TOC entry 3050 (class 2606 OID 26094)
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- TOC entry 3052 (class 2606 OID 26096)
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- TOC entry 3056 (class 2606 OID 26098)
-- Name: starts starts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT starts_pkey PRIMARY KEY (id);


--
-- TOC entry 3060 (class 2606 OID 26100)
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- TOC entry 3062 (class 2606 OID 26102)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 3064 (class 2606 OID 26104)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 3066 (class 2606 OID 26106)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 3068 (class 2606 OID 26108)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 3070 (class 2606 OID 26110)
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- TOC entry 3072 (class 2606 OID 26112)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 3074 (class 2606 OID 26114)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 3058 (class 2606 OID 26116)
-- Name: universe_asteroids uq_asteroid_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT uq_asteroid_id PRIMARY KEY (id);


--
-- TOC entry 3029 (class 2606 OID 26118)
-- Name: containers uq_container_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.containers
    ADD CONSTRAINT uq_container_id PRIMARY KEY (id);


--
-- TOC entry 3031 (class 2606 OID 26120)
-- Name: itemfamilies uq_itemfamily_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemfamilies
    ADD CONSTRAINT uq_itemfamily_id PRIMARY KEY (id);


--
-- TOC entry 3035 (class 2606 OID 26122)
-- Name: items uq_items_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT uq_items_id PRIMARY KEY (id);


--
-- TOC entry 3037 (class 2606 OID 26124)
-- Name: itemtypes uq_itemtype_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT uq_itemtype_id PRIMARY KEY (id);


--
-- TOC entry 3082 (class 2606 OID 26264)
-- Name: sellorders uq_sellorders_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT uq_sellorders_pk PRIMARY KEY (id);


--
-- TOC entry 3041 (class 2606 OID 26126)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 3078 (class 2606 OID 26128)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 3080 (class 2606 OID 26130)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 3032 (class 1259 OID 26131)
-- Name: fki_fk_items_containers; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_containers ON public.items USING btree (containerid);


--
-- TOC entry 3033 (class 1259 OID 26132)
-- Name: fki_fk_items_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_users ON public.items USING btree (createdby);


--
-- TOC entry 3042 (class 1259 OID 26133)
-- Name: fki_fk_ships_containers_cargobay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_cargobay ON public.ships USING btree (cargobay_containerid);


--
-- TOC entry 3043 (class 1259 OID 26134)
-- Name: fki_fk_ships_containers_fittingbay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_fittingbay ON public.ships USING btree (fittingbay_containerid);


--
-- TOC entry 3044 (class 1259 OID 26135)
-- Name: fki_fk_ships_containers_trash; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_trash ON public.ships USING btree (trash_containerid);


--
-- TOC entry 3053 (class 1259 OID 26136)
-- Name: fki_fk_starts_homestations; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_homestations ON public.starts USING btree (homestationid);


--
-- TOC entry 3054 (class 1259 OID 26137)
-- Name: fki_fk_starts_systems; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_systems ON public.starts USING btree (systemid);


--
-- TOC entry 3075 (class 1259 OID 26301)
-- Name: fki_fk_users_containers_escrow; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_containers_escrow ON public.users USING btree (escrow_containerid);


--
-- TOC entry 3076 (class 1259 OID 26138)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 3098 (class 2606 OID 26139)
-- Name: universe_asteroids fk_asteroids_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT fk_asteroids_itemtypes FOREIGN KEY (ore_itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 3099 (class 2606 OID 26144)
-- Name: universe_asteroids fk_asteroids_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT fk_asteroids_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3083 (class 2606 OID 26149)
-- Name: items fk_items_containers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_containers FOREIGN KEY (containerid) REFERENCES public.containers(id);


--
-- TOC entry 3084 (class 2606 OID 26154)
-- Name: items fk_items_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 3085 (class 2606 OID 26159)
-- Name: items fk_items_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_users FOREIGN KEY (createdby) REFERENCES public.users(id);


--
-- TOC entry 3086 (class 2606 OID 26164)
-- Name: itemtypes fk_itemtypes_itemfamilies; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT fk_itemtypes_itemfamilies FOREIGN KEY (family) REFERENCES public.itemfamilies(id);


--
-- TOC entry 3109 (class 2606 OID 26270)
-- Name: sellorders fk_sellorders_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT fk_sellorders_items FOREIGN KEY (itemid) REFERENCES public.items(id);


--
-- TOC entry 3108 (class 2606 OID 26265)
-- Name: sellorders fk_sellorders_stations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT fk_sellorders_stations FOREIGN KEY (universe_stationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 3111 (class 2606 OID 26280)
-- Name: sellorders fk_sellorders_users_buyers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT fk_sellorders_users_buyers FOREIGN KEY (buyer_userid) REFERENCES public.users(id);


--
-- TOC entry 3110 (class 2606 OID 26275)
-- Name: sellorders fk_sellorders_users_sellers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT fk_sellorders_users_sellers FOREIGN KEY (seller_userid) REFERENCES public.users(id);


--
-- TOC entry 3087 (class 2606 OID 26169)
-- Name: ships fk_ships_containers_cargobay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_cargobay FOREIGN KEY (cargobay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 3088 (class 2606 OID 26174)
-- Name: ships fk_ships_containers_fittingbay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_fittingbay FOREIGN KEY (fittingbay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 3089 (class 2606 OID 26179)
-- Name: ships fk_ships_containers_trash; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_trash FOREIGN KEY (trash_containerid) REFERENCES public.containers(id);


--
-- TOC entry 3090 (class 2606 OID 26184)
-- Name: ships fk_ships_dockstations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_dockstations FOREIGN KEY (dockedat_stationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 3091 (class 2606 OID 26189)
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 3092 (class 2606 OID 26194)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3093 (class 2606 OID 26199)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 3094 (class 2606 OID 26204)
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id);


--
-- TOC entry 3095 (class 2606 OID 26209)
-- Name: starts fk_starts_homestations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_homestations FOREIGN KEY (homestationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 3096 (class 2606 OID 26214)
-- Name: starts fk_starts_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 3097 (class 2606 OID 26219)
-- Name: starts fk_starts_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_systems FOREIGN KEY (systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3105 (class 2606 OID 26224)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 3107 (class 2606 OID 26296)
-- Name: users fk_users_containers_escrow; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_containers_escrow FOREIGN KEY (escrow_containerid) REFERENCES public.containers(id);


--
-- TOC entry 3106 (class 2606 OID 26229)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 3100 (class 2606 OID 26234)
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3101 (class 2606 OID 26239)
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3102 (class 2606 OID 26244)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3103 (class 2606 OID 26249)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3104 (class 2606 OID 26254)
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2021-01-24 01:21:22 EST

--
-- PostgreSQL database dump complete
--

