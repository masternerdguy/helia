--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2020-11-23 23:09:04

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
-- TOC entry 202 (class 1259 OID 49625)
-- Name: itemfamilies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.itemfamilies (
    id character varying(16) NOT NULL,
    friendlyname character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemfamilies OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 49631)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id uuid NOT NULL,
    itemtypeid uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL,
    createdby uuid,
    createdreason character varying(64) NOT NULL
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 49637)
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
-- TOC entry 205 (class 1259 OID 49643)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    userid uuid NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 49646)
-- Name: ships; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ships (
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    userid uuid NOT NULL,
    pos_x double precision DEFAULT 0 NOT NULL,
    pos_y double precision DEFAULT 0 NOT NULL,
    created time with time zone DEFAULT now() NOT NULL,
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
    destroyedat timestamp with time zone
);


ALTER TABLE public.ships OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 49660)
-- Name: shiptemplates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptemplates (
    id uuid NOT NULL,
    created time with time zone DEFAULT now() NOT NULL,
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
    slotlayout jsonb DEFAULT '{}'::jsonb NOT NULL
);


ALTER TABLE public.shiptemplates OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 49682)
-- Name: shiptypes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptypes (
    id uuid NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.shiptypes OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 49828)
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
-- TOC entry 209 (class 1259 OID 49685)
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
-- TOC entry 210 (class 1259 OID 49691)
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
-- TOC entry 211 (class 1259 OID 49697)
-- Name: universe_regions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_regions (
    id uuid NOT NULL,
    regionname character varying(32) NOT NULL
);


ALTER TABLE public.universe_regions OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 49700)
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
-- TOC entry 213 (class 1259 OID 49706)
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
-- TOC entry 214 (class 1259 OID 49712)
-- Name: universe_systems; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_systems (
    id uuid NOT NULL,
    systemname character varying(32) NOT NULL,
    regionid uuid NOT NULL
);


ALTER TABLE public.universe_systems OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 49715)
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
-- TOC entry 2974 (class 0 OID 49625)
-- Dependencies: 202
-- Data for Name: itemfamilies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.itemfamilies (id, friendlyname, meta) FROM stdin;
gun_turret	Gun Turret	{}
missile_launcher	Missile Launcher	{}
\.


--
-- TOC entry 2975 (class 0 OID 49631)
-- Dependencies: 203
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, itemtypeid, meta, created, createdby, createdreason) FROM stdin;
297d14b2-e6b5-4301-ba79-222546a394a0	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 16:03:29.743038-05	7168a9e6-09dc-44f8-9943-bd1231df7b9f	Module for starter ship of new player
93e4f414-c917-4333-8b97-02f0d8985c45	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 16:03:29.750024-05	7168a9e6-09dc-44f8-9943-bd1231df7b9f	Module for starter ship of new player
b0c0566f-b85d-4d24-a295-6cdc305be833	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 16:05:14.65662-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for starter ship of new player
42bc15d0-1891-4898-bd18-42d8ff04600e	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 16:05:14.66362-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for starter ship of new player
98d8c6fa-e83c-4cb5-97e0-179e8be27c6e	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 18:38:07.63895-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for new noob ship for player
8a18ab6d-63ce-401c-ac40-f3ead327c089	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 18:38:07.641982-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for new noob ship for player
d26843f5-6c94-4c54-8eca-bf96ab790d51	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:08:06.921454-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for new noob ship for player
c63f5c44-c55a-4729-a240-52d2f8c02d73	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:08:06.925453-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for new noob ship for player
7283c2de-d72a-49db-bb75-0a13e17da259	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:14:30.033992-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for new noob ship for player
a430b5f1-faae-4d06-afa6-0d21069c32c1	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:14:30.039992-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for new noob ship for player
88f7e2a5-f8f4-4a57-af65-a4ed6d58fa75	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:21:47.30866-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for new noob ship for player
0a21e964-7c7d-49da-80f8-cefcd0a68e06	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:21:47.311659-05	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	Module for new noob ship for player
6976fd79-4655-47eb-a1f3-bf09c8ae3441	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:40:16.611039-05	e92c5611-f424-4f2f-abf9-35f727a2bb21	Module for new noob ship for player
81cd155e-1431-47d2-8c67-b4106dd2c21a	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:40:16.615039-05	e92c5611-f424-4f2f-abf9-35f727a2bb21	Module for new noob ship for player
9935e11b-715b-4ae8-86e6-3efc188ed35e	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:46:38.198542-05	7168a9e6-09dc-44f8-9943-bd1231df7b9f	Module for new noob ship for player
7698fec9-2d24-476e-9898-15fa5a01dfd7	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 22:46:38.203542-05	7168a9e6-09dc-44f8-9943-bd1231df7b9f	Module for new noob ship for player
5f12240a-dc3b-4138-b6b7-47e203cf7fd8	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 23:04:32.05429-05	e92c5611-f424-4f2f-abf9-35f727a2bb21	Module for new noob ship for player
12d11742-dc0a-4568-921d-efe3851ce807	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-11-23 23:04:32.059261-05	e92c5611-f424-4f2f-abf9-35f727a2bb21	Module for new noob ship for player
\.


--
-- TOC entry 2976 (class 0 OID 49637)
-- Dependencies: 204
-- Data for Name: itemtypes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.itemtypes (id, family, name, meta) FROM stdin;
9d1014c5-3422-4a0f-9839-f585269b4b16	gun_turret	Basic Laser Tool	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}
\.


--
-- TOC entry 2977 (class 0 OID 49643)
-- Dependencies: 205
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
ad31d465-b7db-4ace-b2bf-3bb3fbe0ddfb	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9
a4548449-1e3b-496a-bf56-ee2562c8a40d	e92c5611-f424-4f2f-abf9-35f727a2bb21
c5244653-ca2e-4822-878a-e859e2e3795d	7168a9e6-09dc-44f8-9943-bd1231df7b9f
\.


--
-- TOC entry 2978 (class 0 OID 49646)
-- Dependencies: 206
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat) FROM stdin;
682b6f69-454b-4c59-91cd-b28efcc03b31	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	207988.4762577953	-334644.17859510664	22:21:47.378659-05	asdf's Starter Ship	Sparrow	211.13184408045646	-2.858842583109887	2.592006108276746	0	3.0761647867025217	200	250.06610998232955	0.006914535107179481	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "88f7e2a5-f8f4-4a57-af65-a4ed6d58fa75", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "0a21e964-7c7d-49da-80f8-cefcd0a68e06", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	f	\N
2b653c97-e671-419c-bdc4-ac51d3d29561	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	e92c5611-f424-4f2f-abf9-35f727a2bb21	24771.795632863843	-9938.30877953488	23:04:32.180262-05	aaaa's Starter Ship	Sparrow	230.26468458476336	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	cf07bba9-90b2-4599-b1e3-84d797a67f0a	{"a_rack": [{"item_id": "5f12240a-dc3b-4138-b6b7-47e203cf7fd8", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "12d11742-dc0a-4568-921d-efe3851ce807", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	f	\N
9fd579ce-a902-4b2d-aaa9-92e192e65b92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	7168a9e6-09dc-44f8-9943-bd1231df7b9f	25402.520390005866	22969.6735888212	16:03:29.80897-05	nwiehoff's Starter Ship	Sparrow	6.71281429025629	6.5850663939902895e-167	-3.1958686874314914e-167	0	3.9054829655719496	0	-2.265560781579748e-06	0	0.7400000000002933	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "297d14b2-e6b5-4301-ba79-222546a394a0", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "93e4f414-c917-4333-8b97-02f0d8985c45", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	t	2020-11-23 22:46:38.18657-05
f0b1bcfb-3dd9-48b6-a18f-3c6388da4691	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	23928.54514305847	-10754.15744184927	16:05:14.729621-05	asdf's Starter Ship	Sparrow	242.88276782449734	-2.0305346126718823e-97	6.974949556860364e-97	0	0	0	252.3329006044156	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "b0c0566f-b85d-4d24-a295-6cdc305be833", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "42bc15d0-1891-4898-bd18-42d8ff04600e", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	t	2020-11-23 18:52:09.777582-05
6aaf3890-450a-46d9-9951-3792bf4e2d1b	edf08406-0879-4141-8af1-f68d32e31c8d	7168a9e6-09dc-44f8-9943-bd1231df7b9f	-593105.1401454965	351417.1020575273	22:46:38.280558-05	nwiehoff's Starter Ship	Sparrow	163.17859010995917	-1.2879685114642694e-23	-2.064661009830277e-24	0	62.97648906617315	110.16612072264648	103.17770818990802	126.16090678434998	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "9935e11b-715b-4ae8-86e6-3efc188ed35e", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "7698fec9-2d24-476e-9898-15fa5a01dfd7", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	f	\N
07870a03-913b-4430-a42a-cd1d7a94fb10	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	23926.20860334656	-10745.975049228044	22:08:07.003453-05	asdf's Starter Ship	Sparrow	136.313965978422	-1.1215133471123992e-28	-1.071218537955603e-28	0	3.137702061778015	0	260.23998597682646	61.30000000000106	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "d26843f5-6c94-4c54-8eca-bf96ab790d51", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "c63f5c44-c55a-4729-a240-52d2f8c02d73", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	t	2020-11-23 22:14:30.004993-05
0f030b14-7239-4e54-82b7-271e079a4e9a	edf08406-0879-4141-8af1-f68d32e31c8d	e92c5611-f424-4f2f-abf9-35f727a2bb21	-591263.2330057462	349756.282160357	22:40:16.677043-05	aaaa's Starter Ship	Sparrow	137.90675796336305	-1.8287761883474274	-0.6928506356017361	0	2.762271030879232	0	108.41960552242853	359.58203542434995	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "6976fd79-4655-47eb-a1f3-bf09c8ae3441", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "81cd155e-1431-47d2-8c67-b4106dd2c21a", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	t	2020-11-23 23:04:32.043548-05
05171d64-bf9d-4612-84f6-0afa534b2630	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	23927.592214583863	-10776.780874378657	18:38:07.715928-05	asdf's Starter Ship	Sparrow	203.48225691739913	-1.8826987743846525e-219	8.179275792709127e-220	0	0	0	262.5595870856634	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "98d8c6fa-e83c-4cb5-97e0-179e8be27c6e", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "8a18ab6d-63ce-401c-ac40-f3ead327c089", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	t	2020-11-23 22:09:33.022188-05
e39dff9d-5686-4011-b7fd-0d5c581279b2	edf08406-0879-4141-8af1-f68d32e31c8d	664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	-598411.9958169686	351865.8952203155	22:14:30.120993-05	asdf's Starter Ship	Sparrow	329.82647997035565	3.728475306676788e-25	2.183566104093122e-25	0	0	0	262.0334485065684	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "7283c2de-d72a-49db-bb75-0a13e17da259", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "a430b5f1-faae-4d06-afa6-0d21069c32c1", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	t	2020-11-23 22:21:47.296712-05
\.


--
-- TOC entry 2979 (class 0 OID 49660)
-- Dependencies: 207
-- Data for Name: shiptemplates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout) FROM stdin;
8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	12:34:26.294578-04	Sparrow	Sparrow	12.5	4.3	100	4.7	209	6	169	135	265	737	11	138	9	e364a553-1dc5-4e8d-9195-0ca4989bec49	{"a_slots": [{"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}]}
\.


--
-- TOC entry 2980 (class 0 OID 49682)
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
-- TOC entry 2988 (class 0 OID 49828)
-- Dependencies: 216
-- Data for Name: starts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid) FROM stdin;
49f06e89-29fb-4a02-a034-4b5d0702adac	Test Start	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": null}	2020-11-23 15:31:55.475609-05	t	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cf07bba9-90b2-4599-b1e3-84d797a67f0a
\.


--
-- TOC entry 2981 (class 0 OID 49685)
-- Dependencies: 209
-- Data for Name: universe_jumpholes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_jumpholes (id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
0602c3bf-7d70-4a4f-9ebd-77e7cec7ff12	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	edf08406-0879-4141-8af1-f68d32e31c8d	Test Jumphole 1	25000	25000	Jumphole	250	9999999	25
834572ef-b709-4ea7-9cd9-b526744f38cc	edf08406-0879-4141-8af1-f68d32e31c8d	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Jumphole 2	-600000	350000	Jumphole	250	9999999	112
\.


--
-- TOC entry 2982 (class 0 OID 49691)
-- Dependencies: 210
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	207460	-335230	vh_unshaded\\planet03	3960	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	-303350	-100230	vh_unshaded\\planet09	2352	2000000	112
\.


--
-- TOC entry 2983 (class 0 OID 49697)
-- Dependencies: 211
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 2984 (class 0 OID 49700)
-- Dependencies: 212
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	10570	15130	vh_main_sequence\\star_yellow03	28380	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	32670	-11350	vh_main_sequence\\star_red01	15010	1000000000	0
\.


--
-- TOC entry 2985 (class 0 OID 49706)
-- Dependencies: 213
-- Data for Name: universe_stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stations (id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
cf07bba9-90b2-4599-b1e3-84d797a67f0a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Station	24771.795632863843	-9938.30877953488	Sunfarm	740	25300	112.4
526f57f5-09e0-41c7-9a89-cd803ec0a065	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Station	207460	-335230	Fleet Armory	810	65000	23.33
\.


--
-- TOC entry 2986 (class 0 OID 49712)
-- Dependencies: 214
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 2987 (class 0 OID 49715)
-- Dependencies: 215
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid, startid) FROM stdin;
664a4f4d-4cb1-4e4f-a5c6-c9a970d87bf9	asdf	7ab518aea3a72084bf1f1e1ddeb7fa25563a47b73b643452960187b6405b72f1	2020-11-23 16:05:14.583626-05	0	682b6f69-454b-4c59-91cd-b28efcc03b31	49f06e89-29fb-4a02-a034-4b5d0702adac
7168a9e6-09dc-44f8-9943-bd1231df7b9f	nwiehoff	02c4bf7d7e35c6bab999ac03ece60b8586a27f7ecd4830983b138b74262bf3f9	2020-11-23 16:03:29.677724-05	0	6aaf3890-450a-46d9-9951-3792bf4e2d1b	49f06e89-29fb-4a02-a034-4b5d0702adac
e92c5611-f424-4f2f-abf9-35f727a2bb21	aaaa	1af66cfb4ce4d7e11f4b67079e2c3c889a0f6a085e91107d49fcb63e3fc2ef42	2020-11-23 22:40:16.53804-05	0	2b653c97-e671-419c-bdc4-ac51d3d29561	49f06e89-29fb-4a02-a034-4b5d0702adac
\.


--
-- TOC entry 2794 (class 2606 OID 49719)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2798 (class 2606 OID 49721)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 2800 (class 2606 OID 49723)
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- TOC entry 2802 (class 2606 OID 49725)
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- TOC entry 2804 (class 2606 OID 49727)
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- TOC entry 2829 (class 2606 OID 49835)
-- Name: starts starts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT starts_pkey PRIMARY KEY (id);


--
-- TOC entry 2806 (class 2606 OID 49729)
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- TOC entry 2808 (class 2606 OID 49731)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 2810 (class 2606 OID 49733)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 2812 (class 2606 OID 49735)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2814 (class 2606 OID 49737)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 2816 (class 2606 OID 49739)
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- TOC entry 2818 (class 2606 OID 49741)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 2820 (class 2606 OID 49743)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2787 (class 2606 OID 49745)
-- Name: itemfamilies uq_itemfamily_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemfamilies
    ADD CONSTRAINT uq_itemfamily_id PRIMARY KEY (id);


--
-- TOC entry 2790 (class 2606 OID 49747)
-- Name: items uq_items_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT uq_items_id PRIMARY KEY (id);


--
-- TOC entry 2792 (class 2606 OID 49749)
-- Name: itemtypes uq_itemtype_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT uq_itemtype_id PRIMARY KEY (id);


--
-- TOC entry 2796 (class 2606 OID 49751)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 2823 (class 2606 OID 49753)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2825 (class 2606 OID 49755)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 2788 (class 1259 OID 49854)
-- Name: fki_fk_items_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_users ON public.items USING btree (createdby);


--
-- TOC entry 2826 (class 1259 OID 49860)
-- Name: fki_fk_starts_homestations; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_homestations ON public.starts USING btree (homestationid);


--
-- TOC entry 2827 (class 1259 OID 49848)
-- Name: fki_fk_starts_systems; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_systems ON public.starts USING btree (systemid);


--
-- TOC entry 2821 (class 1259 OID 49756)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 2830 (class 2606 OID 49757)
-- Name: items fk_items_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 2831 (class 2606 OID 49849)
-- Name: items fk_items_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_users FOREIGN KEY (createdby) REFERENCES public.users(id);


--
-- TOC entry 2832 (class 2606 OID 49762)
-- Name: itemtypes fk_itemtypes_itemfamilies; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT fk_itemtypes_itemfamilies FOREIGN KEY (family) REFERENCES public.itemfamilies(id);


--
-- TOC entry 2836 (class 2606 OID 49767)
-- Name: ships fk_ships_dockstations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_dockstations FOREIGN KEY (dockedat_stationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2835 (class 2606 OID 49772)
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2833 (class 2606 OID 49777)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2834 (class 2606 OID 49782)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 2837 (class 2606 OID 49787)
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id);


--
-- TOC entry 2847 (class 2606 OID 49855)
-- Name: starts fk_starts_homestations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_homestations FOREIGN KEY (homestationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2845 (class 2606 OID 49836)
-- Name: starts fk_starts_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2846 (class 2606 OID 49843)
-- Name: starts fk_starts_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_systems FOREIGN KEY (systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2843 (class 2606 OID 49792)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 2844 (class 2606 OID 49797)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 2838 (class 2606 OID 49802)
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2839 (class 2606 OID 49807)
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2840 (class 2606 OID 49812)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2841 (class 2606 OID 49817)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2842 (class 2606 OID 49822)
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2020-11-23 23:09:05

--
-- PostgreSQL database dump complete
--

