FROM golang:1.22.3-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o temperature-postal-code ./cmd

FROM scratch
WORKDIR /app
COPY --from=builder /app/cmd/.env .
COPY --from=builder /app/temperature-postal-code .

ENTRYPOINT ["/app/temperature-postal-code"]