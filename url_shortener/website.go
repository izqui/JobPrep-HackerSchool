package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)

type Website struct{}

func (w *Website) Index(render render.Render) {

	render.HTML(200, "index", nil)
}

func (w *Website) NewURL(link Link, render render.Render, request *http.Request, db linkdb) {

	if key, err := db.Save(link); err == nil {

		render.Redirect(fmt.Sprintf("/link/%s", key))

	} else {

		render.Error(404)
	}
}

func (w *Website) LinkInfo(render render.Render, params martini.Params, db linkdb) {

	var linkParam = params["link"]

	if linkParam != "" {

		link, err := db.Get(linkParam)

		if err == nil {

			url := fmt.Sprintf("%s%s", baseurl, linkParam)
			data := struct {
				Short string
				Link  Link
			}{
				url,
				link,
			}

			render.HTML(200, "link", data)
			return
		}

	}

	render.Error(404)
}

func (w *Website) Link(render render.Render, params martini.Params, db linkdb) {

	var linkParam = params["link"]

	if linkParam != "" {

		link, err := db.Get(linkParam)

		if err == nil {

			render.Redirect(link.Url)
			db.SaveVisit(link)
		}
	}

	render.Error(404)

}
