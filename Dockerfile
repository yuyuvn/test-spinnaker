FROM golang:1.12.7-alpine3.10

ENV webserver_path /go/src/github.com/yuyuvn/test-spinnaker/
ENV PATH $PATH:$webserver_path

WORKDIR $webserver_path
COPY sourcecode/ .

RUN go build .

ENTRYPOINT ./test-spinnaker

EXPOSE 80