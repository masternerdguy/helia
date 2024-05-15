ALTER TABLE IF EXISTS public.users
    ADD COLUMN is_trueoriginal boolean NOT NULL DEFAULT false;

UPDATE public.users SET is_trueoriginal = true;
