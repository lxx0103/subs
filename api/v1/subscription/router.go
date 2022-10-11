package subscription

import "github.com/gin-gonic/gin"

func AuthRouter(g *gin.RouterGroup) {
	g.POST("/subscriptions", NewSubscription)
	g.GET("/subscriptions", GetSubscriptionList)
	g.PUT("/subscriptions/:id", UpdateSubscription)

}
