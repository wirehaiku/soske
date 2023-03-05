"""
Tests for 'soske.comms.set'.
"""

from soske.comms import set_
from soske.tools.test import conn, run


def test_set_(conn, run):
    # success - new key
    code, echo = run(conn, "set", "test", "one", "two")
    row = conn.execute("select * from GoodVals where key='test'").fetchone()
    assert code == 0
    assert echo == ""
    assert row["body"] == "one\ntwo"

    # success - existing key
    code, echo = run(conn, "set", "test", "three", "four")
    row = conn.execute("select * from GoodVals where key='test'").fetchone()
    assert code == 0
    assert echo == ""
    assert row["body"] == "three\nfour"