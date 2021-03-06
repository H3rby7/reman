package actions

import (
	"github.com/h3rby7/reman/models"
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/mw-csrf"
	"github.com/gobuffalo/mw-i18n"
	"github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/packr/v2"
	"time"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_reman_session",
			//Prefix: "/reman",
		})

		//app.GET("/", func(c buffalo.Context) error {
		//	return c.Redirect(http.StatusSeeOther, "/reman")
		//})

		base := app.Group("/reman")

		// Automatically redirect to SSL
		// app.Use(forceSSL())

		// Log request parameters (filters apply).
		base.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		base.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		base.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		base.Use(translations())

		base.Use(timeNow)

		base.GET("/", HomeHandler)
		eventsResource := EventsResource{}
		base.GET("/events/upcoming", eventsResource.ListUpcoming)
		base.Resource("/events", eventsResource)

		base.ServeFiles("/", assetsBox) // serve files from the public directory
		app.ServeFiles("/", assetsBox)  // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("app:locales", "../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

func timeNow(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Set("now", fmt.Sprintf("%.2d:%.2d", time.Now().UTC().Hour(), time.Now().UTC().Minute()))
		return next(c)
	}
}
