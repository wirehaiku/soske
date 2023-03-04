package soske

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// TestSchema is the default database testing schema.
const TestSchema = `
insert into Keys values ('alpha', false, 1000);
insert into Keys values ('bravo', false, 2000);
insert into Keys values ('dead',  true,  3000);
insert into Keys values ('empty', false, 4000);

insert into Vals values ('alpha', 'Alpha one.', 1001);
insert into Vals values ('alpha', 'Alpha two.', 1002);
insert into Vals values ('bravo', 'Bravo one.', 2001);
insert into Vals values ('dead',  'Dead key.',  3001);
`

// TestDB returns an in-memory database populated with test data.
func TestDB() *sqlx.DB {
	Debug = true
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(Pragma + Schema + TestSchema)
	return db
}
