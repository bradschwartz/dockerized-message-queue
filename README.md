# Multi-Language Messaging System Test

Goal: Run multiple Docker containers, where each will:
  - Run it's own pub/sub messaging system
  - Single, unique language (MVP: Python, Golang | Reach: Scala (Utilizing Java))
  - Coordinate from single space (MVP: compose | Reach: swarm, kubernetes)

## Architecture

MVP:
- [ ]  (2) Containers -  pub/sub
    - [ ] Python
    - [ ] Golang
- [x]  (1) Docker Container - MQ server, above connect to
- [ ]  (1) docker-compose.yml - Define containers, ports, hosts by name for containers
Reach:
- [ ]  (1) Scala/JVM Container -  pub/sub
- [ ]  (1) Database Container - History of comms

## Instructions for Running Locally

To run locally, ensure [Docker](https://www.docker.com/), and [RabbitMQ](https://www.rabbitmq.com/)
 are installed. For OSX, everything can be installed using [Homebrew](https://brew.sh/):
```
brew update
brew cask install docker
brew install rabbitmq
```

Start the Docker runtime, and bring up the MQ server and pub/sub containers
with `docker-compose up`.

No publisher container exists yet, so that everything is working by manually
submitting a message to the existing queues:
```
rabbitmqadmin list queues # Get queue name
rabbitmqadmin publish routing_key=worker payload='Hello there!'
```

In the log output from the listener containers up, you should see the payload
printed back out.