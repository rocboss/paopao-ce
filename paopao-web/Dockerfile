FROM library/nginx

USER root

# copy static files
COPY ./dist/ /usr/share/nginx/html/

# HEALTHCHECK
HEALTHCHECK --interval=5s --timeout=3s  --retries=3  CMD service nginx status | grep running || exit 1