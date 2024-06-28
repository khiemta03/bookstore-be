# bookstore-be

This project provides APIs for online book stores by applying [Microservices Architecture](https://microservices.io/). There are five primary services included:
  - **User Service**: used to manage information related to users.
  - **Authentication Service**: authenticates users before using other features.
  - **Book Service**: used to manage information about books, authors, and publishers.
  - **Order Service**: manages and handles order requests.
  - **API Getway**: the main broker between other services.


## Tech Stack

**Server:** Gin (Golang), gRPC, SQLC, REST API, Microservices Architecture, JSON Web Token

**Testing**: Go Mock, Postman

**Documentation**: Postman

**Deployment**: Docker

**Database:** PostgreSQL

## API Examples

#### Get Specified User

```
  GET {BASE_URL}/users/:id
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required**. Id of user to retrieve |

#### Get Specified Book

```
  GET {BASE_URL}/books/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of book to fetch |

#### Add New Item Into The Shopping Cart

```
  POST {BASE_URL}/items
```

| Body Field | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `book_id`      | `string` | **Required**. Id of book to add |
| `quantity`      | `int` | **Required**. Quantity of book to add |

And more ...

## Documentation

Please follow this file for more API details: [docs](https://github.com/khiemta03/bookstore-be/blob/main/API-Getway/public/docs/postman.json). I suggest you download and open it in [Postman](https://www.postman.com/).


## Installation

### Clone the Repository
Firstly, clone this repository to your local machine:

```bash
  git clone https://github.com/khiemta03/bookstore-be.git
  cd bookstore-be
```

### Run with Docker (Recommended)
To run all projects using Docker, ensure Docker is installed on your machine. If it's not installed yet, you can download it from [Docker's official website](https://www.docker.com/):

#### Start Docker Compose
```bash
  docker compose up
``` 
This command will build and start all necessary services defined in your docker-compose.yml file.

### Run Manually Project by Project
You can run all projects without Docker by following these steps for each project:

#### 1. Setup PostgreSQL Database:

```bash
  make postgres
  make createdb
``` 

#### 2. Database Migration:

```bash
  make migration
``` 

#### 3. Run the Server:

```bash
  make server
``` 

### Access the Main Server
After starting the projects, the main server will be accessible at port 3000.

## Authors
- [@khiemta03](https://www.github.com/khiemta03)



