# Development image for practicing LeetCode / NeetCode problems in Go.
FROM golang:alpine

# git is handy for pulling deps; bash/make for a nicer dev shell.
RUN apk add --no-cache git bash make

WORKDIR /app

# No go.mod is baked in — run `go mod init` inside the container once started.
CMD ["sleep", "infinity"]
