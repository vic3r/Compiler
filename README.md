# Compiler in Go

Compiler in Go which is intended to make a lexical, syntactic and semantic analysis from a source program written in a kind of C (written in spanish).

## Getting Started

Download Go and install all the required dependeciens which are needed (in some cases). Clone the project and then run the commands specified in the section below.

### Prerequisites

- You need to install Go to run this project

- The binary can be downloaded from: https://golang.org/dl/

- Configure properly your `GOPATH` and `GOROOT`


### Installing
Run the following commands in the root of the project:

- `cd main/`

- `go build -o 'filename'`

- `./filename`

Or simply: 

- `go run cmd/main.go cmd/lexical.go cmd/syntactical.go`

The result of this will be a table of tokens, errors and symbols in case of Lexical analysis

The result of this will be a table of errors and symbols in case of Syntactic analysis

## Running the tests


### And coding style tests


## Versioning

## Authors

- [@vic3r](https://github.com/vic3r)

## Acknowledgments

* Lex & Yacc example of running programs
