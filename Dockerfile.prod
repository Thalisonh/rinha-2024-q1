FROM golang AS builder

WORKDIR /app

COPY . /app

RUN go mod download

EXPOSE 8080

RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api ./
EXPOSE 8080

ENTRYPOINT [ "./api" ]

