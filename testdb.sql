--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2020-06-28 20:08:03

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
    theta integer DEFAULT 0 NOT NULL,
    vel_x double precision DEFAULT 0 NOT NULL,
    vel_y double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.ships OWNER TO postgres;

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
-- TOC entry 2860 (class 0 OID 24609)
-- Dependencies: 205
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (id, userid) FROM stdin;
6ff30dd2-6093-4a9d-9f69-1ce5cb25d1ec	502b5a79-f5d4-46ff-8e94-653c0778bb6e
38c95eda-861f-40ea-adb2-f023b14a9624	298011ef-0eb1-42d3-9aed-28e7166da350
8178feb8-7d5e-4f55-88ff-4660c391260f	ded647c9-2799-4fb6-9887-a32ca11d0c4c
\.


--
-- TOC entry 2861 (class 0 OID 24616)
-- Dependencies: 206
-- Data for Name: ships; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ships (id, universe_systemid, userid, pos_x, pos_y, created, shipname, texture, theta, vel_x, vel_y) FROM stdin;
272041ee-f373-49ea-8e37-379ed9fa84da	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	502b5a79-f5d4-46ff-8e94-653c0778bb6e	100	150	08:27:26.737175-04	nwiehoff's Starter Ship	Mass Testing Brick	0	0	0
a50c0ca6-004f-47c1-9447-c9b8cdea2e67	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	ded647c9-2799-4fb6-9887-a32ca11d0c4c	-65	-300	07:41:30.998232-04	masternerdguy's Starter Ship	Mass Testing Brick	0	0	0
c5f7d737-c8ed-42ce-98f0-25dbf3433196	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	194ea29b-b0ec-4d80-8f42-be72dd582736	0	0	18:15:18.72366-04	56yutjhgf's Starter Ship	Mass Testing Brick	0	0	0
f98d0360-09b8-4b27-9cb8-72b6481509ce	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	298011ef-0eb1-42d3-9aed-28e7166da350	0	0	18:15:39.257237-04	yujinhgr45's Starter Ship	Mass Testing Brick	0	0	0
4dccf469-9397-44b1-8884-8f08d9b519d6	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	8f2dc811-3a4f-4d2b-b051-ebc671221814	0	0	18:42:43.875555-04	67u8iyjgnhfbd's Starter Ship	Mass Testing Brick	0	0	0
0d82574e-a16c-42f9-a1c7-d1db9546a338	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	7aa62345-9e11-4dcd-8110-b9cc2e684e03	0	0	18:43:46.257894-04	67u8iyjgnhfbdw's Starter Ship	Mass Testing Brick	0	0	0
d4226474-de7d-4606-81e1-844b60b37531	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	56e974fd-5e95-4467-9e28-746fda9fbec8	0	0	18:44:11.861133-04	e456tyrhtrge's Starter Ship	Mass Testing Brick	0	0	0
13b163d6-22e1-4f76-8d1c-2b2c8fc600e1	1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	5838b541-601d-4652-9d86-0a4f61068f2e	0	0	18:46:13.203242-04	r56tyrjgnf's Starter Ship	Mass Testing Brick	0	0	0
\.


--
-- TOC entry 2857 (class 0 OID 24583)
-- Dependencies: 202
-- Data for Name: universe_regions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_regions (id, regionname) FROM stdin;
bfca1f47-e182-4b4d-8632-48d8ead08647	The Core
\.


--
-- TOC entry 2858 (class 0 OID 24590)
-- Dependencies: 203
-- Data for Name: universe_systems; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.universe_systems (id, systemname, regionid) FROM stdin;
1d4e0a33-9f67-4f24-8b7b-1af4d5aa2ef1	Test System	bfca1f47-e182-4b4d-8632-48d8ead08647
edf08406-0879-4141-8af1-f68d32e31c8d	Another System	bfca1f47-e182-4b4d-8632-48d8ead08647
\.


--
-- TOC entry 2859 (class 0 OID 24602)
-- Dependencies: 204
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, hashpass, registered, banned, current_shipid) FROM stdin;
502b5a79-f5d4-46ff-8e94-653c0778bb6e	nwiehoff	02c4bf7d7e35c6bab999ac03ece60b8586a27f7ecd4830983b138b74262bf3f9	2020-06-16 08:27:26.545781-04	0	272041ee-f373-49ea-8e37-379ed9fa84da
ded647c9-2799-4fb6-9887-a32ca11d0c4c	masternerdguy	71e58265a4cc0edc6b4654593ec35094222c88edbb51e9439b47ba86999b3f86	2020-06-18 07:41:30.776233-04	0	a50c0ca6-004f-47c1-9447-c9b8cdea2e67
1ff92074-68b6-4a74-ab52-315101575e79	rtyhfbgd	6d923d645872b8c15135eb8a5ea2bfd6a7cfc208ed25afa9265ff5a72b28602a	2020-06-28 18:13:42.938684-04	0	\N
194ea29b-b0ec-4d80-8f42-be72dd582736	56yutjhgf	e5776ffa00458184f796ab025c043fbd3b8556d9947820fe328de75fbafaffde	2020-06-28 18:15:18.643667-04	0	c5f7d737-c8ed-42ce-98f0-25dbf3433196
298011ef-0eb1-42d3-9aed-28e7166da350	yujinhgr45	20cc6c8a9135011b81116d32f64db810f612b552423dd0ea2dbd4a400613e7e6	2020-06-28 18:15:39.176236-04	0	f98d0360-09b8-4b27-9cb8-72b6481509ce
8f2dc811-3a4f-4d2b-b051-ebc671221814	67u8iyjgnhfbd	6c06ea295fc0960be027c8adbfb947e97702d1fda5772c359ccdf49f6404395f	2020-06-28 18:42:43.798555-04	0	4dccf469-9397-44b1-8884-8f08d9b519d6
7aa62345-9e11-4dcd-8110-b9cc2e684e03	67u8iyjgnhfbdw	ad3a3942ae14a1b5d88e9892641c1a6892e616b25f4470a2456b9ebcf247bd3e	2020-06-28 18:43:46.165894-04	0	0d82574e-a16c-42f9-a1c7-d1db9546a338
56e974fd-5e95-4467-9e28-746fda9fbec8	e456tyrhtrge	d6311e908d1ee5e964cce536e272e2701c7f07ef64bc24a85043934d6a5310ff	2020-06-28 18:44:11.779132-04	0	d4226474-de7d-4606-81e1-844b60b37531
5838b541-601d-4652-9d86-0a4f61068f2e	r56tyrjgnf	01740c51823ee93bd4d8294c772a33f62f12252f7e7278aa4ae7f210c43a8414	2020-06-28 18:46:13.118242-04	0	13b163d6-22e1-4f76-8d1c-2b2c8fc600e1
\.


--
-- TOC entry 2722 (class 2606 OID 24613)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2726 (class 2606 OID 24620)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 2709 (class 2606 OID 24589)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 2711 (class 2606 OID 24587)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2713 (class 2606 OID 24596)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 2715 (class 2606 OID 24594)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2724 (class 2606 OID 24615)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 2718 (class 2606 OID 24606)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2720 (class 2606 OID 24608)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 2716 (class 1259 OID 24641)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 2729 (class 2606 OID 24621)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2730 (class 2606 OID 24626)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 2727 (class 2606 OID 24597)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 2728 (class 2606 OID 24636)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


-- Completed on 2020-06-28 20:08:04

--
-- PostgreSQL database dump complete
--

