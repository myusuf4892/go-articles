package data

import (
	"articles/features/articles"

	"gorm.io/gorm"
)

type mysqlPostRepo struct {
	db *gorm.DB
}

func NewArticleRepo(conn *gorm.DB) articles.Data {
	return &mysqlPostRepo{
		db: conn,
	}
}

func (repo *mysqlPostRepo) Insert(dataReq articles.Core) (row int, err error) {
	data := fromCoreToRepo(&dataReq)

	srv := repo.db.Create(&data)

	return int(srv.RowsAffected), srv.Error
}

func (repo *mysqlPostRepo) Get() (dataRes []articles.Core, err error) {
	data := []Article{}

	srv := repo.db.Preload("Category").Find(&data)

	dataRes = repoToCoreList(data)

	return dataRes, srv.Error
}
