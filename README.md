# keynote-api

A Go HTTP API for user authentication and session management, built around PostgreSQL and Ent. The project currently focuses on account registration, login, JWT issuance, JWKS discovery, and refresh-token-backed sessions.

## Project overview

The application exposes a small REST API using Huma on top of Chi, with database access through Ent and PostgreSQL. It is structured as a modular API service and includes a draft access-control design document for future multi-tenant workspace/project permissions.

Main components:

- API entrypoint: `cmd/api`
- Configuration: `config/`
- Database connection and setup: `database/`
- HTTP server and route registration: `internal/httpserver/`
- Auth features: `internal/features/auth/`
- Ent schemas: `internal/ent/schema/`
- JWT helpers and token logic: `internal/common/tokens/`
- Database migration runner: `migrate/`

## Tech stack

- Go
- Chi router
- Huma API framework
- Ent ORM
- PostgreSQL
- JWT with RSA signing keys
- Argon2 password hashing

## Features

- User registration at `POST /auth/register`
- User login at `POST /auth/login`
- Public JWKS endpoint at `GET /.well-known/jwks.json`
- Health endpoint at `GET /health`
- Refresh-token session persistence with cookie-based refresh token handling
- Password hashing using Argon2
- PostgreSQL-backed persistence and schema migrations

## Repository structure

```text
.
├── cmd/
│   ├── api/                 # application entrypoint
│   └── migrate/             # migration runner
├── config/                 # environment/config parsing
├── database/               # PostgreSQL connection utilities
├── internal/
│   ├── common/             # shared helpers (argon, jwt signer)
│   ├── ent/
│   │   ├── schema/         # Ent models for users, credentials, sessions
│   │   └── gen/            # generated Ent client code
│   ├── errors/             # API error helpers
│   ├── features/
│   │   └── auth/           # register/login/jwks routes and services
│   ├── httpserver/         # HTTP server wiring and routes
│   └── types/              # DTO and shared types
├── migrate/
│   ├── migrations/         # SQL migrations
│   └── migrate.go          # migration runner
├── .env.example            # sample environment file
├── dev-jwt.pem            # local JWT private key for development
├── keynote-api-access-control-design.md
├── go.mod
├── go.sum
├── README.md
└── atlas.hcl              # Atlas migration metadata
```

## Prerequisites

- Go 1.26.8 (as declared in `go.mod`)
- PostgreSQL running locally or in a reachable environment
- A JWT private key in PEM format

## Environment configuration

Copy `.env.example` to `.env` and fill in the required values. The application loads environment variables with `godotenv` and parses them using `env`.

Example:

```env
HTTP_ADDR=8080
HTTP_READTIMEOUT=10s
HTTP_WRITETIMEOUT=10s

DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=keynote
DATABASE_SSLMODE=disable
DATABASE_MAX_OPEN_CONNECTIONS=15
DATABASE_MAX_IDLE_CONNECTIONS=5
DATABASE_CONNECTION_TIMEOUT=15m

JWT_PRIVATE_KEY_FILE=./dev-jwt.pem
JWT_ISSUER=keynote-api
JWT_AUDIENCE=keynote-api
JWT_ACCESS_TOKEN_TTL=15m
JWT_REFRESH_TOKEN_TTL=24h
```

## Running the app

Install dependencies and start the API:

```bash
go mod download
go run ./cmd/api
```

By default the server starts on the `HTTP_ADDR` value (for example `:8080`). It logs the API address and the OpenAPI docs URL.

## Database migrations

The project includes a custom migration runner and SQL migrations under `migrate/migrations`.

Apply migrations:

```bash
go run ./cmd/migrate
```

This opens the PostgreSQL connection from the environment config and runs pending migration files in filename order.

## API

### Health check

```http
GET /health
```

Response example:

```json
{
  "status": "Ok",
  "timestamp": "2026-10-13T00:00:00Z"
}
```

### Register

```http
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

### Login

```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

On success, the response includes an access token, token type, expiry, and a `refresh_token` cookie (`HttpOnly`, secure cookie handling for refreshed sessions).

### JWKS

```http
GET /.well-known/jwks.json
```

This returns the public signing keys used to validate JWTs.

## Authentication model

The project currently stores:

- `users`: basic user identity with email and timestamps
- `user_credentials`: hashed password data for each user
- `user_sessions`: refresh-token session records including expiry and revocation tracking

User passwords are hashed with Argon2 before storage, and JWTs are signed with a private key defined by `JWT_PRIVATE_KEY_FILE`.

## Access-control design notes

The repository also contains `keynote-api-access-control-design.md`, which describes a future multi-tenant authorization model based on scopes, roles, and permissions (tenant → workspace → project). That design is intentionally separate from the current auth endpoints and reflects a planned expansion beyond the current API surface.

## Notes

- The project is still in active development and the access-control design is a draft.
- The service is intentionally narrow in scope right now: authentication and session management are the implemented core features.
