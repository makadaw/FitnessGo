package services

import (
	"fitness/errors"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"reflect"
)

/*
 * Service layer on martini render
 */
type Render struct {
	r render.Render
}

func Renderer() martini.Handler {
	return func(c martini.Context, r render.Render) {
		c.Map(Render{r})
		c.Next()
	}
}

func (self Render) One(v interface{}) {
	self.r.JSON(http.StatusOK, v)
}

func (self Render) All(v interface{}) {
	log.Println("JSON array", v)
	if v == nil {
		self.r.JSON(http.StatusOK, []string{"test"})
	} else {
		self.r.JSON(http.StatusOK, v)
	}
}

func (self Render) Error(v error) {
	st := reflect.TypeOf(v)
	tError := reflect.TypeOf((*errors.ApiErrorI)(nil)).Elem()
	if st.Implements(tError) {
		e := v.(errors.ApiError)
		errorsMap := make(map[string][]errors.ApiError)
		errorsMap["errors"] = []errors.ApiError{e}
		self.r.JSON(e.HttpStatus(), errorsMap)
	} else {
		self.r.JSON(http.StatusInternalServerError, v)
	}

}
