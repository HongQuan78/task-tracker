# Task Tracker

Task Tracker is a small CLI for managing tasks directly in your terminal. Tasks are stored in a local JSON file, so there’s no database or external dependency.

## Requirements
- Go 1.25+ (per `go.mod`)
- OS permissions to write `internal/storage/data.json`

## Quick start
1) Clone the repo.
2) Run commands directly: `go run . --help`
3) Build a binary: `go build -o task.exe .` and run `./task.exe ...`

## Commands
- Add a task: `go run . add "Write README"`
- List tasks: `go run . list` or filter by status `go run . list -s in-progress`
- Update description: `go run . update 3 -d "Polish docs"`
- Change status: `go run . status 3 --set-status done` (`todo | in-progress | done`)
- Delete a task: `go run . delete 3`

## Data storage
- All tasks live in `internal/storage/data.json`
- Each task: `Id`, `Description`, `Status` (`todo | in-progress | done`), `CreatedAt`, `UpdatedAt`
- Back up this file to move data between machines

## Dev notes
- `cmd/` holds Cobra commands; `internal/` has models and storage
- `main.go` calls `cmd.Execute()` to register commands

## Tips
- Get help: `go run . --help` or `go run . <command> --help`
- Reset data by deleting `internal/storage/data.json` (back it up first if needed)
