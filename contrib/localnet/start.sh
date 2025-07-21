#!/bin/bash

CHAIN_ID=${CHAIN_ID:-tacchain_2391-1}
TACCHAIND=${TACCHAIND:-$(which tacchaind)}
HOMEDIR=${HOMEDIR:-$HOME/.tacchaind}

sed -i.bak -E 's@(address = "tcp://)[^:/]+(:[0-9]+")@\10.0.0.0\2@g' "$HOMEDIR/config/app.toml"

$TACCHAIND start --chain-id $CHAIN_ID \
  --home $HOMEDIR
