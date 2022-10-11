package subscription

import (
	"subs/core/response"
	"subs/service"

	"github.com/gin-gonic/gin"
)

// @Summary 新建订阅
// @Id 201
// @Tags 订阅管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param subscription body SubscriptionNew true "组织信息"
// @Success 200 object response.SuccessRes{data=string} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /subscriptions [POST]
func NewSubscription(c *gin.Context) {
	var subscription SubscriptionNew
	if err := c.ShouldBindJSON(&subscription); err != nil {
		response.ResponseError(c, "数据格式错误", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	subscription.OpenID = claims.OpenID
	subscriptionService := NewSubscriptionService()
	err := subscriptionService.NewSubscription(subscription)
	if err != nil {
		response.ResponseError(c, "内部错误", err)
		return
	}
	response.Response(c, "创建成功")
}

// @Summary 订阅列表
// @Id 202
// @Tags 订阅管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param page_id query int true "页码"
// @Param page_size query int true "每页行数"
// @Param cat_id query int false "学生ID"
// @Param sub_cat_id query int false "员工ID"
// @Param sub_type_id query int false "学生姓名"
// @Param source_id query int false "员工姓名"
// @Param driver query string false "扫码时间开始"
// @Param pay_type_id query int false "扫码时间结束"
// @Success 200 object response.ListRes{data=[]SubscriptionResponse} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /subscriptions [GET]
func GetSubscriptionList(c *gin.Context) {
	var filter SubscriptionFilter
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.ResponseError(c, "数据格式错误", err)
		return
	}
	subscriptionService := NewSubscriptionService()
	count, list, err := subscriptionService.GetSubscriptionList(filter)
	if err != nil {
		response.ResponseError(c, "内部错误", err)
		return
	}
	response.ResponseList(c, filter.PageID, filter.PageSize, count, list)
}

// @Summary 根据ID更新订阅
// @Id 203
// @Tags 订阅管理
// @version 1.0
// @Accept application/json
// @Produce application/json
// @Param id path int true "订阅ID"
// @Param subscription_info body SubscriptionUpdate true "订阅信息"
// @Success 200 object response.SuccessRes{data=string} 成功
// @Failure 400 object response.ErrorRes 内部错误
// @Router /subscriptions/:id [PUT]
func UpdateSubscription(c *gin.Context) {
	var uri SubscriptionID
	if err := c.ShouldBindUri(&uri); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	var subscription SubscriptionUpdate
	if err := c.ShouldBindJSON(&subscription); err != nil {
		response.ResponseError(c, "BindingError", err)
		return
	}
	claims := c.MustGet("claims").(*service.CustomClaims)
	subscription.OpenID = claims.OpenID
	subscriptionService := NewSubscriptionService()
	err := subscriptionService.UpdateSubscription(uri.ID, subscription)
	if err != nil {
		response.ResponseError(c, "DatabaseError", err)
		return
	}
	response.Response(c, "更新成功")
}
