package codegen

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func SyncProjectFiles(model Model, project *Project, files []string) error {
	logrus.Infoln("Syncing project files locally...")
	group, _ := errgroup.WithContext(context.Background())

	if len(files) < 1 {
		files = project.UnSyncedFiles()
	}
	group.SetLimit(len(files))
	for _, filePath := range files {
		filePath := filePath
		group.Go(func() error {
			err := syncFile(model, project, filePath)
			if err == nil {
				project.FileSynced(filePath)
			}

			return err
		})
	}

	err := group.Wait()
	if err == nil {
		logrus.Infof("Done.")
	}

	return err
}

func syncFile(model Model, project *Project, filePath string) error {
	logrus.Infof("Creating file %s...\n", filePath)
	file, err := getFileContent(model, project, filePath)
	if err != nil {
		return fmt.Errorf("failed to get file content for file[%s]: %v", filePath, err)
	}

	err = writeToFile(file)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

type GetFileContent struct {
	FilePath string
	Metadata string
}

func getFileContent(model Model, project *Project, filePath string) (*File, error) {
	system, err := executeTemplate("existing_project_system.goprompt", project)
	if err != nil {
		return nil, err
	}

	metadata, err := project.MetadataJSON(filePath)
	if err != nil {
		return nil, err
	}

	user, err := executeTemplate("file_content.goprompt", GetFileContent{
		FilePath: filePath,
		Metadata: string(metadata),
	})
	if err != nil {
		return nil, err
	}

	request := openai.ChatCompletionRequest{
		Model: model.Name,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: system,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: user,
			},
		},
		Temperature: 0,
	}

	content, err := model.Respond(request)
	if err != nil {
		return nil, err
	}

	return &File{
		Path:    filePath,
		Content: []byte(content),
	}, nil
}
