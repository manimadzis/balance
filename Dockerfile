FROM golang:latest

ADD . $GOROOT/app
WORKDIR $GOROOT/app
RUN go build -o server cmd/main.go
ENTRYPOINT ["./server",  "-f", "config.toml"]


