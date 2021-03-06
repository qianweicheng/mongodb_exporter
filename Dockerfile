FROM       golang:alpine as builder

RUN apk --no-cache add curl git make perl
RUN curl -s https://glide.sh/get | sh
COPY . /go/src/github.com/qianweicheng/mongodb_exporter
RUN cd /go/src/github.com/qianweicheng/mongodb_exporter && make release

FROM       alpine:3.4
MAINTAINER David Cuadrado <dacuad@facebook.com>
EXPOSE     9001

RUN apk add --update ca-certificates
COPY --from=builder /go/src/github.com/qianweicheng/mongodb_exporter/release/mongodb_exporter-linux-amd64 /usr/local/bin/mongodb_exporter

ENTRYPOINT [ "mongodb_exporter" ]
