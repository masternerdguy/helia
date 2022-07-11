
CREATE PROCEDURE public.sp_purgeallusers()
    LANGUAGE sql
    AS $$-- delete user data
update users set current_shipid = null;
select id from users;
delete from ships where userid in (select id from users);
delete from sellorders where seller_userid in (select id from users);
delete from sellorders where buyer_userid in (select id from users);
delete from sessions;
delete from schematicruns;

-- delete items
delete from items;

-- move non-npc users into neutral faction so they can be deleted
update users set current_factionid = '42b937ad-0000-46e9-9af9-fc7dbf878e6a' where isnpc = 'f';

-- null owners on non-npc factions so their owners can be deleted
update factions set ownerid = null where isnpc = false and id != '42b937ad-0000-46e9-9af9-fc7dbf878e6a';

-- delete users
delete from users;

-- delete containers
delete from containers;

-- delete player factions
delete from factions where isnpc = false and id != '42b937ad-0000-46e9-9af9-fc7dbf878e6a';

-- delete action reports
delete from actionreports;$$;


ALTER PROCEDURE public.sp_purgeallusers() OWNER TO heliaagent;
