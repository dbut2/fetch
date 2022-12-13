## Fetch

Download and open Advent of Code puzzle and inputs files automatically

### Installation

1. Install and follow installation steps for [scarvalhojr/aoc-cli](https://github.com/scarvalhojr/aoc-cli)

```shell
go install github.com/dbut2/fetch@latest
```

### Usage

From your aoc directory root, run fetch
```shell                  
fetch --help

Usage of fetch:
  -day int
        day to download puzzle for (default 25)
  -files string
        comma seperated list of files to open automatically (default "puzzle.md")
  -ide string
        ide command to open files, must support opening files like "$ {IDE} example" (default "goland")
  -template string
        template folder (default "template")
  -year int
        year to download puzzle for (default 2022)

```

Set up your arguments, eg
```shell
fetch --files=puzzle.md --ide=goland --year=2020 --day=01 --template=template
```

Optionally create an alias run quicker
```shell
alias f="fetch --files=puzzle.md --ide=goland --template=template"
```

Optionally provide 