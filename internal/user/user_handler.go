package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, c.Query("username"))

	profile, err := h.Service.GetProfile(c, c.Query("username"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, profile)
}