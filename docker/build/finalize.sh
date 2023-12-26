#!/bin/sh

set -xe

# Create paopao user for paopao-ce
addgroup -S paopao
adduser -G paopao -H -D -g 'paopao User' paopao -h /app -s /bin/sh && usermod -p '*' paopao && passwd -u paopao
# echo "export PAOPAO_CUSTOM=${PAOPAO_CUSTOM}" >> /etc/profile

# Final cleaning
mv /app/docker/config.yaml /app/config.yaml
rm -rf /app/docker/build
rm /app/docker/README.md
