package service

import (
	"eShop/models"
	"eShop/pkg/repository"
)

func GetAllRacks() ([]models.Rack, error) {
	return repository.GetAllRacks()
}
