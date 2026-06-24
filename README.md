# leetcode-go

LeetCode / NeetCode practice in Go, run inside Docker.

## Setup

Bring up the dev container:

```sh
docker compose up -d --build
```

Open a shell inside it:

```sh
docker compose exec go sh
```

Initialize the Go module (first time only):

```sh
go mod init github.com/kobaryubi/leetcode-go
```

## Usage

The repo is bind-mounted into the container, so edits made on the host are
picked up immediately — write solutions in your editor and work from the
container shell.

Stop the container when done:

```sh
docker compose down
```

## Tests

Each problem lives in its own directory (e.g. `0006-zigzag-conversion/`)
with a solution file and a table-driven `_test.go` file.

Run every test in the repo:

```sh
go test ./...
```

Run the tests for a single problem:

```sh
go test ./0006-zigzag-conversion/
go test ./0007-reverse-integer/
```

Add `-v` to see each case, or `-run` to filter by name:

```sh
go test -v -run TestConvert ./0006-zigzag-conversion/
go test -v -run TestReverse ./0007-reverse-integer/
```

## Formatting

Format Go source with `gofmt` (see https://pkg.go.dev/cmd/gofmt).

`gofmt` takes file or directory paths (it recurses into directories on
its own) — it does not understand the `./...` package pattern.

List files that need formatting under the current directory:

```sh
gofmt -l .
```

Show the diff for one directory without writing:

```sh
gofmt -d ./0006-zigzag-conversion
gofmt -d ./0007-reverse-integer
```

Write formatting changes in place:

```sh
gofmt -w ./0006-zigzag-conversion
gofmt -w ./0007-reverse-integer
```

`go fmt ./...` is the package-aware wrapper that formats every package
in the module — use this when you do want the `./...` pattern.

## Running a program

For directories with a `package main` and a `main()` function (e.g.
`learn-fuzz/`), compile and run in one step:

```sh
go run ./learn-fuzz
```

The path is the package import path — `./` makes it relative to the repo
root. From inside the directory, `go run .` works the same way.

Build a standalone binary instead:

```sh
go build ./learn-fuzz   # produces ./learn-fuzz
./learn-fuzz
```
