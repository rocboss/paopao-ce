# build app
FROM golang AS build-env

ADD . /paopao-api

WORKDIR /paopao-api

RUN CGO_ENABLED=0 go build .

# safe image
FROM alpine

ENV TZ=Asia/Shanghai

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=build-env /paopao-api/paopao-api /usr/bin/paopao-api
COPY --from=build-env /paopao-api/comic.ttf /comic.ttf
COPY --from=build-env /paopao-api/qqwry.dat /qqwry.dat
COPY --from=build-env /paopao-api/configs /configs

EXPOSE 8000

CMD ["paopao-api"]

# HEALTHCHECK
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep paopao-api || exit 1
