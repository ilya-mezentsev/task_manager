#!/usr/bin/env bash

function prepareFolders() {
  mkdir $1/backend/test_report
  mkdir $1/backend/data
}

function prepareFiles() {
  touch $1/.env
  touch $1/backend/data/data.db
  touch $1/backend/data/test_data.db
}

function compileAngularProject() {
  cd $1/frontend && npm install
}

rootFolder=$1
if [[ ${rootFolder} = '' ]]; then
  echo 'root folder was not provided'
  exit 1
fi

envVars=(
  "ENV_VARS_WERE_SET=1"
  "PROJECT_ROOT=${rootFolder}"
  "REPORT_FOLDER=${rootFolder}/backend/test_report"
  "GOPATH=${rootFolder}/backend"
  "TEST_DB_FILE=${rootFolder}/backend/data/test_data.db"
  "DB_FILE=${rootFolder}/backend/data/data.db"
  "CODER_KEY=123456789012345678901234"
  "FRONTEND_DIR=${rootFolder}/frontend"
  "STATIC_DIR=${rootFolder}/frontend/dist/task_manager"
)

prepareFolders ${rootFolder}
prepareFiles ${rootFolder}
compileAngularProject ${rootFolder}

for envVar in ${envVars[@]}; do
  echo ${envVar} >> ${rootFolder}/.env
done
