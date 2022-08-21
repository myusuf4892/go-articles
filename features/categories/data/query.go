package data

import (
	"articles/features/categories"

	"gorm.io/gorm"
)

type mysqlCategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(conn *gorm.DB) categories.Data {
	return &mysqlCategoryRepo{
		db: conn,
	}
}

func (repo *mysqlCategoryRepo) Insert(dataReq categories.Core) (row int, err error) {
	data := fromCoreToRepo(&dataReq)

	srv := repo.db.Create(&data)

	return int(srv.RowsAffected), srv.Error
}

func (repo *mysqlCategoryRepo) Get() (dataRes []categories.Core, err error) {
	data := []Category{}

	srv := repo.db.Find(&data)

	dataRes = repoToCoreList(data)

	return dataRes, srv.Error
}
