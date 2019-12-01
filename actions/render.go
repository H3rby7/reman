package actions

import (
	"encoding/json"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

var r *render.Engine
var assetsBox = packr.New("app:assets", "../public")

func loadManifest() map[string]string {
	manifest, err := assetsBox.FindString("assets/manifest.json")
	if err != nil {
		App().Logger.Warnf("could not read webpack manifest file %v", err)
		return map[string]string{}
	}
	m := map[string]string{}
	err = json.Unmarshal([]byte(manifest), &m)
	if err != nil {
		panic(err)
	}
	return m
}

func assetPath(file string) string {
	m := loadManifest()
	return m[file]
}

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			"env":        func() string { return app.Env },
			"asset_path": assetPath,
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			// forms.FormKey:     forms.Form,
			// forms.FormForKey:  forms.FormFor,
		},
	})
}
