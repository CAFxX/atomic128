cat << 'DOCKERFILE' > Dockerfile.test
FROM golang:1.18
WORKDIR /app
COPY . /app
RUN go test -v ./...
DOCKERFILE

# This uses docker buildx to build for arm64 platform which relies on qemu integration if we are on amd64 host
docker buildx build --platform linux/arm64 -f Dockerfile.test .
