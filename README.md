# go-shell

A POSIX-compliant shell implementation written in Go. This project implements a functional shell capable of interpreting commands, running external programs, and handling builtin commands.

## Features

- Command parsing and execution
- Builtin commands: `cd`, `pwd`, `echo`, `type`, `exit`
- External program execution with PATH resolution
- REPL (Read-Eval-Print Loop) interface

## Requirements

- Go 1.25 or higher

## Usage

Run the shell:
```sh
./your_program.sh
```

Or build and run directly:
```sh
go build -o shell app/main.go
./shell
```

## Project Structure

```
.
├── app/
│   └── main.go          # Main shell implementation
├── .codecrafters/       # CodeCrafters configuration
└── your_program.sh      # Shell runner script
```

## Development

This project was built as part of the [CodeCrafters "Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview).

abc

