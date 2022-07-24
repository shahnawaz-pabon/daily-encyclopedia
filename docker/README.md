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
