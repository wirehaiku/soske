"""
Unit testing definitions and helpers.
"""

import sqlite3

import click.testing
import pytest
from soske.comms import app
from soske.tools import sqls

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


@pytest.fixture(scope="function")
def conn():
    """
    Return a connection to an in-memory database populated with test data.
    """

    conn = sqlite3.connect(":memory:")
    conn.row_factory = sqlite3.Row
    conn.executescript(sqls.PRAGMA + sqls.SCHEMA + TEST_SCHEMA)
    return conn


@pytest.fixture(scope="function")
def run():
    """
    Return a function that returns the status code and output of a Click command.
    """

    def cmdfunc(conn: sqlite3.Connection, *args: str) -> tuple[int, str]:
        runner = click.testing.CliRunner()
        result = runner.invoke(app.soske, args, obj=conn)
        if result.exception:
            raise result.exception

        return result.exit_code, result.output

    return cmdfunc
