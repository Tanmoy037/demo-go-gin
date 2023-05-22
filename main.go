package main

import (
	"main"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/route"
)

func main(){
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context){
		
		//call the HTML method of the Context to render a template
		c.HTML(
			//set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			//pass the data that the page uses (in the case, 'title')
			git.H{
				"title": "Home Page",
			},
		)






	})

}