package rings

import (
	"github.com/sirupsen/logrus"
)

var R *Ring = New()

type Ring struct {
	Logger *logrus.Logger
}

func New() *Ring {
	return &Ring{
		Logger: logrus.New(),
	}
}

// TODO:
// func (r *Ring) Trace(next http.Handler) http.Handler
// func (r *Ring) Session(next http.Handler) http.Handler
