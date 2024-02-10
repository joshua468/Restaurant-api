package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/joshua468/restaurant-management-app/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.POST("/orderItems-order/order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.PATCH("/orderItems/orderItem_id", controller.UpdateOrderItem())
}
