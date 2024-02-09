# URL Minifier
API that allows to create smaller URLs for any project.

## Description
In many cases, some URLs that we want to share, or directly from our own project, are really lengthy and verbous. With this applications the user is able to retrieve a extreainmy short version of a URL through a simple HTTP POST request.

The complete application is written in GO.

## Overview

### Features
- The user is able to obtain a minified URL
- The user is able to check the original URL based on a given secret key.
- The user is able to remove the minified URL based on a given secreat


### Exposed API

| Endpoint            | HTTP Verb | Request Body    | Action |
| ------------------- | --------- | --------------- | ------ |
| /                   | GET       |                 | Returns "Hello World!"  |
| /url                | POST      | Your target URL | Returns the created `url_key` with additional info such a the `secret_key` |
| /{`url_key`}        | GET       |                 | Redirects to the target URL |
| /url/{`secret_key`} | GET       |                 | Returns the original URL the shorten version is based of |
| /url/{`secret_key`} | DELETE    |                 | Removes the shortened URL |
| /admin              | POST      | Password        | Returns a list of all the minified URLs |
