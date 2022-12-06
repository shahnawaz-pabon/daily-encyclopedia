<div align="center">
  <a href="https://kafka.apache.org/">
    <img alt="kafka" src="../logos/kafka.png" height="100" width="100"/>
  </a>
  <h1>Kafka</h1>
</div>

# Table of Contents

- [Setup kafka via docker](#setup-kafka-via-docker)

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
