FROM golang:1.20-alpine
RUN apk add git make \
    && git clone https://github.com/chenrizhi/alertmanager-wechatrobot-webhook.git \
    && cd alertmanager-wechatrobot-webhook \
    && make


FROM alpine:3.18.2

ENV PATH /usr/local/bin:$PATH
ENV LANG C.UTF-8

ENV TZ=Asia/Shanghai

RUN set -ex \
    && apk update && apk upgrade \
    && apk --no-cache add openssl wget bash tzdata curl ca-certificates \
    && update-ca-certificates

COPY --from=0 /go/alertmanager-wechatrobot-webhook/bin/wechat-webhook /usr/local/bin/wechat-webhook

ENTRYPOINT ["wechat-webhook"]
