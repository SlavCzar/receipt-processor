
# Receipt Processor

This project has been implemented using the golang clean architecture comprising of model, service, and repository layers and multiple tests.

## Building the Project

To build this project using Docker, run:

```
docker build -t receipt-processor .
```

## Running the Docker Image

To run the docker image:

```
docker run -p 8080:8080 receipt-processor
```

This will start the webserver, and you can then test the APIs using Postman or similar tools on port 8080.

## API Endpoints

The following API endpoints are available:

- `POST http://localhost:8080/receipts/process`
- `GET http://localhost:8080/receipts/{id}/points`

Feel free to use these endpoints to interact with the Receipt Processor application.


## Running Tests
```
go test -count=1 ./...
```
