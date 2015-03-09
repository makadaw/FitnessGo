package main

import (
	"fitness/conf"
	"fitness/controllers"
	"fitness/db"
	"fitness/lib/sessionauth"
	"fitness/models"
	"fitness/services"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	server := martini.Classic()

	server.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	server.Use(render.Renderer(render.Options{
		IndentJSON: false,
	}))
	server.Use(services.Renderer())

	initDB()
	initSession(server)
	initAuth(server)
	initRoutes(server)

	log.Println("Start server on", conf.Config.HostString())
	server.RunOnAddr(conf.Config.HostString())
}

func initSession(server *martini.ClassicMartini) {
	store := sessions.NewCookieStore([]byte(conf.Config.SessionSecret))
	server.Use(sessions.Sessions("my_session", store))
}

func initAuth(server *martini.ClassicMartini) {
	server.Use(sessionauth.AuthSession(models.GenerateAnonymousSession))
}

func initDB() {
	db.Connect(conf.Config.DbHostString(), conf.Config.DbName)
	db.RegisterAllIndexes()
}

func initRoutes(server *martini.ClassicMartini) {
	server.Get("/", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"error": "Please use API methods"})
	})
	controllers.RegisterController(controllers.SessionsController{}, "sessions", server)
	controllers.RegisterController(controllers.UsersController{}, "users", server)
}
