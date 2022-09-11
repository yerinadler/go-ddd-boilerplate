package order

type OrderItem struct {
	ProductId string `json:"productId" bson:"productId"`
	Quantity  int    `json:"quantity" bson:"quantity"`
	Price     int    `json:"price" bson:"price"`
}
