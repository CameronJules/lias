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

## Commands

| Description                        | Command                                                      |
|------------------------------------|--------------------------------------------------------------|
| Create a project                   | `lias add project [project name]`                            |
| Create a function in a project     | `lias add function [project name] [alias] ["command to run"]`|
| List projects                      | `lias list projects`                                         |
| List functions in a project        | `lias list functions [project name]`                         |
| Delete a project                   | `lias delete project [project name]`                         |
| Delete a function                  | `lias delete function [project name] [function name]`        |
| Run an alias                       | `lias run [project name] [alias]`                            |





