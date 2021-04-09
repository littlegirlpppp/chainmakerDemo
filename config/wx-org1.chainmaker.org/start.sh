#!/bin/bash
set -ev
export MSYS_NO_PATHCONV=1
docker-compose -f docker-compose.yml up -d chainmaker
docker ps