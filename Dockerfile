FROM golang:1.12.6-alpine

MAINTAINER devops787

WORKDIR /go/src/go-kubernetes
COPY . /go/src/go-kubernetes

RUN apk --no-cache add git dep
RUN rm -rf ./vendor && dep ensure && go build -o app

EXPOSE 3000

ENTRYPOINT ["./app"]