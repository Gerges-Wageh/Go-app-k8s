# ---------- Build Stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY main.go .

RUN go build -o app main.go


# ---------- Runtime Stage ----------
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 3000

CMD ["./app"]

