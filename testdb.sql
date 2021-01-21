--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2021-01-21 00:25:07

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
-- TOC entry 202 (class 1259 OID 50877)
-- Name: containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.containers (
    id uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.containers OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 50883)
-- Name: itemfamilies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.itemfamilies (
    id character varying(16) NOT NULL,
    friendlyname character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemfamilies OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 50889)
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
-- TOC entry 205 (class 1259 OID 50897)
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
-- TOC entry 206 (class 1259 OID 50903)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    userid uuid NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 50906)
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
    trash_containerid uuid NOT NULL
);


ALTER TABLE public.ships OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 50922)
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
-- TOC entry 209 (class 1259 OID 50945)
-- Name: shiptypes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptypes (
    id uuid NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.shiptypes OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 50948)
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
-- TOC entry 218 (class 1259 OID 51149)
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
-- TOC entry 211 (class 1259 OID 50956)
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
-- TOC entry 212 (class 1259 OID 50962)
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
-- TOC entry 213 (class 1259 OID 50968)
-- Name: universe_regions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_regions (
    id uuid NOT NULL,
    regionname character varying(32) NOT NULL
);


ALTER TABLE public.universe_regions OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 50971)
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
-- TOC entry 215 (class 1259 OID 50977)
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
-- TOC entry 216 (class 1259 OID 50983)
-- Name: universe_systems; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_systems (
    id uuid NOT NULL,
    systemname character varying(32) NOT NULL,
    regionid uuid NOT NULL
);


ALTER TABLE public.universe_systems OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 50986)
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
-- TOC entry 3001 (class 0 OID 50877)
-- Dependencies: 202
-- Data for Name: containers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.containers (id, meta, created) FROM stdin;
7bf8f0fc-8d62-4178-99d4-c79be5d31dc1	{}	2021-01-18 22:15:44.865334-05
ab59dcf4-867b-49eb-9cec-ec576b3343ab	{}	2021-01-18 22:15:44.86803-05
62d60ec7-2314-4c3a-a1f6-be0e6dee8f44	{}	2021-01-18 22:15:44.937632-05
f14e2b57-fe09-474d-99c1-a96faa6bf2e7	{}	2021-01-20 23:19:55.878703-05
8ba3f99a-95dd-4e2d-b0e9-f7dae82ae4e7	{}	2021-01-20 23:19:55.882704-05
fde70085-e845-43d9-9677-0ef6ad59d231	{}	2021-01-20 23:19:55.975702-05
a420c826-33d5-4f1e-b2d4-90a26bb20ee9	{}	2021-01-20 23:22:39.114089-05
93bca1ac-0843-4b7d-85c8-7853e1fa1322	{}	2021-01-20 23:22:39.118088-05
58da71d3-f66e-487f-8879-10c713750044	{}	2021-01-20 23:22:39.195094-05
1cb7675d-bcb4-41c1-b651-df462b8e7d1f	{}	2021-01-21 00:20:33.454296-05
94ef2a5b-3e37-46b8-bde0-5b6fecbf8abf	{}	2021-01-21 00:20:33.455294-05
d79d2bf7-b7a5-47ed-a27e-19e278cae35c	{}	2021-01-21 00:20:33.456296-05
\.


--
-- TOC entry 3002 (class 0 OID 50883)
-- Dependencies: 203
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
-- TOC entry 3003 (class 0 OID 50889)
-- Dependencies: 204
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, itemtypeid, meta, created, createdby, createdreason, containerid, quantity, ispackaged) FROM stdin;
b9f5aed1-f1cb-456f-a389-d503ca378910	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-18 22:15:45.074086-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Module for new noob ship for player	62d60ec7-2314-4c3a-a1f6-be0e6dee8f44	1	f
cad1206b-997d-4788-9cda-6eb5a866f2db	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-18 22:15:45.152077-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Module for new noob ship for player	62d60ec7-2314-4c3a-a1f6-be0e6dee8f44	1	f
0a3e786e-f2c8-4372-89cf-d07158079c27	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-18 22:15:45.228293-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Module for new noob ship for player	62d60ec7-2314-4c3a-a1f6-be0e6dee8f44	1	f
240d170e-c23e-4aaf-8a97-75d6a2c220ed	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-18 22:15:45.309292-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Module for new noob ship for player	62d60ec7-2314-4c3a-a1f6-be0e6dee8f44	1	f
65cc5524-1b54-40d5-b581-42c0aa220ac7	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-18 22:15:45.395292-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Module for new noob ship for player	62d60ec7-2314-4c3a-a1f6-be0e6dee8f44	1	f
9d871fae-0958-42a7-ba34-5328f43dbfac	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:17:56.314467-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Mined ore	7bf8f0fc-8d62-4178-99d4-c79be5d31dc1	119	t
bec48914-2347-4993-b117-96613bdfc631	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:04:20.938512-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Mined ore	ab59dcf4-867b-49eb-9cec-ec576b3343ab	0	t
3037b3a0-6d3f-4120-a724-d6e3efec5e72	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-20 23:19:56.168724-05	c6235bd4-f586-433f-95e9-60020b6f564e	Module for new noob ship for player	fde70085-e845-43d9-9677-0ef6ad59d231	1	f
e8f59164-87ed-48dd-baab-4629ab6a902f	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-20 23:19:56.258703-05	c6235bd4-f586-433f-95e9-60020b6f564e	Module for new noob ship for player	fde70085-e845-43d9-9677-0ef6ad59d231	1	f
0954d23b-0451-4147-a3ce-e46f58c704ba	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-20 23:19:56.344703-05	c6235bd4-f586-433f-95e9-60020b6f564e	Module for new noob ship for player	fde70085-e845-43d9-9677-0ef6ad59d231	1	f
ca12e643-6152-4ca4-b0a0-1b6a884baf8c	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:10:40.202733-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Stack split	7bf8f0fc-8d62-4178-99d4-c79be5d31dc1	41	t
62d17925-486d-4b5e-86ff-ce9d0ba52a6e	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:10:56.145837-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Stack split	7bf8f0fc-8d62-4178-99d4-c79be5d31dc1	77	t
c5d268ff-f0ff-40aa-be28-521e0e74e4e9	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:11:00.704564-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Stack split	7bf8f0fc-8d62-4178-99d4-c79be5d31dc1	1	t
9560464f-ccf8-4fe7-ae26-9298fe7e0a0a	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-20 23:19:56.437703-05	c6235bd4-f586-433f-95e9-60020b6f564e	Module for new noob ship for player	fde70085-e845-43d9-9677-0ef6ad59d231	1	f
e9581554-b3ca-45cf-9763-5b1c0ed3e50b	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-20 23:19:56.542703-05	c6235bd4-f586-433f-95e9-60020b6f564e	Module for new noob ship for player	fde70085-e845-43d9-9677-0ef6ad59d231	1	f
cbfe99cc-9da9-4439-9fbf-a1fe93ea1728	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-20 23:22:39.354087-05	a0fd0657-776b-424a-afe7-db6633f7e283	Module for new noob ship for player	58da71d3-f66e-487f-8879-10c713750044	1	f
79020503-827b-457d-8a62-10bd9e33a3c2	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-20 23:22:39.444088-05	a0fd0657-776b-424a-afe7-db6633f7e283	Module for new noob ship for player	58da71d3-f66e-487f-8879-10c713750044	1	f
030ca987-c825-4a80-9ab1-6b7c893f81ab	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-20 23:22:39.547089-05	a0fd0657-776b-424a-afe7-db6633f7e283	Module for new noob ship for player	58da71d3-f66e-487f-8879-10c713750044	1	f
c7d20552-21c9-47ad-9a0e-8e568280d191	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-20 23:22:39.634088-05	a0fd0657-776b-424a-afe7-db6633f7e283	Module for new noob ship for player	58da71d3-f66e-487f-8879-10c713750044	1	f
a1e7d39f-a623-4ab6-b0ce-fee765591afa	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-20 23:22:39.720088-05	a0fd0657-776b-424a-afe7-db6633f7e283	Module for new noob ship for player	58da71d3-f66e-487f-8879-10c713750044	1	f
72f57ac6-51d8-4bac-9e8e-f51184f64aa3	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:37.902384-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
47a57339-aa13-481a-ba7d-e1db1c76e203	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:51:27.167689-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	a420c826-33d5-4f1e-b2d4-90a26bb20ee9	46	t
9c8a8e84-c872-4f30-a884-b0de49057346	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:51:32.995289-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Mined ore	7bf8f0fc-8d62-4178-99d4-c79be5d31dc1	44	t
135d5d81-f176-4079-87ee-96049afee88e	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:42.417675-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
ef086d06-51c2-48b3-9bbc-aeb25ad93b1c	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:21.48616-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
eadf0b72-6fa6-4a57-9632-a7812809d8e5	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:57:04.222605-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Mined ore	7bf8f0fc-8d62-4178-99d4-c79be5d31dc1	101	t
3c1e6dfd-7789-4627-bc22-941057e30bf8	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:22.202822-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
f6c2a501-6744-407f-8a6f-17ef78137a62	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:43.119581-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
36b7b7ab-394b-4d3a-8e87-8d68c564f44e	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:57:11.249709-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	a420c826-33d5-4f1e-b2d4-90a26bb20ee9	83	t
142938c0-9bc3-4ed7-87e3-98683a816afd	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:26.691146-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
f56dfe0d-eb0c-483d-bebf-e838f3c722ec	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:47.664936-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
67775574-56f0-4dcd-b797-ca7077cbbdef	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:27.426796-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
daf3ec71-dc6d-404f-986e-7aa1e6463eae	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:31.917058-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
d11e63cf-ecbb-4536-92da-e22e60075fb6	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:48.330694-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
c9121ad9-e780-4957-8196-7a912eaa31f3	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:32.652518-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
94f9365f-a040-40cb-9622-18bb62b32861	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:52.810212-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
52893fba-bb4e-4028-a968-9f2171ddd172	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:37.158932-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
40829749-8846-4479-b358-4c2f70fee624	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:53.555826-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
53d9a1f6-2aaa-4553-ae0e-b28a14442f21	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:58.846719-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	a420c826-33d5-4f1e-b2d4-90a26bb20ee9	118	t
c369b896-3b01-4ffa-9eac-da92911fd3ed	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:23:58.072719-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	0	t
be07f8b9-5d84-4e34-b8b2-b731e1b617f2	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-20 23:43:27.507133-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	a420c826-33d5-4f1e-b2d4-90a26bb20ee9	2	t
be94bc50-6ac8-4a8d-9d85-ab35046d1918	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-21 00:20:33.458295-05	c0c5c4ee-5eb8-41da-8d4f-091905a890e9	Module for new noob ship for player	d79d2bf7-b7a5-47ed-a27e-19e278cae35c	1	f
3791b5e9-2892-4bc8-9cff-0600f5c967c5	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-21 00:20:33.461295-05	c0c5c4ee-5eb8-41da-8d4f-091905a890e9	Module for new noob ship for player	d79d2bf7-b7a5-47ed-a27e-19e278cae35c	1	f
14abb5ca-dab3-46ac-abb9-c6648a380762	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-21 00:20:33.462295-05	c0c5c4ee-5eb8-41da-8d4f-091905a890e9	Module for new noob ship for player	d79d2bf7-b7a5-47ed-a27e-19e278cae35c	1	f
a2f36b0f-2ee2-4bc6-bff1-f388560134cf	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-21 00:20:33.465353-05	c0c5c4ee-5eb8-41da-8d4f-091905a890e9	Module for new noob ship for player	d79d2bf7-b7a5-47ed-a27e-19e278cae35c	1	f
3e514085-1058-4797-9b3b-3c90092a84c4	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-21 00:20:33.474295-05	c0c5c4ee-5eb8-41da-8d4f-091905a890e9	Module for new noob ship for player	d79d2bf7-b7a5-47ed-a27e-19e278cae35c	1	f
2addd0be-b022-4df7-9bcb-2911b51f4001	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-21 00:21:21.15612-05	540cfb79-19aa-4d1c-a71d-e08c1220045c	Mined ore	ab59dcf4-867b-49eb-9cec-ec576b3343ab	62	t
44154ad2-3ead-434f-9b38-02d8813795b4	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-21 00:21:37.878012-05	a0fd0657-776b-424a-afe7-db6633f7e283	Mined ore	93bca1ac-0843-4b7d-85c8-7853e1fa1322	56	t
9c60b0c0-e173-432a-a3c7-3055b88cfdc7	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	{"hp": 1, "volume": 1}	2021-01-21 00:21:33.206142-05	c0c5c4ee-5eb8-41da-8d4f-091905a890e9	Mined ore	94ef2a5b-3e37-46b8-bde0-5b6fecbf8abf	58	t
\.


--
-- TOC entry 3004 (class 0 OID 50897)
-- Dependencies: 205
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
-- TOC entry 3005 (class 0 OID 50903)
-- Dependencies: 206
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
b0b4310a-45eb-4b7b-805a-32d8edea365e	540cfb79-19aa-4d1c-a71d-e08c1220045c
1168b868-9f09-4d10-ba16-bfe2b7802a9c	a0fd0657-776b-424a-afe7-db6633f7e283
52b02fba-6bb7-493d-a982-f04e054682ce	c0c5c4ee-5eb8-41da-8d4f-091905a890e9
\.


--
-- TOC entry 3006 (class 0 OID 50906)
-- Dependencies: 207
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid, fittingbay_containerid, remaxdirty, trash_containerid) FROM stdin;
d072ee04-48a8-4ce0-9d0a-48fd35121df8	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	c0c5c4ee-5eb8-41da-8d4f-091905a890e9	-62440.20029149502	-34326.15701334437	2021-01-21 00:20:33.477323-05	ccc's Starter Ship	Sparrow	345.9129045486018	1.13481963177171	0.06960249163488458	209	244	135	262.8719311022382	350.72544063805526	78.74956249916377	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "be94bc50-6ac8-4a8d-9d85-ab35046d1918", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "3791b5e9-2892-4bc8-9cff-0600f5c967c5", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "14abb5ca-dab3-46ac-abb9-c6648a380762", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "a2f36b0f-2ee2-4bc6-bff1-f388560134cf", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "3e514085-1058-4797-9b3b-3c90092a84c4", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	94ef2a5b-3e37-46b8-bde0-5b6fecbf8abf	d79d2bf7-b7a5-47ed-a27e-19e278cae35c	f	1cb7675d-bcb4-41c1-b651-df462b8e7d1f
7251798a-8c0a-4f76-b4ad-e777928f61e7	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	540cfb79-19aa-4d1c-a71d-e08c1220045c	-62250.119016853176	-33953.501670691134	2021-01-18 22:15:45.402292-05	nwiehoff's Starter Ship	Sparrow	225.08398708157884	-0.6699886453908002	0.998913562487287	209	244	135	262.2355630450278	344.14238139351147	80.9186731024476	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "b9f5aed1-f1cb-456f-a389-d503ca378910", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "cad1206b-997d-4788-9cda-6eb5a866f2db", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "0a3e786e-f2c8-4372-89cf-d07158079c27", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "240d170e-c23e-4aaf-8a97-75d6a2c220ed", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "65cc5524-1b54-40d5-b581-42c0aa220ac7", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	ab59dcf4-867b-49eb-9cec-ec576b3343ab	62d60ec7-2314-4c3a-a1f6-be0e6dee8f44	f	7bf8f0fc-8d62-4178-99d4-c79be5d31dc1
a3709785-5ada-4bfd-a3c6-7bedb704d630	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	a0fd0657-776b-424a-afe7-db6633f7e283	-62882.02253021875	-34266.43640170791	2021-01-20 23:22:39.728087-05	bbb's Starter Ship	Sparrow	56.791630547759496	0.8528300845589802	-2.09566352300149	209	244	135	264.77312102493636	350.69215342341414	79.8099832170426	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "cbfe99cc-9da9-4439-9fbf-a1fe93ea1728", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "79020503-827b-457d-8a62-10bd9e33a3c2", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "030ca987-c825-4a80-9ab1-6b7c893f81ab", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "c7d20552-21c9-47ad-9a0e-8e568280d191", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "a1e7d39f-a623-4ab6-b0ce-fee765591afa", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	93bca1ac-0843-4b7d-85c8-7853e1fa1322	58da71d3-f66e-487f-8879-10c713750044	f	a420c826-33d5-4f1e-b2d4-90a26bb20ee9
\.


--
-- TOC entry 3007 (class 0 OID 50922)
-- Dependencies: 208
-- Data for Name: shiptemplates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume) FROM stdin;
8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	2020-11-23 22:14:30.004993-05	Sparrow	Sparrow	12.5	4.3	100	4.7	209	6	169	135	265	737	11	138	9	e364a553-1dc5-4e8d-9195-0ca4989bec49	{"a_slots": [{"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}]}	120
\.


--
-- TOC entry 3008 (class 0 OID 50945)
-- Dependencies: 209
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
-- TOC entry 3009 (class 0 OID 50948)
-- Dependencies: 210
-- Data for Name: starts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid) FROM stdin;
49f06e89-29fb-4a02-a034-4b5d0702adac	Test Start	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	2020-11-23 15:31:55.475609-05	t	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cf07bba9-90b2-4599-b1e3-84d797a67f0a
\.


--
-- TOC entry 3017 (class 0 OID 51149)
-- Dependencies: 218
-- Data for Name: universe_asteroids; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_asteroids (id, universe_systemid, ore_itemtypeid, name, texture, radius, theta, pos_x, pos_y, yield, mass) FROM stdin;
231ac943-7fca-42db-a8ef-07f35690af3b	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	XYZ-123	Mega\\asteroidR4	220	45.35	-62454	-34091	1.25	10000
\.


--
-- TOC entry 3010 (class 0 OID 50956)
-- Dependencies: 211
-- Data for Name: universe_jumpholes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_jumpholes (id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
0602c3bf-7d70-4a4f-9ebd-77e7cec7ff12	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	edf08406-0879-4141-8af1-f68d32e31c8d	Test Jumphole 1	25000	25000	Jumphole	250	9999999	25
834572ef-b709-4ea7-9cd9-b526744f38cc	edf08406-0879-4141-8af1-f68d32e31c8d	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Jumphole 2	-600000	350000	Jumphole	250	9999999	112
\.


--
-- TOC entry 3011 (class 0 OID 50962)
-- Dependencies: 212
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	207460	-335230	vh_unshaded\\planet03	3960	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	-303350	-100230	vh_unshaded\\planet09	2352	2000000	112
\.


--
-- TOC entry 3012 (class 0 OID 50968)
-- Dependencies: 213
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 3013 (class 0 OID 50971)
-- Dependencies: 214
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	10570	15130	vh_main_sequence\\star_yellow03	28380	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	32670	-11350	vh_main_sequence\\star_red01	15010	1000000000	0
\.


--
-- TOC entry 3014 (class 0 OID 50977)
-- Dependencies: 215
-- Data for Name: universe_stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stations (id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
526f57f5-09e0-41c7-9a89-cd803ec0a065	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Station	207460	-335230	Fleet Armory	810	65000	23.33
cf07bba9-90b2-4599-b1e3-84d797a67f0a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Station	24771.795632863843	-9938.30877953488	Sunfarm	740	25300	112.4
\.


--
-- TOC entry 3015 (class 0 OID 50983)
-- Dependencies: 216
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 3016 (class 0 OID 50986)
-- Dependencies: 217
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid, startid) FROM stdin;
540cfb79-19aa-4d1c-a71d-e08c1220045c	nwiehoff	02c4bf7d7e35c6bab999ac03ece60b8586a27f7ecd4830983b138b74262bf3f9	2021-01-18 22:15:44.780494-05	0	7251798a-8c0a-4f76-b4ad-e777928f61e7	49f06e89-29fb-4a02-a034-4b5d0702adac
c6235bd4-f586-433f-95e9-60020b6f564e	aaa	09b6eea09b48ac41c136ef0515927f468dd413d64406b97660a4b1de14c15d0c	2021-01-20 23:19:55.684703-05	0	\N	49f06e89-29fb-4a02-a034-4b5d0702adac
a0fd0657-776b-424a-afe7-db6633f7e283	bbb	80e6fd623cea60db763983c636dbeb91618681bc2e5217dd13fbd6be536eea8b	2021-01-20 23:22:38.965089-05	0	a3709785-5ada-4bfd-a3c6-7bedb704d630	49f06e89-29fb-4a02-a034-4b5d0702adac
c0c5c4ee-5eb8-41da-8d4f-091905a890e9	ccc	9f8b26b5f16ae88d75289bf5f5d3302fc92caa37d04a0f9e99072471fb5e9562	2021-01-21 00:20:33.430296-05	0	d072ee04-48a8-4ce0-9d0a-48fd35121df8	49f06e89-29fb-4a02-a034-4b5d0702adac
\.


--
-- TOC entry 2810 (class 2606 OID 50990)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2817 (class 2606 OID 50992)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 2819 (class 2606 OID 50994)
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- TOC entry 2821 (class 2606 OID 50996)
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- TOC entry 2823 (class 2606 OID 50998)
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- TOC entry 2827 (class 2606 OID 51000)
-- Name: starts starts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT starts_pkey PRIMARY KEY (id);


--
-- TOC entry 2829 (class 2606 OID 51002)
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- TOC entry 2831 (class 2606 OID 51004)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 2833 (class 2606 OID 51006)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 2835 (class 2606 OID 51008)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2837 (class 2606 OID 51010)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 2839 (class 2606 OID 51012)
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- TOC entry 2841 (class 2606 OID 51014)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 2843 (class 2606 OID 51016)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2850 (class 2606 OID 51153)
-- Name: universe_asteroids uq_asteroid_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT uq_asteroid_id PRIMARY KEY (id);


--
-- TOC entry 2800 (class 2606 OID 51018)
-- Name: containers uq_container_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.containers
    ADD CONSTRAINT uq_container_id PRIMARY KEY (id);


--
-- TOC entry 2802 (class 2606 OID 51020)
-- Name: itemfamilies uq_itemfamily_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemfamilies
    ADD CONSTRAINT uq_itemfamily_id PRIMARY KEY (id);


--
-- TOC entry 2806 (class 2606 OID 51022)
-- Name: items uq_items_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT uq_items_id PRIMARY KEY (id);


--
-- TOC entry 2808 (class 2606 OID 51024)
-- Name: itemtypes uq_itemtype_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT uq_itemtype_id PRIMARY KEY (id);


--
-- TOC entry 2812 (class 2606 OID 51026)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 2846 (class 2606 OID 51028)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2848 (class 2606 OID 51030)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 2803 (class 1259 OID 51031)
-- Name: fki_fk_items_containers; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_containers ON public.items USING btree (containerid);


--
-- TOC entry 2804 (class 1259 OID 51032)
-- Name: fki_fk_items_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_users ON public.items USING btree (createdby);


--
-- TOC entry 2813 (class 1259 OID 51033)
-- Name: fki_fk_ships_containers_cargobay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_cargobay ON public.ships USING btree (cargobay_containerid);


--
-- TOC entry 2814 (class 1259 OID 51034)
-- Name: fki_fk_ships_containers_fittingbay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_fittingbay ON public.ships USING btree (fittingbay_containerid);


--
-- TOC entry 2815 (class 1259 OID 51035)
-- Name: fki_fk_ships_containers_trash; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_trash ON public.ships USING btree (trash_containerid);


--
-- TOC entry 2824 (class 1259 OID 51036)
-- Name: fki_fk_starts_homestations; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_homestations ON public.starts USING btree (homestationid);


--
-- TOC entry 2825 (class 1259 OID 51037)
-- Name: fki_fk_starts_systems; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_systems ON public.starts USING btree (systemid);


--
-- TOC entry 2844 (class 1259 OID 51038)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 2874 (class 2606 OID 51159)
-- Name: universe_asteroids fk_asteroids_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT fk_asteroids_itemtypes FOREIGN KEY (ore_itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 2873 (class 2606 OID 51154)
-- Name: universe_asteroids fk_asteroids_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT fk_asteroids_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2851 (class 2606 OID 51039)
-- Name: items fk_items_containers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_containers FOREIGN KEY (containerid) REFERENCES public.containers(id);


--
-- TOC entry 2852 (class 2606 OID 51044)
-- Name: items fk_items_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 2853 (class 2606 OID 51049)
-- Name: items fk_items_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_users FOREIGN KEY (createdby) REFERENCES public.users(id);


--
-- TOC entry 2854 (class 2606 OID 51054)
-- Name: itemtypes fk_itemtypes_itemfamilies; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT fk_itemtypes_itemfamilies FOREIGN KEY (family) REFERENCES public.itemfamilies(id);


--
-- TOC entry 2855 (class 2606 OID 51059)
-- Name: ships fk_ships_containers_cargobay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_cargobay FOREIGN KEY (cargobay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2856 (class 2606 OID 51064)
-- Name: ships fk_ships_containers_fittingbay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_fittingbay FOREIGN KEY (fittingbay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2857 (class 2606 OID 51069)
-- Name: ships fk_ships_containers_trash; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_trash FOREIGN KEY (trash_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2858 (class 2606 OID 51074)
-- Name: ships fk_ships_dockstations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_dockstations FOREIGN KEY (dockedat_stationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2859 (class 2606 OID 51079)
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2860 (class 2606 OID 51084)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2861 (class 2606 OID 51089)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 2862 (class 2606 OID 51094)
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id);


--
-- TOC entry 2863 (class 2606 OID 51099)
-- Name: starts fk_starts_homestations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_homestations FOREIGN KEY (homestationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2864 (class 2606 OID 51104)
-- Name: starts fk_starts_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2865 (class 2606 OID 51109)
-- Name: starts fk_starts_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_systems FOREIGN KEY (systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2871 (class 2606 OID 51114)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 2872 (class 2606 OID 51119)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 2866 (class 2606 OID 51124)
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2867 (class 2606 OID 51129)
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2868 (class 2606 OID 51134)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2869 (class 2606 OID 51139)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2870 (class 2606 OID 51144)
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2021-01-21 00:25:08

--
-- PostgreSQL database dump complete
--

