package repository

import (
	"world-heritage-scrape/app/worldheritage/domain/model"
)

// TODO:
// WorldherirageRepository : Worldheritage における Repository のインターフェース
//  -> 依存性逆転の法則により infrastructure 層は domain 層（本インターフェース）に依存
type WorldheritageRepository interface {
	Store(data []model.Worldheritage) ([]*model.Worldheritage, error)
}
