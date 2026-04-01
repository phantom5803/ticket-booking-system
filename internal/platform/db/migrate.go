package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

const schemaSQL = `
CREATE TABLE IF NOT EXISTS flights (
  id TEXT PRIMARY KEY,
  flight_number TEXT NOT NULL,
  source_airport TEXT NOT NULL,
  destination_airport TEXT NOT NULL,
  departure_time TIMESTAMPTZ NOT NULL,
  arrival_time TIMESTAMPTZ NOT NULL,
  available_seats INT NOT NULL CHECK (available_seats >= 0)
);

CREATE TABLE IF NOT EXISTS bookings (
  id BIGSERIAL PRIMARY KEY,
  flight_id TEXT NOT NULL REFERENCES flights(id),
  passenger_name TEXT NOT NULL,
  passenger_email TEXT NOT NULL,
  seat_count INT NOT NULL CHECK (seat_count > 0),
  status TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`

const seedFlightsSQL = `
INSERT INTO flights (id, flight_number, source_airport, destination_airport, departure_time, arrival_time, available_seats)
VALUES
  ('F1', 'AI-101', 'DEL', 'BOM', '2026-04-02T09:00:00Z', '2026-04-02T11:10:00Z', 60),
  ('F2', 'AI-202', 'BLR', 'DEL', '2026-04-02T14:15:00Z', '2026-04-02T16:50:00Z', 40)
ON CONFLICT (id) DO NOTHING;
`

func MigrateAndSeed(ctx context.Context, pool DBTX) error {
	if _, err := pool.Exec(ctx, schemaSQL); err != nil {
		return err
	}
	if _, err := pool.Exec(ctx, seedFlightsSQL); err != nil {
		return err
	}
	return nil
}

type DBTX interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}
