
// folowing example

user-service/
├── cmd/              # Entry point(s)
│   └── http/        # Main REST/HTTP API
│       └── main.go
├── internal/         # Core application logic
│   ├── config/      # Configuration files
│   ├── db/          # Database layer
│   │   ├── migrations/ # SQL schema changes
│   │   ├── models/  # DB-specific structs
│   │   └── repos/   # Repository implementations
│   ├── http/        # HTTP handlers
│   └── services/    # Business logic
└── infrastructure/   # External systems
