-- add flag for users to be marked as developers
ALTER TABLE IF EXISTS public.users
    ADD COLUMN isdev boolean NOT NULL DEFAULT 'f';

-- make me a dev
update users set isdev = true where id = 'ca23e222-bf73-4f40-968f-715c033f12b2'
