package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kviatkovsky/auth_service/internal/user"
)

func ApiKeyAuth(h *user.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Api-key")

		if len(apiKey) == 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "Missing API key"})
			c.Abort()
			return
		}

		_, err := h.Service.GetAuthByApiKey(c, apiKey)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}
	}
}