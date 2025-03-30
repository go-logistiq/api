FROM golang:alpine AS app

WORKDIR /app

COPY . ./

RUN go mod download && \
    GOOS=linux GOARCH=amd64 go build -o /out/logistiq-api

FROM alpine:latest

WORKDIR /app

COPY --from=app /out/logistiq-api ./

EXPOSE 3000

CMD ["./logistiq-api"]