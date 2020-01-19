#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

cd ${PROJECT_ROOT} && docker-compose -f docker-compose.dev.yaml -f backend/docker-compose.yaml up $1
