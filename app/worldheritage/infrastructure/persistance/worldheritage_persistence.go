package persistance

import (
	"context"
	"world-heritage-scrape/app/worldheritage/domain/model"
	"world-heritage-scrape/app/worldheritage/domain/repository"

	"gorm.io/gorm"
)

// WorldheritageにおけるPersistenceのインターフェース
type worldheritagePersistence struct {
	Conn *gorm.DB
}

// Worldheritageデータに関するPersistenceを生成
func NewworldheritagePersistence(conn *gorm.DB) repository.WorldheritageRepository {
	return &worldheritagePersistence{Conn: conn}
}

func (bp worldheritagePersistence) GetAll(context.Context) ([]*model.Worldheritage, error) {
	book1 := model.Worldheritage{}
	book1.Name = "test1"

	book2 := model.Worldheritage{}
	book2.Name = "test2"

	return []*model.Worldheritage{&book1, &book2}, nil
}

func (bp worldheritagePersistence) Store(data []model.Worldheritage) ([]*model.Worldheritage, error) {
	model.Worldheritages = data

	return []*model.Worldheritage{}, nil
}
