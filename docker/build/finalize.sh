#!/bin/sh

set -xe

# Create paopao user for paopao-ce
addgroup -S paopao
adduser -G paopao -H -D -g 'paopao User' paopao -h /data/paopao -s /bin/bash && usermod -p '*' paopao && passwd -u paopao
echo "export PAOPAO_CUSTOM=${PAOPAO_CUSTOM}" >> /etc/profile

# Final cleaning
mv /app/docker/paopao-ce_config.yaml /app/config.yaml
rm -rf /app/docker/build
rm /app/docker/nsswitch.conf
rm /app/docker/README.md
