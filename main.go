package main

import (
	"log"
	"os"

	"github.com/anekkanti/go-codegen/pkg"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "go-generate-interface-impl",
		Usage: "Generate implementations for interfaces",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "input-directory",
				Usage:    "The input directory path containing the code to parse",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.StringSliceFlag{
				Name:     "names",
				Usage:    "The names to run the generator for",
				Aliases:  []string{"n"},
				Required: true,
			},
			&cli.PathFlag{
				Name:     "renderer-plugin-path",
				Usage:    "The path to the renderer go plugin to load and use",
				Aliases:  []string{"r"},
				Required: true,
			},
			&cli.PathFlag{
				Name:     "output-path",
				Usage:    "The output directory path to write the generated implementations to",
				Aliases:  []string{"o"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			in := &pkg.Input{
				DirectoryPath: c.String("input-directory"),
				PluginPath:    c.String("renderer-plugin-path"),
				OutputPath:    c.String("output-path"),
				NameFilters:   c.StringSlice("names"),
			}
			return pkg.Generate(in)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
