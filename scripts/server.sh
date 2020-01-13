#!/usr/bin/env bash
if [[ $ENV_VARS_WERE_SET != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

cd ${GOPATH} && go run main.go
