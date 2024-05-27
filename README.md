`goread` is a command-line interface (CLI) utility written in Go, designed to replicate the functionality of the classic
Unix utility cat.

## How to use

To use `goread`, clone and `cd` into the repository and run the following command to run the provided binary:

```sh
./goread file
```

### Flags

- `-i`: Read from stdin.
- `-n`: Number the lines in the output.

### Examples

#### Read from stdin

To read from stdin, use the `-i` flag:

```sh
head -n2 test.txt| ./goread -i
```

#### Concatenate multiple files

To concatenate multiple files simply pass the desired filenames as arguments:

```sh
./goread test.txt test2.txt
```

#### Numbered lines in output

To number the lines in the output, use the `-n` flag:

```sh
./goread test.txt -n
```
