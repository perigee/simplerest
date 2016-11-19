FROM golang:1.7-alpine

ARG proxy
ENV http_proxy $proxy
ENV https_proxy $proxy


ENV PRJNAME terrant
ENV PRJPATH /go/src/github.com/perigee/$PRJNAME
COPY . $PRJPATH

WORKDIR $PRJPATH
RUN apk add --update --no-cache git openssl \
    && go install

EXPOSE 8080

ENTRYPOINT ["./$PRJPATH/infra"]