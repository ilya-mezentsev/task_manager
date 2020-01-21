### Task manager
#### Application for managing tasks by users (admin, group leads and group workers)

### Local deployment
#### Prepare local workspace (run in root project folder):
```bash
$ bash prepare_workspace.sh $(pwd)
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
