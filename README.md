# Timecard Service

[![Go Report Card](https://goreportcard.com/badge/github.com/patricksferraz/timecard-service)](https://goreportcard.com/report/github.com/patricksferraz/timecard-service)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/patricksferraz/timecard-service?status.svg)](https://godoc.org/github.com/patricksferraz/timecard-service)

A modern, scalable timecard management service built with Go, featuring gRPC and REST APIs, event-driven architecture, and comprehensive monitoring.

## 🚀 Features

- **Dual API Support**: REST and gRPC endpoints for maximum flexibility
- **Event-Driven Architecture**: Powered by Apache Kafka for reliable event processing
- **Database Support**: PostgreSQL integration with GORM for robust data persistence
- **API Documentation**: Swagger/OpenAPI documentation for easy API exploration
- **Monitoring**: Elastic APM integration for comprehensive application monitoring
- **Containerized**: Docker and Docker Compose support for easy deployment
- **Development Tools**: PGAdmin included for database management
- **Kafka Management**: Confluent Control Center for Kafka monitoring and management

## 🛠️ Tech Stack

- **Language**: Go 1.16+
- **Framework**: Gin (REST), gRPC
- **Database**: PostgreSQL
- **ORM**: GORM
- **Message Broker**: Apache Kafka
- **Monitoring**: Elastic APM
- **Container**: Docker
- **Documentation**: Swagger/OpenAPI

## 📋 Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Make (for using Makefile commands)

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/patricksferraz/timecard-service.git
   cd timecard-service
   ```

2. Copy the environment file and configure it:
   ```bash
   cp .env.example .env
   ```

3. Start the services using Docker Compose:
   ```bash
   make up
   ```

4. The service will be available at:
   - REST API: http://localhost:8080
   - gRPC: localhost:50051
   - Swagger UI: http://localhost:8080/swagger/index.html
   - PGAdmin: http://localhost:9000
   - Kafka Control Center: http://localhost:9021

## 🏗️ Project Structure

```
.
├── application/     # Application layer (use cases)
├── cmd/            # Command line interface
├── domain/         # Domain models and interfaces
├── infrastructure/ # Infrastructure implementations
├── proto/          # Protocol buffer definitions
└── utils/          # Utility functions and helpers
```

## 🔧 Development

### Available Make Commands

```bash
# Docker Compose Operations
make up          # Start services in detached mode
make down        # Stop and remove containers, networks, and volumes
make start       # Start existing services
make stop        # Stop running services
make ps          # List running services
make logs        # View service logs
make attach      # Attach to a service container (requires SERVICE=service_name)
make build       # Build service images
make prune       # Remove unused Docker resources

# Testing
make test        # Run tests using docker-compose.test.yml
make gtest       # Run Go tests with coverage report

# Code Generation
make gen         # Generate Go code from Protocol Buffer definitions
```

### Running Tests

You can run tests in two ways:

1. Using Docker Compose (recommended for CI/CD):
   ```bash
   make test
   ```

2. Running Go tests directly with coverage:
   ```bash
   make gtest
   ```

## 📚 API Documentation

Once the service is running, you can access the Swagger documentation at:
http://localhost:8080/swagger/index.html

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- Patrick Ferraz - Initial work

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Confluent Kafka](https://www.confluent.io/)
- [Elastic APM](https://www.elastic.co/apm)
