
## Setup & Installation (as a contributor, or developer 🚀)
### Tools

* TaskFile
  ```sh
  sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d
  # or, for mac users
  brew install go-task/tap/go-task
  ```
* PreCommit
  ```sh
  brew install pre-commit
  # or, using pip
  pip install pre-commit
  ```
* GoReleaser
  ```sh
  brew install goreleaser/tap/goreleaser
  brew install goreleaser

  ```

### Local development
Most of the common operations are encapsulated using [TaskFile](https://www.taskfile.dev). You can run the following commands to get started:
* Run the CLI in development mode

```yaml
    run-dev:
        dir: app/cli
        cmds:
            - go run main.go {{.CLI_ARGS}}

```
```bash
    task run-dev -- --help
```
* Run the CLI in binary mode

```yaml
    run-binary:
        dir: app/cli
        cmds:
            - go build -o hello-world-go main.go
            - ./hello-world-go {{.CLI_ARGS}}

```
```bash
    task run-binary -- --help
```

* Run the CLI in docker mode

```yaml
    docker-build:
        dir: app/cli
        cmds:
            - docker build -t hello-world-go:latest .

    docker-run:
        dir: app/cli
        cmds:
            - docker run hello-world-go:latest {{.CLI_ARGS}}

```
```bash
    task docker-build
    task docker-run -- --help
```
