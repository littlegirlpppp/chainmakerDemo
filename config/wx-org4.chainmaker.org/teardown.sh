#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -e
rm -rf ./data
rm -rf ./logs
# Shut down the Docker containers for the system tests.
docker-compose -f docker-compose.yml kill && docker-compose -f docker-compose.yml down  --volumes --remove-orphans

# remove chaincode docker images
docker rm -f $(docker ps -aq)
docker rmi -f $(docker images dev-* -q)
docker volume prune
# Your system is now clean
