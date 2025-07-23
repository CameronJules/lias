# Lias-CLI

An easy-to-use command-line tool to manage your aliases

## ✨ Features

- 💾 Save frequently used functions as aliases
- 📁  Create projects to group aliases
- 🧰 Works on macOS and Linux


## 🚀 Installation

### Option 1: Using `go install`
```
git clone https://github.com/CameronJules/lias.git
cd lias
go install
export $PATH
```

## 🔧 Commands

| Description                        | Command                                                      |
|------------------------------------|--------------------------------------------------------------|
| Create a project                   | `lias project add [project name]`                            |
| List projects                      | `lias project list`                                         |
| Delete a project                   | `lias project delete [project name]`                         |
| Create a function in a project     | `lias function add [project name] [alias] ["command to run"]`|
| List functions in a project        | `lias function list [project name]`                         |
| Delete a function                  | `lias function delete [project name] [function name]`        |
| Run an alias                       | `lias run [project name] [alias]`                            |





