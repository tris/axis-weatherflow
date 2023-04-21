ARG ARCH=armv7hf
ARG GOARCH=arm
ARG VERSION=1.3
ARG UBUNTU_VERSION=22.04
ARG REPO=axisecp
ARG SDK=acap-native-sdk

FROM golang:1.20-alpine AS builder
ARG GOARCH
ENV GOARCH=${GOARCH}
WORKDIR /opt/build
RUN apk add --no-cache upx
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o weatherflow .
RUN upx --lzma weatherflow

FROM ${REPO}/${SDK}:${VERSION}-${ARCH}-ubuntu${UBUNTU_VERSION} AS sdk
COPY ./app /opt/app/
COPY --from=builder /opt/build/weatherflow /opt/app/lib/
WORKDIR /opt/app
RUN . /opt/axis/acapsdk/environment-setup* && acap-build ./

FROM scratch
COPY --from=sdk /opt/app/*.eap /
CMD ""
