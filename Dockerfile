FROM golang:alpine

#Install dev tools
RUN apk add --update --no-cache alpine-sdk bash ca-certificates \
      libressl \
      tar \
      git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils

# Copy our source code into the container.
WORKDIR /go/src/tap_exporter
ADD . /go/src/tap_exporter

RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /go/bin/tap_exporter .

CMD ["./tap_exporter"]
