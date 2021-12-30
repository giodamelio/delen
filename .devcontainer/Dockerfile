# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.208.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.17-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# Install Just
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | sudo bash -s -- --to /usr/local/bin/

# Install some system utils
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends bzip2

# Install Goose Db Migrations tools
RUN wget https://github.com/pressly/goose/releases/download/v3.5.0/goose_linux_x86_64 -O goose \
  && install goose /usr/local/bin

# Install SQLC to generate Go sql code
RUN wget https://github.com/kyleconroy/sqlc/releases/download/v1.11.0/sqlc_1.11.0_linux_amd64.zip -O sqlc.zip \
  && unzip sqlc.zip \
  && install sqlc /usr/local/bin

# Install xh for making http requests
RUN wget https://github.com/ducaale/xh/releases/download/v0.14.1/xh-v0.14.1-x86_64-unknown-linux-musl.tar.gz -O xh.tar.gz \
  && tar -xf xh.tar.gz \
  && install xh-*/xh /usr/local/bin/ \
  && cp xh-*/doc/xh.1 /usr/share/man/man1/ \
  && mandb

# Download usql sql cli
# Due to an outdated glibc in bullseye, install it via go get
# That doesn't guerentee the version, but that should be fine in this case
# RUN wget https://github.com/xo/usql/releases/download/v0.9.5/usql-0.9.5-linux-amd64.tar.bz2 -O usql.zip \
#   && tar -xf usql.zip \
#   && install usql /usr/local/bin
USER vscode
RUN go install github.com/xo/usql@master