package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"desafio-goweb-abrilgil/cmd/server/handler"
	"desafio-goweb-abrilgil/internal/domain"
	"desafio-goweb-abrilgil/internal/tickets"

	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	list, err := LoadTicketsFromFile("./tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}
	ticketsRepo := tickets.NewRepository(list)
	ticketsService := tickets.NewService(ticketsRepo)
	ticketsController := handler.NewController(ticketsService)

	r := gin.Default()
	versionOne := r.Group("/v1")
	{
		versionOne.GET("/ticket/getByCountry/:dest", ticketsController.GetTicketsByCountry)
		versionOne.GET("/ticket/getAverage/:dest", ticketsController.AverageDestination)
	}


	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
