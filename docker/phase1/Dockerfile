FROM debian:buster-slim as graylog-downloader

ARG VCS_REF
ARG GRAYLOG_VERSION

WORKDIR /tmp

# hadolint ignore=DL3008,DL3015
RUN \
  apt-get update  > /dev/null && \
  apt-get install --assume-yes \
    ca-certificates \
    curl > /dev/null

RUN \
  curl \
    --silent \
    --location \
    --retry 3 \
    --output "/tmp/graylog-${GRAYLOG_VERSION}.tgz" \
    "https://packages.graylog2.org/releases/graylog/graylog-${GRAYLOG_VERSION}.tgz"

RUN \
  curl \
    --silent \
    --location \
    --retry 3 \
    --output "/tmp/graylog-${GRAYLOG_VERSION}.tgz.sha256.txt" \
    "https://packages.graylog2.org/releases/graylog/graylog-${GRAYLOG_VERSION}.tgz.sha256.txt"

RUN \
  sha256sum --check "graylog-${GRAYLOG_VERSION}.tgz.sha256.txt"

RUN \
  mkdir /opt/graylog && \
  tar --extract --gzip --file "/tmp/graylog-${GRAYLOG_VERSION}.tgz" --strip-components=1 --directory /opt/graylog

RUN \
  install \
    --directory \
    --mode=0750 \
    /opt/graylog/data \
    /opt/graylog/data/journal \
    /opt/graylog/data/log \
    /opt/graylog/data/config \
    /opt/graylog/data/plugin \
    /opt/graylog/data/data

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name="Internal Services" \
      org.label-schema.description="Greylog2" \
      org.label-schema.url="https://docker.freighttrust.com/" \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/freight-trust/action-docker.git" \
      org.label-schema.vendor="Freight Trust & Clearing" \
      org.label-schema.version=$VERSION \
      org.label-schema.schema-version="1.0"