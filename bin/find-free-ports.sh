#!/usr/bin/env bash
set -euo pipefail

# Emit PORT1=<port>, PORT2=<port>, ... to stdout.
# Usage: find-free-ports.sh [count]   (default: 3)
# Picks ports in 10000-49999 to avoid the typical ephemeral range
# and leave headroom for callers that need a paired companion port.

count=${1:-3}

in_use() {
  lsof -iTCP:"$1" -sTCP:LISTEN -nP 2>/dev/null | grep -q .
}

pick_port() {
  local picked="$1" p
  while :; do
    p=$(( ((RANDOM << 15) | RANDOM) % 40000 + 10000 ))
    [[ " $picked " == *" $p "* ]] && continue
    in_use "$p" && continue
    echo "$p"
    return
  done
}

picked=""
for ((i=1; i<=count; i++)); do
  p=$(pick_port "$picked")
  picked="$picked $p"
  echo "PORT${i}=$p"
done
