package entity

type Order struct {
	OrderId    uint32
	Customer   string
	Items      []uint32
	TotalPrice float32
}
