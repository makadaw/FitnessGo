package controllers

import (
	"fitness/db"
	"fitness/lib/sessionauth"
	"fitness/models"
	"fitness/services"
	"github.com/go-martini/martini"
	"net/http"
)

type SessionsController struct {
	Controller
}

func (c SessionsController) Routes() []Route {
	return []Route{
        Route{"GET", "/:id", []martini.Handler{c.Check}},
		Route{"POST", "/create", []martini.Handler{c.Create}},
	}
}

func (s *SessionsController) Check(req *http.Request, r services.Render) {

}

func (s *SessionsController) Create(req *http.Request, r services.Render) {
	user, err := models.FindUserByIdAndPassword(req.FormValue("email"), req.FormValue("password"))

	if err == nil {
		session := models.CreateSessionForUser(user)
		sessionauth.AuthenticateSession(session)
		_, err = db.Upsert(session)
		if err != nil {
			r.Error(err)
		} else {
			r.One(session)
		}
	} else {
		r.Error(err)
	}

}
