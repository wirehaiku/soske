"""
Command definition for 'list'.
"""

import sqlite3

import click
from soske.comms.app import soske
from soske.tools import data


@soske.command(name="list", short_help="List all keys.")
@click.pass_obj
def list_(conn: sqlite3.Connection):
    """
    List all existing keys.
    """

    for row in data.rows(conn, "select * from GoodKeys"):
        click.echo(row["name"])
