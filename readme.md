# SynaPedia

SynaPedia is a simple e-commerce API implementation.

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Technologies](#technologies)
- [Architecture](#architecture)
- [Authentication](#authentication)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Description

SynaPedia is a basic e-commerce API designed to demonstrate the implementation of an e-commerce system. It is built using Golang and PostgreSQL, following the principles of Clean Architecture. The project also uses JWT tokens for secure authentication.

## Features

- User registration and authentication
- Product listing and details
- Shopping cart management
- Order processing

## Technologies

- **Programming Language:** Golang
- **Database:** PostgreSQL
- **Architecture:** Clean Architecture
- **Authentication:** JWT tokens

## Architecture

The project follows the Clean Architecture principles, which emphasize separation of concerns and maintainability. The main layers include:

- **Handlers:**  User-interface level implementation of the Request-Response model
- **Use Cases:** Application-specific business logic and rules
- **Repositories:** Data layer to communicate with database

## Authentication

Authentication is handled using JWT tokens. This ensures secure and stateless user sessions.

## Installation

1. **Clone the repository:**
    ```bash
    git clone https://github.com/YeremiahKevin/synapedia.git
    cd synapedia
    ```

2. **Install dependencies:**
   Ensure you have Golang and PostgreSQL installed on your machine. Install the necessary Go packages:
    ```bash
    go get ./...
    ```

3. **Set up PostgreSQL:**
   Create a PostgreSQL database and update the connection string in the configuration file.

4. **Set up .env:**
   Create new .env file that follows this template (please adjust the value based on your needs)
   ```code
   DATABASE_HOST=
   DATABASE_PORT=
   DATABASE_USERNAME=
   DATABASE_PASSWORD=
   DATABASE_NAME=
   JWT_KEY=
   ```

## Usage

1. **Run the application:**
    ```bash
    go run main.go
    ```

2. **Server:**
   The HTTP server will listen on port `8080`.

3. **API Endpoints:**
   Use tools like Postman or cURL to interact with the API. Detailed API documentation will be provided in a separate file (e.g., `API_DOCS.md`).

