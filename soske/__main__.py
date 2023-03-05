"""
Main program function.
"""

from soske.comms import app


def main(args: list[str] | None = None):
    """
    Run the main Soske program.
    """

    app.soske.main(args)


if __name__ == "__main__":
    main()
