# build frontend
FROM node:19-alpine as frontend
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
WORKDIR /web
COPY web/ ./
RUN [ $EMBED_UI != yes ] || [ $USE_API_HOST != yes ] || echo "VITE_HOST='$API_HOST'">.env.local
RUN [ $EMBED_UI != yes ] || [ $USE_DIST != no ] || (yarn && yarn build)
RUN [ $EMBED_UI = yes ] || mkdir dist || echo ""

# build backend
FROM golang:1.20-alpine AS backend
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
RUN apk --no-cache --no-progress add --virtual \
  build-deps \
  build-base \
  git

WORKDIR /paopao-ce
COPY . .
COPY --from=frontend /web/dist ./web/dist
ENV GOPROXY=https://goproxy.cn
RUN [ $EMBED_UI != yes ] || make build TAGS='embed go_json'
RUN [ $EMBED_UI = yes ] || make build TAGS='go_json'

FROM alpine:3.17
ARG API_HOST
ARG USE_API_HOST=yes
ARG EMBED_UI=yes
ARG USE_DIST=no
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /app/paopao-ce
COPY --from=backend /paopao-ce/release/paopao-ce .

VOLUME ["/app/paopao-ce/custom"]
EXPOSE 8008
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep paopao-ce || exit 1
ENTRYPOINT ["/app/paopao-ce/paopao-ce"]
