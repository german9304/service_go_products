# restful api

stand alone service application

## Development

* docker-compose build - builds images
* docker-compose up - starts services
* docker-compose start - database service will start and run 

## Built with 

* Docker
* docker-compose
* PostgreSQL
* golang

## Testing

`go run test ./...`

## Environment

To add a database url add an `.env` file with docker:

`MODE=DOCKER`
`DATABASE_DOCKER_URL={url}`

To add a database url add an `.env` file with no docker:

`MODE=NODOCKER`
`DATABASE_DEV_URL={url}`

## How to run the server:

Inside cmd/ folder there is an executable main file 
with endpoints hosted on port `localhost:80`. 

Curl

`curl -v localhost:8080/api/products`