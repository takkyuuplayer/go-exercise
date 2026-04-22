#!/usr/bin/env bash
set -euo pipefail

# Emit REDIS_CLUSTER_NODE_{1,2,3}_PORT=<port> to stdout.
# Ports are picked in 7000-54999 so that the cluster bus port
# (announce-port + 10000) stays within the 65535 TCP limit.

in_use() {
  lsof -iTCP:"$1" -sTCP:LISTEN -nP 2>/dev/null | grep -q .
}

pick_port() {
  local picked="$1" p
  while :; do
    p=$(( ((RANDOM << 15) | RANDOM) % 48000 + 7000 ))
    [[ " $picked " == *" $p "* ]] && continue
    in_use "$p" && continue
    echo "$p"
    return
  done
}

picked=""
for i in 1 2 3; do
  p=$(pick_port "$picked")
  picked="$picked $p"
  echo "REDIS_CLUSTER_NODE_${i}_PORT=$p"
done
