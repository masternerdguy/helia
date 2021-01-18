--
-- PostgreSQL database dump
--

-- Dumped from database version 10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)

-- Started on 2021-01-17 20:18:24 EST

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
-- TOC entry 3234 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 25698)
-- Name: containers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.containers (
    id uuid NOT NULL,
    meta jsonb NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.containers OWNER TO postgres;

--
-- TOC entry 197 (class 1259 OID 25704)
-- Name: itemfamilies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.itemfamilies (
    id character varying(16) NOT NULL,
    friendlyname character varying(64) NOT NULL,
    meta jsonb NOT NULL
);


ALTER TABLE public.itemfamilies OWNER TO postgres;

--
-- TOC entry 198 (class 1259 OID 25710)
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
-- TOC entry 199 (class 1259 OID 25718)
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
-- TOC entry 200 (class 1259 OID 25724)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    userid uuid NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 25727)
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
-- TOC entry 202 (class 1259 OID 25743)
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
-- TOC entry 203 (class 1259 OID 25766)
-- Name: shiptypes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptypes (
    id uuid NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.shiptypes OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 25769)
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
-- TOC entry 205 (class 1259 OID 25777)
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
-- TOC entry 206 (class 1259 OID 25783)
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
-- TOC entry 207 (class 1259 OID 25789)
-- Name: universe_regions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_regions (
    id uuid NOT NULL,
    regionname character varying(32) NOT NULL
);


ALTER TABLE public.universe_regions OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 25792)
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
-- TOC entry 209 (class 1259 OID 25798)
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
-- TOC entry 210 (class 1259 OID 25804)
-- Name: universe_systems; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_systems (
    id uuid NOT NULL,
    systemname character varying(32) NOT NULL,
    regionid uuid NOT NULL
);


ALTER TABLE public.universe_systems OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 25807)
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
-- TOC entry 3211 (class 0 OID 25698)
-- Dependencies: 196
-- Data for Name: containers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.containers (id, meta, created) FROM stdin;
33d0e729-b352-4979-86bc-6d92a25bd753	{}	2021-01-17 19:22:31.282482-05
02bfbe50-116f-4cf0-8c1e-024f377de9a5	{}	2021-01-17 19:22:31.284819-05
95a7a02c-325d-4802-985c-b895151add23	{}	2021-01-17 19:22:31.291712-05
da2ddd1d-6cf7-43f6-8fb3-43507bb4229a	{}	2021-01-17 19:26:14.57823-05
65d8f341-0fc7-4241-a884-42865a4d2042	{}	2021-01-17 19:26:14.580587-05
ec7f31bd-e30d-42bb-8346-09b3098de3bf	{}	2021-01-17 19:26:14.587558-05
aa74d378-bcf6-43ef-ac78-542a946af38a	{}	2021-01-17 19:30:32.842208-05
123c9fd8-b3f7-4ec5-af12-50fe7df4298c	{}	2021-01-17 19:30:32.844337-05
44a753f5-a5fe-4202-9790-2cf8ff84999c	{}	2021-01-17 19:30:32.850784-05
fa4e6246-7386-4d5e-9bda-083b130e7c1c	{}	2021-01-17 19:35:11.240611-05
77091e39-b19d-4bb6-b9ab-24eb1851f390	{}	2021-01-17 19:35:11.242406-05
28a510c5-bb86-41bb-a1b5-eeaf7db6c91a	{}	2021-01-17 19:35:11.249396-05
384f4ff1-5f66-4a4d-81f9-fc03679a43f2	{}	2021-01-17 19:48:45.597558-05
f032a365-e099-441e-86f6-71a72bf3e808	{}	2021-01-17 19:48:45.60112-05
ca14571d-373e-40af-9d44-7fb30bcc7f74	{}	2021-01-17 19:48:45.609816-05
3786ae3f-acc3-473b-a8fe-ef75e28d9044	{}	2021-01-17 19:52:15.8217-05
43783fa6-c74f-46f2-b216-6d8537c1f7fd	{}	2021-01-17 19:52:15.823485-05
a2f1216f-f0a6-447b-bf8e-536945569ca8	{}	2021-01-17 19:52:15.831048-05
70edc0af-6119-4c3f-956a-1a50730b171c	{}	2021-01-17 20:03:27.283175-05
d74f7f91-d39b-4a36-8923-7f92f4f9074c	{}	2021-01-17 20:03:27.285401-05
17e399f8-c3c9-4be4-9ff3-7b35be5ba8ea	{}	2021-01-17 20:03:27.292364-05
\.


--
-- TOC entry 3212 (class 0 OID 25704)
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
\.


--
-- TOC entry 3213 (class 0 OID 25710)
-- Dependencies: 198
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, itemtypeid, meta, created, createdby, createdreason, containerid, quantity, ispackaged) FROM stdin;
2a91d12e-be8d-4f0c-b97a-324811a7da2c	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 20:03:27.341305-05	fcc1e295-1fd1-4af7-be83-8e9f13fc6821	Module for new noob ship for player	17e399f8-c3c9-4be4-9ff3-7b35be5ba8ea	1	f
58114f4e-7574-4387-9c22-19552961ff2c	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 20:03:27.30416-05	fcc1e295-1fd1-4af7-be83-8e9f13fc6821	Module for new noob ship for player	d74f7f91-d39b-4a36-8923-7f92f4f9074c	1	f
80384f8a-2ec5-4411-8faf-3c3238b708c4	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:22:31.304311-05	9a1bd390-37e7-4d51-bf75-f6793b863122	Module for new noob ship for player	95a7a02c-325d-4802-985c-b895151add23	1	f
6eb8f142-1a28-4d62-927b-329596ec9ee5	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:22:31.313579-05	9a1bd390-37e7-4d51-bf75-f6793b863122	Module for new noob ship for player	95a7a02c-325d-4802-985c-b895151add23	1	f
63b8bd6c-2f90-4416-b28a-2944cb52e91f	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-17 19:22:31.330757-05	9a1bd390-37e7-4d51-bf75-f6793b863122	Module for new noob ship for player	95a7a02c-325d-4802-985c-b895151add23	1	f
8682521e-c9b0-4dfb-9e69-885d0ef5c4dd	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 19:22:31.338448-05	9a1bd390-37e7-4d51-bf75-f6793b863122	Module for new noob ship for player	95a7a02c-325d-4802-985c-b895151add23	1	f
a8dc4107-eaed-4316-aece-3100889520d0	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 19:22:31.323052-05	9a1bd390-37e7-4d51-bf75-f6793b863122	Module for new noob ship for player	95a7a02c-325d-4802-985c-b895151add23	1	f
550f9e51-d28a-4b44-aee3-0c4bb27be663	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:26:14.600431-05	6ee90928-956d-4347-a504-e1edc27dddc1	Module for new noob ship for player	ec7f31bd-e30d-42bb-8346-09b3098de3bf	1	f
0b6cc6b4-a71a-4c3a-9433-b8377a628593	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:26:14.611618-05	6ee90928-956d-4347-a504-e1edc27dddc1	Module for new noob ship for player	ec7f31bd-e30d-42bb-8346-09b3098de3bf	1	f
f411fed4-3923-40a8-830e-8cc32d8a96c7	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-17 19:26:14.630018-05	6ee90928-956d-4347-a504-e1edc27dddc1	Module for new noob ship for player	ec7f31bd-e30d-42bb-8346-09b3098de3bf	1	f
d2c8ac95-42ca-41f6-af3f-b06b65be4108	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 19:26:14.64023-05	6ee90928-956d-4347-a504-e1edc27dddc1	Module for new noob ship for player	ec7f31bd-e30d-42bb-8346-09b3098de3bf	1	f
1a384bce-37e9-4e27-b229-8be3a3099b88	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 19:26:14.621451-05	6ee90928-956d-4347-a504-e1edc27dddc1	Module for new noob ship for player	ec7f31bd-e30d-42bb-8346-09b3098de3bf	1	f
e823c1ca-00a1-4844-8aa8-2f368b702381	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:30:32.864823-05	ae7ee448-e913-4f37-9513-5434198c521c	Module for new noob ship for player	44a753f5-a5fe-4202-9790-2cf8ff84999c	1	f
26b3645d-714a-4e9f-b005-f442849c3f5a	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 19:30:32.883001-05	ae7ee448-e913-4f37-9513-5434198c521c	Module for new noob ship for player	44a753f5-a5fe-4202-9790-2cf8ff84999c	1	f
bcb18f7f-5382-4e02-add9-9b4916b24bf6	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 20:03:27.313991-05	fcc1e295-1fd1-4af7-be83-8e9f13fc6821	Module for new noob ship for player	17e399f8-c3c9-4be4-9ff3-7b35be5ba8ea	1	f
a7191190-672a-446a-9f56-c0cf968b099b	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:30:32.87342-05	ae7ee448-e913-4f37-9513-5434198c521c	Module for new noob ship for player	44a753f5-a5fe-4202-9790-2cf8ff84999c	1	f
6075e39a-07c8-4339-9686-7374deceb782	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 20:03:27.324452-05	fcc1e295-1fd1-4af7-be83-8e9f13fc6821	Module for new noob ship for player	17e399f8-c3c9-4be4-9ff3-7b35be5ba8ea	1	f
6369b09d-2681-4c8c-a956-533b22125a00	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-17 19:30:32.892777-05	ae7ee448-e913-4f37-9513-5434198c521c	Module for new noob ship for player	44a753f5-a5fe-4202-9790-2cf8ff84999c	1	f
bc490ce1-46af-4ee6-a764-463e6403f522	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 19:30:32.90265-05	ae7ee448-e913-4f37-9513-5434198c521c	Module for new noob ship for player	44a753f5-a5fe-4202-9790-2cf8ff84999c	1	f
33ad0a14-a957-48f3-a6c7-b4109b11d824	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:35:11.262451-05	1bbcaf85-d2e6-4ef3-9195-be9cfe3954f5	Module for new noob ship for player	28a510c5-bb86-41bb-a1b5-eeaf7db6c91a	1	f
e9aa1530-0b4a-48c5-8844-c5ae1d0d34cb	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:35:11.272173-05	1bbcaf85-d2e6-4ef3-9195-be9cfe3954f5	Module for new noob ship for player	28a510c5-bb86-41bb-a1b5-eeaf7db6c91a	1	f
c6a7af2e-8ba1-462b-b062-0470dccd84c7	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 19:35:11.281602-05	1bbcaf85-d2e6-4ef3-9195-be9cfe3954f5	Module for new noob ship for player	28a510c5-bb86-41bb-a1b5-eeaf7db6c91a	1	f
7c5771f1-6c62-4afb-9624-0f23e3ee3656	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-17 19:35:11.291428-05	1bbcaf85-d2e6-4ef3-9195-be9cfe3954f5	Module for new noob ship for player	28a510c5-bb86-41bb-a1b5-eeaf7db6c91a	1	f
acebc975-3680-484c-a4cf-5f3ae3552877	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 19:35:11.30076-05	1bbcaf85-d2e6-4ef3-9195-be9cfe3954f5	Module for new noob ship for player	28a510c5-bb86-41bb-a1b5-eeaf7db6c91a	1	f
5b634029-0e7f-485a-97ba-f0bea2f457b2	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:48:45.625476-05	b4a5d5a6-ef08-4746-a8d9-85ce2d8e5220	Module for new noob ship for player	ca14571d-373e-40af-9d44-7fb30bcc7f74	1	f
4b21c71b-596b-489d-80de-ca18b3a99a0f	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 19:48:45.650937-05	b4a5d5a6-ef08-4746-a8d9-85ce2d8e5220	Module for new noob ship for player	ca14571d-373e-40af-9d44-7fb30bcc7f74	1	f
765db619-c0d8-4aa6-ac13-a04c558ceeab	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-17 19:48:45.663467-05	b4a5d5a6-ef08-4746-a8d9-85ce2d8e5220	Module for new noob ship for player	ca14571d-373e-40af-9d44-7fb30bcc7f74	1	f
91777cb5-63ed-4f23-a5d3-b36d1f608111	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 19:48:45.67746-05	b4a5d5a6-ef08-4746-a8d9-85ce2d8e5220	Module for new noob ship for player	ca14571d-373e-40af-9d44-7fb30bcc7f74	1	f
672285cc-7ea7-47b8-afe6-63de9cb54dc2	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:48:45.638757-05	b4a5d5a6-ef08-4746-a8d9-85ce2d8e5220	Module for new noob ship for player	f032a365-e099-441e-86f6-71a72bf3e808	1	f
17f847c7-3e2a-4f6e-90fd-f9fb8868cf4a	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2021-01-17 20:03:27.332523-05	fcc1e295-1fd1-4af7-be83-8e9f13fc6821	Module for new noob ship for player	d74f7f91-d39b-4a36-8923-7f92f4f9074c	1	t
1dc0ee27-e64f-44ed-b0bb-36decf74a29e	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:52:15.856486-05	8f55859b-abf5-46a2-84a9-a8c30d8c04af	Module for new noob ship for player	a2f1216f-f0a6-447b-bf8e-536945569ca8	1	f
5c6d338b-11d2-4eb9-86ce-097e297b2153	b481a521-1b12-4ffa-ac2f-4da015036f7f	{}	2021-01-17 19:52:15.876356-05	8f55859b-abf5-46a2-84a9-a8c30d8c04af	Module for new noob ship for player	43783fa6-c74f-46f2-b216-6d8537c1f7fd	1	t
c8b0efa1-b760-43e0-8604-f3d35578d187	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 19:52:15.885775-05	8f55859b-abf5-46a2-84a9-a8c30d8c04af	Module for new noob ship for player	a2f1216f-f0a6-447b-bf8e-536945569ca8	1	f
55842ea0-d0a5-41d8-bc20-15eeb9d726d9	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 19:52:15.866541-05	8f55859b-abf5-46a2-84a9-a8c30d8c04af	Module for new noob ship for player	a2f1216f-f0a6-447b-bf8e-536945569ca8	1	f
a4d45c4e-ba3f-4b69-8af8-9523174d2cd6	9d1014c5-3422-4a0f-9839-f585269b4b16	{}	2021-01-17 19:52:15.845814-05	8f55859b-abf5-46a2-84a9-a8c30d8c04af	Module for new noob ship for player	43783fa6-c74f-46f2-b216-6d8537c1f7fd	1	t
\.


--
-- TOC entry 3214 (class 0 OID 25718)
-- Dependencies: 199
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
-- TOC entry 3215 (class 0 OID 25724)
-- Dependencies: 200
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
ddc918fb-7ee1-4256-95b2-36c297f6a1fa	9a1bd390-37e7-4d51-bf75-f6793b863122
37e5ef07-0f9f-43a1-b0c4-aeee0a31856a	6ee90928-956d-4347-a504-e1edc27dddc1
1bfe041d-7b88-4503-9753-3f673a25af30	ae7ee448-e913-4f37-9513-5434198c521c
035bdf8f-2061-42c1-ae76-40bb5b709ea9	1bbcaf85-d2e6-4ef3-9195-be9cfe3954f5
2937653d-2aae-4835-8285-754015a86292	b4a5d5a6-ef08-4746-a8d9-85ce2d8e5220
848b470f-310c-4823-8fe7-5b2075c32f45	8f55859b-abf5-46a2-84a9-a8c30d8c04af
44e8c009-6be9-492d-b715-b7c62ffb573f	fcc1e295-1fd1-4af7-be83-8e9f13fc6821
\.


--
-- TOC entry 3216 (class 0 OID 25727)
-- Dependencies: 201
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid, fittingbay_containerid, remaxdirty, trash_containerid) FROM stdin;
4c226a60-3cd0-420f-b717-4b2b4548d873	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	9a1bd390-37e7-4d51-bf75-f6793b863122	48081	-22113	2021-01-17 19:22:31.341492-05	nwiehoff's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "80384f8a-2ec5-4411-8faf-3c3238b708c4", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "6eb8f142-1a28-4d62-927b-329596ec9ee5", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "a8dc4107-eaed-4316-aece-3100889520d0", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "63b8bd6c-2f90-4416-b28a-2944cb52e91f", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "8682521e-c9b0-4dfb-9e69-885d0ef5c4dd", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	02bfbe50-116f-4cf0-8c1e-024f377de9a5	95a7a02c-325d-4802-985c-b895151add23	t	33d0e729-b352-4979-86bc-6d92a25bd753
b52dbde8-5644-4d9e-923d-53dd3c3c918a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	6ee90928-956d-4347-a504-e1edc27dddc1	48081	-22113	2021-01-17 19:26:14.644246-05	lll's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "550f9e51-d28a-4b44-aee3-0c4bb27be663", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "0b6cc6b4-a71a-4c3a-9433-b8377a628593", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "1a384bce-37e9-4e27-b229-8be3a3099b88", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "f411fed4-3923-40a8-830e-8cc32d8a96c7", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "d2c8ac95-42ca-41f6-af3f-b06b65be4108", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	65d8f341-0fc7-4241-a884-42865a4d2042	ec7f31bd-e30d-42bb-8346-09b3098de3bf	t	da2ddd1d-6cf7-43f6-8fb3-43507bb4229a
39e28249-9498-4d28-9f23-003f2ea29a75	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	ae7ee448-e913-4f37-9513-5434198c521c	-18153	34059	2021-01-17 19:30:32.906564-05	ooo's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "e823c1ca-00a1-4844-8aa8-2f368b702381", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "a7191190-672a-446a-9f56-c0cf968b099b", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "26b3645d-714a-4e9f-b005-f442849c3f5a", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "6369b09d-2681-4c8c-a956-533b22125a00", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "bc490ce1-46af-4ee6-a764-463e6403f522", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	123c9fd8-b3f7-4ec5-af12-50fe7df4298c	44a753f5-a5fe-4202-9790-2cf8ff84999c	t	aa74d378-bcf6-43ef-ac78-542a946af38a
33437b16-caee-47f9-9f37-83e884816dc1	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	1bbcaf85-d2e6-4ef3-9195-be9cfe3954f5	-47919	-8682	2021-01-17 19:35:11.30371-05	eee's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "33ad0a14-a957-48f3-a6c7-b4109b11d824", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "e9aa1530-0b4a-48c5-8844-c5ae1d0d34cb", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "c6a7af2e-8ba1-462b-b062-0470dccd84c7", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "7c5771f1-6c62-4afb-9624-0f23e3ee3656", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "acebc975-3680-484c-a4cf-5f3ae3552877", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	77091e39-b19d-4bb6-b9ab-24eb1851f390	28a510c5-bb86-41bb-a1b5-eeaf7db6c91a	t	fa4e6246-7386-4d5e-9bda-083b130e7c1c
a10254e6-25fe-4098-b559-cf2bb1a65abf	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	b4a5d5a6-ef08-4746-a8d9-85ce2d8e5220	48081	-22113	2021-01-17 19:48:45.682083-05	uuu's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "5b634029-0e7f-485a-97ba-f0bea2f457b2", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "672285cc-7ea7-47b8-afe6-63de9cb54dc2", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "4b21c71b-596b-489d-80de-ca18b3a99a0f", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "765db619-c0d8-4aa6-ac13-a04c558ceeab", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "91777cb5-63ed-4f23-a5d3-b36d1f608111", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	f032a365-e099-441e-86f6-71a72bf3e808	ca14571d-373e-40af-9d44-7fb30bcc7f74	t	384f4ff1-5f66-4a4d-81f9-fc03679a43f2
b45f0037-f075-4548-89e9-c0bf3879b9e2	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	8f55859b-abf5-46a2-84a9-a8c30d8c04af	-18153	34059	2021-01-17 19:52:15.88943-05	jjj's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "a4d45c4e-ba3f-4b69-8af8-9523174d2cd6", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "1dc0ee27-e64f-44ed-b0bb-36decf74a29e", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "55842ea0-d0a5-41d8-bc20-15eeb9d726d9", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "5c6d338b-11d2-4eb9-86ce-097e297b2153", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "c8b0efa1-b760-43e0-8604-f3d35578d187", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	43783fa6-c74f-46f2-b216-6d8537c1f7fd	a2f1216f-f0a6-447b-bf8e-536945569ca8	t	3786ae3f-acc3-473b-a8fe-ef75e28d9044
43fd0312-074d-47f2-bfa6-9bc1659a68a8	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	fcc1e295-1fd1-4af7-be83-8e9f13fc6821	48081	-22113	2021-01-17 20:03:27.344616-05	ggg's Starter Ship	Sparrow	0	0	0	209	169	135	265	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "58114f4e-7574-4387-9c22-19552961ff2c", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "bcb18f7f-5382-4e02-add9-9b4916b24bf6", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "6075e39a-07c8-4339-9686-7374deceb782", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "17f847c7-3e2a-4f6e-90fd-f9fb8868cf4a", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "2a91d12e-be8d-4f0c-b97a-324811a7da2c", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	d74f7f91-d39b-4a36-8923-7f92f4f9074c	17e399f8-c3c9-4be4-9ff3-7b35be5ba8ea	t	70edc0af-6119-4c3f-956a-1a50730b171c
\.


--
-- TOC entry 3217 (class 0 OID 25743)
-- Dependencies: 202
-- Data for Name: shiptemplates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shiptemplates (id, created, shiptemplatename, texture, radius, baseaccel, basemass, baseturn, baseshield, baseshieldregen, basearmor, basehull, basefuel, baseheatcap, baseheatsink, baseenergy, baseenergyregen, shiptypeid, slotlayout, basecargobayvolume) FROM stdin;
8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	2020-11-23 22:14:30.004993-05	Sparrow	Sparrow	12.5	4.29999999999999982	100	4.70000000000000018	209	6	169	135	265	737	11	138	9	e364a553-1dc5-4e8d-9195-0ca4989bec49	{"a_slots": [{"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "gun"}, {"hp_pos": [0, 0, 0], "volume": 10, "mod_family": "utility"}], "b_slots": [{"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 8, "mod_family": "any"}], "c_slots": [{"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}, {"hp_pos": [0, 0, 0], "volume": 6, "mod_family": "any"}]}	120
\.


--
-- TOC entry 3218 (class 0 OID 25766)
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
-- TOC entry 3219 (class 0 OID 25769)
-- Dependencies: 204
-- Data for Name: starts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.starts (id, name, shiptemplateid, shipfitting, created, available, systemid, homestationid) FROM stdin;
49f06e89-29fb-4a02-a034-4b5d0702adac	Test Start	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	{"a_rack": [{"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {}], "b_rack": [{"item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {}], "c_rack": [{"item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	2020-11-23 15:31:55.475609-05	t	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cf07bba9-90b2-4599-b1e3-84d797a67f0a
\.


--
-- TOC entry 3220 (class 0 OID 25777)
-- Dependencies: 205
-- Data for Name: universe_jumpholes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_jumpholes (id, universe_systemid, out_systemid, jumpholename, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
0602c3bf-7d70-4a4f-9ebd-77e7cec7ff12	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	edf08406-0879-4141-8af1-f68d32e31c8d	Test Jumphole 1	25000	25000	Jumphole	250	9999999	25
834572ef-b709-4ea7-9cd9-b526744f38cc	edf08406-0879-4141-8af1-f68d32e31c8d	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Jumphole 2	-600000	350000	Jumphole	250	9999999	112
\.


--
-- TOC entry 3221 (class 0 OID 25783)
-- Dependencies: 206
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	207460	-335230	vh_unshaded\\planet03	3960	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	-303350	-100230	vh_unshaded\\planet09	2352	2000000	112
\.


--
-- TOC entry 3222 (class 0 OID 25789)
-- Dependencies: 207
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 3223 (class 0 OID 25792)
-- Dependencies: 208
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	10570	15130	vh_main_sequence\\star_yellow03	28380	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	32670	-11350	vh_main_sequence\\star_red01	15010	1000000000	0
\.


--
-- TOC entry 3224 (class 0 OID 25798)
-- Dependencies: 209
-- Data for Name: universe_stations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stations (id, universe_systemid, stationname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
526f57f5-09e0-41c7-9a89-cd803ec0a065	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Station	207460	-335230	Fleet Armory	810	65000	23.3299999999999983
cf07bba9-90b2-4599-b1e3-84d797a67f0a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Station	24771.7956328638429	-9938.30877953488016	Sunfarm	740	25300	112.400000000000006
\.


--
-- TOC entry 3225 (class 0 OID 25804)
-- Dependencies: 210
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 3226 (class 0 OID 25807)
-- Dependencies: 211
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid, startid) FROM stdin;
9a1bd390-37e7-4d51-bf75-f6793b863122	nwiehoff	02c4bf7d7e35c6bab999ac03ece60b8586a27f7ecd4830983b138b74262bf3f9	2021-01-17 19:22:31.2711-05	0	4c226a60-3cd0-420f-b717-4b2b4548d873	49f06e89-29fb-4a02-a034-4b5d0702adac
6ee90928-956d-4347-a504-e1edc27dddc1	lll	6c7ed138a187034ead7b715c6eeb86ee1fcedc73cbc2fdc07cb02fd751601022	2021-01-17 19:26:14.566828-05	0	b52dbde8-5644-4d9e-923d-53dd3c3c918a	49f06e89-29fb-4a02-a034-4b5d0702adac
ae7ee448-e913-4f37-9513-5434198c521c	ooo	bfac7f1ae62d4e667ecdc4b64d7059cfc53e155bab826f8bfff41749d816fa07	2021-01-17 19:30:32.823417-05	0	39e28249-9498-4d28-9f23-003f2ea29a75	49f06e89-29fb-4a02-a034-4b5d0702adac
1bbcaf85-d2e6-4ef3-9195-be9cfe3954f5	eee	92b163cfac723e7724f5b1a3f14e8c2590f503cd489de5f1c0c9e094e009e9cb	2021-01-17 19:35:11.229805-05	0	33437b16-caee-47f9-9f37-83e884816dc1	49f06e89-29fb-4a02-a034-4b5d0702adac
b4a5d5a6-ef08-4746-a8d9-85ce2d8e5220	uuu	c41d2894a8006f7dd2380c94c6989c4cdb6220ac37ff7ae45c35e8f0e5ef7306	2021-01-17 19:48:45.585163-05	0	a10254e6-25fe-4098-b559-cf2bb1a65abf	49f06e89-29fb-4a02-a034-4b5d0702adac
8f55859b-abf5-46a2-84a9-a8c30d8c04af	jjj	27900e03a78a054db2a71f21ac5f1e58afe7af7891ffd49268c6b31f33629290	2021-01-17 19:52:15.801916-05	0	b45f0037-f075-4548-89e9-c0bf3879b9e2	49f06e89-29fb-4a02-a034-4b5d0702adac
fcc1e295-1fd1-4af7-be83-8e9f13fc6821	ggg	0ca0eeb79267c8565f8744d299b3ef5fc0fc80d9ce9c5527f8c52077e25c4a26	2021-01-17 20:03:27.270528-05	0	43fd0312-074d-47f2-bfa6-9bc1659a68a8	49f06e89-29fb-4a02-a034-4b5d0702adac
\.


--
-- TOC entry 3029 (class 2606 OID 25811)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 3036 (class 2606 OID 25813)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 3038 (class 2606 OID 25815)
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- TOC entry 3040 (class 2606 OID 25817)
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- TOC entry 3042 (class 2606 OID 25819)
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- TOC entry 3046 (class 2606 OID 25821)
-- Name: starts starts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT starts_pkey PRIMARY KEY (id);


--
-- TOC entry 3048 (class 2606 OID 25823)
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- TOC entry 3050 (class 2606 OID 25825)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 3052 (class 2606 OID 25827)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 3054 (class 2606 OID 25829)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 3056 (class 2606 OID 25831)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 3058 (class 2606 OID 25833)
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- TOC entry 3060 (class 2606 OID 25835)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 3062 (class 2606 OID 25837)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 3019 (class 2606 OID 25839)
-- Name: containers uq_container_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.containers
    ADD CONSTRAINT uq_container_id PRIMARY KEY (id);


--
-- TOC entry 3021 (class 2606 OID 25841)
-- Name: itemfamilies uq_itemfamily_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemfamilies
    ADD CONSTRAINT uq_itemfamily_id PRIMARY KEY (id);


--
-- TOC entry 3025 (class 2606 OID 25843)
-- Name: items uq_items_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT uq_items_id PRIMARY KEY (id);


--
-- TOC entry 3027 (class 2606 OID 25845)
-- Name: itemtypes uq_itemtype_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT uq_itemtype_id PRIMARY KEY (id);


--
-- TOC entry 3031 (class 2606 OID 25847)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 3065 (class 2606 OID 25849)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 3067 (class 2606 OID 25851)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 3022 (class 1259 OID 25852)
-- Name: fki_fk_items_containers; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_containers ON public.items USING btree (containerid);


--
-- TOC entry 3023 (class 1259 OID 25853)
-- Name: fki_fk_items_users; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_items_users ON public.items USING btree (createdby);


--
-- TOC entry 3032 (class 1259 OID 25854)
-- Name: fki_fk_ships_containers_cargobay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_cargobay ON public.ships USING btree (cargobay_containerid);


--
-- TOC entry 3033 (class 1259 OID 25855)
-- Name: fki_fk_ships_containers_fittingbay; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_fittingbay ON public.ships USING btree (fittingbay_containerid);


--
-- TOC entry 3034 (class 1259 OID 25856)
-- Name: fki_fk_ships_containers_trash; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_ships_containers_trash ON public.ships USING btree (trash_containerid);


--
-- TOC entry 3043 (class 1259 OID 25857)
-- Name: fki_fk_starts_homestations; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_homestations ON public.starts USING btree (homestationid);


--
-- TOC entry 3044 (class 1259 OID 25858)
-- Name: fki_fk_starts_systems; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_starts_systems ON public.starts USING btree (systemid);


--
-- TOC entry 3063 (class 1259 OID 25859)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 3068 (class 2606 OID 25860)
-- Name: items fk_items_containers; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_containers FOREIGN KEY (containerid) REFERENCES public.containers(id);


--
-- TOC entry 3069 (class 2606 OID 25865)
-- Name: items fk_items_itemtypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_itemtypes FOREIGN KEY (itemtypeid) REFERENCES public.itemtypes(id);


--
-- TOC entry 3070 (class 2606 OID 25870)
-- Name: items fk_items_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_items_users FOREIGN KEY (createdby) REFERENCES public.users(id);


--
-- TOC entry 3071 (class 2606 OID 25875)
-- Name: itemtypes fk_itemtypes_itemfamilies; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.itemtypes
    ADD CONSTRAINT fk_itemtypes_itemfamilies FOREIGN KEY (family) REFERENCES public.itemfamilies(id);


--
-- TOC entry 3072 (class 2606 OID 25880)
-- Name: ships fk_ships_containers_cargobay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_cargobay FOREIGN KEY (cargobay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 3073 (class 2606 OID 25885)
-- Name: ships fk_ships_containers_fittingbay; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_fittingbay FOREIGN KEY (fittingbay_containerid) REFERENCES public.containers(id);


--
-- TOC entry 3074 (class 2606 OID 25890)
-- Name: ships fk_ships_containers_trash; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_containers_trash FOREIGN KEY (trash_containerid) REFERENCES public.containers(id);


--
-- TOC entry 3075 (class 2606 OID 25895)
-- Name: ships fk_ships_dockstations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_dockstations FOREIGN KEY (dockedat_stationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 3076 (class 2606 OID 25900)
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 3077 (class 2606 OID 25905)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3078 (class 2606 OID 25910)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 3079 (class 2606 OID 25915)
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id);


--
-- TOC entry 3080 (class 2606 OID 25920)
-- Name: starts fk_starts_homestations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_homestations FOREIGN KEY (homestationid) REFERENCES public.universe_stations(id);


--
-- TOC entry 3081 (class 2606 OID 25925)
-- Name: starts fk_starts_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id);


--
-- TOC entry 3082 (class 2606 OID 25930)
-- Name: starts fk_starts_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.starts
    ADD CONSTRAINT fk_starts_systems FOREIGN KEY (systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3088 (class 2606 OID 25935)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 3089 (class 2606 OID 25940)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 3083 (class 2606 OID 25945)
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3084 (class 2606 OID 25950)
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3085 (class 2606 OID 25955)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3086 (class 2606 OID 25960)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 3087 (class 2606 OID 25965)
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2021-01-17 20:18:25 EST

--
-- PostgreSQL database dump complete
--

