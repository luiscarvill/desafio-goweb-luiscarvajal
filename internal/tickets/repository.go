package tickets

import (
	"context"
	domain "desafio-go-web/internal/domains"
	"fmt"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetAverageTickets() ([]domain.Ticket, error)
	GetCounterTickets() ([]domain.Ticket, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}
func (r *repository) GetAverageTickets() ([]domain.Ticket, error) {
	fmt.Println("here the average", r)
	return []domain.Ticket{}, fmt.Errorf("Error al sacar el promedio")
}
func (r *repository) GetCounterTickets() ([]domain.Ticket, error) {
	fmt.Println("here the counter", r)
	return []domain.Ticket{}, fmt.Errorf("Error al sacar el promedio")
}
