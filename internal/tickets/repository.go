package tickets

import (
	"desafio-goweb-abrilgil/internal/domain"
)

type Repository interface {
	GetAll() []domain.Ticket
	GetTicketByDestination(destination string) []domain.Ticket
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Ticket) {

	if len(r.db) == 0 {
		return []domain.Ticket{}
	}

	return r.db
}

func (r *repository) GetTicketByDestination(destination string) []domain.Ticket {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest
}
