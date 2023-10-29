# ReadSpace Backend

This project provides backend support for the ReadSpace application, supporting functionalities like user authentication, managing reading sessions, and more. It is built with Go, and uses Docker for containerization and Postgres for data persistence.

## Setup Instructions

### Prerequisites:

1. Ensure you have **Docker** and **Docker Compose** installed on your machine. If not, you can get Docker [here](https://docs.docker.com/get-docker/) and Docker Compose [here](https://docs.docker.com/compose/install/).
2. Clone the repository to your local machine.

### Configuration:

1. **Environment Variables**:
   - At the root of the `backend` directory, create a `.env` file.
   - Use the provided `.env.example` as a template to fill in the required environment variables.

### Steps to Run:

1. **Build the Docker containers**:
   Navigate to the directory containing the `docker-compose.yml` file and run:

   ```bash
   docker-compose build
   ```

2. **Run the Docker containers**:

   ```bash
   docker-compose up
   ```

3. After executing the above steps, the backend service should be up and running, and you can access it at `http://localhost:8080`.

### Development:

- If you make changes to the Go code or environment variables, you will need to stop the containers, rebuild, and restart them. Use the following commands:

  ```bash
  docker-compose down
  docker-compose build
  docker-compose up
  ```

### Features

1. **User Authentication & Management**

   - Register a new user
   - Authenticate and log in a user
   - Log out a user
   - Delete a user account
   - Update user information

2. **User Reading Sessions**

   - Create a reading session
   - Save a reading session
   - Edit a reading session
   - Delete a reading session

3. **Unit Testing**
   - Comprehensive unit tests covering all functionalities to ensure robustness and correctness.

### Folder Structure

- `cmd`: Contains the application's entry points.

  - `api`: API initialization and routes.
  - `main.go`: Starts up the server.

- `internal`: Application's internal logic.

  - `entity`: Core domain objects or structs.
  - `handler`: Interfaces directly with the HTTP layer.
    - `rest`: Defines RESTful endpoints.
  - `repository`: Abstractions over data storage.
    - `dbrepo/postgres`: Implementation for PostgreSQL database.
  - `useCase`: Contains business logic.

- `pkg`: Libraries or code intended to be used by other services or applications.

- `testdata`: Contains static files or data for testing.

### API Endpoints

**Authentication**

- `POST /register`: Register a new user.
- `POST /login`: Authenticate a user.
- `POST /logout`: Log out an authenticated user.
- `DELETE /user`: Delete the user account.
- `PUT /user`: Update user information.

**Reading Sessions**

- `POST /session`: Create a new reading session.
- `GET /session/{sessionID}`: Retrieve details of a specific reading session.
- `PUT /session/{sessionID}`: Edit a specific reading session.
- `DELETE /session/{sessionID}`: Delete a specific reading session.

### TODO

- Book crud, get books,
- Create a book shelf for the user
- Reading session
- History of user reading session and focus

### Testing

The project's test coverage aims to be comprehensive, covering all application functionalities:

1. **Running Tests**:

   ```bash
   go test ./...
   ```

2. Ensure mocks are set up for dependencies to test in isolation.

3. Regularly check test coverage to ensure critical paths are tested.

### Contributions

Feel free to fork this repository, submit PRs or raise issues if you find any.

---
