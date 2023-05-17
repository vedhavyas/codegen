# Codegen

The Codegen is a command-line interface that helps you create a new project based on the provided specification using the power of the AI. It allows you to choose the model, provide a project name, and sync files locally.

Note: The CLI will only get you started with the project with most of the work already done. As a developer, you still need to fix the mistakes in the Code. The more specific specification is, the better the results are.

## Usage

Firstly, install the necessary dependencies and build the binary:

```bash
go build ./cmd/codegen/...
```

Then you can use the CLI as follows:

```bash
./codegen COMMAND [FLAGS]
```

## Commands

### Create

The `create` command is used to create a new project based on the provided specifications.

Flags:

- `--name`: Name of the project. (Required)
- `--specification`: Project specification file. (Default: "specification.md")
- `--model`: GPT model to use. (Default: "gpt-4")
- `--retries`: Number of retries for the model API. (Default: 5)
- `--sync`: Create files locally. (Default: true)
- `--file`: Specify files that are to be created locally instead of all files in the project.

Example:

```bash
./codegen create --name="MyProject" --specification="project_spec.md" --model="gpt-4" --retries=5 --sync
```

This command will create a new project named "MyProject" based on the "project_spec.md" file. It will use the "gpt-4" model with 5 retries, and will sync the generated files locally.

## Configuration

Make sure you have set the `OPEN_AI_API_KEY` with your valid API key. You can set the environment variable as shown below:

```bash
export OPEN_AI_API_KEY="Your-API-Key"
```

After setting up the environment variable, you can run the Codegen CLI tool.
