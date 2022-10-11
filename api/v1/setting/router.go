package setting

import "github.com/gin-gonic/gin"

func AuthRouter(g *gin.RouterGroup) {
	g.GET("/cats", GetCatList)
}
