--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2020-07-26 21:12:46

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
    shield double precision NOT NULL,
    armor double precision NOT NULL,
    hull double precision NOT NULL,
    fuel double precision NOT NULL,
    heat double precision NOT NULL,
    energy double precision NOT NULL,
    shiptemplateid uuid NOT NULL
);


ALTER TABLE public.ships OWNER TO postgres;

--
-- TOC entry 212 (class 1259 OID 49217)
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
    shiptypeid uuid NOT NULL
);


ALTER TABLE public.shiptemplates OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 49210)
-- Name: shiptypes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shiptypes (
    id uuid NOT NULL,
    name character varying(64) NOT NULL
);


ALTER TABLE public.shiptypes OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 49192)
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
-- TOC entry 209 (class 1259 OID 41032)
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
-- TOC entry 2773 (class 2606 OID 24613)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 2777 (class 2606 OID 24620)
-- Name: ships ships_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT ships_pkey PRIMARY KEY (id);


--
-- TOC entry 2791 (class 2606 OID 49236)
-- Name: shiptemplates shiptemplates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT shiptemplates_pkey PRIMARY KEY (id);


--
-- TOC entry 2787 (class 2606 OID 49216)
-- Name: shiptypes shiptypes_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_name_uq UNIQUE (name);


--
-- TOC entry 2789 (class 2606 OID 49214)
-- Name: shiptypes shiptypes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptypes
    ADD CONSTRAINT shiptypes_pkey PRIMARY KEY (id);


--
-- TOC entry 2785 (class 2606 OID 49199)
-- Name: universe_jumpholes universe_jumphole_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT universe_jumphole_pk PRIMARY KEY (id);


--
-- TOC entry 2781 (class 2606 OID 41026)
-- Name: universe_planets universe_planet_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT universe_planet_pk PRIMARY KEY (id);


--
-- TOC entry 2760 (class 2606 OID 24589)
-- Name: universe_regions universe_regions_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_name_uq UNIQUE (regionname);


--
-- TOC entry 2762 (class 2606 OID 24587)
-- Name: universe_regions universe_regions_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_regions
    ADD CONSTRAINT universe_regions_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2779 (class 2606 OID 41009)
-- Name: universe_stars universe_star_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT universe_star_pk PRIMARY KEY (id);


--
-- TOC entry 2783 (class 2606 OID 41039)
-- Name: universe_stations universe_station_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT universe_station_pk PRIMARY KEY (id);


--
-- TOC entry 2764 (class 2606 OID 24596)
-- Name: universe_systems universe_systems_name_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_name_uq UNIQUE (systemname);


--
-- TOC entry 2766 (class 2606 OID 24594)
-- Name: universe_systems universe_systems_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT universe_systems_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2775 (class 2606 OID 24615)
-- Name: sessions userid_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT userid_uq UNIQUE (userid);


--
-- TOC entry 2769 (class 2606 OID 24606)
-- Name: users users_pk_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk_uq PRIMARY KEY (id);


--
-- TOC entry 2771 (class 2606 OID 24608)
-- Name: users users_username_uq; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_uq UNIQUE (username);


--
-- TOC entry 2767 (class 1259 OID 24641)
-- Name: fki_fk_users_ships; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_fk_users_ships ON public.users USING btree (current_shipid);


--
-- TOC entry 2796 (class 2606 OID 49242)
-- Name: ships fk_ships_shiptemplates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_shiptemplates FOREIGN KEY (shiptemplateid) REFERENCES public.shiptemplates(id) NOT VALID;


--
-- TOC entry 2794 (class 2606 OID 24621)
-- Name: ships fk_ships_systems; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_systems FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2795 (class 2606 OID 24626)
-- Name: ships fk_ships_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ships
    ADD CONSTRAINT fk_ships_users FOREIGN KEY (userid) REFERENCES public.users(id);


--
-- TOC entry 2802 (class 2606 OID 49237)
-- Name: shiptemplates fk_shiptemplates_shiptypes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shiptemplates
    ADD CONSTRAINT fk_shiptemplates_shiptypes FOREIGN KEY (shiptypeid) REFERENCES public.shiptypes(id) NOT VALID;


--
-- TOC entry 2792 (class 2606 OID 24597)
-- Name: universe_systems fk_system_region; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_systems
    ADD CONSTRAINT fk_system_region FOREIGN KEY (regionid) REFERENCES public.universe_regions(id);


--
-- TOC entry 2793 (class 2606 OID 24636)
-- Name: users fk_users_ships; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_ships FOREIGN KEY (current_shipid) REFERENCES public.ships(id);


--
-- TOC entry 2801 (class 2606 OID 49205)
-- Name: universe_jumpholes jumphole_out_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_out_fk FOREIGN KEY (out_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2800 (class 2606 OID 49200)
-- Name: universe_jumpholes jumphole_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_jumpholes
    ADD CONSTRAINT jumphole_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2798 (class 2606 OID 41027)
-- Name: universe_planets planet_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_planets
    ADD CONSTRAINT planet_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2797 (class 2606 OID 41010)
-- Name: universe_stars star_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stars
    ADD CONSTRAINT star_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


--
-- TOC entry 2799 (class 2606 OID 41040)
-- Name: universe_stations station_system_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.universe_stations
    ADD CONSTRAINT station_system_fk FOREIGN KEY (universe_systemid) REFERENCES public.universe_systems(id);


-- Completed on 2020-07-26 21:12:46

--
-- PostgreSQL database dump complete
--

