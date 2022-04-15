package repository

import (
	"context"

	"world-heritage-scrape/app/worldheritage/domain/model"
)

// TODO:
// BookRepository : Book における Repository のインターフェース
//  -> 依存性逆転の法則により infra 層は domain 層（本インターフェース）に依存
type WorldheritageRepository interface {
	Store(context.Context) ([]*model.Rows, error)
}
