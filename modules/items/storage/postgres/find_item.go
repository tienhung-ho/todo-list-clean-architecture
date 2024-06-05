package postgres

import (
	"context"
	"todo-list/common"
	"todo-list/modules/items/model"

	"gorm.io/gorm"
)

func (p *postgresStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := p.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)

	}

	return &data, nil
}
