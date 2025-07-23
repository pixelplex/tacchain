#!/bin/bash

CHAIN_ID=${CHAIN_ID:-tacchain_2391-1}
TACCHAIND=${TACCHAIND:-$(which tacchaind)}
HOMEDIR=${HOMEDIR:-$HOME/.tacchaind}

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
  --api.enabled-unsafe-cors
--home $HOMEDIR
