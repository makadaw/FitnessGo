package controllers

import (
	"fitness/errors"
	"fitness/lib/sessionauth"
	"fitness/models"
	"fitness/services"
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

type UsersController struct {
	Controller
}

func (c UsersController) Routes() []Route {
	return []Route{
		Route{"GET", "/:id", []martini.Handler{sessionauth.LoginRequired, c.GetUser}},
		Route{"POST", "", []martini.Handler{c.SignUp}},
		Route{"PUT", "", []martini.Handler{sessionauth.LoginRequired, c.UpdateProfile}},
	}
}

func (ctrl UsersController) Index(params martini.Params, r services.Render) {
	users := models.User{}.All()
	r.All(users)
}

func (ctrl UsersController) GetUser(params martini.Params, r services.Render) {
	log.Println("Get user", params["id"])
	r.Error(errors.ApiError{401, "Method not implemented", 0})
}

func (ctrl UsersController) SignUp(req *http.Request, r services.Render) {
	user, err := models.RegisterUser(req.FormValue("email"), req.FormValue("password"), req.FormValue("confirm_password"))
	if err != nil {
		r.Error(err)
	} else {
		r.One(user)
	}
}

func (ctrl UsersController) UpdateProfile(r services.Render) {
	r.Error(errors.ApiError{401, "Method not implemented", 0})
}
