#!/bin/bash

CHAIN_ID=${CHAIN_ID:-tacchain_2391-1}
TACCHAIND=${TACCHAIND:-$(which tacchaind)}
HOMEDIR=${HOMEDIR:-$HOME/.tacchaind}

P2P_LADDR=${P2P_LADDR:-tcp://0.0.0.0:26656}
P2P_EXTERNAL_ADDRESS=${P2P_EXTERNAL_ADDRESS:-${P2P_LADDR}}
RPC_LADDR=${RPC_LADDR:-tcp://127.0.0.1:26657}
JSON_RPC_ADDR=${JSON_RPC_ADDR:-127.0.0.1:8545}
JSON_RPC_WS_ADDR=${JSON_RPC_WS_ADDR:-127.0.0.1:8546}
GRPC_LADDR=${GRPC_LADDR:-0.0.0.0:9090}

sed -i.bak -E 's@(address = "tcp://)[^:/]+(:[0-9]+")@\10.0.0.0\2@g' "$HOMEDIR/config/app.toml"

$TACCHAIND start --chain-id $CHAIN_ID \
  --home $HOMEDIR \
  --p2p.laddr $P2P_LADDR \
  --p2p.external-address $P2P_EXTERNAL_ADDRESS \
  --rpc.laddr $RPC_LADDR \
  --json-rpc.address $JSON_RPC_ADDR \
  --json-rpc.ws-address $JSON_RPC_WS_ADDR \
  --json-rpc.enable \
  --grpc.address $GRPC_LADDR \
  --grpc.enable=true \
  --api.enable=true \
  --api.enabled-unsafe-cors \
  --home $HOMEDIR
