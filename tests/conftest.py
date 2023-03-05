"""
Global unit-testing fixtures and configuration.
"""

import sqlite3

import click.testing
import pytest
from soske.comms import app
from soske.tools import sqls


@pytest.fixture(scope="function")
def conn():
    """
    Return a connection to an in-memory database populated with test data.
    """

    conn = sqlite3.connect(":memory:")
    conn.row_factory = sqlite3.Row
    conn.executescript(sqls.PRAGMA + sqls.SCHEMA + sqls.TEST_SCHEMA)
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
