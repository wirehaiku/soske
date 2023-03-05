"""
Command definition for 'get'.
"""

import sqlite3

import click
from soske.comms.app import soske
from soske.tools import data


@soske.command(short_help="Get a key.")
@click.argument("key")
@click.pass_obj
def get(conn: sqlite3.Connection, key: str):
    """
    Get the latest value of an existing key.
    """

    key = key.lower()
    if row := data.row(conn, "select * from GoodVals where key=?", key):
        click.echo(row["body"])
