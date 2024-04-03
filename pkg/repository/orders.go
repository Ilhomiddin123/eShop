package repository

import (
	"eShop/db"
	"eShop/models"
	"fmt"
)

func GetAllOrderByRackID(rackID int, orderIDs string) (orders []models.Order, err error) {
	sqlQuery := fmt.Sprintf(`SELECT r.number,
				   o.id     as order_id,
				   p.title  as product_name,
				   p.id  as product_id,
				   od.count as order_count,
				   (SELECT string_agg(r2.number::varchar, ', ')
					FROM racks_products
					JOIN racks r2 on r2.id = racks_products.rack_id
					WHERE product_id = p.id
					  AND priority = 'secondary') as secondary_racks
			FROM racks r
					 JOIN racks_products rp on r.id = rp.rack_id
					 JOIN products p on p.id = rp.product_id
					 JOIN order_details od on p.id = od.product_id
					 JOIN orders o on o.id = od.order_id
			WHERE r.id = %v
			  AND o.id in (%v)
			  AND rp.priority = 'primary';`, rackID, orderIDs)

	if err = db.GetDBConn().Raw(sqlQuery).Scan(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}
