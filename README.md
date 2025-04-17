# Sports Tracking Microservices Application

A microservices-based application for tracking sports sessions, progress, and fitness goals.

## Architecture Overview

The application is built using a microservices architecture with the following services:

### Core Services

1. **User Service**

   - Handles user authentication and profile management
   - Manages user preferences and settings
   - gRPC API for user operations
   - Written in Go

2. **Workout Service**

   - Manages workout sessions and exercises
   - Tracks workout history
   - Provides workout statistics and analytics

3. **Progress Service**

   - Tracks user's physical measurements
   - Manages fitness goals
   - Provides progress analytics

4. **Dashboard Service**

   - Aggregates data from other services
   - Generates visualizations and reports
   - Provides motivational insights

5. **Notification Service**
   - Handles user notifications
   - Sends reminders and motivational messages
   - Manages communication preferences

### Technology Stack

- **Backend**:
  - Go (for User Service)
  - gRPC for inter-service communication
- **Database**: PostgreSQL
- **Message Broker**: RabbitMQ (for asynchronous communication between services)
- **API Gateway**: Traefik (for routing and load balancing)
- **Containerization**: Docker
- **Orchestration**: Docker Compose (for development)
- **Frontend**: To be implemented

## Project Structure

```
microservices_interface/                    # Root directory of the project
├── README.md                              # Project documentation and setup instructions
├── microservices/                         # Directory containing all microservices
│   ├── user/                             # User management service
│   │   ├── cmd/                          # Main application entry points
│   │   ├── internal/                     # Internal package structure following clean architecture
│   │   │   ├── config/                   # Configuration management
│   │   │   ├── infrastructure/           # External interfaces (database, gRPC, etc.)
│   │   │   ├── domain/                   # Business logic and entities
│   │   │   └── application/              # Use cases and application services
│   │   ├── proto/                        # gRPC service definitions and protobuf files
│   │   ├── Dockerfile                    # Container configuration for the service
│   │   ├── docker-compose.yml            # Service-specific Docker Compose configuration
│   │   └── go.mod                        # Go module dependencies
│   ├── workout/                          # Workout tracking and management service
│   ├── progress/                         # Progress tracking and analytics service
│   ├── notification/                     # Notification and messaging service
│   └── dashboard/                        # Data aggregation and visualization service
└── interface/                            # Frontend application (to be implemented)
```

## Getting Started

1. Clone the repository
2. Install dependencies for each service
3. Set up the required environment variables
4. Start the services using Docker Compose

## Development

Each microservice is independently deployable and maintainable. They communicate through gRPC.

## API Documentation

API documentation for each service will be available through their respective gRPC definitions in the `proto` directory of each service.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
