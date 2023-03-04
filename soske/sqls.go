package soske

// Pragma is the default database pragma.
const Pragma = `
pragma encoding = "UTF-8";
pragma foreign_keys = ON;
`

// Schema is the default database schema.
const Schema = `
create table Keys (
    name text    not null unique,
    dead boolean not null default false,
    init integer not null default (unixepoch())
);

create table Vals (
    key  text    not null references Keys(name),
    body text    not null,
    init integer not null default (unixepoch())
);

create view DeadKeys as
    select * from Keys where dead=true
    order by name asc;

create view GoodKeys as
    select * from Keys where (select count(*) from Vals where key=name)>0 and dead=false
    order by name asc;

create view GoodVals as 
    select * from Vals where (select dead from Keys where name=key)=false
    order by init desc;

create view ZeroKeys as
    select * from Keys where (select count(*) from Vals where key=name)=0
    order by name asc;
`
