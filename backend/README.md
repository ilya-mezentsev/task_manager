### Golang API overview

#### On each error API returns:
```json5
{
  "status": "error",
  "error_detail": "error description"
}
```

#### Authorization endpoints:
##### GET /api/session/
```bash
$ curl localhost/api/session/ -H "Cookie: TM-Session-Token=$sessionToken"
```
##### Returns user session by session token:
```json5
{
  "status":"ok",
  "data": {
    "id": 0,
    "name": "tm_login",
    "role": "admin",
    "group_id": 0
  }
}
```
##### POST /api/session/login
```bash
$ curl -X POST localhost/api/session/login -d '{"user_name": "tm_login", "user_password": "tm_password"}' -i
```
##### Returns header with cookie and user session:
```
HTTP/1.1 200 OK
Server: nginx
Date: Sun, 19 Jan 2020 10:12:53 GMT
Content-Type: application/json
Content-Length: 64
Connection: keep-alive
Set-Cookie: TM-Session-Token=sessionToken; Path=/; Max-Age=3600; HttpOnly
```
```json5
{
  "status":"ok",
  "data": {
    "id": 0,
    "name": "tm_login",
    "role": "admin",
    "group_id": 0
  }
}
```
##### POST /api/session/logout
```bash
$ curl -X POST localhost/api/session/logout -H "Cookie: TM-Session-Token=$sessionToken" -i
```
##### Returns header with empty cookie and default body:
```
HTTP/1.1 200 OK
Server: nginx
Date: Sun, 19 Jan 2020 10:15:54 GMT
Content-Type: application/json
Content-Length: 27
Connection: keep-alive
Set-Cookie: TM-Session-Token=; Path=/; Expires=Thu, 01 Jan 1970 00:00:00 GMT
```
```json5
{
  "status": "ok",
  "data": null
}
```

#### Admin endpoints
##### GET /api/admin/groups
```bash
$ curl localhost/api/admin/groups -H "Cookie: TM-Session-Token=$sessionToken"
```
##### Returns list of all groups:
```json5
{
  "status": "ok",
  "data": [
    {"id":  1, "name":  "group1"},
    {"id":  2, "name":  "group2"}
  ]
}
```
##### POST /api/admin/group
```bash
$ curl -X POST localhost/api/admin/group -H "Cookie: TM-Session-Token=$sessionToken" -d '{"group_name": "group3"}'
```
##### Creates new group and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```
##### DELETE /api/admin/group
```bash
$ curl -X DELETE localhost/api/admin/group -H "Cookie: TM-Session-Token=$sessionToken" -d '{"group_id": 1}'
```
##### Deletes group by passed id and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```
##### GET /api/admin/users
```bash
$ curl localhost/api/admin/users -H "Cookie: TM-Session-Token=$sessionToken"
```
##### Returns list of all users:
```json5
{
  "status": "ok",
  "data": [
    {"id":  1, "name":  "name1", "group_id":  1, "password":  "some_pass", "is_group_lead":  false},
    {"id":  2, "name":  "name2", "group_id":  1, "password":  "another_pass", "is_group_lead":  true},
    {"id":  3, "name":  "name3", "group_id":  2, "password":  "another_pass", "is_group_lead":  true}
  ]
}
```
##### POST /api/admin/user
```bash
$ curl -X POST localhost/api/admin/users -H "Cookie: TM-Session-Token=$sessionToken" -d '{"user": {"name": "name4", "group_id": 1, "is_group_lead": false}}'
```
##### Creates new user and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```
##### DELETE /api/admin/user
```bash
$ curl -X DELETE localhost/api/admin/user -H "Cookie: TM-Session-Token=$sessionToken" -d '{"user_id": 1}'
```
##### Deletes user by passed id and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```
##### GET /api/admin/tasks
```bash
$ curl localhost/api/admin/tasks -H "Cookie: TM-Session-Token=$sessionToken"
```
##### Returns list of all tasks:
```json5
{
  "status": "ok",
  "data": [
    {"id":  1, "title":  "title1", "description":  "some description", "group_id":  1, "user_id":  1, "is_complete":  false, "comment":  ""},
    {"id":  2, "title":  "title2", "description":  "description", "group_id":  1, "user_id":  2, "is_complete":  true, "comment":  "done"},
  ]
}
```
##### POST /api/admin/tasks
```bash
$ curl -X POST localhost/api/admin/tasks -H "Cookie: TM-Session-Token=$sessionToken" -d '{"group_id": 1, "tasks": [{"title": "some_title", "description": "hello world"}]}'
```
##### Assign passed tasks to work group and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```
##### DELETE /api/admin/task
```bash
$ curl -X DELETE localhost/api/admin/user -H "Cookie: TM-Session-Token=$sessionToken" -d '{"task_id": 1}'
```
##### Deletes task by passed id and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```

#### Group lead endpoints
##### GET /api/group/lead/users
```bash
$ curl localhost/api/group/lead/users -H "Cookie: TM-Session-Token=$sessionToken" -d '{"group_id": 1}'
```
##### Returns list of users belong to group:
```json5
{
  "status": "ok",
  "data": [
    {"id":  1, "name":  "name1", "group_id":  1, "password":  "some_pass", "is_group_lead":  false},
    {"id":  2, "name":  "name2", "group_id":  1, "password":  "another_pass", "is_group_lead":  true}
  ]
}
```
##### GET /api/group/lead/tasks
```bash
$ curl localhost/api/group/lead/tasks -H "Cookie: TM-Session-Token=$sessionToken" -d '{"group_id": 1}'
```
##### Returns list of tasks belong to group:
```json5
{
  "status": "ok",
  "data": [
    {"id":  1, "title":  "title1", "description":  "some description", "group_id":  1, "user_id":  1, "is_complete":  false, "comment":  ""},
    {"id":  2, "title":  "title2", "description":  "description", "group_id":  1, "user_id":  2, "is_complete":  true, "comment":  "done"},
  ]
}
```
##### POST /api/group/lead/task
```bash
$ curl -X POST localhost/api/group/lead/task -H "Cookie: TM-Session-Token=$sessionToken" -d '{"user_id": 1, "task": {"id": 1}}'
```
##### Assign task to worker and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```

#### Group worker endpoints
##### GET /api/group/worker/tasks
```bash
$ curl localhost/api/group/worker/tasks -H "Cookie: TM-Session-Token=$sessionToken" -d '{"user_id": 1}'
```
##### Returns tasks belong to user:
```json5
{
  "status": "ok",
  "data": [
    {"id":  1, "title":  "title1", "description":  "some description", "group_id":  1, "user_id":  1, "is_complete":  false, "comment":  ""}
  ]
}
```
##### POST /api/group/worker/task/comment
```bash
$ curl localhost/api/group/worker/task/comment -H "Cookie: TM-Session-Token=$sessionToken" -d '{"task_id": 1, "comment": "hello world"}'
```
##### Adds comment to task and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```
##### POST /api/group/worker/task/complete
```bash
$ curl localhost/api/group/worker/task/complete -H "Cookie: TM-Session-Token=$sessionToken" -d '{"task_id": 1}'
```
##### Mark task as complete and returns default body:
```json5
{
  "status": "ok",
  "data": null
}
```
