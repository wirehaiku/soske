"""
Tests for 'soske.comms.set'.
"""

from soske.comms import set_


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

    # success - existing key, same value
    code, echo = run(conn, "set", "test", "three", "four")
    row = conn.execute("select count(*) from GoodVals where key='test'").fetchone()
    assert row["count(*)"] == 2
