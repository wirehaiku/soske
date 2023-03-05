"""
Tests for 'soske.comms.list_'.
"""

from soske.comms import list_


def test_list_(conn, run):
    # success
    code, echo = run(conn, "list")
    assert code == 0
    assert echo == "alpha\nbravo\n"
