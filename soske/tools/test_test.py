"""
Tests for 'soske.tools.test'.
"""

import sqlite3

from soske.tools import sqls, test


def test_module():
    # setup
    db = sqlite3.connect(":memory:")

    # success
    db.executescript(sqls.PRAGMA + sqls.SCHEMA + test.TEST_SCHEMA)
    curs = db.execute("select count(*) from SQLITE_SCHEMA")
    assert curs.fetchone()[0] != 0
