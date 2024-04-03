package models

type Rack struct {
	ID     int
	Number string
}

type Order struct {
	Number         string
	OrderID        int
	ProductName    string
	ProductID      int
	OrderCount     int
	SecondaryRacks string
}
