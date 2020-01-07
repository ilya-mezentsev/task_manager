#!/usr/bin/env bash

function run() {
  if [[ -f ./.env ]]; then
    set -o allexport
    source ./.env
    set +o allexport
    scriptName=$1
    shift
    bash ${PROJECT_ROOT}/scripts/${scriptName} "$(echo $*)"
  fi
}

function showHelp {
  echo 'usage bash run.sh <command>'
  echo 'available commands:'
  echo -e '\t help (show this help)'
  echo -e '\t push_all (push all files to git)'
  echo -e '\t go_tests (run tests for golang)'
}

if [[ $1 = '-h' || $1 = 'help' || $1 = '-help' ]]; then
  showHelp
  exit 0
fi

run "$1.sh" $*
