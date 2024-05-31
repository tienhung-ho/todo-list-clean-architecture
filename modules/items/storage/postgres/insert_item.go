package postgres

import (
	"context"
	"todo-list/modules/items/model"
)

func (p *postgresStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := p.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
