#!/bin/bash

hash = tacchaind tx upgrade software-upgrade liquidstake \
  --upgrade-height ${1} \
  --title "Liquidstake Module Upgrade" \
  --summary "Upgrade to add liquidstake functionality" \
  --upgrade-info '{}' \
  --no-validate \
  --deposit "10000000000000000utac" \
  --from validator \
  --gas-prices "25000000000utac" \
  --yes | jq -r ".hash"

tacchain query tx BD49A41AF9568C12E90B67F5DB50D973FF2BB88784BFFED0C71C2F4967F9C4CA | jq -r ".raw_log"

sleep .5

tacchaind tx gov vote ${2} yes --from validator --gas-prices "25000000000utac" --yes

