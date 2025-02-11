# Backend Documentation

## File Structure

```bash
/backend
│── /cmd                     # Entry points (main.go per service)
│   ├── /api                 # API service entry point
│   │   ├── main.go          # API bootstrap logic
│   ├── /migrate             # Database migration service entry point
│   │   ├── main.go          # Database migration logic
│
│── /config                  # Configuration management
│   ├── config.go            # Load and parse configurations
│   ├── config.yaml          # Default app config
│   ├── config_test.go       # Config unit tests
│
│── /internal                # Private application logic (not exposed as a module)
│   ├── /app                 # Application-specific logic
│   │   ├── /services        # Business logic services
│   │   │   ├── user.go      # Example: User service
│   │   │   ├── order.go     # Example: Order service
│   │   │   ├── ...
│   │   │   ├── service_test.go  # Unit tests for services
│   │   ├── /repositories    # Database interaction layer (ORM or raw SQL)
│   │   │   ├── user_repo.go
│   │   │   ├── order_repo.go
│   │   │   ├── repository_test.go
│   │   ├── /models          # Data models (structs)
│   │   │   ├── user.go
│   │   │   ├── order.go
│   │   ├── /validators      # Input validation logic
│   │   │   ├── user_validator.go
│   │   ├── /middlewares     # Middleware (auth, logging, rate limiting)
│   │   │   ├── auth.go
│   │   │   ├── logging.go
│   │   ├── /utils           # Utility functions (helpers)
│   │   │   ├── hash.go
│   │   │   ├── jwt.go
│
│── /api                     # HTTP Handlers (Gin/Fiber/Echo)
│   ├── /v1
│   │   ├── user_handler.go
│   │   ├── order_handler.go
│   ├── router.go            # Router setup
│   ├── middleware.go        # Global middlewares
│
│── /db                      # Database-related files
│   ├── migrations           # Database migrations
│   ├── seeders              # Sample data
│
│── /pkg                     # Shared code that can be imported in other projects
│   ├── logger               # Logging package
│   ├── cache                # Redis caching
│   ├── email                # Email utilities
│   ├── queue                # Message queue (Kafka, RabbitMQ)
│   ├── auth                 # Auth package (JWT, OAuth, etc.)
│
│── /tests                   # Integration & End-to-End (E2E) tests
│   ├── e2e_test.go
│   ├── api_test.go
│
│── /scripts                 # Helper scripts (database reset, deployment, etc.)
│   ├── migrate.sh
│   ├── seed.sh
│
│── /deploy                  # Deployment files
│   ├── Dockerfile
│   ├── docker-compose.yaml
│   ├── k8s                  # Kubernetes manifests
│
│── go.mod                    # Go module file
│── go.sum                    # Go dependencies
│── Makefile                  # Task automation (build, run, test)
│── README.md                 # Documentation
```