#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# make sure the collector had enough time to launch pipeline:
sleep 5

while true
do
  msg=$(uptime)
  payload="<150>`env LANG=us_US.UTF-8 date "+%b %d %H:%M:%S"` 127.0.0.1 $msg"
  echo "Sending [$payload] to OTel collector"
  echo $payload | nc -v otel-collector 54526 -w 3
  sleep 1
done
