FROM node as node

WORKDIR /work/

COPY . .

RUN yarn
RUN yarn build

FROM nginx as nginx
# copy static files
COPY --from=node /work/dist/ /usr/share/nginx/html/

# HEALTHCHECK
HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD service nginx status | grep running || exit 1
