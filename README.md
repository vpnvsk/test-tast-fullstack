# test-tast-fullstack

## Overview

This project is a full-stack test task that demonstrates the integration of Golang and React for building a simple User CRUD application. The backend is developed in Golang, utilizing the Gin web framework, while the frontend is built with React. The communication between the two is facilitated through RESTful endpoints.


## Docker Compose

A Docker Compose configuration is included in the project to facilitate easy deployment and management of the Golang backend and React frontend components in a single environment.

## Swagger Documentation

Swagger documentation is available at `http://localhost:8000/docs/index.html`. This documentation provides insights into the API endpoints, making it easier for developers to understand and interact with the backend. The backend is accessible at `localhost:8000`, while the frontend is available at `localhost:3000`.

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/vpnvsk/test-tast-fullstack.git
```

2. Start Docker Compose

 ```bash
 docker-compose up --build
 ```

## Tests

Both frontend and backend is covered with tests
