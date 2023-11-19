package main

import (
	"fmt"
	"log"

	"github.com/RaniaMidaoui/goMart-gateway/pkg/auth"
	"github.com/RaniaMidaoui/goMart-gateway/pkg/config"
	"github.com/RaniaMidaoui/goMart-gateway/pkg/order"
	"github.com/RaniaMidaoui/goMart-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	fmt.Println("Another test")
	r.Run(c.Port)
}
