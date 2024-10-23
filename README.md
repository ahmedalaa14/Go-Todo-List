# Go-Web-API

A simple TodoList application written in Go that allows users to view and manage their tasks through HTTP endpoints.

## Description

This project is a basic web server that provides endpoints to display a list of tasks and greet the user. It demonstrates basic HTTP handling and string manipulation in Go.

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/ahmedalaa14/Go-Web-API.git
    cd Go-Web-API
    ```

2. **Install Go**:
    Ensure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

3. **Run the application**:
    ```sh
    go run main.go
    ```

## Usage

Once the application is running, you can interact with it using a web browser or tools like `curl`.

- **Home Page**:
    ```sh
    curl http://localhost:8080/
    ```
    This will return a welcome message.

- **Show Tasks**:
    ```sh
    curl http://localhost:8080/show-tasks
    ```
    This will return the list of tasks.

## Endpoints

- `GET /`: Returns a welcome message.
- `GET /show-tasks`: Returns the list of tasks.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.
