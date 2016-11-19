FROM golang:1.7.3-wheezy

ARG proxy
ENV http_proxy $proxy
ENV https_proxy $proxy


ENV PRJNAME terrant
ENV PRJPATH /go/src/github.com/perigee/$PRJNAME
COPY . $PRJPATH

WORKDIR $PRJPATH
RUN go install

EXPOSE 8090

ENTRYPOINT ["./$PRJPATH/infra"]