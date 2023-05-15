package codegen

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"text/template"
)

type Project struct {
	Name          string `toml:"name"`
	Specification string `toml:"specification"`
	mu            sync.RWMutex
	files         map[string]bool `toml:"files"`
}

func (p *Project) Files() []string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	var list []string
	for k := range p.files {
		list = append(list, k)
	}

	return list
}

func (p *Project) UnSyncedFiles() []string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	var list []string
	for k, v := range p.files {
		if v {
			continue
		}

		list = append(list, k)
	}

	return list
}

func (p *Project) FileSynced(filePath string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.files[filePath] = true
}

func (p *Project) IsFileSynced(filePath string) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.files[filePath]
}

type File struct {
	Path    string
	Content []byte
}

//go:embed prompts
var prompts embed.FS
var templates template.Template

func init() {
	dir, err := prompts.ReadDir("prompts")
	if err != nil {
		panic(err)
	}

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		data, err := prompts.ReadFile(fmt.Sprintf("prompts/%s", entry.Name()))
		if err != nil {
			panic(err)
		}

		tmpl := templates.New(name)
		_, err = tmpl.Parse(string(data))
		if err != nil {
			panic(err)
		}
	}
}

func executeTemplate(name string, data any) (string, error) {
	buf := &bytes.Buffer{}
	err := templates.ExecuteTemplate(buf, name, data)
	return buf.String(), err
}

func CreateProject(model Model, name string, specificationFile string) (*Project, error) {
	logrus.Infof("Creating project %s...\n", name)
	specification, err := os.ReadFile(specificationFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read specification: %v", err)
	}

	project := &Project{
		Name:          name,
		Specification: string(specification),
	}

	filesPaths, err := getProjectFiles(model, project)
	if err != nil {
		return nil, fmt.Errorf("failed to generate project file list: %v", err)
	}

	project.files = filesPaths
	logrus.Infof("Done.")
	return project, nil
}

func getProjectFiles(model Model, project *Project) (map[string]bool, error) {
	system, err := executeTemplate("create_project_structure_system.goprompt", project)
	if err != nil {
		return nil, err
	}

	user, err := executeTemplate("list_files.prompt", nil)
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
		Temperature: 0.3,
	}

	var response struct {
		Files []string `json:"files"`
	}

	err = model.RespondJSON(request, &response)
	if err != nil {
		return nil, err
	}

	files := make(map[string]bool)
	for _, file := range response.Files {
		files[file] = false
	}
	return files, nil
}