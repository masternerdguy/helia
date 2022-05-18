-- View: public.vw_modules_needsschematics

-- DROP VIEW public.vw_modules_needsschematics;

CREATE OR REPLACE VIEW public.vw_modules_needsschematics
 AS
 SELECT itemtypes.id,
    itemtypes.family,
    itemtypes.name,
    itemtypes.meta
   FROM itemtypes
  WHERE NOT (itemtypes.id IN ( SELECT processoutputs.itemtypeid
           FROM processoutputs
          WHERE (processoutputs.processid IN ( SELECT ((itemtypes_1.meta::json -> 'industrialmarket'::text) ->> 'process_id'::text)::uuid AS proccessid
                   FROM itemtypes itemtypes_1
                  WHERE itemtypes_1.family::text = 'schematic'::text)))) AND itemtypes.name::text !~~ '%`%'::text AND (itemtypes.family::text <> ALL (ARRAY['nothing'::character varying, 'ore'::character varying, 'ice'::character varying, 'trade_good'::character varying, 'ship'::character varying, 'schematic'::character varying, 'power_cell'::character varying, 'depleted_cell'::character varying, 'widget'::character varying]::text[]));

ALTER TABLE public.vw_modules_needsschematics
    OWNER TO heliaagent;

