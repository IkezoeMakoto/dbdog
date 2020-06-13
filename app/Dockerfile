FROM golang:1.14-alpine

RUN apk update && \
    apk add make git && \
    rm -rf /var/cache/apk/*

RUN apk update && \
    apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

WORKDIR /app
COPY . /app

CMD make -C src build