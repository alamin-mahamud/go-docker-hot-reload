FROM golang:1.10

WORKDIR /go/src/github.com/alamin-mahamud/go-docker-hot-reload-example

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"
