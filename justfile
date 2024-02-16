set dotenv-load := true

install:
    brew install go just jq watchexec
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    go install gotest.tools/gotestsum@latest
    go install mvdan.cc/gofumpt@latest
    go install github.com/gobench-io/gobench@master
    go install go.uber.org/nilaway/cmd/nilaway@latest

build: format lint test
    go build -n ./...

format:
    gofumpt -l -w .
    go mod tidy

lint:
    go vet ./...
    golangci-lint run ./...

typecheck:
    nilaway ./...

audit:
    not_implemented

year := "2021"
start:
    go test ./{{year}}/... -count 1 -json | jq 'select(.Test and .Elapsed > 0).Elapsed' | jq -s 'add'

dev:
    just devtest

debug:
    not_implemented

test:
    gotestsum ./... -- -count 1

devtest:
    gotestsum --hide-summary=skipped,output --watch -- -count 1

coverage:
    go test ./... -coverprofile=.coverage.out

bench:
    gobench ./...
