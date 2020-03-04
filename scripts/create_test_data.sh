#!/usr/bin/env bash

cookieHeader="Cookie: TM-Session-Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoie1wiaWRcIjowLFwibmFtZVwiOlwidG1fbG9naW5cIixcInJvbGVcIjpcImFkbWluXCIsXCJncm91cF9pZFwiOjB9In0.qOZxKqeER85HbC62rKu4Uhtca7X8BMcDJoY69ZwxKYk"
cookieGroupLead="Cookie: TM-Session-Token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoie1wiaWRcIjoxLFwibmFtZVwiOlwia2x5a292XCIsXCJyb2xlXCI6XCJncm91cF9sZWFkXCIsXCJncm91cF9pZFwiOjF9In0.FwP77nC1bfp930P9NsfdVQXhgtRX3R-y1H3EebnZcbU"
#echo 'create group...'
#curl -X POST localhost/api/admin/group -H "${cookieHeader}" -d '{"group_name": "my_group"}' && echo
#curl -X POST localhost/api/admin/group -H "${cookieHeader}" -d '{"group_name": "my_group2"}' && echo
#curl -X POST localhost/api/admin/group -H "${cookieHeader}" -d '{"group_name": "ilyapidor"}' && echo

#curl localhost/api/admin/groups -H "${cookieHeader}" && echo

#echo 'create group lead...'
#curl -X POST localhost/api/admin/user -H "${cookieHeader}" \
#-d '{"user": {"name": "lead", "group_id": 1, "is_group_lead": true}}' && echo

#echo 'create group worker...'
#curl -X POST localhost/api/admin/user -H "${cookieHeader}" \
#-d '{"user": {"name": "worker", "group_id": 1, "is_group_lead": false}}' && echo

#echo 'create tasks...'
#curl -X POST localhost/api/admin/tasks -H "${cookieHeader}" \
#  -d '{"group_id": 1, "tasks": [{"title": "task2282e2", "description": "perform taske2e2228"}]}'\
#  && echo

#curl -X POST localhost/api/admin/user -H "${cookieHeader}" \
# -d '{"user": {"name": "nam", "group_id": 15, "is_group_lead": false}}'\
#&& echo

#curl localhost/api/admin/users -H "${cookieHeader}" && echo
#echo "pidor"
#curl -X POST localhost/api/group/lead/users -H "${cookieGroupLead}" -d '{"group_id": 1}' && echo
#echo "dude"
#curl -X POST localhost/api/group/lead/tasks -H "${cookieGroupLead}" -d '{"group_id": 1}' && echo

#curl -X POST localhost/api/group/lead/task -H "${cookieGroupLead}" -d '{"user_id": 2, "task": {"id": 1}}'

curl localhost/api/group/worker/tasks -H "${cookieGroupLead}" -d '{"user_id": 1}' && echo
