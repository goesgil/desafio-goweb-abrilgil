package tickets

type Service interface {
	GetTotalTickets(destination string) int
	AverageDestination(destination string) float64
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(destination string) int {

	tickets := s.repository.GetTicketByDestination(destination)

	return len(tickets)
}

func (s *service) AverageDestination(destination string) float64 {

	ticketsByDestination := s.repository.GetTicketByDestination(destination)

	// promedio de personas que viajan en un dia a ese pais
	var total float64
	total = float64(len(ticketsByDestination)) * 100 / float64(len(s.repository.GetAll()))
	return total
}
