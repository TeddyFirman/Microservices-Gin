package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	// router.GET("/", func(c *gin.Context) {

	// 	//Call the HTML method of the context to render a template
	// 	c.HTML(
	// 		//set the http status to 200 (ok)
	// 		http.StatusOK,
	// 		//use the index.html template
	// 		"index.html",
	// 		//pass the data that page uses (in this case, 'title')
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	//Handle Index
	router.GET("/", showIndexPage)

	//Hanle GET requests at /article/view.some_article_id
	router.GET("/article/view/:article_id", getArticle)

	router.Run()

	//Start serving the apps
	router.Run()
}

