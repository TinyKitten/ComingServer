FROM golang:1.12-alpine

WORKDIR /go/src/github.com/TinyKitten/ComingServer
COPY . .

RUN apk add make alpine-sdk

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["ComingServer"]
