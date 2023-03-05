"""
Database pragma and schema definitions.
"""

PRAGMA = """
pragma encoding = "UTF-8";
pragma foreign_keys = ON;
"""

SCHEMA = """
create table Keys (
    name text not null unique,
    dead bool not null,
    init real not null
);

create table Vals (
    key  text not null references Keys(name),
    body text not null,
    init real not null
);

create view DeadKeys as
    select * from Keys where dead=true
    order by name asc;

create view EmptyKeys as
    select * from Keys where (select count(*) from Vals where key=name)=0
    order by name asc;

create view GoodKeys as
    select * from Keys where (select count(*) from Vals where key=name)>0 and dead=false
    order by name asc;

create view GoodVals as 
    select * from Vals where (select dead from Keys where name=key)=false
    order by init desc;
"""

TEST_SCHEMA = """
insert into Keys values ('alpha', false,        1000.1);
insert into Vals values ('alpha', 'Alpha one.', 1000.2);
insert into Vals values ('alpha', 'Alpha two.', 1000.3);

insert into Keys values ('bravo', false,        2000.1); 
insert into Vals values ('bravo', 'Bravo one.', 2000.2);

insert into Keys values ('dead_key', true,        3000.1);
insert into Vals values ('dead_key', 'Dead key.', 3000.2);

insert into Keys values ('empty_key', false, 4000.1);
"""
