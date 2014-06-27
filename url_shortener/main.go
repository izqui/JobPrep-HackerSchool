package main

import (
	"os"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

var server *martini.Martini
var baseurl string

func init() {

	server = martini.New()
	baseurl = "http://localhost:8080/"
}

func main() {

	setupServer()
}

func setupServer() {

	os.Setenv("PORT", "8080")

	server.Use(martini.Logger())
	server.Use(render.Renderer())
	server.Use(DB())

	websiteRoutes := new(Website)
	router := martini.NewRouter()

	router.Get("/", websiteRoutes.Index)
	router.Post("/new", binding.Form(Link{}), websiteRoutes.NewURL)
	router.Get("/link/:link", websiteRoutes.LinkInfo)

	router.Get("/:link", websiteRoutes.Link)

	server.Action(router.Handle)
	server.Run()
}
