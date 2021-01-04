--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2021-01-03 21:11:10

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
-- TOC entry 202 (class 1259 OID 50603)
-- Name: containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.containers (
    id uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.containers OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 50609)
-- Name: itemfamilies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.itemfamilies (
    id character varying(16) NOT NULL,
    friendlyname character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemfamilies OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 50615)
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
-- TOC entry 205 (class 1259 OID 50621)
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
-- TOC entry 206 (class 1259 OID 50627)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    userid uuid NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 50630)
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
-- TOC entry 208 (class 1259 OID 50646)
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
-- TOC entry 209 (class 1259 OID 50669)
-- Name: shiptypes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptypes (
    id uuid NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.shiptypes OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 50672)
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
-- TOC entry 211 (class 1259 OID 50680)
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
-- TOC entry 212 (class 1259 OID 50686)
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
-- TOC entry 213 (class 1259 OID 50692)
-- Name: universe_regions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_regions (
    id uuid NOT NULL,
    regionname character varying(32) NOT NULL
);


ALTER TABLE public.universe_regions OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 50695)
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
-- TOC entry 215 (class 1259 OID 50701)
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
-- TOC entry 216 (class 1259 OID 50707)
-- Name: universe_systems; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_systems (
    id uuid NOT NULL,
    systemname character varying(32) NOT NULL,
    regionid uuid NOT NULL
);


ALTER TABLE public.universe_systems OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 50710)
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
-- TOC entry 2993 (class 0 OID 50603)
-- Dependencies: 202
-- Data for Name: containers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.containers (id, meta, created) FROM stdin;
aa22ba1f-80e2-4d15-a83d-a0d60c03543c	{}	2021-01-03 21:01:00.664807-05
5a4d2ed1-f597-4713-b251-61396dc6e61f	{}	2021-01-03 21:01:00.666807-05
f01c22f8-1c65-418d-ac6a-3ae2490b12e4	{}	2021-01-03 21:01:00.73886-05
6aa7c441-e38b-4d26-bb4d-7c23b1bea2a4	{}	2021-01-03 21:06:37.009478-05
1d6833a4-2076-4a35-832c-5880d1925948	{}	2021-01-03 21:06:37.01249-05
0b23a6bf-1078-49ff-bd26-cb87c717178f	{}	2021-01-03 21:06:37.086546-05
\.


--
-- TOC entry 2994 (class 0 OID 50609)
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
\.


--
-- TOC entry 2995 (class 0 OID 50615)
-- Dependencies: 204
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, itemtypeid, meta, created, createdby, createdreason, containerid, quantity, ispackaged) FROM stdin;
a2076201-4d22-437d-88dd-e1b631be74ba	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-03 21:06:37.387738-05	7635aca2-3536-4727-a1a5-2b6255982ec2	Module for new noob ship for player	1d6833a4-2076-4a35-832c-5880d1925948	1	f
9cf91783-4b98-4665-b1b0-d0b83ed58f9e	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2021-01-03 21:06:37.469964-05	7635aca2-3536-4727-a1a5-2b6255982ec2	Module for new noob ship for player	1d6833a4-2076-4a35-832c-5880d1925948	1	t
1c5e7d03-5a64-4001-b871-ccef274c4359	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-03 21:01:01.050813-05	c27609bb-8eaf-4053-a4e1-6e47a76cb8e4	Module for new noob ship for player	f01c22f8-1c65-418d-ac6a-3ae2490b12e4	1	f
a8c4429b-3c3e-4837-89ff-f36e8c5b1524	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-03 21:01:01.210875-05	c27609bb-8eaf-4053-a4e1-6e47a76cb8e4	Module for new noob ship for player	f01c22f8-1c65-418d-ac6a-3ae2490b12e4	1	f
4ccad356-93e8-43f8-aa04-71bccfb5b6e7	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-03 21:01:01.130875-05	c27609bb-8eaf-4053-a4e1-6e47a76cb8e4	Module for new noob ship for player	5a4d2ed1-f597-4713-b251-61396dc6e61f	1	f
fe104590-0e2c-458c-bc5e-94e46a722e46	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2021-01-03 21:01:00.885875-05	c27609bb-8eaf-4053-a4e1-6e47a76cb8e4	Module for new noob ship for player	5a4d2ed1-f597-4713-b251-61396dc6e61f	0	t
3cd14dd0-9ed5-445d-a538-69861615efae	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2021-01-03 21:01:00.971807-05	c27609bb-8eaf-4053-a4e1-6e47a76cb8e4	Module for new noob ship for player	5a4d2ed1-f597-4713-b251-61396dc6e61f	2	t
c3242572-4dad-42e1-8376-4bbb16142f8f	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2021-01-03 21:06:37.23148-05	7635aca2-3536-4727-a1a5-2b6255982ec2	Module for new noob ship for player	1d6833a4-2076-4a35-832c-5880d1925948	2	t
8b65e978-211f-473c-8c14-64c64e9a4dbe	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-03 21:06:37.551965-05	7635aca2-3536-4727-a1a5-2b6255982ec2	Module for new noob ship for player	0b23a6bf-1078-49ff-bd26-cb87c717178f	1	f
04328a44-4425-491d-8154-5fa5b36b62c8	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2021-01-03 21:06:37.309643-05	7635aca2-3536-4727-a1a5-2b6255982ec2	Module for new noob ship for player	1d6833a4-2076-4a35-832c-5880d1925948	0	t
\.


--
-- TOC entry 2996 (class 0 OID 50621)
-- Dependencies: 205
-- Data for Name: itemtypes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.itemtypes (id, family, name, meta) FROM stdin;
9d1014c5-3422-4a0f-9839-f585269b4b16	gun_turret	Basic Laser Tool	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}
b481a521-1b12-4ffa-ac2f-4da015036f7f	fuel_tank	Basic Fuel Tank	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}
09172710-740c-4d1c-9fc0-43cb62e674e7	shield_booster	Basic Shield Booster	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}
c311df30-c21e-4895-acb0-d8808f99710e	armor_plate	Basic Armor Plate	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}
91ec9901-ea7e-476a-bc65-7da4523dca38	nothing	Nothing	{"volume": 0}
\.


--
-- TOC entry 2997 (class 0 OID 50627)
-- Dependencies: 206
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
2439b10d-6440-4c20-8a91-d592842c8322	c27609bb-8eaf-4053-a4e1-6e47a76cb8e4
67b80662-a240-421a-82af-f719ee43189d	7635aca2-3536-4727-a1a5-2b6255982ec2
\.


--
-- TOC entry 2998 (class 0 OID 50630)
-- Dependencies: 207
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid, fittingbay_containerid, remaxdirty, trash_containerid) FROM stdin;
e928d764-c270-424b-8969-939186074344	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	7635aca2-3536-4727-a1a5-2b6255982ec2	24771.795632863843	-9938.30877953488	2021-01-03 21:06:37.558964-05	asdf's Starter Ship	Sparrow	207.57863264297916	0	0	209	244	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	cf07bba9-90b2-4599-b1e3-84d797a67f0a	{"a_rack": null, "b_rack": null, "c_rack": [{"item_id": "8b65e978-211f-473c-8c14-64c64e9a4dbe", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	1d6833a4-2076-4a35-832c-5880d1925948	0b23a6bf-1078-49ff-bd26-cb87c717178f	f	6aa7c441-e38b-4d26-bb4d-7c23b1bea2a4
0bebd4f2-4e91-4352-ae78-7472702c6bdc	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	c27609bb-8eaf-4053-a4e1-6e47a76cb8e4	48237.077569892834	-22194.52128189182	2021-01-03 21:01:01.216874-05	nwiehoff's Starter Ship	Sparrow	0	9.424037815194766e-90	-4.922293727526296e-90	209	244	135	295	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "fe104590-0e2c-458c-bc5e-94e46a722e46", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "3cd14dd0-9ed5-445d-a538-69861615efae", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "1c5e7d03-5a64-4001-b871-ccef274c4359", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_id": "4ccad356-93e8-43f8-aa04-71bccfb5b6e7", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "a8c4429b-3c3e-4837-89ff-f36e8c5b1524", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	5a4d2ed1-f597-4713-b251-61396dc6e61f	f01c22f8-1c65-418d-ac6a-3ae2490b12e4	f	aa22ba1f-80e2-4d15-a83d-a0d60c03543c
\.


--
-- TOC entry 2999 (class 0 OID 50646)
-- Dependencies: 208
-- Data for Name: shiptemplates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume) FROM stdin;
8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	2020-11-23 22:14:30.004993-05	Sparrow	Sparrow	12.5	4.3	100	4.7	209	6	169	135	265	737	11	138	9	e364a553-1dc5-4e8d-9195-0ca4989bec49	{"a_slots": [{"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}]}	120
\.


--
-- TOC entry 3000 (class 0 OID 50669)
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
-- TOC entry 3001 (class 0 OID 50672)
-- Dependencies: 210
-- Data for Name: starts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid) FROM stdin;
49f06e89-29fb-4a02-a034-4b5d0702adac	Test Start	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	2020-11-23 15:31:55.475609-05	t	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cf07bba9-90b2-4599-b1e3-84d797a67f0a
\.


--
-- TOC entry 3002 (class 0 OID 50680)
-- Dependencies: 211
-- Data for Name: universe_jumpholes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_jumpholes (id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
0602c3bf-7d70-4a4f-9ebd-77e7cec7ff12	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	edf08406-0879-4141-8af1-f68d32e31c8d	Test Jumphole 1	25000	25000	Jumphole	250	9999999	25
834572ef-b709-4ea7-9cd9-b526744f38cc	edf08406-0879-4141-8af1-f68d32e31c8d	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Jumphole 2	-600000	350000	Jumphole	250	9999999	112
\.


--
-- TOC entry 3003 (class 0 OID 50686)
-- Dependencies: 212
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	207460	-335230	vh_unshaded\\planet03	3960	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	-303350	-100230	vh_unshaded\\planet09	2352	2000000	112
\.


--
-- TOC entry 3004 (class 0 OID 50692)
-- Dependencies: 213
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 3005 (class 0 OID 50695)
-- Dependencies: 214
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	10570	15130	vh_main_sequence\\star_yellow03	28380	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	32670	-11350	vh_main_sequence\\star_red01	15010	1000000000	0
\.


--
-- TOC entry 3006 (class 0 OID 50701)
-- Dependencies: 215
-- Data for Name: universe_stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stations (id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
cf07bba9-90b2-4599-b1e3-84d797a67f0a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Station	24771.795632863843	-9938.30877953488	Sunfarm	740	25300	112.4
526f57f5-09e0-41c7-9a89-cd803ec0a065	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Station	207460	-335230	Fleet Armory	810	65000	23.33
\.


--
-- TOC entry 3007 (class 0 OID 50707)
-- Dependencies: 216
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 3008 (class 0 OID 50710)
-- Dependencies: 217
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid, startid) FROM stdin;
c27609bb-8eaf-4053-a4e1-6e47a76cb8e4	nwiehoff	02c4bf7d7e35c6bab999ac03ece60b8586a27f7ecd4830983b138b74262bf3f9	2021-01-03 21:01:00.577808-05	0	0bebd4f2-4e91-4352-ae78-7472702c6bdc	49f06e89-29fb-4a02-a034-4b5d0702adac
7635aca2-3536-4727-a1a5-2b6255982ec2	asdf	7ab518aea3a72084bf1f1e1ddeb7fa25563a47b73b643452960187b6405b72f1	2021-01-03 21:06:36.925479-05	0	e928d764-c270-424b-8969-939186074344	49f06e89-29fb-4a02-a034-4b5d0702adac
\.


--
-- TOC entry 2806 (class 2606 OID 50714)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2813 (class 2606 OID 50716)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 2815 (class 2606 OID 50718)
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- TOC entry 2817 (class 2606 OID 50720)
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- TOC entry 2819 (class 2606 OID 50722)
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- TOC entry 2823 (class 2606 OID 50724)
-- Name: starts starts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT starts_pkey PRIMARY KEY (id);


--
-- TOC entry 2825 (class 2606 OID 50726)
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- TOC entry 2827 (class 2606 OID 50728)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 2829 (class 2606 OID 50730)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 2831 (class 2606 OID 50732)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2833 (class 2606 OID 50734)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 2835 (class 2606 OID 50736)
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- TOC entry 2837 (class 2606 OID 50738)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 2839 (class 2606 OID 50740)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2796 (class 2606 OID 50742)
-- Name: containers uq_container_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.containers
    ADD CONSTRAINT uq_container_id PRIMARY KEY (id);


--
-- TOC entry 2798 (class 2606 OID 50744)
-- Name: itemfamilies uq_itemfamily_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemfamilies
    ADD CONSTRAINT uq_itemfamily_id PRIMARY KEY (id);


--
-- TOC entry 2802 (class 2606 OID 50746)
-- Name: items uq_items_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT uq_items_id PRIMARY KEY (id);


--
-- TOC entry 2804 (class 2606 OID 50748)
-- Name: itemtypes uq_itemtype_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT uq_itemtype_id PRIMARY KEY (id);


--
-- TOC entry 2808 (class 2606 OID 50750)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 2842 (class 2606 OID 50752)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2844 (class 2606 OID 50754)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 2799 (class 1259 OID 50755)
-- Name: fki_fk_items_containers; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_containers ON public.items USING btree (containerid);


--
-- TOC entry 2800 (class 1259 OID 50756)
-- Name: fki_fk_items_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_users ON public.items USING btree (createdby);


--
-- TOC entry 2809 (class 1259 OID 50757)
-- Name: fki_fk_ships_containers_cargobay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_cargobay ON public.ships USING btree (cargobay_containerid);


--
-- TOC entry 2810 (class 1259 OID 50758)
-- Name: fki_fk_ships_containers_fittingbay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_fittingbay ON public.ships USING btree (fittingbay_containerid);


--
-- TOC entry 2811 (class 1259 OID 50874)
-- Name: fki_fk_ships_containers_trash; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_trash ON public.ships USING btree (trash_containerid);


--
-- TOC entry 2820 (class 1259 OID 50759)
-- Name: fki_fk_starts_homestations; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_homestations ON public.starts USING btree (homestationid);


--
-- TOC entry 2821 (class 1259 OID 50760)
-- Name: fki_fk_starts_systems; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_systems ON public.starts USING btree (systemid);


--
-- TOC entry 2840 (class 1259 OID 50761)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 2845 (class 2606 OID 50762)
-- Name: items fk_items_containers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_containers FOREIGN KEY (containerid) REFERENCES public.containers(id);


--
-- TOC entry 2846 (class 2606 OID 50767)
-- Name: items fk_items_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 2847 (class 2606 OID 50772)
-- Name: items fk_items_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_users FOREIGN KEY (createdby) REFERENCES public.users(id);


--
-- TOC entry 2848 (class 2606 OID 50777)
-- Name: itemtypes fk_itemtypes_itemfamilies; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT fk_itemtypes_itemfamilies FOREIGN KEY (family) REFERENCES public.itemfamilies(id);


--
-- TOC entry 2849 (class 2606 OID 50782)
-- Name: ships fk_ships_containers_cargobay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_cargobay FOREIGN KEY (cargobay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2850 (class 2606 OID 50787)
-- Name: ships fk_ships_containers_fittingbay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_fittingbay FOREIGN KEY (fittingbay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2855 (class 2606 OID 50869)
-- Name: ships fk_ships_containers_trash; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_trash FOREIGN KEY (trash_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2851 (class 2606 OID 50792)
-- Name: ships fk_ships_dockstations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_dockstations FOREIGN KEY (dockedat_stationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2852 (class 2606 OID 50797)
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2853 (class 2606 OID 50802)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2854 (class 2606 OID 50807)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 2856 (class 2606 OID 50812)
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id);


--
-- TOC entry 2857 (class 2606 OID 50817)
-- Name: starts fk_starts_homestations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_homestations FOREIGN KEY (homestationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2858 (class 2606 OID 50822)
-- Name: starts fk_starts_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2859 (class 2606 OID 50827)
-- Name: starts fk_starts_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_systems FOREIGN KEY (systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2865 (class 2606 OID 50832)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 2866 (class 2606 OID 50837)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 2860 (class 2606 OID 50842)
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2861 (class 2606 OID 50847)
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2862 (class 2606 OID 50852)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2863 (class 2606 OID 50857)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2864 (class 2606 OID 50862)
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2021-01-03 21:11:11

--
-- PostgreSQL database dump complete
--

