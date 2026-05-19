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
