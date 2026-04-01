# Airline Ticket Booking System (Go Backend)

Starter backend for an airplane ticket booking system using Go.

## What is included

- Go HTTP API server
- Basic routes for flights and bookings
- Professional layered folder structure (`domain`, `repository`, `service`, `transport/http`)
- PostgreSQL-backed repositories with startup schema/seed migration
- `docker-compose` for PostgreSQL and Redis

## API Endpoints

- `GET /health`
- `GET /flights`
- `GET /bookings`
- `POST /bookings`

Example booking body:

```json
{
  "flightId": "F1",
  "passengerName": "Sarthak",
  "passengerEmail": "sarthak@example.com",
  "seatCount": 2
}
```

## Run locally

1. Make sure Go 1.22+ is installed.
2. Start infra:
   - `docker compose up -d`
   - Postgres will be exposed on `localhost:5433`
3. Run API:
   - `go run ./cmd/server`
4. Test:
   - `curl http://localhost:8080/health`
   - `curl http://localhost:8080/flights`

## Suggested next upgrades

- Add user auth (JWT)
- Add payment workflow + idempotency keys
- Add booking cancellation and seat release
