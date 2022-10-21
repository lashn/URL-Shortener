# Build state
FROM golang:alpine AS builder

RUN mkdir /build

WORKDIR /build
COPY . .
RUN go build -o main app/main.go


#Run stage
FROM alpine:latest
WORKDIR /build
COPY --from=builder /build/main .

EXPOSE 8000

CMD ["/build/main"]



