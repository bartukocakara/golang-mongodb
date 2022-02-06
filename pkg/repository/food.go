package repository

type FoodRepository interface {
	GetFood(int)
	GetFoods()
	CreateFood()
	UpdateFood(int)
	DeleteFood(int)
}