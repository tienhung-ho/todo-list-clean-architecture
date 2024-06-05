package postgres

import (
	"context"
	"todo-list/common"
	"todo-list/modules/items/model"
)

func (p *postgresStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	if err := p.db.Table(model.TodoItem{}.TableName()).Where(cond).
		Updates(map[string]interface{}{
			"status": model.ItemStatusDeleted}).Error; err != nil {
		return common.NewErrorResponse(err, "Error delete items from database", err.Error(), "CouldNotDelete")
	}
	return nil
}
