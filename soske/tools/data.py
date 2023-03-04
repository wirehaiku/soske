"""
Low-level database handler functions.
"""

import os.path
import sqlite3
from typing import Any


def open(path: str, pragma: str, schema: str) -> sqlite3.Connection:
    """
    Return a new database connection. Pragma is executed on a successful connection,
    schema is executed if the database is in-memory or a file that does not exist.
    """

    exst = False if path == ":memory:" else os.path.isfile(path)
    conn = sqlite3.connect(path)
    conn.row_factory = sqlite3.Row

    with conn:
        conn.executescript(pragma)
        if schema and not exst:
            conn.executescript(schema)

    return conn


def exec(conn: sqlite3.Connection, sql: str, *args: Any):
    """
    Execute and commit a database query.
    """

    with conn:
        conn.execute(sql, args)


def row(conn: sqlite3.Connection, sql: str, *args: Any) -> sqlite3.Row:
    """
    Return the first resulting row of a database query.
    """

    return conn.execute(sql, args).fetchone()


def rows(conn: sqlite3.Connection, sql: str, *args: Any) -> list[sqlite3.Row]:
    """
    Return all resulting rows of a database query.
    """

    return conn.execute(sql, args).fetchall()
