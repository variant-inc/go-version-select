# go-versions-select

`go-versions-select` is a command-line tool written in Go that filters a list of versions based on a given constraint and returns the newest candidate that matches the constraint.

## Features

- Accepts a list of versions as input.
- Supports version constraint filtering.
- Returns the newest version that satisfies the given constraint.

## Installation

To install the CLI tool, use:

```sh
# Clone the repository (if applicable)
git clone https://github.com/your-repo/go-versions-select.git
cd go-versions-select

# Build the binary
go build -o go-versions-select

# Move the binary to a location in your PATH
mv go-versions-select /usr/local/bin/
```

## Usage

```sh
go-versions-select --versions "1.0.0,1.5.6,2.0.0" --constraint ">=1.0.0"
```

### Flags

- `--versions` (required): A comma-separated list of versions.
- `--constraint` (required): A constraint string that defines the filtering rule.

### Example

#### Input

```sh
go-versions-select --versions "1.0.0,1.5.6,2.0.0" --constraint ">=1.5.0"
```

#### Output

```txt
2.0.0
```
