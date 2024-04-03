package service

import (
	"eShop/models"
	"eShop/pkg/repository"
	"fmt"
)

func BulkGetOrdersByIDs(IDs string) error {
	racks, err := repository.GetAllRacks()
	if err != nil {
		return err
	}

	for _, rack := range racks {
		rackOrders, err := repository.GetAllOrderByRackID(rack.ID, IDs)
		if err != nil {
			return err
		}
		if len(rackOrders) == 0 {
			continue
		}
		fmt.Println("===", rack.Number)
		formattedOrderOutput(rackOrders)
	}

	return nil
}

func formattedOrderOutput(orders []models.Order) {
	for _, order := range orders {
		fmt.Printf("%s (id=%d)\n", order.ProductName, order.ProductID)
		fmt.Printf("заказ %d, %dшт\n", order.OrderID, order.OrderCount)
		if order.SecondaryRacks != "" {
			fmt.Printf("доп стеллаж: %s\n", order.SecondaryRacks)
		}
		fmt.Printf("\n")
	}
}
