# Soske

**Soske** (*Stephen's Old-School Key Engine*, pronounced *sos·kay*) is a command-line key-value store written in [Go 1.20][gver] by [Stephen Malone][stvm].
It's designed to be a simple, portable plaintext data-dumpster.

```bash
# TODO: A slick Asciinema demo.
```

## Installation

Install Soske with `go install`...

```
go install github.com/wirehaiku/Soske@latest
```

...or download the [latest release][rels] for your platform.

## Configuration.

Soske stores all its data in a single [Bolt][bolt] database in either [`$XDG_CONFIG_HOME`][xdgs] or [`%APPDATA%`][appd], depending on your platform.
The database itself is an open standard and there are commands to easily import and export your data.

## Commands

Soske's syntax uses the form `soske <command> <args...>` and is always variadic, meaning you can include multiple final arguments to apply the command multiple times.

### `del KEYS...`

Delete each key if it exists, or do nothing:

```
$ soske del foo
```

### `get KEYS...`

Print the value of each key if it exists, or do nothing:

```
$ soske get foo
Bar bar bar bar Barbara Ann...
```

### `set KEY VALUES...`

Set the value of a new or existing key. 
Multiple arguments are combined into a single newline-separated value.

```
$ soske set foo one two three
$ soske get foo
one
two
three
```

### `lst SUBSTRINGS...`

List all existing keys containing a substring, or all keys if no arguments are provided.
Multiple arguments repeatedly filter the list.

```
$ soske lst
foo
bar
baz

$ soske lst ba ar
bar
```


### `exp FILES...`

Export the entire database as compact JSON to each file.

```
$ soske exp test.json
$ cat test.json
{"foo":"bar"}
```

### `imp TYPE FILES...`

Import each JSON file into the database.

```
$ echo '{"foo":"bar"}' > test.json
$ soske imp test.json
$ soske get foo
bar
```

## Contributing

- Please submit bug reports and feature requests to the [issue tracker][bugs].
- Soske's only dependencies are [Bolt][bolt] (database) and [Testify][test] (unit testing). 
- The `extra` directory contains helper scripts for building and testing Soske.

[appd]: https://ss64.com/nt/syntax-variables.html
[bolt]: https://github.com/etcd-io/bbolt
[bugs]: https://github.com/wirehaiku/Soske/issues
[gver]: https://go.dev/doc/go1.20
[rels]: https://github.com/wirehaiku/Soske/releases/latest
[stvm]: https://wirehaiku.org/
[test]: https://github.com/stretchr/testify
[xdgs]: https://wiki.archlinux.org/title/XDG_Base_Directory

