package repository

type OrderRepository interface {
	GetOrder(int)
	GetOrders()
	CreateOrder()
	UpdateOrder(int)
	DeleteOrder(int)
}