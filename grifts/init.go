package grifts

import (
	"github.com/h3rby7/reman/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
