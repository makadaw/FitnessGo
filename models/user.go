package models

import (
	"fitness/db"
	"fitness/errors"
	"golang.org/x/crypto/bcrypt"
	"labix.org/v2/mgo/bson"
	"regexp"
	"time"
)

type User struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username  string        `bson:"username" json:"username"`
	Password  []byte        `bson:"password" json:"-"`
	Email     string        `bson:"email,omitempty" json:"email"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}

// DB settings
func init() {
	db.Register(&User{})
}
func (u User) Indexes() [][]string {
	return [][]string{[]string{"email"}}
}
func (u User) Collection() string {
	return "users"
}
func (u User) Unique() bson.M {
	return bson.M{"email": u.Email}
}
func (u User) PreSave() {
}

// Model methods
func (self User) All() []User {
	var users []User = []User{}
	db.Find(self, bson.M{}).All(&users)
	return users
}

func (self User) FindById(id string) User {
	return User{}
}

func (self User) Validate() (bool, error) {
	//check email
	regex, _ := regexp.Compile("(\\w[-._\\w]*\\w@\\w[-._\\w]*\\w\\.\\w{2,3})")
	if !regex.MatchString(self.Email) {
		return false, errors.InvalidEmail
	}
	return true, nil
}

// Password
func (u *User) hashPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	u.Password = hashedPassword
}

//
func RegisterUser(email, password, confirm_password string) (*User, error) {
	user := User{Email: email, CreatedAt: time.Now()}
	if password != confirm_password {
		return nil, errors.PasswordNotEqualConfirmPassword
	} else {
		user.hashPassword(password)
		valid, err := user.Validate()
		if valid {
			if db.Exists(user) {
				return nil, errors.EmailExist
			}
			_, err = db.Upsert(user)
			if err != nil {
				return nil, errors.EmailExist
			}
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func FindUserByIdAndPassword(email string, password string) (User, error) {
	user := User{}
	db.Find(user, bson.M{"email": email}).One(user)
	if len(user.Email) > 0 {
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
		if err != nil {
			return nil, errors.BadUsernameOrPassword
		}
		return user, nil
	} else {
		return nil, errors.BadUsernameOrPassword
	}
}
