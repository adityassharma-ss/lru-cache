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

- Additional Notes

Running Just the Backend Without Docker

If you want to run only the backend without Docker, follow these steps:

    Ensure you have Go installed on your machine.
    Navigate to the project root directory.
    Run the backend using the following command:

    ``` go run main.go ```

    The backend server will start on http://localhost:8080.
Project Structure

The project structure is as follows:

    client/: Contains the React frontend code.
    go.mod and go.sum: Go module files for backend dependencies.
    main.go: Main backend Go file containing the cache management logic.


- WORKING SCREENSHOTS

- Backend:

![Screenshot from 2024-06-28 16-13-11](https://github.com/adityassharma-ss/lru-cache/assets/82082352/1cb98f40-c949-402d-a3fc-5075d2a96d74)

- Frontend & Backend:
![Screenshot from 2024-06-28 16-12-35](https://g![Screenshot from 2024-06-28 16-12-17](https://github.com/adityassharma-ss/lru-cache/assets/82082352/3f105219-2a32-4911-8b04-03e6d76d47a4)
ithub.com/adityassharma-ss/lru-cache/assets/82082352/784283d5-d88b-4459-bbba-a7fcc362fd9c)




