package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kviatkovsky/auth_service/internal/services/auth"
	"github.com/kviatkovsky/auth_service/internal/user"
)

const (
	api_v1 = "/api/v1/"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()
	r.Use(auth.ApiKeyAuth(userHandler))

	r.GET(fmt.Sprintf("%s/profile", api_v1), userHandler.GetProfile)

}

func Start(addr string) {
	if err := r.Run(addr); err != nil {
		panic(err)
	}
}
