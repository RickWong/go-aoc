# Advent of Code in Go

Fast and simple Advent of Code solutions written in Go. First make it work, then make it right, then make it fast. Going for all puzzles < 1s.

| Year | Stars | Runtime |
|------|------:|--------:|
| 2021 |    15 |  0.22 s |
| 2023 |    50 |  0.84 s |

## Installation

```sh
# get brew from https://brew.sh

# install just
brew install just

# install dependencies
just install

# build project
just build
```

## Usage

```sh
# run with automatic restarts
just dev

# run tests with just once
just test

# run tests with automatic restarts
just devtest

# debug on port 5678 with automatic restarts
just debug

# format all the code
just format

# typecheck all the code
just typecheck

# lint all the code (lints are extra rules agreed upon by the team)
just lint

# audit all the code
just audit
```

## VS Code extensions

- Go (official)
- Golang postfix code completion
- Go Snippets
- Tabnine AI
- Task (official)

## GoLand plugins

- GitHub Copilot
- Just
- Rainbow Brackets
- Tabnine AI

## License

This project is licensed under the GNU General Public License version 3.0 (GPL-3.0, or GPLv3).
