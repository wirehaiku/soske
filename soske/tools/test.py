"""
Unit testing definitions and helpers.
"""

import sqlite3

import pytest
from soske.tools import sqls

TEST_SCHEMA = """
insert into Keys values ('alpha', false, 1000);
insert into Vals values ('alpha', 'Alpha one.', 1001);
insert into Vals values ('alpha', 'Alpha two.', 1002);

insert into Keys values ('bravo', false, 2000); 
insert into Vals values ('bravo', 'Bravo one.', 2001);

insert into Keys values ('dead_key', true, 3000);
insert into Vals values ('dead_key', 'Dead key.', 3001);

insert into Keys values ('empty_key', false, 4000);
"""


@pytest.fixture(scope="function")
def conn():
    conn = sqlite3.connect(":memory:")
    conn.row_factory = sqlite3.Row
    conn.executescript(sqls.PRAGMA + sqls.SCHEMA + TEST_SCHEMA)
    return conn
