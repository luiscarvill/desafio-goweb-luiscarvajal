package tickets

import (
	domain "desafio-go-web/internal/domains"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(c *gin.Context, str string) ([]domain.Ticket, error)
	AverageDestination(c *gin.Context, dest string) ([]domain.Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTotalTickets(c *gin.Context, dest string) ([]domain.Ticket, error) {
	return s.repository.GetCounterTickets()
}

func (s *service) AverageDestination(c *gin.Context, dest string) ([]domain.Ticket, error) {
	return s.repository.GetAverageTickets()
}
