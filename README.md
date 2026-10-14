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
- Auth middleware: `internal/common/middleware/auth/`
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
- Protected authenticated user endpoint at `GET /auth/me`
- Public JWKS endpoint at `GET /.well-known/jwks.json`
- Health endpoint at `GET /health`
- Refresh-token session persistence with cookie-based refresh token handling
- Password hashing using Argon2
- PostgreSQL-backed persistence and schema migrations
- Protected route grouping via JWT verification middleware

## Repository structure

```text
.
├── cmd/
│   ├── api/                    # application entrypoint
│   └── migrate/                # migration runner
├── config/
│   └── config.go               # environment/config parsing
├── database/
│   ├── database.go             # PostgreSQL connection helpers
│   └── ent.go                  # Ent client wrapper setup
├── internal/
│   ├── common/
│   │   ├── middleware/
│   │   │   └── auth/
│   │   │       └── verifyjwt.go    # JWT verification middleware
│   │   ├── password/
│   │   │   └── argon.go
│   │   └── tokens/
│   │       └── signer.go           # JWT signing + verification helpers
│   ├── ent/
│   │   ├── gen/                   # generated Ent client
│   │   ├── schema/
│   │   │   ├── users.go
│   │   │   ├── usercredentials.go
│   │   │   └── usersessions.go
│   │   └── generate.go
│   ├── errors/
│   │   ├── apperr.go
│   │   └── httperr.go
│   ├── features/
│   │   └── auth/
│   │       ├── jwks/
│   │       │   └── route.go
│   │       ├── login/
│   │       │   ├── route.go
│   │       │   └── service.go
│   │       ├── me/
│   │       │   ├── route.go
│   │       │   └── service.go
│   │       ├── register/
│   │       │   ├── route.go
│   │       │   └── service.go
│   │       └── routes.go           # public/protected auth mount points
│   ├── httpserver/
│   │   ├── health.go
│   │   ├── routes.go               # protected group registration
│   │   └── server.go
│   └── types/
│       └── user.go
├── migrate/
│   ├── embed.go
│   ├── migrate.go
│   └── migrations/                # SQL migration files
├── .env.example                    # sample environment file
├── .gitignore
├── dev-jwt.pem                     # local JWT private key for development
├── atlas.hcl                       # Atlas migration metadata
├── keynote-api-access-control-design.md
├── go.mod
├── go.sum
├── README.md
└── .env                            # local env file (not committed)
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

Generating Ent code and Atlas migrations
Generate Ent code

After modifying or adding Ent schemas under internal/ent/schema/, regenerate the Ent client and related code:

```bas
go generate ./internal/ent/
```

Review the generated changes before committing them.

Generate Atlas migrations

After updating the Ent schemas, generate a new Atlas migration using:

```bash
atlas migrate diff {name} --env local
```

Replace {name} with a descriptive migration name, such as add_user_sessions.

The generated SQL migration files are stored under migrate/migrations/. Review the generated SQL before applying the migration.

Apply database migrations

Once the migration files have been generated and reviewed, apply pending migrations to the configured PostgreSQL database:

go run ./cmd/migrate


Recommended workflow:

Update the Ent schemas in internal/ent/schema/.
Regenerate the Ent client using go generate ./internal/ent/.
Generate an Atlas migration using atlas migrate diff {name} --env local.
Review the generated migration SQL.
Apply the migration using go run ./cmd/migrate.

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

### Me (authenticated user)

```http
GET /auth/me
Authorization: Bearer <token>
```

This route is registered on the protected API group and requires a valid JWT. The middleware validates the Authorization header and injects the parsed claims into the request context before the handler runs.

### JWKS

```http
GET /.well-known/jwks.json
```

This returns the public signing keys used to validate JWTs.

## Protected routes and JWT middleware

Protected endpoints are mounted on a dedicated Huma group that applies the JWT middleware before route handlers execute.

```go
protectedAPI := huma.NewGroup(api)
protectedAPI.UseMiddleware(authMiddleware.VerifyJWT(api, signer))

auth.Mount(api, protectedAPI, db, signer)
```

The middleware:

- reads the `Authorization: Bearer <token>` header
- validates the token with the configured signer
- stores the claims in the context under `jwt_claims`
- rejects requests with `401 Unauthorized` when the JWT is missing, malformed, or invalid

This pattern is what secures the `GET /auth/me` route. The handler pulls the subject from the context and resolves the current user.

## Authentication model

The project currently stores:

- `users`: basic user identity with email and timestamps
- `user_credentials`: hashed password data for each user
- `user_sessions`: refresh-token session records including expiry and revocation tracking

User passwords are hashed with Argon2 before storage, and JWTs are signed with a private key defined by `JWT_PRIVATE_KEY_FILE`.


