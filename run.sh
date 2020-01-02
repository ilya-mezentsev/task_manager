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

  exit 0
}

if [[ $1 = 'push_all' ]]; then
  shift
  run push_all.sh $*
elif [[ $1 = 'go_tests' ]]; then
  shift
  run go_tests.sh $*
elif [[ $1 = 'help' ]]; then
  showHelp
else
  showHelp
fi
