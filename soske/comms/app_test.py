"""
Testss for 'soske.comms.app'.
"""

import click
from soske.comms import app
from soske.tools.test import conn, run


def test_soske(conn, run):
    # setup
    @app.soske.command()
    @click.pass_obj
    def mock(conn):
        for row in conn.execute("select * from GoodKeys").fetchall():
            click.echo(row["name"])

    # success
    code, echo = run(conn, "mock")
    assert code == 0
    assert echo == "alpha\nbravo\n"
