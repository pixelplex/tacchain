# docker build . -t tacchaind:latest
# docker run --rm -it tacchaind:latest tacchaind --help

FROM golang:1.23.8-bullseye AS go-builder

# Install build dependencies
RUN apt-get update && apt-get install -y \
    ca-certificates \
    build-essential \
    git \
    curl \
    wget \
    libusb-1.0-0-dev \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /code
COPY . /code/

# force it to use static lib (from above) not standard libgo_cosmwasm.so file
RUN make build
RUN LEDGER_ENABLED=false make build

FROM ubuntu:22.04

COPY --from=go-builder /code/build/tacchaind /usr/bin/tacchaind

WORKDIR /opt

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

CMD ["/usr/bin/tacchaind", "version"]
