# Airline Ticket Booking System (Go Backend)

Starter backend for an airplane ticket booking system using Go.

## What is included

- Go HTTP API server
- Basic routes for flights and bookings
- In-memory seat reservation logic (prevents overbooking)
- `docker-compose` for PostgreSQL and Redis (ready for next phase)

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
2. (Optional) start infra:
   - `docker compose up -d`
3. Run API:
   - `go run ./cmd/server`
4. Test:
   - `curl http://localhost:8080/health`
   - `curl http://localhost:8080/flights`

## Suggested next upgrades

- Move from in-memory store to PostgreSQL
- Add user auth (JWT)
- Add payment workflow + idempotency keys
- Add booking cancellation and seat release
