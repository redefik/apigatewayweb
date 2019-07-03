#build stage
FROM golang:1.11 AS build-env

WORKDIR /go/src/github.com/redefik/apigatewayweb

COPY . .

RUN go get -d -v ./...

RUN cd cmd && CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o /go/bin/apigatewayweb 

#production stage
FROM alpine:latest

WORKDIR /root/

COPY --from=build-env /go/bin/apigatewayweb .

EXPOSE 80

CMD ["./apigatewayweb"]
