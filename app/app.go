package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/controllers"
	"os"
)

func StartApp() {
	router := gin.Default()
	router.GET("/items/:itemID", controllers.GetItem)
	err := router.Run(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error running app", err)
	}

}
