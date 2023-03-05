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

Print the value of a key (if it exists):

```
$ soske get alpha
Aye aye aye! 
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

- Please submit bug reports and feature requests to the [issue tracker][bugs].
- Soske's only dependencies are [SQLite][sqli], [Click][clck] and [Testify][test]. 

[bugs]: https://github.com/wirehaiku/Soske/issues
[clck]: https://click.palletsprojects.com/en/8.1.x/
[rels]: https://github.com/wirehaiku/Soske/releases/latest
[p310]: https://www.python.org/downloads/release/python-3100/
[sqli]: https://www.sqlite.org/index.html
[stvm]: https://wirehaiku.org/
[test]: https://github.com/stretchr/testify
