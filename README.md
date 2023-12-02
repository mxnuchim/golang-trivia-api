# Trivia REST API with Go, PostgreSQL, Fiber, and Docker

This repository contains a simple Trivia REST API built with Go, PostgreSQL, Fiber, and Docker.

## Setup

Follow the steps below to set up and run the Trivia API on your local machine.

### Prerequisites

Make sure you have the following tools installed on your machine:

- Docker
- Go (1.19.0 or later)

## Project Structure

- **cmd/api/main.go:** Main application entry point.
- **handlers:** Contains HTTP request handlers.
- **models:** Defines the data model for trivia facts.
- **routes.go:** Defines the API routes.

## Database

The API uses PostgreSQL as the database. You can access the database on `localhost:5432` with the provided credentials.

## API Endpoints

- **List Facts:** `GET /`
- **Create Fact:** `POST /fact`

## Environment Variables

Create a `.env` file in the root of the project with the following content:

```env
DB_USER=<db-user>
DB_PASSWORD=<db-password>
DB_NAME=<db-name>
```

These variables are used to configure the PostgreSQL database for the Trivia API. Adjust the values according to your requirements.

### Clone the Repository

```bash
    git clone https://github.com/mxnuchim/golang-trivia-api.git
    cd golang-trivia-api
```

### Running the API

To start the API, run the following commands:

```bash
make build   # Build the Docker image
make run     # Start the Docker containers and the api
```

## Makefile Commands

- **start_web:** Run a bash shell in the web service container with service ports exposed.

```bash
  make start_web
```

- **build:** Build the Docker images.

  ```bash
  make build
  ```

- **run:** Start the containers.

  ```bash
  make run
  ```
