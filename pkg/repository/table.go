package repository

type TableRepository interface {
	GetTable(int)
	GetTables()
	CreateTable()
	UpdateTable(int)
	DeleteTable(int)
}