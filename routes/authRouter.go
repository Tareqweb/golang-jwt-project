package routes

import(
	controllers "github.com/tareq/golang-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incoming Routes *gin.Engine){
	incomingRoutes.POST("users/signup", controllers.signup)
	incomingRoutes.POST("users/login", controllers.login)


}