package controller

import "eShop/pkg/service"

func BulkGetOrdersByIDs(IDs string) error {
	return service.BulkGetOrdersByIDs(IDs)
}
