#!/usr/bin/env bash
if [[ ${ENV_VARS_WERE_SET} != '1' ]]; then
  echo 'env variables are not set'
  exit 1
fi

folders=()
function setFoldersWithTests() {
  cd ${GOPATH}/src

  for dir in $(find . -type d)
  do
    if tests=$(find ${GOPATH}/src/${dir} -maxdepth 1 -name '*_test.go'); [[ ${tests} != "" ]]; then
      folders+=(${dir})
    fi
  done
}

setFoldersWithTests
for dir in "${folders[@]}"
do
  reportFileName=$(echo -n ${dir} | md5sum | awk '{print $1}')
  cd ${GOPATH}/src/${dir} && go test -coverprofile=${REPORT_FOLDER}/${reportFileName}.out
  if [[ $1 = html ]]; then # open reports in browser
    go tool cover -html=${REPORT_FOLDER}/${reportFileName}.out
  fi
done
