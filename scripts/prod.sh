#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

cd ${FRONTEND_DIR} && npm run build && cd ${PROJECT_ROOT} \
&& docker-compose -f docker-compose.prod.yaml -f backend/docker-compose.yaml up $1
