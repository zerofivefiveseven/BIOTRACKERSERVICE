FROM golang:1.22.2-alpine3.18 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o cmd/service/bin/main ./cmd/service/

FROM alpine:latest
WORKDIR /BIOTRACKERSERVICE
COPY /.env .
COPY /config/config.yaml ./config/
COPY --from=builder /app/cmd/service/bin/main .
EXPOSE 8080
CMD ["./main"]