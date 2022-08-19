package data

import (
	"articles/features/articles"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepo(conn *gorm.DB) articles.Data {
	return &mysqlArticleRepository{
		db: conn,
	}
}

func (repo *mysqlArticleRepository) Insert(dataReq articles.Core) (row int, err error) {
	data := fromCore(&dataReq)

	srv := repo.db.Create(&data)

	return int(srv.RowsAffected), srv.Error
}

func (repo *mysqlArticleRepository) Get() (dataRes []articles.Core, err error) {
	data := []Article{}

	srv := repo.db.Preload("Category").Find(&data)

	dataRes = toCoreList(data)

	return dataRes, srv.Error
}
