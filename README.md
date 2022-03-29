# API Eventos
## Introduction

This API is responsible to provide a RESTFUL interface to do a facade between a broker and the producers.
For this example, we will use as broker Kafka and we use Golang as a programming language.

## Features

- One endpoint http://localhost:7777/events/ with the GET verb
- One endpoint http://localhost:7777/events/ with the POST verb
- One endpoint http://localhost:7777/topics with the POST verb
- A service with the main responsability of write a event on Kafka topic

## Tech

API Eventos uses only Golang but we will specify the most important libs:

- github.com/gin-gonic/gin -> Used to create the web server facilities
- github.com/segmentio/kafka-go -> Used to communicate with Kafka
- context -> Used to capture the background context of the application
- encoding/json -> Used to encode my struct to json format as byte array

## Installation

To run this application you only have to open your terminal, go for the project directory and put this command:
"go run ."

## Development

This project was my initiative Gabriel Araujo Dantas a brazilian Computer Engineer and my collegue Lucas Henrique A. de Paula a brazilian Computer Scientist.

## Docker

We begin the creation of Devops part that isn't the way we want, but we start the develop of this feature. For now, we have some important considerations:

- When you work with two docker-compose you need to create the network first (
  ```
  docker network create --driver=bridge  --subnet=172.18.0.0/16  --ip-range=172.18.0.0/24  --gateway=172.18.0.1   my_network
  ```
  ), this is necessary because you have use same network on all docker-compose
  
- After that you have to add on your kafka docker-compose this (
```
networks: 
  default: 
    external: 
      name: kafka_confluent_network
```
      )
- And now you have to run all docker-compose and you are ready to use this API
```
docker-compose up -d
```
or
```
docker-compose up -d --build
```
.

## Important information

We use the JSON cloudevent.io specification so is important to send a event with these characteristics.
```
{
    "id": "ac47faa1-3f11-4d3e-8e53-41e29cdf4b0c", 
    "specVersion": "1.0", 
    "source": "/product/domain/subdomain/service", 
    "type": "br.com.example.correctTopic", 
    "time": "2022-03-22T17:41:02",
    "subject": "Estamos inserindo um evento no kafka da Confluent espia.", 
    "correlationID": "", 
    "dataContentType": "application/json", 
    "data": "{info1: \"A\"}"
}
```

If you want to create a topic you have to send this body:

```
{
    "topicName": "br.com.example.correctTopic",
    "numPartition": 1,
    "replicationFactor": 1
}
```

Another important information is we use the confluent kafka community for this test. You can find that on this link: https://github.com/confluentinc/cp-all-in-one/tree/7.0.1-post/cp-all-in-one-community 
