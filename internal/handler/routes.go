package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wbl0/internal/service"
)

type Handler struct {
	Store *service.Store
}

func NewHandler(services *service.Store) *Handler {
	return &Handler{Store: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Handle("GET", "/", h.Home)
	router.Handle("GET", "/show", h.ShowData)

	fileServer := http.FileServer(http.Dir("ui/html/static/"))
	router.GET("/static/*filepath", gin.WrapH(http.StripPrefix("/static/", fileServer)))

	return router
}
