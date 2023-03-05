# Soske

**Soske** (*Stephen's Old-School Key Engine*, pronounced *sos·kay*) is a command-line key-value store written in [Python 3.10][p310] by [Stephen Malone][stvm].
It's designed to be a simple, portable plaintext data-dumpster with robust history tracking and import/export features.

```bash
# TODO: A slick Asciinema demo.
```

## Installation

Install Soske with `pip install`...

```
pip install soske
```

...or download the [latest release][rels] for your platform.

## Configuration

Soske stores all its data in a single [SQLite][sqli] database in your home directory at `$HOME/.soske`.
The database itself is an open standard and there are commands to easily import and export your data.

## Commands

### `list`

List all existing keys.

```
$ soske lst
alpha
bravo
charlie
```

### `get KEY`

Print the value of a key (if it exists).

```
$ soske get alpha
Aye aye aye! 
```

### `delete KEY`

Delete a key and its values (if it exists).
By default the key is soft-deleted, meaning it is *marked* as deleted in the database but no actual data is lost.

- Use the `-h --hard` flag to hard-delete and actually destroy the key's data.
- Use the `-y --yes` flag to bypass the "Are you sure?" prompt.

```
$ soske delete alpha
$ soske delete bravo --hard
Are you sure you want to hard-delete 'bravo'? [y/N]: y
Key deleted.
```

### `set KEY VALUES...`

Set the value of a new or existing key. 
- Multiple arguments are combined into a single newline-separated value.

```
$ soske set alpha one two three
$ soske get alpha
one
two
three
```

## Contributing

Please submit bug reports and feature requests to the [issue tracker][bugs], thank you.

## Credits

Soske is made possible by these third-party libraries:

- [click](https://click.palletsprojects.com)
- [mypy](https://www.mypy-lang.org/)
- [py.test](https://docs.pytest.org/)
- [pytest-mypy](https://pypi.org/project/pytest-mypy/)

[bugs]: https://github.com/wirehaiku/Soske/issues
[rels]: https://github.com/wirehaiku/Soske/releases/latest
[p310]: https://www.python.org/downloads/release/python-3100/
[sqli]: https://www.sqlite.org/index.html
[stvm]: https://wirehaiku.org/
