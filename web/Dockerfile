# build frontend
FROM node:19-alpine as frontend
ARG API_HOST
ARG USE_DIST=no
WORKDIR /web
COPY . .
RUN [ empty$API_HOST = empty ] || echo "VITE_HOST=$API_HOST" >.env.local
RUN [ $USE_DIST != no ] || (yarn && yarn build)

FROM library/nginx
ARG API_HOST
ARG USE_DIST=no
USER root

# copy static files
COPY --from=frontend /web/dist/ /usr/share/nginx/html/

# HEALTHCHECK
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD service nginx status | grep running || exit 1
