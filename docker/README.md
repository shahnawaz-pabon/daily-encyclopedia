<div align="center">
  <a href="https://www.docker.com/">
    <img alt="docker" src="../logos/docker.png"/>
  </a>
  <h1>Docker</h1>
</div>

# Table of Contents

- [Installing Docker](#installing-docker)
- [Run docker as a non-root user](#run-docker-as-a-non-root-user)
- [See docker networks](#see-docker-networks)
- [Link ngrok to the running docker container](#link-ngrok-to-the-running-docker-container)
- [Setup SQL Server](#setup-sql-server)
- [Docker Multistage Builds](#docker-multistage-builds)
- [List and remove all images](#list-and-remove-all-images)
- [List and remove all exited containers](#list-and-remove-all-exited-containers)
- [Restart container at computer boot](#restart-container-at-computer-boot)
- [Dockerize Laravel application with newrelic](#dockerize-laravel-application-with-newrelic)
- [Dockerize FastAPI application with newrelic](#dockerize-fastapi-application-with-newrelic)
- [Install python packages into the docker container](#install-python-packages-into-the-docker-container)
- [Copy a file from host to docker container and vice versa](#copy-a-file-from-host-to-docker-container-and-vice-versa)
- [Get fake REST APIs with json-server](#get-fake-rest-apis-with-json-server)

## Installing Docker

This procedure was applied on the `Ubuntu 20.04` OS

```sh
# updating existing list of packages
sudo apt update
# installing a few prerequisite packages which let apt use packages over HTTPS
sudo apt install apt-transport-https ca-certificates curl software-properties-common
# adding GPG key for the official Docker repository in the system
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
# adding the Docker repository to APT sources, this will also update our package database with the Docker packages from the newly added repo
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable"
# checking whether docker is going to be installed from the Docker repo instead of the default Ubuntu repo:
apt-cache policy docker-ce
# installing Docker
sudo apt install docker-ce
# check that it’s running
sudo systemctl status docker
```

<br>

## Run docker as a non-root user

Docker commands can be run under `root` user or users that are under `docker` group.

```sh
# add docker group if it doesn't already exist
sudo groupadd docker
# -aG option will append the current user to the docker group
sudo usermod -aG docker ${USER}
# user that is not logged in to the system
sudo usermod -aG docker username
# Log out and log back in so that your group membership is re-evaluated or run this
# The newgrp command is used to change the current group ID during a login session
newgrp docker # activate the changes to groups
docker images # check docker images without sudo
```

<br>

## See docker networks

```sh
# This command lists all the docker networks. docker-compose will create a new network when you run docker-compose up
docker network ls
```

<br>

## Link ngrok to the running docker container

```sh
# The --rm flag tells docker that the container should automatically be removed after we close docker.
# The -it flag tells docker that it should open an interactive container instance, -it is short for --interactive + --tty. When you docker run with this command it takes you straight inside the container.
docker run --rm -it --link <container-name> --net <docker-network> shkoliar/ngrok http <container-name>:<running-port> --authtoken <auth-token-from-ngrok-dashboard>
```

**NOTE**

- You can see container list by running `docker container ls`
- Get your authtoken from the [ngrok's dashboard](https://dashboard.ngrok.com/login).

<br>

## Setup SQL Server

```sh
# Pull, run and bind with local port
sudo docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=<YourStrong@Passw0rd>" \
   -p 1433:1433 --name sql1 --hostname sql1 \
   -d mcr.microsoft.com/mssql/server:2019-latest
```

You can connect to the SQL Server using the `sqlcmd` tool inside of the container by using the following command.

```sh
docker exec -it <container_id|container_name> /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P <your_password>
```

Alternatively you can do the following as well

```sh
# Here sql1 is the hostname
docker exec -it sql1 "bash"
/opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P "<YourNewStrong@Passw0rd>"
```

Test whether you can see databases

```sh
CREATE DATABASE TestDB
SELECT Name from sys.Databases
```

Don't forget to run `GO` in the end for executing the previous commands.

```sh
GO
```

<br>

## Docker Multistage Builds

Defining multiple images inside a single `Dockerfile` to build the final image is multistage build. Multistage build is used to keep images as small as possible in production. [See this](single-stage/README.md) with a real-life example.

<br>

## List and remove all images

- List images

```sh
docker images -a
```

- Remove all images

```sh
docker rmi $(docker images -a -q)
```

<br>

## List and remove all exited containers

- List exited containers

```sh
docker ps -a -f status=exited
```

- Remove exited containers

```sh
docker rm $(docker ps -a -f status=exited -q)
```

<br>

## Restart container at computer boot

```sh
docker update --restart unless-stopped <container_id>
```

<br>

## Dockerize Laravel application with newrelic

**`Dockerfile`**

```Dockerfile
FROM ubuntu:bionic

RUN apt-get update && apt-get -y upgrade && apt-get install -y software-properties-common tzdata
RUN echo "Asia/Karachi" > /etc/timezone && rm -f /etc/localtime && dpkg-reconfigure -f noninteractive tzdata
RUN add-apt-repository -y ppa:ondrej/php
RUN apt-get update
RUN apt install -y wget php8.0 php8.0-fpm php8.0-common php8.0-mysql php8.0-xml php8.0-dev php8.0-xmlrpc php8.0-curl php8.0-mbstring php8.0-zip
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
RUN php composer-setup.php
RUN mv composer.phar /usr/local/bin/composer
RUN apt-get install -y php8.0-mongodb unzip
RUN apt-get install -y php8.0-gd
#RUN apt-get install -y mysql-client
WORKDIR /app
COPY . .



## New relic installation
#RUN wget -r -l1 -nd -A"linux.tar.gz" https://download.newrelic.com/php_agent/release/
RUN wget https://download.newrelic.com/php_agent/release/newrelic-php5-9.18.1.303-linux.tar.gz

# Unzip it and enter it
RUN gzip -dc newrelic*.tar.gz | tar xf -
WORKDIR '/app/newrelic-php5-9.18.1.303-linux'
RUN ls -l

# Delete any previous New Relic extension
RUN rm -f /usr/lib/php/20200930/newrelic.so

# Copy the downloaded extension to your php's extension directory
# If is ZTS is compiled in, the file name should have "-zts" appended. See examples below.
RUN cp ./agent/x64/newrelic-20200930.so /usr/lib/php/20200930/newrelic.so

# Install New Relic's daemon
RUN cp ./daemon/newrelic-daemon.x64 /usr/bin/newrelic-daemon

WORKDIR /app

# Copy your newrelic.ini custom file to your php conf.d dir. Make sure your php reads .ini files from this directory.
RUN cp /app/newrelic.ini /etc/php/8.0/fpm/conf.d/20-newrelic.ini
RUN cp /app/newrelic.ini /etc/php/8.0/cli/conf.d/20-newrelic.ini
```

<br>

**`init.sh`**

```sh
#!/bin/bash
#mysql -uroot -pexample --host mysql -e "CREATE DATABASE IF NOT EXISTS maya_local;"
composer install --ignore-platform-reqs
#php artisan sentry:publish --dsn=http://cbe2201f8f374c9785b10697886ce902@3.0.117.15:9000/3
php artisan optimize:clear
php artisan cache:clear
php artisan config:clear
php artisan key:generate
php artisan passport:keys
#php artisan migrate
#php artisan db:seed

#New relic
/usr/bin/newrelic-daemon start

# Server application
php artisan serve --host 0.0.0.0
```

<br>

**`docker-compose.yml`**

```yml
version: "3"

services:
  client:
    build: ./
    command: tail -f /dev/null
    ports:
      - 5500:8000
    #    depends_on:
    #      - mysql
    volumes:
      - ./:/app
    entrypoint: ./init.sh

  #  mysql:
  #    image: mariadb
  #    environment:
  #      MYSQL_ROOT_PASSWORD: example
  #      MYSQL_USER: root
  #    volumes:
  #      - ./mysqldbdir:/var/lib/mysql
  #    ports:
  #      - 3307:3306

  redis:
    image: redis:latest
    ports:
      - 6379:6379

networks:
  default:
    external:
      name: custom-bot
```

<br>

## Dockerize FastAPI application with newrelic

**`Dockerfile`**

```Dockerfile
# set base image (host OS)
FROM python:3.8-buster

# set work directory
WORKDIR /app

# set env variables
## Prevents Python from writing pyc files to disc (equivalent to python -B)
ENV PYTHONDONTWRITEBYTECODE 1
## PYTHONUNBUFFERED: Prevents Python from buffering stdout and stderr (equivalent to python -u)
ENV PYTHONUNBUFFERED 1

RUN pip install --no-cache-dir newrelic

# copy the dependencies file to the working directory
COPY requirements.txt .

# install dependencies
RUN pip install -r requirements.txt

ENV NEW_RELIC_LOG=stderr \
    NEW_RELIC_LOG_LEVEL=info \
    NEW_RELIC_ENABLED=true \
    NEW_RELIC_DISTRIBUTED_TRACING_ENABLED=true \
    NEW_RELIC_LICENSE_KEY=<license_key> \
    NEW_RELIC_APP_NAME="app-name"

# copy project
COPY . .

CMD ["newrelic-admin", "run-program", "gunicorn", "app.main:app", "--workers 2", "--worker-class", "uvicorn.workers.UvicornWorker", "--bind", "0.0.0.0"]
```

<br>

**`docker-compose.yml`**

```yml
# docker-compose.yml
version: "3"
services:
  web:
    build: .
    # command: gunicorn app.main:app --workers 2 --worker-class uvicorn.workers.UvicornWorker --bind 0.0.0.0
    volumes:
      - .:/app
    ports:
      - 8001:8000
```

<br>

## Install python packages into the docker container

```bash
$ docker exec -it <CONTAINER_ID> /bin/bash # open the container's console
$ pip install <package_name>
```

<br>

## Copy a file from host to docker container and vice versa

Copy from host to docker container.

```bash
docker cp <container_id>:/path/to/file.ext .
```

<br>

Copy from docker container to host.

```sh
docker cp file.ext <container_id>:/path/to/file.ext
```

<br>

## Get fake REST APIs with json-server

Get a full fake REST API with zero coding in less than 30 seconds with [json-server](https://github.com/typicode/json-server).

> **`docker-compose.yml`**

```Dockerfile
version: "3.8"
services:
  json-server:
    image: vimagick/json-server
    command: -H 0.0.0.0 -p 3000 -w db.json
    init: true
    ports:
      - "3000:3000"
    volumes:
      - ./data:/data
    restart: unless-stopped
```
