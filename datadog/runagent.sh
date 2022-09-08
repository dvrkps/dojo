#!/bin/bash

docker run -d --rm --name dd-agent \
  -p 127.0.0.1:8126:8126/tcp \
  -v /var/run/docker.sock:/var/run/docker.sock:ro \
  -v /proc/:/host/proc/:ro \
  -v /sys/fs/cgroup/:/host/sys/fs/cgroup:ro \
  -e DD_API_KEY=$1 \
  -e DD_APM_ENABLED=true \
  -e DD_SITE="datadoghq.eu" \
  gcr.io/datadoghq/agent:latest
