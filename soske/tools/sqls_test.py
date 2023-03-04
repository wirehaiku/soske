"""
Tests for 'soske.tools.sqls'.
"""

import sqlite3

from soske.tools import sqls


def test_module():
    # setup
    db = sqlite3.connect(":memory:")

    # success
    db.executescript(sqls.PRAGMA + sqls.SCHEMA)
    curs = db.execute("select count(*) from SQLITE_SCHEMA")
    assert curs.fetchone()[0] != 0
