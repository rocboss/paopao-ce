# build app
FROM golang AS build-env

ADD . /paopao-ce

WORKDIR /paopao-ce

RUN CGO_ENABLED=0 go build .

# safe image
FROM alpine

ENV TZ=Asia/Shanghai

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=build-env /paopao-ce/paopao-ce /usr/bin/paopao-ce
COPY --from=build-env /paopao-ce/assets/comic.ttf /assets/comic.ttf
COPY --from=build-env /paopao-ce/configs /configs
COPY config.yaml .

EXPOSE 8000

CMD ["paopao-ce"]

# HEALTHCHECK
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep paopao-ce || exit 1
