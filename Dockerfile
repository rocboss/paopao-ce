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

WORKDIR /app/paopao-ce
COPY --from=backend /paopao-ce/release/paopao-ce .
COPY assets ./assets
COPY configs ./configs

VOLUME ["/app/paopao-ce/assets", "/app/paopao-ce/configs"]
EXPOSE 8008
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD ps -ef | grep paopao-ce || exit 1
ENTRYPOINT ["/app/paopao-ce/paopao-ce"]
