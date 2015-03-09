package sessionauth

import (
	"github.com/go-martini/martini"
	//"github.com/martini-contrib/render"
	"log"
	"net/http"
)

var (
	SessionKey       string = "TEST"
	AuthorizationKey string = "Authorization"
)

type Session interface {
	IsAuthenticated() bool

	Login()
	Logout()
	FindById(id interface{}) error
}

// newSession - function return default anon session
func AuthSession(newSession func() Session) martini.Handler {
	return func(c martini.Context, req *http.Request) {
		authHeader := req.Header.Get(AuthorizationKey)
		sessionId := authKeyFromHeader(authHeader)
		session := newSession()

		if len(sessionId) > 0 {
			err := session.FindById(sessionId)
			if err == nil {
				log.Println("Loaded session", session)
			}
		}

		c.MapTo(session, (*Session)(nil))
	}
}

func authKeyFromHeader(header string) string {
	return header
}

func AuthenticateSession(s Session) {
	s.Login()
}

func Logout(s Session) {

}

func LoginRequired(session Session, w http.ResponseWriter, req *http.Request) {
	if session.IsAuthenticated() == false {
		http.Redirect(w, req, "/login", 301)
	}
}
