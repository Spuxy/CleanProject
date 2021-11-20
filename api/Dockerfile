FROM golang:1.14-alpine as base

WORKDIR /api

FROM aquasec/trivy:0.4.4 as trivy 

RUN trivy --debug --timeout 4m golang:1.14-alpine && \
  echo "No image vulnerabilities" > result

FROM base as dev

COPY go.* ./

RUN go mod download

ENV GOPATH /go

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go env

RUN go get github.com/go-delve/delve/cmd/dlv \
  && go get github.com/githubnemo/CompileDaemon

EXPOSE 4000 2345
