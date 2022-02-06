package repository

type MenuRepository interface {
	GetMenu(int)
	GetMenus()
	CreateMenu()
	UpdateMenu(int)
	DeleteMenu(int)
}