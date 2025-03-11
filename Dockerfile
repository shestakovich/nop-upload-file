FROM golang:1.22-alpine3.19 AS compile

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o fp


FROM alpine:3.19

WORKDIR /app

COPY --from=compile /build/fp /app/

ENV GIN_MODE=release

CMD ["/app/fp"]
