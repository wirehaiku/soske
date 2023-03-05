"""
Tests for 'soske.comms.get'.
"""

from soske.comms import get


def test_get(conn, run):
    # success
    code, echo = run(conn, "get", "alpha")
    assert code == 0
    assert echo == "Alpha two.\n"
