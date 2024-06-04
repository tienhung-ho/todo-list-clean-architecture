package postgres

import (
	"context"
	"todo-list/modules/items/model"
)

func (p *postgresStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := p.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
