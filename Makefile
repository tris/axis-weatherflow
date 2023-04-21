.PHONY: all clean

all:
	DOCKER_BUILDKIT=1 docker build --build-arg ARCH=armv7hf --build-arg GOARCH=arm -o type=local,dest=. .
	DOCKER_BUILDKIT=1 docker build --build-arg ARCH=aarch64 --build-arg GOARCH=arm64 -o type=local,dest=. .

clean:
	rm -rf *.eap
