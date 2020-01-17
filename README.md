### Task manager
#### Application for managing tasks by users (admin, group leads and group workers)

### Local deployment
#### Create DB file:
```bash
$ cd task_manager/backend

$ mkdir data/

$ touch data/data.db

$ touch data/test_data.db # testing purposes
```
#### Create .env file (in root project folder) and replace default values in created .env file by your local:
```bash
$ cp .example.env .env
```
#### Build Angular project:
```bash
$ cd frontend/

$ npm install && npm run build
```
#### Start static Nginx and Golang API by docker-compose:
```bash
$ source .env && docker-compose up --build
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
##### This script can run each *.sh file from scripts folder with environment variables from .env file
