
FROM golang:1.16.5

ADD . /go/src/github.com/TIove/links_cutter

RUN go install github.com/TIove/links_cutter@latest

ENTRYPOINT ["/go/bin/links_cutter"]

EXPOSE 50051
EXPOSE 5432
