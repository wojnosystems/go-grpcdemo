FROM golang AS BUILDER
WORKDIR /build
ADD . .
ENV CGO_ENABLED=false
RUN go build -o /tmp/server -tags netgo cmd/server/main.go

FROM alpine
COPY --from=BUILDER /tmp/server /srv/server
ENTRYPOINT ["/srv/server"]
