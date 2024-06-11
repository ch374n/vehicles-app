### Project Overview

This project implements a RESTful API for Vehicles-App in Golang, leveraging the gorilla-mux tool and mongodb for persistance. It offers the following features:

- **Create New Manufacturer**
- **Delete Manufacturer By Id**
- **Get All Manufacturers**
- **Update Manufacturer**
#### Prerequisites

- Golang (version 1.22 or later): https://go.dev/doc/install
- Docker (optional, but recommended for easy deployment): https://www.docker.com/
- Make (version 3.81 or later) : https://formulae.brew.sh/formula/make
####  Installation
1. Clone the repository:

```bash
git clone https://github.com/ch374n/vehicles-app.git
cd vehicles-app
```
2. Install dependencies:
```bash
go mod download
```
3. Create a .env file in the project root directory (copy from .env.example if provided) and define environment variables:
```bash
APP_ENV=dev  # Environment type 
APP_PORT=8080     # Port on which the server will listen
APP_MONGO_URI=mongsrv+ # Mongodb connection URL
```
4. Execute `make run` command

#### Explaination of Makefile targets:
- **help**: Prints all available options
- **clean**: Removes generated code
- **run**: Starts the server
- **test**: Runs test cases

#### API Usage with Swagger
Swagger documentation is available at http://localhost:8080/api-docs/. 
Refer to the generated API definitions for details on endpoints and request/response formats.

#### Building and Running the Docker Image
1. Build the Docker image
```bash
docker build -t your-vehicles-app-service . 
```

2. Run the container
```bash
docker run -p $PORT:$PORT your-vehicles-app-service
```
#### API Endpoints

| Endpoint           | HTTP Method | Description                                                               |
|--------------------|-------------|---------------------------------------------------------------------------|
| /api/v1/manufacturers    | GET        | Get all Manufacturers.          |
| /api/v1/manufacturers    | POST        | Create a new Manufacturer.          |
| /api/v1/manufacturers/{id} | GET         | Get Manufacturer by ID. |
| /api/v1/manufacturers/{id}    | PUT         | Update manufacturer by ID.        |
| /api/v1/manufacturers/{id}    | DELETE         | Delete manufacturer by ID.        |

