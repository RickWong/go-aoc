# Go: Advent of Code 2021

## Installation

```sh
# get brew from https://brew.sh

# install go
brew install go

# install go-task
go install github.com/go-task/task/v3/cmd/task@latest

# install go-releaser
go install github.com/goreleaser/goreleaser@latest

# install gotest for colorized test output
go install github.com/rakyll/gotest

# build project
task build
```

## Usage

```sh
# run once
task start

# run with automatic restarts
task dev

# run tests with just once
task test

# run tests with automatic restarts
task devtest

# debug on port 5678 with automatic restarts
task debug

# format all the code
task format

# typecheck all the code
task typecheck

# lint all the code (lints are extra rules agreed upon by the team)
task lint

# audit all the code
task audit

# run all build steps: lint, format, typecheck, audit, test and coverage
task build
```

## VS Code extensions

- Go (official)
- Golang postfix code completion
- Go Snippets
- Tabnine AI
- Task (official)

## GoLand plugins

- GitHub Copilot
- Rainbow Brackets
- Tabnine AI

## License

This project is licensed under the GNU General Public License version 3.0 (GPL-3.0, or GPLv3).
