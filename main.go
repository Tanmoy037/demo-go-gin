package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/route"
)

var router *gin.Engine

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
			gin.H{
				"title": "Home Page",
			},
		)

	})

	router.Run()

}