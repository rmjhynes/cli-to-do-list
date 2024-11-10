# CLI To Do List in Go
A CLI based to-do list app in Go. This is my first project in Go, and the aim was purely to familiarise myself with the logic and syntax by building a very simple project.

## Installation Instructions
Add the absolute path to your `data.csv` file (stored in /data) to a `.env` file in the root directory. The code is setup to use the environment variable `TASK_DATA_FILE` but this can be changed in `constants/paths.go`.

## Logic
The files in the `cmd` directory register CLI commands using the Cobra package.
The files in the `logic` directory contain the logic behind the CLI commands.


## Run Instructions
- To list the tasks run `go run main.go`
- To add a task to the list run `go run main.go add`
- To mark a task as complete run `go run main.go complete`
