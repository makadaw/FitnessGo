package models

import (
	"crypto/rand"
	"fitness/db"
	"fitness/lib/sessionauth"
	"labix.org/v2/mgo/bson"
)

type Session struct {
	Id            bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Code          string        `bson:"code" json:"code"`
	UserId        bson.ObjectId `bson:"userId" json:"-"`
	authenticated bool          `bson:"-" json:"-"`
}

// DB settings
func init() {
	db.Register(&Session{})
}
func (u Session) Indexes() [][]string {
	return [][]string{[]string{"_id"}}
}
func (u Session) Collection() string {
	return "sessions"
}
func (u Session) Unique() bson.M {
	return bson.M{"_id": u.Id}
}
func (u Session) PreSave() {
}

// Model methods
func (self Session) All() []Session {
	var sessions []Session = []Session{}
	db.Find(self, bson.M{}).All(&sessions)
	return sessions
}

func (self Session) FindById(id interface{}) error {
	return db.Find(self, bson.M{"id": id}).One(&self)
}

// Session interface
func GenerateAnonymousSession() sessionauth.Session {
	return &Session{}
}

func (self Session) IsAuthenticated() bool {
	return self.authenticated
}

func (self Session) Login() {
	self.authenticated = true
}

func (self Session) Logout() {
	self.authenticated = false
}

func CreateSessionForUser(user *User) Session {
	session := Session{Code: rand_str(30), UserId: user.Id}

	return session
}

func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
