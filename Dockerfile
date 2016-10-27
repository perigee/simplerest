FROM golang:1.7-alpine

ARG proxy
ENV http_proxy $proxy
ENV https_proxy $proxy


ENV PRJNAME simplerest
ENV PRJPATH /go/src/github.com/perigee/$PRJNAME
COPY . $PRJPATH

WORKDIR $PRJPATH
RUN mkdir -p $PRJPATH \
    && apk add --update --no-cache git openssl \
    && go get && go install

EXPOSE 8080

ENTRYPOINT ["/go/bin/simplerest"]