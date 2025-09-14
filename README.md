# Binary Tree CLI in Go

This project provides a simple **binary tree implementation in Go**, with a command-line interface (CLI) to insert, search, and manipulate nodes.  
It was created for learning purposes, focusing on data structures and Go’s testing practices.

---

## Features

- Insert nodes into the binary tree
- Prevent duplicate insertions
- Search for nodes
- (Optional) Rebalance the tree
- Simple CLI to interact with the tree

---

## Requirements

- Go 1.23.4 or higher

---

## Installation

Clone this repository and build the project:

```bash
git clone https://github.com/your-username/binary-tree-go.git
cd binary-tree-go
go build ./cmd/binary-tree-go
```

## Usage

Run the CLI:

```bash
go run ./cmd/binary-tree-go
```

Example session:

```bash
Binary Tree CLI. Type 'exit' to quit.
> insert 5
inserted 5
> insert 3
inserted 3
> insert 5
5 is already inserted
> search 3
found 3
> search 10
10 not found
> exit
```

## Testing

Unit tests are written with GO's `testing` package.
You can run all tests with:

```bash
go test ./...
```

## Project Structure

```bash
.
├── cmd/
│   └── binary-tree-go/   # CLI entry point
├── internal/
│   └── tree/             # Binary tree implementation
├── go.mod
└── README.md
```