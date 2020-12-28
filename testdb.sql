--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2020-12-27 21:36:52

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
    quantity integer DEFAULT 1 NOT NULL
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
    remaxdirty boolean DEFAULT true NOT NULL
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
-- TOC entry 2990 (class 0 OID 50603)
-- Dependencies: 202
-- Data for Name: containers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.containers (id, meta, created) FROM stdin;
402440d9-7f17-435c-9bef-c0ea6260187b	{}	2020-12-15 08:20:53.110268-05
6266f40f-8226-43e3-a980-d4f03f5957c2	{}	2020-12-15 08:20:53.112155-05
bc398fe3-5097-46fe-acaa-6b1eaa1e3f9c	{}	2020-12-18 16:31:59.448904-05
e8ab90af-3189-400f-8876-2b6cb2f6abea	{}	2020-12-18 16:31:59.450532-05
7066c0b1-c781-46a3-bdfe-8fe3544c14c5	{}	2020-12-26 21:52:04.9213-05
6ba417b0-0950-4e99-8a74-9dbf55ce523b	{}	2020-12-26 21:52:04.925294-05
551e34f4-0d8f-4bb9-89b9-f856919c05db	{}	2020-12-27 21:32:15.967738-05
340629e5-f1d0-48ad-b17d-dee87d8d3a77	{}	2020-12-27 21:32:15.969738-05
\.


--
-- TOC entry 2991 (class 0 OID 50609)
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
-- TOC entry 2992 (class 0 OID 50615)
-- Dependencies: 204
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, itemtypeid, meta, created, createdby, createdreason, containerid, quantity) FROM stdin;
e2d1b362-3b5b-49ba-8e42-7691f5dfa69b	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 08:20:53.178609-05	fd7e6075-43ca-46a7-acd4-754212a72615	Module for new noob ship for player	6266f40f-8226-43e3-a980-d4f03f5957c2	1
4ac481b2-1990-4984-a352-ebaa2139be8f	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-15 08:20:53.250832-05	fd7e6075-43ca-46a7-acd4-754212a72615	Module for new noob ship for player	6266f40f-8226-43e3-a980-d4f03f5957c2	1
bbe6a120-5d5a-485c-afd1-fe269a288060	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-15 08:20:53.320153-05	fd7e6075-43ca-46a7-acd4-754212a72615	Module for new noob ship for player	6266f40f-8226-43e3-a980-d4f03f5957c2	1
f48e9a66-e854-416f-8ae3-31c0e56e3816	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2020-12-15 08:20:53.389244-05	fd7e6075-43ca-46a7-acd4-754212a72615	Module for new noob ship for player	6266f40f-8226-43e3-a980-d4f03f5957c2	1
a5af4784-b10d-4d70-a0c1-af2aec86bcc0	c311df30-c21e-4895-acb0-d8808f99710e	{}	2020-12-15 08:20:53.461005-05	fd7e6075-43ca-46a7-acd4-754212a72615	Module for new noob ship for player	6266f40f-8226-43e3-a980-d4f03f5957c2	1
a7692aa0-8fb6-4301-aaa3-0ca8abba1d8d	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-18 16:31:59.454193-05	626641b5-5e2b-4238-91d4-a455b5679998	Module for new noob ship for player	e8ab90af-3189-400f-8876-2b6cb2f6abea	1
0c87fb93-41a4-43ce-ba68-9097cce64168	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-18 16:31:59.460082-05	626641b5-5e2b-4238-91d4-a455b5679998	Module for new noob ship for player	e8ab90af-3189-400f-8876-2b6cb2f6abea	1
f250fe93-ab25-414a-b535-7c06658a6a4c	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-18 16:31:59.466065-05	626641b5-5e2b-4238-91d4-a455b5679998	Module for new noob ship for player	e8ab90af-3189-400f-8876-2b6cb2f6abea	1
0aba534a-d2c1-407e-82f4-042aa7c241e0	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2020-12-18 16:31:59.47184-05	626641b5-5e2b-4238-91d4-a455b5679998	Module for new noob ship for player	e8ab90af-3189-400f-8876-2b6cb2f6abea	1
7d7b1b8b-b336-4db3-b513-61aa11446358	c311df30-c21e-4895-acb0-d8808f99710e	{}	2020-12-18 16:31:59.477633-05	626641b5-5e2b-4238-91d4-a455b5679998	Module for new noob ship for player	e8ab90af-3189-400f-8876-2b6cb2f6abea	1
bfd7f3d1-2bf1-4ada-8e9d-1c2d06a256ff	91ec9901-ea7e-476a-bc65-7da4523dca38	{}	2020-12-18 00:00:00-05	fd7e6075-43ca-46a7-acd4-754212a72615	Manually added for testing	402440d9-7f17-435c-9bef-c0ea6260187b	1
0e1d36b0-4701-48f3-bd1d-dcc6c45d2791	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-26 21:52:04.996308-05	2a7a8af8-f8a1-40da-a41d-87c6f3c79b71	Module for new noob ship for player	6ba417b0-0950-4e99-8a74-9dbf55ce523b	1
8e4d5009-217b-4c26-8672-ac970adf469e	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-26 21:52:05.157362-05	2a7a8af8-f8a1-40da-a41d-87c6f3c79b71	Module for new noob ship for player	6ba417b0-0950-4e99-8a74-9dbf55ce523b	1
1a7f41ca-e90e-419b-90fc-6187420ec776	c311df30-c21e-4895-acb0-d8808f99710e	{}	2020-12-26 21:52:05.317294-05	2a7a8af8-f8a1-40da-a41d-87c6f3c79b71	Module for new noob ship for player	7066c0b1-c781-46a3-bdfe-8fe3544c14c5	1
6e6f4206-58c2-489d-8553-1f65384e6666	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-26 21:52:05.080294-05	2a7a8af8-f8a1-40da-a41d-87c6f3c79b71	Module for new noob ship for player	7066c0b1-c781-46a3-bdfe-8fe3544c14c5	1
1bffea1e-74c5-49d7-b3e8-af3926a398ff	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2020-12-26 21:52:05.239294-05	2a7a8af8-f8a1-40da-a41d-87c6f3c79b71	Module for new noob ship for player	7066c0b1-c781-46a3-bdfe-8fe3544c14c5	1
c720dd09-53a5-4d7b-ae4c-e01bd98795f8	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-27 21:32:16.040749-05	ceb14463-1720-4eb6-bfd1-a71e59591262	Module for new noob ship for player	340629e5-f1d0-48ad-b17d-dee87d8d3a77	1
74b2d151-06e6-4779-bb9f-90c6dd9ff771	c311df30-c21e-4895-acb0-d8808f99710e	{}	2020-12-27 21:32:16.343871-05	ceb14463-1720-4eb6-bfd1-a71e59591262	Module for new noob ship for player	340629e5-f1d0-48ad-b17d-dee87d8d3a77	1
3c0bd013-e896-4cdb-8ba6-dca54cb3905e	09172710-740c-4d1c-9fc0-43cb62e674e7	{}	2020-12-27 21:32:16.190866-05	ceb14463-1720-4eb6-bfd1-a71e59591262	Module for new noob ship for player	551e34f4-0d8f-4bb9-89b9-f856919c05db	1
d0c3ed58-471e-4bab-9e7c-56ce9968c68a	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2020-12-27 21:32:16.11586-05	ceb14463-1720-4eb6-bfd1-a71e59591262	Module for new noob ship for player	551e34f4-0d8f-4bb9-89b9-f856919c05db	1
ef81665f-34e7-465b-8b71-a7e0ea7536ad	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2020-12-27 21:32:16.266862-05	ceb14463-1720-4eb6-bfd1-a71e59591262	Module for new noob ship for player	551e34f4-0d8f-4bb9-89b9-f856919c05db	1
\.


--
-- TOC entry 2993 (class 0 OID 50621)
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
-- TOC entry 2994 (class 0 OID 50627)
-- Dependencies: 206
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
c13b3cf7-c4ed-4d43-b3f8-d23ebac16dc7	2a7a8af8-f8a1-40da-a41d-87c6f3c79b71
afdfd2a7-daa9-46d0-af88-d66dd3739389	fd7e6075-43ca-46a7-acd4-754212a72615
aad90676-3f75-4920-97e1-e48ba89bfced	ceb14463-1720-4eb6-bfd1-a71e59591262
ee5de67e-7acf-45be-8733-69a9bc043694	626641b5-5e2b-4238-91d4-a455b5679998
\.


--
-- TOC entry 2995 (class 0 OID 50630)
-- Dependencies: 207
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid, fittingbay_containerid, remaxdirty) FROM stdin;
0976b13e-80ad-4c44-ac24-dac6b15e653e	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	fd7e6075-43ca-46a7-acd4-754212a72615	25622.789966669417	-8098.701810015407	2020-12-15 08:20:53.531321-05	nwiehoff's Starter Ship	Sparrow	106.11341823308929	-1.5734954116528082e-115	-5.426308845209466e-115	209	244	135	275.61404525205137	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "e2d1b362-3b5b-49ba-8e42-7691f5dfa69b", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "4ac481b2-1990-4984-a352-ebaa2139be8f", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "bbe6a120-5d5a-485c-afd1-fe269a288060", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_id": "f48e9a66-e854-416f-8ae3-31c0e56e3816", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "a5af4784-b10d-4d70-a0c1-af2aec86bcc0", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	402440d9-7f17-435c-9bef-c0ea6260187b	6266f40f-8226-43e3-a980-d4f03f5957c2	f
aedbd4af-1cf5-4c4f-8264-4bafe19bcaeb	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	626641b5-5e2b-4238-91d4-a455b5679998	31557.378479230338	14086.66345472699	2020-12-18 16:31:59.483604-05	asdf's Starter Ship	Sparrow	32.756093721776864	9.4e-323	-9.4e-323	209	244	135	282.8972016766344	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "a7692aa0-8fb6-4301-aaa3-0ca8abba1d8d", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "0c87fb93-41a4-43ce-ba68-9097cce64168", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "f250fe93-ab25-414a-b535-7c06658a6a4c", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_id": "0aba534a-d2c1-407e-82f4-042aa7c241e0", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "7d7b1b8b-b336-4db3-b513-61aa11446358", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	bc398fe3-5097-46fe-acaa-6b1eaa1e3f9c	e8ab90af-3189-400f-8876-2b6cb2f6abea	f
df800325-651b-4f51-8930-c65dd3b2a8e7	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	2a7a8af8-f8a1-40da-a41d-87c6f3c79b71	24771.795632863843	-9938.30877953488	2020-12-26 21:52:05.403294-05	lol's Starter Ship	Sparrow	216.86989764584402	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "0e1d36b0-4701-48f3-bd1d-dcc6c45d2791", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "8e4d5009-217b-4c26-8672-ac970adf469e", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": null}	f	\N	7066c0b1-c781-46a3-bdfe-8fe3544c14c5	6ba417b0-0950-4e99-8a74-9dbf55ce523b	f
3776ff2c-998d-40d9-a77f-08ffadeb93f4	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	ceb14463-1720-4eb6-bfd1-a71e59591262	24771.795632863843	-9938.30877953488	2020-12-27 21:32:16.423914-05	reee's Starter Ship	Sparrow	207.578632642979	0	0	209	244	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	cf07bba9-90b2-4599-b1e3-84d797a67f0a	{"a_rack": [{"item_id": "c720dd09-53a5-4d7b-ae4c-e01bd98795f8", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": null, "c_rack": [{"item_id": "74b2d151-06e6-4779-bb9f-90c6dd9ff771", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	551e34f4-0d8f-4bb9-89b9-f856919c05db	340629e5-f1d0-48ad-b17d-dee87d8d3a77	f
\.


--
-- TOC entry 2996 (class 0 OID 50646)
-- Dependencies: 208
-- Data for Name: shiptemplates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume) FROM stdin;
8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	2020-11-23 22:14:30.004993-05	Sparrow	Sparrow	12.5	4.3	100	4.7	209	6	169	135	265	737	11	138	9	e364a553-1dc5-4e8d-9195-0ca4989bec49	{"a_slots": [{"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}]}	120
\.


--
-- TOC entry 2997 (class 0 OID 50669)
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
-- TOC entry 2998 (class 0 OID 50672)
-- Dependencies: 210
-- Data for Name: starts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid) FROM stdin;
49f06e89-29fb-4a02-a034-4b5d0702adac	Test Start	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	2020-11-23 15:31:55.475609-05	t	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cf07bba9-90b2-4599-b1e3-84d797a67f0a
\.


--
-- TOC entry 2999 (class 0 OID 50680)
-- Dependencies: 211
-- Data for Name: universe_jumpholes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_jumpholes (id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
0602c3bf-7d70-4a4f-9ebd-77e7cec7ff12	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	edf08406-0879-4141-8af1-f68d32e31c8d	Test Jumphole 1	25000	25000	Jumphole	250	9999999	25
834572ef-b709-4ea7-9cd9-b526744f38cc	edf08406-0879-4141-8af1-f68d32e31c8d	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Jumphole 2	-600000	350000	Jumphole	250	9999999	112
\.


--
-- TOC entry 3000 (class 0 OID 50686)
-- Dependencies: 212
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	207460	-335230	vh_unshaded\\planet03	3960	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	-303350	-100230	vh_unshaded\\planet09	2352	2000000	112
\.


--
-- TOC entry 3001 (class 0 OID 50692)
-- Dependencies: 213
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 3002 (class 0 OID 50695)
-- Dependencies: 214
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	10570	15130	vh_main_sequence\\star_yellow03	28380	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	32670	-11350	vh_main_sequence\\star_red01	15010	1000000000	0
\.


--
-- TOC entry 3003 (class 0 OID 50701)
-- Dependencies: 215
-- Data for Name: universe_stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stations (id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
cf07bba9-90b2-4599-b1e3-84d797a67f0a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Station	24771.795632863843	-9938.30877953488	Sunfarm	740	25300	112.4
526f57f5-09e0-41c7-9a89-cd803ec0a065	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Station	207460	-335230	Fleet Armory	810	65000	23.33
\.


--
-- TOC entry 3004 (class 0 OID 50707)
-- Dependencies: 216
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 3005 (class 0 OID 50710)
-- Dependencies: 217
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid, startid) FROM stdin;
fd7e6075-43ca-46a7-acd4-754212a72615	nwiehoff	02c4bf7d7e35c6bab999ac03ece60b8586a27f7ecd4830983b138b74262bf3f9	2020-12-15 08:20:53.033058-05	0	0976b13e-80ad-4c44-ac24-dac6b15e653e	49f06e89-29fb-4a02-a034-4b5d0702adac
626641b5-5e2b-4238-91d4-a455b5679998	asdf	7ab518aea3a72084bf1f1e1ddeb7fa25563a47b73b643452960187b6405b72f1	2020-12-18 16:31:59.439553-05	0	aedbd4af-1cf5-4c4f-8264-4bafe19bcaeb	49f06e89-29fb-4a02-a034-4b5d0702adac
2a7a8af8-f8a1-40da-a41d-87c6f3c79b71	lol	7c17eeca2d07d6757b8e11e789522b94c73fb83f229574594566011bf2c0dcf5	2020-12-26 21:52:04.83012-05	0	df800325-651b-4f51-8930-c65dd3b2a8e7	49f06e89-29fb-4a02-a034-4b5d0702adac
ceb14463-1720-4eb6-bfd1-a71e59591262	reee	d40a31b9652b592fa3f200d4be20f5d48ea80988d14bd034d864d35599aa1a0f	2020-12-27 21:32:15.886738-05	0	3776ff2c-998d-40d9-a77f-08ffadeb93f4	49f06e89-29fb-4a02-a034-4b5d0702adac
\.


--
-- TOC entry 2805 (class 2606 OID 50714)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2811 (class 2606 OID 50716)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 2813 (class 2606 OID 50718)
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- TOC entry 2815 (class 2606 OID 50720)
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- TOC entry 2817 (class 2606 OID 50722)
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- TOC entry 2821 (class 2606 OID 50724)
-- Name: starts starts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT starts_pkey PRIMARY KEY (id);


--
-- TOC entry 2823 (class 2606 OID 50726)
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- TOC entry 2825 (class 2606 OID 50728)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 2827 (class 2606 OID 50730)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 2829 (class 2606 OID 50732)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2831 (class 2606 OID 50734)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 2833 (class 2606 OID 50736)
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- TOC entry 2835 (class 2606 OID 50738)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 2837 (class 2606 OID 50740)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2795 (class 2606 OID 50742)
-- Name: containers uq_container_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.containers
    ADD CONSTRAINT uq_container_id PRIMARY KEY (id);


--
-- TOC entry 2797 (class 2606 OID 50744)
-- Name: itemfamilies uq_itemfamily_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemfamilies
    ADD CONSTRAINT uq_itemfamily_id PRIMARY KEY (id);


--
-- TOC entry 2801 (class 2606 OID 50746)
-- Name: items uq_items_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT uq_items_id PRIMARY KEY (id);


--
-- TOC entry 2803 (class 2606 OID 50748)
-- Name: itemtypes uq_itemtype_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT uq_itemtype_id PRIMARY KEY (id);


--
-- TOC entry 2807 (class 2606 OID 50750)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 2840 (class 2606 OID 50752)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2842 (class 2606 OID 50754)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 2798 (class 1259 OID 50755)
-- Name: fki_fk_items_containers; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_containers ON public.items USING btree (containerid);


--
-- TOC entry 2799 (class 1259 OID 50756)
-- Name: fki_fk_items_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_users ON public.items USING btree (createdby);


--
-- TOC entry 2808 (class 1259 OID 50757)
-- Name: fki_fk_ships_containers_cargobay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_cargobay ON public.ships USING btree (cargobay_containerid);


--
-- TOC entry 2809 (class 1259 OID 50758)
-- Name: fki_fk_ships_containers_fittingbay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_fittingbay ON public.ships USING btree (fittingbay_containerid);


--
-- TOC entry 2818 (class 1259 OID 50759)
-- Name: fki_fk_starts_homestations; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_homestations ON public.starts USING btree (homestationid);


--
-- TOC entry 2819 (class 1259 OID 50760)
-- Name: fki_fk_starts_systems; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_systems ON public.starts USING btree (systemid);


--
-- TOC entry 2838 (class 1259 OID 50761)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 2843 (class 2606 OID 50762)
-- Name: items fk_items_containers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_containers FOREIGN KEY (containerid) REFERENCES public.containers(id);


--
-- TOC entry 2844 (class 2606 OID 50767)
-- Name: items fk_items_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 2845 (class 2606 OID 50772)
-- Name: items fk_items_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_users FOREIGN KEY (createdby) REFERENCES public.users(id);


--
-- TOC entry 2846 (class 2606 OID 50777)
-- Name: itemtypes fk_itemtypes_itemfamilies; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT fk_itemtypes_itemfamilies FOREIGN KEY (family) REFERENCES public.itemfamilies(id);


--
-- TOC entry 2847 (class 2606 OID 50782)
-- Name: ships fk_ships_containers_cargobay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_cargobay FOREIGN KEY (cargobay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2848 (class 2606 OID 50787)
-- Name: ships fk_ships_containers_fittingbay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_fittingbay FOREIGN KEY (fittingbay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 2849 (class 2606 OID 50792)
-- Name: ships fk_ships_dockstations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_dockstations FOREIGN KEY (dockedat_stationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2850 (class 2606 OID 50797)
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2851 (class 2606 OID 50802)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2852 (class 2606 OID 50807)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 2853 (class 2606 OID 50812)
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id);


--
-- TOC entry 2854 (class 2606 OID 50817)
-- Name: starts fk_starts_homestations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_homestations FOREIGN KEY (homestationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 2855 (class 2606 OID 50822)
-- Name: starts fk_starts_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 2856 (class 2606 OID 50827)
-- Name: starts fk_starts_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_systems FOREIGN KEY (systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2862 (class 2606 OID 50832)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 2863 (class 2606 OID 50837)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 2857 (class 2606 OID 50842)
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2858 (class 2606 OID 50847)
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2859 (class 2606 OID 50852)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2860 (class 2606 OID 50857)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2861 (class 2606 OID 50862)
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2020-12-27 21:36:53

--
-- PostgreSQL database dump complete
--

