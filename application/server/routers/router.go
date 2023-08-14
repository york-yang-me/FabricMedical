package routers

import (
	v1 "application/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter init router information
func InitRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/hello", v1.Hello)
		apiV1.POST("/queryAccountList", v1.QueryAccountList)
		apiV1.POST("/createRealSequence", v1.CreateRealSequence)
		apiV1.POST("/queryRealSequenceList", v1.QueryRealSequenceList)
	}
	// static files router
	r.StaticFS("/web", http.Dir("./dist/"))
	return r
}
