package controller

import "gorm.io/gorm"

type Controller struct {
	Courier  CourierController
	Customer CustomerController
	Product  ProductController
	Shipping ShippingController
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		Courier:  *NewCourierController(db),
		Customer: *NewCustomerController(db),
		Product:  *NewProductController(db),
		Shipping: *NewShippingController(db),
	}
}
