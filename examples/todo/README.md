# Todo List Application

## Description

A simple Todo List application that allows users to add, edit, mark as complete, and delete tasks. The application is developed using a Svelte as frontend and Go as backend to store the state.

## Installing

1. Clone the repository
2. Install Docker and Docker Compose

## Usage

To run the application locally, use `docker-compose -f infra/docker-compose.dev.yml up`.

## Tests

Backend unit tests can be run using `go test ./backend`. For frontend tests, use `npm test` inside the `frontend` directory.

## Deploying

Deploy the application using `docker-compose -f infra/docker-compose.prod.yml up -d`.

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request