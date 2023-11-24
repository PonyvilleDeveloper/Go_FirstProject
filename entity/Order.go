package entity

type Order struct {
	OrderId    uint32
	Customer   string
	Items      []Item
	TotalPrice float32
}
