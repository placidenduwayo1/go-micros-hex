package routes

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	HandleCreateOrder(ctx *gin.Context)
	HandleGetOrderByID(ctx *gin.Context)
}

type RouteRegistration struct {
	handler OrderHandler
}

func NewRouteRegistration(handler OrderHandler) *RouteRegistration {
	return &RouteRegistration{handler: handler}
}

func (rr *RouteRegistration) RegisterRoutes() *gin.Engine {
	engine := gin.Default()

	api := engine.Group("/api/v1")
	api.POST("/orders", rr.handler.HandleCreateOrder)
	api.GET("/orders/:id", rr.handler.HandleGetOrderByID)
	return engine
}
