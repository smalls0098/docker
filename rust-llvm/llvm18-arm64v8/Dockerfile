FROM arm64v8/ubuntu:20.04

LABEL maintainer="smalls smalls0098@gmail.com"

ENV LLVM_VERSION=18

ENV DEBIAN_FRONTEND=noninteractive
RUN export DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install --no-install-recommends -y \
    # install deps
    lsb-release \
    clang \
    tzdata \
    wget \
    curl \
    git \
    cmake \
    software-properties-common \
    llvm-dev && \
    gnupg && \
    # install llvm packages
    wget https://apt.llvm.org/llvm.sh && \
    chmod +x llvm.sh && \
    ./llvm.sh ${LLVM_VERSION} all && \
    rm llvm.sh && \
    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs > rustup.sh && \
    chmod +x ./rustup.sh && \
    ./rustup.sh -y && \
    root/.cargo/bin/rustup update && \
    root/.cargo/bin/rustup default nightly && \
    rm rustup.sh && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*