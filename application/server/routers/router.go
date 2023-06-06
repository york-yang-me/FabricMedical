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
		apiV1.POST("/createAuthorizing", v1.CreateSelling)
		apiV1.POST("/createAuthorizingByBuy", v1.CreateAuthorizingByBuy)
		apiV1.POST("/queryAuthorizingList", v1.QueryAuthorizingList)
		apiV1.POST("/queryAuthorizingListByBuyer", v1.QueryAuthorizingListByBuyer)
		apiV1.POST("/updateAuthorizing", v1.UpdateAuthorizing)
		apiV1.POST("/createAppointing", v1.CreateAppointing)
		apiV1.POST("/queryAppointingList", v1.QueryAppointingList)
		apiV1.POST("/queryAppointingListByHospital", v1.QueryAppointingListByHospital)
		apiV1.POST("/updateAppointing", v1.UpdateAppointing)
	}
	// static files router
	r.StaticFS("/web", http.Dir("./dist/"))
	return r
}
