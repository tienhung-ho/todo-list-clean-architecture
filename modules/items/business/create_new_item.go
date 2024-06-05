package business

import (
	"context"
	"strings"
	"todo-list/common"
	"todo-list/modules/items/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBusiness struct {
	store CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{store: store}
}

func (biz *createItemBusiness) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		common.NewErrorResponse(model.ErrTitleBlank, "Can not leave the title blank", "Title empty!", "TitleEmpty")
		return model.ErrTitleBlank
	}

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil

}
