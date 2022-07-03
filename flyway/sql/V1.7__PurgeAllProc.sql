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

ALTER PROCEDURE public.sp_purgeallusers() OWNER TO heliaagent;
