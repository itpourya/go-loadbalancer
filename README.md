# Golang Load Balancer

This is a **Load Balancer** implemented in **Golang** that distributes incoming traffic across multiple backend servers. The load balancer supports different load-balancing strategies and ensures high availability and fault tolerance for backend services.

## Features

- **Least Connections**: Routes traffic to the server with the fewest active connections.
- **Health Checks**: Periodically checks the health of backend servers and removes unhealthy ones from the rotation.
- **High Availability**: Routes traffic only to available and healthy servers.
- **Logging**: Provides logs for monitoring traffic distribution and backend health.

## Prerequisites

To run this project, you'll need:

- [Golang](https://golang.org/dl/) (version 1.16 or above)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/itpourya/go-loadbalancer.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-loadbalancer
   ```

3. Install the dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

1. Open the `internal/example` file (if applicable) and configure your backend servers with python Flask:

## Running the Load Balancer

To start the load balancer:

```bash
go run cmd/main.go
```

The load balancer will start and listen for incoming traffic, forwarding requests to available backend servers according to the specified algorithm.

## Health Checks

The load balancer periodically checks the health of backend servers by sending HTTP requests. Unhealthy servers are automatically removed from the rotation until they recover.

## Project Structure

- `main.go`: The entry point of the load balancer.
- `internal/loadbalancer/`: Contains the logic for different load balancing strategies.
- `internal/health/`: Manages health checks for backend servers.

## Contribution

Feel free to contribute to this project by submitting pull requests:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.
