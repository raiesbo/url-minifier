# URL Minifier

API that allows to create smaller URLs for any project.

In many cases, some URLs that we want to share, or directly from our own project, are really lengthy and verbose. With this application, the user is able to retrieve a short version of a URL through a simple HTTP POST request.

It has two different interfaces. On one hand we find the Visual interface found under the root URL `/` which will render a web page with a form to introduce the URL to shorten. On the other hand, the application exposes a public JSON API, which allows integrating third party applications with the same goal.

The complete application is fully written in GO.

## Overview

### Features

- The user is able to obtain a shortened URL through a graphical interface
- The user is able to obtain a minified URL through a JSON API
- The user is able to check the original URL based on a given secret key.
- The user is able to remove the minified URL based on a given secreat

### Exposed API

| Endpoint            | HTTP Verb | Request Body    | Action |
| ------------------- | --------- | --------------- | ------ |
| /url                | POST      | Your target URL | Creates URL |
| /`{url_key}`        | GET       |                 | Forwards to target URL |
| /url/`{secret_key}` | GET       |                 | Retrieves URL information |
| /url/`{secret_key}` | DELETE    |                 | Removes URL |
| /admin              | POST      | Secret password | Retrives list of URLs |

### Getting Started

### Setup

1. Install Golang v1.21 or above
2. Create a `.env` file in the root of the project with the variables listed below
3. Run `go mod download` to install all the dependencies
4. Run `go build -o app` to compile the project
5. Run `./app` to run the server from the created binary file

### Environment Variables

- DB_URI: MongoDB connection URI connection
- PORT: Port in which the project will run

## Tech stack

- Built with Go version 1.21
- Uses the [Chi router](https://github.com/go-chi/chi/v5) package
- Uses the [godotenv](https://github.com/joho/godotenv) package
- Uses the [mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) package
