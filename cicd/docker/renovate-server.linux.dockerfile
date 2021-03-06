ARG MATRIX_ARCH
ARG MATRIX_ROOTFS

FROM ghcr.io/arhat-dev/builder-go:alpine as builder

ARG MATRIX_ARCH

COPY . /app
RUN dukkha golang local build renovate-server -m kernel=linux -m arch=${MATRIX_ARCH}

FROM ghcr.io/arhat-dev/go:${MATRIX_ROOTFS}-${MATRIX_ARCH}

LABEL org.opencontainers.image.source https://github.com/arhat-dev/renovate-server

ARG MATRIX_ROOTFS
COPY cicd/docker/setup.sh /setup.sh
RUN sh /setup.sh && rm -f /setup.sh

ARG MATRIX_ARCH
COPY --from=builder /app/build/renovate-server.linux.${MATRIX_ARCH} /renovate-server

ENTRYPOINT [ "/renovate-server" ]
