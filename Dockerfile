
FROM golang:1.22-bookworm as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -v -o server


FROM debian:bookworm-slim

COPY --from=builder /app/server /app/server

CMD ["/app/server"]
