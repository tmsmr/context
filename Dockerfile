FROM golang:1.21-alpine3.19 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w"

FROM scratch
COPY --from=builder /build/context /context
ENTRYPOINT ["/context"]
