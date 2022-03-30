# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.222.0/containers/rust/.devcontainer/base.Dockerfile

# [Choice] Debian OS version (use bullseye on local arm64/Apple Silicon): buster, bullseye
ARG VARIANT="buster"
FROM mcr.microsoft.com/vscode/devcontainers/rust:0-${VARIANT}

# Install Just
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | sudo bash -s -- --to /usr/local/bin/

# Install xh for making http requests
# RUN cargo install xh # Takes too long to build
ARG XH_VERSION="v0.15.0"
RUN wget https://github.com/ducaale/xh/releases/download/${XH_VERSION}/xh-${XH_VERSION}-x86_64-unknown-linux-musl.tar.gz -O xh.tar.gz \
  && tar -xf xh.tar.gz \
  && install xh-*/xh /usr/local/bin/ \
  && cp xh-*/doc/xh.1 /usr/share/man/man1/ \
  && mandb
