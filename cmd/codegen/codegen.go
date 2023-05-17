package main

import (
	"codegen"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.App{
		Name:  "codegen",
		Usage: "AI Developer",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Creates a new project based on the Specification provided",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Required: true,
						Usage:    "Name of the project.",
					},
					&cli.StringFlag{
						Name:  "specification",
						Value: "specification.md",
						Usage: "Project specification file.",
					},
					&cli.StringFlag{
						Name:  "model",
						Value: "gpt-4",
						Usage: "GPT model to use",
					},
					&cli.IntFlag{
						Name:  "retries",
						Value: 5,
					},
					&cli.BoolFlag{
						Name:  "sync",
						Value: true,
						Usage: "Create files locally.",
					},
					&cli.StringSliceFlag{Name: "file"},
				},
				Action: func(context *cli.Context) error {
					apiKey := os.Getenv("OPEN_AI_API_KEY")
					model := codegen.NewModel(context.String("model"), context.Int("retries"), apiKey)
					project, err := codegen.CreateProject(model, context.String("name"), context.String("specification"))
					if err != nil {
						return err
					}
					defer codegen.SaveProject(project)

					if !context.Bool("sync") {
						return nil
					}

					return codegen.SyncProjectFiles(model, project, context.StringSlice("file"))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
