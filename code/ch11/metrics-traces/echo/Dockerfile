FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN GOPROXY=off go build -o echo -mod=vendor main.go

FROM alpine
WORKDIR /app
COPY --from=builder /build/echo /app/echo
CMD ["./echo"]
