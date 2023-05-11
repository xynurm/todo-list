package routes

import (
	"test/handlers"
	"test/pkg/mysql"
	"test/repositories"

	"github.com/gin-gonic/gin"
)

func ActivityRoutes(r *gin.RouterGroup) {
	activityRepository := repositories.RepositoryActivity(mysql.DB)
	h := handlers.HandlerActivity(activityRepository)
	r.GET("/activity-groups", h.GetAllActivity)
	r.POST("/activity-groups", h.CreateActivity)
	r.GET("/activity-groups/:id", h.GetActivityID)
	r.DELETE("/activity-groups/:id", h.DeleteActivity)
	r.PATCH("/activity-groups/:id", h.UpdateActivity)

}
