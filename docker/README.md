<div align="center">
  <a href="https://www.docker.com/">
    <img alt="docker" src="../logos/docker.png"/>
  </a>
  <h1>Docker</h1>
</div>

# Table of Contents

- [Installing Docker](#installing-docker)
- [See docker networks](#see-docker-networks)
- [Link ngrok to the running docker container](#link-ngrok-to-the-running-docker-container)
- [Docker Multistage Builds](#docker-multistage-builds)

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

## See docker networks

```sh
# This command lists all the docker networks. docker-compose will create a new network when you run docker-compose up
docker network ls
```

## Link ngrok to the running docker container

```sh
# The --rm flag tells docker that the container should automatically be removed after we close docker.
# The -it flag tells docker that it should open an interactive container instance, -it is short for --interactive + --tty. When you docker run with this command it takes you straight inside the container.
docker run --rm -it --link <container-name> --net <docker-network> shkoliar/ngrok http <container-name>:<running-port> --authtoken <auth-token-from-ngrok-dashboard>
```

**NOTE**

- You can see container list by running `docker container ls`
- Get your authtoken from the [ngrok's dashboard](https://dashboard.ngrok.com/login).

## Docker Multistage Builds

Defining multiple images inside a single `Dockerfile` to build the final image is multistage build. Multistage build is used to keep images as small as possible in production.
