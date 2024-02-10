package handler

import (
	"net/http"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/dto"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetHealth(c *gin.Context) {
	health := dto.NewHealthOK()
	c.IndentedJSON(http.StatusOK, health)
}
