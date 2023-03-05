"""
Command definition for 'set'.
"""

import sqlite3
import time

import click
from soske.comms.app import soske
from soske.tools import data


@soske.command(name="set", short_help="Set a key.")
@click.argument("key")
@click.argument("values", nargs=-1)
@click.pass_obj
def set_(conn: sqlite3.Connection, key: str, values: str):
    """
    Set the value of a new or existing key.
    """

    key = key.lower()
    val = "\n".join(values)
    now = time.time()
    row = data.row(conn, "select * from GoodVals where key=?", key)

    if not row or row["body"] != val:
        data.exec(conn, "insert or ignore into Keys values (?, false, ?)", key, now)
        data.exec(conn, "insert into Vals values (?, ?, ?)", key, val, now)
