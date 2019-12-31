#!/usr/bin/env bash

function run() {
  if [[ -f ./.env ]]; then
    set -o allexport
    source ./.env
    set +o allexport
    scriptName=$1
    shift
    bash ${PROJECT_ROOT}/scripts/${scriptName} $*
  fi
}

function showHelp {
  echo 'usage bash run.sh command'
  echo 'available commands:'
  echo -e '\t -h (show this help)'
  echo -e '\t push_all (push all files to git)'

  exit 0
}

if [[ $1 = 'push_all' ]]; then
  shift
  run pushAll.sh "$(echo $*)"
elif [[ $1 = '-h' ]]; then
  showHelp
else
  showHelp
fi
