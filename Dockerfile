FROM golang:1.12.6-stretch as builder

RUN mkdir -p /go/src/server
WORKDIR /go/src/server

COPY . .

RUN go get .
RUN go build server.go

FROM alpine:edge

RUN mkdir /app
COPY --from=builder /go/src/server/server /app/server
CMD ["/app/server"]
