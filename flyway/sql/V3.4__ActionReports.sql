drop view if exists vw_flattenactionreportsbyparty;

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

-- View: public.vw_actionreports_summary

-- DROP VIEW public.vw_actionreports_summary;

CREATE OR REPLACE VIEW public.vw_actionreports_summary
 AS
 SELECT s.victim_isnpc,
    s.victim_name,
    s.victim_shiptemplatename,
    f.ticker AS victim_ticker,
    s.solarsystemname,
    s.regionname,
    s.parties,
    s."timestamp",
    s.search_userid,
    s.id
   FROM ( SELECT (((ar.actionreport ->> 'header'::text)::jsonb) ->> 'isNPC'::text)::boolean AS victim_isnpc,
            ((ar.actionreport ->> 'header'::text)::jsonb) ->> 'victimName'::text AS victim_name,
            ((ar.actionreport ->> 'header'::text)::jsonb) ->> 'victimShipTemplateName'::text AS victim_shiptemplatename,
            (((ar.actionreport ->> 'header'::text)::jsonb) ->> 'victimFactionID'::text)::uuid AS victim_factionid,
            ((ar.actionreport ->> 'header'::text)::jsonb) ->> 'solarSystemName'::text AS solarsystemname,
            ((ar.actionreport ->> 'header'::text)::jsonb) ->> 'regionName'::text AS regionname,
            jsonb_array_length(jsonb_path_query_array((ar.actionreport ->> 'involvedParties'::text)::jsonb, '$."userID"'::jsonpath)) AS parties,
            ar."timestamp",
            r.userid AS search_userid,
            ar.id
           FROM ( SELECT q.id,
                    q.userid
                   FROM ( SELECT vwa.id,
                            TRIM(BOTH '""'::text FROM vwa.partyid)::uuid AS userid
                           FROM vw_actionreports_involvedpartyids vwa
                        UNION
                         SELECT actionreports.id,
                            TRIM(BOTH '""'::text FROM ((actionreports.actionreport ->> 'header'::text)::jsonb) ->> 'victimID'::text)::uuid AS userid
                           FROM actionreports) q
                  GROUP BY q.id, q.userid) r
             JOIN actionreports ar ON r.id = ar.id) s
     LEFT JOIN factions f ON f.id = s.victim_factionid;

ALTER TABLE public.vw_actionreports_summary
    OWNER TO heliaagent;

