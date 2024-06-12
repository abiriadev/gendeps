package main

import (
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

var gendepsPlugin = api.Plugin{
	Name: "gendeps",
	Setup: func(build api.PluginBuild) {
		build.OnResolve(
			api.OnResolveOptions{Filter: `^[^./].*$`},
			func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				println(args.Path)

				return api.OnResolveResult{}, nil
			})
	},
}

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"./src/index.ts"},
		Bundle:      true,
		Outfile:     "out.js",
		Plugins:     []api.Plugin{gendepsPlugin},
		Write:       true,
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}
}
