package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	// analogous to factory
	router := gin.Default()

	// we added a base router
	api := router.Group("/api")
	{	//'hello' is the endpoint;
		// this sdk takes the refrence to our request (like HttpRequest+HttpResponse in java)
		api.GET("/hello", func(context *gin.Context){
			context.JSON(http.StatusOK, gin.H{
				"message":"hello World", //our actual response
			})
		})
	}
	// expose to this port
	router.Run(":8001")

}

