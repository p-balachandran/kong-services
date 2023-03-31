package db

import (
	"github.com/p-balachandran/kong-services/server/models"
)

type DBService interface {
	KongSvcs() ([]models.Services, error)
	KingSvc(int) (*models.Service, error)
}
