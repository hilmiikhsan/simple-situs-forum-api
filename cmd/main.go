package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/situs-forum/internal/handlers/memberships"
)

func main() {
	router := gin.Default()

	membershipHandler := memberships.NewHandler(router)
	membershipHandler.RegisterRoute()

	router.Run(":9988")
}
