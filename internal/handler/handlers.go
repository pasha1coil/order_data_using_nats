package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wbl0/internal/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Service: services}
}

func (h *Handler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Home Page!",
	})
}

func (h *Handler) ShowData(c *gin.Context) {
	uid := c.Param("uid")
	data := h.Service.GetFromCacheByUID(uid)
	c.JSON(http.StatusOK, data)
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Handle("GET", "/", h.Home)
	router.Handle("GET", "/show/:uid", h.ShowData)

	fileServer := http.FileServer(http.Dir("ui/html/static/"))
	router.GET("/static/*filepath", gin.WrapH(http.StripPrefix("/static/", fileServer)))

	return router
}
