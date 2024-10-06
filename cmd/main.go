package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/configs"
	membershipsHandler "github.com/hilmiikhsan/situs-forum/internal/handlers/memberships"
	membershipsRepository "github.com/hilmiikhsan/situs-forum/internal/repository/memberships"
	"github.com/hilmiikhsan/situs-forum/pkg/internal_sql"
)

func main() {
	router := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("failed to initialize config: ", err)
	}

	cfg = configs.Get()

	db, err := internal_sql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	_ = membershipsRepository.NewRepository(db)

	membershipHandler := membershipsHandler.NewHandler(router)
	membershipHandler.RegisterRoute()

	router.Run(cfg.Service.Port)
}
