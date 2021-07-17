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
    meta jsonb NOT NULL,
    ticker character varying(3) DEFAULT '???'::character varying NOT NULL,
    isnpc boolean DEFAULT false NOT NULL,
    isjoinable boolean DEFAULT false NOT NULL,
    canholdsov boolean DEFAULT false NOT NULL,
    isclosed boolean DEFAULT false NOT NULL,
    reputationsheet jsonb DEFAULT '{}'::jsonb NOT NULL
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

