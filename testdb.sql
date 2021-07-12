--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7 (Ubuntu 12.7-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.7 (Ubuntu 12.7-0ubuntu0.20.04.1)

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
-- Name: containers; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.containers (
    id uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.containers OWNER TO developer;

--
-- Name: factions; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.factions (
    id uuid NOT NULL,
    name character varying(32) NOT NULL,
    description character varying(512) NOT NULL,
    isnpc bit(1) NOT NULL,
    isjoinable bit(1) NOT NULL,
    canholdsov bit(1) NOT NULL,
    isclosed bit(1) NOT NULL,
    meta jsonb NOT NULL,
    ticker character varying(3) DEFAULT '???'::character varying NOT NULL
);


ALTER TABLE public.factions OWNER TO developer;

--
-- Name: itemfamilies; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.itemfamilies (
    id character varying(16) NOT NULL,
    friendlyname character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemfamilies OWNER TO developer;

--
-- Name: items; Type: TABLE; Schema: public; Owner: developer
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


ALTER TABLE public.items OWNER TO developer;

--
-- Name: itemtypes; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.itemtypes (
    id uuid NOT NULL,
    family character varying(16) NOT NULL,
    name character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemtypes OWNER TO developer;

--
-- Name: processes; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.processes (
    id uuid NOT NULL,
    name character varying(32) NOT NULL,
    meta jsonb DEFAULT '{}'::jsonb NOT NULL,
    "time" integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.processes OWNER TO developer;

--
-- Name: COLUMN processes."time"; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON COLUMN public.processes."time" IS 'Manufacturing time in seconds.';


--
-- Name: processinputs; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.processinputs (
    id uuid NOT NULL,
    itemtypeid uuid NOT NULL,
    quantity integer DEFAULT 0 NOT NULL,
    meta jsonb DEFAULT '{}'::jsonb NOT NULL,
    processid uuid NOT NULL
);


ALTER TABLE public.processinputs OWNER TO developer;

--
-- Name: processoutputs; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.processoutputs (
    id uuid NOT NULL,
    itemtypeid uuid NOT NULL,
    quantity integer DEFAULT 0 NOT NULL,
    meta jsonb DEFAULT '{}'::jsonb NOT NULL,
    processid uuid NOT NULL
);


ALTER TABLE public.processoutputs OWNER TO developer;

--
-- Name: sellorders; Type: TABLE; Schema: public; Owner: developer
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


ALTER TABLE public.sellorders OWNER TO developer;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    userid uuid NOT NULL
);


ALTER TABLE public.sessions OWNER TO developer;

--
-- Name: ships; Type: TABLE; Schema: public; Owner: developer
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


ALTER TABLE public.ships OWNER TO developer;

--
-- Name: shiptemplates; Type: TABLE; Schema: public; Owner: developer
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


ALTER TABLE public.shiptemplates OWNER TO developer;

--
-- Name: shiptypes; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.shiptypes (
    id uuid NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.shiptypes OWNER TO developer;

--
-- Name: starts; Type: TABLE; Schema: public; Owner: developer
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
    wallet double precision DEFAULT 0 NOT NULL,
    factionid uuid NOT NULL
);


ALTER TABLE public.starts OWNER TO developer;

--
-- Name: stationprocesses; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.stationprocesses (
    id uuid NOT NULL,
    universe_stationid uuid NOT NULL,
    processid uuid NOT NULL,
    progress integer DEFAULT 0 NOT NULL,
    installed boolean DEFAULT false NOT NULL,
    internalstate jsonb DEFAULT '{}'::jsonb NOT NULL,
    meta jsonb DEFAULT '{}'::jsonb NOT NULL
);


ALTER TABLE public.stationprocesses OWNER TO developer;

--
-- Name: COLUMN stationprocesses.progress; Type: COMMENT; Schema: public; Owner: developer
--

COMMENT ON COLUMN public.stationprocesses.progress IS 'Progress of manufacturing job in seconds.';


--
-- Name: universe_asteroids; Type: TABLE; Schema: public; Owner: developer
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


ALTER TABLE public.universe_asteroids OWNER TO developer;

--
-- Name: universe_jumpholes; Type: TABLE; Schema: public; Owner: developer
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


ALTER TABLE public.universe_jumpholes OWNER TO developer;

--
-- Name: universe_planets; Type: TABLE; Schema: public; Owner: developer
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


ALTER TABLE public.universe_planets OWNER TO developer;

--
-- Name: universe_regions; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.universe_regions (
    id uuid NOT NULL,
    regionname character varying(32) NOT NULL
);


ALTER TABLE public.universe_regions OWNER TO developer;

--
-- Name: universe_stars; Type: TABLE; Schema: public; Owner: developer
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


ALTER TABLE public.universe_stars OWNER TO developer;

--
-- Name: universe_stations; Type: TABLE; Schema: public; Owner: developer
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
    theta double precision DEFAULT 0 NOT NULL,
    factionid uuid NOT NULL
);


ALTER TABLE public.universe_stations OWNER TO developer;

--
-- Name: universe_systems; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.universe_systems (
    id uuid NOT NULL,
    systemname character varying(32) NOT NULL,
    regionid uuid NOT NULL,
    holding_factionid uuid DEFAULT '42b937ad-0000-46e9-9af9-fc7dbf878e6a'::uuid NOT NULL
);


ALTER TABLE public.universe_systems OWNER TO developer;

--
-- Name: users; Type: TABLE; Schema: public; Owner: developer
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    username character varying(16) NOT NULL,
    hashpass character(64) NOT NULL,
    registered timestamp with time zone NOT NULL,
    banned bit(1) NOT NULL,
    current_shipid uuid,
    startid uuid NOT NULL,
    escrow_containerid uuid NOT NULL,
    current_factionid uuid NOT NULL
);


ALTER TABLE public.users OWNER TO developer;

--
-- Data for Name: containers; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.containers (id, meta, created) FROM stdin;
d5dc0b21-e586-4bf9-8ac6-f2a0ea0ca39c	{}	2021-07-11 20:05:33.487824-04
ce37676b-0c76-4b54-96bc-d26e27e96b0a	{}	2021-07-11 20:05:33.498068-04
fde1ef61-3019-4a0c-9ec1-f2c5b5400eef	{}	2021-07-11 20:05:33.500776-04
5786d9a7-425f-4fae-a2d6-4a203c2dd595	{}	2021-07-11 20:05:33.503514-04
\.


--
-- Data for Name: factions; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.factions (id, name, description, isnpc, isjoinable, canholdsov, isclosed, meta, ticker) FROM stdin;
a8a28085-e7b4-48f5-b8cb-1465ccab82a5	Test Starter Faction	Temporary starter faction for use when a player is created.	0	0	0	0	{}	TSF
42b937ad-0000-46e9-9af9-fc7dbf878e6a	Neutral	Not associated with any faction.	0	0	0	1	{}	   
\.


--
-- Data for Name: itemfamilies; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.itemfamilies (id, friendlyname, meta) FROM stdin;
gun_turret	Gun Turret	{}
missile_launcher	Missile Launcher	{}
shield_booster	Shield Booster	{}
fuel_tank	Fuel Tank	{}
armor_plate	Armor Plate	{}
nothing	Empty Space	{}
ore	Ore	{}
widget	Widget	{}
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.items (id, itemtypeid, meta, created, createdby, createdreason, containerid, quantity, ispackaged) FROM stdin;
006d352c-4fd2-4690-94c0-e938b0063b47	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "tracking": 4.2, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "industrialmarket": {"maxprice": 210, "minprice": 125, "silosize": 1000}, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-07-11 20:05:33.507451-04	a49cb4c7-c7d3-4d89-b7b9-74215db417e3	Module for new noob ship for player	5786d9a7-425f-4fae-a2d6-4a203c2dd595	1	f
6ee35bd3-3337-40ab-a8c1-717d90e7f2e7	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "tracking": 4.2, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "industrialmarket": {"maxprice": 210, "minprice": 125, "silosize": 1000}, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-07-11 20:05:33.514271-04	a49cb4c7-c7d3-4d89-b7b9-74215db417e3	Module for new noob ship for player	5786d9a7-425f-4fae-a2d6-4a203c2dd595	1	f
79a2e10d-9a63-4f6a-85de-7652b7d4e56e	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "industrialmarket": {"maxprice": 650, "minprice": 275, "silosize": 1000}, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-07-11 20:05:33.519128-04	a49cb4c7-c7d3-4d89-b7b9-74215db417e3	Module for new noob ship for player	5786d9a7-425f-4fae-a2d6-4a203c2dd595	1	f
b88ac39c-8f43-4761-be64-21d9596c25e1	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30, "industrialmarket": {"maxprice": 25, "minprice": 10, "silosize": 1000}}	2021-07-11 20:05:33.523607-04	a49cb4c7-c7d3-4d89-b7b9-74215db417e3	Module for new noob ship for player	5786d9a7-425f-4fae-a2d6-4a203c2dd595	1	f
d37f8734-99b3-4214-a730-a31d50af6dc4	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75, "industrialmarket": {"maxprice": 65, "minprice": 30, "silosize": 1000}}	2021-07-11 20:05:33.528289-04	a49cb4c7-c7d3-4d89-b7b9-74215db417e3	Module for new noob ship for player	5786d9a7-425f-4fae-a2d6-4a203c2dd595	1	f
\.


--
-- Data for Name: itemtypes; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.itemtypes (id, family, name, meta) FROM stdin;
91ec9901-ea7e-476a-bc65-7da4523dca38	nothing	Nothing	{"volume": 0}
c311df30-c21e-4895-acb0-d8808f99710e	armor_plate	Basic Armor Plate	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75, "industrialmarket": {"maxprice": 65, "minprice": 30, "silosize": 1000}}
b481a521-1b12-4ffa-ac2f-4da015036f7f	fuel_tank	Basic Fuel Tank	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30, "industrialmarket": {"maxprice": 25, "minprice": 10, "silosize": 1000}}
09172710-740c-4d1c-9fc0-43cb62e674e7	shield_booster	Basic Shield Booster	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "industrialmarket": {"maxprice": 650, "minprice": 275, "silosize": 1000}, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}
dd522f03-2f52-4e82-b2f8-d7e0029cb82f	ore	Testite	{"hp": 1, "volume": 1, "industrialmarket": {"maxprice": 3, "minprice": 1, "silosize": 300000000}}
9d1014c5-3422-4a0f-9839-f585269b4b16	gun_turret	Basic Laser Tool	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "tracking": 4.2, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "industrialmarket": {"maxprice": 210, "minprice": 125, "silosize": 1000}, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}
e2d342a9-d3f2-4f60-8dcf-0ed5bece0da3	widget	Widget	{"volume": 1, "industrialmarket": {"maxprice": 15, "minprice": 10, "silosize": 1000}}
\.


--
-- Data for Name: processes; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.processes (id, name, meta, "time") FROM stdin;
0f33d5ce-2b0b-4f80-80eb-506dd803ac4c	Make Widget	{}	30
\.


--
-- Data for Name: processinputs; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.processinputs (id, itemtypeid, quantity, meta, processid) FROM stdin;
c23de2ef-74e3-4510-8b18-b6f5bcd2f4e9	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	1	{}	0f33d5ce-2b0b-4f80-80eb-506dd803ac4c
\.


--
-- Data for Name: processoutputs; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.processoutputs (id, itemtypeid, quantity, meta, processid) FROM stdin;
71d9ccfa-549b-42be-b1ed-7f5a54c05581	e2d342a9-d3f2-4f60-8dcf-0ed5bece0da3	4	{}	0f33d5ce-2b0b-4f80-80eb-506dd803ac4c
\.


--
-- Data for Name: sellorders; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.sellorders (id, universe_stationid, itemid, seller_userid, askprice, created, bought, buyer_userid) FROM stdin;
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.sessions (id, userid) FROM stdin;
5f47e79d-d571-45d8-9dca-c34cea94286e	a49cb4c7-c7d3-4d89-b7b9-74215db417e3
\.


--
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid, fittingbay_containerid, remaxdirty, trash_containerid, wallet) FROM stdin;
046880ff-d59c-4b1f-9aff-e153d05338ac	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	a49cb4c7-c7d3-4d89-b7b9-74215db417e3	24771.795632863843	-9938.30877953488	2021-07-11 20:05:33.531591-04	aaa's Starter Ship	Sparrow	124.93200320539722	0	0	209	244	135	295	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	cf07bba9-90b2-4599-b1e3-84d797a67f0a	{"a_rack": [{"item_id": "006d352c-4fd2-4690-94c0-e938b0063b47", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "6ee35bd3-3337-40ab-a8c1-717d90e7f2e7", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "79a2e10d-9a63-4f6a-85de-7652b7d4e56e", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "b88ac39c-8f43-4761-be64-21d9596c25e1", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "d37f8734-99b3-4214-a730-a31d50af6dc4", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	fde1ef61-3019-4a0c-9ec1-f2c5b5400eef	5786d9a7-425f-4fae-a2d6-4a203c2dd595	f	ce37676b-0c76-4b54-96bc-d26e27e96b0a	10000
\.


--
-- Data for Name: shiptemplates; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume) FROM stdin;
8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	2020-11-23 22:14:30.004993-05	Sparrow	Sparrow	12.5	4.3	100	4.7	209	6	169	135	265	737	11	138	9	e364a553-1dc5-4e8d-9195-0ca4989bec49	{"a_slots": [{"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}]}	120
\.


--
-- Data for Name: shiptypes; Type: TABLE DATA; Schema: public; Owner: developer
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
-- Data for Name: starts; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid, wallet, factionid) FROM stdin;
49f06e89-29fb-4a02-a034-4b5d0702adac	Test Start	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	2020-11-23 15:31:55.475609-05	t	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cf07bba9-90b2-4599-b1e3-84d797a67f0a	10000	a8a28085-e7b4-48f5-b8cb-1465ccab82a5
\.


--
-- Data for Name: stationprocesses; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.stationprocesses (id, universe_stationid, processid, progress, installed, internalstate, meta) FROM stdin;
966aa59b-68d0-4328-8baf-30616900bcbd	cf07bba9-90b2-4599-b1e3-84d797a67f0a	0f33d5ce-2b0b-4f80-80eb-506dd803ac4c	30	t	{"inputs": {"dd522f03-2f52-4e82-b2f8-d7e0029cb82f": {"price": 1, "quantity": 289593623}}, "outputs": {"e2d342a9-d3f2-4f60-8dcf-0ed5bece0da3": {"price": 10, "quantity": 1000}}, "isRunning": true}	{}
\.


--
-- Data for Name: universe_asteroids; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.universe_asteroids (id, universe_systemid, ore_itemtypeid, name, texture, radius, theta, pos_x, pos_y, yield, mass) FROM stdin;
231ac943-7fca-42db-a8ef-07f35690af3b	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	dd522f03-2f52-4e82-b2f8-d7e0029cb82f	XYZ-123	Mega\\asteroidR4	220	45.35	-62454	-34091	1.25	10000
\.


--
-- Data for Name: universe_jumpholes; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.universe_jumpholes (id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
0602c3bf-7d70-4a4f-9ebd-77e7cec7ff12	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	edf08406-0879-4141-8af1-f68d32e31c8d	Test Jumphole 1	25000	25000	Jumphole	250	9999999	25
834572ef-b709-4ea7-9cd9-b526744f38cc	edf08406-0879-4141-8af1-f68d32e31c8d	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Jumphole 2	-600000	350000	Jumphole	250	9999999	112
\.


--
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	207460	-335230	vh_unshaded\\planet03	3960	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	-303350	-100230	vh_unshaded\\planet09	2352	2000000	112
\.


--
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	10570	15130	vh_main_sequence\\star_yellow03	28380	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	32670	-11350	vh_main_sequence\\star_red01	15010	1000000000	0
\.


--
-- Data for Name: universe_stations; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.universe_stations (id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta, factionid) FROM stdin;
cf07bba9-90b2-4599-b1e3-84d797a67f0a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Station	24771.795632863843	-9938.30877953488	Sunfarm	740	25300	112.4	a8a28085-e7b4-48f5-b8cb-1465ccab82a5
526f57f5-09e0-41c7-9a89-cd803ec0a065	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Station	207460	-335230	Fleet Armory	810	65000	23.33	a8a28085-e7b4-48f5-b8cb-1465ccab82a5
\.


--
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.universe_systems (id, systemname, regionid, holding_factionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647	42b937ad-0000-46e9-9af9-fc7dbf878e6a
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647	42b937ad-0000-46e9-9af9-fc7dbf878e6a
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: developer
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid, startid, escrow_containerid, current_factionid) FROM stdin;
a49cb4c7-c7d3-4d89-b7b9-74215db417e3	aaa	09b6eea09b48ac41c136ef0515927f468dd413d64406b97660a4b1de14c15d0c	2021-07-11 20:05:33.491602-04	0	046880ff-d59c-4b1f-9aff-e153d05338ac	49f06e89-29fb-4a02-a034-4b5d0702adac	d5dc0b21-e586-4bf9-8ac6-f2a0ea0ca39c	a8a28085-e7b4-48f5-b8cb-1465ccab82a5
\.


--
-- Name: factions factions_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.factions
    ADD CONSTRAINT factions_pkey PRIMARY KEY (id);


--
-- Name: processes pk_processes_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.processes
    ADD CONSTRAINT pk_processes_uq PRIMARY KEY (id);


--
-- Name: processinputs pk_processinputs_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.processinputs
    ADD CONSTRAINT pk_processinputs_uq PRIMARY KEY (id);


--
-- Name: processoutputs pk_processoutputs_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.processoutputs
    ADD CONSTRAINT pk_processoutputs_uq PRIMARY KEY (id);


--
-- Name: stationprocesses pk_stationprocess_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.stationprocesses
    ADD CONSTRAINT pk_stationprocess_uq PRIMARY KEY (id);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- Name: starts starts_pkey; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT starts_pkey PRIMARY KEY (id);


--
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- Name: universe_asteroids uq_asteroid_id; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT uq_asteroid_id PRIMARY KEY (id);


--
-- Name: containers uq_container_id; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.containers
    ADD CONSTRAINT uq_container_id PRIMARY KEY (id);


--
-- Name: factions uq_factions_name; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.factions
    ADD CONSTRAINT uq_factions_name UNIQUE (name);


--
-- Name: factions uq_factions_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.factions
    ADD CONSTRAINT uq_factions_pk UNIQUE (id);


--
-- Name: factions uq_factions_ticker; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.factions
    ADD CONSTRAINT uq_factions_ticker UNIQUE (ticker);


--
-- Name: itemfamilies uq_itemfamily_id; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.itemfamilies
    ADD CONSTRAINT uq_itemfamily_id PRIMARY KEY (id);


--
-- Name: items uq_items_id; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT uq_items_id PRIMARY KEY (id);


--
-- Name: itemtypes uq_itemtype_id; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT uq_itemtype_id PRIMARY KEY (id);


--
-- Name: sellorders uq_sellorders_pk; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT uq_sellorders_pk PRIMARY KEY (id);


--
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- Name: fki_fk_items_containers; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_items_containers ON public.items USING btree (containerid);


--
-- Name: fki_fk_items_users; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_items_users ON public.items USING btree (createdby);


--
-- Name: fki_fk_ships_containers_cargobay; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_ships_containers_cargobay ON public.ships USING btree (cargobay_containerid);


--
-- Name: fki_fk_ships_containers_fittingbay; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_ships_containers_fittingbay ON public.ships USING btree (fittingbay_containerid);


--
-- Name: fki_fk_ships_containers_trash; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_ships_containers_trash ON public.ships USING btree (trash_containerid);


--
-- Name: fki_fk_starts_homestations; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_starts_homestations ON public.starts USING btree (homestationid);


--
-- Name: fki_fk_starts_systems; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_starts_systems ON public.starts USING btree (systemid);


--
-- Name: fki_fk_users_containers_escrow; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_users_containers_escrow ON public.users USING btree (escrow_containerid);


--
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: developer
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- Name: universe_asteroids fk_asteroids_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT fk_asteroids_itemtypes FOREIGN KEY (ore_itemtypeid) REFERENCES public.itemtypes(id);


--
-- Name: universe_asteroids fk_asteroids_systems; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_asteroids
    ADD CONSTRAINT fk_asteroids_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- Name: items fk_items_containers; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_containers FOREIGN KEY (containerid) REFERENCES public.containers(id);


--
-- Name: items fk_items_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- Name: items fk_items_users; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_users FOREIGN KEY (createdby) REFERENCES public.users(id);


--
-- Name: itemtypes fk_itemtypes_itemfamilies; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT fk_itemtypes_itemfamilies FOREIGN KEY (family) REFERENCES public.itemfamilies(id);


--
-- Name: processinputs fk_processinputs_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.processinputs
    ADD CONSTRAINT fk_processinputs_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- Name: processinputs fk_processinputs_processes; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.processinputs
    ADD CONSTRAINT fk_processinputs_processes FOREIGN KEY (processid) REFERENCES public.processes(id);


--
-- Name: processoutputs fk_processoutputs_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.processoutputs
    ADD CONSTRAINT fk_processoutputs_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- Name: processoutputs fk_processoutputs_processes; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.processoutputs
    ADD CONSTRAINT fk_processoutputs_processes FOREIGN KEY (processid) REFERENCES public.processes(id);


--
-- Name: sellorders fk_sellorders_items; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT fk_sellorders_items FOREIGN KEY (itemid) REFERENCES public.items(id);


--
-- Name: sellorders fk_sellorders_stations; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT fk_sellorders_stations FOREIGN KEY (universe_stationid) REFERENCES public.universe_stations(id);


--
-- Name: sellorders fk_sellorders_users_buyers; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT fk_sellorders_users_buyers FOREIGN KEY (buyer_userid) REFERENCES public.users(id);


--
-- Name: sellorders fk_sellorders_users_sellers; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.sellorders
    ADD CONSTRAINT fk_sellorders_users_sellers FOREIGN KEY (seller_userid) REFERENCES public.users(id);


--
-- Name: ships fk_ships_containers_cargobay; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_cargobay FOREIGN KEY (cargobay_containerid) REFERENCES public.containers(id);


--
-- Name: ships fk_ships_containers_fittingbay; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_fittingbay FOREIGN KEY (fittingbay_containerid) REFERENCES public.containers(id);


--
-- Name: ships fk_ships_containers_trash; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_trash FOREIGN KEY (trash_containerid) REFERENCES public.containers(id);


--
-- Name: ships fk_ships_dockstations; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_dockstations FOREIGN KEY (dockedat_stationid) REFERENCES public.universe_stations(id);


--
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id);


--
-- Name: starts fk_starts_factions; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_factions FOREIGN KEY (factionid) REFERENCES public.factions(id);


--
-- Name: starts fk_starts_homestations; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_homestations FOREIGN KEY (homestationid) REFERENCES public.universe_stations(id);


--
-- Name: starts fk_starts_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- Name: starts fk_starts_systems; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_systems FOREIGN KEY (systemid) REFERENCES public.universe_systems(id);


--
-- Name: stationprocesses fk_stationprocess_process; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.stationprocesses
    ADD CONSTRAINT fk_stationprocess_process FOREIGN KEY (processid) REFERENCES public.processes(id);


--
-- Name: universe_systems fk_system_faction; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_faction FOREIGN KEY (holding_factionid) REFERENCES public.factions(id);


--
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- Name: stationprocesses fk_universe_station_process; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.stationprocesses
    ADD CONSTRAINT fk_universe_station_process FOREIGN KEY (universe_stationid) REFERENCES public.universe_stations(id);


--
-- Name: users fk_users_containers_escrow; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_containers_escrow FOREIGN KEY (escrow_containerid) REFERENCES public.containers(id);


--
-- Name: users fk_users_factions; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_factions FOREIGN KEY (current_factionid) REFERENCES public.factions(id);


--
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- Name: universe_stations station_faction_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_faction_fk FOREIGN KEY (factionid) REFERENCES public.factions(id);


--
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: developer
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- PostgreSQL database dump complete
--

