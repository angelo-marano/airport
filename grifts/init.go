package grifts

import (
	"github.com/angelo-marano/airport/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
