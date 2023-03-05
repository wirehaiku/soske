"""
Tests for 'soske.tools.data'.
"""

from soske.tools import data, sqls
from soske.tools.test import conn


def test_open(tmp_path):
    # setup
    path = str(tmp_path / "test.db")

    # success - pragma & schema
    conn = data.open(path, sqls.PRAGMA, "create table Tests(t)")
    curs = conn.execute("select count(*) from SQLITE_SCHEMA")
    assert curs.fetchone()["count(*)"] != 0

    # success - pragma only
    conn = data.open(path, sqls.PRAGMA, "drop table Tests")
    curs = conn.execute("select count(*) from SQLITE_SCHEMA")
    assert curs.fetchone()["count(*)"] != 0


def test_exec(conn):
    # success
    data.exec(conn, "insert into Keys values ('test', false, 123)")
    curs = conn.execute("select * from Keys where name=?", ["test"])
    assert curs.fetchone()["name"] == "test"


def test_row(conn):
    # success
    row = data.row(conn, "select * from Keys where name=?", "alpha")
    assert row["name"] == "alpha"


def test_rows(conn):
    # success
    rows = data.rows(conn, "select * from Keys where name=?", "alpha")
    assert [row["name"] for row in rows] == ["alpha"]
