FROM golang:1.19.3-alpine3.16 AS builder
ENV GOPROXY "https://goproxy.cn,direct"
RUN apk add --no-cache g++ git
WORKDIR /go/src/app
COPY go.mod go.sum /go/src/app/
RUN go mod download
COPY . /go/src/app/
RUN CGO_ENABLED=1 GO111MODULE=on GOOS=linux go build -o main main.go

FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder /go/src/app/main ./main
EXPOSE 8081
VOLUME ["/app/data"]
CMD ["/app/main"]
