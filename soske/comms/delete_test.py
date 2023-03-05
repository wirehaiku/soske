"""
Tests for 'soske.comms.delete'.
"""

from soske.comms import delete
from soske.tools.test import conn, run


def test_delete(conn, run):
    # success - soft delete
    code, echo = run(conn, "delete", "alpha")
    row = conn.execute("select * from Keys where name='alpha'").fetchone()
    assert code == 0
    assert echo == ""
    assert row["dead"] == True

    # success - hard delete, yes
    code, echo = run(conn, "delete", "bravo", "--hard", "--yes")
    row1 = conn.execute("select count(*) from Keys where name='bravo'").fetchone()
    row2 = conn.execute("select count(*) from Vals where key='bravo'").fetchone()
    assert code == 0
    assert echo == "Key deleted.\n"
    assert row1["count(*)"] == row2["count(*)"] == 0
