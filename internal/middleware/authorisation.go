package middleware

import (
	"errors"
	"net/http"

	"github.com/chindi-m/goapi/api"
	"github.com/chindi-m/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorisedError = errors.New("Invalid username or token. ")

func Authorisation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorisation")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorisedError)
			api.RequestErrorHandler(w, UnAuthorisedError)
			return 
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return 
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*&database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (loginDetails).AuthToken)) {
			log.Error(UnAuthorisedError)
			api.RequestErrorHandler(w, UnAuthorisedError)
			return 
		}

		next.ServeHTTP(w, r)
	})
}