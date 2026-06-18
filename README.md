# Problem Service

gRPC service for problem CRUD, test cases, code execution via NATS, and submission tracking.

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `ENVIRONMENT` | No | `development` | Runtime environment |
| `MONGODBURL` | Yes | `mongodb://localhost:27017` | MongoDB connection URI |
| `PROBLEMSERVICE` | No | `50055` | gRPC listen port |
| `NATSURL` | No | `nats://localhost:4222` | NATS server URL |
| `REDISURL` | No | `localhost:6379` | Redis address |
