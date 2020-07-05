--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2020-07-04 21:04:48

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
-- TOC entry 205 (class 1259 OID 24609)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    userid uuid NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 24616)
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
    accel double precision DEFAULT 0 NOT NULL,
    radius double precision DEFAULT 0 NOT NULL,
    mass double precision DEFAULT 0 NOT NULL,
    turn double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.ships OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 41019)
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
-- TOC entry 202 (class 1259 OID 24583)
-- Name: universe_regions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_regions (
    id uuid NOT NULL,
    regionname character varying(32) NOT NULL
);


ALTER TABLE public.universe_regions OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 41005)
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
-- TOC entry 203 (class 1259 OID 24590)
-- Name: universe_systems; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.universe_systems (
    id uuid NOT NULL,
    systemname character varying(32) NOT NULL,
    regionid uuid NOT NULL
);


ALTER TABLE public.universe_systems OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 24602)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    username character varying(32) NOT NULL,
    hashpass character(64) NOT NULL,
    registered timestamp with time zone NOT NULL,
    banned bit(1) NOT NULL,
    current_shipid uuid
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 2884 (class 0 OID 24609)
-- Dependencies: 205
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
8a210e3d-766d-4eae-af42-4b2628e885c4	ec5c12e6-04b1-4ea3-b1c7-a8e5615b012a
229a9af0-db46-4b41-8b23-a39310ca2373	502b5a79-f5d4-46ff-8e94-653c0778bb6e
ce2f7901-8896-41a5-a641-4b24ddd212de	ded647c9-2799-4fb6-9887-a32ca11d0c4c
38c95eda-861f-40ea-adb2-f023b14a9624	298011ef-0eb1-42d3-9aed-28e7166da350
a9de6c9b-3d4b-4b8e-9d31-5dd96a57d086	57194133-619d-4239-bc10-ee4a2d9fdd62
\.


--
-- TOC entry 2885 (class 0 OID 24616)
-- Dependencies: 206
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y, accel, radius, mass, turn) FROM stdin;
272041ee-f373-49ea-8e37-379ed9fa84da	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	502b5a79-f5d4-46ff-8e94-653c0778bb6e	-1.4330223655533643	573.9666142622724	08:27:26.737175-04	nwiehoff's Starter Ship	Goanna	38.93751964251964	5.3854493871049e-36	-3.144057927553592e-36	0.1	167.5	10000	1
a50c0ca6-004f-47c1-9447-c9b8cdea2e67	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	ded647c9-2799-4fb6-9887-a32ca11d0c4c	-95.26682399238562	-30.077679998955638	07:41:30.998232-04	masternerdguy's Starter Ship	Mass Testing Brick	225	1.3057592317242641e-208	9.068857015040889e-209	1	12.5	100	10
390387f4-794b-4e28-b7e5-9be1c282bf80	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	57194133-619d-4239-bc10-ee4a2d9fdd62	-427.23521394489956	530.3861094507423	16:06:48.42982-04	asdf's Starter Ship	Mass Testing Brick	0	-4.975612916091584e-203	8.663208787828185e-202	1	12.5	100	10
\.


--
-- TOC entry 2887 (class 0 OID 41019)
-- Dependencies: 208
-- Data for Name: universe_planets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_planets (id, universe_systemid, planetname, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
695f28ec-941d-405c-b1ca-fbeace169d92	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test Planet	-1000	-1000	vh_unshaded\\planet03	100000	1000000	23
e20d3b80-f44f-4e16-91c7-d5489a95bf4a	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Another Planet	60000	450000	vh_unshaded\\planet09	8000	2000000	112
\.


--
-- TOC entry 2881 (class 0 OID 24583)
-- Dependencies: 202
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 2886 (class 0 OID 41005)
-- Dependencies: 207
-- Data for Name: universe_stars; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_stars (id, universe_systemid, pos_x, pos_y, texture, radius, mass, theta) FROM stdin;
058618fc-1bb3-42a6-a240-14922782de41	edf08406-0879-4141-8af1-f68d32e31c8d	25	-25	vh_main_sequence\\star_yellow03	600000	30000000000	0
42805a07-9f38-484c-9aaf-9ec55e74f725	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	10000000	-500000	vh_main_sequence\\star_red01	200000	1000000000	0
\.


--
-- TOC entry 2882 (class 0 OID 24590)
-- Dependencies: 203
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 2883 (class 0 OID 24602)
-- Dependencies: 204
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid) FROM stdin;
502b5a79-f5d4-46ff-8e94-653c0778bb6e	nwiehoff	02c4bf7d7e35c6bab999ac03ece60b8586a27f7ecd4830983b138b74262bf3f9	2020-06-16 08:27:26.545781-04	0	272041ee-f373-49ea-8e37-379ed9fa84da
ded647c9-2799-4fb6-9887-a32ca11d0c4c	masternerdguy	71e58265a4cc0edc6b4654593ec35094222c88edbb51e9439b47ba86999b3f86	2020-06-18 07:41:30.776233-04	0	a50c0ca6-004f-47c1-9447-c9b8cdea2e67
57194133-619d-4239-bc10-ee4a2d9fdd62	asdf	7ab518aea3a72084bf1f1e1ddeb7fa25563a47b73b643452960187b6405b72f1	2020-07-03 16:06:48.339821-04	0	390387f4-794b-4e28-b7e5-9be1c282bf80
\.


--
-- TOC entry 2740 (class 2606 OID 24613)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2744 (class 2606 OID 24620)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 2748 (class 2606 OID 41026)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 2727 (class 2606 OID 24589)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 2729 (class 2606 OID 24587)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2746 (class 2606 OID 41009)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 2731 (class 2606 OID 24596)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 2733 (class 2606 OID 24594)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2742 (class 2606 OID 24615)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 2736 (class 2606 OID 24606)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2738 (class 2606 OID 24608)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 2734 (class 1259 OID 24641)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 2751 (class 2606 OID 24621)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2752 (class 2606 OID 24626)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 2749 (class 2606 OID 24597)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 2750 (class 2606 OID 24636)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 2754 (class 2606 OID 41027)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2753 (class 2606 OID 41010)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2020-07-04 21:04:48

--
-- PostgreSQL database dump complete
--

