package controllers

import (
	"github.com/go-martini/martini"
	"log"
)

type Controller struct {
	ctx *martini.ClassicMartini
}

type ControllerI interface {
	Routes() []Route
	setContext(ctx *martini.ClassicMartini)
}

type Route struct {
	Method   string
	Pattern  string
	Handlers []martini.Handler
}

func (c Controller) Routes() []Route {
	return nil
}

func (c Controller) setContext(ctx *martini.ClassicMartini) {
	c.ctx = ctx
}

func RegisterController(ctr ControllerI, resource string, web *martini.ClassicMartini, h ...martini.Handler) {
	log.Println("Add routing for resource:", resource)
	ctr.setContext(web)
	web.Group("/"+resource, func(r martini.Router) {
		routes := ctr.Routes()
		for _, route := range routes {
			r.AddRoute(route.Method, route.Pattern, route.Handlers...)
		}
	}, h...)
}
