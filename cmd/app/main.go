package main

import (
	"github.com/kviatkovsky/auth_service/db"
	"github.com/kviatkovsky/auth_service/internal/config"
	"github.com/kviatkovsky/auth_service/internal/user"
	"github.com/kviatkovsky/auth_service/router"
)

func main() {
	cfg := config.MustLoad()
	db, err := db.NewDatabase(cfg)
	if err != nil {
		panic(err)
	}

	userRep := user.NewRepository(db.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
