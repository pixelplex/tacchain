# docker build . -t tacchaind:latest
# docker run --rm -it tacchaind:latest tacchaind --help

FROM golang:1.23.8-bookworm AS go-builder

RUN apt-get update && apt-get install -y \
    ca-certificates \
    build-essential \
    libusb-1.0-0-dev \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /code
COPY . /code/

RUN LEDGER_ENABLED=true make build


# --------------------------------------------------------
FROM debian:bookworm-slim

COPY --from=go-builder /code/build/tacchaind /usr/bin/tacchaind

RUN wget -P /usr/lib https://github.com/CosmWasm/wasmvm/releases/download/v2.2.1/libwasmvm.x86_64.so
RUN wget -P /usr/lib https://github.com/CosmWasm/wasmvm/releases/download/v2.1.0/libwasmvm.aarch64.so

WORKDIR /opt

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["/usr/bin/tacchaind", "version"]
