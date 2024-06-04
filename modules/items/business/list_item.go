package business

import (
	"context"
	"todo-list/common"
	"todo-list/modules/items/model"
)

type ListItemStorage interface {
	ListItem(ctx context.Context, filter *model.Filter,
		paging *common.Paging, morekeys ...string) ([]model.TodoItem, error)
}

type listItemBusiness struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItemBusiness {
	return &listItemBusiness{store: store}
}

func (biz *listItemBusiness) GetListItem(ctx context.Context, filter *model.Filter,
	paging *common.Paging) ([]model.TodoItem, error) {
	data, err := biz.store.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return data, nil
}
