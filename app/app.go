package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	"go-api/controllers"
)

func StartApp() {
	// Crea una instancia del enrutador de Gin
	router := gin.Default()

	// Crea un middleware CORS con opciones predeterminadas
	c := cors.Default()

	// Envolvemos el enrutador en el middleware CORS
	handler := c.Handler(router)

	// Agrega tus rutas de API al enrutador aquí
	router.GET("/items/:itemID", controllers.GetItem)

	// Inicie el servidor en el puerto 8000
	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler)
	if err != nil {
		fmt.Println("Error running app", err)
	}
}
