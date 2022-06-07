# build frontend
FROM node:18-alpine as frontend
WORKDIR /web
COPY web/ ./
RUN echo 'VITE_HOST=""'>.env && yarn && yarn build

# build backend
FROM golang:1.18-alpine AS backend
RUN apk --no-cache --no-progress add --virtual \
  build-deps \
  build-base \
  git

WORKDIR /paopao-ce
COPY . .
COPY --from=frontend /web/dist ./web/dist
ENV GOPROXY=https://goproxy.cn
RUN make build TAGS='embed'

FROM alpine:3.16
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=build-env /paopao-ce/paopao-ce /usr/bin/paopao-ce
COPY --from=build-env /paopao-ce/assets/comic.ttf /assets/comic.ttf
COPY --from=build-env /paopao-ce/configs /configs
COPY config.yaml .

EXPOSE 8000

CMD ["paopao-ce"]
