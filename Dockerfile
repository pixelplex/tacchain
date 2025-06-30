# docker build . -t tacchaind:latest
# docker run --rm -it tacchaind:latest tacchaind --help

FROM golang:1.23.8-bookworm AS go-builder

# Install build dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    build-essential \
    git \
    curl \
    wget \
    libusb-1.0-0-dev \
    libc6 \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /code
COPY . /code/

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN make build
RUN LEDGER_ENABLED=false make build

FROM ubuntu:22.04

COPY --from=go-builder /code/build/tacchaind /usr/bin/tacchaind
# To run a localnet --------------------------------------
COPY --from=go-builder /code/contrib/localnet/init.sh /scripts/init.sh
COPY --from=go-builder /code/contrib/localnet/init-multi-node.sh /scripts/init-multi-node.sh
COPY --from=go-builder /code/contrib/localnet/start.sh /scripts/start.sh
COPY --from=go-builder /code/contrib/localnet/start-additional-node.sh /scripts/start-additional-node.sh
RUN chmod +x /scripts/*.sh

RUN apt-get update && apt-get install -y \
    wget \
    jq \
    bc \
    && rm -rf /var/lib/apt/lists/*

RUN wget -P /usr/lib https://github.com/CosmWasm/wasmvm/releases/download/v2.1.0/libwasmvm.x86_64.so

WORKDIR /scripts

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["/usr/bin/tacchaind", "version"]
