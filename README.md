# goread

`goread` is a simple command-line interface (CLI) program written in Go.

## Features

- Read input from stdin or files.
- Option to number the lines in the output.

## Usage

### Flags

- `-i`: Read from stdin.
- `-n`: Number the lines in the output.

### Examples

#### Read from stdin

To read from stdin, use the `-i` flag:

```head -n2 test.txt| ./goread -i

```
