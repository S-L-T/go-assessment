# Go Assessment

## Requirements

The application requires `docker` and `docker-compose` to run. 
`make` greatly eases the installation process, but is not strictly required.

## Commands

Run `make <command>` in order to execute one of the following commands:

| Command              | Description                                              |
|----------------------|----------------------------------------------------------|
| `start`              | Launches the application.                                |
| `start-no-cache`     | Launches the application and rebuilds the docker images. |
| `stop`               | Stops the application.                                   |
| `reset-db`           | Resets the database with some initial dummy data.        |
| `run-tests`          | Runs all the application tests.                          |

## Setup

1) Create a file `.env` on the project root and add the following contents:
```
JWT_KEY=Iamverysecret
DB_HOST=db
DB_PORT=3306
DB_USERNAME=demo
DB_PASSWORD=demo
```
2) `make start`
3) `make reset-db`

After the initial setup, you only need run `make start` to launch the application.
Adminer provides a simple UI to the database, allowing for easy visualization and manipulation
of the data.

| Application   | Port   |
|---------------|--------|
| HTTP Endpoint | `8080` |
| MySQL         | `3306` |
| Adminer       | `8888` |

## Issues & possible improvements

* Inserting a new company, although functional, does not return the inserted object's ID.
* Attempting to retrieve a company while providing a valid but non-existing UUID will erroneously return a 200 response with a default body
* The application does not shut down gracefully in case of SIGINT and SIGTERM signals.
* gRPC and Kafka events have not been implemented.
* The error handling/logging could be improved.

## Sample requests

You can use the following requests to test various aspects of the application.
cURL can be imported into Postman for even easier testing.
When performing GET, PATCH and DELETE requests, keep in mind that the ID must exist in your database.

Add the following header to your requests for PUT, PATCH and DELETE calls:
`Authorization: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRlbW8ifQ.kcTpIyaazrW3jrr3ydVwS1K0YM6IxgxJeQ1omMnV2uo"`

### Get company
```
curl --location --request GET 'localhost:8080/company' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "330542E99A5B11EDAE360242C0A83002"
}
'
```

### Delete company
```
curl --location --request DELETE 'localhost:8080/company' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRlbW8ifQ.kcTpIyaazrW3jrr3ydVwS1K0YM6IxgxJeQ1omMnV2uo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "330542E99A5B11EDAE360242C0A83002"
}
'
```

### Create company
```
curl --location --request PUT 'localhost:8080/company' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRlbW8ifQ.kcTpIyaazrW3jrr3ydVwS1K0YM6IxgxJeQ1omMnV2uo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Company8",
    "description": "Company8 description",
    "total_employees": 10,
    "is_registered": true,
    "type": 4
}'
```

### Update company
```
curl --location --request PATCH 'localhost:8080/company' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRlbW8ifQ.kcTpIyaazrW3jrr3ydVwS1K0YM6IxgxJeQ1omMnV2uo' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "6056190C9A9A11ED9C9B0242AC1D0002",
    "name": "Company9",
    "description": "Company9 description",
    "total_employees": 10,
    "is_registered": false,
    "type": 3
}'
```
