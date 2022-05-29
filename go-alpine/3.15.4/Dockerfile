FROM alpine:3.15.4

LABEL maintainer="smalls smalls0098@gmail.com"

ENV TZ=Asia/Shanghai

RUN echo 'http://mirrors.ustc.edu.cn/alpine/v3.15/main' > /etc/apk/repositories \
    && echo 'http://mirrors.ustc.edu.cn/alpine/v3.15/community' >>/etc/apk/repositories

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/$TZ /etc/localtime && \
    apk add --update --no-cache ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/* && \
    update-ca-certificates