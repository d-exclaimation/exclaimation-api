# Getting Golang v1.16.3 Coded Buster
FROM golang:1.16.3-buster

# Using /app directory and moving things there, to avoid using the current working directory
RUN mkdir /app
ADD . /app
WORKDIR /app

# Enviroment setup
ENV GO111MODULE=on
ENV ENV_MODE=PRODUCTION

# Dependecies
RUN go mod download

# Build
RUN go build -o main .

# Start command
CMD ["/app/main"]
