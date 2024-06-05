package ginitem

import (
	"net/http"
	"strconv"
	"todo-list/common"
	"todo-list/modules/items/business"
	"todo-list/modules/items/model"
	"todo-list/modules/items/storage/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := postgres.NewPostgresStore(db)
		biz := business.NewUpdateItemBiz(store)

		if err := biz.UpdateItemById(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccesResponse(true))

	}
}
