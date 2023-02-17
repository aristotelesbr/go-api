package routes

import (
	list_user_controller "github.com/aristotelesbr/go-api/src/controller/users_controller/list"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/list-users", list_user_controller.Call)
}
