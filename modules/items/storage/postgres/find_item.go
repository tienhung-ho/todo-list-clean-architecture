package postgres

import (
	"context"
	"todo-list/modules/items/model"
)

func (p *postgresStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := p.db.Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
