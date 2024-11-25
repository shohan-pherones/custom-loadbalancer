# Custom Loadbalancer

## Overview

This project is a simple **load balancer** implementation using Go and Docker. The load balancer distributes incoming requests to multiple backend servers in a round-robin fashion. It helps to balance the load among different servers to ensure high availability and fault tolerance. The backend servers are simple HTTP servers that simulate different service nodes. The project utilizes Docker to containerize all components and Docker Compose to orchestrate the services.

## Features

- **Load Balancing:** Distributes requests among multiple backend servers.
- **Health Checks:** Ensures only active servers handle requests.
- **Dockerized:** All components are containerized using Docker.
- **Scalable:** Easily scale the number of backend servers by adding more in the Docker Compose file.
- **Round-robin:** A simple round-robin algorithm for load distribution.

## Components

- **Load Balancer:** Written in Go, the load balancer listens on port `8080` and proxies incoming requests to the backend servers.
- **Backend Servers:** Three backend servers (`server1`, `server2`, and `server3`) each listening on ports `8081`, `8082`, and `8083`, respectively. Each backend server returns a response indicating the server that handled the request.

## Prerequisites

- Docker (version 20.10.0 or later)
- Docker Compose (version 1.27.0 or later)

## Getting Started

Follow the steps below to build and run the load balancer and its backend servers locally using Docker.

### 1. Clone the repository

```bash
git clone https://github.com/shohan-pherones/custom-loadbalancer.git
cd custom-loadbalancer
```

### 2. Build and start the services with Docker Compose

```bash
docker-compose up --build
```

This will:

- Build the Docker images for the load balancer and backend servers.
- Start the containers and connect them according to the `docker-compose.yml` configuration.
- Expose the load balancer on port `8080`, server1 on `8081`, server2 on `8082`, and server3 on `8083`.

### 3. Testing the Load Balancer

Once the services are up and running, you can test the load balancer by visiting `http://localhost:8080/` in your browser. The load balancer will forward the request to one of the active backend servers.

To test load balancing:

- Refresh the page multiple times, and you should see the requests being handled by different servers (Server 1, Server 2, or Server 3).

### 4. Shutting Down the Services

```bash
docker-compose down
```

This will stop and remove all containers defined in the `docker-compose.yml` file.

## How It Works

- **Load Balancer Logic (Go):** The load balancer runs an HTTP server on port `8080`. It maintains a list of active backend servers and distributes incoming requests to them in a round-robin fashion. If a backend server is inactive or unreachable, the load balancer skips it and tries another one.

- **Backend Servers:** Each backend server listens on a specific port (`8081`, `8082`, or `8083`) and responds with a message indicating which server handled the request. This is useful for testing and debugging the load balancer.

- **Docker Compose:** The project uses Docker Compose to simplify the process of managing multiple Docker containers. It defines the services (load balancer and backend servers) and their relationships, making it easy to scale and manage the infrastructure.

## Conclusion

This Load Balancer project demonstrates a simple yet effective way to distribute incoming HTTP requests across multiple backend servers, ensuring high availability and better resource utilization. By leveraging Docker and Docker Compose, the project enables easy containerization, management, and scaling of services. The load balancer, written in Go, efficiently distributes requests in a round-robin manner, ensuring an even distribution of traffic among the available backend servers.

Key takeaways:

- **Scalable Architecture:** The use of Docker Compose makes it easy to scale the system by adding more backend servers.
- **Fault Tolerance:** The load balancer ensures that only active servers handle requests, improving system resilience.
- **Dockerized Services:** Docker containers provide a consistent and isolated environment for each service, simplifying deployment and maintenance.
- **Versioning and Deployment:** Docker Hub integration allows you to tag and push your images with version control, making it easier to manage different versions of the application.

This project serves as a foundational framework for building more complex and robust load balancing solutions, which can be extended to include features such as health checks, logging, and more advanced load balancing algorithms.
