FROM ubuntu:20.04

LABEL maintainer="smalls smalls0098@gmail.com"

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /tmp

RUN sed -i 's/archive.ubuntu/mirrors.aliyun/g' /etc/apt/sources.list
RUN apt-get update &&  \
    apt-get install -y --no-install-recommends \
            git build-essential cmake pkg-config unzip libgtk2.0-dev \
            curl ca-certificates libcurl4-openssl-dev libssl-dev \
            libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev \
            libjpeg-dev libpng-dev libtiff-dev libdc1394-22-dev && \
            rm -rf /var/lib/apt/lists/*

RUN ln -s /usr/bin/python3 /usr/bin/python

RUN apt-get autoclean && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*