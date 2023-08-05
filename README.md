## ReadSpace Backend

ReadSpace is a backend service built with Go using the Gin framework. The project adheres to the principles of clean architecture, ensuring a clear separation of concerns, scalability, and maintainability.

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

### Setting Up

1. **Database Setup**: Ensure you have PostgreSQL installed and running. Create a database and note down the credentials.

2. **Environment Variables**: Set up required environment variables like `DATABASE_URL`, `JWT_SECRET_KEY`, etc.

3. **Running the Application**:
   ```bash
   go run cmd/api/main.go
   ```

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

This README provides a holistic view of your project while emphasizing the clean architecture and importance of unit testing. As you progress in development, you might want to add more details or refine certain sections based on the specifics of your application.