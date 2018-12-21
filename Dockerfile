FROM alpine:latest as certs
MAINTAINER kubernetes@platform9.com

RUN apk --no-cache --update upgrade && apk --no-cache add ca-certificates

FROM scratch
COPY pa-proxy /
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/pa-proxy"]