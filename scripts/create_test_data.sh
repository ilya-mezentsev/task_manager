#!/usr/bin/env bash

cookieHeader="Cookie: TM-Session-Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoie1wiaWRcIjowLFwibmFtZVwiOlwidG1fbG9naW5cIixcInJvbGVcIjpcImFkbWluXCIsXCJncm91cF9pZFwiOjB9In0.qOZxKqeER85HbC62rKu4Uhtca7X8BMcDJoY69ZwxKYk"

#echo 'create group...'
#curl -X POST localhost/api/admin/group -H "${cookieHeader}" -d '{"group_name": "my_group"}' && echo
#curl -X POST localhost/api/admin/group -H "${cookieHeader}" -d '{"group_name": "my_group2"}' && echo
#curl -X POST localhost/api/admin/group -H "${cookieHeader}" -d '{"group_name": "my_group3"}' && echo

curl localhost/api/admin/groups -H "${cookieHeader}" && echo

#echo 'create group lead...'
#curl -X POST localhost/api/admin/user -H "${cookieHeader}" \
#-d '{"user": {"name": "lead", "group_id": 1, "is_group_lead": true}}' && echo

#echo 'create group worker...'
#curl -X POST localhost/api/admin/user -H "${cookieHeader}" \
#-d '{"user": {"name": "worker", "group_id": 1, "is_group_lead": false}}' && echo

echo 'create tasks...'
curl -X POST localhost/api/admin/tasks -H "${cookieHeader}" \
  -d '{"group_id": 15, "tasks": [{"title": "task1", "description": "perform task1"},{"title": "task2", "description": "perform task2"}]}'\
  && echo
