"""
Main application command group.
"""

import os.path

import click
from soske.tools import data, sqls

PATH_TYPE = click.Path(dir_okay=False, writable=True)


@click.group()
@click.option("--path", default="~/.soske", help="Database path.", type=PATH_TYPE)
@click.pass_context
def soske(ctx: click.Context, path: str):
    """
    Soske: Stephen's Old-School Key Engine.
    """

    if not getattr(ctx, "obj", None):
        path = os.path.expanduser(path)
        path = os.path.expandvars(path)
        ctx.obj = data.open(path, sqls.PRAGMA, sqls.SCHEMA)
