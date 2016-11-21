FROM golang:1.7.3-wheezy

ARG proxy
ENV http_proxy $proxy
ENV https_proxy $proxy

## AWS Credential
ARG aws_key_id
ARG aws_access_key
ENV AWS_ACCESS_KEY_ID       $aws_key_id
ENV AWS_SECRET_ACCESS_KEY   $aws_access_key
ENV AWS_DEFAULT_REGION us-east-1

## OPENSTACK Credential

ARG op_tenant_name
ARG op_username
ARG op_password
ARG op_auth_url
ENV OS_TENANT_NAME  $op_tenant_name
ENV OS_USERNAME     $op_username
ENV OS_PASSWORD     $op_password
ENV OS_AUTH_URL     $op_auth_url

ENV PRJNAME terrant
ENV PRJPATH /go/src/github.com/perigee/$PRJNAME
COPY . $PRJPATH

WORKDIR $PRJPATH
RUN go install

EXPOSE 8090

ENTRYPOINT ["./$PRJPATH/infra"]