package codegen

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

func writeToFile(file *File) error {
	dirPath := filepath.Dir(file.Path)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(file.Path, file.Content, 0644)
	if err != nil {
		return err
	}

	return nil
}

type wp struct {
	*Project
	Files map[string]bool `toml:"files"`
}

func SaveProject(project *Project) error {
	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	err := enc.Encode(wp{Project: project, Files: project.files})
	if err != nil {
		return fmt.Errorf("failed to encode project: %v", err)
	}

	return writeToFile(&File{
		Path:    ".codegen/project.toml",
		Content: buf.Bytes(),
	})
}

func LoadProject() (*Project, error) {
	var project wp
	_, err := toml.DecodeFile(".codegen/project.toml", &project)
	if err != nil {
		return nil, fmt.Errorf("failed to decode project")
	}

	files := project.Files
	p := project.Project
	p.files = files
	return p, nil
}
