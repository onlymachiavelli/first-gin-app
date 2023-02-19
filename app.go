package main 


import (
	"fmt" 
	"net/http" 
	"github.com/gin-gonic/gin"
	"onlymachiavelli/fga/utils"
)
func main () {

	app:= gin.Default() 
	app.GET("/" , func(g * gin.Context) {
		g.JSON(http.StatusOK, gin.H{
			"msg" : "Hello world" , 
 		})
	} )
		//connect
		db := utils.Connect()
		fmt.Println(db)
		
	

	
	app.Run()
}