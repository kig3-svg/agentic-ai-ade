# Development Guide

## Prerequisites

- Go 1.22 or higher
- PostgreSQL 13 or higher (optional, for persistence)
- Docker & Docker Compose (optional, for containerized development)
- Make (optional, for convenience)

## Quick Start

### 1. Clone the Repository
```bash
git clone https://github.com/kig3-svg/agentic-ai-ade.git
cd agentic-ai-ade
```

### 2. Set Up Environment
```bash
cp .env.example .env
# Edit .env with your configuration
```

### 3. Install Dependencies
```bash
make deps
# or
go mod download
```

### 4. Build and Run
```bash
# Using Make
make build
make run

# Or directly with Go
go run ./cmd/ade
```

## Development Workflow

### Running Tests
```bash
make test
```

### Code Formatting
```bash
make fmt
```

### Linting
```bash
make lint
```

### Docker Development

Build and run with Docker Compose:
```bash
docker-compose up -d
```

Stop containers:
```bash
docker-compose down
```

View logs:
```bash
docker-compose logs -f ade
```

## Project Structure

```
agentic-ai-ade/
├── cmd/
│   └── ade/                 # Main application entry point
│       └── main.go
├── internal/
│   ├── app/                 # Application orchestration
│   │   ├── app.go
│   │   └── server.go
│   ├── agent/               # Agent implementation
│   │   ├── agent.go
│   │   ├── pool.go
│   │   └── errors.go
│   └── config/              # Configuration management
│       └── config.go
├── pkg/
│   └── logger/              # Logging utilities
│       └── logger.go
├── go.mod                   # Go module definition
├── Makefile                 # Build automation
├── Dockerfile               # Container image
├── docker-compose.yml       # Development environment
└── README.md                # Project documentation
```

## API Endpoints

### Health Check
```bash
curl http://localhost:8080/health
```

### List Agents
```bash
curl http://localhost:8080/api/v1/agents
```

### Create Task
```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"name":"example-task","input":{}}'
```

## Debugging

### Enable Debug Logging
Set `LOG_LEVEL=debug` in your `.env` file:
```bash
LOG_LEVEL=debug
```

### Using Delve Debugger
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug ./cmd/ade
```

## Contributing

1. Create a feature branch: `git checkout -b feature/your-feature`
2. Make changes and commit: `git commit -am 'Add feature'`
3. Push to the branch: `git push origin feature/your-feature`
4. Submit a pull request

## Common Issues

### Connection Refused
If you get connection refused errors:
1. Check that the server is running: `make run`
2. Verify the port is not already in use: `lsof -i :8080`
3. Check your `.env` configuration

### Database Connection Issues
1. Ensure PostgreSQL is running
2. Verify connection string in `.env`
3. Check database user permissions

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Logrus Logger](https://github.com/sirupsen/logrus)
- [Cobra CLI](https://github.com/spf13/cobra)
