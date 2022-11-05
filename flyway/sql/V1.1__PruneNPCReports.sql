-- PROCEDURE: public.sp_quarantineactionreports()

-- DROP PROCEDURE IF EXISTS public.sp_quarantineactionreports();

CREATE OR REPLACE PROCEDURE public.sp_quarantineactionreports(
	)
LANGUAGE 'sql'
AS $BODY$
-- delete copied NPC-only action reports from main action reports table
delete from actionreports where id in
(
	select id from actionreports_npc
);

-- copy NPC-only (no human involvement) action reports to quarantine table
INSERT INTO actionreports_npc 
(
	SELECT * FROM actionreports WHERE id in
	(
		 select id from (
			 select id, actionreport::text as txtreport from actionreports
		 ) q
		 where q.txtreport not like '%"isNPC": false%'
	)
);

-- delete copied NPC-only action reports from main action reports table again
delete from actionreports where id in
(
	select id from actionreports_npc
);

-- delete NPC-only action reports that are >30 days old
delete from actionreports_npc
where timestamp < (CURRENT_DATE - INTERVAL '30 DAY')::DATE
$BODY$;
ALTER PROCEDURE public.sp_quarantineactionreports()
    OWNER TO heliaagent;

-- delete 
