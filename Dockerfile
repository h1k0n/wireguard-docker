FROM golang:1.13.4-stretch AS builder

# Install application
WORKDIR /go/src/wgconf
ADD . .

# Build application
RUN go build

FROM debian:stretch

RUN echo "deb http://deb.debian.org/debian/ unstable main" > /etc/apt/sources.list.d/unstable-wireguard.list && \
 printf 'Package: *\nPin: release a=unstable\nPin-Priority: 90\n' > /etc/apt/preferences.d/limit-unstable

RUN apt update && \
 apt install -y --no-install-recommends curl openresolv procps wireguard-tools iptables nano net-tools && \
 apt clean

WORKDIR /scripts
ENV PATH="/scripts:${PATH}"
COPY --from=builder /go/src/wgconf/wgconf /scripts
COPY install-module /scripts
COPY run /scripts
COPY wgconf /scripts
COPY genkeys /scripts
RUN chmod 755 /scripts/*

CMD ["run"]
