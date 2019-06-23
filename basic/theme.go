package main

import (
	"fmt"

	"github.com/aeridya/core"
	"github.com/aeridya/theme"
)

type ATheme struct {
	theme.Theme
}

func (t *ATheme) Init(opts ...theme.Option) error {
	t.ParseOpts(opts)
	theme.Register(t)
	return nil
}

func (t *ATheme) Serve(resp *core.Response) {
	if resp.R.URL.Path == "/" {
		resp.Good(200)
		fmt.Fprintf(resp.W, "Hello Aeridya!\n")
		return
	}
	resp.Bad(404, "Basic theme only supports /")
}

func (t *ATheme) Error(resp *core.Response) {
	fmt.Fprintf(resp.W, "Error: %d\n%s\n", resp.Status, resp.Err)
	return
}
