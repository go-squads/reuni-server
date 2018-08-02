# Reuni API

## Read this docs for more complete documentation: [https://go-squads.gitbook.io/reuni]

Reuni system is a centralized service re-configuration manager for microservices architecture.

This repository is intended for the RESTful API server for the reuni system. This API provide:
- Service Configuration 
- Authorization Service for reuni system
- Authentication Service for reuni system

## Development Environment


### Dockerized 

Prerequisite:
- Docker 18.03.1 or later
- docker-compose 1.21.1 or later

To build application:
`docker-compose build`

To start application:
`docker-compose up`

### Non-Docker

#### MacOS / OSX

Prerequisite:
- Go 1.10.3 - `brew install go`
- PostgreSQl 10.4 - `brew install postgresql`
- Homebrew

To build application:
`make build`

To test application:
`make test`

To run application:
`make run`

To install application:
`make install`

Authors: GO-SQUAD Tech 2.0 Team Bravo
- [malamsyah](https://github.com/malamsyah)
- [rifkiadrn](https://github.com/rifkiadrn)
- [vin0298](https://github.com/vin0298)
- [vincentius15](https://github.com/vincentius15)
- [reynaldipane](https://github.com/reynaldipane)
- [kanisiuskenneth](https://github.com/kanisiuskenneth)
