set dotenv-load := true

year := "2023"

build: format lint test
    go build -n ./...

format:
    go fmt ./...

lint:
    go vet ./...
    golangci-lint run ./...

typecheck:
    nilaway ./...

audit:
    not_implemented

start:
    just test

dev:
    just devtest

test:
    gotestsum --packages=./{{year}}/... -- -count 1

devtest:
    gotestsum --hide-summary=skipped,output --watch --packages=./{{year}}/...

coverage:
    go test ./{{year}}/... -coverprofile=.coverage.out

bench:
    gobench ./{{year}}/...
