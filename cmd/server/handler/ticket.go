package handler

import (
	"desafio-goweb-abrilgil/internal/tickets"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetTicketsByCountry(c *gin.Context)
	AverageDestination(c *gin.Context)
}

type controller struct {
	service tickets.Service
}

func NewController(s tickets.Service) Controller {
	return &controller{
		service: s,
	}
}

/*
Requerimiento 3:
	Desarrollar los métodos correspondientes a la estructura Ticket.
	Uno de ellos debe devolver la cantidad de tickets de un destino.
	El otro método debe devolver el promedio de personas que viajan
	a un país determinado en un dia.
*/

func (s *controller) GetTicketsByCountry(c *gin.Context) {

	destination := c.Param("dest")

	tickets := s.service.GetTotalTickets(destination)

	c.JSON(200, tickets)
}

func (s *controller) AverageDestination(c *gin.Context) {

	destination := c.Param("dest")

	avg := s.service.AverageDestination(destination)

	c.JSON(200, avg)
}
