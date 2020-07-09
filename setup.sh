#!/usr/bin/env bash

if [ ! -f ./meilisearch ]; then
  # Install Meilisearch
  curl -L https://install.meilisearch.com | sh
fi

go build
