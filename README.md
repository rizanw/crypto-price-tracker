# Crypto Price Tracker

# Overview

Backend restful API to track cryptocurrency in IDR using golang and sqlite. It has authentication using jwt token and
session manager. the session manager is simulated using golang hashmap variable, for large scale up in future we can
replace it using cache or persistent storage. Crypto rates source is from https://docs.coincap.io.

## Features

endpoint list:

- sign-up (email, password, password_confirm)
- sign-in (email, password)
- sign-out (bearer token)
- get user coins (bearer token)
- add coin to user (bearer token | coin)
- remove coin to user (bearer token | coin)

further endpoint description, kindly check postman collection.

# Postman Collection

Check the postman collection here:
https://www.postman.com/navigation-candidate-18708542/workspace/crypto-price-tracker/overview

# Local Development

## Prerequisites

Make sure you have installed all the following prerequisites on your development machine:

* go version : [1.19](https://golang.org/dl/)

## Local Run Guides:

To clone this repo:

```bash
git clone https://github.com/rizanw/crypto-price-tracker.git
```

To build and start the apps:

- build the binaries:

```bash 
make build
```

- start the app:

```bash 
make run
```

## Unit Test

To run unit test

```bash
make test
```

# Project Structure

- `bin/` is directory for compiled binary
- `cmd/` is the main program directory
- `files/` contains app files (including db & config)
    - `file/db` contains sqlite db directory
    - `file/etc/crypto-tracker` contains app config files
- `internal/` contains the whole logic of the app
    - `internal/common` contains helper functions
        - `internal/common/middleware` is for http middleware for client
        - `internal/common/session` is the session manager for auth
    - `internal/config` is the config of the app, has relation to files directory
    - `internal/handler` is application logic interface between this app with client
    - `internal/model` is model business design
    - `internal/repo` is the repositories to fetch/store data of this app
    - `internal/usecase` is main business logic
- `go.mod` the golang dependencies list