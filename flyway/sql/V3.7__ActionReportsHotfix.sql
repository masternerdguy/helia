CREATE OR REPLACE VIEW public.vw_actionreports_summary
 AS
 SELECT s.victim_isnpc,
    s.victim_name,
    s.victim_shiptemplatename,
    f.ticker AS victim_ticker,
    s.solarsystemname,
    s.regionname,
    COALESCE(jsonb_array_length(ids_arr), 0) as parties,
    s."timestamp",
    s.search_userid,
    s.id
   FROM ( SELECT (((ar.actionreport ->> 'header'::text)::jsonb) ->> 'isNPC'::text)::boolean AS victim_isnpc,
            ((ar.actionreport ->> 'header'::text)::jsonb) ->> 'victimName'::text AS victim_name,
            ((ar.actionreport ->> 'header'::text)::jsonb) ->> 'victimShipTemplateName'::text AS victim_shiptemplatename,
            (((ar.actionreport ->> 'header'::text)::jsonb) ->> 'victimFactionID'::text)::uuid AS victim_factionid,
            ((ar.actionreport ->> 'header'::text)::jsonb) ->> 'solarSystemName'::text AS solarsystemname,
            ((ar.actionreport ->> 'header'::text)::jsonb) ->> 'regionName'::text AS regionname,
            ( SELECT jsonb_agg(t.j ->> 'userID'::text) AS jsonb_agg
                   FROM jsonb_array_elements((ar.actionreport ->> 'involvedParties'::text)::jsonb) t(j)) AS ids_arr,
            ar."timestamp",
            r.userid AS search_userid,
            ar.id
           FROM ( SELECT q.id,
                    q.userid
                   FROM ( SELECT vwa.id,
                            btrim(vwa.partyid, '""'::text)::uuid AS userid
                           FROM vw_actionreports_involvedpartyids vwa
                        UNION
                         SELECT actionreports.id,
                            btrim(((actionreports.actionreport ->> 'header'::text)::jsonb) ->> 'victimID'::text, '""'::text)::uuid AS userid
                           FROM actionreports) q
                  GROUP BY q.id, q.userid) r
             JOIN actionreports ar ON r.id = ar.id) s
     LEFT JOIN factions f ON f.id = s.victim_factionid;

ALTER TABLE public.vw_actionreports_summary
    OWNER TO heliaagent;
