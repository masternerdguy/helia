--
-- PostgreSQL database dump
--

-- Dumped from database version 10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.15 (Ubuntu 10.15-0ubuntu0.18.04.1)

-- Started on 2021-01-17 19:07:59 EST

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
b6dfc5a9-e9a5-46c5-b2ce-a9e7080fbe69	{}	2021-01-17 18:52:36.100391-05
48b28e75-01e1-443e-97eb-f2b2cc8073b0	{}	2021-01-17 18:52:36.102707-05
8cf5bf0f-d44d-4619-a489-e8d0e3ea7a85	{}	2021-01-17 18:52:36.109642-05
a719d8f6-e4ec-49d6-8175-f8cf1db19207	{}	2021-01-17 18:57:02.553541-05
c3f7a739-ac00-440e-b1cd-221c858e9df7	{}	2021-01-17 18:57:02.555676-05
4dd6b7e1-bf69-42a4-a695-74325c0768d4	{}	2021-01-17 18:57:02.561201-05
b7aa4c01-4168-419c-87ca-c7f1855a8311	{}	2021-01-17 19:02:52.422907-05
43315ba6-2040-414d-843a-7dca5c7dfd08	{}	2021-01-17 19:02:52.425086-05
2c7b5627-c8d3-40e3-b3f6-223375a7e46b	{}	2021-01-17 19:02:52.432304-05
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
531e80ea-455e-47a6-b896-0071338b16d7	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 18:52:36.136023-05	b37596b0-9841-4e27-a72c-eed954bb8c2d	Module for new noob ship for player	48b28e75-01e1-443e-97eb-f2b2cc8073b0	1	f
f2e8acbb-6ea5-4837-bdc8-05a59e3eb8b7	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-17 18:52:36.155718-05	b37596b0-9841-4e27-a72c-eed954bb8c2d	Module for new noob ship for player	48b28e75-01e1-443e-97eb-f2b2cc8073b0	1	f
081ea8f3-c85d-409c-88a6-964b4fbb4395	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 18:57:02.575627-05	cda4d030-ab22-45e3-b83e-2f5c233ce3fe	Module for new noob ship for player	4dd6b7e1-bf69-42a4-a695-74325c0768d4	1	f
babfeb23-ebd5-49f6-b473-1501c1be546c	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 18:57:02.593723-05	cda4d030-ab22-45e3-b83e-2f5c233ce3fe	Module for new noob ship for player	4dd6b7e1-bf69-42a4-a695-74325c0768d4	1	f
64f25d5c-8b76-452f-bee0-b81c88295092	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 18:57:02.608239-05	cda4d030-ab22-45e3-b83e-2f5c233ce3fe	Module for new noob ship for player	4dd6b7e1-bf69-42a4-a695-74325c0768d4	1	f
dcf866d4-c3a3-46c6-918b-c8bc2ea0aeef	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:02:52.456653-05	648e4e7c-2328-413f-94cf-38a0a554d6ae	Module for new noob ship for player	2c7b5627-c8d3-40e3-b3f6-223375a7e46b	1	f
5c12a7a2-2f0c-4e62-8a74-daf5c049ffcd	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-17 19:02:52.477324-05	648e4e7c-2328-413f-94cf-38a0a554d6ae	Module for new noob ship for player	2c7b5627-c8d3-40e3-b3f6-223375a7e46b	1	f
3dfc6052-44d5-4d6e-ac3d-0778ce33f7f9	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 19:02:52.486448-05	648e4e7c-2328-413f-94cf-38a0a554d6ae	Module for new noob ship for player	2c7b5627-c8d3-40e3-b3f6-223375a7e46b	1	f
102db1aa-ff16-43e7-a671-bed4535d7299	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 18:52:36.12453-05	b37596b0-9841-4e27-a72c-eed954bb8c2d	Module for new noob ship for player	48b28e75-01e1-443e-97eb-f2b2cc8073b0	1	f
cda043ad-dfef-47cf-8fe7-996ad09124ca	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 18:52:36.145444-05	b37596b0-9841-4e27-a72c-eed954bb8c2d	Module for new noob ship for player	48b28e75-01e1-443e-97eb-f2b2cc8073b0	1	f
0b2ce719-4b04-4648-9b4d-df6310959584	c311df30-c21e-4895-acb0-d8808f99710e	{"hp": 75, "rack": "c", "volume": 6, "armor_max_add": 75}	2021-01-17 18:52:36.166349-05	b37596b0-9841-4e27-a72c-eed954bb8c2d	Module for new noob ship for player	48b28e75-01e1-443e-97eb-f2b2cc8073b0	1	f
0cbbc87e-54f3-4abe-be16-a1d2bb8dd5f1	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 18:57:02.585482-05	cda4d030-ab22-45e3-b83e-2f5c233ce3fe	Module for new noob ship for player	4dd6b7e1-bf69-42a4-a695-74325c0768d4	1	f
2907bb02-fd20-437b-a2bc-7f81fdd91b39	b481a521-1b12-4ffa-ac2f-4da015036f7f	{"hp": 15, "rack": "c", "volume": 3, "fuel_max_add": 30}	2021-01-17 18:57:02.602252-05	cda4d030-ab22-45e3-b83e-2f5c233ce3fe	Module for new noob ship for player	4dd6b7e1-bf69-42a4-a695-74325c0768d4	1	f
fb96c08d-09cb-40ff-9293-33e964363761	9d1014c5-3422-4a0f-9839-f585269b4b16	{"hp": 10, "rack": "a", "range": 1528, "volume": 4, "falloff": "linear", "cooldown": 5, "hull_damage": 4, "armor_damage": 1, "can_mine_gas": false, "can_mine_ice": false, "can_mine_ore": true, "needs_target": true, "shield_damage": 2, "activation_heat": 30, "activation_energy": 5, "ore_mining_volume": 1, "activation_gfx_effect": "basic_laser_tool"}	2021-01-17 19:02:52.445344-05	648e4e7c-2328-413f-94cf-38a0a554d6ae	Module for new noob ship for player	2c7b5627-c8d3-40e3-b3f6-223375a7e46b	1	f
6f04e2ad-3e75-4096-89f5-a790190338a2	09172710-740c-4d1c-9fc0-43cb62e674e7	{"hp": 5, "rack": "b", "volume": 4, "cooldown": 7, "needs_target": false, "activation_heat": 65, "activation_energy": 15, "shield_boost_amount": 20, "activation_gfx_effect": "basic_shield_booster"}	2021-01-17 19:02:52.467195-05	648e4e7c-2328-413f-94cf-38a0a554d6ae	Module for new noob ship for player	2c7b5627-c8d3-40e3-b3f6-223375a7e46b	1	f
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
7b3b103c-2e7a-456b-a40e-243b818b3fc8	b37596b0-9841-4e27-a72c-eed954bb8c2d
80677c5d-7503-46f3-8a6f-6157318cdda0	cda4d030-ab22-45e3-b83e-2f5c233ce3fe
a2ad7499-b58a-4b6e-9e85-a58243337802	648e4e7c-2328-413f-94cf-38a0a554d6ae
\.


--
-- TOC entry 3216 (class 0 OID 25727)
-- Dependencies: 201
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, shield, armor, hull, fuel, heat, energy, shiptemplateid, dockedat_stationid, fitting, destroyed, destroyedat, cargobay_containerid, fittingbay_containerid, remaxdirty, trash_containerid) FROM stdin;
1248859c-8266-4aa1-b784-f8dd82caf026	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	b37596b0-9841-4e27-a72c-eed954bb8c2d	48440.9354921281774	-22300.9988440157395	2021-01-17 18:52:36.169252-05	asdf's Starter Ship	Sparrow	0	6.00617997963371166e-89	-3.1371035027552126e-89	209	244	135	295	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "102db1aa-ff16-43e7-a671-bed4535d7299", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "531e80ea-455e-47a6-b896-0071338b16d7", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "cda043ad-dfef-47cf-8fe7-996ad09124ca", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_id": "f2e8acbb-6ea5-4837-bdc8-05a59e3eb8b7", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "0b2ce719-4b04-4648-9b4d-df6310959584", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	48b28e75-01e1-443e-97eb-f2b2cc8073b0	8cf5bf0f-d44d-4619-a489-e8d0e3ea7a85	f	b6dfc5a9-e9a5-46c5-b2ce-a9e7080fbe69
6ca2cbfd-d0b9-4b99-8244-ca085c22fdbb	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	cda4d030-ab22-45e3-b83e-2f5c233ce3fe	36426.3381159103737	-16025.6232074325362	2021-01-17 18:57:02.610964-05	111's Starter Ship	Sparrow	0	-3.15777765492855925e-74	1.64934706848437887e-74	209	244	135	295	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	\N	{"a_rack": [{"item_id": "081ea8f3-c85d-409c-88a6-964b4fbb4395", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "0cbbc87e-54f3-4abe-be16-a1d2bb8dd5f1", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}], "b_rack": [{"item_id": "babfeb23-ebd5-49f6-b473-1501c1be546c", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}], "c_rack": [{"item_id": "2907bb02-fd20-437b-a2bc-7f81fdd91b39", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "64f25d5c-8b76-452f-bee0-b81c88295092", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	c3f7a739-ac00-440e-b1cd-221c858e9df7	4dd6b7e1-bf69-42a4-a695-74325c0768d4	f	a719d8f6-e4ec-49d6-8175-f8cf1db19207
fc64f4a1-79fb-427c-bc95-7f81da58a0e9	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	648e4e7c-2328-413f-94cf-38a0a554d6ae	24771.7956328638429	-9938.30877953488016	2021-01-17 19:02:52.490243-05	ppp's Starter Ship	Sparrow	188.126500442953471	0	0	209	169	135	295	0	138	8d9e032c-d9b1-4a36-8bbf-1448fa60a09a	cf07bba9-90b2-4599-b1e3-84d797a67f0a	{"a_rack": [{"item_id": "dcf866d4-c3a3-46c6-918b-c8bc2ea0aeef", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "fb96c08d-09cb-40ff-9293-33e964363761", "item_type_id": "9d1014c5-3422-4a0f-9839-f585269b4b16"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "b_rack": [{"item_id": "6f04e2ad-3e75-4096-89f5-a790190338a2", "item_type_id": "09172710-740c-4d1c-9fc0-43cb62e674e7"}, {"item_id": "00000000-0000-0000-0000-000000000000", "item_type_id": "00000000-0000-0000-0000-000000000000"}], "c_rack": [{"item_id": "5c12a7a2-2f0c-4e62-8a74-daf5c049ffcd", "item_type_id": "b481a521-1b12-4ffa-ac2f-4da015036f7f"}, {"item_id": "3dfc6052-44d5-4d6e-ac3d-0778ce33f7f9", "item_type_id": "c311df30-c21e-4895-acb0-d8808f99710e"}]}	f	\N	43315ba6-2040-414d-843a-7dca5c7dfd08	2c7b5627-c8d3-40e3-b3f6-223375a7e46b	f	b7aa4c01-4168-419c-87ca-c7f1855a8311
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
b37596b0-9841-4e27-a72c-eed954bb8c2d	asdf	7ab518aea3a72084bf1f1e1ddeb7fa25563a47b73b643452960187b6405b72f1	2021-01-17 18:52:36.08725-05	0	1248859c-8266-4aa1-b784-f8dd82caf026	49f06e89-29fb-4a02-a034-4b5d0702adac
cda4d030-ab22-45e3-b83e-2f5c233ce3fe	111	a6f3c209d805744c6a6f8895d6b15e24a424aff9287e2dde56755ea2d968b6f1	2021-01-17 18:57:02.541376-05	0	6ca2cbfd-d0b9-4b99-8244-ca085c22fdbb	49f06e89-29fb-4a02-a034-4b5d0702adac
648e4e7c-2328-413f-94cf-38a0a554d6ae	ppp	1615e807844001b17b2b9f8bc92b85f5ef70823ec4ca405d80da1793582484a2	2021-01-17 19:02:52.410437-05	0	fc64f4a1-79fb-427c-bc95-7f81da58a0e9	49f06e89-29fb-4a02-a034-4b5d0702adac
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


-- Completed on 2021-01-17 19:07:59 EST

--
-- PostgreSQL database dump complete
--

