-- Table: public.universe_artifacts

-- DROP TABLE IF EXISTS public.universe_artifacts;

CREATE TABLE IF NOT EXISTS public.universe_artifacts
(
    id uuid NOT NULL,
    universe_systemid uuid NOT NULL,
    artifactname character varying(64) COLLATE pg_catalog."default" NOT NULL,
    pos_x double precision NOT NULL,
    pos_y double precision NOT NULL,
    texture character varying(64) COLLATE pg_catalog."default" NOT NULL,
    radius double precision NOT NULL DEFAULT 0,
    mass double precision NOT NULL DEFAULT 0,
    theta double precision NOT NULL DEFAULT 0,
    meta jsonb NOT NULL DEFAULT '{}'::jsonb,
    CONSTRAINT universe_artifact_pk PRIMARY KEY (id),
    CONSTRAINT artifact_system_fk FOREIGN KEY (universe_systemid)
        REFERENCES public.universe_systems (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.universe_artifacts
    OWNER to heliaagent;
