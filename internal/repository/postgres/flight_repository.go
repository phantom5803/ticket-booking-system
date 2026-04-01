package postgres

import (
	"context"

	"airline-booking/internal/domain"
	"airline-booking/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FlightRepository struct {
	pool *pgxpool.Pool
}

func NewFlightRepository(pool *pgxpool.Pool) *FlightRepository {
	return &FlightRepository{pool: pool}
}

func (r *FlightRepository) List() ([]domain.Flight, error) {
	rows, err := r.pool.Query(
		context.Background(),
		`SELECT id, flight_number, source_airport, destination_airport, departure_time::text, arrival_time::text, available_seats
		 FROM flights ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flights := make([]domain.Flight, 0)
	for rows.Next() {
		var f domain.Flight
		if err := rows.Scan(&f.ID, &f.FlightNumber, &f.From, &f.To, &f.DepartureTime, &f.ArrivalTime, &f.AvailableSeats); err != nil {
			return nil, err
		}
		flights = append(flights, f)
	}
	return flights, rows.Err()
}

func (r *FlightRepository) ReserveSeats(flightID string, seats int) error {
	ctx := context.Background()
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	tag, err := tx.Exec(
		ctx,
		`UPDATE flights
		 SET available_seats = available_seats - $2
		 WHERE id = $1 AND available_seats >= $2`,
		flightID, seats,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		var exists bool
		if err := tx.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM flights WHERE id = $1)`, flightID).Scan(&exists); err != nil {
			return err
		}
		if !exists {
			return repository.ErrFlightNotFound
		}
		return repository.ErrNotEnoughSeats
	}

	return tx.Commit(ctx)
}
