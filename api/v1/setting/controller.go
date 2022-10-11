package setting

import (
	"subs/core/response"

	"github.com/gin-gonic/gin"
)

// @Summary 类别列表
// @Id 301
// @Tags 类别管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数（5/10/15/20）"
// @Param name query string false "类别名称"
// @Success 200 object response.ListRes{data=[]CatResponse} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /cats [GET]
func GetCatList(c *gin.Context) {
	var filter CatFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	settingService := NewSettingService()
	count, list, err := settingService.GetCatList(filter)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.ResponseList(c, filter.PageID, filter.PageSize, count, list)
}
