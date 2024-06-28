LRU Cache Management Application

This project implements a Least Recently Used (LRU) Cache Management application using Go for the backend and React for the frontend.
Features

    Set Operation: Allows setting key-value pairs with an optional expiration time.
    Get Operation: Retrieves the value associated with a given key if it exists and is not expired.
    Delete Operation: Removes a key-value pair from the cache.

Technologies Used

    Backend: Go (Golang)
    Frontend: React
    Containerization: Docker, Docker Compose

Setup Instructions
Prerequisites

    Docker
    Node.js (for frontend development)

Steps to Run

    Clone the Repository:

    ```
    git clone <repository-url>
    cd lru-cache
    ```

    Build and Run with Docker Compose:

    ```docker-compose up --build
```

    This command will build and start the frontend and backend services defined in docker-compose.yml.

    Access the Application:

    Once the containers are up and running, you can access the frontend application in your browser at http://localhost:3000.

Project Structure

The project structure is as follows:

    client/: Contains the React frontend code.
    go.mod and go.sum: Go module files for backend dependencies.
    main.go: Main backend Go file containing the cache management logic.

Backend API Endpoints

    POST /set: Sets a key-value pair in the LRU cache.
    GET /get: Retrieves the value associated with a key.
    DELETE /delete: Deletes a key-value pair from the LRU cache.

Additional Notes

    Adjust CORS settings in main.go if deploying to a different environment.