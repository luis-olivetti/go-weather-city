FROM golang:1.21.3 AS builder
WORKDIR /app
COPY . .
RUN go install github.com/google/wire/cmd/wire@latest
RUN $GOPATH/bin/wire ./cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goweather ./cmd

FROM scratch
WORKDIR /app
COPY --from=builder /app/goweather .
ENTRYPOINT ["./goweather"]