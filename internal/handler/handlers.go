package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pasha1coil/order_data_using_nats/internal/service"
	"net/http"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Service: services}
}

func (h *Handler) ShowData(c *gin.Context) {
	uid := c.Param("uid")
	data := h.Service.GetFromCacheByUID(uid)
	if data.OrderUid == "" {
		c.JSON(http.StatusOK, "Not such data")
	} else {
		c.JSON(http.StatusOK, data)
	}

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	fileServer := http.FileServer(http.Dir("internal/ui/html/static/"))
	router.GET("/", gin.WrapH(http.StripPrefix("/", fileServer)))
	router.Handle("GET", "/show/:uid", h.ShowData)

	return router
}
