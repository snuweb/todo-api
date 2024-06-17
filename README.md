# To-Do List API

Welcome to the To-Do List API, a simple RESTful API built using Go to manage to-do items. This project demonstrates basic CRUD operations and how to set up a RESTful service with Go.

## Features

- Retrieve all to-do items
- Retrieve a specific to-do item by ID
- Create a new to-do item
- Update an existing to-do item
- Delete a to-do item

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (latest version recommended)

## Installation

Follow these steps to get the project up and running:

1. **Clone the repository**:
    ```bash
    git clone https://github.com/snuweb/todo-api.git
    cd todo-api
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

## Running the API

To start the server, run:
```bash
go run main.go


List All To-Dos
Endpoint: /todos
Method: GET


[
  {
    "id": 1,
    "title": "Learn Go",
    "status": "pending"
  },
  {
    "id": 2,
    "title": "Build a REST API",
    "status": "completed"
  }
]

Get To-Do by ID
Endpoint: /todos/{id}
Method: GET
Response:
json

{
  "id": 1,
  "title": "Learn Go",
  "status": "p
ending"
}


Sure! Here's an example README.md file for your Go To-Do app project. This file includes sections that describe the project, how to set it up, how to run it, and how to use the API endpoints.

README.md
markdown
Copy code
# To-Do List API

Welcome to the To-Do List API, a simple RESTful API built using Go to manage to-do items. This project demonstrates basic CRUD operations and how to set up a RESTful service with Go.

## Features

- Retrieve all to-do items
- Retrieve a specific to-do item by ID
- Create a new to-do item
- Update an existing to-do item
- Delete a to-do item

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (latest version recommended)

## Installation

Follow these steps to get the project up and running:

1. **Clone the repository**:
    ```bash
    git clone https://github.com/snuweb/todo-api.git
    cd todo-api
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

## Running the API

To start the server, run:
```bash
go run main.go
By default, the server will run on http://localhost:8080.

API Endpoints
Here are the available endpoints:

List All To-Dos
Endpoint: /todos
Method: GET
Response:
json
Copy code
[
  {
    "id": 1,
    "title": "Learn Go",
    "status": "pending"
  },
  {
    "id": 2,
    "title": "Build a REST API",
    "status": "completed"
  }
]
Get To-Do by ID
Endpoint: /todos/{id}
Method: GET
Response:
json
Copy code
{
  "id": 1,
  "title": "Learn Go",
  "status": "pending"
}
Create a New To-Do
Endpoint: /todos
Method: POST
{
  "title": "New Task",
  "status": "pending"
}
Response
{
  "id": 3,
  "title": "New Task",
  "status": "pending"
}
Update a To-Do
Endpoint: /todos/{id}
Method: PUT
Request Body

{
  "title": "Updated Task",
  "status": "completed"
}
Response
{
  "id": 1,
  "title": "Updated Task",
  "status": "completed"
}
Delete a To-Do
Endpoint: /todos/{id}
Method: DELETE
Response:

[
  {
    "id": 2,
    "title": "Build a REST API",
    "status": "completed"
  }
]

Sure! Here's an example README.md file for your Go To-Do app project. This file includes sections that describe the project, how to set it up, how to run it, and how to use the API endpoints.

README.md
markdown
Copy code
# To-Do List API

Welcome to the To-Do List API, a simple RESTful API built using Go to manage to-do items. This project demonstrates basic CRUD operations and how to set up a RESTful service with Go.

## Features

- Retrieve all to-do items
- Retrieve a specific to-do item by ID
- Create a new to-do item
- Update an existing to-do item
- Delete a to-do item

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (latest version recommended)

## Installation

Follow these steps to get the project up and running:

1. **Clone the repository**:
    ```bash
    git clone https://github.com/snuweb/todo-api.git
    cd todo-api
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

## Running the API

To start the server, run:
```bash
go run main.go
By default, the server will run on http://localhost:8080.

API Endpoints
Here are the available endpoints:

List All To-Dos
Endpoint: /todos
Method: GET
Response:
json
Copy code
[
  {
    "id": 1,
    "title": "Learn Go",
    "status": "pending"
  },
  {
    "id": 2,
    "title": "Build a REST API",
    "status": "completed"
  }
]
Get To-Do by ID
Endpoint: /todos/{id}
Method: GET
Response:
json
Copy code
{
  "id": 1,
  "title": "Learn Go",
  "status": "pending"
}
Create a New To-Do
Endpoint: /todos
Method: POST
Request Body:
json
Copy code
{
  "title": "New Task",
  "status": "pending"
}
Response:
json
Copy code
{
  "id": 3,
  "title": "New Task",
  "status": "pending"
}
Update a To-Do
Endpoint: /todos/{id}
Method: PUT
Request Body:
json
Copy code
{
  "title": "Updated Task",
  "status": "completed"
}
Response:
json
Copy code
{
  "id": 1,
  "title": "Updated Task",
  "status": "completed"
}
Delete a To-Do
Endpoint: /todos/{id}
Method: DELETE
Response:
json
Copy code
[
  {
    "id": 2,
    "title": "Build a REST API",
    "status": "completed"
  }
]
Contributing
Contributions are welcome! If you have any suggestions or improvements, please feel free to open an issue or submit a pull request. For major changes, please discuss them first via an issue.

License
This project is licensed under the MIT License. See the LICENSE file for more details.
