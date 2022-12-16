# Reverse Proxy Server
> This application base on simple reverse proxy server in Go.

## Table of Contents
* [General Info](#general-information)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Setup](#setup)
* [Usage](#usage)
* [Project Status](#project-status)
* [Acknowledgements](#acknowledgements)


## General Information
- This project was created because I wanted to try to implement simple algorithm with Golang.
- I wanted to learn more about the proxy servers. 


## Technologies Used
- Go - version 1.19.2


## Features
List the ready features here:
- create origin server
- create reverse proxy server
- forward a client request to the origin server via reverse proxy
- copy origin server response
- implementation in docker


## Setup
For start application with docker you need [Docker](https://docs.docker.com/get-docker/) and [docker-compose](https://docs.docker.com/compose/install/).


## Usage
The application can be build from sources or can be run in docker.

##### Build from sources
```bash
$ # Move to directory
$ cd folder/to/clone-into/
$
$ # Clone the sources
$ git clone https://github.com/mateuszgua/reverse_proxy_server.git
$
$ # Move into folder
$ cd reverse_proxy_server
$
$ # Start app
$ go run .
$ #2022/12/31 23:59:59 Starting server on http://localhost:8090 ...  
```

##### Start the app in Docker
```bash
$ # Move to directory
$ cd folder/to/clone-into/
$
$ # Clone the sources
$ git clone https://github.com/mateuszgua/reverse_proxy_server.git
$
$ # Move into folder
$ cd reverse_proxy_server
$
$ # Start app
$ docker-compose up --build
$ # ...
$ #app_1  | 2022/12/31 23:59:59 Starting server on http://localhost:8090 ...
```

## Project Status
Project is: in_progress 


## Acknowledgements
- This project was based on [this tutorial](https://dev.to/b0r/implement-reverse-proxy-in-gogolang-2cp4).