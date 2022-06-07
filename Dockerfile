# build backend
FROM golang:1.18-alpine AS backend
RUN apk --no-cache --no-progress add --virtual build-deps build-base git

WORKDIR /paopao-ce
COPY . .
RUN CGO_ENABLED=0 go build .

FROM alpine:3.16
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=backend /paopao-ce/paopao-ce /usr/bin/paopao-ce
COPY --from=backend /paopao-ce/assets/comic.ttf /assets/comic.ttf
COPY --from=backend /paopao-ce/configs /configs
COPY config.yaml .

EXPOSE 8000

CMD ["paopao-ce"]
