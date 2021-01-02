FROM golang AS BUILDER
WORKDIR /build
ADD . .
ENV CGO_ENABLED=false
RUN go build -o /tmp/client -tags netgo cmd/cli/main.go cmd/cli/dogsay.go

FROM alpine
COPY --from=BUILDER /tmp/client /usr/bin/client
ENTRYPOINT ["/usr/bin/client"]
