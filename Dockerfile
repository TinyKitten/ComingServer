FROM golang:1.12-alpine

WORKDIR /go/src/github.com/TinyKitten/ComingServer
COPY . .

RUN apk add make alpine-sdk

ADD ./jwtkey/jwt.key ./jwtkey
ADD ./jwtkey/jwt.pub ./jwtkey

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["ComingServer"]
