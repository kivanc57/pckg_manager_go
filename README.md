
# Go Modular Project

## Overview

This project provides a modular Go-based solution comprising two command-line tools:

- **Downloader**: Downloads files from specified URLs, manages extraction, and processes packages.
- **Package Manager**: Processes JSON package lists, checks availability in Chocolatey repositories, and logs missing packages.

*The package manager on the new version is implemented both using Go and Python because of maintability. Feel free to use your flavor of choice.*

```

## Installation

Clone this repository and navigate to the project directory:

```bash
git clone https://github.com/your-username/your-project.git
cd your-project
```

Initialize Go modules:

```bash
go mod tidy
```

## Environment Configuration

Create a `.env` file at the root with necessary environment variables:

```env
LINKS=https://example.com/download?token=12345
REPOSITORIES=https://repo.example.com
EXCLUDED_PACKAGES=package1,package2
EXTENSIONS=.zip,.exe,.nupkg
```

## Running the Scripts

### Downloader

To execute the downloader tool:

```bash
go run cmd/downloader/main.go
```

### Package Manager

To execute the package manager:

```bash
go run cmd/packagemanager/main.go
```

## Output

Results will be stored in the `output` directory. Ensure it exists or will be created automatically by the scripts.

## Dependencies

- [godotenv](https://github.com/joho/godotenv) for loading `.env` variables

Install dependencies using:

```bash
go get github.com/joho/godotenv
```
