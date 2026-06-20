# problem service

grpc service for problem crud, test case management, and code execution coordination. stores problems and submissions in mongodb, publishes execution requests to the engine via nats, and caches with redis.

## env vars

| var | default | description |
|-----|---------|-------------|
| environment | development | runtime environment |
| mongourl | mongodb://localhost:27017 | mongodb connection uri |
| problemserviceport | 50055 | grpc listen port |
| natsurl | nats://localhost:4222 | nats server url |
| redisurl | localhost:6379 | redis address |

## grpc services

- problem crud (create, read, update, delete) with metadata and test cases
- code execution (run user code against test cases, submit for scoring)
- leaderboard queries (global top k, entity-specific)
- submission history and problem statistics
