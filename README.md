# API Eventos
## Introduction

This API is responsible to provide a RESTFUL interface to do a facade between a broker and the producers.
For this example, we will use as broker Kafka and we use Golang as a programming language.

## Features

- One endpoint http://localhost:7777/events/ with the GET verb
- One endpoint http://localhost:7777/events/ with the POST verb
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

- When you work with two docker-compose you need to create the network first (docker network create your_network)
- You have to put this network on all docker-compose (Now all container will be in the same subnet)
- And that is our challenge how I can connect kafka in a container with a client in a container? I seriously don't know yet.