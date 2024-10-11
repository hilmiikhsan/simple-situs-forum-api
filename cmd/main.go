package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/configs"
	membershipsHandler "github.com/hilmiikhsan/situs-forum/internal/handlers/memberships"
	postsHandler "github.com/hilmiikhsan/situs-forum/internal/handlers/posts"
	membershipsRepository "github.com/hilmiikhsan/situs-forum/internal/repository/memberships"
	postsRepository "github.com/hilmiikhsan/situs-forum/internal/repository/posts"
	membersipsService "github.com/hilmiikhsan/situs-forum/internal/service/memberships"
	postsService "github.com/hilmiikhsan/situs-forum/internal/service/posts"
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

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	memberShipRepo := membershipsRepository.NewRepository(db)
	postRepo := postsRepository.NewRepository(db)

	membershipService := membersipsService.NewService(cfg, memberShipRepo)
	postService := postsService.NewService(cfg, postRepo)

	membershipHandler := membershipsHandler.NewHandler(router, membershipService)
	postHandler := postsHandler.NewHandler(router, postService)

	membershipHandler.RegisterRoute()
	postHandler.RegisterRoute()

	router.Run(cfg.Service.Port)
}
