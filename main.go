package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/oentoro/ms.account/service"
	"github.com/oentoro/ms.account/dbclient"
)

var appName = "main service"

func initializeBoltClient(){
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}

func main(){
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8004")
}