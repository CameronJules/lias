# Lias-CLI

An easy-to-use command-line tool to manage your aliases.

This project was created to learn Golang, all critiques are welcome.

## ✨ Features

- 💾 Save frequently used functions as aliases
- 📁  Create projects to group aliases
- 🧰 Works on macOS and Linux


## 🚀 Installation

### Option 1: Using `go install`
```
git clone https://github.com/CameronJules/lias.git
cd lias
go build
go install
export PATH=$HOME/go/bin:$PATH
```

## 🔧 Commands

| Description                        | Command                                                      |
|------------------------------------|--------------------------------------------------------------|
| Create a project                   | `lias project add [project name]`                            |
| List projects                      | `lias project list`                                         |
| Delete a project                   | `lias project delete [project name]`                         |
| Use a project to store and run commands from    | `lias use [project name]`|
| View the current active project   | `lias active`|
| Create a function in the active project     | `lias function add [alias] ["command to run"]`|
| List functions in the active project        | `lias function list`                         |
| Delete a function from the active project                | `lias function delete [function name]`        |
| Run an alias from the active project                       | `lias run [alias]`                            |





