package entity

type Item struct {
	ItemId  uint32  `json:"ItemId"`
	Article uint32  `json:"Article"`
	Name    string  `json:"Name"`
	Price   float32 `json:"Price"`
	Creator string  `json:"Creator"`
}
