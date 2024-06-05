package postgres

import (
	"context"
	"todo-list/common"
	"todo-list/modules/items/model"
)

func (p *postgresStore) ListItem(ctx context.Context, filter *model.Filter, paging *common.Paging, morekeys ...string) ([]model.TodoItem, error) {
	var result []model.TodoItem

	db := p.db.Where("status != ?", model.ItemStatusDeleted)

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", f.Status)
		}
	}

	if err := db.Table(model.TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.NewErrorResponse(err, "Error count items from database", err.Error(), "CouldNotCount")
	}

	if err := db.Order("id desc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
		return nil, common.NewErrorResponse(err, "Error fetching items from database", err.Error(), "ErrorFetching")
	}

	return result, nil
}
