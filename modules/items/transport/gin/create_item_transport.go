package ginitem

import (
	"net/http"
	"todo-list/common"
	"todo-list/modules/items/business"
	"todo-list/modules/items/model"
	"todo-list/modules/items/storage/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if data.Status == nil {
			defaultStatus := model.ItemStatus(0)
			data.Status = &defaultStatus
		}

		store := postgres.NewPostgresStore(db)

		biz := business.NewCreateItemBiz(store)

		if err := biz.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(data.Id))

	}
}
