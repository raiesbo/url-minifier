# URL Minifier

API that allows creating shorter URLs for any project.

In many cases, some URLs that we want to share, or directly from our own project, are excessively long and verbose. With this application, users can obtain a concise version of a URL through a straightforward HTTP POST request.

The application provides two distinct interfaces. On one hand, there is the Visual interface located at the root URL /, which renders a web page with a form for entering the URL to be shortened. On the other hand, the application exposes a public JSON API, enabling the integration of third-party applications with the same objective.

The entire application is written entirely in GO.

## Overview

### Features

- The user is able to obtain a shortened URL through a graphical interface
- The user is able to obtain a minified URL through a JSON API
- The user is able to check the original URL based on a given secret key.
- The user is able to remove the minified URL based on a given secreat

### Exposed API

| Endpoint                | HTTP Verb | Request Body    | Action |
| ----------------------- | --------- | --------------- | ------ |
| `/{url_key}`            | GET       |                 | Forwards to target URL |
| `/api/url`              | POST      | Your target URL | Creates URL |
| `/api/url/{secret_key}` | GET       |                 | Retrieves URL information |
| `/api/url/{secret_key}` | DELETE    |                 | Removes URL |
| `/api/admin`            | POST      | Secret password | Retrives list of URLs |

### Getting Started

### Setup

1. Install Go v1.21 or above
2. Create a `.env` file in the root of the project with the variables listed below
3. Run `go mod download` to install all the dependencies
4. Run `go build -o app` to compile the project
5. Run `./app` to run the server from the created binary file

### Environment Variables

- `DB_URI`: MongoDB connection URI connection
- `PORT`: Port in which the project will run

## Tech stack

- Built with [Go](https://go.dev/) version 1.21
- Uses the [Chi router](https://github.com/go-chi/chi/v5) package
- Uses the [godotenv](https://github.com/joho/godotenv) package
- Uses the [mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) package
- Uses the [htmx](https://htmx.org/) framework
