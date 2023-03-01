// Package sqls implements SQL schema definitions.
package sqls

// BasePragma is the baseline required database pragma.
const BasePragma = `
pragma encoding = 'UTF-8';
pragma foreign_keys = ON;
`

// BaseSchema is the baseline required database schema.
const BaseSchema = `
create table Keys (
    name text not null unique,
    init int  not null default (unixepoch())
);

create table Vals (
    key  text not null references Keys(name),
    body text not null,
    init int  not null default (unixepoch())
);
`

// ProdSchema is the default schema for production applications.
const ProdSchema = BaseSchema + `
insert into Keys (name) values ("readme");
insert into Vals (key, body) values ("readme", "Welcome to Soske!");
`

// TestSchema is the default schema for testing applications.
const TestSchema = BaseSchema + `
insert into Keys (name, init) values ("alpha", 100);
insert into Keys (name, init) values ("bravo", 200);
insert into Keys (name, init) values ("charlie", 300);
insert into Vals (key, body, init) values ("alpha", "Alpha one.", 400);
insert into Vals (key, body, init) values ("bravo", "Alpha two.", 500);
insert into Vals (key, body, init) values ("bravo", "Bravo one.", 600);
`
