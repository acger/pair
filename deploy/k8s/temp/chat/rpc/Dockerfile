FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl && apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
COPY service/chat/rpc/etc /app/etc
RUN rm go.mod go.sum && go mod init pair && go mod tidy
RUN go build -ldflags="-s -w" -o /app/chat service/chat/rpc/chat.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/chat /app/chat
COPY --from=builder /app/etc /app/etc

CMD ["./chat", "-f", "etc/chat.yaml"]
