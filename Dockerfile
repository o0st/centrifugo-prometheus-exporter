FROM golang:1.11.6-alpine3.9 as builder

LABEL maintainer="Oleg Ozimok oleg.ozimok@corp.kismia.com"

COPY . /go/src/github.com/kismia/centrifugo-prometheus-exporter

WORKDIR /go/src/github.com/kismia/centrifugo-prometheus-exporter

RUN go build -o /centrifugo-prometheus-exporter ./cmd/centrifugo-prometheus-exporter

FROM alpine:3.9

COPY --from=builder /centrifugo-prometheus-exporter /usr/bin/centrifugo-prometheus-exporter

EXPOSE 9100

STOPSIGNAL SIGTERM

ENTRYPOINT ["/usr/bin/centrifugo-prometheus-exporter"]