package routers

import (
	"net/http"

	"./controllers"
	"./middlewares"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

// registers all routes for the application.
func GetRouter() *pat.Router {
	// url paths imported from urls package
	url_patterns := ReturnURLS()
	medium := pat.New()
	medium.Get(url_patterns.HOME_PATH, controllers.HomeController)
	common := pat.New()
	// static route
	common.PathPrefix(url_patterns.STATIC_PATH).Handler(
		http.StripPrefix(url_patterns.STATIC_PATH, http.FileServer(http.Dir("static"))))
	// applying middlewares
	common.PathPrefix(url_patterns.MEDIUM_PATH).Handler(
		negroni.New(
			negroni.HandlerFunc(
				middlewares.LoggingMiddleware),
			negroni.Wrap(medium),
		),
	)
	common.Get(url_patterns.LOGIN_PATH, controllers.LoginController)
	common.Post(url_patterns.LOGIN_PATH, controllers.LoginController)
	return common
}
