# Task Manager

A RESTful task management microservice built in Go for the GRAPH technical assessment.

## Tech Stack

- **Go 1.25** — language
- **Gin** — HTTP framework
- **PostgreSQL** (pgx/v5) — primary database
- **Redis** (go-redis/v9) — cache layer
- **Prometheus** — metrics
- **OpenTelemetry** — distributed tracing
- **Swagger/OpenAPI** — API documentation
- **Docker / Docker Compose** — containerization
- **sql-migrate** — database migrations

## Features

- CRUD operations for tasks (Create, Read, Update, Delete)
- Task status: `todo` | `in_progress` | `done`
- Pagination and filtering by status and assignee
- Cache-aside pattern with Redis (invalidation on update/delete)
- Prometheus metrics: request count, latency histogram, tasks gauge
- OpenTelemetry tracing with otelgin middleware
- Swagger UI at `/swagger/index.html`
- `/metrics` endpoint for Prometheus scraping
- pprof profiling endpoint on port `6060`
- Health check endpoint
- 75%+ unit and integration test coverage

## Project Structure

```
graph/
├── delivery/httpserver/       # HTTP layer (Gin server, handlers, middleware)
│   ├── middleware/metrics.go  # Prometheus middleware
│   ├── taskhandler/           # CRUD route handlers
│   └── tracing.go             # OpenTelemetry tracer init
├── entity/                    # Domain models (Task, TaskStatus)
├── metrics/                   # Prometheus metric definitions
├── param/                     # Request/response DTOs
├── repository/
│   ├── psql/psqltask/         # PostgreSQL repository
│   └── redis/redistask/       # Redis cache adapter
├── service/                   # Business logic layer
├── migrations/                # SQL migration files (sql-migrate)
├── docs/                      # Swagger generated files
├── Dockerfile
├── docker-compose.yml
└── main.go
```

## Getting Started

### Prerequisites

- Go 1.25+
- Docker and Docker Compose

### Run with Docker Compose

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`.

### Run locally

Start dependencies:

```bash
docker-compose up postgres redis -d
```

Run database migrations:

```bash
sql-migrate up
```

Start the server:

```bash
go run main.go
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health-check` | Health check |
| GET | `/task/` | List tasks (paginated) |
| GET | `/task/:id` | Get task by ID |
| POST | `/task/` | Create a task |
| PUT | `/task/:id` | Update a task |
| DELETE | `/task/:id` | Delete a task |
| GET | `/metrics` | Prometheus metrics |
| GET | `/swagger/index.html` | Swagger UI |

### Query Parameters for `GET /task/`

| Param | Type | Default | Description |
|-------|------|---------|-------------|
| `page` | int | 1 | Page number |
| `page_size` | int | 10 | Items per page (max 100) |
| `status` | string | — | Filter by status |
| `assignee` | string | — | Filter by assignee |

### Example: Create a task

```bash
curl -X POST http://localhost:8080/task/ \
  -H "Content-Type: application/json" \
  -d '{"title":"My Task","description":"Details here","assignee":"alice"}'
```

## Running Tests

```bash
go test ./...
```

With coverage:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Load Test Results

Tested with [`hey`](https://github.com/rakyll/hey) at 50 concurrent connections:

| Endpoint | Req/sec | p50 | p99 |
|----------|---------|-----|-----|
| `GET /task/` | 1,880 | 24ms | 77ms |
| `GET /task/:id` | 2,424 | 18ms | 76ms |
| `POST /task/` | 2,228 | 7ms | 46ms |
| `GET /health-check` | 7,941 | 10ms | 52ms |

## Profiling

A pprof server runs on port `6060`. To capture a CPU profile while under load:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

```bash
curl http://localhost:6060/debug/pprof/
```
