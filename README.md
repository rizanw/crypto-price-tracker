# Crypto Price Tracker

# Overview

Backend restful API to track cryptocurrency in IDR using golang and sqlite. It has authentification using jwt token and
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

further endpoint description, kindly check postman colection.

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

- Install golang dependency:

```bash 
go mod tidy
```

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