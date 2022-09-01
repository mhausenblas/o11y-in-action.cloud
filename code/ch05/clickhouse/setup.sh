#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

# make sure ClickHouse had enough time to come up and the OTel exporter did its
# properly (set up the table):
sleep 5

echo "DESCRIBE otel_logs" | curl "http://clickhouse-server:8123/" --data-binary @-
echo "ClickHouse set up for OTel ingestion"
