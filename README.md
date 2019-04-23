[![Build Status](https://travis-ci.org/777777miSSU7777777/github-aggregator.svg?branch=master)](https://travis-ci.org/777777miSSU7777777/github-aggregator)
[![](https://images.microbadger.com/badges/version/777777missu7777777/github-aggregator-web.svg)](https://hub.docker.com/r/777777missu7777777/github-aggregator-web/)
[![](https://images.microbadger.com/badges/image/777777missu7777777/github-aggregator-web.svg)](https://microbadger.com/images/777777missu7777777/github-aggregator-web)
[![codecov](https://codecov.io/gh/777777miSSU7777777/github-aggregator/branch/master/graph/badge.svg)](https://codecov.io/gh/777777miSSU7777777/github-aggregator) 
[![Go Report Card](https://goreportcard.com/badge/github.com/777777miSSU7777777/github-aggregator)](https://goreportcard.com/report/github.com/777777miSSU7777777/github-aggregator) 
[![GoDoc](https://godoc.org/github.com/777777miSSU7777777/github-aggregator?status.svg)](https://godoc.org/github.com/777777miSSU7777777/github-aggregator) 

# Github Aggregator
Aggregator is an application that represents GitHub Pull Requests and Activity stream in a more configurable and detailed way.
It helps organize code review process by monitoring events from organizations and projects, structuring and aggregating them in the fashion that fits a user best.

# Installation

## Default
First of all you need to install:

 - [git](https://git-scm.com/)
 - [golang](https://golang.org/)
 
Then you need to make following steps:
 - Make directory src/github.com/777777miSSU7777777 in $GOPATH:

       mkdir -p src/github.com/777777miSSU7777777 
	   && cd src/github.com/777777miSSU7777777

 - Clone git repository: 
 
	   git clone https://github.com/777777miSSU7777777/github-aggregator.git

- Build application:

      go build cmd/web-app/main.go

## Docker
You need to install [docker](https://www.docker.com/).
Then pull docker [image from dockerhub](https://hub.docker.com/r/777777missu7777777/github-aggregator-web/):

    docker pull 777777missu7777777/github-aggregator-web

# Usage
## Default
Run compiled binary:

    ./main
## Docker
Run pulled image in container:

    docker run -p 8080:8080 777777missu7777777/github-aggregator-web

## Common
Then go to 127.0.0.1:8080 in browser to use application.  

**Notice: you need browser with ES6 support.**

# Supported cmd flags
- **-host -h** - defines host ip (default is 0.0.0.0).
-  **-port -p** - defines host port  (default is 8080).
-  **-data-source** - defines data source (default is "rest-api").
