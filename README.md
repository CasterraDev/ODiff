# ODiff

Outputs the lines that are different from two files. Outputs to `stdout`

# Install

You can install it by using go install:

```go install -v github.com/CasterraDev/odiff@latest```

# Usage

```
> cat list.txt
cats
dogs
cows

> cat map.txt
dogs
cows
horses
chickens

> odiff list.txt map.txt
horses
chickens
cats

```
