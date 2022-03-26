# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.222.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.17-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# Install Just
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | sudo bash -s -- --to /usr/local/bin/

# Install xh for making http requests
RUN wget https://github.com/ducaale/xh/releases/download/v0.14.1/xh-v0.14.1-x86_64-unknown-linux-musl.tar.gz -O xh.tar.gz \
  && tar -xf xh.tar.gz \
  && install xh-*/xh /usr/local/bin/ \
  && cp xh-*/doc/xh.1 /usr/share/man/man1/ \
  && mandb

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# Install the SQLBoiler tools
USER vscode
RUN go install github.com/volatiletech/sqlboiler/v4@latest \
  && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest
