# restful api

## Development

* docker-compose build - builds images
* docker-compose up - starts services
* docker-compose start - database service will start and run 

## Built with 

* Docker
* PostgreSQL
* golang

## Example

Inside cmd/ folder there is an executable main file 
with endpoints (below) hosted on port `localhost:80`. 

* localhost:8080/api/products - GET
* localhost:8080/api/product?id=1234 - GET
* localhost:8080/api/product/ - POST

Curl

`curl -v localhost:8080/api/products`