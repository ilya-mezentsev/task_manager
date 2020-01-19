### Task manager
#### Application for managing tasks by users (admin, group leads and group workers)

### Local deployment
#### Create DB file:
```bash
$ cd task_manager/backend

$ mkdir data/

$ touch data/data.db
```
#### Create .env file (in root project folder) and replace default values in created .env file by your local:
```bash
$ cp .example.env .env
```
#### This command will start Angular serve (hot-reload), Golang API and Nginx (proxy server to another two containers):
```bash
$ bash run.sh dev [--build] # use --build flag for first building of containers
```
#### Open in browser:
```
http://localhost:80
```

#### You can use task runner - script run.sh
#### To see what it supports, type (without arguments):
```bash
$ bash run.sh
```
#### This script can run each *.sh file from scripts folder with environment variables from .env file
