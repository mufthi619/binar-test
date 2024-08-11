FROM golang:1.22-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config/app.yaml ./config/app.yaml
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Etc/UTC

RUN mkdir -p /root/uploads

EXPOSE 8080

CMD ["./main"]