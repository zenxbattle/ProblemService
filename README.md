# xcode Problems Service

A high-performance, distributed coding problems platform built with Go that provides problem management, code execution, and leaderboard functionality. The service uses gRPC for communication, MongoDB for data persistence, Redis for caching and leaderboards, and NATS for asynchronous messaging.

## Architecture

The service follows a microservices architecture with the following components:

- **gRPC Server**: Handles problem-related operations and user submissions
- **Repository Layer**: Manages data persistence with MongoDB
- **Service Layer**: Implements business logic and orchestrates operations
- **Cache Layer**: Redis-based caching for performance optimization
- **Leaderboard System**: Real-time ranking using RedisBoard
- **Message Queue**: NATS for asynchronous communication
- **Logging**: Structured logging with Zap and BetterStack integration

## Tech Stack

- **Runtime**: Go 1.24
- **Database**: MongoDB for document storage
- **Cache**: Redis for high-performance caching and leaderboards
- **Messaging**: NATS for asynchronous communication
- **API**: gRPC with Protocol Buffers
- **Logging**: Zap logger with BetterStack integration
- **Containerization**: Docker with Docker Compose for development

## Key Features

### Problem Management
- Create, update, and manage coding problems
- Support for multiple programming languages
- Test case validation and management
- Problem difficulty categorization and tagging

### Submission System
- Code submission processing and validation
- Real-time execution results
- Submission history tracking
- Language-specific code validation

### Leaderboard System
- **High-Performance Rankings**: Leverages RedisBoard for ultra-fast rank pickups and queries
- **Real-time Updates**: Instant leaderboard updates with sub-millisecond response times
- **Global & Country Rankings**: Dual-tier ranking system with efficient entity-based filtering
- **Scalable Architecture**: Supports 1M+ users with 200+ entities (countries)
- **Periodic Sync**: Automated MongoDB to Redis synchronization via cron jobs

### Caching Layer
- Redis-based caching for frequently accessed data
- Performance optimization for problem retrieval
- Session and temporary data management

## Internal Workflow

### 1. Problem Submission Flow
- Client submits code via gRPC
- Service validates submission against test cases
- Results are processed and stored in MongoDB
- Leaderboard is updated in real-time
- Response is sent back to client

### 2. Leaderboard Management
- **RedisBoard Integration**: Leverages RedisBoard library for O(log n) rank operations
- **Fast Rank Pickups**: Sub-millisecond rank queries vs traditional database sorting
- **Dual Storage Strategy**: MongoDB for persistence, Redis for real-time rankings
- **Periodic Sync**: Hourly cron job ensures data consistency between stores
- **Incremental Updates**: Real-time score updates without full leaderboard recalculation

### 3. Data Persistence
- MongoDB stores problems, submissions, and user data
- Redis caches frequently accessed data
- Atomic operations ensure data consistency

## Configuration

Environment variables (`.env` file):

```bash
# Service Ports
APIGATEWAYPORT=7000
USERGRPCPORT=50051
PROBLEMSERVICE=50055

# Database Connections
MONGODBURL=mongodb://localhost:27017
REDISURL=localhost:6379
NATSURL=nats://localhost:4222

# Logging
ENVIRONMENT=development
BETTERSTACKSOURCETOKEN=<your_token>
BETTERSTACKUPLOADURL=<logging_endpoint>
```

## Development Setup

### Prerequisites
- Go 1.24+
- Docker and Docker Compose
- MongoDB
- Redis
- NATS Server

### Local Development

1. **Clone and setup**:
```bash
git clone <repository-url>
cd ProblemService
go mod download
```

2. **Start infrastructure services**:
```bash
docker-compose up -d
```

3. **Configure environment**:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. **Run the service**:
```bash
go run cmd/main.go
```

### Docker Deployment

```bash
# Build the application
docker build -t problem-service .

# Run with Docker Compose
docker-compose up --build
```

## API Endpoints

The service exposes gRPC endpoints for:

- **Problem Management**: Create, read, update, delete problems
- **Submission Processing**: Handle code submissions and validation
- **Leaderboard Queries**: Retrieve rankings and user statistics
- **Challenge Management**: Handle coding challenges and competitions

## Monitoring and Logging

- **Structured Logging**: Zap logger with configurable levels
- **External Logging**: BetterStack integration for log aggregation
- **Performance Monitoring**: Request tracing and execution metrics


## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.