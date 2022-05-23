-- View: public.vw_actionreports_involvedpartyids

-- DROP VIEW public.vw_actionreports_involvedpartyids;

CREATE OR REPLACE VIEW public.vw_actionreports_involvedpartyids
 AS
 SELECT r.id,
    r.length,
    jsonb_array_elements(r.ids_arr)::text AS partyid
   FROM ( SELECT q.id,
            jsonb_array_length(q.parties) AS length,
            jsonb_path_query_array(q.parties, '$."userID"'::jsonpath) AS ids_arr
           FROM ( SELECT actionreports.id,
                    (actionreports.actionreport ->> 'involvedParties'::text)::jsonb AS parties
                   FROM actionreports) q) r
  GROUP BY r.id, r.length, (jsonb_array_elements(r.ids_arr)::text);

ALTER TABLE public.vw_actionreports_involvedpartyids
    OWNER TO heliaagent;

