package main

import (
	"bookshop/Api"
	"bookshop/Entry/getDB"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db, GetDbErr := getDB.GetDB()
	if GetDbErr != nil {
		fmt.Println("Get DB err")
		return
	}
	router := gin.Default()
	router = Api.APIs(db, router)
	RunServiceErr := router.Run("0.0.0.0:5000")
	if RunServiceErr != nil {
		fmt.Println("Run service err")
		return
	}

}
