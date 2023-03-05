"""
Tests for 'soske.comms.get'.
"""

from soske.comms import get
from soske.tools.test import conn, run


def test_get(conn, run):
    # success
    code, echo = run(conn, "get", "alpha")
    assert code == 0
    assert echo == "Alpha two.\n"
