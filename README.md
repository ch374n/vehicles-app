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

# Running on Minikube

## Prerequisites

- Minikube installed and running
- kubectl command-line tool installed

## Steps

1. **Start Minikube**

   If you haven't started Minikube yet, run the following command:
```
minikube start
```

2. **Create a Secret for MongoDB URI**

Create a Kubernetes Secret to store the MongoDB URI. Replace `<your-mongo-uri>` with your actual MongoDB connection string.

```bash
kubectl create secret generic app-secret --from-literal=APP_MONGO_URI='<your-mongo-uri>'
```

Alternatively, you can edit the Secret.yml file and replace <your-mongo-uri> with your MongoDB connection string:

```
# Secret.yml
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
type: Opaque
data:
  APP_MONGO_URI: <base64-encoded-mongo-uri>
```

Then, apply the Secret:

```
kubectl apply -f Secret.yml
```

3. **Set the Container Image in Deployment.yml**
Open the Deployment.yml file and replace <your-image> with the actual container image name and tag for your application:

```
# Deployment.yml
...
spec:
  containers:
  - name: your-app
    image: <your-image>
    ...
```

4. **Apply the Kubernetes Manifests**
Apply the Kubernetes manifests to create the Deployment, Service, and any other resources:
