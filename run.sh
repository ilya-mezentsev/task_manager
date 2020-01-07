#!/usr/bin/env bash

SCRIPTS_FOLDER=$(pwd)/scripts

function run() {
  if [[ -f ./.env ]]; then
    set -o allexport
    source ./.env
    set +o allexport
    scriptName=$1
    shift
    bash ${SCRIPTS_FOLDER}/${scriptName} "$(echo $*)"
  else
    echo file $(pwd)/.env not found
    exit 1
  fi
}

function showHelp {
  echo 'usage bash run.sh <command>'
  echo 'available commands:'
  printf '\t-h, -help, help - show this help\n'
  find ${SCRIPTS_FOLDER} -type f -printf "\t%f\n" | sed 's/\.sh$//1'
}

if [[ $1 = '-h' || $1 = 'help' || $1 = '-help' || $1 = '' ]]; then
  showHelp
  exit 0
fi

scriptName="$1.sh"
if [[ -f ${SCRIPTS_FOLDER}/${scriptName} ]]; then
  run ${scriptName} $*
else
  echo file ${SCRIPTS_FOLDER}/${scriptName} not found
  exit 1
fi
