package main

import (
	"codegen"
	"fmt"
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
					},
					&cli.StringFlag{
						Name:  "specification",
						Value: "specification.md",
					},
					&cli.StringFlag{
						Name:  "model",
						Value: "gpt-4",
					},
					&cli.IntFlag{
						Name:  "retries",
						Value: 5,
					},
				},
				Action: func(context *cli.Context) error {
					apiKey := os.Getenv("OPEN_AI_API_KEY")
					model := codegen.NewModel(context.String("model"), context.Int("retries"), apiKey)
					project, err := codegen.CreateProject(model, context.String("name"), context.String("specification"))
					if err != nil {
						return err
					}

					return codegen.SaveProject(project)
				},
			},
			{
				Name:  "sync",
				Usage: "Sync a project files created based on the Specification provided",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "model",
						Value: "gpt-4",
					},
					&cli.IntFlag{
						Name:  "retries",
						Value: 5,
					},
					&cli.StringSliceFlag{Name: "file"},
				},
				Action: func(context *cli.Context) error {
					apiKey := os.Getenv("OPEN_AI_API_KEY")
					model := codegen.NewModel(context.String("model"), context.Int("retries"), apiKey)
					project, err := codegen.LoadProject()
					if err != nil {
						return fmt.Errorf("failed to load project. please run `create` first")
					}
					defer codegen.SaveProject(project)

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
