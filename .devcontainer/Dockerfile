# [Choice] Go version (use -bookworm, or -bullseye variants on local arm64/Apple Silicon): 1, 1.21, 1.20, 1-bookworm, 1.21-bookworm, 1.20-bookworm, 1-bullseye, 1.21-bullseye, 1.20-bullseye
ARG VARIANT=1.21-bookworm
FROM golang:${VARIANT}

# Follow install instructions for ebitengine https://ebitengine.org/en/documents/install.html
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends gcc

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config
