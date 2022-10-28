FROM golang:1.19-alpine

RUN apk add --no-cache make

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY Makefile .


COPY cmd ./cmd



RUN make build


CMD ["./bin/encodr"]
