FROM golang:latest

WORKDIR /backend/

COPY . /backend/

CMD [ "go", "run", "main.go" ]