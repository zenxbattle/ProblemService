# ZenXBattle Problem Service

gRPC service for coding problem management. Handles problem CRUD, test cases, search/filter, and execution statistics.

## Architecture

```
ApiGateway → gRPC (:50053) → ProblemService
                                 ├── MongoDB (problems, test cases)
                                 ├── Redis (cache)
                                 └── NATS → CodeExecutionEngine (validation)
```

## Tech Stack

- **Go** + gRPC
- **MongoDB** for problem storage
- **Redis** for caching
- **NATS** for async code validation dispatch
- **RedisBoard** for problem-level leaderboards

## gRPC Endpoints

| Method | Description |
|--------|-------------|
| `CreateProblem` | Add new problem with test cases |
| `GetProblem` | Fetch problem by ID |
| `ListProblems` | Paginated problem list with filters |
| `UpdateProblem` | Modify problem/tests |
| `DeleteProblem` | Remove problem |
| `SearchProblems` | Full-text search |
| `GetProblemStats` | Submission/acceptance stats |
| `ValidateSolution` | Async: run solution against test cases |

## Quick Start

```bash
export MONGO_URI=mongodb://localhost:27017
export REDIS_ADDR=localhost:6379
export NATS_URL=nats://localhost:4222

go run cmd/main.go
# → gRPC server on :50053
```

## Data Model

```json
{
  "id": "uuid",
  "title": "Two Sum",
  "description": "...",
  "difficulty": "easy|medium|hard",
  "tags": ["array", "hashmap"],
  "testCases": [
    {"input": "[2,7,11,15]\n9", "expected": "[0,1]"}
  ],
  "timeLimit": 2.0,
  "memoryLimit": 256
}
```

## Docker

```bash
docker build -t zenxbattle-problems .
docker run -p 50053:50053 zenxbattle-problems
```

## Related Services

- [ApiGateway](https://github.com/zenxbattle/ApiGateway) — REST proxy
- [CodeExecutionEngine](https://github.com/zenxbattle/CodeExecutionEngine) — code validation
- [RedisBoard](https://github.com/zenxbattle/RedisBoard) — problem leaderboards
- [Frontend](https://github.com/zenxbattle/Frontend) — problem browser/editor
- [CommonProto](https://github.com/zenxbattle/CommonProto) — protobuf definitions
