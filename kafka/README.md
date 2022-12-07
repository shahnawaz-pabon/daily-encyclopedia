<div align="center">
  <a href="https://kafka.apache.org/">
    <img alt="kafka" src="../logos/kafka.png" height="100" width="100"/>
  </a>
  <h1>Kafka</h1>
</div>

# Table of Contents

- [Setup kafka via docker](#setup-kafka-via-docker)
- [Kafka commands inside the containers](#kafka-commands-inside-the-containers)

<br>

## Setup kafka via docker

```yml
version: "3"

services:
  zookeeper:
    image: wurstmeister/zookeeper
    container_name: zookeeper
    ports:
      - "2181:2181"
  kafka:
    image: wurstmeister/kafka
    container_name: kafka
    ports:
      - "9092:9092"
    environment:
      KAFKA_ADVERTISED_HOST_NAME: localhost
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
```

<br>

## Kafka commands inside the containers

```sh
## List Brokers
docker exec -ti kafka /usr/bin/broker-list.sh

## List Topics
docker exec -ti kafka /opt/kafka/bin/kafka-topics.sh --list --zookeeper zookeeper:2181

## Create a Topic
docker exec -ti kafka /opt/kafka/bin/kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic test2
```
