FROM golang:1.12-alpine

WORKDIR /go/src/github.com/TinyKitten/ComingServer
COPY . .

RUN apk add make mysql mysql-client alpine-sdk

RUN go get -d -v ./...
RUN go install -v ./...

RUN go get github.com/rubenv/sql-migrate/...

CMD ["ComingServer"]
