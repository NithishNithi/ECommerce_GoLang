package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kishorens18/ecommerce/cmd/client/handlercontrollers"
)




func AppRoutes(r *gin.Engine) {
	r.POST("/signup", handlercontrollers.SignUp)
	r.POST("/signin", handlercontrollers.SignIn)
	r.GET("/deletecustomer",handlercontrollers.DeleteCustomer )
	r.POST("/updatecustomer", handlercontrollers.UpdateCustomer)
	r.GET("/getbyid", handlercontrollers.GetByCustomerId)
	r.POST("/updatepassword",handlercontrollers.UpdatePassword )
}
