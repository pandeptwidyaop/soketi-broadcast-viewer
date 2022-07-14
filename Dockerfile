FROM golang:1.17-alpine AS BUILDER

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o app .

FROM alpine:3.14

RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime \
    && echo "Asia/Jakarta" > /etc/timezone

COPY --from=BUILDER ["/build/app","/"]


EXPOSE 8000

CMD ["/app"]
