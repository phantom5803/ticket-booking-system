package postgres

import (
	"context"
	"strconv"

	"airline-booking/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingRepository struct {
	pool *pgxpool.Pool
}

func NewBookingRepository(pool *pgxpool.Pool) *BookingRepository {
	return &BookingRepository{pool: pool}
}

func (r *BookingRepository) List() ([]domain.Booking, error) {
	rows, err := r.pool.Query(
		context.Background(),
		`SELECT id, flight_id, passenger_name, passenger_email, seat_count, status
		 FROM bookings ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookings := make([]domain.Booking, 0)
	for rows.Next() {
		var (
			id int64
			b  domain.Booking
		)
		if err := rows.Scan(&id, &b.FlightID, &b.PassengerName, &b.PassengerEmail, &b.SeatCount, &b.Status); err != nil {
			return nil, err
		}
		b.ID = strconv.FormatInt(id, 10)
		bookings = append(bookings, b)
	}
	return bookings, rows.Err()
}

func (r *BookingRepository) Create(input domain.CreateBookingRequest) (domain.Booking, error) {
	var id int64
	err := r.pool.QueryRow(
		context.Background(),
		`INSERT INTO bookings (flight_id, passenger_name, passenger_email, seat_count, status)
		 VALUES ($1, $2, $3, $4, 'CONFIRMED')
		 RETURNING id`,
		input.FlightID, input.PassengerName, input.PassengerEmail, input.SeatCount,
	).Scan(&id)
	if err != nil {
		return domain.Booking{}, err
	}

	return domain.Booking{
		ID:             strconv.FormatInt(id, 10),
		FlightID:       input.FlightID,
		PassengerName:  input.PassengerName,
		PassengerEmail: input.PassengerEmail,
		SeatCount:      input.SeatCount,
		Status:         "CONFIRMED",
	}, nil
}
