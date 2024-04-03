package repository

import (
	"eShop/db"
	"eShop/models"
)

func GetAllRacks() (racks []models.Rack, err error) {
	if err = db.GetDBConn().Raw("SELECT * FROM racks").Scan(&racks).Error; err != nil {
		return nil, err
	}

	return racks, nil
}
