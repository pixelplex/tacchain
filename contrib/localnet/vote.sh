#!/bin/bash

tacchaind tx upgrade software-upgrade liquidstake \
  --upgrade-height ${1} \
  --title "Liquidstake Module Upgrade" \
  --summary "Upgrade to add liquidstake functionality" \
  --upgrade-info '{}' \
  --no-validate \
  --deposit "10000000000000000utac" \
  --from validator \
  --gas-prices "25000000000utac" \
  --yes | jq -r ".hash"

sleep .5

tacchaind tx gov vote ${2} yes --from validator --gas-prices "25000000000utac" --yes

