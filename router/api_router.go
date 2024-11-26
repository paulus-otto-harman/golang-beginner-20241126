package router

import (
	"20241126/controller"

	"github.com/gin-gonic/gin"
)

func APIRouter(r *gin.Engine, ctl controller.Controller) {
	r.GET("/couriers", ctl.Courier.GetCouriers)
	r.POST("/couriers", ctl.Courier.CreateCourier)

	r.GET("/shipping", ctl.Shipping.Cost)

	r.GET("/customers", ctl.Customer.GetCustomers)
	r.POST("/customers", ctl.Customer.CreateCustomer)

	r.GET("/products", ctl.Product.GetProducts)
	r.POST("/products", ctl.Product.CreateProduct)
}
