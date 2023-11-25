package entity

type Cart struct {
	UserId   uint64
	ItemId   uint64
	BrandId  uint64
	MnfctrId uint64
	Count    uint16
	CartId   uint32
}
