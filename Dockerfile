FROM golang:1.19 as dev

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o prefix-matcher-service ./src/
ENTRYPOINT [ "/go/app/prefix-matcher-service" ]
FROM dev as debug

RUN go get -d -v ./...

FROM dev as build


