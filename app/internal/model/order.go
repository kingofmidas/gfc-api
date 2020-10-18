package model

// Order ...
type Order struct {
	ID       int         `json:"id"`
	ItemList []OrderItem `json:"item_list"`
	Status   string      `json:"status"`
}

// OrderItem ...
type OrderItem struct {
	ItemID int `json:"item_id"`
	Count  int `json:"count"`
}
