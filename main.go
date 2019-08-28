package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/oentoro/ms.account/dbclient"
	"github.com/oentoro/ms.account/service"
)

var appName = "main service"

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/accounts/:accountid", func(c *gin.Context) {
		accountid := c.Param("accountid")
		account, err := service.DBClient.QueryAccount(accountid)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		if account.Name == "" {
			c.JSON(404, gin.H{
				"message": "account not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": account,
		})
	})

	r.Run(":8004")
}
