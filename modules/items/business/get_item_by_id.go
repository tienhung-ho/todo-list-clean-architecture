package business

import (
	"context"
	"todo-list/modules/items/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type getItemBusiness struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBusiness {
	return &getItemBusiness{store: store}
}

func (biz *getItemBusiness) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	return data, nil

}
