-- add meta columns to celestials
ALTER TABLE IF EXISTS public.universe_asteroids
    ADD COLUMN meta jsonb NOT NULL DEFAULT '{}';

ALTER TABLE IF EXISTS public.universe_planets
    ADD COLUMN meta jsonb NOT NULL DEFAULT '{}';

ALTER TABLE IF EXISTS public.universe_stars
    ADD COLUMN meta jsonb NOT NULL DEFAULT '{}';
