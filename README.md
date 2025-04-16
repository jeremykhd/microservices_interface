# Sports Tracking Microservices Application

A microservices-based application for tracking sports sessions, progress, and fitness goals.

## Architecture Overview

The application is built using a microservices architecture with the following services:

### Core Services

1. **User Service**

   - Handles user authentication and profile management
   - Manages user preferences and settings
   - REST API endpoints for user operations

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

- **Backend**: Node.js with Express
- **Database**: MongoDB (for each service)
- **Message Broker**: RabbitMQ (for inter-service communication)
- **API Gateway**: Express Gateway
- **Frontend**: React with TypeScript
- **Containerization**: Docker
- **Orchestration**: Docker Compose (for development)

## Project Structure

```
microservices/
├── user-service/
├── workout-service/
├── progress-service/
├── dashboard-service/
└── notification-service/

interface/
└── frontend/
```

## Getting Started

1. Clone the repository
2. Install dependencies for each service
3. Set up the required environment variables
4. Start the services using Docker Compose

## Development

Each microservice is independently deployable and maintainable. They communicate through REST APIs and message queues.

## API Documentation

API documentation for each service will be available at:

- User Service: `/api/users/docs`
- Workout Service: `/api/workouts/docs`
- Progress Service: `/api/progress/docs`
- Dashboard Service: `/api/dashboard/docs`
- Notification Service: `/api/notifications/docs`
