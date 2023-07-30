-- Table: public.outposttemplates

-- DROP TABLE IF EXISTS public.outposttemplates;

CREATE TABLE IF NOT EXISTS public.outposttemplates
(
    id uuid NOT NULL,
    created timestamp with time zone NOT NULL DEFAULT now(),
    outposttemplatename character varying(32) COLLATE pg_catalog."default" NOT NULL,
    texture character varying(32) COLLATE pg_catalog."default" NOT NULL DEFAULT 'Mass Testing Brick'::character varying,
    radius double precision NOT NULL DEFAULT 0,
    basemass double precision NOT NULL DEFAULT 0,
    baseshield double precision NOT NULL DEFAULT 0,
    baseshieldregen double precision NOT NULL DEFAULT 0,
    basearmor double precision NOT NULL DEFAULT 0,
    basehull double precision NOT NULL DEFAULT 0,
    itemtypeid uuid NOT NULL,
    wrecktexture character varying(16) COLLATE pg_catalog."default" NOT NULL DEFAULT 'basic-wreck'::character varying,
    explosiontexture character varying(16) COLLATE pg_catalog."default" NOT NULL DEFAULT 'basic_explosion'::character varying,
    CONSTRAINT outposttemplates_pkey PRIMARY KEY (id),
    CONSTRAINT fk_outposttemplates_itemtypes FOREIGN KEY (itemtypeid)
        REFERENCES public.itemtypes (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.outposttemplates
    OWNER to heliaagent;
-- Index: fki_fk_outposttemplates_itemtypes

-- DROP INDEX IF EXISTS public.fki_fk_outposttemplates_itemtypes;

CREATE INDEX IF NOT EXISTS fki_fk_outposttemplates_itemtypes
    ON public.outposttemplates USING btree
    (itemtypeid ASC NULLS LAST)
    TABLESPACE pg_default;

-- Table: public.outposts

-- DROP TABLE IF EXISTS public.outposts;

CREATE TABLE IF NOT EXISTS public.outposts
(
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    universe_stationid uuid NOT NULL,
    outpostname character varying(24) COLLATE pg_catalog."default" NOT NULL,
    pos_x double precision NOT NULL,
    pos_y double precision NOT NULL,
    theta double precision NOT NULL DEFAULT 0,
    userid uuid NOT NULL,
    shield double precision NOT NULL,
    armor double precision NOT NULL,
    hull double precision NOT NULL,
    wallet double precision NOT NULL DEFAULT 0,
    destroyed boolean NOT NULL DEFAULT false,
    destroyedat timestamp with time zone,
    outposttemplateid uuid NOT NULL,
    created timestamp with time zone NOT NULL,
    CONSTRAINT outpost_pk PRIMARY KEY (id),
    CONSTRAINT fk_outpost_users FOREIGN KEY (userid)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_outposts_stations FOREIGN KEY (universe_stationid)
        REFERENCES public.universe_stations (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT outpost_system_fk FOREIGN KEY (universe_systemid)
        REFERENCES public.universe_systems (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.outposts
    OWNER to heliaagent;

INSERT INTO public.itemfamilies(
	id, friendlyname, meta)
	VALUES ('outpost_kit', 'Outpost Construction Kit', '{}'::jsonb);

INSERT INTO public.itemtypes (id, family, name, meta) VALUES ('59851ea9-5f78-41c9-9cc2-2a7b1bbc6e72', 'outpost_kit', 'Test Outpost Please Ignore', '{"volume": 350775, "outposttemplateid": "188a3e34-0662-480a-8df8-d4b038e8a8c3"}');

INSERT INTO public.outposttemplates (id, created, outposttemplatename, texture, radius, basemass, baseshield, baseshieldregen, basearmor, basehull, itemtypeid, wrecktexture, explosiontexture) VALUES ('188a3e34-0662-480a-8df8-d4b038e8a8c3', '2023-05-06 20:41:21.336072-04', 'Test Template Please Ignore', 'kingdom-7', 588, 47283, 113293, 328, 782171, 588924, '59851ea9-5f78-41c9-9cc2-2a7b1bbc6e72', 'basic-wreck', 'basic_explosion');

-- Column: public.universe_stations.isoutpostshim

-- ALTER TABLE IF EXISTS public.universe_stations DROP COLUMN IF EXISTS isoutpostshim;

ALTER TABLE IF EXISTS public.universe_stations
    ADD COLUMN isoutpostshim boolean NOT NULL DEFAULT false;
