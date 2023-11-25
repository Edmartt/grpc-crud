# gRPC Test Server

gRPC is an implementation of RPC (remote procedure calls) originally designed by Google. It is open source and is used for communications with client-server architecture.

gRPC can use protocol buffers as an interface definition language and as a format for message exchange.

If you want to know more visit: [introduction to gRPC](https://grpc.io/docs/what-is-grpc/introduction/)

This small project is a simple CRUD in which gRPC is used to register people's data such as name, surname and email, containing the service code, in a separate package the gRPC client and also an HTTP client through which in a simple way and in JSON format the data can be sent to avoid complications configuring clients such as POSTMAN or Insomnia for gRPC.

## Requirements

- Go 1.19+
- SQLite
- http client: POSTMAN, Insomnia, cURL

### Running Locally

```
git clone https://github.com/Edmartt/grpc-crud.git
```

or ssh instead:

```
git clone git@github.com:Edmartt/grpc-crud.git
```

browse into project directory:

```
cd grpc-crud/
```

download dependencies

```
go mod tidy
```

set environment variables following the [.env.example](https://github.com/Edmartt/grpc-test-server/blob/main/.env.example) file and run

```
go run main.go
```

The command above will run the server-side code

Now you will need to run this next command for the http client:

```
go run pkg/client/http/main.go
```

#### Note

You can check api docs in: [api docs](http://localhost:8080/api/v1/swagger/index.html)
