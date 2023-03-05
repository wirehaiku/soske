"""
Command definition for 'delete'.
"""

import sqlite3

import click
from soske.comms.app import soske
from soske.tools import data


@soske.command(short_help="Delete a key.")
@click.argument("key")
@click.option("-h", "--hard", help="Hard-delete key.", is_flag=True)
@click.option("-y", "--yes", help="Bypass confirmation.", is_flag=True)
@click.pass_obj
def delete(conn: sqlite3.Connection, key: str, hard: bool, yes: bool):
    """
    Delete an existing key.
    """

    key = key.lower()

    if row := data.row(conn, "select * from GoodKeys where name=?", key):
        if hard:
            if yes or click.confirm(f"Are you sure you want to delete {key!r}?"):
                data.exec(conn, "delete from Vals where key=?", key)
                data.exec(conn, "delete from Keys where name=?", key)
                click.echo("Key deleted.")
            else:
                click.echo("Deletion cancelled.")

        else:
            data.exec(conn, "update Keys set dead=true where name=?", key)
