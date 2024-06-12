#!/bin/bash

# This script monitors the verification progress of zkEVM batches.
# Usage: ./batch_verification_monitor <verified_batches_target> <timeout>

# The number of batches to be verified.
verified_batches_target="$1"

# The script timeout (in seconds).
timeout="$2"

start_time=$(date +%s)
end_time=$((start_time + timeout))

rpc_url="$(kurtosis port print cdk-v1 zkevm-node-rpc-001 http-rpc)"
while true; do
  verified_batches="$(cast to-dec "$(cast rpc --rpc-url "$rpc_url" zkevm_verifiedBatchNumber | sed 's/"//g')")"
  echo "[$(date '+%Y-%m-%d %H:%M:%S')] Verified Batches: $verified_batches"

  current_time=$(date +%s)
  if (( current_time > end_time )); then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] ❌ Exiting... Timeout reached!"
    exit 1
  fi

  if (( verified_batches > verified_batches_target )); then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] ✅ Exiting... $verified_batches batches were verified!"
    exit 0
  fi

  sleep 10
done