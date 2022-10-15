## Project description

This project is an example of integration between client and service using the gRPC protocol.

The project consists of 2 microservices and each runs in a container in docker-compose:

1. Client - go service for calling second service using gRPC protocol.
2. Service - go service, that get request from client, starts some calculating script and returns response back to
   client.

## Start

To start the project, you must run the command in the terminal:

```bash
make dockerup
```

To stop, run the command:

```bash
make dockerstop
```

## Work

To check alive of gRPC service make GET request on http://localhost:8000/health.
If client and service are alive you get JSON:

```bash
{
    "status": "ok"
}
```

To get the main calculation send a POST request to http://localhost:8000/execute.
The example of request JSON you con find in directory:

```bash
client/testdata/input.json
```
