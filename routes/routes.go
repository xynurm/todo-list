package routes

import "github.com/gin-gonic/gin"

func RoutesInit(r *gin.RouterGroup) {
	ActivityRoutes(r)
	TodoRoutes(r)
}
