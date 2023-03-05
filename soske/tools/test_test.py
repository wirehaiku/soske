"""
Tests for 'soske.tools.test'.
"""

import sqlite3

from soske.tools import sqls, test


def test_schema():
    # setup
    conn = sqlite3.connect(":memory:")

    # success
    conn.executescript(sqls.PRAGMA + sqls.SCHEMA + test.TEST_SCHEMA)
    curs = conn.execute("select count(*) from SQLITE_SCHEMA")
    assert curs.fetchone()[0] != 0
