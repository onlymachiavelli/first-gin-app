package app 


import (
	"fmt" 
	"net/http" 
	"github.com/gin-gonic/gin"
)
func main () {

	app:= gin.Default() 
	app.GET("/" , func(g * gin.Context) {
		g.JSON(http.StatusOK, gin.H{
			"msg" : "Hello world"
 		})
	} )

	app.Run()
}