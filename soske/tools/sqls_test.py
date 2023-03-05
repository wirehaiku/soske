"""
Tests for 'soske.tools.sqls'.
"""

import sqlite3

from soske.tools import sqls


def test_schema():
    # setup
    conn = sqlite3.connect(":memory:")

    # success
    conn.executescript(sqls.PRAGMA + sqls.SCHEMA)
    curs = conn.execute("select count(*) from SQLITE_SCHEMA")
    assert curs.fetchone()[0] != 0
