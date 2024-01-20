### Dockerfile builer pre-build images

```sh
docker buildx build -t bitbus/paopao-ce-backend-builder:latest -f Dockerfile.backend-builder .
docker buildx build -t bitbus/paopao-ce-backend-runner:latest -f Dockerfile.backend-runner .
cd ../../ && docker buildx build -t bitbus/paopao-ce-allinone-runner:latest -f scripts/docker/Dockerfile.allinone-runner .
```