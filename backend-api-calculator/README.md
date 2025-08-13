# Backend API - Calculator

The goal of this project is to create an http+json API for a calculator service.

## Overview

## Requirements

The API should conform to the given OpenAPI spec found in this directory, which can also be viewed at this URL.


#### Logging

In order to be able to debug issues that occur, you're going to want to log out each request as it comes in, as well as any associated data such as the status code, ip address, and what the request path was.

For logging in Go, I recommend using the `log/slog` package, which provides structured logging in either JSON, or Text format. You can create a logger using the following snippet:

##### Text
```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
```

##### JSON
```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
```
## Recommended Packages

Some of the packages that I used for my implementation include:

- `net/http` - This is the http package of the standard library which I use for routing and setting up an http server. If you want to know how to use this package for advanced routing, [I have a video you can check out](https://youtu.be/H7tbjKFSg58).
- `encoding/json` - This package is used for encoding and decoding JSON from the request and to the response bodies.
- `log/slog` - For structured logging
- `github.com/rs/cors` - For cors, if you need it.


## Additional Tasks

- Add in token authentication to prevent anyone unauthorized from using the API
- Add in a database to keep track of all of the calculations that have taken place
- Add in support for floating point numbers as well.
- Add in a middleware that adds a request ID to the http.Request object.