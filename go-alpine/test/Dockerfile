FROM golang:1.18.1-alpine AS builder

ARG GOPROXY="https://goproxy.cn,direct"
ENV GOPROXY=$GOPROXY

# 设置工作目录
WORKDIR /app
# 拷贝项目到容器
COPY go.mod ./
COPY main.go ./

RUN go build -o ./main .


FROM smalls0098/go-alpine:3.15.4

LABEL maintainer="smalls"

ENV TZ=Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/main ./

CMD ["/app/main"]