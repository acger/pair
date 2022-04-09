FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
COPY service/pair/api/etc /app/etc
RUN rm go.mod go.sum && go mod init pair && go mod tidy
RUN go build -ldflags="-s -w" -o /app/pair service/pair/api/pair.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/pair /app/pair
COPY --from=builder /app/etc /app/etc

CMD ["./pair", "-f", "etc/pair-api.yaml"]
