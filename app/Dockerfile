FROM golang:1.11.1-alpine AS builder
WORKDIR /
COPY app.go .
ARG VERSION
ARG COLOR
RUN GOOS=linux GOARCH=amd64 go build -o app -ldflags "-w -s -X main.version=$VERSION -X main.color=$COLOR" app.go

FROM alpine
RUN apk --no-cache add curl
WORKDIR /
COPY --from=builder /app .
HEALTHCHECK CMD curl -f http://127.0.0.1:8080/healthcheck || exit 1
CMD ["./app"]
