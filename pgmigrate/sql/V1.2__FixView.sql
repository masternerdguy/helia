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
                  WHERE itemtypes_1.family::text = 'schematic'::text)))) AND itemtypes.name::text !~~ '%`%'::text AND (itemtypes.family::text <> ALL (ARRAY['nothing'::character varying::text, 'ore'::character varying::text, 'repair_kit'::character varying::text, 'mod_kit'::character varying::text, 'ammunition'::character varying::text, 'fuel'::character varying::text, 'ice'::character varying::text, 'trade_good'::character varying::text, 'ship'::character varying::text, 'schematic'::character varying::text, 'power_cell'::character varying::text, 'depleted_cell'::character varying::text, 'widget'::character varying::text, 'outpost_kit'::character varying::text, 'gas'::character varying::text]));

ALTER TABLE public.vw_modules_needsschematics
    OWNER TO heliaagent;

